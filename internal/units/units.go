package units

import (
	"errors"
	"fmt"
	"html/template" // Needed for the return type of the main formatter
	"math"
	"strconv"
	"strings"

	"github.com/TJN25/recipe-site/internal/model"
	"github.com/TJN25/recipe-site/internal/store"
	log "github.com/sirupsen/logrus"
)

type UnitType string

const (
	TypeWeight      UnitType = "weight"
	TypeVolume      UnitType = "volume"
	TypeCount       UnitType = "count"       // Represents countable items where conversion depends on the item (e.g., clove, unit)
	TypeDescriptive UnitType = "descriptive" // Non-convertible units (e.g., pinch, to taste)
)

type UnitSystem string

const (
	SystemMetric UnitSystem = "metric"
	SystemUS     UnitSystem = "us_customary"
	SystemOther  UnitSystem = "other" // For count/descriptive
)

type UnitDefinition struct {
	Type            UnitType // Is it weight, volume, count, or descriptive?
	FactorToStdBase float64  // Factor to convert 1 of THIS unit TO its standard base (g or ml). 0 if not applicable (count/desc).
	// System          UnitSystem // Optional: Keep if useful for selecting rules later
}

var UnitDefinitions = map[string]UnitDefinition{
	// Volume -> ml
	"ml":    {Type: TypeVolume, FactorToStdBase: 1.0},
	"l":     {Type: TypeVolume, FactorToStdBase: 1000.0},
	"tsp":   {Type: TypeVolume, FactorToStdBase: 4.92892},
	"tbsp":  {Type: TypeVolume, FactorToStdBase: 14.7868},
	"fl oz": {Type: TypeVolume, FactorToStdBase: 29.5735},
	"cup":   {Type: TypeVolume, FactorToStdBase: 236.588},
	"pt":    {Type: TypeVolume, FactorToStdBase: 473.176},
	"qt":    {Type: TypeVolume, FactorToStdBase: 946.353},
	"gal":   {Type: TypeVolume, FactorToStdBase: 3785.41},
	// Weight -> g
	"g":  {Type: TypeWeight, FactorToStdBase: 1.0},
	"kg": {Type: TypeWeight, FactorToStdBase: 1000.0},
	"oz": {Type: TypeWeight, FactorToStdBase: 28.3495},
	"lb": {Type: TypeWeight, FactorToStdBase: 453.592},
	// Count/Desc -> Base conversion is item-specific or N/A
	"cm":       {Type: TypeCount, FactorToStdBase: 0},
	"inch":     {Type: TypeCount, FactorToStdBase: 0},
	"clove":    {Type: TypeCount, FactorToStdBase: 0},
	"unit":     {Type: TypeCount, FactorToStdBase: 0},
	"slice":    {Type: TypeCount, FactorToStdBase: 0},
	"head":     {Type: TypeCount, FactorToStdBase: 0},
	"leaf":     {Type: TypeCount, FactorToStdBase: 0},
	"piece":    {Type: TypeCount, FactorToStdBase: 0},
	"sprig":    {Type: TypeCount, FactorToStdBase: 0},
	"bunch":    {Type: TypeCount, FactorToStdBase: 0},
	"can":      {Type: TypeCount, FactorToStdBase: 0},
	"stalk":    {Type: TypeCount, FactorToStdBase: 0},
	"pinch":    {Type: TypeDescriptive, FactorToStdBase: 0},
	"dash":     {Type: TypeDescriptive, FactorToStdBase: 0},
	"garnish":  {Type: TypeDescriptive, FactorToStdBase: 0},
	"to taste": {Type: TypeDescriptive, FactorToStdBase: 0},
}

var UnitNormalizationMap = map[string]string{
	// Volume - Metric
	"ml":          "ml",
	"milliliter":  "ml",
	"millilitre":  "ml",
	"milliliters": "ml",
	"millilitres": "ml",
	"l":           "l", // Use 'l' for litre internally? Or L? Consistency needed. Let's use 'l'.
	"litre":       "l",
	"liter":       "l",
	"litres":      "l",
	"liters":      "l",

	// Volume - US Customary
	"tsp":          "tsp",
	"t":            "tsp", // Common abbreviation
	"teaspoon":     "tsp",
	"teaspoons":    "tsp",
	"tbsp":         "tbsp",
	"tbs":          "tbsp", // Common abbreviation
	"tablespoon":   "tbsp",
	"tablespoons":  "tbsp",
	"fl oz":        "fl oz", // Keep space for clarity? Or fl_oz? Let's use space.
	"floz":         "fl oz",
	"fluid ounce":  "fl oz",
	"fluid ounces": "fl oz",
	"cup":          "cup",
	"c":            "cup",
	"cups":         "cup",
	"pt":           "pt",
	"pint":         "pt",
	"pints":        "pt",
	"qt":           "qt",
	"quart":        "qt",
	"quarts":       "qt",
	"gal":          "gal",
	"gallon":       "gal",
	"gallons":      "gal",

	// Weight - Metric
	"g":         "g",
	"gram":      "g",
	"grams":     "g",
	"kg":        "kg",
	"kilogram":  "kg",
	"kilograms": "kg",

	// Weight - Imperial
	"oz":     "oz",
	"ounce":  "oz",
	"ounces": "oz",
	"lb":     "lb",
	"pound":  "lb",
	"pounds": "lb",

	// Count / Descriptive
	"clove":    "clove",
	"cloves":   "clove",
	"unit":     "unit", // For things like '1 egg', '1 onion'
	"units":    "unit",
	"slice":    "slice",
	"slices":   "slice",
	"head":     "head",
	"heads":    "head",
	"sprig":    "sprig",
	"sprigs":   "sprig",
	"bunch":    "bunch",
	"bunches":  "bunch",
	"can":      "can",
	"cans":     "can",
	"pinch":    "pinch", // Descriptive, non-convertible
	"pinches":  "pinch",
	"dash":     "dash", // Descriptive, non-convertible
	"dashes":   "dash",
	"to taste": "to taste", // Descriptive, non-convertible
}

