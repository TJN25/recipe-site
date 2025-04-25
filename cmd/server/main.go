package main

import (
	"errors"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/TJN25/recipe-site/internal/model"
	"github.com/TJN25/recipe-site/internal/store"
	log "github.com/sirupsen/logrus"
)

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

	// --- Setup Routes ---
	http.HandleFunc("/", handleIndexPage)
	http.HandleFunc("/recipe/", handleShowRecipe)
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
	log.Info("--- Running handleIndexPage ---") // Updated log message
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	recipes := store.GetRecipes() // Use the store function

	data := map[string]interface{}{
		"Recipes":     recipes, // Pass the recipes from the store
		"CurrentYear": time.Now().Year(),
	}

	// Retrieve the pre-parsed set for "index"
	tmplSet, ok := templateSets["index"]
	if !ok {
		log.Error("Template set 'index' not found")
		http.Error(w, "Internal Server Error: Template set 'index' missing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Execute "base.html" WITHIN the 'index' set
	err := tmplSet.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Errorf("Error executing index template set: %v", err)
		return
	}
	log.Info("Served index page")
}

// handleShowRecipe handles requests for the main recipe page
func handleShowRecipe(w http.ResponseWriter, r *http.Request) {
	log.Info("--- Running handleShowRecipe ---")

	// Extract ID (can keep this part here or move to helper)
	idStr := strings.TrimPrefix(r.URL.Path, "/recipe/")
	idStr = strings.TrimSuffix(idStr, "/")
	recipeID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Warnf("Invalid recipe ID requested: %s", idStr)
		http.NotFound(w, r)
		return
	}

	// Get recipe from the store by ID
	recipe, err := store.GetRecipeByID(recipeID)
	if err != nil {
		log.Errorf("Error getting recipe ID %d from store: %v", recipeID, err)
		if errors.Is(err, store.ErrNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	data := map[string]interface{}{
		"Recipe":      recipe, // Pass the fully assembled recipe from the store
		"CurrentYear": time.Now().Year(),
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
	err = tmplSet.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Errorf("Error executing recipe page template set for ID %d: %v", recipeID, err)
		return
	}
	log.Infof("Served recipe page for: %s (ID: %d)", recipe.Title, recipe.ID)
}

func handleUpdateServings(w http.ResponseWriter, r *http.Request) {
	log.Info("HandleUpdateServings: Received request")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusInternalServerError) // Corrected Status Code
		return
	}
	_, newServings, baseRecipe, ok := getUpdateServingsDetails(r)
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	scalingFactor := float32(newServings) / float32(baseRecipe.Servings)

	modifiedRecipe := *baseRecipe
	for idx, recipe_step := range baseRecipe.RecipeSteps {
		modifiedRecipe.RecipeSteps[idx].Ingredients = calculateAdjustedIngredients(&recipe_step, scalingFactor)
		log.Infof("HandleUpdateServings: Calculated adjustedIngredients (len %d): %+v", len(modifiedRecipe.RecipeSteps[idx].Ingredients), modifiedRecipe.RecipeSteps[idx].Ingredients)
	}

	tmplSet, found := templateSets["ingredient-list"]
	if !found {
		log.Error("Template set 'ingredient-list' not found")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	targetID := "ingredients-list-container"
	fmt.Fprintf(w, `<div id="%s">`, targetID)

	// Execute the defined name "ingredient-list"
	err := tmplSet.ExecuteTemplate(w, "ingredient-list", modifiedRecipe.RecipeSteps)
	if err != nil { /* ... */
	}
	fmt.Fprintln(w, `</div>`)

	log.Infof("Served updated ingredients list for '%s' (%d servings)", baseRecipe.Title, newServings)

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
	baseRecipe, err := store.GetRecipeByID(recipeID)
	if err != nil {
		log.Errorf("Recipe ID %d not found during update servings request", recipeID)
		return recipeID, newServings, nil, false
	}
	if baseRecipe.Servings <= 0 {
		log.Errorf("Base recipe '%s' (ID %d) has invalid original servings: %d", baseRecipe.Title, baseRecipe.ID, baseRecipe.Servings)
		return recipeID, newServings, nil, false
	}
	return recipeID, newServings, baseRecipe, true
}

func calculateAdjustedIngredients(recipeStep *model.RecipeStep, scalingFactor float32) []model.RecipeIngredient {
	// 1. Aggregate all ingredients from the base recipe's steps
	baseIngredients := recipeStep.Ingredients

	// 2. Apply scaling factor to the aggregated list
	adjustedIngredients := make([]model.RecipeIngredient, len(baseIngredients))
	for i, ing := range baseIngredients {
		// Create a copy of the FoodItem to avoid modifying the original
		foodItemCopy := ing.FoodItem
		adjustedIngredients[i] = model.RecipeIngredient{
			FoodItem:   foodItemCopy,                 // Use the copy
			Quantity:   ing.Quantity * scalingFactor, // Scale the quantity
			Unit:       ing.Unit,
			IsOptional: ing.IsOptional,
			Purpose:    ing.Purpose,
		}
	}
	return adjustedIngredients
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
	"len":            func(s []model.RecipeStep) int { return len(s) },
	"add":            func(a, b int) int { return a + b },
	"default": func(value, defaultValue string) string { // Keep if used
		if value == "" {
			return defaultValue
		}
		return value
	},
	"truncate": func(s string, length int) string { // New function
		if len(s) <= length {
			return s
		}
		// Consider rune length for Unicode safety if needed
		// return string([]rune(s)[:length]) + "..."
		return s[:length] + "..." // Simpler byte slice version
	},
}
