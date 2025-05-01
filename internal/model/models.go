package model

type CleaningDifficulty int

const (
	CleaningDifficultyNone   CleaningDifficulty = iota // or near none e.g. paper towel, baking paper, teaspoon for tasting etc.
	CleaningDifficultyEasy                             // something that goes in the dishwasher or sink and requires little space e.g. Spatula
	CleaningDifficultyMedium                           // something that requires some attention, or is a bit bigger e.g. chef's knife, small chopping board
	CleaningDifficultyHard                             // requires care, has things baked on, is very large, can't be cleaned while cooking e.g. caserole dish
)

type RecipeStep struct {
	ID       int64 `json:"id"`        // Will be DB ID later
	RecipeID int64 `json:"recipe_id"` // can be used in multiple recipes
	// RecipeIDs   []int64            `json:"recipe_ids"` // can be used in multiple recipes
	StepOrder   int                `json:"step_order"` // We will just determine this by it's position in the array. Allows it to be moved around without needing to renumber
	Title       string             `json:"title"`      // e.g., "Prepare Chicken", "Make Sauce"
	Description string             `json:"description,omitempty"`
	Notes       string             `json:"notes,omitempty"`
	Ingredients []RecipeIngredient `json:"ingredients"`  // Ingredients *specific* to this step
	MethodSteps []MethodStep       `json:"method_steps"` // Sub-steps *within* this step
	Equipment   []Equipment        `json:"equipment"`    // Equipment *specific* to this step
	// Servings    int                `json:"servings"`     // Define the scaling used in the section. We may need to ensure this is consistent across steps
	// Tags        []Tag           `json:"tags"`        // We need this for enabling searching for steps as though they are recipes e.g. bao bun chicken, roux for mac and cheese, mexican chicken marinade.
}

type RecipeIngredient struct {
	FoodItem FoodItem `json:"food_item"` // The generic food item details (Name, Nutrition etc.) fetched via join

	// Fields specific to the recipe usage (from recipe_ingredients join table)
	Quantity   float32 `json:"quantity"`          // How much of the item is needed for the recipe servings
	Unit       string  `json:"unit"`              // Unit for the quantity (e.g., "cup", "tbsp", "g", "clove")
	IsOptional bool    `json:"is_optional"`       // Is this ingredient optional for the recipe?
	Purpose    string  `json:"purpose,omitempty"` // Why this ingredient is used (e.g., "thickener", "acidity")
}

type FoodItem struct {
	ID                      int64     `json:"id"`
	Name                    string    `json:"name"`                // Unique name, e.g., "All-Purpose Flour"
	BaseUnit                string    `json:"base_unit"`           // Canonical unit for price/nutrition (e.g., "g", "ml", "whole")
	PricePerBaseUnit        float32   `json:"price_per_base_unit"` // Price for one BaseUnit
	Nutrition               Nutrition `json:"nutrition"`           // Embedded nutrition info per BaseUnit
	AlternateUnits          []string  `json:"alternate_units"`
	AlternateUnitConversion []float32 `json:"alternate_unit_conversion"`
	// Potential future fields: category, brand, substitutes[]
}

type FoodItemGeneric struct {
	ID   int64  `json:"id"`   // Database primary key for the generic item
	Name string `json:"name"` // The common name (e.g., "Garlic") - Should be unique
	// Potential future fields: DefaultBaseUnit? Category ("Spice", "Vegetable", "Dairy")?
}

type Nutrition struct {
	CarbsGrams   float32 `json:"carbs_grams"`
	FatGrams     float32 `json:"fat_grams"`
	ProteinGrams float32 `json:"protein_grams"`
	CaloriesKcal float32 `json:"calories_kcal"`
}

type MethodStep struct {
	ID              int64  `json:"id"`
	RecipeID        int64  `json:"-"`                 // Foreign key back to Recipe (db only)
	StepNumber      int    `json:"step_number"`       // Order of the step (unique per recipe)
	Instruction     string `json:"instruction"`       // What to do in this step
	PrepTimeMinutes int    `json:"prep_time_minutes"` // Active time for this step
	CookTimeMinutes int    `json:"cook_time_minutes"` // Passive/cooking time for this step
	// Potential V2: IngredientsUsed []int64, EquipmentUsed []int64
}

type Equipment struct {
	ID                 int64              `json:"id"`
	Name               string             `json:"name"` // Unique name, e.g., "Large Saucepan"
	Type               string             `json:"type"` // e.g., "Cookware", "Utensil", "Appliance"
	CleaningDifficulty CleaningDifficulty `json:"cleaning_difficulty"`
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"` // e.g., "Indian", "Quick", "Winter", "Pasta Bake"
	Type string `json:"type"` // e.g., "cuisine", "duration", "season", "dish_type", "technique", "mood"
}