// FormattingRule defines how to display a value within a specific range.
type FormattingRule struct {
	Threshold      float64        // Apply this rule if value_in_canonical_base >= Threshold
	DisplayUnit    string         // The unit to display (e.g., "kg", "lb", "cup", "tsp")
	FactorFromBase float64        // Factor to convert FROM canonical base TO DisplayUnit (e.g., for kg: 0.001; for cup: 1.0/236.588)
	RoundingRule   RoundingDetail // Specific rounding details for this range/unit
}

// RoundingDetail specifies how to round the value IN THE DISPLAY UNIT.
type RoundingDetail struct {
	Type      string  // "decimal", "fraction", "nearest_int", "nearest_multiple"
	Precision float64 // For "decimal": number of places; For "fraction": denominator; For "nearest_multiple": the multiple (e.g., 5 for rounding to nearest 5g)
}

// FormattingRuleSet holds the ordered rules for a specific Type + System.
type FormattingRuleSet []FormattingRule

// --- Define the actual Rule Sets ---

// MetricWeightRules (Input is grams)
var MetricWeightRules = FormattingRuleSet{
	{Threshold: 1000.0, DisplayUnit: "kg", FactorFromBase: 0.001, RoundingRule: RoundingDetail{Type: "decimal", Precision: 1}},       // 1.2 kg
	{Threshold: 500.0, DisplayUnit: "g", FactorFromBase: 1.0, RoundingRule: RoundingDetail{Type: "nearest_multiple", Precision: 50}}, // 550 g
	{Threshold: 100.0, DisplayUnit: "g", FactorFromBase: 1.0, RoundingRule: RoundingDetail{Type: "nearest_multiple", Precision: 10}}, // 110 g
	{Threshold: 20.0, DisplayUnit: "g", FactorFromBase: 1.0, RoundingRule: RoundingDetail{Type: "nearest_multiple", Precision: 5}},   // 25 g
	{Threshold: 5.0, DisplayUnit: "g", FactorFromBase: 1.0, RoundingRule: RoundingDetail{Type: "nearest_int"}},                       // 6 g
	{Threshold: 0.0, DisplayUnit: "g", FactorFromBase: 1.0, RoundingRule: RoundingDetail{Type: "decimal", Precision: 1}},             // 0.5 g (Small amounts)
}

// USWeightRules (Input is grams)
var USWeightRules = FormattingRuleSet{
	// TBD: Define thresholds/units (lb, oz) and rounding (e.g., 1/4 lb, oz, 1/8 oz?)
	// Example structure:
	{Threshold: 453.592, DisplayUnit: "lb", FactorFromBase: 1.0 / 453.592, RoundingRule: RoundingDetail{Type: "fraction", Precision: 4}}, // Round to 1/4 lb > 1lb
	{Threshold: 28.3495 * 4, DisplayUnit: "oz", FactorFromBase: 1.0 / 28.3495, RoundingRule: RoundingDetail{Type: "nearest_int"}},        // Round to nearest oz > 4oz?
	{Threshold: 0.0, DisplayUnit: "oz", FactorFromBase: 1.0 / 28.3495, RoundingRule: RoundingDetail{Type: "fraction", Precision: 8}},     // Round to 1/8 oz < 4oz?
}

// MetricVolumeRules (Input is ml)
var MetricVolumeRules = FormattingRuleSet{
	{Threshold: 1000.0, DisplayUnit: "l", FactorFromBase: 0.001, RoundingRule: RoundingDetail{Type: "decimal", Precision: 1}},         // 1.5 L
	{Threshold: 500.0, DisplayUnit: "ml", FactorFromBase: 1.0, RoundingRule: RoundingDetail{Type: "nearest_multiple", Precision: 50}}, // 550 ml
	{Threshold: 100.0, DisplayUnit: "ml", FactorFromBase: 1.0, RoundingRule: RoundingDetail{Type: "nearest_multiple", Precision: 5}},  // 105 ml
	{Threshold: 20.0, DisplayUnit: "ml", FactorFromBase: 1.0, RoundingRule: RoundingDetail{Type: "nearest_int"}},                      // 21 ml
	{Threshold: 0.0, DisplayUnit: "ml", FactorFromBase: 1.0, RoundingRule: RoundingDetail{Type: "decimal", Precision: 1}},             // 5.5 ml (Maybe nearest 5ml?)
}

