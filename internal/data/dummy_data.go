package data

import (
	"github.com/TJN25/recipe-site/internal/model"
)

var DummyFoodItems = []model.FoodItem{
	{ID: 10, Name: "Macaroni Pasta", BaseUnit: "g", PricePerBaseUnit: 0.005},
	{ID: 11, Name: "Everyday Cheese", BaseUnit: "g", PricePerBaseUnit: 0.02},
	{ID: 12, Name: "Tasty Cheddar", BaseUnit: "g", PricePerBaseUnit: 0.025},
	{ID: 13, Name: "Butter", BaseUnit: "g", PricePerBaseUnit: 0.01},
	{ID: 14, Name: "All‑Purpose Flour", BaseUnit: "g", PricePerBaseUnit: 0.002}, // Price per gram might be better base
	{ID: 15, Name: "Full Fat Milk", BaseUnit: "ml", PricePerBaseUnit: 0.002},
	{ID: 16, Name: "Salt", BaseUnit: "g", PricePerBaseUnit: 0.001},
	{ID: 17, Name: "Mustard Powder", BaseUnit: "g", PricePerBaseUnit: 0.05},
	{ID: 18, Name: "Smoked Paprika", BaseUnit: "g", PricePerBaseUnit: 0.06},
	{ID: 19, Name: "Chipotle Powder", BaseUnit: "g", PricePerBaseUnit: 0.08},
	{ID: 20, Name: "Mexican Chicken", BaseUnit: "g", PricePerBaseUnit: 0.015}, // Assume pre-cooked/shredded
	// ... add food items for Bao Buns etc ...
}

var DummyEquipment = []model.Equipment{
	{ID: 50, Name: "Large Saucepan", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyEasy},
	{ID: 51, Name: "Cheese Grater", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyMedium},
	{ID: 52, Name: "Casserole Dish", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyHard},
	{ID: 53, Name: "Whisk", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyMedium},
}

var DummyTags = []model.Tag{
	{ID: 80, Name: "American", Type: "cuisine"},
	{ID: 81, Name: "Comfort", Type: "mood"},
	{ID: 82, Name: "Pasta", Type: "dish_type"},
	// ... add tags for Bao Buns etc ...
}

var DummyRecipes = []struct {
	ID          int64
	Title       string
	Description string
	Servings    int
	Notes       string
	ImagePath   string
}{
	{
		ID:          1,
		Title:       "Chipotle Mexican Chicken Mac and Cheese",
		Description: "A smoky, spicy twist on classic mac and cheese...",
		Servings:    2,
		Notes:       "Let rest 5 minutes...",
		ImagePath:   "img/mac-and-cheese.png",
	},
	// ... add core data for Bao Buns Recipe ID 2 ...
}

