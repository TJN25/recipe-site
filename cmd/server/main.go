package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/TJN25/recipe-site/internal/model"
	log "github.com/sirupsen/logrus"
)

// some relational linking of all the recipes by a score (shared ingredients, shared tags, shared equipment)
// Matrix of Recipes, by Recipes with score (but maybe more efficient?)

var templates *template.Template

var dummyRecipes = []model.Recipe{
	{
		ID:          1,
		Title:       "Chipotle Mexican Chicken Mac and Cheese",
		Description: "A smoky, spicy twist on classic mac and cheese, featuring chipotle chicken.",
		Servings:    2,
		Notes:       "Let rest 5 minutes before eating. Best fresh, but refrigerates up to 3 days.",
		ImagePath:   "recipe-images/mac-and-cheese.jpg",
		Tags: []model.Tag{
			{ID: 1, Name: "American", Type: "cuisine"},
			{ID: 2, Name: "Comfort", Type: "mood"},
		},
		Equipment: []model.Equipment{
			{ID: 1, Name: "Large Saucepan", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyEasy},
			{ID: 2, Name: "Cheese Grater", Type: "Utensil", CleaningDifficulty: model.CleaningDifficultyEasy},
			{ID: 3, Name: "Casserole Dish", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyHard},
		},
		Ingredients: []model.RecipeIngredient{
			{FoodItem: model.FoodItem{Name: "Everyday Cheese", BaseUnit: "g", PricePerBaseUnit: 0.02}, Quantity: 112, Unit: "g", Purpose: "melting base"},
			{FoodItem: model.FoodItem{Name: "Tasty Cheddar", BaseUnit: "g", PricePerBaseUnit: 0.025}, Quantity: 113, Unit: "g", Purpose: "flavor"},
			{FoodItem: model.FoodItem{Name: "Macaroni Pasta", BaseUnit: "g", PricePerBaseUnit: 0.005}, Quantity: 150, Unit: "g", Purpose: "starch base"},
			{FoodItem: model.FoodItem{Name: "Butter", BaseUnit: "g", PricePerBaseUnit: 0.01}, Quantity: 38, Unit: "g", Purpose: "richness"},
			{FoodItem: model.FoodItem{Name: "All‑Purpose Flour", BaseUnit: "tbsp", PricePerBaseUnit: 0.03}, Quantity: 2.25, Unit: "tbsp", Purpose: "thickener"},
			{FoodItem: model.FoodItem{Name: "Full Fat Milk", BaseUnit: "ml", PricePerBaseUnit: 0.002}, Quantity: 265, Unit: "ml", Purpose: "sauce base"},
			{FoodItem: model.FoodItem{Name: "Salt", BaseUnit: "tsp", PricePerBaseUnit: 0.005}, Quantity: 1.5, Unit: "tbsp", Purpose: "seasoning"},
			{FoodItem: model.FoodItem{Name: "Mustard Powder", BaseUnit: "tsp", PricePerBaseUnit: 0.02}, Quantity: 0.25, Unit: "tsp", Purpose: "enhance cheese flavor"},
			{FoodItem: model.FoodItem{Name: "Smoked Paprika", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 0.75, Unit: "tsp", Purpose: "smoky depth"},
			{FoodItem: model.FoodItem{Name: "Chipotle Powder", BaseUnit: "tsp", PricePerBaseUnit: 0.04}, Quantity: 0.5, Unit: "tsp", IsOptional: true, Purpose: "smoky heat"},
			{FoodItem: model.FoodItem{Name: "Mexican Chicken", BaseUnit: "g", PricePerBaseUnit: 0.015}, Quantity: 150, Unit: "g", Purpose: "protein"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Grate cheeses and set aside 74g for topping.", Stage: "Prep/Roux", PrepTimeMinutes: 5},
			{StepNumber: 2, Instruction: "Cook macaroni pasta until 1–2 minutes past al dente.", Stage: "Pasta", CookTimeMinutes: 12},
			{StepNumber: 3, Instruction: "Melt butter, whisk in flour, cook 2 minutes.", Stage: "Prep/Roux", PrepTimeMinutes: 3},
			{StepNumber: 4, Instruction: "Whisk in milk gradually, cook until thickened.", Stage: "Prep/Roux", CookTimeMinutes: 5},
			{StepNumber: 5, Instruction: "Remove from heat, stir in cheese mix and spices.", Stage: "Prep/Roux", PrepTimeMinutes: 2},
			{StepNumber: 6, Instruction: "Fold in chicken and pasta, transfer to dish.", Stage: "Assembly", PrepTimeMinutes: 3},
			{StepNumber: 7, Instruction: "Top with reserved cheese, grill 5–6 minutes.", Stage: "Assembly", CookTimeMinutes: 6},
		},
	},

	{
		ID:          2,
		Title:       "Spiced Chicken Bao Buns with Sriracha Mayo",
		Description: "Crispy spiced chicken tucked into soft bao buns with cool slaw and spicy mayo.",
		Servings:    4,
		Notes:       "Assemble just before eating to keep chicken crispy.",
		ImagePath:   "recipe-images/bao-buns.jpg",
		Tags: []model.Tag{
			{ID: 3, Name: "Asian", Type: "cuisine"},
			{ID: 4, Name: "Snack", Type: "meal_type"},
		},
		Equipment: []model.Equipment{
			{ID: 4, Name: "Wok", Type: "Cookware", CleaningDifficulty: model.CleaningDifficultyMedium},
			{ID: 5, Name: "Steamer", Type: "Appliance", CleaningDifficulty: model.CleaningDifficultyMedium},
		},
		Ingredients: []model.RecipeIngredient{
			{FoodItem: model.FoodItem{Name: "Chicken Thigh", BaseUnit: "g", PricePerBaseUnit: 0.015}, Quantity: 300, Unit: "g", Purpose: "protein"},
			{FoodItem: model.FoodItem{Name: "Soy Sauce", BaseUnit: "tbsp", PricePerBaseUnit: 0.02}, Quantity: 1, Unit: "tbsp", Purpose: "umami"},
			{FoodItem: model.FoodItem{Name: "Sesame Oil", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 1, Unit: "tsp", Purpose: "aromatic"},
			{FoodItem: model.FoodItem{Name: "Garlic Powder", BaseUnit: "tsp", PricePerBaseUnit: 0.02}, Quantity: 1, Unit: "tsp", Purpose: "flavor"},
			{FoodItem: model.FoodItem{Name: "Ginger", BaseUnit: "tsp", PricePerBaseUnit: 0.02}, Quantity: 1, Unit: "tsp", Purpose: "warmth"},
			{FoodItem: model.FoodItem{Name: "Sichuan Pepper", BaseUnit: "tsp", PricePerBaseUnit: 0.04}, Quantity: 0.5, Unit: "tsp", Purpose: "numbing heat"},
			{FoodItem: model.FoodItem{Name: "Allspice", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 0.25, Unit: "tsp", Purpose: "complex spice"},
			{FoodItem: model.FoodItem{Name: "Smoked Paprika", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 0.5, Unit: "tsp", Purpose: "smoky depth"},
			{FoodItem: model.FoodItem{Name: "Bird’s Eye Chili", BaseUnit: "whole", PricePerBaseUnit: 0.05}, Quantity: 1, Unit: "whole", Purpose: "heat"},
			{FoodItem: model.FoodItem{Name: "Egg", BaseUnit: "whole", PricePerBaseUnit: 0.20}, Quantity: 1, Unit: "whole", Purpose: "binding"},
			{FoodItem: model.FoodItem{Name: "Cornflour", BaseUnit: "tbsp", PricePerBaseUnit: 0.01}, Quantity: 3, Unit: "tbsp", Purpose: "crispy coating"},
			{FoodItem: model.FoodItem{Name: "Canola Oil", BaseUnit: "tbsp", PricePerBaseUnit: 0.01}, Quantity: 3, Unit: "tbsp", Purpose: "frying"},
			{FoodItem: model.FoodItem{Name: "Kewpie Mayo", BaseUnit: "tbsp", PricePerBaseUnit: 0.05}, Quantity: 3, Unit: "tbsp", Purpose: "creamy base"},
			{FoodItem: model.FoodItem{Name: "Sriracha Sauce", BaseUnit: "tsp", PricePerBaseUnit: 0.03}, Quantity: 2, Unit: "tsp", Purpose: "heat"},
			{FoodItem: model.FoodItem{Name: "Honey", BaseUnit: "tsp", PricePerBaseUnit: 0.04}, Quantity: 0.5, Unit: "tsp", Purpose: "sweetness"},
			{FoodItem: model.FoodItem{Name: "Lemon Juice", BaseUnit: "tsp", PricePerBaseUnit: 0.01}, Quantity: 1, Unit: "tsp", Purpose: "brightness"},
			{FoodItem: model.FoodItem{Name: "Bao Buns", BaseUnit: "whole", PricePerBaseUnit: 0.50}, Quantity: 4, Unit: "whole", Purpose: "vessel"},
			{FoodItem: model.FoodItem{Name: "Cabbage", BaseUnit: "g", PricePerBaseUnit: 0.005}, Quantity: 100, Unit: "g", Purpose: "crunch"},
			{FoodItem: model.FoodItem{Name: "Carrot", BaseUnit: "g", PricePerBaseUnit: 0.005}, Quantity: 50, Unit: "g", Purpose: "sweetness"},
			{FoodItem: model.FoodItem{Name: "Cucumber", BaseUnit: "g", PricePerBaseUnit: 0.005}, Quantity: 50, Unit: "g", Purpose: "coolness"},
			{FoodItem: model.FoodItem{Name: "Spring Onion", BaseUnit: "whole", PricePerBaseUnit: 0.10}, Quantity: 2, Unit: "whole", Purpose: "garnish"},
			{FoodItem: model.FoodItem{Name: "Roasted Peanuts", BaseUnit: "tbsp", PricePerBaseUnit: 0.02}, Quantity: 2, Unit: "tbsp", Purpose: "texture"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Slice chicken and prepare marinade mix.", Stage: "Prep", PrepTimeMinutes: 10},
			{StepNumber: 2, Instruction: "Marinate chicken in fridge 15-20 minutes.", Stage: "Prep", CookTimeMinutes: 20},
			{StepNumber: 3, Instruction: "Fry chicken in batches until golden.", Stage: "Cooking", CookTimeMinutes: 3},
			{StepNumber: 4, Instruction: "Steam bao buns per package instructions.", Stage: "Assembly", CookTimeMinutes: 4},
			{StepNumber: 5, Instruction: "Assemble buns with slaw, chicken, mayo, peanuts, onions.", Stage: "Assembly", PrepTimeMinutes: 5},
		},
	},
}

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	// --- Parse Templates ---
	// Parse all files in layouts and the specific page file
	// IMPORTANT: This simple parsing happens on every run. Better to parse once.
	// The base template must be listed FIRST if using ParseFiles with layouts
	var err error
	templates, err = template.ParseFiles(
		filepath.Join("web", "template", "layouts", "base.html"),
		filepath.Join("web", "template", "recipe_page.html"),
		filepath.Join("web", "template", "partials", "ingredients.html"),
		// Add other page templates here as you create them
	)
	if err != nil {
		log.Fatalf("Could not parse templates: %s", err)
	} else {
		log.Info("Templates parsed successfully")
	}

	http.HandleFunc("/", handleShowRecipe)
	http.HandleFunc("/update-servings", handleUpdateServings)
	fs := http.FileServer(http.Dir("./web/static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Info("Serving static files from ./web/static/ at /static/")

	port := ":8080"
	log.Infof("Starting server on http://localhost%s", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

// handleShowRecipe handles requests for the main recipe page
func handleShowRecipe(w http.ResponseWriter, r *http.Request) {
	// --- Prepare Dummy Data ---
	// In a real app, you'd fetch this from a database based on r.URL.Path maybe
	dummyRecipe := model.RecipeData{
		Title:       "Simple Example Recipe",
		Description: "A basic recipe structure to demonstrate Go templates and quantity adjustment.",
		Servings:    2, // This recipe is originally for 2 servings
		Ingredients: []model.Ingredient{
			{Name: "Flour", Quantity: 1.5, Unit: "cup"}, // Example using float
			{Name: "Large Egg", Quantity: 1, Unit: "whole"},
			{Name: "Milk", Quantity: 0.5, Unit: "cup"},
			{Name: "Salt", Quantity: 1, Unit: "pinch"}, // Unit can be descriptive
		},
		Method: []string{
			"Combine dry ingredients.",
			"Whisk in egg and milk.",
			"Cook on medium heat.",
		},
		Notes: "This is just a placeholder structure. Try adjusting the servings!",
	}

	// Data to pass to the template
	// Note the structure matches how we access it in the template (e.g., .Recipe.Title)
	data := map[string]interface{}{
		"Recipe":      dummyRecipe,
		"CurrentYear": time.Now().Year(), // Pass the year for the footer
	}

	// --- Execute Template ---
	// We want to execute the "base" template, which will in turn include
	// the correct "content" block from recipe_page.html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, "base", data) // Execute the "base" template block
	if err != nil {
		log.Errorf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Info("Served recipe page")
}

func handleUpdateServings(w http.ResponseWriter, r *http.Request) {
	// 1. Parse the form data submitted (required for POST requests)
	err := r.ParseForm()
	if err != nil {
		log.Errorf("Error parsing form: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 2. Get the new servings value from the form
	servingsStr := r.FormValue("servings") // Get value by input's "name" attribute
	newServings, err := strconv.Atoi(servingsStr)
	if err != nil || newServings <= 0 {
		log.Warnf("Invalid servings value received: %s", servingsStr)
		// Optional: could return an error message snippet to HTMX
		http.Error(w, "Invalid servings number", http.StatusBadRequest)
		return
	}

	log.Infof("Received request to update servings to: %d", newServings)

	// 3. Get the original recipe data (using the same dummy data for now)
	// In a real app, fetch this from DB or a more persistent source
	// NOTE: For simplicity, we recreate the dummy data here. A better approach
	// would be to have the base recipe defined once globally or loaded on startup.
	baseRecipe := model.RecipeData{
		// Title:       "Simple Example Recipe", // Not needed for calculation
		// Description: "...",                  // Not needed
		Servings: 2, // Base servings
		Ingredients: []model.Ingredient{
			{Name: "Flour", Quantity: 1.5, Unit: "cup"},
			{Name: "Large Egg", Quantity: 1, Unit: "whole"},
			{Name: "Milk", Quantity: 0.5, Unit: "cup"},
			{Name: "Salt", Quantity: 1, Unit: "pinch"},
		},
		// Method: []string{...}, // Not needed
		// Notes:  "...",         // Not needed
	}

	// 4. Calculate the scaling factor
	// Ensure float division
	scalingFactor := float32(newServings) / float32(baseRecipe.Servings)

	// 5. Create the new list of ingredients with adjusted quantities
	adjustedIngredients := make([]model.Ingredient, len(baseRecipe.Ingredients))
	for i, ing := range baseRecipe.Ingredients {
		adjustedIngredients[i] = model.Ingredient{
			Name: ing.Name,
			// Scale quantity, handle potential floating point inaccuracies if needed later
			Quantity: ing.Quantity * scalingFactor,
			Unit:     ing.Unit,
		}
	}

	// 6. Execute *only* the partial template with the adjusted ingredients
	// The data passed is the slice itself, matching what the partial expects
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// We need to re-wrap the output in the target div because we used outerHTML swap
	// Alternatively, target the ul directly and use innerHTML swap in the form
	fmt.Fprintln(w, `<div id="ingredients-list-container">`) // Start the replacement container
	err = templates.ExecuteTemplate(w, "ingredient-list", adjustedIngredients)
	if err != nil {
		log.Errorf("Error executing ingredient-list template: %v", err)
		// Don't write http.Error here if you already started writing the response
		return // Or handle more gracefully
	}
	fmt.Fprintln(w, `</div>`) // End the replacement container
	log.Infof("Served updated ingredients list for %d servings", newServings)
}
