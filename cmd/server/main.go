package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/TJN25/recipe-site/internal/model"
	log "github.com/sirupsen/logrus"
)

// some relational linking of all the recipes by a score (shared ingredients, shared tags, shared equipment)
// Matrix of Recipes, by Recipes with score (but maybe more efficient?)

var templateSets map[string]*template.Template

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	// Initialize the map
	templateSets = make(map[string]*template.Template)

	// --- PARSE TEMPLATES AT STARTUP (Separate Sets Pattern) ---
	log.Println("Parsing templates...")
	templateDir := "web/template"
	baseFile := filepath.Join(templateDir, "layouts", "base.html")

	// 1. Parse template set for the Index Page
	indexPageFiles := []string{
		baseFile,
		filepath.Join(templateDir, "index.html"),
	}
	log.Printf("Parsing set 'index': %v", indexPageFiles)
	// Parse using the base name of baseFile ("base.html") as the root name
	indexSet := template.Must(template.New(filepath.Base(baseFile)).
		Funcs(funcMap). // Include funcs if base or index needs them
		ParseFiles(indexPageFiles...))
	templateSets["index"] = indexSet
	log.Printf("Stored template set: index")

	// 2. Parse template set for the Recipe Page (keep for later)
	recipePageFiles := []string{
		baseFile, // web/template/layouts/base.html
		filepath.Join(templateDir, "recipe_page.html"),
		// We can omit ingredients.html for this minimal test if recipe_page doesn't {{template}} it
		filepath.Join(templateDir, "partials", "ingredients.html"),
	}
	log.Printf("Parsing set 'recipe': %v", recipePageFiles)
	recipeSet := template.Must(template.New(filepath.Base(baseFile)). // Rooted at base.html
										Funcs(funcMap).
										ParseFiles(recipePageFiles...))
	templateSets["recipe"] = recipeSet
	log.Printf("Stored template set: recipe")

	// 3. Parse standalone partial for HTMX swap (ingredients.html) (keep for later)
	ingredientPartialFiles := []string{
		filepath.Join(templateDir, "partials", "ingredients.html"),
	}
	log.Printf("Parsing set 'ingredient-list': %v", ingredientPartialFiles)

	// *** CHANGE NAME IN New() HERE ***
	// Use a different root name for the set, e.g., "ingredients-root" or just the filename base
	ingredientSet := template.Must(template.New(filepath.Base(ingredientPartialFiles[0])). // Use "ingredients.html" as root name
												Funcs(funcMap).
												ParseFiles(ingredientPartialFiles...))
	// **********************************

	// Keep storing under the logical key "ingredient-list"
	templateSets["ingredient-list"] = ingredientSet
	log.Printf("Stored template set: ingredient-list")
	// --- END TEMPLATE PARSING ---

	// --- Setup Routes ---
	http.HandleFunc("/", handleIndexPage)         // Focus on this route
	http.HandleFunc("/recipe/", handleShowRecipe) // Keep other routes defined
	http.HandleFunc("/update-servings", handleUpdateServings)

	// --- Serve Static Files ---
	fs := http.FileServer(http.Dir("./web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Info("Serving static files from ./web/static/ at /static/")

	// --- Start Server ---
	port := ":8080"
	log.Infof("Starting server on http://localhost%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func handleIndexPage(w http.ResponseWriter, r *http.Request) {
	log.Info("--- Running handleIndexPage (Simplified) ---")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Minimal data, or nil if index.html and base.html don't need it
	data := map[string]interface{}{
		"CurrentYear": time.Now().Year(), // Keep if base.html uses it
	}

	// Retrieve the pre-parsed set for "index"
	tmplSet, ok := templateSets["index"]
	if !ok {
		log.Error("Template set 'index' not found")
		http.Error(w, "Internal Server Error: Template set 'index' missing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// *** EXECUTE "base.html" within the 'index' set ***
	// This matches the successful pattern from your minimal example
	err := tmplSet.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Errorf("Error executing simplified index template set: %v", err)
		// Don't write http.Error if potentially already written headers/body
		return
	}
	log.Info("Served simplified index page")
}

// handleShowRecipe handles requests for the main recipe page
func handleShowRecipe(w http.ResponseWriter, r *http.Request) {
	log.Info("--- Running handleShowRecipe (Simplified Test) ---")
	// ... (keep logic to extract ID and find recipe) ...
	recipeID, recipeCopy, found := getRecipeDetailsFromRequest(r)
	if !found {
		http.NotFound(w, r)
		return
	}

	// Data includes the Recipe struct needed by the simplified template
	data := map[string]interface{}{
		"Recipe":      recipeCopy,
		"CurrentYear": time.Now().Year(), // If base.html uses it
	}

	// Retrieve the pre-parsed set for "recipe"
	tmplSet, ok := templateSets["recipe"]
	if !ok {
		log.Error("Template set 'recipe' not found")
		http.Error(w, "Internal Server Error: Template set 'recipe' missing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// *** EXECUTE "base.html" within the 'recipe' set ***
	err := tmplSet.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Errorf("Error executing simplified recipe page template set for ID %d: %v", recipeID, err)
		return
	}
	log.Infof("Served simplified recipe page for: %s (ID: %d)", recipeCopy.Title, recipeCopy.ID)
}

func handleUpdateServings(w http.ResponseWriter, r *http.Request) {
	log.Info("HandleUpdateServings: Received request")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusInternalServerError) // Corrected Status Code
		return
	}
	recipeID, newServings, baseRecipe, ok := getUpdateServingsDetails(r)
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	scalingFactor := float32(newServings) / float32(baseRecipe.Servings)
	adjustedIngredients := calculateAdjustedIngredients(baseRecipe, scalingFactor)

	// **** ADD THIS LOG ****
	log.Infof("HandleUpdateServings: Calculated adjustedIngredients (len %d): %+v", len(adjustedIngredients), adjustedIngredients)
	// *********************

	tmplSet, found := templateSets["ingredient-list"]
	if !found {
		log.Error("Template set 'ingredient-list' not found")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Wrap the output because the target expects the container div
	fmt.Fprintln(w, `<div id="ingredients-list-container">`)
	// Execute the "ingredients.html" template directly (it's the only one in its set)
	// We use the partial's *filename* here as the template name within its set
	err := tmplSet.ExecuteTemplate(w, "ingredient-list", adjustedIngredients)
	if err != nil {
		log.Errorf("Error executing ingredient-list template set for recipe ID %d: %v", recipeID, err)
		// Avoid sending another error if headers/body already partially sent
		fmt.Fprintln(w, `<!-- Error executing template -->`) // Add comment for debugging
	}
	fmt.Fprintln(w, `</div>`)

	log.Infof("Served updated ingredients list for '%s' (%d servings)", baseRecipe.Title, newServings)

}

func getRecipeDetailsFromRequest(r *http.Request) (id int64, recipe model.Recipe, found bool) {
	idStr := strings.TrimPrefix(r.URL.Path, "/recipe/")
	idStr = strings.TrimSuffix(idStr, "/")
	recipeID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Warnf("Invalid recipe ID requested: %s", idStr)
		return 0, model.Recipe{}, false
	}
	recipePtr, found := findRecipeByID(recipeID)
	if !found {
		log.Warnf("Recipe ID not found: %d", recipeID)
		return recipeID, model.Recipe{}, false
	}
	recipeCopy := *recipePtr
	recipeCopy.CalculateTotals()
	return recipeID, recipeCopy, true
}

func getUpdateServingsDetails(r *http.Request) (id int64, servings int, recipe *model.Recipe, ok bool) {
	err := r.ParseForm()
	if err != nil {
		log.Errorf("Error parsing form: %v", err)
		return 0, 0, nil, false
	}
	recipeIDStr := r.FormValue("recipe_id")
	recipeID, err := strconv.ParseInt(recipeIDStr, 10, 64)
	if err != nil {
		log.Warnf("Invalid recipe_id value received in form: '%s'", recipeIDStr)
		return 0, 0, nil, false
	}
	servingsStr := r.FormValue("servings")
	newServings, err := strconv.Atoi(servingsStr)
	if err != nil || newServings <= 0 {
		log.Warnf("Invalid servings value received: '%s'", servingsStr)
		return recipeID, 0, nil, false
	}
	baseRecipe, found := findRecipeByID(recipeID)
	if !found {
		log.Errorf("Recipe ID %d not found during update servings request", recipeID)
		return recipeID, newServings, nil, false
	}
	if baseRecipe.Servings <= 0 {
		log.Errorf("Base recipe '%s' (ID %d) has invalid original servings: %d", baseRecipe.Title, baseRecipe.ID, baseRecipe.Servings)
		return recipeID, newServings, nil, false
	}
	return recipeID, newServings, baseRecipe, true
}

func calculateAdjustedIngredients(baseRecipe *model.Recipe, scalingFactor float32) []model.RecipeIngredient {
	adjustedIngredients := make([]model.RecipeIngredient, len(baseRecipe.Ingredients))
	for i, ing := range baseRecipe.Ingredients {
		foodItemCopy := ing.FoodItem
		adjustedIngredients[i] = model.RecipeIngredient{
			FoodItem:   foodItemCopy,
			Quantity:   ing.Quantity * scalingFactor,
			Unit:       ing.Unit,
			IsOptional: ing.IsOptional,
			Purpose:    ing.Purpose,
		}
	}
	return adjustedIngredients
}

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

func formatQuantity(q float32) string {
	// Check if the number is effectively an integer (within a small tolerance)
	if math.Abs(float64(q)-math.Round(float64(q))) < 0.001 {
		return fmt.Sprintf("%.0f", q)
	}

	// strconv 'g' might still produce more than 2 decimal places in some cases (e.g., scientific notation for very small/large)
	// Let's explicitly format to 2dp and trim if needed for typical recipe quantities.
	sFixed := fmt.Sprintf("%.2f", q)
	// Remove trailing ".00"
	if strings.HasSuffix(sFixed, ".00") {
		return sFixed[:len(sFixed)-3]
	}
	// Remove trailing "0" for things like "2.50" -> "2.5"
	if strings.HasSuffix(sFixed, "0") {
		return sFixed[:len(sFixed)-1]
	}
	// Otherwise return the 2dp fixed version
	return sFixed
}

// Create a FuncMap to register the function
var funcMap = template.FuncMap{
	"formatQuantity": formatQuantity,
}

func findRecipeByID(id int64) (*model.Recipe, bool) {
	for i := range dummyRecipes {
		if dummyRecipes[i].ID == id {
			// Return a pointer to the recipe in the slice
			return &dummyRecipes[i], true
		}
	}
	return nil, false // Not found
}