type Recipe struct {
	ID          int64  `json:"id"` // Database primary key
	Title       string `json:"title"`
	Description string `json:"description"`
	Servings    int    `json:"servings"`        // Base servings for the whole recipe
	Notes       string `json:"notes,omitempty"` // Overall notes
	ImagePath   string `json:"image_path,omitempty"`

	// Relationships / Components
	RecipeSteps []RecipeStep `json:"recipe_steps"` // The ordered list of stages/steps
	Tags        []Tag        `json:"tags"`         // Overall tags for the recipe

	// --- Optional: Add aggregated list for convenience? ---
	// AllIngredients []RecipeIngredient `json:"-"` // Could be calculated, not stored directly

	// Calculated/Derived fields
	TotalPrepTimeMinutes int       `json:"total_prep_time_minutes"`
	TotalCookTimeMinutes int       `json:"total_cook_time_minutes"`
	TotalCleaningScore   int       `json:"total_cleaning_score"`
	CalculatedNutrition  Nutrition `json:"calculated_nutrition"`
	CalculatedPrice      float32   `json:"calculated_price"`
}

// Helper function to calculate total nutrition for a recipe (example)
func (r *Recipe) CalculateTotals() {
	// Reset calculated fields
	r.TotalPrepTimeMinutes = 0
	r.TotalCookTimeMinutes = 0
	r.TotalCleaningScore = 0
	r.CalculatedPrice = 0.0
	r.CalculatedNutrition = Nutrition{} // Zero out nutrition

	// Use maps to track unique equipment and aggregate ingredients
	uniqueEquipment := make(map[int64]Equipment) // Assuming Equipment has an ID
	allIngredients := []RecipeIngredient{}       // To calculate total nutrition/price

	// Iterate through each major step/stage (RecipeStep)
	for _, recipeStep := range r.RecipeSteps {
		// Aggregate ingredients from this step
		allIngredients = append(allIngredients, recipeStep.Ingredients...)

		// Sum time from sub-steps (MethodStep)
		for _, methodStep := range recipeStep.MethodSteps {
			r.TotalPrepTimeMinutes += methodStep.PrepTimeMinutes
			r.TotalCookTimeMinutes += methodStep.CookTimeMinutes
		}

		// Track unique equipment used in this step
		// Note: Assumes Equipment struct has an ID field populated.
		// If not, use Name as the map key.
		for _, eq := range recipeStep.Equipment {
			if eq.ID != 0 { // Use ID if available
				uniqueEquipment[eq.ID] = eq
			} else if eq.Name != "" { // Fallback to Name if ID is 0 (dummy data)
				// Need a way to handle equipment without ID/Name uniquely if possible
				// For now, just add based on Name if ID is missing
				key := int64(0) // Or generate a temporary key/hash if needed
				if _, exists := uniqueEquipment[key]; !exists || uniqueEquipment[key].Name != eq.Name {
					// Crude check, assumes name is unique if ID is 0
					uniqueEquipment[key] = eq
					key-- // Ensure next potential 0-ID item gets a different temp key
				}
			}
		}
	}

	// Calculate cleaning score from unique equipment
	for _, eq := range uniqueEquipment {
		r.TotalCleaningScore += int(eq.CleaningDifficulty)
	}

	// --- Calculate price and nutrition from aggregated ingredients ---
	var totalNutrition Nutrition
	var totalPrice float32
	// IMPORTANT: This aggregation assumes quantities are directly comparable
	//            WITHOUT UNIT CONVERSION, which is INCORRECT for real use.
	//            We need the FoodItem lookup and conversion logic here eventually.
	for _, ri := range allIngredients {
		if ri.FoodItem.ID == 0 && ri.FoodItem.Name == "" {
			continue
		} // Skip empty items

		// !!! --- Placeholder --- !!! Needs unit conversion
		factor := ri.Quantity
		if ri.FoodItem.BaseUnit != "" && ri.Unit == ri.FoodItem.BaseUnit { // Basic check
			totalPrice += ri.FoodItem.PricePerBaseUnit * factor
			totalNutrition.CarbsGrams += ri.FoodItem.Nutrition.CarbsGrams * factor
			totalNutrition.FatGrams += ri.FoodItem.Nutrition.FatGrams * factor
			totalNutrition.ProteinGrams += ri.FoodItem.Nutrition.ProteinGrams * factor
			totalNutrition.CaloriesKcal += ri.FoodItem.Nutrition.CaloriesKcal * factor
		} else {
			// TODO: Implement unit conversion logic here
			// log.Printf("Warning: Unit conversion needed for %s (%s vs %s)", ri.FoodItem.Name, ri.Unit, ri.FoodItem.BaseUnit)
		}
		// !!! --- End Placeholder --- !!!
	}

	// Calculate per serving
	if r.Servings > 0 {
		r.CalculatedPrice = totalPrice / float32(r.Servings)
		r.CalculatedNutrition.CarbsGrams = totalNutrition.CarbsGrams / float32(r.Servings)
		r.CalculatedNutrition.FatGrams = totalNutrition.FatGrams / float32(r.Servings)
		r.CalculatedNutrition.ProteinGrams = totalNutrition.ProteinGrams / float32(r.Servings)
		r.CalculatedNutrition.CaloriesKcal = totalNutrition.CaloriesKcal / float32(r.Servings)
	}

	// Optional: Store the aggregated list if needed elsewhere
	// r.AllIngredients = allIngredients
}
