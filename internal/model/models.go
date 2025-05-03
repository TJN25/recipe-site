package model

type CleaningDifficulty int

const (
	CleaningDifficultyNone   CleaningDifficulty = iota // or near none e.g. paper towel, baking paper, teaspoon for tasting etc.
	CleaningDifficultyEasy                             // something that goes in the dishwasher or sink and requires little space e.g. Spatula
	CleaningDifficultyMedium                           // something that requires some attention, or is a bit bigger e.g. chef's knife, small chopping board
	CleaningDifficultyHard                             // requires care, has things baked on, is very large, can't be cleaned while cooking e.g. caserole dish
)

type RecipeStep struct {
	ID          int64              `json:"id"`         // Will be DB ID later
	RecipeID    int64              `json:"recipe_id"`  // can be used in multiple recipes
	StepOrder   int                `json:"step_order"` // We will just determine this by it's position in the array. Allows it to be moved around without needing to renumber
	Title       string             `json:"title"`      // e.g., "Prepare Chicken", "Make Sauce"
	Description string             `json:"description,omitempty"`
	Notes       string             `json:"notes,omitempty"`
	Ingredients []RecipeIngredient `json:"ingredients"`  // Ingredients *specific* to this step
	MethodSteps []MethodStep       `json:"method_steps"` // Sub-steps *within* this step
	Equipment   []Equipment        `json:"equipment"`    // Equipment *specific* to this step

	// Dependencies []RecipeDependency `json:"dependencies,omitempty"`
	BaseServings int `json:"base_servings"` // Define the scaling used in the section. We may need to ensure this is consistent across steps
}

type RecipeIngredient struct {
	FoodItemID    int64   `json:"food_item_id"`   // Link to the FoodItem concept (e.g., Garlic ID, Cheddar Cheese ID)
	FoodItemName  string  `json:"food_item_name"` // Link to the FoodItem concept (e.g., Garlic ID, Cheddar Cheese ID)
	FormName      string  `json:"form_name"`
	Quantity      float32 `json:"quantity"`
	SpecifiedUnit string  `json:"unit"`

	IsOptional  bool   `json:"is_optional"`       // Is this ingredient optional for the recipe?
	Purpose     string `json:"purpose,omitempty"` // Why this ingredient is used (e.g., "thickener", "acidity")
	Preparation string `json:"preparation,omitempty"`
}

type FoodItem struct {
	ID                 int64                          `json:"id"`
	Name               string                         `json:"name"`           // The primary/generic name (e.g., "Garlic", "Cheddar Cheese", "All-Purpose Flour")
	FormComparisonUnit string                         `json:"canonical_unit"` // The standard unit for internal calculations & nutrition (e.g., "g")
	Forms              map[string]FoodItemFormDetails `json:"forms"`          // Holding details of each form e.g. Clove and the values for Clove.
	DefaultFormName    string                         `json:"default_form_name"`
}

type FoodItemFormDetails struct {
	FormName              string             `json:"form_name"`               // redundant based on maps, but keeping it in case we need to use it
	Unit                  string             `json:"unit"`                    // e.g. clove, tsp, g, cup
	ConversionScaleFactor float32            `json:"conversion_scale_factor"` // e.g. 5.0 if CanonicalUnit is 'g' and this form is 'clove'.
	ToGrams               float32            `json:"to_grams"`
	ToMl                  float32            `json:"to_ml"`
	PricePerUnit          float32            `json:"price_per_unit,omitempty"`
	NutritionPerUnit      float32            `json:"nutrition_per_unit,omitempty"`
	UnitConversions       map[string]float32 `json:"unit_conversions,omitempty"` // e.g., For Flour (Form: Default, Unit: g), UnitConversions: {"cup": 120.0, "tbsp": 7.5}
	MakeableRecipeStepIDs []int64            `json:"makeable_recipe_ids"`        // ID that points to the recipe steps needed to make the recipe
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

type RecipeVariant struct {
	// ID   int64  `json:"id"` // Optional: A unique ID for the variant itself
	Key  string `json:"key"`  // Required: A unique machine-readable key (e.g., "chipotle_chicken", "vegetarian")
	Name string `json:"name"` // Required: User-friendly display name (e.g., "Chipotle Chicken Version")

	// Option B: Variant contains steps to ADD or MODIFY base steps (More Complex)
	AddedSteps []RecipeStep `json:"added_steps,omitempty"` // Steps to add
	// Need info on WHERE to add them (e.g., before/after which base StepID?)
	// InsertionPoints map[string]int64 `json:"insertion_points,omitempty"` // e.g., {"add_after_step_id": 102}

	// Could also include modified ingredients/quantities, but that gets very complex.

	// Let's assume Option A initially for simplicity: The variant defines the *complete* step sequence when selected.
	RecipeSteps []RecipeStep `json:"recipe_steps"`
}

type Recipe struct {
	ID            int64   `json:"id"` // Database primary key
	RecipeStepIds []int64 `json:"recipe_step_ids"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Servings      int     `json:"servings"`        // Base servings for the whole recipe
	Notes         string  `json:"notes,omitempty"` // Overall notes
	ImagePath     string  `json:"image_path,omitempty"`

	// Relationships / Components
	RecipeSteps []RecipeStep    `json:"recipe_steps"` // The ordered list of stages/steps
	Tags        []Tag           `json:"tags"`         // Overall tags for the recipe
	Variants    []RecipeVariant `json:"variants,omitempty"`

	// --- Optional: Add aggregated list for convenience? ---
	// AllIngredients []RecipeIngredient `json:"-"` // Could be calculated, not stored directly

	// Calculated/Derived fields
	TotalPrepTimeMinutes int       `json:"total_prep_time_minutes"`
	TotalCookTimeMinutes int       `json:"total_cook_time_minutes"`
	TotalCleaningScore   int       `json:"total_cleaning_score"`
	CalculatedNutrition  Nutrition `json:"calculated_nutrition"`
	CalculatedPrice      float32   `json:"calculated_price"`
}

type GlobalUserSettings struct {
	DefaultServings   int
	DisplayUnitSystem string
	ShowOptions       bool
}

type RecipeUserConfig struct {
	ActiveRecipeStepIDs []int64
	//      Let's use a composite key string "RecipeID-RecipeStepID-IngredientIndex" for now, needs refinement.
	IngredientOverrides map[string]IngredientUserOverride
	ShowOptionals       bool
}

type IngredientUserOverride struct {
	SelectedFormName   string
	SelectedUnitName   string
	QuantityMultiplier float64 //
}