type dummyRecipeStep struct {
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

var DummyRecipeSteps = []dummyRecipeStep{
	{ // Mac&Cheese - Step 1: Prep Pasta & Cheese
		ID:        101,
		RecipeID:  1,
		StepOrder: 1,
		Title:     "Prepare Pasta & Cheese",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 10, Quantity: 150, Unit: "g", Purpose: "starch base"},  // Macaroni Pasta
			{FoodItemID: 11, Quantity: 112, Unit: "g", Purpose: "melting base"}, // Everyday Cheese
			{FoodItemID: 12, Quantity: 113, Unit: "g", Purpose: "flavor"},       // Tasty Cheddar
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Cook macaroni pasta... Drain well."},
			{StepNumber: 2, Instruction: "While pasta cooks, grate cheeses..."},
		},
		EquipmentIDs: []int64{50, 51}, // Large Saucepan, Cheese Grater
	},
	{ // Mac&Cheese - Step 2: Make Cheese Sauce
		ID:        102,
		RecipeID:  1,
		StepOrder: 2,
		Title:     "Make Cheese Sauce",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 13, Quantity: 38, Unit: "g", Purpose: "richness"},                    // Butter
			{FoodItemID: 14, Quantity: 20, Unit: "g", Purpose: "thickener"},                   // Flour (adjust quantity for base unit g?)
			{FoodItemID: 15, Quantity: 265, Unit: "ml", Purpose: "sauce base"},                // Milk
			{FoodItemID: 16, Quantity: 3, Unit: "g", Purpose: "seasoning"},                    // Salt (adjust quantity?)
			{FoodItemID: 17, Quantity: 1, Unit: "g", Purpose: "enhance cheese"},               // Mustard Powder (adjust quantity?)
			{FoodItemID: 18, Quantity: 2, Unit: "g", Purpose: "smoky depth"},                  // Paprika (adjust quantity?)
			{FoodItemID: 19, Quantity: 1, Unit: "g", IsOptional: true, Purpose: "smoky heat"}, // Chipotle (adjust quantity?)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In the same saucepan... melt butter..."},
			{StepNumber: 2, Instruction: "Whisk in flour and cook..."},
			{StepNumber: 3, Instruction: "Gradually whisk in the milk..."},
			{StepNumber: 4, Instruction: "Remove from heat. Stir in cheese... Stir in spices..."},
		},
		EquipmentIDs: []int64{50, 53}, // Large Saucepan, Whisk
	},
	{ // Mac&Cheese - Step 3: Combine and Finish
		ID:        103,
		RecipeID:  1,
		StepOrder: 3,
		Title:     "Combine and Finish",
		Ingredients: []struct {
			FoodItemID int64
			Quantity   float32
			Unit       string
			IsOptional bool
			Purpose    string
		}{
			{FoodItemID: 20, Quantity: 150, Unit: "g", Purpose: "protein"}, // Mexican Chicken
			// Note: Reserved cheese isn't an *added* ingredient here, just used.
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Add the cooked pasta and chicken... Fold gently..."},
			{StepNumber: 2, Instruction: "Transfer the mixture to a casserole dish."},
			{StepNumber: 3, Instruction: "Sprinkle the reserved cheese..."},
			{StepNumber: 4, Instruction: "Place under a hot grill..."},
			{StepNumber: 5, Instruction: "Remove from grill and let rest..."},
		},
		EquipmentIDs: []int64{52}, // Casserole Dish
	},
	// ... add dummyRecipeSteps for Bao Buns Recipe ID 2 ...
}

var DummyRecipeTags = []struct {
	RecipeID int64
	TagID    int64
}{
	{RecipeID: 1, TagID: 80}, // Mac&Cheese -> American
	{RecipeID: 1, TagID: 81}, // Mac&Cheese -> Comfort
	{RecipeID: 1, TagID: 82}, // Mac&Cheese -> Pasta
	// ... add links for Bao Buns Recipe ID 2 ...
}

