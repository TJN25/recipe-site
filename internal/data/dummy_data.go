package data

import (
	"github.com/TJN25/recipe-site/internal/model"
)

var placeholderNutrition = model.Nutrition{}

var DummyFoodItems = []model.FoodItem{
	{
		ID:              10,
		Name:            "Macaroni Pasta",
		CanonicalUnit:   "g",
		DefaultFormName: "Default", // Assuming standard dry pasta form
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0}, // Base unit is grams
		},
		PricePerCanonicalUnit: 0.005,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              11,
		Name:            "Everyday Cheese",
		CanonicalUnit:   "g",
		DefaultFormName: "Block", // Assume sold/used as a block primarily
		Forms: map[string]model.FoodItemFormDetails{
			"Block": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"cup shredded": 113.0, "oz": 28.35}},
			"Shredded": {
				Unit:                  "cup shredded", // The unit for this specific form entry
				ConversionToCanonical: 113.0,          // 1 'cup shredded' = 113 'g' (canonical unit)
				// UnitConversions could potentially convert to other volumes if needed, e.g. {"litre": 0.236}
			},
		},
		PricePerCanonicalUnit: 0.02,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              12,
		Name:            "Cheddar Cheese",
		CanonicalUnit:   "g",
		DefaultFormName: "Block",
		Forms: map[string]model.FoodItemFormDetails{
			"Block": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"cup shredded": 113.0, "oz": 28.35}},
		},
		PricePerCanonicalUnit: 0.025,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              13,
		Name:            "Butter",
		CanonicalUnit:   "g",
		DefaultFormName: "Default", // Assuming standard butter block/stick form
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tbsp": 14.2, "cup": 227.0}},
		},
		PricePerCanonicalUnit: 0.01,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              14,
		Name:            "All‑Purpose Flour",
		CanonicalUnit:   "g",
		DefaultFormName: "Default", // Standard bagged flour
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"cup": 120.0, "tbsp": 7.5}},
		},
		PricePerCanonicalUnit: 0.002,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              15,
		Name:            "Full Fat Milk",
		CanonicalUnit:   "ml",
		DefaultFormName: "Default", // Standard liquid form
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"cup": 240.0, "tbsp": 15.0, "litre": 1000.0}},
		},
		PricePerCanonicalUnit: 0.002,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              16,
		Name:            "Salt",
		CanonicalUnit:   "g",
		DefaultFormName: "Kosher", // Specify the default type/form
		Forms: map[string]model.FoodItemFormDetails{
			"Kosher": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 6.0}},
			"Coarse": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 4.0}}, // Example if you add coarse salt later
		},
		PricePerCanonicalUnit: 0.001,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              17,
		Name:            "Mustard Powder",
		CanonicalUnit:   "g",
		DefaultFormName: "Powder",
		Forms: map[string]model.FoodItemFormDetails{
			"Powder": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 3.0}}, // Estimate ~3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              18,
		Name:            "Smoked Paprika",
		CanonicalUnit:   "ml",
		DefaultFormName: "Ground",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.06,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              19,
		Name:            "Chipotle Powder",
		CanonicalUnit:   "g",
		DefaultFormName: "Powder",
		Forms: map[string]model.FoodItemFormDetails{
			"Powder": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.5}}, // Estimate ~2.5g/tsp
		},
		PricePerCanonicalUnit: 0.08,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              20,
		Name:            "Mexican Chicken", // Assume pre-cooked/shredded
		CanonicalUnit:   "g",
		DefaultFormName: "Shredded",
		Forms: map[string]model.FoodItemFormDetails{
			"Shredded": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"cup": 140.0}}, // Estimate ~140g/cup cooked shredded
		},
		PricePerCanonicalUnit: 0.015,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              124,
		Name:            "Ground Coriander", // New
		CanonicalUnit:   "ml",
		DefaultFormName: "Ground",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.0}}, // Estimate ~2g/tsp
		},
		PricePerCanonicalUnit: 0.06, // Guess
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              125,
		Name:            "Green Beans", // New
		CanonicalUnit:   "g",
		DefaultFormName: "Whole",
		Forms: map[string]model.FoodItemFormDetails{
			"Whole": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"handful": 40.0}}, // Estimate ~40g/handful
		},
		PricePerCanonicalUnit: 0.01, // Guess
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              126,
		Name:            "Roasted Capsicum Strips", // New (Assume jarred)
		CanonicalUnit:   "g",
		DefaultFormName: "Strips",
		Forms: map[string]model.FoodItemFormDetails{
			"Strips": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"cup": 150.0}}, // Estimate ~150g/cup drained
		},
		PricePerCanonicalUnit: 0.03, // Guess
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              127,
		Name:            "Fresh Parsley", // New
		CanonicalUnit:   "g",
		DefaultFormName: "Chopped",
		Forms: map[string]model.FoodItemFormDetails{
			"Chopped": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tbsp": 4.0}}, // Estimate ~4g/tbsp chopped
			"Bunch":   {Unit: "g", ConversionToCanonical: 30.0},                                                  // Estimate ~30g/bunch
		},
		PricePerCanonicalUnit: 0.10, // Guess
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              128,
		Name:            "Fresh Basil", // New
		CanonicalUnit:   "g",
		DefaultFormName: "Chopped",
		Forms: map[string]model.FoodItemFormDetails{
			"Chopped": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tbsp": 3.0}}, // Estimate ~3g/tbsp chopped
			"Bunch":   {Unit: "g", ConversionToCanonical: 25.0},                                                  // Estimate ~25g/bunch
		},
		PricePerCanonicalUnit: 0.12, // Guess
		Nutrition:             placeholderNutrition,
	},
	// --- Add conversions/forms to existing items if needed ---
	// Example: Add Paste/Ground forms to Ginger (ID 30)
	// Example: Add form to Onion (ID 29)

	// --- New Items from Recipes ---
	{ID: 21, Name: "Bread Flour", CanonicalUnit: "g", PricePerCanonicalUnit: 0.003},
	{ID: 22, Name: "Active Dry Yeast", CanonicalUnit: "g", PricePerCanonicalUnit: 0.04},
	{ID: 23, Name: "Raw Sugar", CanonicalUnit: "g", PricePerCanonicalUnit: 0.002},
	{ID: 24, Name: "Olive Oil", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.02},
	// {ID: 25, Name: "Water", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.0001},
	{
		ID:              25,
		Name:            "Water",
		CanonicalUnit:   "ml",
		DefaultFormName: "Default",
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	// {ID: 26, Name: "Chicken Thighs", CanonicalUnit: "g", PricePerCanonicalUnit: 0.012}, // Boneless, skinless
	{
		ID:              26,
		Name:            "Chicken Thigh", // Assume pre-cooked/shredded
		CanonicalUnit:   "g",
		DefaultFormName: "Default",
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0}, // Estimate ~140g/cup cooked shredded
		},
		PricePerCanonicalUnit: 0.012,
		Nutrition:             placeholderNutrition,
	},
	{ID: 27, Name: "Greek Yogurt", CanonicalUnit: "g", PricePerCanonicalUnit: 0.02},
	{ID: 28, Name: "Heavy Cream", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.03},
	{
		ID:              29,
		Name:            "Onion", // Assuming brown/yellow
		CanonicalUnit:   "g",
		DefaultFormName: "Whole",
		Forms: map[string]model.FoodItemFormDetails{
			"Whole":  {Unit: "unit", ConversionToCanonical: 150.0}, // Estimate avg onion ~150g
			"Sliced": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"cup": 115.0}},
		},
		PricePerCanonicalUnit: 0.002,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              30,
		Name:            "Ginger",
		CanonicalUnit:   "g",
		DefaultFormName: "Paste", // Assuming fresh root is default unless specified
		Forms: map[string]model.FoodItemFormDetails{
			"Fresh Root": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tbsp minced": 6.0, "inch": 15.0}},
			"Paste":      {Unit: "ml", ConversionToCanonical: 4.0}, // Estimate ~4g/tsp paste
			"Ground":     {Unit: "ml", ConversionToCanonical: 2.0}, // Estimate ~2g/tsp ground
		},
		PricePerCanonicalUnit: 0.02,
		Nutrition:             placeholderNutrition,
	},
	// {ID: 31, Name: "Garlic", CanonicalUnit: "g", PricePerCanonicalUnit: 0.03}, // Fresh garlic cloves/minced
	{
		ID:              31,
		Name:            "Garlic", // Assume pre-cooked/shredded
		CanonicalUnit:   "cloves",
		DefaultFormName: "Cloves",
		Forms: map[string]model.FoodItemFormDetails{
			"Cloves": {Unit: "cloves", ConversionToCanonical: 1.0}, // Estimate ~140g/cup cooked shredded
		},
		PricePerCanonicalUnit: 0.03,
		Nutrition:             placeholderNutrition,
	},
	{ID: 32, Name: "Cashews", CanonicalUnit: "g", PricePerCanonicalUnit: 0.05}, // Raw, unsalted
	{ID: 33, Name: "Canned Tomatoes", CanonicalUnit: "g", PricePerCanonicalUnit: 0.004},
	// {ID: 34, Name: "Canola Oil", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.01},
	{
		ID:              34,
		Name:            "Canola Oil",
		CanonicalUnit:   "ml",
		DefaultFormName: "Default", // Standard liquid form
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"cup": 240.0, "tbsp": 15.0, "litre": 1000.0}},
		},
		PricePerCanonicalUnit: 0.002,
		Nutrition:             placeholderNutrition,
	},
	{ID: 35, Name: "Baking Soda", CanonicalUnit: "g", PricePerCanonicalUnit: 0.01},
	{ID: 36, Name: "Fenugreek Leaves", CanonicalUnit: "g", PricePerCanonicalUnit: 0.08}, // Kasoori Methi, dried
	// {ID: 37, Name: "Garam Masala", CanonicalUnit: "g", PricePerCanonicalUnit: 0.06},     // Spice blend
	{
		ID:              37,
		Name:            "Garam Masala",
		CanonicalUnit:   "ml",
		DefaultFormName: "Ground",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{ID: 38, Name: "Dried Chili", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.15}, // Generic whole dried chili
	{ID: 39, Name: "Black Cardamom Pod", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.10},
	{ID: 40, Name: "Clove", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.05},      // Whole clove
	{ID: 41, Name: "Basmati Rice", CanonicalUnit: "g", PricePerCanonicalUnit: 0.005}, // Dry weight
	// {ID: 42, Name: "White Pepper", CanonicalUnit: "g", PricePerCanonicalUnit: 0.04},  // Ground
	{
		ID:              42,
		Name:            "White Pepper",
		CanonicalUnit:   "ml",
		DefaultFormName: "Corns",
		Forms: map[string]model.FoodItemFormDetails{
			"Corns": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{ID: 43, Name: "Spicy Burrito Sauce", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.03}, // Derived item (potentially homemade)
	{ID: 44, Name: "Corn Kernels", CanonicalUnit: "g", PricePerCanonicalUnit: 0.006},        // Frozen or Canned
	{ID: 45, Name: "Lime Juice", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.03},          // Fresh or bottled
	{ID: 46, Name: "Cilantro", CanonicalUnit: "g", PricePerCanonicalUnit: 0.05},             // Fresh coriander
	{ID: 47, Name: "Baguette", CanonicalUnit: "unit", PricePerCanonicalUnit: 1.50},
	{ID: 48, Name: "Duck Breast", CanonicalUnit: "g", PricePerCanonicalUnit: 0.03},
	{ID: 49, Name: "Daikon Radish", CanonicalUnit: "g", PricePerCanonicalUnit: 0.008},
	// {ID: 50, Name: "Carrot", CanonicalUnit: "g", PricePerCanonicalUnit: 0.003},
	{
		ID:              50,
		Name:            "Carrot", // Assume pre-cooked/shredded
		CanonicalUnit:   "g",
		DefaultFormName: "Default",
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0}, // Estimate ~140g/cup cooked shredded
		},
		PricePerCanonicalUnit: 0.01,
		Nutrition:             placeholderNutrition,
	},
	{ID: 51, Name: "Cucumber", CanonicalUnit: "g", PricePerCanonicalUnit: 0.005},
	{ID: 52, Name: "Bird Eye Chilies", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.20}, // Fresh
	{ID: 53, Name: "Spring Onion", CanonicalUnit: "g", PricePerCanonicalUnit: 0.04},        // Scallion
	{ID: 54, Name: "Mayonnaise", CanonicalUnit: "g", PricePerCanonicalUnit: 0.015},
	// {ID: 55, Name: "Soy Sauce", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.01},
	{
		ID:              55,
		Name:            "Soy sauce", // Assume pre-cooked/shredded
		CanonicalUnit:   "ml",
		DefaultFormName: "Default",
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "ml", ConversionToCanonical: 1.0}, // Estimate ~140g/cup cooked shredded
		},
		PricePerCanonicalUnit: 0.01,
		Nutrition:             placeholderNutrition,
	},
	{ID: 56, Name: "Tabasco", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.08},
	// {ID: 57, Name: "Honey", CanonicalUnit: "g", PricePerCanonicalUnit: 0.02},
	{
		ID:              57,
		Name:            "Honey",
		CanonicalUnit:   "g",
		DefaultFormName: "Default",
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{ID: 58, Name: "White Vinegar", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.002}, // Distilled white vinegar
	{ID: 59, Name: "Agria Potato", CanonicalUnit: "g", PricePerCanonicalUnit: 0.004},
	{ID: 60, Name: "Avocado Oil", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.03},
	{ID: 61, Name: "Garlic Powder", CanonicalUnit: "g", PricePerCanonicalUnit: 0.04},
	{
		ID:              62,
		Name:            "Black Pepper",
		CanonicalUnit:   "ml",
		DefaultFormName: "Corns",
		Forms: map[string]model.FoodItemFormDetails{
			"Corns": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{ID: 63, Name: "Mixed Herbs", CanonicalUnit: "g", PricePerCanonicalUnit: 0.05}, // Dried blend (e.g., Italian seasoning)
	{
		ID:              64,
		Name:            "Cayenne pepper",
		CanonicalUnit:   "ml",
		DefaultFormName: "Ground",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{ID: 65, Name: "Parmesan Cheese", CanonicalUnit: "g", PricePerCanonicalUnit: 0.04},    // Parmigiano Reggiano or Grana Padano
	{ID: 66, Name: "Flour Tortilla", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.50},  // Large size
	{ID: 67, Name: "Lettuce", CanonicalUnit: "g", PricePerCanonicalUnit: 0.01},            // Generic (e.g., Iceberg, Romaine)
	{ID: 68, Name: "Pickled Red Onions", CanonicalUnit: "g", PricePerCanonicalUnit: 0.02}, // Derived item (Recipe ID 3)
	{ID: 69, Name: "Tomato", CanonicalUnit: "g", PricePerCanonicalUnit: 0.006},            // Fresh whole tomato
	{ID: 70, Name: "Black Beans", CanonicalUnit: "g", PricePerCanonicalUnit: 0.005},       // Canned/Cooked, drained weight
	{ID: 71, Name: "Avocado", CanonicalUnit: "unit", PricePerCanonicalUnit: 2.00},
	{ID: 72, Name: "Sour Cream", CanonicalUnit: "g", PricePerCanonicalUnit: 0.015},
	{ID: 73, Name: "Pickled Jalapeños", CanonicalUnit: "g", PricePerCanonicalUnit: 0.03},
	{
		ID:              74,
		Name:            "Egg",
		CanonicalUnit:   "unit",
		DefaultFormName: "Whole",
		Forms: map[string]model.FoodItemFormDetails{
			"Whole": {Unit: "unit", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{ID: 75, Name: "Trim Milk", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.0018}, // Low-fat milk
	{ID: 76, Name: "Fettuccine", CanonicalUnit: "g", PricePerCanonicalUnit: 0.006},  // Assumed dry pasta
	{
		ID:              77,
		Name:            "Mushrooms", // Assume pre-cooked/shredded
		CanonicalUnit:   "g",
		DefaultFormName: "Default",
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0}, // Estimate ~140g/cup cooked shredded
		},
		PricePerCanonicalUnit: 0.01,
		Nutrition:             placeholderNutrition,
	},
	{ID: 78, Name: "Celery", CanonicalUnit: "g", PricePerCanonicalUnit: 0.004},
	{ID: 79, Name: "White Wine", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.015},    // Dry white wine for cooking
	{ID: 80, Name: "Chicken Stock", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.003}, // Liquid stock/broth
	// {ID: 81, Name: "Spinach", CanonicalUnit: "g", PricePerCanonicalUnit: 0.02},         // Fresh spinach leaves
	{
		ID:              81,
		Name:            "Spinach",
		CanonicalUnit:   "g",
		DefaultFormName: "Default",
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{ID: 82, Name: "Cherry Tomato", CanonicalUnit: "g", PricePerCanonicalUnit: 0.01},
	{ID: 83, Name: "Chives", CanonicalUnit: "g", PricePerCanonicalUnit: 0.06}, // Fresh
	{ID: 84, Name: "Feta Cheese", CanonicalUnit: "g", PricePerCanonicalUnit: 0.03},
	{ID: 85, Name: "Lemon Juice", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.03}, // Fresh or bottled
	{ID: 86, Name: "Brown Sugar", CanonicalUnit: "g", PricePerCanonicalUnit: 0.003}, // Light or dark
	{ID: 87, Name: "Chicken Stock Cube", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.20},
	{ID: 88, Name: "Red Taco Sauce", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.025}, // Derived item (Recipe ID TBD)
	{ID: 89, Name: "Worcestershire Sauce", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.02},
	// {ID: 90, Name: "Egg Noodles", CanonicalUnit: "g", PricePerCanonicalUnit: 0.007}, // Assumed dry
	{
		ID:              90,
		Name:            "Egg Noodles",
		CanonicalUnit:   "g",
		DefaultFormName: "Default",
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{ID: 91, Name: "Fresh Hot Chilies", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.30}, // Generic fresh hot chili (e.g., jalapeño, serrano)
	{ID: 92, Name: "Canned Peach Halves", CanonicalUnit: "g", PricePerCanonicalUnit: 0.008}, // Drained weight, in syrup or juice
	{ID: 93, Name: "Cardamom", CanonicalUnit: "g", PricePerCanonicalUnit: 0.12},             // Ground green cardamom
	{ID: 94, Name: "Cabbage", CanonicalUnit: "g", PricePerCanonicalUnit: 0.003},             // Green or white cabbage
	{ID: 95, Name: "Buttermilk", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.005},
	{ID: 96, Name: "Granulated Sugar", CanonicalUnit: "g", PricePerCanonicalUnit: 0.002}, // White sugar
	{ID: 97, Name: "Smooth Peanut Butter", CanonicalUnit: "g", PricePerCanonicalUnit: 0.01},
	{ID: 98, Name: "Chilli Powder", CanonicalUnit: "g", PricePerCanonicalUnit: 0.06},       // Generic chili powder blend
	{ID: 99, Name: "Kashmiri Chilies", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.40}, // Dried whole, mild heat, good color
	{ID: 100, Name: "Chipotle Pepper", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.50}, // Dried whole smoked jalapeño
	{ID: 101, Name: "Chorizo Sausage", CanonicalUnit: "g", PricePerCanonicalUnit: 0.02},    // Assumed fresh Mexican chorizo
	{ID: 102, Name: "Brioche Buns", CanonicalUnit: "unit", PricePerCanonicalUnit: 1.00},
	{ID: 103, Name: "Smoky Salsa Verde", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.03}, // Derived item (potentially homemade)
	{ID: 104, Name: "Canned Diced Tomatoes", CanonicalUnit: "g", PricePerCanonicalUnit: 0.004},
	{ID: 105, Name: "Tomato Paste", CanonicalUnit: "g", PricePerCanonicalUnit: 0.01},
	{ID: 106, Name: "Dried Oregano", CanonicalUnit: "g", PricePerCanonicalUnit: 0.04},
	{ID: 107, Name: "Red Pepper Flakes", CanonicalUnit: "g", PricePerCanonicalUnit: 0.05},  // Crushed red pepper
	{ID: 108, Name: "Chipotles in Adobo", CanonicalUnit: "g", PricePerCanonicalUnit: 0.07}, // Whole chipotles packed in adobo sauce
	// {ID: 109, Name: "Cumin", CanonicalUnit: "g", PricePerCanonicalUnit: 0.05},              // Ground
	{
		ID:              109,
		Name:            "Cumin",
		CanonicalUnit:   "ml",
		DefaultFormName: "Ground",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	{ID: 110, Name: "Apple Cider Vinegar", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.003},
	{ID: 111, Name: "Milk Powder", CanonicalUnit: "g", PricePerCanonicalUnit: 0.015}, // Non-fat dry milk powder
	{ID: 112, Name: "Diastatic Malt Powder", CanonicalUnit: "g", PricePerCanonicalUnit: 0.05},
	{ID: 113, Name: "Generic Rice", CanonicalUnit: "g", PricePerCanonicalUnit: 0.003},             // Medium or long grain white rice, dry
	{ID: 114, Name: "Quick Pizza Sauce", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.02},        // Derived item (Recipe ID 4)
	{ID: 115, Name: "Ginger-Peach Fire Sauce", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.035}, // Derived item (Recipe ID TBD)
	// {ID: 116, Name: "Sesame Oil", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.05},
	{
		ID:              116,
		Name:            "Sesame Oil",
		CanonicalUnit:   "ml",
		DefaultFormName: "Ground",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"tsp": 2.3}}, // Estimate ~2.3g/tsp
		},
		PricePerCanonicalUnit: 0.05,
		Nutrition:             placeholderNutrition,
	},
	// {ID: 117, Name: "Sichuan Pepper", CanonicalUnit: "g", PricePerCanonicalUnit: 0.15}, // Ground
	{
		ID:              117,
		Name:            "Sichuan Pepper",
		CanonicalUnit:   "g",
		DefaultFormName: "Default", // Assuming fresh root is default unless specified
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: 0.15,
		Nutrition:             placeholderNutrition,
	},
	{ID: 118, Name: "Allspice", CanonicalUnit: "g", PricePerCanonicalUnit: 0.07}, // Ground
	// {ID: 119, Name: "Cornflour", CanonicalUnit: "g", PricePerCanonicalUnit: 0.005}, // Cornstarch
	{
		ID:              119,
		Name:            "Cornflour",
		CanonicalUnit:   "ml",
		DefaultFormName: "Default", // Standard liquid form
		Forms: map[string]model.FoodItemFormDetails{
			"Default": {Unit: "ml", ConversionToCanonical: 1.0, UnitConversions: map[string]float32{"cup": 240.0, "tbsp": 15.0, "litre": 1000.0}},
		},
		PricePerCanonicalUnit: 0.002,
		Nutrition:             placeholderNutrition,
	},
	{ID: 120, Name: "Kewpie Mayo", CanonicalUnit: "g", PricePerCanonicalUnit: 0.025},
	{ID: 121, Name: "Sriracha Sauce", CanonicalUnit: "ml", PricePerCanonicalUnit: 0.03},
	{ID: 122, Name: "Bao Buns", CanonicalUnit: "unit", PricePerCanonicalUnit: 0.80}, // Steamed buns
	{ID: 123, Name: "Roasted Peanuts", CanonicalUnit: "g", PricePerCanonicalUnit: 0.02},
}

var DummyEquipment = []model.Equipment{
	// --- Original Items ---
	{ID: 50, Name: "Large Saucepan", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 51, Name: "Cheese Grater", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyMedium},
	{ID: 52, Name: "Medium Casserole Dish", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyHard}, // Specified size
	{ID: 53, Name: "Whisk", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyMedium},

	// --- New Items from Recipes ---
	{ID: 54, Name: "Large Mixing Bowl", Type: "Container", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 55, Name: "Medium Mixing Bowl", Type: "Container", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 56, Name: "Small Mixing Bowl", Type: "Container", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 57, Name: "Chef's Knife", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 58, Name: "Chopping Board", Type: "Surface", CleaningDifficulty: model.CleaningDifficultyMedium}, // Can vary
	{ID: 59, Name: "Wire Rack", Type: "Cooling", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 60, Name: "Measuring Cups", Type: "Measuring", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 61, Name: "Measuring Spoons", Type: "Measuring", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 62, Name: "Fork", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},                 // Eating/Mixing
	{ID: 63, Name: "Stirring Spoon", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},       // Wooden or silicone
	{ID: 64, Name: "Cast Iron Pan", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyHard},       // Requires seasoning care
	{ID: 65, Name: "Oven", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyMedium},             // Interior cleaning
	{ID: 66, Name: "Stovetop/Cooktop", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyMedium}, // Surface cleaning
	{ID: 67, Name: "Heatproof Spatula", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},    // Silicone/Rubber
	{ID: 68, Name: "Countertop Blender", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyHard}, // Blade area
	{ID: 69, Name: "Small Saucepan", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 70, Name: "Baking Tray / Sheet Pan", Type: "Bakeware", CleaningDifficulty: model.CleaningDifficultyMedium},
	{ID: 71, Name: "Kitchen Tongs", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 72, Name: "Vegetable Peeler", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 73, Name: "Toaster Oven", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyMedium},
	{ID: 74, Name: "Ramekin", Type: "Bakeware", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 75, Name: "Round Cutter", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy}, // e.g., biscuit cutter
	{ID: 76, Name: "Glass Jar", Type: "Container", CleaningDifficulty: model.CleaningDifficultyEasy},  // For pickles, etc.
	{ID: 77, Name: "Sandwich Press / Toastie Machine", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyMedium},
	{ID: 78, Name: "Bench Scraper", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 79, Name: "Fine Mesh Strainer / Sieve", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyMedium},
	{ID: 80, Name: "Slotted Spoon", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 81, Name: "Microwave", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 82, Name: "Spray Bottle", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy}, // For water/oil
	{ID: 83, Name: "Instant Read Thermometer", Type: "Measuring", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 84, Name: "Wok", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyMedium},
	{ID: 85, Name: "Spice Grinder", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyMedium}, // Electric or manual
	{ID: 86, Name: "Immersion Blender", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyMedium},
	{ID: 87, Name: "Lame / Scoring Tool", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},     // For bread
	{ID: 88, Name: "Small Oven-safe Dish", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyMedium}, // For steam, broiling small amounts
	{ID: 89, Name: "Food Processor", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyHard},
	{ID: 90, Name: "Kitchen Towel / Dish Cloth", Type: "Accessory", CleaningDifficulty: model.CleaningDifficultyEasy}, // Reusable textile
	{ID: 91, Name: "Mortar and Pestle", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyMedium},
	{ID: 92, Name: "Rolling Pin", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy}, // Alternative for crushing
	{ID: 93, Name: "Plate", Type: "Serving", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 94, Name: "Steamer", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyMedium}, // Bamboo or metal insert
	// Excluded consumables like paper towels, parchment paper, plastic wrap
}

var DummyTags = []model.Tag{
	{ID: 80, Name: "American", Type: "cuisine"},
	{ID: 81, Name: "Comfort", Type: "mood"},
	{ID: 82, Name: "Pasta", Type: "dish_type"},
	{ID: 89, Name: "Asian Fusion", Type: "cuisine"},
	{ID: 90, Name: "Chinese Inspired", Type: "cuisine"},
	{ID: 91, Name: "Buns", Type: "dish_type"},
	{ID: 92, Name: "Spicy", Type: "flavor_profile"},
	{ID: 93, Name: "Street Food", Type: "mood"},
	// ... add tags for Bao Buns etc ...
}

var DummyRecipes = []struct {
	ID            int64
	RecipeStepIds []int64
	Title         string
	Description   string
	Servings      int
	Notes         string
	ImagePath     string
}{
	// --- Existing Recipe (Updated Notes) ---
	{
		ID:            1,
		RecipeStepIds: []int64{102, 103},
		Title:         "Chipotle Mexican Chicken Mac and Cheese",
		Description:   "A smoky, spicy twist on classic mac and cheese, combined with Mexican chicken. Can be made without chicken/spice for a standard version.",
		Servings:      2,
		Notes:         "Let rest 5 minutes before serving. Uses full fat milk for creaminess. Sauce should look slightly too saucy before baking; it will thicken. For standard Mac & Cheese, omit chipotle powder, paprika, Mexican chicken, and burrito sauce. For extra texture, add 3 tbsp frozen corn with chicken. A squeeze of lime or cilantro garnish adds freshness. Best eaten fresh, refrigerate up to 3 days. Reheat on stovetop with a splash of milk.",
		ImagePath:     "img/mac-and-cheese.png", // Shared image for both versions
	},

	// --- New Recipes ---
	{
		ID:            2,
		RecipeStepIds: []int64{101, 201, 202, 203, 204},
		Title:         "Pizza Dough - Same Day Preparation",
		Description:   "Soft, chewy pizza dough with 75% hydration, ready in 4-5 hours. Makes two 12-14\" pizzas.",
		Servings:      2, // Makes 2 pizza bases
		Notes:         "Total dough weight ≈ 490g (2x 245g portions). Use stretch and fold technique. Cook on stovetop in cast iron then finish in hot oven. See detailed notes for overnight fermentation adjustments (reduce yeast, fewer folds, refrigerate). Handle overproofed dough gently. Allow cold dough to rest before stretching.",
		ImagePath:     "img/pizza-dough.png",
	},
	{
		ID:            3,
		RecipeStepIds: []int64{301, 302, 303, 304, 305},
		Title:         "Butter Chicken (Murgh Makhani)",
		Description:   "Classic creamy Indian chicken dish with a rich tomato and cashew sauce.",
		Servings:      3,
		Notes:         "Marinate chicken up to 24 hours ahead. Cook onions slowly with baking soda for deep browning (may cook faster than stated). Heat level is mild-medium; adjust chili. Sauce keeps 1 week refrigerated; freeze sauce without cream for 3 months. For extra richness, add more butter at the end. Serve with Basmati rice (60g dry per person). User Notes (2025-02-22): Tasted great, good texture. Chicken was quite salty, be cautious adding salt to sauce. Used lime juice.",
		ImagePath:     "img/butter-chicken.png",
	},
	{
		ID:            4,
		RecipeStepIds: []int64{401, 402, 403, 404},
		Title:         "Banh Mi with Duck",
		Description:   "Vietnamese-style baguette sandwich featuring crispy duck breast and quick pickled vegetables.",
		Servings:      2,
		Notes:         "Requires quick pickled daikon/carrot (see Recipe ID 14). Score duck skin well. Cook duck to 63°C for medium-rare. Toast baguettes before assembly. User Notes (2025-02-24): Tasted nice but needed more sauce (spicy, simple, creamy/thick suggested). Toast buns in oven. Duck was overcooked (>70°C), monitor closely.",
		ImagePath:     "img/banh-mi.png",
	},
	{
		ID:            5,
		RecipeStepIds: []int64{501, 502},
		Title:         "Agria Potato Wedges",
		Description:   "Crispy baked potato wedges using starchy Agria potatoes.",
		Servings:      2,
		Notes:         "Leave skin on. Soak cut wedges in cold water then dry thoroughly for crispiness. For extra crispiness, parboil 5 mins before seasoning. Add parmesan in last 5 mins for flavor. Serve with dip (e.g., yogurt/paprika/garlic). User Notes: Added ½ tsp cayenne, used canola oil instead of avocado. Took ~25 mins in oven. Great taste and texture.",
		ImagePath:     "img/potato-wedges.png",
	},
	{
		ID:            6,
		RecipeStepIds: []int64{601, 602, 603, 604},
		Title:         "Crispy Mexican Chicken Burrito",
		Description:   "A toasted burrito filled with Mexican chicken, cheese, corn, and fresh ingredients.",
		Servings:      1,
		Notes:         "Requires pre-cooked Mexican chicken (ID 20) and pickled onions (ID 14). Microwave tortilla briefly to make pliable. Layering order is important. Roll tightly. Cook seam-side down in sandwich press/toastie machine. Enhancements: add cilantro, lime juice, avocado, black beans, or rice.",
		ImagePath:     "img/chicken-burrito.png",
	},
	{
		ID:            7,
		RecipeStepIds: []int64{701, 702, 703, 704, 705},
		Title:         "English Muffins",
		Description:   "Small, yeasted muffins typically split, toasted, and buttered. Oven-baked method provided.",
		Servings:      4, // Makes 3-4 muffins
		Notes:         "Uses stretch and fold. Dough is sticky; use flour generously or oil hands/scraper. Ramekin used for cutting. Second rise on floured parchment. Oven bake method avoids stovetop griddling. Flour dusting substitutes for cornmeal/semolina but gives different texture. User Notes: Came out well. Good for making ahead and toasting.",
		ImagePath:     "img/english-muffins.png",
	},
	{
		ID:            8,
		RecipeStepIds: []int64{801, 802, 803, 804},
		Title:         "Fettuccine with Mushroom White Wine Sauce",
		Description:   "Pasta dish with a savory sauce made from mushrooms, mirepoix, white wine, and chicken stock.",
		Servings:      2,
		Notes:         "Prep all ingredients first (mise en place). Reserve pasta water to adjust sauce consistency. Serve immediately. User Notes (21-02-2025): Turned out well with fresh ingredients. Slightly salty with fresh parmesan topping. Used homemade wide-cut fettuccine. (05-03-2025): Substituted white wine vinegar + water for wine (used too much vinegar), everyday parmesan, store-bought fettuccine. Enjoyed but dial back vinegar, maybe add mustard powder. (06-03-2025): Added mustard powder, smoked paprika, soy sauce, pepper. No pasta water needed (salty). Used reduced stock (maybe too much). Suggests thinner pasta, slightly less butter, pair with salad (lettuce, spinach, tomato, red onion, spring onion, chives, maybe feta, lemon/pepper dressing).",
		ImagePath:     "img/fettuccine-mushroom.png",
	},
	{
		ID:            9,
		RecipeStepIds: []int64{901, 902, 903},
		Title:         "Fusion Taco Noodles with Red Sauce",
		Description:   "Egg noodles stir-fried with aromatics, egg, and a savory-tangy sauce featuring red taco sauce.",
		Servings:      2,
		Notes:         "Uses Red Taco Sauce (ID 16). Adjust sauce balance (sweetness/acidity) to taste. Cook aromatics carefully to avoid burning garlic. Break egg up quickly while wet. Add reserved noodle water if needed to adjust consistency. Serve immediately.",
		ImagePath:     "img/taco-noodles.png",
	},
	{
		ID:            10,
		RecipeStepIds: []int64{1001, 1002, 1003},
		Title:         "Ginger-Peach Fire Sauce (Component)",
		Description:   "A spicy-sweet-zingy sauce with ginger warmth, inspired by Zambrero's.",
		Servings:      10, // Makes ~1.5-2 cups
		Notes:         "Use fresh ginger (recipe notes possibly too much). Adjust chilies to taste. Use canned peaches. Taste before adding sugar (canned peaches add sweetness). Blend until smooth, strain for best texture. Store refrigerated up to 1 week, freeze excess. Shake well before use.",
		ImagePath:     "img/ginger-peach-sauce.png",
	},
	{
		ID:            11,
		RecipeStepIds: []int64{1101, 1102},
		Title:         "KFC Style Coleslaw",
		Description:   "A copycat recipe for KFC's sweet and tangy coleslaw.",
		Servings:      6,
		Notes:         "Chop cabbage very finely (rice-sized). Requires refrigeration for minimum 1 hour (longer is better) for flavors to meld. Stir before serving.",
		ImagePath:     "img/kfc-coleslaw.png",
	},
	// ID 12 is skipped as the standard Mac & Cheese was merged into ID 1.
	{
		ID:            12, // Renumbered from 13
		RecipeStepIds: []int64{1201, 1202, 1203},
		Title:         "Mi Goreng-Inspired Satay Noodles",
		Description:   "Egg noodles with a spicy peanut satay sauce, inspired by Indonesian Mi Goreng flavors.",
		Servings:      2,
		Notes:         "Use smooth peanut butter. Adjust chilies/chili powder for heat. Keep sauce warm and well-mixed. Watch aromatics carefully. Break egg up quickly. Final sauce should be glossy. Best eaten fresh.",
		ImagePath:     "img/satay-noodles.png",
	},
	{
		ID:            13, // Renumbered from 14
		RecipeStepIds: []int64{1301, 1302, 1303, 1304},
		Title:         "Pambazo (Mexican Chorizo-Potato Sandwich)",
		Description:   "A Mexican sandwich featuring chorizo-potato filling in chili-dipped, fried bread.",
		Servings:      2,
		Notes:         "Uses Kashmiri chilies as substitute for Guajillo. Requires Feta/Parmesan as substitute for Cotija, and Pickled Onions (ID 14). Toast chilies briefly. Par-boil potatoes for faster filling prep. Use sturdy rolls. User Notes (2025-02-23): Very wet meal. Suggests pre-frying buns slightly before dipping. Minimize wet ingredients (salsa, lettuce, onions). Tasted good combined but individual components lacked complexity. Minimize feta or crumble finely. Needs work on salt/acid/texture.",
		ImagePath:     "img/pambazo.png",
	},
	{
		ID:            14, // Renumbered from 15
		RecipeStepIds: []int64{1401},
		Title:         "Quick Pickled Red Onions (Component)",
		Description:   "Simple and fast method for pickling red onions.",
		Servings:      10, // Makes ~200ml jar
		Notes:         "Slice onion thinly root-to-tip. Use hot brine. Ready in 2 hours, best after 24 hours. Stores up to 3 weeks refrigerated. Tap jar to release air bubbles.",
		ImagePath:     "img/pickled-onions.png", // Re-use image
	},
	{
		ID:            15, // Renumbered from 16
		RecipeStepIds: []int64{1501},
		Title:         "Quick Pizza Sauce (Component)",
		Description:   "A fast, blended pizza sauce using canned tomatoes and paste.",
		Servings:      4, // Enough for ~2-4 pizzas
		Notes:         "Drain tomatoes but reserve liquid. Blend until smooth. Should be thick; add reserved liquid only if needed to adjust consistency.",
		ImagePath:     "img/pizza-sauce.png", // Re-use image
	},
	{
		ID:            16, // Renumbered from 17
		RecipeStepIds: []int64{1601, 1602, 1603, 1604},
		Title:         "Red Taco Sauce (Component)",
		Description:   "A smoky and tangy red chili sauce for tacos and other Mexican dishes.",
		Servings:      12, // Makes ~300-350ml
		Notes:         "Uses Kashmiri and chipotle chilies. Roasts garlic and tomatoes for depth. Simmer to thicken. Strain for smoother sauce (optional). Keeps 1 week refrigerated, freezes well. User Notes (Multiple Dates): Tasted good, versatile. Didn't need second heating step. Adjusted sugar/vinegar. Salt adjustment (½ tsp) worked well. Can be quite hot.",
		ImagePath:     "img/red-taco-sauce.png",
	},
	{
		ID:            17, // Renumbered from 18
		RecipeStepIds: []int64{1701, 1702, 1703, 1704, 1705},
		Title:         "Subway-Style Baguette Adaptation",
		Description:   "Multiple attempts documented to replicate soft, thin-crusted Subway-style bread rolls.",
		Servings:      2, // Makes 2 baguettes
		Notes:         "Based on pizza dough recipe (ID 2) with adjustments (less salt, more oil/sugar). Multiple attempts focus on achieving soft crust via temperature, steam (water spray, water pan, covering with glass dish), and post-bake wrapping. Uses stretch & fold. Shaping into logs. Lower baking temps than pizza. User notes track progress through attempts, focusing on crust texture and crumb density. Latest attempts involve higher initial temp, aggressive steam, covering during bake, and immediate wrapping post-bake. Target internal temp 90-93°C.",
		ImagePath:     "img/subway-baguette.png",
	},
	{
		ID:            18, // Assign the next available ID
		RecipeStepIds: []int64{1801, 1802, 1803, 1804},
		Title:         "Spiced Chicken Bao Buns with Sriracha Mayo",
		Description:   "Crispy fried spiced chicken pieces served in soft bao buns with coleslaw, cucumber, peanuts, and a tangy sriracha mayo.",
		Servings:      3, // Makes approx 6-8 buns depending on size/filling
		Notes:         "Marinate chicken for 15-20 mins. Fry chicken in batches. Steam buns just before serving. Assemble just before eating for best texture contrast. Enhancements: Salt cucumber slices briefly before use. Toast peanuts before crushing.",
		ImagePath:     "img/bao-buns-chicken.png", // Assign an appropriate image path
	},
	{
		ID:            19,
		RecipeStepIds: []int64{1901, 1902, 1903, 1904, 1905, 1906},
		Title:         "Crispy Fried Chicken with Vegetable & Egg Noodle Stir-fry",
		Description:   "Crispy fried chicken pieces served over a flavorful stir-fry of vegetables and egg noodles.",
		Servings:      2, // Based on 1-2 range
		Notes:         "Wok Hei: Preheating the wok properly over high heat is crucial for achieving \"wok hei\" - the characteristic smoky flavour of good stir-fries.\nWok Frying: Be mindful when deep-frying/shallow-frying in a wok. The sloped sides mean oil depth varies. Keep pieces moving and adjust heat to prevent burning. Use a wok spatula or spider strainer for removal.\nStir-fry Motion: Use a scooping, tossing motion to move ingredients constantly, ensuring even cooking and preventing sticking. Add sauce around the perimeter to allow it to heat and reduce slightly before coating ingredients.\nSpeed: Wok cooking is fast. Have everything prepped and ready next to the stove before you start heating the wok.",
		ImagePath:     "img/chicken-noodle-stirfry.png", // Placeholder path
	},
}

type DummyRecipeStep struct {
	ID          int64
	RecipeID    int64
	StepOrder   int
	Title       string
	Description string
	Notes       string
	Ingredients []struct { // Ingredient References
		FoodItemID int64
		Quantity   float32
		Unit       string
		IsOptional bool
		Purpose    string
	}
	MethodSteps  []model.MethodStep // Method sub-steps (can be full model struct)
	EquipmentIDs []int64            // Equipment References
}

var DummyRecipeSteps = []DummyRecipeStep{
	// --- Recipe 1: Chipotle Mexican Chicken Mac and Cheese ---
	// Step 1: Prep Pasta & Cheese (Merged Standard & Chipotle)
	{
		ID:          101,
		RecipeID:    1,
		StepOrder:   1,
		Title:       "Prepare Pasta & Cheese",
		Description: "Cook pasta until slightly overcooked. Grate cheeses, reserving some for topping.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 10, Quantity: 150, Unit: "g", Purpose: "starch base"},          // Macaroni Pasta
			{FoodItemID: 11, Quantity: 112, Unit: "g", Purpose: "melting base"},         // Everyday Cheese
			{FoodItemID: 12, Quantity: 113, Unit: "g", Purpose: "flavor"},               // Tasty Cheddar
			{FoodItemID: 16, Quantity: 15, Unit: "g", Purpose: "pasta water seasoning"}, // Salt for pasta water (approx 1.5 tbsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Bring a large saucepan of water to a boil. Add salt."},
			{StepNumber: 2, Instruction: "Add macaroni pasta and cook for 1-2 minutes past al dente (approx. 10-12 minutes total)."},
			{StepNumber: 3, Instruction: "While pasta cooks, grate both cheeses separately. Set aside 37g of each (74g total) for topping."},
			{StepNumber: 4, Instruction: "Drain cooked pasta well, but do not rinse."},
		},
		EquipmentIDs: []int64{50, 51, 54}, // Large Saucepan, Cheese Grater, Mixing Bowl (for cheese)
	},
	// Step 2: Make Cheese Sauce (Merged Standard & Chipotle)
	{
		ID:          102,
		RecipeID:    1,
		StepOrder:   2,
		Title:       "Make Cheese Sauce",
		Description: "Create a roux, add milk, then melt in cheese and seasonings.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 13, Quantity: 38, Unit: "g", Purpose: "richness, roux base"},           // Butter
			{FoodItemID: 14, Quantity: 25, Unit: "g", Purpose: "thickener"},                     // Flour (~2.25 tbsp)
			{FoodItemID: 18, Quantity: 4, Unit: "g", IsOptional: true, Purpose: "smoky depth"},  // Smoked Paprika (~3/4 tsp) - OPTIONAL for standard
			{FoodItemID: 19, Quantity: 1.5, Unit: "g", IsOptional: true, Purpose: "smoky heat"}, // Chipotle Powder (¼-½ tsp) - OPTIONAL for standard
			{FoodItemID: 15, Quantity: 265, Unit: "ml", Purpose: "sauce base"},                  // Full Fat Milk
			{FoodItemID: 11, Quantity: 75, Unit: "g", Purpose: "sauce cheese"},                  // Everyday Cheese (112g total - 37g reserved = 75g)
			{FoodItemID: 12, Quantity: 76, Unit: "g", Purpose: "sauce cheese"},                  // Tasty Cheddar (113g total - 37g reserved = 76g)
			{FoodItemID: 17, Quantity: 1, Unit: "g", Purpose: "enhance cheese"},                 // Mustard Powder (~¼ tsp)
			{FoodItemID: 16, Quantity: 2, Unit: "g", Purpose: "seasoning"},                      // Salt (to taste)
			{FoodItemID: 42, Quantity: 1, Unit: "g", Purpose: "seasoning"},                      // White Pepper (to taste)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In the same large saucepan used for pasta, melt butter over medium heat."},
			{StepNumber: 2, Instruction: "Once melted, add flour. Whisk constantly for 1 minute."},
			{StepNumber: 3, Instruction: "If making Chipotle version, add smoked paprika and chipotle powder. Stir to bloom spices (30 seconds)."},
			{StepNumber: 4, Instruction: "Continue cooking the flour mixture (roux) for a total of 2 minutes, stirring constantly."},
			{StepNumber: 5, Instruction: "Gradually whisk in the full fat milk, starting with about 1/3 cup, whisking until smooth before adding more."},
			{StepNumber: 6, Instruction: "Bring the mixture to a gentle simmer and cook, stirring constantly, until thickened (4-5 minutes)."},
			{StepNumber: 7, Instruction: "Remove the saucepan from heat."},
			{StepNumber: 8, Instruction: "Add the remaining grated cheeses (151g total) and mustard powder."},
			{StepNumber: 9, Instruction: "Season with salt and white pepper to taste."},
			{StepNumber: 10, Instruction: "Stir until the cheese is completely melted and the sauce is smooth."},
		},
		EquipmentIDs: []int64{50, 53}, // Large Saucepan, Whisk
	},
	// Step 3: Combine and Finish (Merged Standard & Chipotle)
	{
		ID:          103,
		RecipeID:    1,
		StepOrder:   3,
		Title:       "Combine, Top, and Grill",
		Description: "Combine pasta with sauce (and optional chicken/burrito sauce), transfer to dish, top with cheese, and grill.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 20, Quantity: 150, Unit: "g", IsOptional: true, Purpose: "protein, flavor"}, // Mexican Chicken - OPTIONAL
			{FoodItemID: 43, Quantity: 22, Unit: "ml", IsOptional: true, Purpose: "extra flavor"},    // Spicy Burrito Sauce (~1.5 tbsp) - OPTIONAL
			{FoodItemID: 44, Quantity: 30, Unit: "g", IsOptional: true, Purpose: "texture"},          // Frozen Corn (~3 tbsp) - OPTIONAL Enhancement
			// Reserved cheese (74g total) from Step 1 is used here
			{FoodItemID: 18, Quantity: 1, Unit: "g", IsOptional: true, Purpose: "visual appeal"}, // Smoked Paprika (light dusting) - OPTIONAL
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Add the cooked pasta to the cheese sauce."},
			{StepNumber: 2, Instruction: "If making Chipotle version, fold in the shredded Mexican chicken, optional burrito sauce, and optional corn."},
			{StepNumber: 3, Instruction: "Mix gently until everything is coated. The mixture should look slightly too saucy."},
			{StepNumber: 4, Instruction: "Transfer the mixture to a medium oven-safe casserole dish."},
			{StepNumber: 5, Instruction: "Sprinkle the reserved cheese mix (74g) evenly over the top."},
			{StepNumber: 6, Instruction: "Optional: Sprinkle a light dusting of smoked paprika on top."},
			{StepNumber: 7, Instruction: "Place the dish under a hot grill (or in a toaster oven on grill setting)."},
			{StepNumber: 8, Instruction: "Grill for 5-6 minutes, or until the top is golden brown and bubbly with some darker spots."},
			{StepNumber: 9, Instruction: "Remove from the grill and let rest for 5 minutes before serving."},
		},
		EquipmentIDs: []int64{50, 52, 69, 63}, // Large Saucepan, Casserole Dish, Toaster Oven / Grill, Spatula
	},

	// --- Recipe 2: Pizza Dough - Same Day Preparation ---
	{
		ID:          201,
		RecipeID:    2,
		StepOrder:   1,
		Title:       "Activate Yeast",
		Description: "Bloom the yeast in warm water with sugar.",
		Notes:       "Water temperature should be 40-43°C.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 25, Quantity: 210, Unit: "ml", Purpose: "hydration"},         // Water
			{FoodItemID: 22, Quantity: 5, Unit: "g", Purpose: "leavening"},            // Active Dry Yeast (~3/4 tsp)
			{FoodItemID: 23, Quantity: 5, Unit: "g", Purpose: "yeast food, browning"}, // Raw Sugar (~3/4 tsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In a small bowl or measuring cup, combine warm water, yeast, and sugar."},
			{StepNumber: 2, Instruction: "Let stand for 10-15 minutes until the mixture is foamy."},
		},
		EquipmentIDs: []int64{54, 58}, // Mixing Bowl (small), Measuring Cups
	},
	{
		ID:          202,
		RecipeID:    2,
		StepOrder:   2,
		Title:       "Make Dough",
		Description: "Combine dry ingredients, add wet ingredients, mix to a shaggy dough, and rest.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 21, Quantity: 280, Unit: "g", Purpose: "structure"},                  // Bread Flour
			{FoodItemID: 16, Quantity: 5, Unit: "g", Purpose: "flavor, fermentation control"}, // Salt (1 tsp)
			{FoodItemID: 24, Quantity: 8, Unit: "ml", Purpose: "tenderness, flavor"},          // Olive Oil (~1.5 tsp, assuming density ~0.92g/ml)
			// Activated yeast mixture from Step 1 is added here
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In a large mixing bowl, whisk together bread flour and salt."},
			{StepNumber: 2, Instruction: "Add the foamy yeast mixture and olive oil to the flour mixture."},
			{StepNumber: 3, Instruction: "Mix with a fork or your hands until just combined and no dry flour remains. The dough will be shaggy."},
			{StepNumber: 4, Instruction: "Cover the bowl and let the dough rest for 10 minutes (autolyse)."},
		},
		EquipmentIDs: []int64{54, 59}, // Mixing Bowl (large), Fork
	},
	{
		ID:          203,
		RecipeID:    2,
		StepOrder:   3,
		Title:       "Develop Dough & Final Proof",
		Description: "Perform stretch and folds, divide, shape, and let rise.",
		Notes:       "Total development time approx. 2 hours.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{}, // No new ingredients added in this step
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Perform the first set of stretch and folds. Cover and rest for 30 minutes."},
			{StepNumber: 2, Instruction: "Perform the second set of stretch and folds. Cover and rest for 30 minutes."},
			{StepNumber: 3, Instruction: "Perform the third set of stretch and folds. Cover and rest for 30 minutes."},
			{StepNumber: 4, Instruction: "Perform a final set of stretch and folds if the dough still feels slack. Cover and rest for 30 minutes."},
			{StepNumber: 5, Instruction: "Calculate total dough weight (should be around 490g). Divide into two equal portions (approx. 245g each)."},
			{StepNumber: 6, Instruction: "Shape each portion into a smooth ball using the edge-folding technique."},
			{StepNumber: 7, Instruction: "Place dough balls on a lightly floured surface or in separate containers, cover, and let rest at room temperature until ready to use (part of the 4-5 hour total time)."},
		},
		EquipmentIDs: []int64{54}, // Mixing Bowl
	},
	{
		ID:          204,
		RecipeID:    2,
		StepOrder:   4,
		Title:       "Stretch and Cook Pizza",
		Description: "Stretch dough, top, cook on stovetop in cast iron, finish in oven.",
		Notes:       "Prepare toppings while dough completes final rise. Rest dough if it resists stretching.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Pizza toppings (sauce, cheese, etc.) are added here but not listed as they vary. Assume use of Recipe ID 15 (Quick Pizza Sauce) and cheese (ID 11/12).
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Preheat oven to 250°C (fan bake setting)."},
			{StepNumber: 2, Instruction: "Place a cold cast iron pan on the stovetop. Lightly spray with cooking spray (optional)."},
			{StepNumber: 3, Instruction: "Gently stretch one dough ball into a 12-14 inch round. Place in the cold pan."},
			{StepNumber: 4, Instruction: "Top the dough with sauce and desired ingredients."},
			{StepNumber: 5, Instruction: "Place the pan over medium-high heat on the stovetop."},
			{StepNumber: 6, Instruction: "Cook for 5-7 minutes, watching for slight bubbling on the top surface of the crust."},
			{StepNumber: 7, Instruction: "Carefully transfer the hot cast iron pan to the preheated oven."},
			{StepNumber: 8, Instruction: "Bake for 10-12 minutes initially."},
			{StepNumber: 9, Instruction: "Check doneness at 10 minutes: cheese should be fully melted and browning, crust edge golden brown, bottom crispy (lift edge with spatula)."},
			{StepNumber: 10, Instruction: "If needed, bake for additional 2-minute intervals until done."},
			{StepNumber: 11, Instruction: "Carefully remove the pan from the oven. Let the pizza rest in the pan for 2-3 minutes before sliding onto a cutting board and slicing."},
			{StepNumber: 12, Instruction: "Repeat for the second dough ball."},
		},
		EquipmentIDs: []int64{60, 62, 61, 63, 56}, // Cast Iron Pan, Stovetop, Oven, Spatula, Cutting Board
	},

	// --- Recipe 3: Butter Chicken ---
	{
		ID:          301,
		RecipeID:    3,
		StepOrder:   1,
		Title:       "Marinate Chicken",
		Description: "Combine marinade ingredients and coat chicken. Refrigerate.",
		Notes:       "Can be done up to 24 hours ahead. Toast fenugreek first.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 36, Quantity: 1.5, Unit: "tsp", Purpose: "herb"},          // Fenugreek Leaves (for toasting)
			{FoodItemID: 26, Quantity: 500, Unit: "g", Purpose: "protein"},         // Chicken Thighs, cut into 2.5cm pieces
			{FoodItemID: 27, Quantity: 30, Unit: "ml", Purpose: "base marinade"},   // Greek Yogurt
			{FoodItemID: 30, Quantity: 8, Unit: "g", Purpose: "base marinade"},     // Ginger, grated (~2cm)
			{FoodItemID: 16, Quantity: 5, Unit: "g", Purpose: "base marinade"},     // Kosher Salt
			{FoodItemID: 37, Quantity: 1.5, Unit: "tsp", Purpose: "base marinade"}, // Garam Masala
			// Fenugreek from above is ground and added here
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Toast 1.5 tsp fenugreek leaves in a dry pan until fragrant (about 30 seconds). Let cool and grind."},
			{StepNumber: 2, Instruction: "In a bowl, mix Greek yogurt, grated ginger, salt, 1.5 tsp garam masala, and the ground toasted fenugreek."},
			{StepNumber: 3, Instruction: "Add the chicken pieces and coat thoroughly."},
			{StepNumber: 4, Instruction: "Cover the bowl and refrigerate for at least 1 hour, or up to 24 hours."},
		},
		EquipmentIDs: []int64{60, 54, 51}, // Cast Iron Pan (for toasting), Mixing Bowl, Grater
	},
	{
		ID:          302,
		RecipeID:    3,
		StepOrder:   2,
		Title:       "Prepare Sauce Components",
		Description: "Toast and grind spices, soak cashews, prep aromatics.",
		Notes:       "Start 1 hour 15 minutes ahead.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 36, Quantity: 2.5, Unit: "tsp", Purpose: "spice blend"}, // Fenugreek Leaves
			{FoodItemID: 38, Quantity: 1, Unit: "unit", Purpose: "spice blend"},  // Dried Chili
			{FoodItemID: 39, Quantity: 1, Unit: "unit", Purpose: "spice blend"},  // Black Cardamom Pod (or 2 green)
			{FoodItemID: 40, Quantity: 1, Unit: "unit", Purpose: "spice blend"},  // Clove
			{FoodItemID: 37, Quantity: 7.5, Unit: "ml", Purpose: "spice blend"},  // Garam Masala
			{FoodItemID: 32, Quantity: 15, Unit: "g", Purpose: "thickener"},      // Cashews
			{FoodItemID: 25, Quantity: 60, Unit: "ml", Purpose: "soaking"},       // Hot Water
			{FoodItemID: 29, Quantity: 75, Unit: "g", Purpose: "aromatic"},       // Onion (~1/2 medium), diced
			{FoodItemID: 30, Quantity: 8, Unit: "g", Purpose: "aromatic"},        // Ginger (~2cm), prepared (minced/grated)
			{FoodItemID: 31, Quantity: 10, Unit: "g", Purpose: "aromatic"},       // Garlic (~2 cloves), prepared (minced)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Toast fenugreek leaves, dried chili, cardamom pod, and clove in a dry pan until fragrant. Let cool."},
			{StepNumber: 2, Instruction: "Grind the toasted spices along with the remaining 1.5 tsp garam masala."},
			{StepNumber: 3, Instruction: "Place cashews in a small bowl and cover with 60ml hot water. Let soak."},
			{StepNumber: 4, Instruction: "Dice the onion. Prepare (mince or grate) the ginger and garlic."},
		},
		EquipmentIDs: []int64{60, 54, 55, 56}, // Cast Iron Pan, Small Bowl, Knife, Chopping Board (Spice grinder assumed)
	},
	{
		ID:          303,
		RecipeID:    3,
		StepOrder:   3,
		Title:       "Make Tomato Cashew Sauce",
		Description: "Brown onions, cook aromatics and spices, add tomatoes and cashews, simmer.",
		Notes:       "Start 1 hour ahead. Onions take 14-17 mins.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 34, Quantity: 30, Unit: "ml", Purpose: "cooking fat"}, // Canola Oil
			// Prepared onion from Step 2
			{FoodItemID: 35, Quantity: 0.125, Unit: "tsp", Purpose: "browning aid"}, // Baking Soda
			// Prepared garlic and ginger from Step 2
			// Ground spice mixture from Step 2
			// Soaked cashews and their water from Step 2
			{FoodItemID: 33, Quantity: 400, Unit: "g", Purpose: "sauce base"}, // Canned Whole Tomatoes
			{FoodItemID: 25, Quantity: 150, Unit: "ml", Purpose: "liquid"},    // Water
			{FoodItemID: 16, Quantity: 3, Unit: "g", Purpose: "seasoning"},    // Salt (to taste)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat canola oil in a large saucepan over medium-high heat."},
			{StepNumber: 2, Instruction: "Add diced onions and baking soda. Cook, stirring occasionally, until deeply browned (14-17 minutes). Adjust heat as needed."},
			{StepNumber: 3, Instruction: "Add prepared ginger and garlic, cook for 1 minute until fragrant."},
			{StepNumber: 4, Instruction: "Add the ground spice mixture, stir constantly for 30 seconds."},
			{StepNumber: 5, Instruction: "Add the soaked cashews along with their soaking water, the canned whole tomatoes (break them up slightly), and 150ml water."},
			{StepNumber: 6, Instruction: "Bring to a simmer, then reduce heat to low, cover partially, and simmer gently for 40 minutes, or until the sauce has reduced and thickened."},
			{StepNumber: 7, Instruction: "Season with salt to taste during simmering."},
		},
		EquipmentIDs: []int64{50, 63}, // Large Saucepan, Spatula/Spoon
	},
	{
		ID:          304,
		RecipeID:    3,
		StepOrder:   4,
		Title:       "Cook Chicken",
		Description: "Broil marinated chicken until cooked through.",
		Notes:       "Start 15 minutes ahead. Target internal temp 75°C.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Marinated chicken from Step 1
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Preheat toaster oven or main oven grill/broiler to high."},
			{StepNumber: 2, Instruction: "Arrange marinated chicken pieces in a single layer on a suitable baking tray or dish."},
			{StepNumber: 3, Instruction: "Broil/grill for 12-15 minutes, turning halfway, until chicken is cooked through and lightly charred (internal temperature reaches 75°C). Check frequently to prevent burning."},
		},
		EquipmentIDs: []int64{69, 66}, // Toaster Oven / Grill, Baking Tray
	},
	{
		ID:          305,
		RecipeID:    3,
		StepOrder:   5,
		Title:       "Finish Sauce and Assemble",
		Description: "Blend sauce, stir in cream and butter, fold in chicken.",
		Notes:       "Start 5 minutes ahead.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Simmered sauce from Step 3
			{FoodItemID: 28, Quantity: 60, Unit: "ml", Purpose: "creaminess"}, // Heavy Cream
			{FoodItemID: 13, Quantity: 30, Unit: "g", Purpose: "richness"},    // Butter
			// Cooked chicken from Step 4
			{FoodItemID: 41, Quantity: 180, Unit: "g", IsOptional: true, Purpose: "serving side"}, // Basmati Rice (60g dry per person for 3 servings)
			{FoodItemID: 28, Quantity: 15, Unit: "ml", IsOptional: true, Purpose: "garnish"},      // Heavy Cream (for drizzle)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Carefully transfer the hot sauce mixture to a blender (or use an immersion blender in the saucepan). Blend until completely smooth."},
			{StepNumber: 2, Instruction: "Return the blended sauce to the saucepan over low heat."},
			{StepNumber: 3, Instruction: "Stir in the heavy cream and butter until the butter is melted and incorporated."},
			{StepNumber: 4, Instruction: "Gently fold in the cooked chicken pieces."},
			{StepNumber: 5, Instruction: "Adjust seasoning (salt) if necessary."},
			{StepNumber: 6, Instruction: "Serve hot, garnished with an optional drizzle of cream, alongside cooked Basmati rice."},
		},
		EquipmentIDs: []int64{64, 50, 63, 65}, // Blender (or immersion), Large Saucepan, Spatula/Spoon, Small Saucepan (for rice)
	},
	// ... Continue for all other recipes: Banh Mi, Wedges, Burrito, Muffins, Fettuccine, Taco Noodles, Ginger-Peach Sauce, Coleslaw, Satay Noodles, Pambazo, Pickled Onions, Pizza Sauce, Red Taco Sauce, Baguette ...
	// Example structure for a Component Recipe (Quick Pickled Onions - ID 14)
	{
		ID:          1401, // Assign appropriate unique ID
		RecipeID:    14,   // Link to the Pickled Onion Recipe ID
		StepOrder:   1,
		Title:       "Prepare and Pickle Onions",
		Description: "Slice onion, heat brine, combine, and let cool.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 29, Quantity: 150, Unit: "g", Purpose: "main ingredient"}, // Red Onion (use correct ID if different from yellow)
			{FoodItemID: 58, Quantity: 125, Unit: "ml", Purpose: "brine acid"},     // White Vinegar
			{FoodItemID: 25, Quantity: 125, Unit: "ml", Purpose: "brine liquid"},   // Water
			{FoodItemID: 16, Quantity: 5, Unit: "g", Purpose: "brine seasoning"},   // Salt
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Slice the red onion very thinly (2-3mm) from root to tip."},
			{StepNumber: 2, Instruction: "Pack the sliced onions tightly into a clean jar (~200ml capacity)."},
			{StepNumber: 3, Instruction: "In a small saucepan, combine white vinegar, water, and salt."},
			{StepNumber: 4, Instruction: "Bring the brine mixture to a boil over medium-high heat, stirring until salt dissolves."},
			{StepNumber: 5, Instruction: "Carefully pour the hot brine over the onions in the jar, ensuring they are fully submerged. Tap the jar gently to release air bubbles."},
			{StepNumber: 6, Instruction: "Let cool completely at room temperature."},
			{StepNumber: 7, Instruction: "Seal the jar and refrigerate for at least 2 hours before using (best after 24 hours)."},
		},
		EquipmentIDs: []int64{55, 56, 71, 65}, // Knife, Chopping Board, Jar, Small Saucepan
	},
	// --- Recipe 4: Banh Mi with Duck ---
	{
		ID:          401,
		RecipeID:    4,
		StepOrder:   1,
		Title:       "Quick Pickle Vegetables",
		Description: "Prepare daikon and carrots and pickle them in a quick brine.",
		Notes:       "This step references the method for Recipe ID 14 (Quick Pickled Red Onions), but uses daikon/carrot. Pickle minimum 30 mins.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 49, Quantity: 150, Unit: "g", Purpose: "pickle veg"}, // Daikon Radish (approx 1 medium)
			{FoodItemID: 50, Quantity: 150, Unit: "g", Purpose: "pickle veg"}, // Carrot (approx 2 medium)
			{FoodItemID: 58, Quantity: 240, Unit: "ml", Purpose: "brine"},     // White Vinegar (1 cup)
			{FoodItemID: 25, Quantity: 240, Unit: "ml", Purpose: "brine"},     // Water (1 cup)
			{FoodItemID: 96, Quantity: 30, Unit: "g", Purpose: "brine"},       // Sugar (2 tbsp)
			{FoodItemID: 16, Quantity: 15, Unit: "g", Purpose: "brine"},       // Salt (1 tbsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Peel and julienne the daikon radish and carrots."},
			{StepNumber: 2, Instruction: "In a small saucepan, combine vinegar, water, sugar, and salt. Heat gently, stirring, until sugar and salt dissolve."},
			{StepNumber: 3, Instruction: "Remove brine from heat and let cool slightly."},
			{StepNumber: 4, Instruction: "Place julienned vegetables in a jar or bowl and pour the warm brine over them."},
			{StepNumber: 5, Instruction: "Ensure vegetables are submerged. Let pickle for at least 30 minutes at room temperature (or longer in the fridge)."},
		},
		EquipmentIDs: []int64{68, 55, 56, 65, 71, 54}, // Peeler, Knife, Chopping Board, Small Saucepan, Jar or Bowl
	},
	{
		ID:          402,
		RecipeID:    4,
		StepOrder:   2,
		Title:       "Prepare Other Components",
		Description: "Prepare duck, slice fresh ingredients, mix sauce.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 48, Quantity: 300, Unit: "g", Purpose: "protein"},                       // Duck Breast
			{FoodItemID: 16, Quantity: 2, Unit: "g", Purpose: "seasoning"},                       // Salt
			{FoodItemID: 62, Quantity: 1, Unit: "g", Purpose: "seasoning"},                       // Black Pepper
			{FoodItemID: 51, Quantity: 100, Unit: "g", Purpose: "fresh veg"},                     // Cucumber (sliced into spears)
			{FoodItemID: 52, Quantity: 3, Unit: "unit", Purpose: "heat"},                         // Bird Eye Chilies (thinly sliced)
			{FoodItemID: 53, Quantity: 30, Unit: "g", Purpose: "aromatic"},                       // Spring Onions (thinly sliced)
			{FoodItemID: 46, Quantity: 10, Unit: "g", Purpose: "herb"},                           // Cilantro (fresh sprigs/leaves)
			{FoodItemID: 54, Quantity: 60, Unit: "g", Purpose: "sauce base"},                     // Mayonnaise (4 tbsp)
			{FoodItemID: 55, Quantity: 15, Unit: "ml", Purpose: "sauce flavor"},                  // Soy Sauce (1 tbsp)
			{FoodItemID: 56, Quantity: 10, Unit: "ml", Purpose: "sauce heat"},                    // Tabasco (1-2 tsp)
			{FoodItemID: 57, Quantity: 5, Unit: "g", IsOptional: true, Purpose: "sauce balance"}, // Honey (1/2 tsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Pat the duck breast dry. Score the skin in a crosshatch pattern, being careful not to cut into the meat. Season skin generously with salt and pepper."},
			{StepNumber: 2, Instruction: "Slice the cucumber into spears."},
			{StepNumber: 3, Instruction: "Thinly slice the chilies (remove seeds for less heat if desired) and spring onions."},
			{StepNumber: 4, Instruction: "Wash cilantro."},
			{StepNumber: 5, Instruction: "In a small bowl, mix together mayonnaise, soy sauce, Tabasco, and optional honey until combined."},
		},
		EquipmentIDs: []int64{55, 56, 54}, // Knife, Chopping Board, Small Bowl
	},
	{
		ID:          403,
		RecipeID:    4,
		StepOrder:   3,
		Title:       "Cook Duck",
		Description: "Render duck fat and cook duck breast in cast iron pan.",
		Notes:       "Target internal temperature 63°C for medium-rare.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Seasoned duck breast from Step 2
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Place the duck breast skin-side down in a cold cast iron pan."},
			{StepNumber: 2, Instruction: "Place the pan over medium heat. Cook for 6-8 minutes, allowing the fat to render slowly. Pour off excess fat periodically if desired (reserve for toasting bread)."},
			{StepNumber: 3, Instruction: "Flip the duck breast and cook on the meat side for 4-5 minutes more for medium-rare (internal temperature 63°C). Cook longer if preferred."},
			{StepNumber: 4, Instruction: "Remove duck from pan and let rest on a cutting board for 5-10 minutes."},
			{StepNumber: 5, Instruction: "Once rested, thinly slice the duck breast against the grain."},
		},
		EquipmentIDs: []int64{60, 62, 67, 80, 56, 55}, // Cast Iron Pan, Stovetop, Tongs, Thermometer, Cutting Board, Knife
	},
	{
		ID:          404,
		RecipeID:    4,
		StepOrder:   4,
		Title:       "Assemble Banh Mi",
		Description: "Toast baguettes and layer all components.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 47, Quantity: 2, Unit: "unit", Purpose: "bread"}, // Baguettes
			// Sauce from Step 2
			// Sliced duck from Step 3
			// Pickled vegetables from Step 1 (drained)
			// Cucumber spears from Step 2
			// Sliced chilies from Step 2
			// Cilantro from Step 2
			// Spring onions from Step 2
			// Optional: Reserved duck fat for toasting
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Split the baguettes lengthwise, leaving one side hinged if desired."},
			{StepNumber: 2, Instruction: "Toast the baguettes lightly (in oven, toaster, or pan. Optional: brush cut sides with reserved duck fat before toasting)."},
			{StepNumber: 3, Instruction: "Spread the prepared mayonnaise sauce mixture generously on the inside of the baguettes."},
			{StepNumber: 4, Instruction: "Layer the sliced duck breast."},
			{StepNumber: 5, Instruction: "Add a generous amount of drained pickled vegetables."},
			{StepNumber: 6, Instruction: "Add cucumber spears."},
			{StepNumber: 7, Instruction: "Top with sliced chilies, fresh cilantro sprigs, and sliced spring onions."},
			{StepNumber: 8, Instruction: "Close the sandwich and serve immediately."},
		},
		EquipmentIDs: []int64{55, 61, 69, 59}, // Knife, Oven or Toaster, Spreader/Spoon
	},

	// --- Recipe 5: Agria Potato Wedges ---
	{
		ID:          501,
		RecipeID:    5,
		StepOrder:   1,
		Title:       "Prepare Potatoes",
		Description: "Wash, cut, soak, and dry potato wedges.",
		Notes:       "Soaking removes excess starch for crispiness. Dry thoroughly.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 59, Quantity: 400, Unit: "g", Purpose: "main"},      // Agria Potatoes (approx 2 medium)
			{FoodItemID: 25, Quantity: 1000, Unit: "ml", Purpose: "soaking"}, // Cold Water
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Preheat air fryer to 200°C or oven to 230°C (fan forced setting). If using oven, place a frying rack on a baking tray."},
			{StepNumber: 2, Instruction: "Wash and scrub the Agria potatoes thoroughly (leave skin on)."},
			{StepNumber: 3, Instruction: "Cut potatoes in half lengthwise, then cut each half into 3-4 wedges of similar thickness."},
			{StepNumber: 4, Instruction: "Place the cut wedges in a large bowl and cover with cold water. Let soak for at least 5 minutes."},
			{StepNumber: 5, Instruction: "Drain the potatoes well and pat them completely dry with paper towels or a clean kitchen towel."},
		},
		EquipmentIDs: []int64{55, 56, 54, 66, 74, 75}, // Knife, Chopping Board, Bowl, Baking Tray, Wire Rack (optional), Paper Towels
	},
	{
		ID:          502,
		RecipeID:    5,
		StepOrder:   2,
		Title:       "Season and Cook Wedges",
		Description: "Toss wedges with oil and seasonings, then bake/air fry until crispy.",
		Notes:       "User notes mention adding cayenne and using canola oil.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Dried potato wedges from Step 1
			{FoodItemID: 60, Quantity: 30, Unit: "ml", Purpose: "coating"},                 // Avocado Oil (2 tbsp) (User substituted Canola Oil ID 34)
			{FoodItemID: 18, Quantity: 5, Unit: "g", Purpose: "seasoning"},                 // Smoked Paprika (1 tsp)
			{FoodItemID: 61, Quantity: 2.5, Unit: "g", Purpose: "seasoning"},               // Garlic Powder (1/2 tsp)
			{FoodItemID: 16, Quantity: 2.5, Unit: "g", Purpose: "seasoning"},               // Salt (1/2 tsp)
			{FoodItemID: 62, Quantity: 1, Unit: "g", Purpose: "seasoning"},                 // Black Pepper (1/4 tsp)
			{FoodItemID: 63, Quantity: 5, Unit: "g", Purpose: "seasoning"},                 // Mixed Herbs (1 tsp)
			{FoodItemID: 64, Quantity: 1.25, Unit: "g", IsOptional: true, Purpose: "heat"}, // Cayenne Pepper (User added 1/2 tsp)
			{FoodItemID: 65, Quantity: 15, Unit: "g", IsOptional: true, Purpose: "flavor"}, // Parmesan, finely grated (Optional enhancement, add last 5 mins)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In a large bowl, combine the oil (avocado or canola), smoked paprika, garlic powder, salt, pepper, mixed herbs, and optional cayenne pepper."},
			{StepNumber: 2, Instruction: "Add the thoroughly dried potato wedges to the bowl and toss gently until evenly coated with the seasoning mixture."},
			{StepNumber: 3, Instruction: "Arrange the seasoned wedges in a single layer on the prepared air fryer basket or baking tray with rack. Ensure they are not overcrowded."},
			{StepNumber: 4, Instruction: "Air fry at 200°C or bake at 220-230°C (fan forced) for 25-30 minutes, flipping halfway through baking (air fryer may not need flipping)."},
			{StepNumber: 5, Instruction: "If adding optional parmesan, sprinkle it over the wedges during the last 5 minutes of cooking."},
			{StepNumber: 6, Instruction: "Wedges are done when golden brown, crispy on the outside, and fluffy on the inside."},
			{StepNumber: 7, Instruction: "Serve immediately."},
		},
		EquipmentIDs: []int64{54, 59, 61, 66, 74}, // Bowl, Spoon/Hands, Oven or Air Fryer, Baking Tray, Wire Rack
	},

	// --- Recipe 6: Crispy Mexican Chicken Burrito ---
	{
		ID:          601,
		RecipeID:    6,
		StepOrder:   1,
		Title:       "Advance Preparation",
		Description: "Cook corn and heat chicken if necessary.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 44, Quantity: 45, Unit: "g", Purpose: "filling"},        // Corn Kernels (1/4 cup)
			{FoodItemID: 25, Quantity: 250, Unit: "ml", Purpose: "cooking corn"}, // Water
			{FoodItemID: 20, Quantity: 50, Unit: "g", Purpose: "filling"},        // Mexican Chicken (ensure pre-cooked)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "If corn is not already cooked: Bring water to a boil in a small saucepan. Add corn, cover, cook for 5 minutes. Drain and set aside."},
			{StepNumber: 2, Instruction: "If chicken is cold: Preheat toaster oven to 180°C. Heat chicken for 10 minutes until hot. Let rest 2 minutes."},
		},
		EquipmentIDs: []int64{65, 62, 69, 66}, // Small Saucepan, Stovetop, Toaster Oven, Small Baking Tray
	},
	{
		ID:          602,
		RecipeID:    6,
		StepOrder:   2,
		Title:       "Prepare Filling Ingredients",
		Description: "Grate cheese, dice tomato, tear lettuce.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 11, Quantity: 50, Unit: "g", Purpose: "filling"},  // Everyday Cheese, grated
			{FoodItemID: 69, Quantity: 60, Unit: "g", Purpose: "filling"},  // Tomato (2 slices), diced
			{FoodItemID: 67, Quantity: 20, Unit: "g", Purpose: "filling"},  // Lettuce Leaf (1), torn
			{FoodItemID: 68, Quantity: 25, Unit: "g", Purpose: "filling"},  // Pickled Onions (small handful)
			{FoodItemID: 88, Quantity: 15, Unit: "ml", Purpose: "filling"}, // Mexican Spicy Sauce (3-4 passes)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Grate the cheese."},
			{StepNumber: 2, Instruction: "Dice the tomato slices."},
			{StepNumber: 3, Instruction: "Tear the lettuce leaf into manageable pieces."},
			{StepNumber: 4, Instruction: "Have pickled onions and spicy sauce ready."},
		},
		EquipmentIDs: []int64{51, 55, 56}, // Grater, Knife, Chopping Board
	},
	{
		ID:          603,
		RecipeID:    6,
		StepOrder:   3,
		Title:       "Assemble and Roll Burrito",
		Description: "Warm tortilla, layer ingredients, and roll tightly.",
		Notes:       "Layering order is important.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 66, Quantity: 1, Unit: "unit", Purpose: "wrap"}, // Large Flour Tortilla
			// All prepared fillings from Steps 1 & 2
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Warm the flour tortilla in a microwave for 20 seconds until pliable."},
			{StepNumber: 2, Instruction: "Lay the warm tortilla flat. Layer ingredients in the center, leaving a border:"},
			{StepNumber: 3, Instruction: "Layer 1: Grated cheese."},
			{StepNumber: 4, Instruction: "Layer 2: Torn lettuce."},
			{StepNumber: 5, Instruction: "Layer 3: Diced tomato."},
			{StepNumber: 6, Instruction: "Layer 4: Pickled onions."},
			{StepNumber: 7, Instruction: "Layer 5: Cooked corn."},
			{StepNumber: 8, Instruction: "Layer 6: Warmed Mexican chicken."},
			{StepNumber: 9, Instruction: "Layer 7: Drizzle Mexican spicy sauce over the filling (3-4 passes)."},
			{StepNumber: 10, Instruction: "Fold the sides of the tortilla inwards towards the middle, over the filling."},
			{StepNumber: 11, Instruction: "Lift the bottom edge (closest to you) up and over the filling, holding the sides in."},
			{StepNumber: 12, Instruction: "Tuck the filling in tightly as you roll the burrito away from you, ensuring the ends remain tucked and sealed. Position seam-side down."},
		},
		EquipmentIDs: []int64{82}, // Microwave
	},
	{
		ID:          604,
		RecipeID:    6,
		StepOrder:   4,
		Title:       "Toast Burrito",
		Description: "Cook the rolled burrito in a sandwich press until crispy.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Assembled burrito from Step 3
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Preheat a sandwich press or toastie machine."},
			{StepNumber: 2, Instruction: "Carefully place the rolled burrito seam-side down onto the hot press."},
			{StepNumber: 3, Instruction: "Close the lid and toast for 2-4 minutes, or until the tortilla is golden brown and crispy."},
			{StepNumber: 4, Instruction: "Remove carefully and serve immediately."},
		},
		EquipmentIDs: []int64{72, 63}, // Sandwich Press, Spatula
	},

	// --- Recipe 7: English Muffins ---
	{
		ID:          701,
		RecipeID:    7,
		StepOrder:   1,
		Title:       "Activate Yeast & Make Dough",
		Description: "Activate yeast, mix ingredients, use stretch and fold.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 75, Quantity: 45, Unit: "ml", Purpose: "liquid, yeast activation"}, // Trim Milk (3 tbsp)
			{FoodItemID: 25, Quantity: 30, Unit: "ml", Purpose: "liquid, yeast activation"}, // Water (2 tbsp)
			{FoodItemID: 96, Quantity: 7.5, Unit: "g", Purpose: "yeast food"},               // Granulated Sugar (1.5 tsp)
			{FoodItemID: 22, Quantity: 2.5, Unit: "g", Purpose: "leavening"},                // Active Dried Yeast (1/2 tsp)
			{FoodItemID: 21, Quantity: 90, Unit: "g", Purpose: "structure"},                 // Bread Flour (3/4 cup)
			{FoodItemID: 16, Quantity: 1.25, Unit: "g", Purpose: "flavor"},                  // Salt (1/4 tsp)
			{FoodItemID: 74, Quantity: 0.25, Unit: "unit", Purpose: "enrichment"},           // Egg, beaten (~1 tbsp)
			{FoodItemID: 13, Quantity: 10, Unit: "g", Purpose: "enrichment, tenderness"},    // Butter, melted (2 tsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Warm milk and water slightly (lukewarm). Stir in sugar and yeast. Let stand 5-10 minutes until foamy."},
			{StepNumber: 2, Instruction: "In a bowl, combine bread flour and salt."},
			{StepNumber: 3, Instruction: "Add the yeast mixture, beaten egg portion, and melted butter to the flour."},
			{StepNumber: 4, Instruction: "Mix until a shaggy dough forms. Cover and rest 10 minutes."},
			{StepNumber: 5, Instruction: "Perform 1-2 sets of stretch and folds instead of kneading (wet hands slightly if needed)."},
		},
		EquipmentIDs: []int64{58, 54, 59}, // Measuring Cups/Spoons, Bowl, Fork/Spoon
	},
	{
		ID:          702,
		RecipeID:    7,
		StepOrder:   2,
		Title:       "First Rise",
		Description: "Let the dough rise until doubled.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Cover the bowl tightly (e.g., with plastic wrap) and let the dough rise in a warm place until doubled in size, about 1 hour."},
		},
		EquipmentIDs: []int64{54, 78}, // Bowl, Plastic Wrap
	},
	{
		ID:          703,
		RecipeID:    7,
		StepOrder:   3,
		Title:       "Shape Muffins",
		Description: "Gently pat out dough on floured surface and cut rounds.",
		Notes:       "Dough is sticky; use generous flour or oiled hands/scraper. Can also portion and flatten balls.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 14, Quantity: 30, Unit: "g", Purpose: "dusting, shaping"}, // All-Purpose Flour for dusting
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Generously flour a clean work surface (chopping board). Also flour your hands well."},
			{StepNumber: 2, Instruction: "Gently scrape the risen dough onto the floured surface. Dust the top liberally with more flour."},
			{StepNumber: 3, Instruction: "Very gently pat the dough down to about 2.5cm (1 inch) thickness."},
			{StepNumber: 4, Instruction: "Flour the rim of a ramekin or round cutter (approx. 7-8cm diameter)."},
			{StepNumber: 5, Instruction: "Cut out 3-4 muffins, trying not to twist the cutter. Re-flour cutter as needed."},
			{StepNumber: 6, Instruction: "Alternative: If too sticky, flour hands well, divide dough into 3-4 equal portions, and gently flatten each into a round shape."},
		},
		EquipmentIDs: []int64{56, 70}, // Chopping Board/Work Surface, Ramekin/Cutter
	},
	{
		ID:          704,
		RecipeID:    7,
		StepOrder:   4,
		Title:       "Second Rise",
		Description: "Let shaped muffins rise on floured parchment.",
		Notes:       "Flour substitutes for cornmeal/semolina.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 14, Quantity: 20, Unit: "g", Purpose: "dusting"}, // All-Purpose Flour for dusting parchment
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Line a small baking tray with parchment paper."},
			{StepNumber: 2, Instruction: "Dust the parchment paper generously with flour."},
			{StepNumber: 3, Instruction: "Carefully transfer the shaped muffins to the floured parchment, leaving space between them."},
			{StepNumber: 4, Instruction: "Dust the tops of the muffins lightly with more flour."},
			{StepNumber: 5, Instruction: "Cover loosely (e.g., with lightly oiled plastic wrap or a clean tea towel) and let rise for 30 minutes."},
		},
		EquipmentIDs: []int64{66, 77, 78}, // Baking Tray, Parchment Paper, Plastic Wrap/Towel
	},
	{
		ID:          705,
		RecipeID:    7,
		StepOrder:   5,
		Title:       "Bake Muffins",
		Description: "Bake in oven, flipping halfway.",
		Notes:       "Target internal temp 90°C or hollow sound when tapped.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Preheat oven to 180°C (350°F), using convection/fan bake if available."},
			{StepNumber: 2, Instruction: "Place the baking tray with the risen muffins into the preheated oven."},
			{StepNumber: 3, Instruction: "Bake for 7-8 minutes."},
			{StepNumber: 4, Instruction: "Carefully remove the tray, flip the muffins over using a spatula."},
			{StepNumber: 5, Instruction: "Return to the oven and bake for another 7-8 minutes."},
			{StepNumber: 6, Instruction: "Muffins are done when golden brown and sound hollow when tapped on the bottom (or reach 90°C internal temperature)."},
			{StepNumber: 7, Instruction: "Transfer baked muffins to a wire rack to cool completely before splitting and toasting."},
		},
		EquipmentIDs: []int64{61, 66, 63, 80, 74}, // Oven, Baking Tray, Spatula, Thermometer, Wire Rack
	},

	// --- Recipe 8: Fettuccine with Mushroom White Wine Sauce ---
	// NOTE: This merges notes from different dates/variations. Base is first attempt.
	{
		ID:          801,
		RecipeID:    8,
		StepOrder:   1,
		Title:       "Mise en Place (Prep)",
		Description: "Prepare all vegetables, garlic, cheese, and have liquids ready.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 77, Quantity: 100, Unit: "g", Purpose: "main veg"},                // Mushrooms, sliced
			{FoodItemID: 29, Quantity: 75, Unit: "g", Purpose: "mirepoix"},                 // Onion (1/2), finely diced
			{FoodItemID: 78, Quantity: 50, Unit: "g", Purpose: "mirepoix"},                 // Celery (1 stalk), finely diced
			{FoodItemID: 50, Quantity: 60, Unit: "g", Purpose: "mirepoix"},                 // Carrot (1), finely diced
			{FoodItemID: 31, Quantity: 15, Unit: "g", Purpose: "aromatic"},                 // Garlic (2-3 cloves), minced
			{FoodItemID: 65, Quantity: 50, Unit: "g", Purpose: "finishing"},                // Parmesan (1/2 cup), grated (Parmigiano Reggiano preferred)
			{FoodItemID: 79, Quantity: 120, Unit: "ml", Purpose: "deglazing, flavor"},      // White Wine (1/2 cup) (Sub: 2 tbsp vinegar ID 58 + 1/4 cup water ID 25)
			{FoodItemID: 80, Quantity: 240, Unit: "ml", Purpose: "sauce liquid"},           // Chicken Stock (1 cup)
			{FoodItemID: 13, Quantity: 50, Unit: "g", Purpose: "cooking fat, finishing"},   // Butter (3-4 tbsp), divided
			{FoodItemID: 24, Quantity: 15, Unit: "ml", Purpose: "cooking fat"},             // Olive Oil (1 tbsp)
			{FoodItemID: 76, Quantity: 150, Unit: "g", Purpose: "pasta base"},              // Fettuccine (User notes: homemade wide, or store-bought 150g)
			{FoodItemID: 16, Quantity: 5, Unit: "g", Purpose: "seasoning"},                 // Salt (to taste)
			{FoodItemID: 62, Quantity: 2, Unit: "g", Purpose: "seasoning"},                 // Pepper (to taste, User note: added later)
			{FoodItemID: 55, Quantity: 15, Unit: "ml", IsOptional: true, Purpose: "umami"}, // Soy Sauce (1 tbsp, User note added later)
			{FoodItemID: 17, Quantity: 1, Unit: "g", IsOptional: true, Purpose: "flavor"},  // Mustard Powder (User note added later)
			{FoodItemID: 18, Quantity: 1, Unit: "g", IsOptional: true, Purpose: "flavor"},  // Smoked Paprika (User note added later)
			// Salad ingredients listed in notes but not part of main recipe steps.
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Finely dice onion, celery, and carrot (mirepoix)."},
			{StepNumber: 2, Instruction: "Slice mushrooms."},
			{StepNumber: 3, Instruction: "Mince garlic."},
			{StepNumber: 4, Instruction: "Grate parmesan cheese."},
			{StepNumber: 5, Instruction: "Measure out wine/vinegar+water, stock, olive oil, butter, and optional soy sauce/spices."},
			{StepNumber: 6, Instruction: "Start bringing a large pot of salted water to boil for the pasta."},
		},
		EquipmentIDs: []int64{55, 56, 51, 58, 50}, // Knife, Chopping Board, Grater, Measuring tools, Large Saucepan (for pasta)
	},
	{
		ID:          802,
		RecipeID:    8,
		StepOrder:   2,
		Title:       "Cook Sauce Base",
		Description: "Sauté mirepoix and mushrooms, add garlic and wine, reduce.",
		Notes:       "User note: taste and salt at garlic stage.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Mirepoix, mushrooms, garlic from Step 1
			// Butter (2 tbsp / ~30g), Olive oil, Wine/Vinegar+Water from Step 1
			// Salt from Step 1
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat a large skillet or cast iron pan over medium-high heat. Add 2 tbsp butter and 1 tbsp olive oil."},
			{StepNumber: 2, Instruction: "Once butter melts, add the diced mirepoix. Sauté for 5-7 minutes until softened."},
			{StepNumber: 3, Instruction: "Add the sliced mushrooms. Cook, stirring occasionally, until golden brown (5-6 minutes)."},
			{StepNumber: 4, Instruction: "Add the minced garlic and cook for 1 minute until fragrant. Taste and season with salt at this point."},
			{StepNumber: 5, Instruction: "Pour in the white wine (or vinegar/water mix). Bring to a simmer and cook until reduced by about half, scraping up any browned bits from the bottom of the pan."},
		},
		EquipmentIDs: []int64{60, 62, 63}, // Skillet/Cast Iron Pan, Stovetop, Spatula/Spoon
	},
	{
		ID:          803,
		RecipeID:    8,
		StepOrder:   3,
		Title:       "Simmer Sauce & Cook Pasta",
		Description: "Add stock to sauce and simmer. Cook pasta.",
		Notes:       "Reserve pasta water.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Sauce base from Step 2
			// Chicken Stock, Fettuccine from Step 1
			// Optional soy sauce, mustard powder, paprika from Step 1 (User note: add here or earlier)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Add the chicken stock to the skillet with the reduced wine mixture. Add optional soy sauce, mustard powder, paprika if using."},
			{StepNumber: 2, Instruction: "Bring back to a simmer, then reduce heat and let the sauce gently simmer and reduce slightly while the pasta cooks."},
			{StepNumber: 3, Instruction: "Add fettuccine to the boiling salted water. Cook according to package directions until al dente."},
			{StepNumber: 4, Instruction: "Before draining the pasta, reserve about 1/2 cup (120ml) of the starchy pasta water."},
			{StepNumber: 5, Instruction: "Drain the pasta."},
		},
		EquipmentIDs: []int64{60, 50, 81, 58}, // Skillet, Large Saucepan, Slotted Spoon/Strainer, Measuring Cup
	},
	{
		ID:          804,
		RecipeID:    8,
		StepOrder:   4,
		Title:       "Finish and Serve",
		Description: "Combine pasta and sauce, add finishing ingredients.",
		Notes:       "Adjust consistency with pasta water. Season to taste.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Cooked pasta and sauce from Step 3
			// Reserved pasta water
			// Grated Parmesan, remaining Butter (1-2 tbsp / ~15-20g), Pepper from Step 1
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Add the drained pasta directly to the skillet with the simmering sauce."},
			{StepNumber: 2, Instruction: "Add the grated parmesan cheese and the remaining butter."},
			{StepNumber: 3, Instruction: "Toss everything together gently over low heat until the butter melts and the sauce coats the pasta. Add splashes of the reserved pasta water as needed to achieve desired sauce consistency (User note: might not be needed if sauce/stock was salty)."},
			{StepNumber: 4, Instruction: "Taste and adjust seasoning with salt and pepper if necessary."},
			{StepNumber: 5, Instruction: "Serve immediately, garnished with additional parmesan, fresh parsley (optional), and black pepper."},
		},
		EquipmentIDs: []int64{60, 67}, // Skillet, Tongs/Pasta Fork
	},

	// --- Recipe 9: Fusion Taco Noodles with Red Sauce ---
	{
		ID:          901,
		RecipeID:    9,
		StepOrder:   1,
		Title:       "Prep Sauce and Aromatics",
		Description: "Mix sauce components, prepare garlic, ginger, onion, spring onions.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 87, Quantity: 0.25, Unit: "unit", Purpose: "sauce base"},           // Chicken Stock Cube (1/4)
			{FoodItemID: 25, Quantity: 30, Unit: "ml", Purpose: "sauce liquid"},             // Very Hot Water (2 tbsp)
			{FoodItemID: 55, Quantity: 30, Unit: "ml", Purpose: "sauce base"},               // Soy Sauce (2 tbsp)
			{FoodItemID: 86, Quantity: 10, Unit: "g", Purpose: "sauce balance"},             // Brown Sugar (1.5-2 tsp)
			{FoodItemID: 58, Quantity: 10, Unit: "ml", Purpose: "sauce acid"},               // White Vinegar (2 tsp)
			{FoodItemID: 85, Quantity: 5, Unit: "ml", Purpose: "sauce acid"},                // Lemon or Lime Juice (1 tsp)
			{FoodItemID: 88, Quantity: 40, Unit: "ml", Purpose: "main flavor"},              // Red Taco Sauce (2-3 tbsp, Recipe ID 16)
			{FoodItemID: 89, Quantity: 2.5, Unit: "ml", IsOptional: true, Purpose: "depth"}, // Worcestershire Sauce (1/2 tsp)
			{FoodItemID: 31, Quantity: 15, Unit: "g", Purpose: "aromatic"},                  // Garlic (2-3 tsp minced)
			{FoodItemID: 30, Quantity: 5, Unit: "g", Purpose: "aromatic"},                   // Ginger Paste (1 tsp)
			{FoodItemID: 29, Quantity: 75, Unit: "g", Purpose: "aromatic"},                  // Onion (1/2), finely diced
			{FoodItemID: 53, Quantity: 30, Unit: "g", Purpose: "aromatic, garnish"},         // Spring Onions (2), whites/greens separated
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Sauce Prep: Dissolve stock cube in very hot water in a small bowl. Stir in soy sauce, brown sugar, white vinegar, and lemon/lime juice. Add red taco sauce and optional Worcestershire sauce. Stir well until combined. Taste and adjust sweetness/acidity. Keep warm if possible."},
			{StepNumber: 2, Instruction: "Aromatics Prep: Finely dice onion. Mince garlic (if not using paste). Slice spring onions, keeping white and green parts separate."},
		},
		EquipmentIDs: []int64{54, 59, 55, 56}, // Small Bowl, Spoon, Knife, Chopping Board
	},
	{
		ID:          902,
		RecipeID:    9,
		StepOrder:   2,
		Title:       "Cook Noodles and Egg",
		Description: "Cook noodles, reserve water. Cook aromatics and egg.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 90, Quantity: 150, Unit: "g", Purpose: "base"},             // Egg Noodles (1-2 servings)
			{FoodItemID: 74, Quantity: 1, Unit: "unit", Purpose: "protein, binder"}, // Egg
			{FoodItemID: 60, Quantity: 30, Unit: "ml", Purpose: "cooking fat"},      // Oil (Avocado or Canola, 2 tbsp)
			// Prepared aromatics (garlic, ginger, onion, spring onion whites) from Step 1
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Cook egg noodles according to package directions until just underdone (al dente)."},
			{StepNumber: 2, Instruction: "Before draining, reserve about 1/4 cup (60ml) of the cooking water."},
			{StepNumber: 3, Instruction: "Drain the noodles and toss with a tiny bit of oil to prevent sticking."},
			{StepNumber: 4, Instruction: "Heat a wok or large cast iron pan over medium-high heat until hot."},
			{StepNumber: 5, Instruction: "Add cooking oil and wait until it shimmers."},
			{StepNumber: 6, Instruction: "Add minced garlic, ginger paste, diced onion, and spring onion whites."},
			{StepNumber: 7, Instruction: "Stir-fry for 3-4 minutes until onions are translucent and starting to brown lightly. Adjust heat to prevent burning garlic."},
			{StepNumber: 8, Instruction: "Push the aromatics to one side of the pan. Crack the egg into the empty space."},
			{StepNumber: 9, Instruction: "Immediately start breaking up the egg with your spatula as the whites turn opaque (about 30 seconds). Keep the pieces small."},
			{StepNumber: 10, Instruction: "Mix the cooked egg pieces with the aromatics."},
		},
		EquipmentIDs: []int64{50, 81, 58, 60, 62, 63}, // Saucepan (for noodles), Strainer, Measuring Cup, Wok/Cast Iron Pan, Stovetop, Spatula
	},
	{
		ID:          903,
		RecipeID:    9,
		StepOrder:   3,
		Title:       "Combine and Finish",
		Description: "Add noodles and sauce, toss, adjust seasoning, garnish.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Cooked noodles, cooked aromatics/egg from Step 2
			// Prepared sauce, reserved noodle water from Step 1/2
			// Spring onion greens from Step 1
			// Optional: extra lemon/lime juice, Worcestershire, sugar, yogurt, soy sauce
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Add the cooked noodles and the prepared sauce to the pan with the aromatics and egg."},
			{StepNumber: 2, Instruction: "Keep heat medium-high and toss continuously until the noodles are evenly coated with the sauce and heated through (1-2 minutes)."},
			{StepNumber: 3, Instruction: "Taste and adjust seasoning: add more lemon/lime for brightness, Worcestershire for depth, sugar/yogurt if too spicy, reserved noodle water (or water+soy+sugar mix) if too dry/thick."},
			{StepNumber: 4, Instruction: "If sauce seems separated, remove from heat and stir vigorously."},
			{StepNumber: 5, Instruction: "Turn off the heat."},
			{StepNumber: 6, Instruction: "Garnish with the sliced spring onion greens."},
			{StepNumber: 7, Instruction: "Serve immediately, with an extra wedge of lemon or lime if desired."},
		},
		EquipmentIDs: []int64{60, 67, 63}, // Wok/Cast Iron Pan, Tongs, Spatula
	},

	// --- Recipe 10: Ginger-Peach Fire Sauce (Component) ---
	{
		ID:          1001,
		RecipeID:    10,
		StepOrder:   1,
		Title:       "Prepare Ingredients",
		Description: "Drain peaches, grate ginger, mince garlic, dice chilies.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 92, Quantity: 80, Unit: "g", Purpose: "fruit base"},    // Canned Peach Halves (~1/3 cup / 2 halves), drained, diced
			{FoodItemID: 30, Quantity: 25, Unit: "g", Purpose: "aromatic"},      // Fresh Ginger (3 tbsp), finely grated
			{FoodItemID: 31, Quantity: 20, Unit: "g", Purpose: "aromatic"},      // Garlic (4 cloves), minced
			{FoodItemID: 91, Quantity: 3, Unit: "unit", Purpose: "heat"},        // Fresh Hot Chilies (2-3), finely diced
			{FoodItemID: 80, Quantity: 180, Unit: "ml", Purpose: "liquid base"}, // Chicken Stock (3/4 cup)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Drain canned peaches well, pat dry, and dice finely."},
			{StepNumber: 2, Instruction: "Peel and finely grate the fresh ginger."},
			{StepNumber: 3, Instruction: "Peel and mince the garlic cloves."},
			{StepNumber: 4, Instruction: "Finely dice the fresh hot chilies (wear gloves if sensitive)."},
			{StepNumber: 5, Instruction: "Measure out chicken stock and other sauce ingredients."},
		},
		EquipmentIDs: []int64{55, 56, 51, 58}, // Knife, Chopping Board, Grater, Measuring tools
	},
	{
		ID:          1002,
		RecipeID:    10,
		StepOrder:   2,
		Title:       "Cook Sauce Base",
		Description: "Make roux, cook aromatics and peaches.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 60, Quantity: 30, Unit: "ml", Purpose: "roux fat"},      // Avocado Oil (2 tbsp)
			{FoodItemID: 14, Quantity: 15, Unit: "g", Purpose: "roux thickener"}, // All-Purpose Flour (2 tbsp)
			// Prepared ginger, garlic, chilies, peaches from Step 1
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat avocado oil in a medium saucepan over medium heat."},
			{StepNumber: 2, Instruction: "Add flour and whisk constantly until the roux turns golden brown (2-3 minutes)."},
			{StepNumber: 3, Instruction: "Add the grated ginger and minced garlic. Cook, stirring, for 1-2 minutes until fragrant."},
			{StepNumber: 4, Instruction: "Add the diced chilies and cook for another minute."},
			{StepNumber: 5, Instruction: "Stir in the diced peaches and cook for just 1 minute more."},
		},
		EquipmentIDs: []int64{50, 62, 53, 59}, // Medium Saucepan, Stovetop, Whisk, Spoon
	},
	{
		ID:          1003,
		RecipeID:    10,
		StepOrder:   3,
		Title:       "Simmer and Finish Sauce",
		Description: "Add liquids and spices, simmer, blend, strain.",
		Notes:       "Taste before adding sugar. Strain for smooth texture.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Cooked base from Step 2
			// Chicken Stock from Step 1
			{FoodItemID: 58, Quantity: 22.5, Unit: "ml", Purpose: "acid"},                  // White Vinegar (1.5 tbsp)
			{FoodItemID: 55, Quantity: 15, Unit: "ml", Purpose: "umami, salt"},             // Soy Sauce (1 tbsp)
			{FoodItemID: 93, Quantity: 2, Unit: "g", Purpose: "spice"},                     // Ground Cardamom (1/2 tsp)
			{FoodItemID: 16, Quantity: 2.5, Unit: "g", Purpose: "seasoning"},               // Salt (1/2 tsp)
			{FoodItemID: 86, Quantity: 5, Unit: "g", IsOptional: true, Purpose: "balance"}, // Brown Sugar (1/2 tsp, optional)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Gradually whisk the chicken stock into the saucepan with the cooked base until smooth."},
			{StepNumber: 2, Instruction: "Stir in the white vinegar, soy sauce, ground cardamom, and salt."},
			{StepNumber: 3, Instruction: "Bring the mixture to a simmer, then reduce heat to low and let it simmer gently, stirring occasionally, for 15-20 minutes, or until slightly thickened."},
			{StepNumber: 4, Instruction: "Taste the sauce. If needed for balance (especially if peaches were tart), stir in the brown sugar."},
			{StepNumber: 5, Instruction: "Carefully transfer the hot sauce to a blender (or use an immersion blender). Blend until completely smooth."},
			{StepNumber: 6, Instruction: "For the smoothest texture, strain the blended sauce through a fine mesh strainer into a clean bowl or container."},
			{StepNumber: 7, Instruction: "Let the sauce cool completely before storing or using."},
		},
		EquipmentIDs: []int64{50, 53, 59, 64, 76, 54}, // Medium Saucepan, Whisk, Spoon, Blender, Strainer, Bowl/Container
	},

	// --- Recipe 11: KFC Style Coleslaw ---
	{
		ID:          1101,
		RecipeID:    11,
		StepOrder:   1,
		Title:       "Prep Vegetables",
		Description: "Finely chop cabbage, shred carrot, mince onion.",
		Notes:       "Cabbage should be very fine, almost rice-sized.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 94, Quantity: 500, Unit: "g", Purpose: "base"},            // Cabbage (~1/2 head / 4 cups chopped)
			{FoodItemID: 50, Quantity: 30, Unit: "g", Purpose: "color, sweetness"}, // Carrot (2 tbsp shredded)
			{FoodItemID: 29, Quantity: 15, Unit: "g", Purpose: "flavor"},           // Onion (1 tbsp minced)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Core the cabbage and chop it very finely (aim for pieces roughly the size of rice grains). A food processor with a shredding/chopping blade can help."},
			{StepNumber: 2, Instruction: "Peel and finely shred or grate the carrot."},
			{StepNumber: 3, Instruction: "Mince the onion very finely."},
			{StepNumber: 4, Instruction: "Combine the prepared cabbage, carrot, and onion in a large bowl."},
		},
		EquipmentIDs: []int64{55, 56, 51, 68, 54}, // Knife, Chopping Board, Grater/Food Processor, Peeler, Large Bowl
	},
	{
		ID:          1102,
		RecipeID:    11,
		StepOrder:   2,
		Title:       "Make Dressing and Combine",
		Description: "Whisk dressing ingredients, pour over vegetables, mix, and chill.",
		Notes:       "Chill for at least 1 hour for flavors to meld.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 54, Quantity: 60, Unit: "g", Purpose: "dressing base"},         // Mayonnaise (1/4 cup)
			{FoodItemID: 95, Quantity: 30, Unit: "ml", Purpose: "dressing tang"},        // Buttermilk (2 tbsp)
			{FoodItemID: 15, Quantity: 30, Unit: "ml", Purpose: "dressing consistency"}, // Milk (Full Fat or Trim) (2 tbsp)
			{FoodItemID: 96, Quantity: 25, Unit: "g", Purpose: "dressing sweetness"},    // Granulated Sugar (2 tbsp)
			{FoodItemID: 85, Quantity: 15, Unit: "ml", Purpose: "dressing brightness"},  // Lemon Juice (1 tbsp)
			{FoodItemID: 58, Quantity: 15, Unit: "ml", Purpose: "dressing acidity"},     // White Vinegar (1 tbsp)
			{FoodItemID: 16, Quantity: 1.25, Unit: "g", Purpose: "dressing seasoning"},  // Salt (1/4 tsp)
			{FoodItemID: 62, Quantity: 0.5, Unit: "g", Purpose: "dressing seasoning"},   // Pepper (Pinch)
			// Prepared vegetables from Step 1
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In a medium bowl, whisk together mayonnaise, buttermilk, milk, granulated sugar, lemon juice, white vinegar, salt, and pepper until smooth and the sugar is dissolved."},
			{StepNumber: 2, Instruction: "Pour the dressing over the prepared vegetables in the large bowl."},
			{StepNumber: 3, Instruction: "Mix thoroughly until all vegetables are evenly coated."},
			{StepNumber: 4, Instruction: "Cover the bowl tightly and refrigerate for at least 1 hour, preferably longer (2-4 hours or overnight)."},
			{StepNumber: 5, Instruction: "Stir the coleslaw again just before serving."},
		},
		EquipmentIDs: []int64{54, 53, 59, 78}, // Medium Bowl, Whisk, Large Bowl (from Step 1), Spoon, Plastic Wrap
	},

	// --- Recipe 12: Mi Goreng-Inspired Satay Noodles ---
	{
		ID:          1201, // Start ID for Recipe 12
		RecipeID:    12,
		StepOrder:   1,
		Title:       "Prep Sauce and Aromatics",
		Description: "Dissolve stock cube, whisk in peanut butter and liquids. Prepare aromatics.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 87, Quantity: 0.25, Unit: "unit", Purpose: "sauce base"},   // Chicken Stock Cube (1/4)
			{FoodItemID: 25, Quantity: 30, Unit: "ml", Purpose: "sauce liquid"},     // Very Hot Water (2 tbsp)
			{FoodItemID: 97, Quantity: 15, Unit: "g", Purpose: "main flavor"},       // Smooth Peanut Butter (1 tbsp)
			{FoodItemID: 55, Quantity: 30, Unit: "ml", Purpose: "sauce base"},       // Soy Sauce (2 tbsp)
			{FoodItemID: 86, Quantity: 7.5, Unit: "g", Purpose: "sauce balance"},    // Brown Sugar (1.5 tsp)
			{FoodItemID: 58, Quantity: 15, Unit: "ml", Purpose: "sauce acid"},       // White Vinegar (1 tbsp)
			{FoodItemID: 31, Quantity: 15, Unit: "g", Purpose: "aromatic"},          // Garlic (3 cloves), minced OR 2-3 tsp paste
			{FoodItemID: 30, Quantity: 3, Unit: "g", Purpose: "aromatic"},           // Ginger (~1cm), finely grated or paste
			{FoodItemID: 52, Quantity: 3, Unit: "unit", Purpose: "heat"},            // Bird's Eye Chilies (2-3), finely chopped OR 1.5 tsp Chilli Powder (ID 98)
			{FoodItemID: 29, Quantity: 75, Unit: "g", Purpose: "aromatic"},          // Onion (1/2), finely diced
			{FoodItemID: 53, Quantity: 30, Unit: "g", Purpose: "aromatic, garnish"}, // Spring Onions (2), whites/greens separated
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Sauce Prep: Dissolve stock cube in very hot water. Immediately whisk in peanut butter until completely smooth. Gradually mix in soy sauce, brown sugar, and vinegar. Taste for balance. Keep warm if possible."},
			{StepNumber: 2, Instruction: "Aromatics Prep: Mince garlic (if using cloves). Finely grate ginger (Microplane recommended). Finely chop chilies (if using fresh). Finely dice onion. Slice spring onions, separating whites and greens."},
		},
		EquipmentIDs: []int64{54, 53, 55, 56, 51}, // Small Bowl, Whisk/Fork, Knife, Chopping Board, Grater
	},
	{
		ID:          1202,
		RecipeID:    12,
		StepOrder:   2,
		Title:       "Cook Noodles, Aromatics, Egg",
		Description: "Cook noodles, reserve water. Stir-fry aromatics and egg.",
		Notes:       "Watch aromatics carefully.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 90, Quantity: 150, Unit: "g", Purpose: "base"},             // Egg Noodles (1-2 servings)
			{FoodItemID: 74, Quantity: 1, Unit: "unit", Purpose: "protein, binder"}, // Egg
			{FoodItemID: 60, Quantity: 30, Unit: "ml", Purpose: "cooking fat"},      // Oil (2 tbsp)
			// Prepared aromatics (garlic, ginger, chili, onion, spring onion whites) from Step 1
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Cook egg noodles until just underdone. Reserve 1/4 cup (60ml) cooking water. Drain noodles, toss with a little oil."},
			{StepNumber: 2, Instruction: "Heat wok or cast iron pan over medium-high heat. Add oil."},
			{StepNumber: 3, Instruction: "Add garlic, ginger, chilies (or chili powder), onion, and spring onion whites. Stir-fry 3-4 minutes until onions are translucent and lightly browned."},
			{StepNumber: 4, Instruction: "Push aromatics aside. Crack egg into empty space, break up immediately, cook until just set but still wet, mixing into aromatics."},
		},
		EquipmentIDs: []int64{50, 81, 58, 60, 62, 63}, // Saucepan, Strainer, Measuring Cup, Wok/Pan, Stovetop, Spatula
	},
	{
		ID:          1203,
		RecipeID:    12,
		StepOrder:   3,
		Title:       "Combine and Finish",
		Description: "Add noodles and sauce, toss, adjust seasoning, garnish.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Cooked noodles, aromatics/egg from Step 2
			// Prepared sauce, reserved noodle water from Step 1/2
			// Spring onion greens from Step 1
			// Optional: Worcestershire sauce (ID 89), extra chili, noodle water+soy+sugar, vinegar
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Add cooked noodles and prepared sauce to the pan."},
			{StepNumber: 2, Instruction: "Toss continuously over medium-high heat until noodles are coated and heated through."},
			{StepNumber: 3, Instruction: "Taste and adjust: add Worcestershire for depth, more chili for heat, noodle water (or water+soy+sugar mix) if too thick/dry, vinegar if flavors muted."},
			{StepNumber: 4, Instruction: "If sauce separates, remove from heat and stir vigorously."},
			{StepNumber: 5, Instruction: "Turn off heat. Garnish with spring onion greens."},
			{StepNumber: 6, Instruction: "Serve immediately."},
		},
		EquipmentIDs: []int64{60, 67, 63}, // Wok/Pan, Tongs, Spatula
	},

	// --- Recipe 13: Pambazo ---
	{
		ID:          1301, // Start ID for Recipe 13
		RecipeID:    13,
		StepOrder:   1,
		Title:       "Make Guajillo (Kashmiri) Chile Sauce",
		Description: "Toast, soak, and blend chilies with garlic and broth.",
		Notes:       "Strain for smoother sauce.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 99, Quantity: 6, Unit: "unit", Purpose: "sauce base, color"},   // Kashmiri Chilies (4-6 dried)
			{FoodItemID: 100, Quantity: 1, Unit: "unit", Purpose: "sauce flavor, heat"}, // Chipotle Pepper (1 dried)
			{FoodItemID: 25, Quantity: 250, Unit: "ml", Purpose: "soaking"},             // Hot Water
			{FoodItemID: 31, Quantity: 10, Unit: "g", Purpose: "sauce aromatic"},        // Garlic (2 cloves)
			{FoodItemID: 80, Quantity: 250, Unit: "ml", Purpose: "sauce liquid"},        // Chicken Broth
			{FoodItemID: 16, Quantity: 3, Unit: "g", Purpose: "seasoning"},              // Salt (to taste)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Remove stems and seeds (optional, for less heat) from Kashmiri and chipotle chilies."},
			{StepNumber: 2, Instruction: "Briefly toast the chilies in a dry cast iron pan over medium heat until fragrant (10-15 seconds per side). Do not burn."},
			{StepNumber: 3, Instruction: "Place toasted chilies in a bowl and cover with hot water. Let soak for 10 minutes until softened."},
			{StepNumber: 4, Instruction: "Transfer soaked chilies (discard soaking water unless needed for blending consistency), garlic cloves, and chicken broth to a blender."},
			{StepNumber: 5, Instruction: "Blend until completely smooth."},
			{StepNumber: 6, Instruction: "Strain the sauce through a fine mesh strainer into a bowl, pressing on solids."},
			{StepNumber: 7, Instruction: "Season the strained sauce with salt to taste."},
		},
		EquipmentIDs: []int64{55, 60, 62, 54, 64, 76}, // Knife, Cast Iron Pan, Stovetop, Bowl, Blender, Strainer
	},
	{
		ID:          1302,
		RecipeID:    13,
		StepOrder:   2,
		Title:       "Make Chorizo-Potato Filling",
		Description: "Cook chorizo and diced potatoes until potatoes are tender.",
		Notes:       "Par-boil potatoes for faster cooking if desired.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 101, Quantity: 230, Unit: "g", Purpose: "filling main"}, // Chorizo Sausage
			{FoodItemID: 59, Quantity: 150, Unit: "g", Purpose: "filling bulk"},  // Agria Potato (1), diced small
			{FoodItemID: 16, Quantity: 2, Unit: "g", Purpose: "seasoning"},       // Salt (to taste)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat a cast iron pan over medium heat."},
			{StepNumber: 2, Instruction: "Add chorizo to the pan, breaking it up with a spoon as it cooks (2-3 minutes). Drain excess fat if desired."},
			{StepNumber: 3, Instruction: "Add the diced Agria potato to the pan with the chorizo."},
			{StepNumber: 4, Instruction: "Cook, stirring occasionally, for about 8-10 minutes, or until the potatoes are tender. (Optional: Cover pan briefly to help potatoes steam)."},
			{StepNumber: 5, Instruction: "Season the filling with salt to taste."},
		},
		EquipmentIDs: []int64{60, 62, 59, 55, 56}, // Cast Iron Pan, Stovetop, Spoon, Knife, Chopping Board
	},
	{
		ID:          1303,
		RecipeID:    13,
		StepOrder:   3,
		Title:       "Prepare Bread",
		Description: "Dip buns in chili sauce and sear in oil.",
		Notes:       "User note: consider pre-frying buns slightly *before* dipping.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 102, Quantity: 4, Unit: "unit", Purpose: "sandwich base"}, // Brioche Buns
			// Prepared chili sauce from Step 1
			{FoodItemID: 34, Quantity: 15, Unit: "ml", Purpose: "searing"}, // Canola Oil (~1 tbsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Pour some of the prepared chile sauce into a shallow dish."},
			{StepNumber: 2, Instruction: "Dip each brioche bun completely into the chile sauce, ensuring it's coated but not overly saturated."},
			{StepNumber: 3, Instruction: "Heat canola oil in the cast iron pan (or another skillet) over medium heat."},
			{StepNumber: 4, Instruction: "Carefully place the sauce-dipped buns, cut-side down first if split, into the hot oil."},
			{StepNumber: 5, Instruction: "Sear for 1-2 minutes per side, pressing gently with a spatula, until toasted and slightly crisped."},
			{StepNumber: 6, Instruction: "Remove buns to a plate."},
		},
		EquipmentIDs: []int64{54, 60, 62, 63, 67}, // Shallow Dish, Cast Iron Pan, Stovetop, Spatula, Tongs
	},
	{
		ID:          1304,
		RecipeID:    13,
		StepOrder:   4,
		Title:       "Assemble Pambazo",
		Description: "Layer filling and toppings onto seared buns.",
		Notes:       "Minimize wet ingredients; crumble feta finely.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Seared buns from Step 3
			// Chorizo-potato filling from Step 2
			{FoodItemID: 84, Quantity: 40, Unit: "g", Purpose: "topping"},   // Feta Cheese, crumbled
			{FoodItemID: 65, Quantity: 20, Unit: "g", Purpose: "topping"},   // Parmesan, grated
			{FoodItemID: 103, Quantity: 30, Unit: "ml", Purpose: "topping"}, // Smoky Salsa Verde
			{FoodItemID: 67, Quantity: 40, Unit: "g", Purpose: "topping"},   // Lettuce, shredded
			{FoodItemID: 68, Quantity: 30, Unit: "g", Purpose: "topping"},   // Pickled Onions (Recipe ID 14)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Place the bottom halves of the seared buns on serving plates."},
			{StepNumber: 2, Instruction: "Spoon a generous amount of the chorizo-potato mixture onto each bottom bun."},
			{StepNumber: 3, Instruction: "Top with crumbled feta cheese (crumble finely) and grated parmesan."},
			{StepNumber: 4, Instruction: "Drizzle with smoky salsa verde."},
			{StepNumber: 5, Instruction: "Add shredded lettuce."},
			{StepNumber: 6, Instruction: "Finish with pickled onions."},
			{StepNumber: 7, Instruction: "Place the top halves of the buns on top. Slice in half if desired and serve immediately."},
		},
		EquipmentIDs: []int64{59}, // Spoon
	},

	// --- Recipe 15: Quick Pizza Sauce (Component) ---
	{
		ID:          1501, // Start ID for Recipe 15
		RecipeID:    15,
		StepOrder:   1,
		Title:       "Blend Sauce",
		Description: "Drain tomatoes, combine ingredients, blend until smooth.",
		Notes:       "Reserve tomato liquid to adjust consistency if needed.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 104, Quantity: 400, Unit: "g", Purpose: "base"},                   // Canned Diced Tomatoes, drained (reserve liquid)
			{FoodItemID: 105, Quantity: 30, Unit: "g", Purpose: "thickener, flavor"},       // Tomato Paste (2 tbsp)
			{FoodItemID: 31, Quantity: 5, Unit: "g", Purpose: "aromatic"},                  // Garlic (1 clove), minced
			{FoodItemID: 106, Quantity: 1, Unit: "g", Purpose: "flavor"},                   // Dried Oregano (1 tsp)
			{FoodItemID: 16, Quantity: 1.25, Unit: "g", Purpose: "seasoning"},              // Salt (1/4 tsp)
			{FoodItemID: 96, Quantity: 1, Unit: "g", Purpose: "balance acidity"},           // Sugar (Pinch)
			{FoodItemID: 107, Quantity: 0.5, Unit: "g", IsOptional: true, Purpose: "heat"}, // Red Pepper Flakes (Optional pinch)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Drain the canned diced tomatoes, reserving the liquid in a separate bowl."},
			{StepNumber: 2, Instruction: "Combine the drained tomatoes, tomato paste, minced garlic, dried oregano, salt, pinch of sugar, and optional red pepper flakes in a blending cup or bowl suitable for an immersion blender (or in a standard blender)."},
			{StepNumber: 3, Instruction: "Blend using an immersion blender (or standard blender) until the sauce reaches desired smoothness."},
			{StepNumber: 4, Instruction: "Check the consistency. If the sauce is too thick, blend in a small amount of the reserved tomato liquid until it reaches the desired thickness (it should be thick enough to spread but not watery)."},
			{StepNumber: 5, Instruction: "Taste and adjust salt if needed."},
		},
		EquipmentIDs: []int64{81, 54, 64, 59}, // Strainer/Sieve, Bowl, Immersion Blender/Blender, Spoon
	},

	// --- Recipe 16: Red Taco Sauce (Component) ---
	{
		ID:          1601, // Start ID for Recipe 16
		RecipeID:    16,
		StepOrder:   1,
		Title:       "Prepare Chilies",
		Description: "Soak dried chilies until soft.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 99, Quantity: 3, Unit: "unit", Purpose: "chili base"},        // Kashmiri Chilies (2-3), stems removed
			{FoodItemID: 108, Quantity: 25, Unit: "g", Purpose: "chili base, flavor"}, // Chipotles in Adobo (1 pepper + 1 tbsp sauce)
			{FoodItemID: 25, Quantity: 240, Unit: "ml", Purpose: "soaking"},           // Hot Water (1 cup)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Place the stemmed Kashmiri chilies in a heatproof bowl."},
			{StepNumber: 2, Instruction: "Pour the hot water over the chilies, ensuring they are submerged (use a small plate to weigh them down if needed)."},
			{StepNumber: 3, Instruction: "Let soak for 30 minutes until the chilies are soft and pliable."},
			{StepNumber: 4, Instruction: "Reserve the soaking liquid."},
		},
		EquipmentIDs: []int64{54, 58}, // Bowl, Measuring Cup
	},
	{
		ID:          1602,
		RecipeID:    16,
		StepOrder:   2,
		Title:       "Roast/Toast Aromatics",
		Description: "Grill garlic and tomatoes, toast cumin.",
		Notes:       "Watch carefully to prevent burning.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 31, Quantity: 20, Unit: "g", Purpose: "aromatic"}, // Garlic (4 cloves), unpeeled
			{FoodItemID: 104, Quantity: 400, Unit: "g", Purpose: "base"},   // Canned Diced Tomatoes (1 can)
			{FoodItemID: 109, Quantity: 5, Unit: "g", Purpose: "spice"},    // Cumin (1 tsp, assumed ground)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Turn on the grill/broiler setting in your toaster oven or main oven."},
			{StepNumber: 2, Instruction: "Place the unpeeled garlic cloves on a small baking tray or piece of foil. Grill until the skin is charred and the garlic feels soft (about 7-10 minutes). Let cool slightly."},
			{StepNumber: 3, Instruction: "Spread the diced tomatoes (drain slightly if very watery) on a baking tray. Grill until slightly charred and some moisture has evaporated (15-20 minutes). Let cool slightly."},
			{StepNumber: 4, Instruction: "Optional (if using whole cumin seeds): Toast cumin seeds in a dry pan until fragrant, then grind."},
		},
		EquipmentIDs: []int64{69, 61, 66}, // Toaster Oven / Oven Grill, Baking Tray
	},
	{
		ID:          1603,
		RecipeID:    16,
		StepOrder:   3,
		Title:       "Blend Sauce",
		Description: "Combine all prepared ingredients and blend until smooth.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Soaked Kashmiri chilies, reserved soaking liquid from Step 1
			// Chipotles in adobo from Step 1
			// Roasted garlic (peeled), charred tomatoes, toasted cumin from Step 2
			{FoodItemID: 106, Quantity: 5, Unit: "g", Purpose: "flavor"},       // Oregano (1 tsp, assumed dried)
			{FoodItemID: 110, Quantity: 30, Unit: "ml", Purpose: "acid"},       // Apple Cider Vinegar (2 tbsp)
			{FoodItemID: 18, Quantity: 5, Unit: "g", Purpose: "flavor, color"}, // Smoked Paprika (1 tsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Peel the roasted garlic cloves."},
			{StepNumber: 2, Instruction: "In a blender, combine the soaked Kashmiri chilies, chipotles and adobo sauce, peeled roasted garlic, charred tomatoes, cumin, oregano, apple cider vinegar, and smoked paprika."},
			{StepNumber: 3, Instruction: "Add about 1/4 cup (60ml) of the reserved chili soaking liquid."},
			{StepNumber: 4, Instruction: "Blend on high speed until the sauce is completely smooth."},
		},
		EquipmentIDs: []int64{64}, // Blender
	},
	{
		ID:          1604,
		RecipeID:    16,
		StepOrder:   4,
		Title:       "Cook and Finish Sauce",
		Description: "Simmer sauce to thicken, season, optionally strain.",
		Notes:       "User notes indicate skipping this simmer step can work.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Blended sauce from Step 3
			// Reserved soaking liquid from Step 1
			{FoodItemID: 57, Quantity: 5, Unit: "g", IsOptional: true, Purpose: "balance"}, // Honey (1 tsp, optional)
			{FoodItemID: 16, Quantity: 3, Unit: "g", Purpose: "seasoning"},                 // Salt (to taste, user note: ~1/2 tsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Pour the blended sauce into a saucepan."},
			{StepNumber: 2, Instruction: "(Optional - based on user notes) Bring the sauce to a simmer over medium heat. Reduce heat to low and simmer gently for 10-15 minutes, stirring occasionally, until slightly thickened."},
			{StepNumber: 3, Instruction: "Stir in honey, if using."},
			{StepNumber: 4, Instruction: "Season with salt to taste (start with 1/2 tsp)."},
			{StepNumber: 5, Instruction: "If the sauce is too thick, add a little more of the reserved chili soaking liquid (or water) until it reaches desired consistency (thin enough to drizzle but thick enough to cling)."},
			{StepNumber: 6, Instruction: "(Optional) For an extra smooth sauce, strain through a fine mesh strainer."},
			{StepNumber: 7, Instruction: "Let cool slightly before using. Sauce will thicken more as it cools."},
		},
		EquipmentIDs: []int64{65, 62, 59, 76}, // Small/Medium Saucepan, Stovetop, Spoon, Strainer (optional)
	},

	// --- Recipe 17: Subway-Style Baguette Adaptation ---
	// Note: This represents the 'base' method, with variations described in the recipe notes.
	{
		ID:          1701, // Start ID for Recipe 17
		RecipeID:    17,
		StepOrder:   1,
		Title:       "Activate Yeast & Make Dough (Subway Style)",
		Description: "Modified pizza dough: less salt, more oil/sugar.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 25, Quantity: 210, Unit: "ml", Purpose: "hydration"},                   // Warm Water (40-46°C)
			{FoodItemID: 22, Quantity: 5, Unit: "g", Purpose: "leavening"},                      // Active Dry Yeast (3/4 tsp)
			{FoodItemID: 23, Quantity: 6, Unit: "g", Purpose: "softness, browning"},             // Raw Sugar (increased)
			{FoodItemID: 21, Quantity: 280, Unit: "g", Purpose: "structure"},                    // Bread Flour
			{FoodItemID: 16, Quantity: 4, Unit: "g", Purpose: "flavor"},                         // Salt (reduced)
			{FoodItemID: 24, Quantity: 11, Unit: "g", Purpose: "tenderness"},                    // Olive Oil (~12ml, increased)
			{FoodItemID: 111, Quantity: 10, Unit: "g", IsOptional: true, Purpose: "tenderness"}, // Milk Powder (Optional Enhancement)
			{FoodItemID: 112, Quantity: 1, Unit: "g", IsOptional: true, Purpose: "aroma"},       // Diastatic Malt Powder (Optional Enhancement, 1/4 tsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Activate yeast in warm water with sugar (10-15 mins until foamy)."},
			{StepNumber: 2, Instruction: "In large bowl, mix flour, salt, optional milk powder/malt powder."},
			{StepNumber: 3, Instruction: "Add yeast mixture and olive oil. Mix until shaggy dough forms."},
			{StepNumber: 4, Instruction: "Rest 10 minutes (autolyse)."},
		},
		EquipmentIDs: []int64{58, 54, 59}, // Measuring tools, Bowl, Fork/Spoon
	},
	{
		ID:          1702,
		RecipeID:    17,
		StepOrder:   2,
		Title:       "Develop Dough (Stretch & Fold)",
		Description: "Perform 3-4 sets of stretch and folds.",
		Notes:       "User timing suggests 20 min rests can work.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Perform first stretch and fold. Cover and rest 30 minutes (or 20 mins)."},
			{StepNumber: 2, Instruction: "Perform second stretch and fold. Cover and rest 30 minutes (or 20 mins)."},
			{StepNumber: 3, Instruction: "Perform third stretch and fold. Cover and rest 30 minutes (or 20 mins)."},
			{StepNumber: 4, Instruction: "Perform final stretch and fold if needed. Cover and rest 30 minutes (or 20 mins)."},
		},
		EquipmentIDs: []int64{54}, // Bowl
	},
	{
		ID:          1703,
		RecipeID:    17,
		StepOrder:   3,
		Title:       "Shape Baguettes",
		Description: "Divide dough, shape into logs.",
		Notes:       "Use minimal flour; consider oiled hands.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 14, Quantity: 10, Unit: "g", Purpose: "shaping"}, // AP Flour (minimal) or Oil for hands
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Divide dough into two equal portions (~247g each)."},
			{StepNumber: 2, Instruction: "On a very lightly floured surface (or using oiled hands), gently flatten one portion into a rectangle (approx 20cm x 15cm)."},
			{StepNumber: 3, Instruction: "Fold one long edge one-third towards the center. Fold the opposite edge over to meet the first fold (letter fold). Seal seam."},
			{StepNumber: 4, Instruction: "Gently roll the dough into a 25-30cm log with slightly tapered ends."},
			{StepNumber: 5, Instruction: "Place seam-side down on a baking tray lined with parchment paper."},
			{StepNumber: 6, Instruction: "Repeat for the second portion."},
		},
		EquipmentIDs: []int64{56, 66, 77}, // Work Surface, Baking Tray, Parchment Paper
	},
	{
		ID:          1704,
		RecipeID:    17,
		StepOrder:   4,
		Title:       "Final Proof",
		Description: "Let shaped loaves rise until nearly doubled.",
		Notes:       "User note: Extended proof (95 mins total from shaping) worked well.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Cover the shaped loaves loosely with lightly oiled plastic wrap or a clean kitchen towel."},
			{StepNumber: 2, Instruction: "Let rise in a warm place until nearly doubled in size (approx 45-60 minutes, or longer as per user notes - up to 95 mins)."},
		},
		EquipmentIDs: []int64{66, 78, 75}, // Baking Tray, Plastic Wrap/Towel
	},
	{
		ID:          1705,
		RecipeID:    17,
		StepOrder:   5,
		Title:       "Bake Baguettes (Attempt 3 / Open Spray Method)",
		Description: "Score, spray aggressively, bake with temp reduction, spray during bake.",
		Notes:       "This reflects Attempt 3/Spray method. Other attempts involve covering with glass dish or different temp profiles. Target internal temp 90-93C.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 25, Quantity: 50, Unit: "ml", Purpose: "steam/crust"}, // Water (for spraying)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Preheat oven to 200°C (fan bake). Place an empty oven-safe dish on the bottom rack if desired for extra steam (add boiling water later)."},
			{StepNumber: 2, Instruction: "Once proofed, make 3-4 shallow diagonal scores on each loaf using a sharp knife or lame."},
			{StepNumber: 3, Instruction: "Spray the loaves generously with water using a spray bottle."},
			{StepNumber: 4, Instruction: "(Optional) Carefully add boiling water to the preheated dish on the bottom rack."},
			{StepNumber: 5, Instruction: "Place the baking tray with loaves into the oven. Spray oven walls quickly with water."},
			{StepNumber: 6, Instruction: "Bake for 6-7 minutes at 200°C."},
			{StepNumber: 7, Instruction: "Quickly open oven, spray loaves and walls again, and reduce temperature to 170°C (fan bake)."},
			{StepNumber: 8, Instruction: "Continue baking, spraying again every 5 minutes (e.g., at ~12 mins, ~17 mins total time)."},
			{StepNumber: 9, Instruction: "Bake for a total of approximately 20-22 minutes. For the last 3-5 minutes, do not spray, allowing the surface to set."},
			{StepNumber: 10, Instruction: "Check for doneness: light golden color, internal temperature of 90-93°C."},
			{StepNumber: 11, Instruction: "Immediately remove loaves from oven and wrap them snugly in a clean kitchen towel."},
			{StepNumber: 12, Instruction: "Let cool completely on a wire rack while still wrapped (at least 60 minutes) before unwrapping and slicing."},
		},
		EquipmentIDs: []int64{55, 79, 61, 66, 52, 80, 75, 74}, // Knife/Lame, Spray Bottle, Oven, Baking Tray, Oven-safe Dish (optional), Thermometer, Kitchen Towel, Wire Rack
	},
	{
		ID:          1801, // Assign next available ID block start
		RecipeID:    18,   // Link to the new Bao Bun recipe ID
		StepOrder:   1,
		Title:       "Prep Chicken & Marinade",
		Description: "Slice chicken, crush chili, mix marinade ingredients and coat chicken.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 26, Quantity: 300, Unit: "g", Purpose: "protein"},        // Chicken Thigh
			{FoodItemID: 52, Quantity: 1, Unit: "unit", Purpose: "heat"},          // Dried Bird's Eye Chili
			{FoodItemID: 55, Quantity: 15, Unit: "ml", Purpose: "marinade base"},  // Soy Sauce (1 Tbsp)
			{FoodItemID: 116, Quantity: 5, Unit: "ml", Purpose: "marinade aroma"}, // Sesame Oil (1 tsp)
			{FoodItemID: 31, Quantity: 5, Unit: "g", Purpose: "marinade flavor"},  // Garlic Powder/Minced (1 tsp)
			{FoodItemID: 30, Quantity: 5, Unit: "g", Purpose: "marinade flavor"},  // Ground Ginger/Paste (1 tsp)
			{FoodItemID: 117, Quantity: 2, Unit: "g", Purpose: "marinade spice"},  // Sichuan Pepper, ground (1/2 tsp)
			{FoodItemID: 118, Quantity: 1, Unit: "g", Purpose: "marinade spice"},  // Allspice, ground (1/4 tsp)
			{FoodItemID: 18, Quantity: 2.5, Unit: "g", Purpose: "marinade spice"}, // Smoked Paprika (1/2 tsp)
			{FoodItemID: 74, Quantity: 1, Unit: "unit", Purpose: "binding"},       // Egg
			{FoodItemID: 119, Quantity: 45, Unit: "ml", Purpose: "coating"},       // Cornflour (3 Tbsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Slice chicken thighs into bite-sized pieces (approx 1cm thick strips or cubes). Place in a medium bowl."},
			{StepNumber: 2, Instruction: "Crush the dried bird's eye chili using a mortar and pestle or the flat side of a knife."},
			{StepNumber: 3, Instruction: "In a small ramekin or bowl, mix soy sauce, sesame oil, garlic, ginger, Sichuan pepper, allspice, smoked paprika, and the crushed chili."},
			{StepNumber: 4, Instruction: "Pour the spice mixture over the chicken pieces and mix thoroughly to coat."},
			{StepNumber: 5, Instruction: "Add the egg and cornflour to the chicken bowl. Mix well until chicken is fully coated."},
			{StepNumber: 6, Instruction: "Cover and marinate in the fridge for 15-20 minutes."},
		},
		EquipmentIDs: []int64{55, 56, 57, 58, 91, 74, 59}, // Medium Bowl, Small Bowl/Ramekin, Knife, Chopping Board, Mortar & Pestle (or Knife), Spoon
	},
	{
		ID:          1802,
		RecipeID:    18,
		StepOrder:   2,
		Title:       "Prep Slaw, Mayo & Garnishes",
		Description: "Prepare coleslaw mix, cucumber, peanuts, sriracha mayo, and spring onions.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 94, Quantity: 100, Unit: "g", Purpose: "slaw base"},       // Cabbage, shredded
			{FoodItemID: 50, Quantity: 50, Unit: "g", Purpose: "slaw color"},       // Carrot, grated
			{FoodItemID: 51, Quantity: 100, Unit: "g", Purpose: "freshness"},       // Cucumber (1/2), thinly sliced
			{FoodItemID: 123, Quantity: 25, Unit: "g", Purpose: "garnish texture"}, // Roasted Peanuts (2 Tbsp), crushed
			{FoodItemID: 120, Quantity: 45, Unit: "g", Purpose: "mayo base"},       // Kewpie Mayo (3 Tbsp)
			{FoodItemID: 121, Quantity: 10, Unit: "ml", Purpose: "mayo heat"},      // Sriracha Sauce (1-2 tsp)
			{FoodItemID: 57, Quantity: 2.5, Unit: "g", Purpose: "mayo balance"},    // Honey (1/2 tsp)
			{FoodItemID: 85, Quantity: 1, Unit: "ml", Purpose: "mayo brightness"},  // Lemon Juice (Few drops)
			{FoodItemID: 116, Quantity: 1, Unit: "ml", Purpose: "mayo aroma"},      // Sesame Oil (1/4 tsp / drop)
			{FoodItemID: 53, Quantity: 20, Unit: "g", Purpose: "garnish"},          // Spring Onion, thinly sliced
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Combine shredded cabbage and grated carrot in a bowl for the coleslaw base."},
			{StepNumber: 2, Instruction: "Thinly slice the cucumber into half-moons or matchsticks. (Optional: salt briefly and pat dry)."},
			{StepNumber: 3, Instruction: "Crush the roasted peanuts using a mortar and pestle or place in a bag and crush with a rolling pin or heavy object."},
			{StepNumber: 4, Instruction: "Make Sriracha Mayo: In a small bowl, mix Kewpie mayo, sriracha sauce (start with less, add more to taste), honey, lemon juice, and sesame oil until combined."},
			{StepNumber: 5, Instruction: "Thinly slice the spring onions."},
		},
		EquipmentIDs: []int64{54, 55, 57, 58, 51, 91, 92, 59}, // Bowl (for slaw), Small Bowl (for mayo), Knife, Chopping Board, Grater, Mortar & Pestle (or Bag/Rolling Pin), Spoon
	},
	{
		ID:          1803,
		RecipeID:    18,
		StepOrder:   3,
		Title:       "Fry Chicken & Steam Buns",
		Description: "Fry marinated chicken until crispy. Steam bao buns.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Marinated Chicken from Step 1
			{FoodItemID: 34, Quantity: 45, Unit: "ml", Purpose: "frying"},   // Canola Oil (3 Tbsp)
			{FoodItemID: 122, Quantity: 8, Unit: "unit", Purpose: "vessel"}, // Bao Buns (adjust number as needed)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat canola oil in a wok or large, deep skillet over medium-high heat."},
			{StepNumber: 2, Instruction: "Once oil is hot (a piece of batter sizzles immediately), carefully add chicken pieces in a single layer, working in batches to avoid overcrowding."},
			{StepNumber: 3, Instruction: "Fry chicken for 2-3 minutes per side, or until golden brown, crispy, and cooked through."},
			{StepNumber: 4, Instruction: "Use tongs or a slotted spoon to remove cooked chicken and place on a plate lined with paper towels to drain."},
			{StepNumber: 5, Instruction: "Repeat with remaining chicken."},
			{StepNumber: 6, Instruction: "While the last batch of chicken is frying, steam the bao buns according to package instructions (usually 3-4 minutes over simmering water)."},
		},
		EquipmentIDs: []int64{84, 62, 67, 80, 93, 94}, // Wok or Skillet, Stovetop, Tongs, Slotted Spoon, Plate, Steamer
	},
	{
		ID:          1804,
		RecipeID:    18,
		StepOrder:   4,
		Title:       "Assemble Bao Buns",
		Description: "Layer mayo, slaw, cucumber, chicken, and garnishes in steamed buns.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Steamed Bao Buns from Step 3
			// Sriracha Mayo from Step 2
			// Coleslaw mix from Step 2
			// Sliced Cucumber from Step 2
			// Crispy Chicken from Step 3
			// Crushed Peanuts from Step 2
			// Sliced Spring Onions from Step 2
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Carefully open a warm steamed bao bun."},
			{StepNumber: 2, Instruction: "Spread a small amount of sriracha mayo on the inside surface(s)."},
			{StepNumber: 3, Instruction: "Add a layer of the cabbage and carrot slaw."},
			{StepNumber: 4, Instruction: "Place a few slices of cucumber on top."},
			{StepNumber: 5, Instruction: "Add 3-4 pieces of the crispy fried chicken."},
			{StepNumber: 6, Instruction: "Drizzle with a bit more sriracha mayo."},
			{StepNumber: 7, Instruction: "Sprinkle generously with crushed peanuts and sliced spring onions."},
			{StepNumber: 8, Instruction: "Repeat for remaining buns and serve immediately."},
		},
		EquipmentIDs: []int64{59}, // Spoon/Spreader
	},
	// --- ADD THESE TO THE DummyRecipeSteps slice ---
	// === Recipe 3: Chicken Stir Fry ===
	// --- Step 1: Preparation Stage ---
	{
		ID:          1901,
		RecipeID:    19,
		StepOrder:   1,
		Title:       "Marinade Chicken",
		Description: "Prepare and marinade chicken.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string // Unit as specified IN THE RECIPE TEXT
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 26, Quantity: 200, Unit: "g", Purpose: "protein"},                         // Chicken Thighs (using midpoint of 150-250g)
			{FoodItemID: 55, Quantity: 15, Unit: "ml", Purpose: "marinade umami"},                  // Soy Sauce
			{FoodItemID: 31, Quantity: 2, Unit: "cloves", Purpose: "marinade aromatic"},            // Garlic (Specified Minced/Powder - store handles default/form later)
			{FoodItemID: 30, Quantity: 5, Unit: "ml", Purpose: "marinade aromatic"},                // Ginger (Specified Paste/Ground)
			{FoodItemID: 117, Quantity: 2, Unit: "g", IsOptional: true, Purpose: "marinade spice"}, // Sichuan Pepper, ground
			{FoodItemID: 18, Quantity: 2.5, Unit: "ml", Purpose: "marinade flavour"},               // Smoked Paprika
			{FoodItemID: 64, Quantity: 1.25, Unit: "ml", Purpose: "marinade heat"},                 // Cayenne Pepper (used pinch estimate)
			{FoodItemID: 109, Quantity: 1.25, Unit: "ml", Purpose: "marinade spice"},               // Cumin, Ground
			{FoodItemID: 124, Quantity: 1.25, Unit: "ml", Purpose: "marinade spice"},               // Coriander, Ground
			{FoodItemID: 16, Quantity: 2, Unit: "g", Purpose: "marinade seasoning"},                // Salt
			{FoodItemID: 62, Quantity: 1.25, Unit: "ml", Purpose: "marinade seasoning"},            // Black Pepper
			{FoodItemID: 74, Quantity: 1, Unit: "unit", Purpose: "binding"},                        // Egg
			{FoodItemID: 119, Quantity: 45, Unit: "ml", Purpose: "coating"},                        // Cornflour
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Chicken Prep: Cut chicken thighs into bite-sized pieces. Place in a medium bowl."},
			{StepNumber: 2, Instruction: "Marinade: Whisk together soy sauce (1 Tbsp), garlic (1 tsp minced/0.5 tsp powder), ginger (1 tsp paste/0.5 tsp ground), Sichuan pepper (if using), smoked paprika, cayenne/chilli, cumin, coriander, salt (0.5 tsp), and black pepper (0.25 tsp)."},
			{StepNumber: 3, Instruction: "Marinate Chicken: Pour marinade over chicken, mix well."},
			{StepNumber: 4, Instruction: "Coat Chicken: Mix egg into chicken. Sprinkle cornflour over, mix until coated. Refrigerate 15-20 minutes."},
		},
		EquipmentIDs: []int64{54, 61, 62, 63, 64, 65}, // Med Bowl, Small Bowl, Colander, Board/Knives, Measure Spoons/Cups, Whisk/Fork
	},
	{
		ID:          1902,
		RecipeID:    19,
		StepOrder:   1,
		Title:       "Preparation Stage",
		Description: "Prepare, vegetables, and sauce.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string // Unit as specified IN THE RECIPE TEXT
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 29, Quantity: 0.5, Unit: "unit", Purpose: "stir-fry aromatic"},                  // Onion
			{FoodItemID: 31, Quantity: 2, Unit: "cloves", Purpose: "stir-fry aromatic"},                  // Garlic
			{FoodItemID: 50, Quantity: 75, Unit: "g", Purpose: "stir-fry veg"},                           // Carrot (assuming 1 medium)
			{FoodItemID: 125, Quantity: 40, Unit: "g", Purpose: "stir-fry veg"},                          // Green Beans (using handful estimate)
			{FoodItemID: 77, Quantity: 50, Unit: "g", Purpose: "stir-fry veg"},                           // Mushrooms (using handful estimate)
			{FoodItemID: 126, Quantity: 60, Unit: "g", Purpose: "stir-fry veg"},                          // Roasted Capsicum (using 1/4 cup estimate)
			{FoodItemID: 81, Quantity: 40, Unit: "g", Purpose: "stir-fry veg"},                           // Spinach (using handful estimate)
			{FoodItemID: 55, Quantity: 30, Unit: "ml", Purpose: "stir-fry sauce"},                        // Soy Sauce
			{FoodItemID: 57, Quantity: 10, Unit: "g", IsOptional: true, Purpose: "stir-fry sauce sweet"}, // Honey
			{FoodItemID: 25, Quantity: 60, Unit: "ml", Purpose: "stir-fry sauce liquid"},                 // Water (using max of 2-4)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Vegetable Prep: Slice onion, mince garlic (1-2 cloves), slice/julienne carrot, trim/halve green beans, slice mushrooms. Chop roasted capsicum. Wash spinach. Chop fresh herbs (if using)."},
			{StepNumber: 2, Instruction: "Stir-fry Sauce: Mix soy sauce (2 Tbsp), honey/sugar (if using), and water/stock (2-4 Tbsp) in a small bowl."},
		},
		EquipmentIDs: []int64{54, 61, 62, 63, 64, 65}, // Med Bowl, Small Bowl, Colander, Board/Knives, Measure Spoons/Cups, Whisk/Fork
	},
	{
		ID:          1903,
		RecipeID:    19,
		StepOrder:   1,
		Title:       "Cook noodles",
		Description: "Prepare noodles.",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string // Unit as specified IN THE RECIPE TEXT
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 90, Quantity: 150, Unit: "g", Purpose: "carb base"}, // Egg Noodles (estimating 1.5 portions @ 100g)
			// Note: Salt/Pepper to taste omitted, handled in method
			// Note: Sesame oil omitted, added in method step
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Noodle Prep: Cook egg noodles according to package directions until al dente. Drain immediately, rinse briefly with cold water. Toss with a tiny drizzle of oil if desired. Set aside."},
		},
		EquipmentIDs: []int64{54, 61, 62, 63, 64, 65}, // Med Bowl, Small Bowl, Colander, Board/Knives, Measure Spoons/Cups, Whisk/Fork
	},
	// --- Step 2: Cooking Stage ---
	{
		ID:          1904,
		RecipeID:    19,
		StepOrder:   2,
		Title:       "Fry Chicken",
		Description: "Fry chicken.",
		Ingredients: []struct { // Ingredients specifically *added* or manipulated in this stage
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Oil for stir-fry is listed as needed in method step 4
			{FoodItemID: 34, Quantity: 60, Unit: "ml", Purpose: "frying oil"}, // canola Oil
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat Wok & Oil: Place wok over high heat until very hot. Add 3-4 Tbsp canola/avocado oil and swirl. Heat oil until shimmering."},
			{StepNumber: 2, Instruction: "Fry Chicken: Carefully add marinated chicken pieces in a single layer (work in batches if needed). Reduce heat slightly if needed. Fry, turning occasionally, for 3-5 minutes per side until golden, crispy, and cooked through (75°C)."},
			{StepNumber: 3, Instruction: "Drain Chicken: Remove cooked chicken to a paper towel-lined plate. Keep warm."},
		},
		EquipmentIDs: []int64{60, 66, 67}, // Wok, Tongs/Spatula, Plate
	},
	{
		ID:          1905,
		RecipeID:    19,
		StepOrder:   2,
		Title:       "Stir-fry Noodles",
		Description: "Stir-fry vegetables and noodles.",
		Ingredients: []struct { // Ingredients specifically *added* or manipulated in this stage
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Oil for stir-fry is listed as needed in method step 4
			{FoodItemID: 34, Quantity: 15, Unit: "ml", Purpose: "cooking medium"}, // canola Oil
			{FoodItemID: 16, Quantity: 0, Unit: "g", Purpose: "final seasoning"},  // Salt
			{FoodItemID: 62, Quantity: 0, Unit: "ml", Purpose: "final seasoning"}, // Black Pepper
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat Wok & Oil: Place wok over high heat until very hot. Add canola/avocado oil and swirl. Heat oil until shimmering."},
			{StepNumber: 2, Instruction: "Stir-fry Veggies: Carefully pour out excess oil, leaving ~1 Tbsp. Return wok to high heat. Add onion, stir-fry ~30 seconds. Add minced garlic, carrots, green beans. Stir-fry constantly for 1-2 minutes."},
			{StepNumber: 3, Instruction: "Add Soft Veg: Add mushrooms, stir-fry 1-1.5 minutes until softening."},
			{StepNumber: 4, Instruction: "Combine & Sauce: Add roasted capsicum. Whisk stir-fry sauce, pour around wok edges. Toss quickly as sauce bubbles (~30 seconds)."},
			{StepNumber: 5, Instruction: "Wilt Spinach & Add Noodles: Add baby spinach, toss until just wilting. Add drained egg noodles. Toss gently but quickly to combine."},
		},
		EquipmentIDs: []int64{60, 66, 67}, // Wok, Tongs/Spatula, Plate
	},
	{
		ID:          1906,
		RecipeID:    19,
		StepOrder:   2,
		Title:       "Assemby",
		Description: "Assemble and serve.",
		Ingredients: []struct { // Ingredients specifically *added* or manipulated in this stage
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			// Oil for stir-fry is listed as needed in method step 4
			{FoodItemID: 116, Quantity: 2.5, Unit: "ml", IsOptional: true, Purpose: "finishing oil"}, // Sesame Oil
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Finish: Turn off heat. Stir through optional sesame oil (0.5 tsp). Season with salt and pepper to taste."},
			{StepNumber: 2, Instruction: "Serve: Transfer noodle stir-fry to bowls/plates. Top with crispy chicken. Garnish with fresh herbs if desired."},
		},
		EquipmentIDs: []int64{60, 66, 67}, // Wok, Tongs/Spatula, Plate
	},
}

var DummyRecipeTags = []struct {
	RecipeID int64
	TagID    int64
}{
	{RecipeID: 1, TagID: 80},  // Mac&Cheese -> American
	{RecipeID: 1, TagID: 81},  // Mac&Cheese -> Comfort
	{RecipeID: 1, TagID: 82},  // Mac&Cheese -> Pasta
	{RecipeID: 18, TagID: 89}, // Link Bao Buns recipe to Asian Fusion tag
	{RecipeID: 18, TagID: 90}, // Link Bao Buns recipe to Chinese Inspired tag
	{RecipeID: 18, TagID: 91}, // Link Bao Buns recipe to Buns tag
	{RecipeID: 18, TagID: 92}, // Link Bao Buns recipe to Spicy tag
	{RecipeID: 18, TagID: 88}, // Link Bao Buns recipe to Chicken tag (assuming ID 88 is Chicken)
	{RecipeID: 18, TagID: 93}, // Link Bao Buns recipe to Street Food tag
	// ... add links for Bao Buns Recipe ID 2 ...
}
