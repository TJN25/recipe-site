package data

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
	"clove":    {Type: TypeCount, FactorToStdBase: 0},
	"unit":     {Type: TypeCount, FactorToStdBase: 0},
	"slice":    {Type: TypeCount, FactorToStdBase: 0},
	"head":     {Type: TypeCount, FactorToStdBase: 0},
	"sprig":    {Type: TypeCount, FactorToStdBase: 0},
	"bunch":    {Type: TypeCount, FactorToStdBase: 0},
	"can":      {Type: TypeCount, FactorToStdBase: 0},
	"pinch":    {Type: TypeDescriptive, FactorToStdBase: 0},
	"dash":     {Type: TypeDescriptive, FactorToStdBase: 0},
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

// --- Helper to select the right ruleset ---
// func SelectFormattingRules(unitType UnitType, system UnitSystem) FormattingRuleSet { ... }