// USVolumeRules (Input is ml) - Implementing your ranges
var USVolumeRules = FormattingRuleSet{
	// Note: Factors convert FROM ml TO DisplayUnit
	// Thresholds based on approx ml values for boundaries
	{Threshold: 250.0, DisplayUnit: "cup", FactorFromBase: 1.0 / 236.588, RoundingRule: RoundingDetail{Type: "fraction", Precision: 4}}, // >= 250ml (~1+ cup), round 1/4 cup
	{Threshold: 60.0, DisplayUnit: "cup", FactorFromBase: 1.0 / 236.588, RoundingRule: RoundingDetail{Type: "fraction", Precision: 8}},  // 60-249ml (~1/4-1 cup), round 1/8 cup? Or to nearest 2 tbsp? Let's try 1/8 cup.
	{Threshold: 15.0, DisplayUnit: "tbsp", FactorFromBase: 1.0 / 14.7868, RoundingRule: RoundingDetail{Type: "fraction", Precision: 2}}, // 15-59ml (~1-4 tbsp), round 1/2 tbsp? (nearest tsp = Precision:3?) Let's try 1/2 tbsp.
	{Threshold: 0.0, DisplayUnit: "tsp", FactorFromBase: 1.0 / 4.92892, RoundingRule: RoundingDetail{Type: "fraction", Precision: 8}},   // < 15ml (~< 3 tsp), round 1/8 tsp
}

var CountRules = FormattingRuleSet{
	{Threshold: 50, DisplayUnit: "unit", RoundingRule: RoundingDetail{Type: "nearest_multiple", Precision: 10}},
	{Threshold: 20, DisplayUnit: "unit", FactorFromBase: 1.0 / 236.588, RoundingRule: RoundingDetail{Type: "nearest_multiple", Precision: 5}},
	{Threshold: 5, DisplayUnit: "unit", RoundingRule: RoundingDetail{Type: "nearest_int"}},
	{Threshold: 3, DisplayUnit: "unit", RoundingRule: RoundingDetail{Type: "fraction", Precision: 2}},
	{Threshold: 0.0, DisplayUnit: "unit", FactorFromBase: 1.0 / 4.92892, RoundingRule: RoundingDetail{Type: "fraction", Precision: 4}},
}

