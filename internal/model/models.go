package model

type CleaningDifficulty int

const (
	CleaningDifficultyNone   CleaningDifficulty = iota // or near none e.g. paper towel, baking paper, teaspoon for tasting etc.
	CleaningDifficultyEasy                             // something that goes in the dishwasher or sink and requires little space e.g. Spatula
	CleaningDifficultyMedium                           // something that requires some attention, or is a bit bigger e.g. chef's knife, small chopping board
	CleaningDifficultyHard                             // requires care, has things baked on, is very large, can't be cleaned while cooking e.g. caserole dish
)

type Nutrition struct {
	CarbsGrams   float32 `json:"carbs_grams"`
	FatGrams     float32 `json:"fat_grams"`
	ProteinGrams float32 `json:"protein_grams"`
	CaloriesKcal float32 `json:"calories_kcal"`
}

type FoodItem struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`                // Unique name, e.g., "All-Purpose Flour"
	BaseUnit         string    `json:"base_unit"`           // Canonical unit for price/nutrition (e.g., "g", "ml", "whole")
	PricePerBaseUnit float32   `json:"price_per_base_unit"` // Price for one BaseUnit
	Nutrition        Nutrition `json:"nutrition"`           // Embedded nutrition info per BaseUnit
	// Potential future fields: category, brand, substitutes[]
}

type RecipeIngredient struct {
	FoodItem FoodItem `json:"food_item"` // The generic food item details (Name, Nutrition etc.) fetched via join

	// Fields specific to the recipe usage (from recipe_ingredients join table)
	Quantity   float32 `json:"quantity"`          // How much of the item is needed for the recipe servings
	Unit       string  `json:"unit"`              // Unit for the quantity (e.g., "cup", "tbsp", "g", "clove")
	IsOptional bool    `json:"is_optional"`       // Is this ingredient optional for the recipe?
	Purpose    string  `json:"purpose,omitempty"` // Why this ingredient is used (e.g., "thickener", "acidity")
}

type MethodStep struct {
	ID              int64  `json:"id"`
	RecipeID        int64  `json:"-"`                 // Foreign key back to Recipe (db only)
	StepNumber      int    `json:"step_number"`       // Order of the step (unique per recipe)
	Instruction     string `json:"instruction"`       // What to do in this step
	Stage           string `json:"stage,omitempty"`   // Optional grouping (e.g., "Prep", "Make Sauce", "Assembly")
	PrepTimeMinutes int    `json:"prep_time_minutes"` // Active time for this step
	CookTimeMinutes int    `json:"cook_time_minutes"` // Passive/cooking time for this step
	// Potential V2: IngredientsUsed []int64, EquipmentUsed []int64
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"` // e.g., "Indian", "Quick", "Winter", "Pasta Bake"
	Type string `json:"type"` // e.g., "cuisine", "duration", "season", "dish_type", "technique", "mood"
}

type Equipment struct {
	ID                 int64              `json:"id"`
	Name               string             `json:"name"` // Unique name, e.g., "Large Saucepan"
	Type               string             `json:"type"` // e.g., "Cookware", "Utensil", "Appliance"
	CleaningDifficulty CleaningDifficulty `json:"cleaning_difficulty"`
}

type Recipe struct {
	ID          int64 // Database primary key
	Title       string
	Description string
	Servings    int
	Notes       string
	ImagePath   string
	// Relationships (populated by joining data from other tables)
	Ingredients []RecipeIngredient // Ingredients needed for the recipe
	MethodSteps []MethodStep       // Steps to make the recipe
	Equipment   []Equipment        // Unique equipment needed
	Tags        []Tag              // Tags associated
	// Calculated/Derived fields (usually not stored in DB)
	TotalPrepTimeMinutes int
	TotalCookTimeMinutes int
	TotalCleaningScore   int
	CalculatedNutrition  Nutrition `json:"calculated_nutrition"` // Estimated nutrition per serving
	CalculatedPrice      float32   `json:"calculated_price"`     // Estimated price per serving
}

// Helper function to calculate total nutrition for a recipe (example)
func (r *Recipe) CalculateTotals() {
	// Reset calculated fields
	r.TotalPrepTimeMinutes = 0
	r.TotalCookTimeMinutes = 0
	r.TotalCleaningScore = 0
	r.CalculatedPrice = 0.0
	r.CalculatedNutrition = Nutrition{} // Zero out nutrition

	// Calculate time from steps
	for _, step := range r.MethodSteps {
		r.TotalPrepTimeMinutes += step.PrepTimeMinutes
		r.TotalCookTimeMinutes += step.CookTimeMinutes
	}

	// Calculate cleaning score from unique equipment
	// Note: Assumes r.Equipment holds unique items for the recipe
	for _, eq := range r.Equipment {
		r.TotalCleaningScore += int(eq.CleaningDifficulty)
	}

	// Calculate price and nutrition (more complex - needs unit conversion logic)
	// This is a placeholder - requires a robust unit conversion system
	var totalNutrition Nutrition
	var totalPrice float32
	for _, ri := range r.Ingredients {
		if ri.FoodItem.ID == 0 {
			continue
		} // Skip if FoodItem wasn't loaded properly

		// !!! --- Placeholder --- !!!
		// !!! This needs proper unit conversion between ri.Unit and ri.FoodItem.BaseUnit !!!
		// For now, assume quantity is directly comparable (highly unlikely in reality)
		// factor := ri.Quantity // This is wrong without conversion

		// Example: If ri.Unit == ri.FoodItem.BaseUnit (simplest case)
		factor := ri.Quantity
		if ri.FoodItem.BaseUnit != "" && ri.Unit == ri.FoodItem.BaseUnit { // Basic check
			totalPrice += ri.FoodItem.PricePerBaseUnit * factor
			totalNutrition.CarbsGrams += ri.FoodItem.Nutrition.CarbsGrams * factor
			totalNutrition.FatGrams += ri.FoodItem.Nutrition.FatGrams * factor
			totalNutrition.ProteinGrams += ri.FoodItem.Nutrition.ProteinGrams * factor
			totalNutrition.CaloriesKcal += ri.FoodItem.Nutrition.CaloriesKcal * factor
		} else {
			// TODO: Implement unit conversion logic here based on ri.Unit and ri.FoodItem.BaseUnit
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
}
