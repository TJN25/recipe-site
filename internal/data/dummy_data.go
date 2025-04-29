package data

import (
	"github.com/TJN25/recipe-site/internal/model"
)

var placeholderNutrition = model.Nutrition{}
var placeHolderPrice = float32(1.0)

var DummyFoodItems = []model.FoodItem{
	{
		ID:              1,
		Name:            "Chicken",
		CanonicalUnit:   "g",
		DefaultFormName: "Chicken Breast",
		Forms: map[string]model.FoodItemFormDetails{
			"Chicken Breast":  {Unit: "g", ConversionToCanonical: 1.0}, // Base unit is grams
			"Chicken Thigh":   {Unit: "g", ConversionToCanonical: 1.0}, // Base unit is grams
			"Mexican Chicken": {Unit: "g", ConversionToCanonical: 1.0}, // Base unit is grams
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              2,
		Name:            "Garlic",
		CanonicalUnit:   "g",
		DefaultFormName: "Garlic Clove",
		Forms: map[string]model.FoodItemFormDetails{
			"Garlic Clove":  {Unit: "clove", ConversionToCanonical: 5.0},
			"Garlic Paste":  {Unit: "ml", ConversionToCanonical: 1.0},
			"Garlic Powder": {Unit: "ml", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              3,
		Name:            "Ginger",
		CanonicalUnit:   "g",
		DefaultFormName: "Fresh Ginger Root",
		Forms: map[string]model.FoodItemFormDetails{
			"Fresh Ginger Root": {Unit: "g", ConversionToCanonical: 1.0}, // often recorded as cm ?
			"Paste":             {Unit: "ml", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              4,
		Name:            "Onion",
		CanonicalUnit:   "g",
		DefaultFormName: "Brown Onion",
		Forms: map[string]model.FoodItemFormDetails{
			"Brown Onion":  {Unit: "unit", ConversionToCanonical: 150.0},
			"Red Onion":    {Unit: "unit", ConversionToCanonical: 150.0},
			"Shallot":      {Unit: "unit", ConversionToCanonical: 150.0}, // placeholder
			"Spring Onion": {Unit: "unit", ConversionToCanonical: 150.0}, // placeholder
			"Leek Onion":   {Unit: "unit", ConversionToCanonical: 150.0}, // placeholder
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              5,
		Name:            "Cheese",
		CanonicalUnit:   "g",
		DefaultFormName: "Cheddar Cheese",
		Forms: map[string]model.FoodItemFormDetails{
			"Cheddar Cheese":    {Unit: "g", ConversionToCanonical: 1.0},
			"Everday Cheese":    {Unit: "g", ConversionToCanonical: 1.0},
			"Parmesan Cheese":   {Unit: "g", ConversionToCanonical: 1.0},
			"Mozzarella Cheese": {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              6,
		Name:            "Tomato",
		CanonicalUnit:   "g",
		DefaultFormName: "Fresh Tomato",
		Forms: map[string]model.FoodItemFormDetails{
			"Fresh Tomato":           {Unit: "unit", ConversionToCanonical: 120.0},
			"Canned Tomates (whole)": {Unit: "unit", ConversionToCanonical: 400.0},
			"Canned Tomates (diced)": {Unit: "unit", ConversionToCanonical: 400.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              7,
		Name:            "Sugar",
		CanonicalUnit:   "g",
		DefaultFormName: "Raw Sugar",
		Forms: map[string]model.FoodItemFormDetails{
			"Raw Sugar":    {Unit: "g", ConversionToCanonical: 1.0},
			"Brown Sugar":  {Unit: "g", ConversionToCanonical: 1.0},
			"Caster Sugar": {Unit: "g", ConversionToCanonical: 1.0},
			"Icing Sugar":  {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              8,
		Name:            "Rice",
		CanonicalUnit:   "g",
		DefaultFormName: "Basmati Rice",
		Forms: map[string]model.FoodItemFormDetails{
			"Basmati Rice": {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              9,
		Name:            "Salt",
		CanonicalUnit:   "g",
		DefaultFormName: "Kosher Salt",
		Forms: map[string]model.FoodItemFormDetails{
			"Kosher Salt": {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              10,
		Name:            "Pepper",
		CanonicalUnit:   "g",
		DefaultFormName: "Ground Black Pepper",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground Black Pepper": {Unit: "g", ConversionToCanonical: 1.0},
			"Ground White Pepper": {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              11,
		Name:            "Pasta",
		CanonicalUnit:   "g",
		DefaultFormName: "Fettuccine",
		Forms: map[string]model.FoodItemFormDetails{
			"Fettuccine": {Unit: "g", ConversionToCanonical: 1.0},
			"Macaroni":   {Unit: "g", ConversionToCanonical: 1.0},
			"Spaghetti":  {Unit: "g", ConversionToCanonical: 1.0},
			"Lasagna":    {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              12,
		Name:            "Cumin",
		CanonicalUnit:   "g",
		DefaultFormName: "Ground Cumin",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground Cumin": {Unit: "g", ConversionToCanonical: 1.0},
			"Cumin Seed":   {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              13,
		Name:            "Coriander Spice",
		CanonicalUnit:   "g",
		DefaultFormName: "Ground Coriander",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground Coriander": {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              14,
		Name:            "Mustard",
		CanonicalUnit:   "g",
		DefaultFormName: "Mustard Powder",
		Forms: map[string]model.FoodItemFormDetails{
			"Mustard Powder":   {Unit: "ml", ConversionToCanonical: 1.0},
			"American Mustard": {Unit: "ml", ConversionToCanonical: 1.0}, // placeholder
			"Dijon Mustard":    {Unit: "ml", ConversionToCanonical: 1.0}, // placeholder
			"Honey Mustard":    {Unit: "ml", ConversionToCanonical: 1.0}, // placeholder
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              15,
		Name:            "Milk",
		CanonicalUnit:   "ml",
		DefaultFormName: "Full Fat Milk",
		Forms: map[string]model.FoodItemFormDetails{
			"Full Fat Milk": {Unit: "ml", ConversionToCanonical: 1.0},
			"Trim Milk":     {Unit: "ml", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              16,
		Name:            "Smoked Paprika",
		CanonicalUnit:   "g",
		DefaultFormName: "Smoked Paprika",
		Forms: map[string]model.FoodItemFormDetails{
			"Smoked Paprika": {Unit: "ml", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
	{
		ID:              17,
		Name:            "Flour",
		CanonicalUnit:   "g",
		DefaultFormName: "Flour",
		Forms: map[string]model.FoodItemFormDetails{
			"Flour":             {Unit: "g", ConversionToCanonical: 1.0},
			"All Purpose Flour": {Unit: "g", ConversionToCanonical: 1.0},
		},
		PricePerCanonicalUnit: placeHolderPrice,
		Nutrition:             placeholderNutrition,
	},
}

var DummyEquipment = []model.Equipment{}

var DummyTags = []model.Tag{}

var DummyRecipes = []struct {
	ID            int64
	RecipeStepIds []int64
	Title         string
	Description   string
	Servings      int
	Notes         string
	ImagePath     string
}{
	{
		ID:            1,
		RecipeStepIds: []int64{101, 102, 103},
		Title:         "Mac and Cheese",
		Description:   "Classic mac and cheese",
		Servings:      2,
		Notes:         "Let rest 5 minutes before serving. Uses full fat milk for creaminess. Sauce should look slightly too saucy before baking; it will thicken.",
		ImagePath:     "img/mac-and-cheese.png",
	},
	{
		ID:            2,
		RecipeStepIds: []int64{1801, 1802, 1803, 1804, 1805},
		Title:         "Spiced Chicken Bao Buns with Sriracha Mayo",
		Description:   "Crispy fried spiced chicken pieces served in soft bao buns with coleslaw, cucumber, peanuts, and a tangy sriracha mayo.",
		Servings:      3,
		Notes:         "Marinate chicken for 15-20 mins. Fry chicken in batches. Steam buns just before serving. Assemble just before eating for best texture contrast. Enhancements: Salt cucumber slices briefly before use. Toast peanuts before crushing.",
		ImagePath:     "img/bao-buns-chicken.png", // Assign an appropriate image path
	},
	{
		ID:            3,
		RecipeStepIds: []int64{1801, 1902, 1803, 1903, 1905, 1906},
		Title:         "Crispy Fried Chicken with Vegetable & Egg Noodle Stir-fry",
		Description:   "Crispy fried chicken pieces served over a flavorful stir-fry of vegetables and egg noodles.",
		Servings:      2,
		Notes:         "Wok Hei: Preheating the wok properly over high heat is crucial for achieving \"wok hei\" - the characteristic smoky flavour of good stir-fries.\nWok Frying: Be mindful when deep-frying/shallow-frying in a wok. The sloped sides mean oil depth varies. Keep pieces moving and adjust heat to prevent burning. Use a wok spatula or spider strainer for removal.\nStir-fry Motion: Use a scooping, tossing motion to move ingredients constantly, ensuring even cooking and preventing sticking. Add sauce around the perimeter to allow it to heat and reduce slightly before coating ingredients.\nSpeed: Wok cooking is fast. Have everything prepped and ready next to the stove before you start heating the wok.",
		ImagePath:     "img/chicken-noodle-stirfry.png", // Placeholder path
	},
}

type dummyIngredientRef struct {
	FoodItemID   int64
	FoodItemName string
	FormName     string
	Quantity     float32
	Unit         string // Keep this for initial loading consistency check
	IsOptional   bool
	Purpose      string
}

type DummyRecipeStep struct {
	ID          int64
	RecipeID    int64
	StepOrder   int
	Title       string
	Description string
	Notes       string
	Ingredients []dummyIngredientRef
	MethodSteps []model.MethodStep // Method sub-steps (can be full model struct)
}

var DummyRecipeSteps = []DummyRecipeStep{
	// --- Recipe 1: Chipotle Mexican Chicken Mac and Cheese ---
	{
		ID:          101,
		RecipeID:    1,
		StepOrder:   1,
		Title:       "Cook Pasta",
		Description: "Cook pasta until slightly overcooked.",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Pasta", FormName: "Macaroni", Quantity: 150, Unit: "g", Purpose: "starch base"},            // Macaroni Pasta
			{FoodItemName: "Salt", FormName: "Kosher Salt", Quantity: 15, Unit: "g", Purpose: "pasta water seasoning"}, // Salt for pasta water (approx 1.5 tbsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Bring a large saucepan of water to a boil. Add salt."},
			{StepNumber: 2, Instruction: "Add macaroni pasta and cook for 1-2 minutes past al dente (approx. 10-12 minutes total)."},
			{StepNumber: 4, Instruction: "Drain cooked pasta well, but do not rinse."},
		},
	},
	{
		ID:          102,
		RecipeID:    1,
		StepOrder:   2,
		Title:       "Make Cheese Sauce",
		Description: "Create a roux, add milk, then melt in cheese and seasonings.",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Cheese", FormName: "Everday Cheese", Quantity: 112, Unit: "g", Purpose: "melting base"}, // Everyday Cheese
			{FoodItemName: "Cheese", FormName: "Cheddar Cheese", Quantity: 113, Unit: "g", Purpose: "flavor"},       // Tasty Cheddar
			{FoodItemName: "Butter", Quantity: 38, Unit: "g", Purpose: "richness, roux base"},                       // Butter
			{FoodItemName: "Flour", FormName: "All Purpose Flour", Quantity: 35, Unit: "ml", Purpose: "thickener"},  // Flour (~2.25 tbsp)
			{FoodItemName: "Smoked Paprika", Quantity: 4, IsOptional: true, Unit: "ml", Purpose: "smoky depth"},     // Smoked Paprika (~3/4 tsp) - OPTIONAL for standard
			{FoodItemName: "Milk", Quantity: 265, Unit: "ml", Purpose: "sauce base"},                                // Full Fat Milk
			{FoodItemName: "Mustard", Quantity: 1.25, Unit: "ml", Purpose: "enhance cheese"},                        // Mustard Powder (~¼ tsp)
			{FoodItemName: "Salt", Quantity: 2, Unit: "g", Purpose: "seasoning"},                                    // Salt (to taste)
			{FoodItemName: "Pepper", Quantity: 1, Unit: "g", Purpose: "seasoning"},                                  // White Pepper (to taste)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "While pasta cooks, grate both cheeses separately. Set aside 1/3 of each for topping."},
			{StepNumber: 2, Instruction: "In the same large saucepan used for pasta, melt butter over medium heat."},
			{StepNumber: 3, Instruction: "Once melted, add flour. Whisk constantly for 1 minute."},
			{StepNumber: 4, Instruction: "If making Chipotle version, add smoked paprika. Stir to bloom spices (30 seconds)."},
			{StepNumber: 5, Instruction: "Continue cooking the flour mixture (roux) for a total of 2 minutes, stirring constantly."},
			{StepNumber: 6, Instruction: "Gradually whisk in the full fat milk, starting with about 1/3 cup, whisking until smooth before adding more."},
			{StepNumber: 7, Instruction: "Bring the mixture to a gentle simmer and cook, stirring constantly, until thickened (4-5 minutes)."},
			{StepNumber: 8, Instruction: "Remove the saucepan from heat."},
			{StepNumber: 9, Instruction: "Add the remaining grated cheeses and mustard powder."},
			{StepNumber: 10, Instruction: "Season with salt and white pepper to taste."},
			{StepNumber: 11, Instruction: "Stir until the cheese is completely melted and the sauce is smooth."},
		},
	},
	// Step 3: Combine and Finish (Merged Standard & Chipotle)
	{
		ID:          103,
		RecipeID:    1,
		StepOrder:   3,
		Title:       "Combine, Top, and Grill",
		Description: "Combine pasta with sauce, transfer to dish, top with cheese, and grill.",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Smoked Paprika", Quantity: 1, Unit: "pinch", IsOptional: true, Purpose: "visual appeal"}, // Smoked Paprika (light dusting) - OPTIONAL
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Add the cooked pasta to the cheese sauce."},
			{StepNumber: 3, Instruction: "Mix gently until everything is coated. The mixture should look slightly too saucy."},
			{StepNumber: 4, Instruction: "Transfer the mixture to a medium oven-safe casserole dish."},
			{StepNumber: 5, Instruction: "Sprinkle the reserved cheese mix (74g) evenly over the top."},
			{StepNumber: 6, Instruction: "Optional: Sprinkle a light dusting of smoked paprika on top."},
			{StepNumber: 7, Instruction: "Place the dish under a hot grill (or in a toaster oven on grill setting)."},
			{StepNumber: 8, Instruction: "Grill for 5-6 minutes, or until the top is golden brown and bubbly with some darker spots."},
			{StepNumber: 9, Instruction: "Remove from the grill and let rest for 5 minutes before serving."},
		},
	},
	{
		ID:          1801, // Assign next available ID block start
		RecipeID:    18,   // Link to the new Bao Bun recipe ID
		StepOrder:   1,
		Title:       "Prep Chicken & Marinade",
		Description: "Slice chicken, crush chili, mix marinade ingredients and coat chicken.",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Chicken", FormName: "Chicken Thigh", Quantity: 300, Unit: "g", Purpose: "protein"},       // Chicken Thigh
			{FoodItemName: "Chili", FormName: "Bird's Eye Chili", Quantity: 1, Unit: "unit", Purpose: "heat"},        // Dried Bird's Eye Chili
			{FoodItemName: "Soy sauce", Quantity: 15, Unit: "ml", Purpose: "marinade base"},                          // Soy Sauce (1 Tbsp)
			{FoodItemName: "Oil", FormName: "Sesame Oil", Quantity: 5, Unit: "ml", Purpose: "marinade aroma"},        // Sesame Oil (1 tsp)
			{FoodItemName: "Garlic", FormName: "Garlic Powder", Quantity: 5, Unit: "ml", Purpose: "marinade flavor"}, // Garlic Powder/Minced (1 tsp)
			{FoodItemName: "ginger", FormName: "Ginger paste", Quantity: 5, Unit: "ml", Purpose: "marinade flavor"},  // Ground Ginger/Paste (1 tsp)
			{FoodItemName: "Sichuan pepper", Quantity: 2, Unit: "g", Purpose: "marinade spice"},                      // Sichuan Pepper, ground (1/2 tsp)
			{FoodItemName: "Allspice", Quantity: 1, Unit: "g", Purpose: "marinade spice"},                            // Allspice, ground (1/4 tsp)
			{FoodItemName: "Smoked paprika", Quantity: 2.5, Unit: "g", Purpose: "marinade spice"},                    // Smoked Paprika (1/2 tsp)
			{FoodItemName: "Egg", Quantity: 1, Unit: "unit", Purpose: "binding"},                                     // Egg
			{FoodItemName: "Cornflour", Quantity: 45, Unit: "ml", Purpose: "coating"},                                // Cornflour (3 Tbsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Slice chicken thighs into bite-sized pieces (approx 1cm thick strips or cubes). Place in a medium bowl."},
			{StepNumber: 2, Instruction: "Crush the dried bird's eye chili using a mortar and pestle or the flat side of a knife."},
			{StepNumber: 3, Instruction: "In a small ramekin or bowl, mix soy sauce, sesame oil, garlic, ginger, Sichuan pepper, allspice, smoked paprika, and the crushed chili."},
			{StepNumber: 4, Instruction: "Pour the spice mixture over the chicken pieces and mix thoroughly to coat."},
			{StepNumber: 5, Instruction: "Add the egg and cornflour to the chicken bowl. Mix well until chicken is fully coated."},
			{StepNumber: 6, Instruction: "Cover and marinate in the fridge for 15-20 minutes."},
		},
	},
	{
		ID:          1802,
		RecipeID:    18,
		StepOrder:   2,
		Title:       "Prep Slaw, Mayo & Garnishes",
		Description: "Prepare coleslaw mix, cucumber, peanuts, sriracha mayo, and spring onions.",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Cabbage", Quantity: 100, Unit: "g", Purpose: "slaw base"},                                 // Cabbage, shredded
			{FoodItemName: "Carrot", Quantity: 50, Unit: "g", Purpose: "slaw color"},                                  // Carrot, grated
			{FoodItemName: "Cucumber", Quantity: 100, Unit: "g", Purpose: "freshness"},                                // Cucumber (1/2), thinly sliced
			{FoodItemName: "Peanut", FormName: "Roasted Peanut", Quantity: 25, Unit: "g", Purpose: "garnish texture"}, // Roasted Peanuts (2 Tbsp), crushed
			{FoodItemName: "Mayonnaise", FormName: "Kewpie Mayo", Quantity: 45, Unit: "g", Purpose: "mayo base"},      // Kewpie Mayo (3 Tbsp)
			{FoodItemName: "Sirracha Sauce", Quantity: 10, Unit: "ml", Purpose: "mayo heat"},                          // Sriracha Sauce (1-2 tsp)
			{FoodItemName: "Honey", Quantity: 2.5, Unit: "g", Purpose: "mayo balance"},                                // Honey (1/2 tsp)
			{FoodItemName: "Lemon", FormName: "Lemon Juice", Quantity: 1, Unit: "ml", Purpose: "mayo brightness"},     // Lemon Juice (Few drops)
			{FoodItemName: "Oil", FormName: "Sesame Oil", Quantity: 1, Unit: "ml", Purpose: "mayo aroma"},             // Sesame Oil (1/4 tsp / drop)
			{FoodItemName: "Onion", FormName: "Spring Onion", Quantity: 20, Unit: "g", Purpose: "garnish"},            // Spring Onion, thinly sliced
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Combine shredded cabbage and grated carrot in a bowl for the coleslaw base."},
			{StepNumber: 2, Instruction: "Thinly slice the cucumber into half-moons or matchsticks. (Optional: salt briefly and pat dry)."},
			{StepNumber: 3, Instruction: "Crush the roasted peanuts using a mortar and pestle or place in a bag and crush with a rolling pin or heavy object."},
			{StepNumber: 4, Instruction: "Make Sriracha Mayo: In a small bowl, mix Kewpie mayo, sriracha sauce (start with less, add more to taste), honey, lemon juice, and sesame oil until combined."},
			{StepNumber: 5, Instruction: "Thinly slice the spring onions."},
		},
	},
	{
		ID:          1803,
		RecipeID:    18,
		StepOrder:   3,
		Title:       "Fry Chicken & Steam Buns",
		Description: "Fry marinated chicken until crispy. Steam bao buns.",
		Ingredients: []dummyIngredientRef{
			// Marinated Chicken from Step 1
			{FoodItemName: "Oil", Quantity: 45, Unit: "ml", Purpose: "frying"}, // Canola Oil (3 Tbsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat canola oil in a wok or large, deep skillet over medium-high heat."},
			{StepNumber: 2, Instruction: "Once oil is hot (a piece of batter sizzles immediately), carefully add chicken pieces in a single layer, working in batches to avoid overcrowding."},
			{StepNumber: 3, Instruction: "Fry chicken for 2-3 minutes per side, or until golden brown, crispy, and cooked through."},
			{StepNumber: 4, Instruction: "Use tongs or a slotted spoon to remove cooked chicken and place on a plate lined with paper towels to drain."},
			{StepNumber: 5, Instruction: "Repeat with remaining chicken."},
		},
	},
	{
		ID:          1804,
		RecipeID:    18,
		StepOrder:   3,
		Title:       "Fry Chicken & Steam Buns",
		Description: "Fry marinated chicken until crispy. Steam bao buns.",
		Ingredients: []dummyIngredientRef{
			// Marinated Chicken from Step 1
			{FoodItemName: "Bun", FormName: "Bao Buns", Quantity: 8, Unit: "unit", Purpose: "vessel"}, // Bao Buns (adjust number as needed)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "While the last batch of chicken is frying, steam the bao buns according to package instructions (usually 3-4 minutes over simmering water)."},
		},
	},
	{
		ID:          1805,
		RecipeID:    18,
		StepOrder:   4,
		Title:       "Assemble Bao Buns",
		Description: "Layer mayo, slaw, cucumber, chicken, and garnishes in steamed buns.",
		Ingredients: []dummyIngredientRef{},
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
	},
	{
		ID:          1902,
		RecipeID:    19,
		StepOrder:   1,
		Title:       "Preparation Stage",
		Description: "Prepare, vegetables, and sauce.",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Onion", Quantity: 0.5, Unit: "unit", Purpose: "stir-fry aromatic"},                         // Onion
			{FoodItemName: "Garlic", Quantity: 2, Unit: "cloves", Purpose: "stir-fry aromatic"},                        // Garlic
			{FoodItemName: "Carrot", Quantity: 75, Unit: "g", Purpose: "stir-fry veg"},                                 // Carrot (assuming 1 medium)
			{FoodItemName: "Green bean", Quantity: 40, Unit: "g", Purpose: "stir-fry veg"},                             // Green Beans (using handful estimate)
			{FoodItemName: "Mushrooms", Quantity: 50, Unit: "g", Purpose: "stir-fry veg"},                              // Mushrooms (using handful estimate)
			{FoodItemName: "Capsicum", FormName: "Roasted Capsicum", Quantity: 60, Unit: "g", Purpose: "stir-fry veg"}, // Roasted Capsicum (using 1/4 cup estimate)
			{FoodItemName: "Spinach", Quantity: 40, Unit: "g", Purpose: "stir-fry veg"},                                // Spinach (using handful estimate)
			{FoodItemName: "Soy Sauce", Quantity: 30, Unit: "ml", Purpose: "stir-fry sauce"},                           // Soy Sauce
			{FoodItemName: "Honey", Quantity: 10, Unit: "g", IsOptional: true, Purpose: "stir-fry sauce sweet"},        // Honey
			{FoodItemName: "Water", Quantity: 60, Unit: "ml", Purpose: "stir-fry sauce liquid"},                        // Water (using max of 2-4)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Vegetable Prep: Slice onion, mince garlic (1-2 cloves), slice/julienne carrot, trim/halve green beans, slice mushrooms. Chop roasted capsicum. Wash spinach. Chop fresh herbs (if using)."},
			{StepNumber: 2, Instruction: "Stir-fry Sauce: Mix soy sauce (2 Tbsp), honey/sugar (if using), and water/stock (2-4 Tbsp) in a small bowl."},
		},
	},
	{
		ID:          1903,
		RecipeID:    19,
		StepOrder:   1,
		Title:       "Cook noodles",
		Description: "Prepare noodles.",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Noodles", FormName: "Egg Noodles", Quantity: 150, Unit: "g", Purpose: "carb base"}, // Egg Noodles (estimating 1.5 portions @ 100g)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Noodle Prep: Cook egg noodles according to package directions until al dente. Drain immediately, rinse briefly with cold water. Toss with a tiny drizzle of oil if desired. Set aside."},
		},
	},
	// --- Step 2: Cooking Stage ---
	{
		ID:          1905,
		RecipeID:    19,
		StepOrder:   2,
		Title:       "Stir-fry Noodles",
		Description: "Stir-fry vegetables and noodles.",
		Ingredients: []dummyIngredientRef{
			// Oil for stir-fry is listed as needed in method step 4
			{FoodItemName: "Oil", Quantity: 15, Unit: "ml", Purpose: "cooking medium"},    // canola Oil
			{FoodItemName: "Salt", Quantity: 0, Unit: "g", Purpose: "final seasoning"},    // Salt
			{FoodItemName: "Pepper", Quantity: 0, Unit: "ml", Purpose: "final seasoning"}, // Black Pepper
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat Wok & Oil: Place wok over high heat until very hot. Add canola/avocado oil and swirl. Heat oil until shimmering."},
			{StepNumber: 2, Instruction: "Stir-fry Veggies: Carefully pour out excess oil, leaving ~1 Tbsp. Return wok to high heat. Add onion, stir-fry ~30 seconds. Add minced garlic, carrots, green beans. Stir-fry constantly for 1-2 minutes."},
			{StepNumber: 3, Instruction: "Add Soft Veg: Add mushrooms, stir-fry 1-1.5 minutes until softening."},
			{StepNumber: 4, Instruction: "Combine & Sauce: Add roasted capsicum. Whisk stir-fry sauce, pour around wok edges. Toss quickly as sauce bubbles (~30 seconds)."},
			{StepNumber: 5, Instruction: "Wilt Spinach & Add Noodles: Add baby spinach, toss until just wilting. Add drained egg noodles. Toss gently but quickly to combine."},
		},
	},
	{
		ID:          1906,
		RecipeID:    19,
		StepOrder:   2,
		Title:       "Assemby",
		Description: "Assemble and serve.",
		Ingredients: []dummyIngredientRef{
			// Oil for stir-fry is listed as needed in method step 4
			{FoodItemName: "Oil", FormName: "Sesame Oil", Quantity: 2.5, Unit: "ml", IsOptional: true, Purpose: "finishing oil"}, // Sesame Oil
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Finish: Turn off heat. Stir through optional sesame oil (0.5 tsp). Season with salt and pepper to taste."},
			{StepNumber: 2, Instruction: "Serve: Transfer noodle stir-fry to bowls/plates. Top with crispy chicken. Garnish with fresh herbs if desired."},
		},
	},
}

var DummyRecipeTags = []struct {
	RecipeID int64
	TagID    int64
}{}