func FormatIngredientForDisplay(
	ingredient model.RecipeIngredient,
	stepBaseServings int,
	targetServings int,
	displaySystemKey string, // e.g., "use_original", "use_metric", "use_us_customary"
) (template.HTML, error) {
	log.Debugf("-> FormatIngredient Entry: ItemID=%d, Form='%s', Qty=%.2f, Unit='%s', BaseServ=%d, TargetServ=%d, System='%s'",
		ingredient.FoodItemID, ingredient.FormName, ingredient.Quantity, ingredient.SpecifiedUnit, stepBaseServings, targetServings, displaySystemKey)

	var finalQuantity float64
	var displayUnit string
	var roundingRule RoundingDetail

	log.Debugf("   SpecifiedUnit to normalize: '%s'", ingredient.SpecifiedUnit)
	normalizedUnit, _ := normalizeUnit(ingredient.SpecifiedUnit)
	log.Debugf("   Normalized unit result: '%s'", normalizedUnit)
	unitDef, ok := UnitDefinitions[normalizedUnit]
	log.Debugf("   Looked up UnitDefinition for '%s'. Found: %t. Type: %s", normalizedUnit, ok, unitDef.Type)
	if ok && unitDef.Type == TypeDescriptive {
		log.Debugf("   Handling as Descriptive Unit.")
		roundingRule = RoundingDetail{Type: "decimal", Precision: 1}
		log.Debugf("   Calling formatFinalOutput (Descriptive): Qty=%.2f, Unit='%s', Name='%s', Rule={%s, %.1f}",
			float64(ingredient.Quantity), normalizedUnit, ingredient.FormName, roundingRule.Type, roundingRule.Precision)
		formattedString := formatFinalOutput(float64(ingredient.Quantity), normalizedUnit, ingredient.FormName, roundingRule) // Pass 0 or actual quantity? Handle in formatter.
		log.Debugf("<- FormatIngredient Exit (Descriptive): Result='%s'", formattedString)
		return template.HTML(formattedString), nil
	}

	log.Debugf("   Scaling quantity: BaseQty=%.2f, BaseServ=%d, TargetServ=%d", ingredient.Quantity, stepBaseServings, targetServings)
	scaledQuantity, err := scaleQuantity(ingredient.Quantity, stepBaseServings, targetServings)
	if err != nil {
		log.Errorf("Error scaling quantity for ingredient %d: %v", ingredient.FoodItemID, err)
		scaledQuantity = ingredient.Quantity // Use original as fallback
	}
	log.Debugf("   Scaled quantity result: %.4f", scaledQuantity)

	// 3. Handle "use_original" display system
	switch displaySystemKey {
	case "use_original":
		log.Debugf("   Handling 'use_original' system.")
		log.Debugf("   SpecifiedUnit for 'use_original': '%s'", ingredient.SpecifiedUnit)
		normUnit, ok := normalizeUnit(ingredient.SpecifiedUnit)
		if !ok {
			log.Warnf("Unknown original unit '%s' for ingredient %d", ingredient.SpecifiedUnit, ingredient.FoodItemID)
			normUnit = ingredient.SpecifiedUnit
		}
		log.Debugf("   Normalized unit for 'use_original': '%s'", normUnit)

		finalQuantity = float64(scaledQuantity)
		displayUnit = normUnit
		roundingRule = RoundingDetail{Type: "decimal", Precision: 1}
		log.Debugf("   'use_original' - finalQuantity=%.4f, displayUnit='%s', roundingRule={%s, %.1f}",
			finalQuantity, displayUnit, roundingRule.Type, roundingRule.Precision)

	case "use_scales_metric":
		log.Debugf("   Handling 'use_scales_metric' system.")

		//Fetch foodItem so we can look up conversions
		foodItem, ok := store.AllFoodItemsCache[ingredient.FoodItemID]
		if !ok {
			errMsg := fmt.Sprintf("Missing %d for %s in AllFoodItemsCache", ingredient.FoodItemID, ingredient.FoodItemName)
			log.Errorf("  %s", errMsg)
			return template.HTML(""), errors.New(errMsg) // Cannot proceed without form details
		}

		// Get details for the ingredient's specific form
		normalisedTargetFormName := store.NormalizeName(ingredient.FormName)
		lookupFormName, _ := store.NormalizedFormNameLookup[normalisedTargetFormName]
		formDetails, formFound := foodItem.Forms[lookupFormName]
		if !formFound {
			errMsg := fmt.Sprintf("form '%s' not found for FoodItem %d ('%s')", ingredient.FormName, foodItem.ID, foodItem.Name)
			log.Errorf("   %s", errMsg)
			return template.HTML(""), errors.New(errMsg) // Cannot proceed without form details
		}

		// Check if the form's defined Unit has a ToGrams factor
		normUnit, _ := normalizeUnit(ingredient.SpecifiedUnit)
		if formDetails.ToGrams <= 0 {
			// Cannot convert this form to grams, fallback to original display
			log.Warnf("   Cannot convert Form '%s' (Unit '%s') to grams for ItemID %d. Missing ToGrams factor > 0. Falling back to original.",
				ingredient.FormName, formDetails.Unit, ingredient.FoodItemID)
			finalQuantity = float64(scaledQuantity)
			displayUnit = normUnit
			roundingRule = RoundingDetail{Type: "decimal", Precision: 1} // Basic fallback rule
		} else {
			// Convert scaledQuantity (in SpecifiedUnit) to the form's native Unit first (if different)
			qtyInFormUnit := float64(scaledQuantity)
			if normUnit != formDetails.Unit {
				// Need to convert specified unit to form unit via canonical FIRST
				// This requires converting both specified and form unit to g/ml etc.
				// Let's simplify: assume for now SpecifiedUnit *matches* formDetails.Unit
				// OR that ToGrams is grams per SpecifiedUnit (needs clarification in model/data)
				// For now, assume scaledQuantity IS in formDetails.Unit
				// TODO: Revisit this conversion if SpecifiedUnit can differ from FormDetails.Unit
				log.Warnf("   Assuming scaledQuantity %.4f is already in form unit '%s' for conversion.", scaledQuantity, formDetails.Unit)
			}

			// Convert quantity (in form's unit) to grams using the form's factor
			canonicalValueGrams := qtyInFormUnit * float64(formDetails.ToGrams)
			log.Debugf("   Converted %.4f %s to %.4f grams", qtyInFormUnit, formDetails.Unit, canonicalValueGrams)

			// Format the gram value using metric weight rules (finds g/kg, applies rounding rule)
			var formatErr error
			finalQuantity, displayUnit, roundingRule, formatErr = FormatWeight(canonicalValueGrams, string(SystemMetric)) // Use "metric" key
			if formatErr != nil {
				log.Errorf("   Error formatting weight for ingredient %d: %v. Falling back to original.", ingredient.FoodItemID, formatErr)
				// Fallback on formatting error
				finalQuantity = float64(scaledQuantity)
				displayUnit = normUnit
				roundingRule = RoundingDetail{Type: "decimal", Precision: 1}
			}
			log.Debugf("   Formatted weight output pre-final: Qty=%.4f, Unit='%s', Rule={%s, %.1f}",
				finalQuantity, displayUnit, roundingRule.Type, roundingRule.Precision)
		}
	default:
		log.Debugf("   Handling '%s' system.", displaySystemKey)

		// 4a. Normalize the *original* unit & get its definition
		log.Debugf("   SpecifiedUnit for conversion: '%s'", ingredient.SpecifiedUnit)
		normUnit, ok := normalizeUnit(ingredient.SpecifiedUnit)
		if !ok {
			errMsg := fmt.Sprintf("unknown unit '%s' for ingredient %d", ingredient.SpecifiedUnit, ingredient.FoodItemID)
			log.Errorf("   %s", errMsg)
			log.Debugf("<- FormatIngredient Exit (Error)")
			return template.HTML(""), fmt.Errorf("unknown unit '%s' for ingredient %d", ingredient.SpecifiedUnit, ingredient.FoodItemID)
		}
		log.Debugf("   Normalized unit for conversion: '%s'", normUnit)

		unitDef, ok := UnitDefinitions[normUnit]
		if !ok {
			errMsg := fmt.Sprintf("no definition found for normalized unit '%s'", normUnit)
			log.Errorf("   %s", errMsg)
			log.Debugf("<- FormatIngredient Exit (Error)")
			return template.HTML(""), fmt.Errorf("no definition found for normalized unit '%s'", normUnit)
		}
		log.Debugf("   Unit definition found: Type='%s', FactorToStdBase=%.4f", unitDef.Type, unitDef.FactorToStdBase)

		// Check if conversion is possible (must have a factor to base)
		var formatErr error
		if unitDef.FactorToStdBase == 0 {
			log.Debugf("Unit '%s' (type %s) is of type count '%s'.", normUnit, unitDef.Type, displaySystemKey)
			finalQuantity = float64(scaledQuantity)
			displayUnit = normUnit
			roundingRule = RoundingDetail{Type: "decimal", Precision: 1} // Default basic rounding
			finalQuantity, displayUnit, roundingRule, formatErr = FormatCount(finalQuantity, displaySystemKey)
			log.Debugf("   FormatWeight/Volume result: finalQuantity=%.4f, displayUnit='%s', roundingRule={%s, %.1f}, err=%v",
				finalQuantity, displayUnit, roundingRule.Type, roundingRule.Precision, formatErr)

			if formatErr != nil {
				log.Errorf("Error formatting value for ingredient %d: %v", ingredient.FoodItemID, formatErr)
				finalQuantity = float64(scaledQuantity)
				displayUnit = normUnit
				roundingRule = RoundingDetail{Type: "decimal", Precision: 1} // Default basic rounding
				log.Debugf("   Fallback after format error - finalQuantity=%.4f, displayUnit='%s', roundingRule={%s, %.1f}",
					finalQuantity, displayUnit, roundingRule.Type, roundingRule.Precision)
			}

		} else {
			// 4b. Convert scaled quantity to canonical base (g or ml)
			canonicalValue := float64(scaledQuantity) * unitDef.FactorToStdBase
			var baseUnitType UnitType = unitDef.Type // Should be weight or volume
			log.Debugf("   Calculated canonical value: %.4f (%s)", canonicalValue, baseUnitType)

			// 4c. Determine Display Unit, Unrounded Quantity, and Rounding Rule using FormatWeight/Volume
			if baseUnitType == TypeWeight {
				log.Debugf("   Calling FormatWeight: baseGrams=%.4f, system='%s'", canonicalValue, displaySystemKey)
				finalQuantity, displayUnit, roundingRule, formatErr = FormatWeight(canonicalValue, displaySystemKey)
			} else if baseUnitType == TypeVolume {
				log.Debugf("   Calling FormatVolume: baseMl=%.4f, system='%s'", canonicalValue, displaySystemKey)
				finalQuantity, displayUnit, roundingRule, formatErr = FormatVolume(canonicalValue, displaySystemKey)
			} else {
				formatErr = fmt.Errorf("unexpected unit type '%s' for conversion", baseUnitType)
			}
			log.Debugf("   FormatWeight/Volume result: finalQuantity=%.4f, displayUnit='%s', roundingRule={%s, %.1f}, err=%v",
				finalQuantity, displayUnit, roundingRule.Type, roundingRule.Precision, formatErr)

			if formatErr != nil {
				log.Errorf("Error formatting value for ingredient %d: %v", ingredient.FoodItemID, formatErr)
				finalQuantity = float64(scaledQuantity)
				displayUnit = normUnit
				roundingRule = RoundingDetail{Type: "decimal", Precision: 1} // Default basic rounding
				log.Debugf("   Fallback after format error - finalQuantity=%.4f, displayUnit='%s', roundingRule={%s, %.1f}",
					finalQuantity, displayUnit, roundingRule.Type, roundingRule.Precision)
			}
		}
	}
	log.Debugf("   Calling formatFinalOutput: Qty=%.4f, Unit='%s', Name='%s', Rule={%s, %.1f}",
		finalQuantity, displayUnit, ingredient.FormName, roundingRule.Type, roundingRule.Precision)

	formattedString := formatFinalOutput(finalQuantity, displayUnit, ingredient.FormName, roundingRule)
	log.Debugf("   formatFinalOutput result: '%s'", formattedString)

	log.Debugf("<- FormatIngredient Exit (Success)")

	return template.HTML(formattedString), nil
}