// var dummyRecipes = []model.Recipe{
// 	{
// 		ID:          1,
// 		Title:       "Chipotle Mexican Chicken Mac and Cheese",
// 		Description: "A smoky, spicy twist on classic mac and cheese, featuring chipotle chicken.",
// 		Servings:    2,
// 		Notes:       "Let rest 5 minutes before eating. Best fresh, but refrigerates up to 3 days.",
// 		ImagePath:   "img/mac-and-cheese.png",
// 		Tags: []model.Tag{
// 			{ID: 1, Name: "American", Type: "cuisine"},
// 			{ID: 2, Name: "Comfort", Type: "mood"},
// 		},
// 		Equipment: []model.Equipment{
// 			{ID: 1, Name: "Large Saucepan", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyMedium},
// 			{ID: 4, Name: "Small Saucepan", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyEasy},
// 			{ID: 2, Name: "Cheese Grater", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyMedium},
// 			{ID: 3, Name: "Knife", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyMedium},
// 			{ID: 5, Name: "Container", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},
// 		},
// 		Ingredients: []model.RecipeIngredient{
// 			{FoodItem: model.FoodItem{Name: "Everyday Cheese", BaseUnit: "g", PricePerBaseUnit: 0.02}, Quantity: 112, Unit: "g", Purpose: "melting base"},
// 			{FoodItem: model.FoodItem{Name: "Tasty Cheddar", BaseUnit: "g", PricePerBaseUnit: 0.025}, Quantity: 113, Unit: "g", Purpose: "flavor"},
// 			{FoodItem: model.FoodItem{Name: "Macaroni Pasta", BaseUnit: "g", PricePerBaseUnit: 0.005}, Quantity: 150, Unit: "g", Purpose: "starch base"},
// 			{FoodItem: model.FoodItem{Name: "Butter", BaseUnit: "g", PricePerBaseUnit: 0.01}, Quantity: 38, Unit: "g", Purpose: "richness"},
// 			{FoodItem: model.FoodItem{Name: "All‑Purpose Flour", BaseUnit: "tbsp", PricePerBaseUnit: 0.03}, Quantity: 2.25, Unit: "tbsp", Purpose: "thickener"},
// 			{FoodItem: model.FoodItem{Name: "Full Fat Milk", BaseUnit: "ml", PricePerBaseUnit: 0.002}, Quantity: 265, Unit: "ml", Purpose: "sauce base"},
// 			{FoodItem: model.FoodItem{Name: "Salt", BaseUnit: "tsp", PricePerBaseUnit: 0.005}, Quantity: 1.5, Unit: "tbsp", Purpose: "seasoning"},
// 			{FoodItem: model.FoodItem{Name: "Mustard Powder", BaseUnit: "tsp", PricePerBaseUnit: 0.02}, Quantity: 0.25, Unit: "tsp", IsOptional: true, Purpose: "enhance cheese flavor"},
// 			{FoodItem: model.FoodItem{Name: "Smoked Paprika", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 0.75, Unit: "tsp", IsOptional: true, Purpose: "smoky depth"},
// 			{FoodItem: model.FoodItem{Name: "Chipotle Powder", BaseUnit: "tsp", PricePerBaseUnit: 0.04}, Quantity: 0.5, Unit: "tsp", IsOptional: true, Purpose: "smoky heat"},
// 			{FoodItem: model.FoodItem{Name: "Mexican Chicken", BaseUnit: "g", PricePerBaseUnit: 0.015}, Quantity: 150, Unit: "g", IsOptional: true, Purpose: "protein"},
// 		},
// 		MethodSteps: []model.MethodStep{
// 			{StepNumber: 1, Instruction: "Grate cheeses and set aside 1/3 for topping.", Stage: "Prep/Roux", PrepTimeMinutes: 5},
// 			{StepNumber: 2, Instruction: "Cook macaroni pasta until 1–2 minutes past al dente.", Stage: "Pasta", CookTimeMinutes: 12},
// 			{StepNumber: 3, Instruction: "Melt butter, whisk in flour, cook 2 minutes.", Stage: "Prep/Roux", PrepTimeMinutes: 3},
// 			{StepNumber: 4, Instruction: "Whisk in milk gradually, cook until thickened.", Stage: "Prep/Roux", CookTimeMinutes: 5},
// 			{StepNumber: 5, Instruction: "Remove from heat, stir in cheese mix and any spices.", Stage: "Prep/Roux", PrepTimeMinutes: 2},
// 			{StepNumber: 6, Instruction: "Fold pasta into sauce in large ovenproof saucepan.", Stage: "Assembly", PrepTimeMinutes: 3},
// 			{StepNumber: 7, Instruction: "Top with reserved cheese, grill 5–6 minutes.", Stage: "Assembly", CookTimeMinutes: 6},
// 		},
// 	},
//
// 	{
// 		ID:          2,
// 		Title:       "Spiced Chicken Bao Buns with Sriracha Mayo",
// 		Description: "Crispy spiced chicken tucked into soft bao buns with cool slaw and spicy mayo.",
// 		Servings:    4,
// 		Notes:       "Assemble just before eating to keep chicken crispy.",
// 		ImagePath:   "img/bao-buns.png",
// 		Tags: []model.Tag{
// 			{ID: 3, Name: "Asian", Type: "cuisine"},
// 			{ID: 4, Name: "Snack", Type: "meal_type"},
// 		},
// 		Equipment: []model.Equipment{
// 			{ID: 4, Name: "Wok", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyMedium},
// 			{ID: 5, Name: "Steamer", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyMedium},
// 		},
// 		Ingredients: []model.RecipeIngredient{
// 			{FoodItem: model.FoodItem{Name: "Chicken Thigh", BaseUnit: "g", PricePerBaseUnit: 0.015}, Quantity: 300, Unit: "g", Purpose: "protein"},
// 			{FoodItem: model.FoodItem{Name: "Soy Sauce", BaseUnit: "tbsp", PricePerBaseUnit: 0.02}, Quantity: 1, Unit: "tbsp", Purpose: "umami"},
// 			{FoodItem: model.FoodItem{Name: "Sesame Oil", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 1, Unit: "tsp", Purpose: "aromatic"},
// 			{FoodItem: model.FoodItem{Name: "Garlic Powder", BaseUnit: "tsp", PricePerBaseUnit: 0.02}, Quantity: 1, Unit: "tsp", Purpose: "flavor"},
// 			{FoodItem: model.FoodItem{Name: "Ginger", BaseUnit: "tsp", PricePerBaseUnit: 0.02}, Quantity: 1, Unit: "tsp", Purpose: "warmth"},
// 			{FoodItem: model.FoodItem{Name: "Sichuan Pepper", BaseUnit: "tsp", PricePerBaseUnit: 0.04}, Quantity: 0.5, Unit: "tsp", Purpose: "numbing heat"},
// 			{FoodItem: model.FoodItem{Name: "Allspice", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 0.25, Unit: "tsp", Purpose: "complex spice"},
// 			{FoodItem: model.FoodItem{Name: "Smoked Paprika", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 0.5, Unit: "tsp", Purpose: "smoky depth"},
// 			{FoodItem: model.FoodItem{Name: "Bird’s Eye Chili", BaseUnit: "whole", PricePerBaseUnit: 0.05}, Quantity: 1, Unit: "whole", Purpose: "heat"},
// 			{FoodItem: model.FoodItem{Name: "Egg", BaseUnit: "whole", PricePerBaseUnit: 0.20}, Quantity: 1, Unit: "whole", Purpose: "binding"},
// 			{FoodItem: model.FoodItem{Name: "Cornflour", BaseUnit: "tbsp", PricePerBaseUnit: 0.01}, Quantity: 3, Unit: "tbsp", Purpose: "crispy coating"},
// 			{FoodItem: model.FoodItem{Name: "Canola Oil", BaseUnit: "tbsp", PricePerBaseUnit: 0.01}, Quantity: 3, Unit: "tbsp", Purpose: "frying"},
// 			{FoodItem: model.FoodItem{Name: "Kewpie Mayo", BaseUnit: "tbsp", PricePerBaseUnit: 0.05}, Quantity: 3, Unit: "tbsp", Purpose: "creamy base"},
// 			{FoodItem: model.FoodItem{Name: "Sriracha Sauce", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 2, Unit: "tsp", Purpose: "heat"},
// 			{FoodItem: model.FoodItem{Name: "Honey", BaseUnit: "tsp", PricePerBaseUnit: 0.04}, Quantity: 0.5, Unit: "tsp", Purpose: "sweetness"},
// 			{FoodItem: model.FoodItem{Name: "Lemon Juice", BaseUnit: "tsp", PricePerBaseUnit: 0.01}, Quantity: 1, Unit: "tsp", Purpose: "brightness"},
// 			{FoodItem: model.FoodItem{Name: "Bao Buns", BaseUnit: "whole", PricePerBaseUnit: 0.50}, Quantity: 4, Unit: "whole", Purpose: "vessel"},
// 			{FoodItem: model.FoodItem{Name: "Cabbage", BaseUnit: "g", PricePerBaseUnit: 0.005}, Quantity: 100, Unit: "g", Purpose: "crunch"},
// 			{FoodItem: model.FoodItem{Name: "Carrot", BaseUnit: "g", PricePerBaseUnit: 0.005}, Quantity: 50, Unit: "g", Purpose: "sweetness"},
// 			{FoodItem: model.FoodItem{Name: "Cucumber", BaseUnit: "g", PricePerBaseUnit: 0.005}, Quantity: 50, Unit: "g", Purpose: "coolness"},
// 			{FoodItem: model.FoodItem{Name: "Spring Onion", BaseUnit: "whole", PricePerBaseUnit: 0.10}, Quantity: 2, Unit: "whole", Purpose: "garnish"},
// 			{FoodItem: model.FoodItem{Name: "Roasted Peanuts", BaseUnit: "tbsp", PricePerBaseUnit: 0.02}, Quantity: 2, Unit: "tbsp", Purpose: "texture"},
// 		},
// 		MethodSteps: []model.MethodStep{
// 			{StepNumber: 1, Instruction: "Slice chicken and prepare marinade mix.", Stage: "Prep", PrepTimeMinutes: 10},
// 			{StepNumber: 2, Instruction: "Marinate chicken in fridge 15-20 minutes.", Stage: "Prep", CookTimeMinutes: 20},
// 			{StepNumber: 3, Instruction: "Fry chicken in batches until golden.", Stage: "Cooking", CookTimeMinutes: 3},
// 			{StepNumber: 4, Instruction: "Steam bao buns per package instructions.", Stage: "Assembly", CookTimeMinutes: 4},
// 			{StepNumber: 5, Instruction: "Assemble buns with slaw, chicken, mayo, peanuts, onions.", Stage: "Assembly", PrepTimeMinutes: 5},
// 		},
// 	},
// }