func scaleQuantity(baseQuantity float32, baseServings int, targetServings int) (float32, error) {
	if baseServings <= 0 || targetServings <= 0 {
		return baseQuantity, fmt.Errorf("invalid servings (base: %d, target: %d)", baseServings, targetServings)
	}
	if baseServings == targetServings {
		return baseQuantity, nil
	}
	scaleFactor := float32(targetServings) / float32(baseServings)
	return baseQuantity * scaleFactor, nil
}

func normalizeUnit(input string) (string, bool) {
	norm, found := UnitNormalizationMap[strings.ToLower(strings.TrimSpace(input))]
	return norm, found
}

func selectWeightFormattingRules(system string) (FormattingRuleSet, error) {
	switch system {
	case string(SystemMetric): // Use defined constants
		return MetricWeightRules, nil
	case string(SystemUS): // Use defined constants
		return USWeightRules, nil
	default:
		return MetricWeightRules, nil
	}
}

func selectCountFormattingRules() FormattingRuleSet {
	return CountRules
}

func FormatCount(inputDisplayQty float64, system string) (displayQty float64, displayUnit string, roundingRule RoundingDetail, err error) {
	rules := selectCountFormattingRules()

	// Ensure rules are ordered high threshold to low in consts.go
	for _, rule := range rules {
		// Use a small tolerance for float comparisons if necessary, though >= should be okay
		if inputDisplayQty >= rule.Threshold {
			// Found the correct rule for this magnitude
			if rule.FactorFromBase == 0 {
				// Avoid division by zero issues if factor is zero (shouldn't happen for weight/volume)
				return 0, "", rule.RoundingRule, fmt.Errorf("invalid zero FactorFromBase for rule threshold %.2f, unit %s", rule.Threshold, rule.DisplayUnit)
			}
			displayUnit = rule.DisplayUnit
			roundingRule = rule.RoundingRule
			return inputDisplayQty, displayUnit, roundingRule, nil // Found match, return
		}
	}

	err = fmt.Errorf("no suitable formatting rule found for weight %.2f g in system '%s'", inputDisplayQty, system)
	return 0, "", RoundingDetail{}, err
}

func FormatWeight(baseGrams float64, system string) (displayQty float64, displayUnit string, roundingRule RoundingDetail, err error) {
	rules, err := selectWeightFormattingRules(system)
	if err != nil {
		return 0, "", RoundingDetail{}, err // Return error if system is invalid
	}

	// Ensure rules are ordered high threshold to low in consts.go
	for _, rule := range rules {
		// Use a small tolerance for float comparisons if necessary, though >= should be okay
		if baseGrams >= rule.Threshold {
			// Found the correct rule for this magnitude
			if rule.FactorFromBase == 0 {
				// Avoid division by zero issues if factor is zero (shouldn't happen for weight/volume)
				return 0, "", rule.RoundingRule, fmt.Errorf("invalid zero FactorFromBase for rule threshold %.2f, unit %s", rule.Threshold, rule.DisplayUnit)
			}
			displayQty = baseGrams * rule.FactorFromBase // Convert grams TO display unit
			displayUnit = rule.DisplayUnit
			roundingRule = rule.RoundingRule
			return displayQty, displayUnit, roundingRule, nil // Found match, return
		}
	}

	err = fmt.Errorf("no suitable formatting rule found for weight %.2f g in system '%s'", baseGrams, system)
	return 0, "", RoundingDetail{}, err
}

func selectVolumeFormattingRules(system string) (FormattingRuleSet, error) {
	switch system {
	case string(SystemMetric):
		return MetricVolumeRules, nil
	case string(SystemUS):
		return USVolumeRules, nil
	default:
		return USVolumeRules, nil
	}
}

// FormatVolume determines the best display unit, unrounded quantity, and rounding rule for a given volume.
func FormatVolume(baseMl float64, system string) (displayQty float64, displayUnit string, roundingRule RoundingDetail, err error) {

	rules, err := selectVolumeFormattingRules(system)
	if err != nil {
		return 0, "", RoundingDetail{}, err // Return error if system is invalid
	}

	// Ensure rules are ordered high threshold to low in consts.go
	for _, rule := range rules {
		// Use a small tolerance for float comparisons if necessary
		if baseMl >= rule.Threshold {
			// Found the correct rule for this magnitude
			if rule.FactorFromBase == 0 {
				// This might happen if converting TO ml/l itself
				if rule.DisplayUnit == "ml" || rule.DisplayUnit == "l" {
					// If DisplayUnit is ml or l, FactorFromBase might be derived differently or handled specially
					// Let's assume FactorFromBase is correctly set (e.g., 1.0 for ml, 0.001 for l)
					// This error check might be less critical here if data is set up correctly.
				} else {
					return 0, "", rule.RoundingRule, fmt.Errorf("invalid zero FactorFromBase for volume rule threshold %.2f, unit %s", rule.Threshold, rule.DisplayUnit)
				}
			}

			// Calculate the quantity in the rule's DisplayUnit
			// Example: baseMl = 500, rule is for "pt" (FactorFromBase = 1.0 / 473.176)
			// displayQty = 500 * (1.0 / 473.176) = ~1.05 pints
			displayQty = baseMl * rule.FactorFromBase

			// Get the target display unit and the rounding rule
			displayUnit = rule.DisplayUnit
			roundingRule = rule.RoundingRule

			// Return the results - WE FOUND THE MATCH, STOP ITERATING
			return displayQty, displayUnit, roundingRule, nil
		}
	}

	// Should not be reached if rules include threshold 0.0
	err = fmt.Errorf("no suitable formatting rule found for volume %.2f ml in system '%s'", baseMl, system)
	return 0, "", RoundingDetail{}, err
}

func unitSpace(unit string) string {
	// Map of units that should NOT have a preceding space (case-sensitive - assuming normalized input)
	noSpaceUnits := map[string]bool{
		"g":  true,
		"kg": true, // Added kg
		"ml": true,
		"l":  true, // Added l
		"%":  true,
		"°C": true,
		"°F": true,
	}
	// Check normalized unit (already done before calling this ideally)
	normalizedUnit := strings.ToLower(unit)
	if _, found := noSpaceUnits[normalizedUnit]; found || normalizedUnit == "" {
		return "" // No space needed
	}
	return " " // Add a space for others (tsp, tbsp, cup, oz, lb, clove, unit, etc.)
}

// pluralize handles basic pluralization for common units.
func pluralize(quantity float64, unit string) string {
	// Using a tolerance for floating point comparisons to 1
	if math.Abs(quantity-1.0) < 0.001 {
		return unit // Return singular if quantity is effectively 1
	}

	// units like ml and g are not pluralized
	if len(unit) < 3 {
		return unit
	}

	// Handle specific irregulars or known non-pluralized units first if any
	switch unit {
	case "to taste", "pinch", "dash": // Don't pluralize descriptive
		return unit
	}

	// Basic pluralization rules (can be expanded)
	// Check for units already ending in 's' - simplistic check
	if strings.HasSuffix(unit, "s") {
		return unit
	}
	// Simple rule: add 's' - covers most common cases like cup, tsp, tbsp, clove, etc.
	// More complex rules for words ending in y, ch, sh, x, z could be added if needed.
	return unit + "s"
}

func FormatFraction(value float64, denominator int) string {
	if denominator <= 0 {
		// Fallback to decimal if denominator is invalid
		return strconv.FormatFloat(value, 'f', -1, 64)
	}

	// tolerance := 1.0 / float64(denominator*2) // Tolerance for rounding

	wholePart := int(math.Floor(value))
	fractionalPart := value - float64(wholePart)

	// Round fractional part to nearest fraction based on denominator
	numerator := int(math.Round(fractionalPart * float64(denominator)))

	// Handle rounding up to the next whole number
	if numerator == denominator {
		wholePart++
		numerator = 0
	}

	// Simplify fraction (find greatest common divisor)
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	if numerator != 0 {
		commonDivisor := gcd(numerator, denominator)
		numerator /= commonDivisor
		denominator /= commonDivisor
	}

	// Build the string
	result := ""
	if wholePart > 0 {
		result += strconv.Itoa(wholePart)
	}

	if numerator > 0 {
		if wholePart > 0 {
			result += " " // Add space between whole and fraction
		}
		result += fmt.Sprintf("%d/%d", numerator, denominator)
	}

	if result == "" { // Handle case where value was 0 or rounded to 0
		return "0"
	}

	return result
}

// Helper for applying rounding rules (used within formatFinalOutput)
func applyRounding(qty float64, rule RoundingDetail) float64 {
	if qty == 0 {
		return 0
	} // Avoid issues with rounding zero

	switch rule.Type {
	case "decimal":
		precision := int(rule.Precision)
		if precision < 0 {
			precision = 0
		}
		factor := math.Pow10(precision)
		return math.Round(qty*factor) / factor
	case "fraction":
		return qty // Let FormatFraction handle it
	case "nearest_int":
		return math.Round(qty)
	case "nearest_multiple":
		multiple := rule.Precision
		if multiple <= 0 {
			return qty
		} // Avoid division by zero or no-op
		return math.Round(qty/multiple) * multiple
	default:
		return math.Round(qty*10) / 10 // Round to 1dp as a fallback
	}
}

func formatFinalOutput(unroundedQty float64, displayUnit string, itemName string, rule RoundingDetail) string {

	// --- 1. Handle non-numeric / descriptive units first ---
	switch displayUnit {
	case "to taste":
		// For "to taste", typically omit quantity entirely
		return fmt.Sprintf("%s %s", displayUnit, itemName)
	case "pinch", "dash":
		// For pinch/dash, often use "a" or "1" unless specifically fractional
		qtyStr := "A" // Default to "a"
		if math.Abs(unroundedQty-1.0) > 0.001 && unroundedQty > 0 {
			// If explicitly not 1 (e.g., 0.5 or 2), format it
			// Use FormatFraction for things like 1/2 pinch? Precision 2 or 4.
			qtyStr = FormatFraction(unroundedQty, 2) // Example: format to nearest 1/2
		}
		// Use singular form for "a"
		unitStr := displayUnit
		if qtyStr == "A" {
			// No pluralization needed
		} else {
			// Need to parse qtyStr back to float for pluralize if it became "1/2" etc.
			// Simplification: just use original unroundedQty for plural check here
			unitStr = pluralize(unroundedQty, displayUnit)
		}
		return fmt.Sprintf("%s%s%s %s", qtyStr, unitSpace(unitStr), unitStr, itemName)
	case "unit":
		unitDef, ok := UnitDefinitions[displayUnit] // Assuming displayUnit is normalized
		if !ok {
			qtyStr := formatQuantityFloat(unroundedQty) // Basic format
			return fmt.Sprintf("%s %s", qtyStr, itemName)
		}

		var qtyStr string
		roundedQty := applyRounding(unroundedQty, rule) // Apply rounding determined by FormatWeight/Volume rule

		denominator := 4
		if rule.Type == "fraction" && rule.Precision > 0 {
			denominator = int(rule.Precision)
		}
		// Only use fraction formatting if the rounding type suggests it might be useful
		// or if the unit itself implies it (like 'unit' for onion/carrot)
		if rule.Type == "fraction" { // Or check specific units: unit, clove?
			qtyStr = FormatFraction(roundedQty, denominator)
		} else {
			// Otherwise, format counts as integers or maybe 1 decimal place
			qtyStr = formatQuantityFloat(roundedQty) // Default to decimal/int format
		}
		pluralUnit := pluralize(roundedQty, displayUnit)

		unitDef, ok = UnitDefinitions[displayUnit]

		if ok && unitDef.Type == TypeCount {
			displayName := itemName
			singularUnit := strings.TrimSuffix(pluralUnit, "s")

			if singularUnit != "" && strings.Contains(strings.ToLower(displayName), strings.ToLower(singularUnit)) {
				splitName := strings.Split(displayName, " ")
				modifiedDisplayName := ""
				count := 0
				for _, word := range splitName {
					if strings.Contains(strings.ToLower(word), strings.ToLower(singularUnit)) {
						continue
					}
					if count > 0 {
						modifiedDisplayName += " "
					}
					modifiedDisplayName += word
					count += 1
				}
				return fmt.Sprintf(" %s %s", qtyStr, modifiedDisplayName)
			}
		}

		// --- 5. Combine ---
		return fmt.Sprintf("%s %s", qtyStr, itemName)

	case "":
		// Handle empty unit (error condition)
		if unroundedQty != 0 {
			qtyStr := formatQuantityFloat(unroundedQty) // Basic format
			return fmt.Sprintf("%s %s ???", qtyStr, itemName)
		}
		return itemName // Just return name if quantity is also zero
	}

	// --- 2. Look up fundamental unit type ---
	unitDef, ok := UnitDefinitions[displayUnit] // Assuming displayUnit is normalized
	if !ok {
		// Fallback if the unit determined by FormatWeight/Volume isn't in UnitDefinitions (shouldn't happen)
		log.Warnf("formatFinalOutput: Unit '%s' not found in UnitDefinitions. Using basic formatting.", displayUnit)
		qtyStr := formatQuantityFloat(unroundedQty) // Basic fallback
		pluralUnit := pluralize(unroundedQty, displayUnit)
		return fmt.Sprintf("%s%s%s %s", qtyStr, unitSpace(pluralUnit), pluralUnit, itemName)
	}

	// --- 3. Apply Rounding and Formatting based on Type ---
	var qtyStr string
	roundedQty := applyRounding(unroundedQty, rule) // Apply rounding determined by FormatWeight/Volume rule

	switch unitDef.Type {
	case TypeCount:
		// Specific formatting for counts (e.g., 1/2 onion, 2 cloves)
		// Often use fractions for halves/quarters. Use rule's precision if fraction, else basic format.
		// Let's default count fractions to quarters if rule allows.
		denominator := 4
		if rule.Type == "fraction" && rule.Precision > 0 {
			denominator = int(rule.Precision)
		}
		// Only use fraction formatting if the rounding type suggests it might be useful
		// or if the unit itself implies it (like 'unit' for onion/carrot)
		if rule.Type == "fraction" { // Or check specific units: unit, clove?
			qtyStr = FormatFraction(roundedQty, denominator)
		} else {
			// Otherwise, format counts as integers or maybe 1 decimal place
			qtyStr = formatQuantityFloat(roundedQty) // Default to decimal/int format
		}

	case TypeVolume:
		// Check if the rule specifies fraction (typical for US units)
		if rule.Type == "fraction" && rule.Precision > 0 {
			qtyStr = FormatFraction(roundedQty, int(rule.Precision))
		} else {
			// Otherwise use standard decimal formatting (typical for metric)
			qtyStr = formatQuantityFloat(roundedQty)
		}

	case TypeWeight:
		// Weights usually use decimal formatting
		qtyStr = formatQuantityFloat(roundedQty)

	default: // Should not happen
		qtyStr = formatQuantityFloat(roundedQty)
	}

	pluralUnit := pluralize(roundedQty, displayUnit)

	unitDef, ok = UnitDefinitions[displayUnit]

	if ok && unitDef.Type == TypeCount {
		displayName := itemName
		singularUnit := strings.TrimSuffix(pluralUnit, "s")

		if singularUnit != "" && strings.Contains(strings.ToLower(displayName), strings.ToLower(singularUnit)) {
			splitName := strings.Split(displayName, " ")
			modifiedDisplayName := ""
			count := 0
			for _, word := range splitName {
				if strings.Contains(strings.ToLower(word), strings.ToLower(singularUnit)) {
					continue
				}
				if count > 0 {
					modifiedDisplayName += " "
				}
				modifiedDisplayName += word
				count += 1
			}
			return fmt.Sprintf("%s %s of %s", qtyStr, pluralUnit, modifiedDisplayName)
		}
	}

	// --- 5. Combine ---
	return fmt.Sprintf("%s%s%s %s", qtyStr, unitSpace(pluralUnit), pluralUnit, itemName)
}

// formatQuantityFloat - adapt your original formatQuantity for float64
func formatQuantityFloat(q float64) string {
	// Check if the number is effectively an integer
	if math.Abs(q-math.Round(q)) < 0.001 {
		return fmt.Sprintf("%.0f", q)
	}
	// Format to 2dp and trim
	sFixed := fmt.Sprintf("%.2f", q)
	if strings.HasSuffix(sFixed, ".00") {
		return sFixed[:len(sFixed)-3]
	}
	if strings.HasSuffix(sFixed, "0") {
		return sFixed[:len(sFixed)-1]
	}
	return sFixed
}
