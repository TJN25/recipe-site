package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/TJN25/recipe-site/internal/model"
	"github.com/TJN25/recipe-site/internal/store"
	"github.com/TJN25/recipe-site/internal/units"
	log "github.com/sirupsen/logrus"
)

var templateSets map[string]*template.Template

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	// Initialize the maps
	store.InitializeCaches()
	store.InitializeUserState()

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
		filepath.Join(templateDir, "partials", "methods.html"),
		filepath.Join(templateDir, "partials", "select-recipe-steps.html"),
		filepath.Join(templateDir, "partials", "zen_mode_content.html"),
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
	ingredientSet := template.Must(template.New(filepath.Base(ingredientPartialFiles[0])). // Use "ingredients.html" as root name
												Funcs(funcMap).
												ParseFiles(ingredientPartialFiles...))
	templateSets["ingredient-list"] = ingredientSet
	log.Printf("Stored template set: ingredient-list")

	methodsPartialFiles := []string{
		filepath.Join(templateDir, "partials", "methods.html"),
	}
	log.Printf("Parsing set 'methods-list': %v", methodsPartialFiles)

	methodsSet := template.Must(template.New(filepath.Base(methodsPartialFiles[0])). // Use "ingredients.html" as root name
												Funcs(funcMap).
												ParseFiles(methodsPartialFiles...))
	templateSets["methods-list"] = methodsSet
	log.Printf("Stored template set: methods-list")

	selectRecipeStepsPartialFiles := []string{
		filepath.Join(templateDir, "partials", "select-recipe-steps.html"),
	}
	log.Printf("Parsing set 'select-recipe-steps': %v", selectRecipeStepsPartialFiles)

	selectRecipeStepsSet := template.Must(template.New(filepath.Base(selectRecipeStepsPartialFiles[0])). // Use "ingredients.html" as root name
														Funcs(funcMap).
														ParseFiles(selectRecipeStepsPartialFiles...))
	templateSets["select-recipe-steps"] = selectRecipeStepsSet
	log.Printf("Stored template set: select-recipe-steps")

	// *** CHANGE NAME IN New() HERE ***
	// Use a different root name for the set, e.g., "ingredients-root" or just the filename base
	// **********************************

	// Keep storing under the logical key "ingredient-list"

	zenPartialFiles := []string{
		filepath.Join(templateDir, "partials", "zen_mode_content.html"),
	}
	log.Printf("Parsing set 'zen-mode-content': %v", zenPartialFiles)

	// *** CHANGE NAME IN New() HERE ***
	// Use a different root name for the set, e.g., "ingredients-root" or just the filename base
	zenSet := template.Must(template.New(filepath.Base(zenPartialFiles[0])). // Use "ingredients.html" as root name
											Funcs(funcMap).
											ParseFiles(zenPartialFiles...))
	// **********************************

	// Keep storing under the logical key "ingredient-list"
	templateSets["zen-mode-content"] = zenSet
	log.Printf("Stored template set: zen-mode-content")

	// --- Setup Routes ---
	http.HandleFunc("/", handleIndexPage)
	http.HandleFunc("/recipe/", handleShowRecipe)
	http.HandleFunc("/update-servings-trigger/{id}", handleUpdateServings)
	http.HandleFunc("/render-full-ingredients/", handleRenderFullIngredients)
	http.HandleFunc("/render-recipe-steps/", handleRenderRecipeSteps)
	http.HandleFunc("/render-zen-mode-content/", handleRenderZenContent)

	// --- Serve Static Files ---
	fs := http.FileServer(http.Dir("./web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Info("Serving static files from ./web/static/ at /static/")

	// --- Start Server ---
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("INFO: No PORT environment variable detected, defaulting to %s", port)
		// Use log.Infof if using a leveled logger like logrus
	}
	listenAddr := ":" + port
	log.Infof("Starting server on http://localhost%s", listenAddr)
	err := http.ListenAndServe(listenAddr, nil)
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

	RecipeStepsAll := recipe.RecipeSteps

	recipeConfig, exists := store.GetRecipeUserConfig(recipeID)
	if exists {

		// obtain complete list of recipe steps
		additionalRecipes := recipeConfig.AdditionalRecipes
		if additionalRecipes != nil && len(additionalRecipes) > 0 {
			log.Infof("HandleShowRecipe: Additional recipes found: %v", additionalRecipes)
			for ID, additionalRecipeSteps := range additionalRecipes {
				additionalRecipe, err := store.GetRecipeByID(ID)
				if err != nil {
					log.Errorf("Cannot get recipe %d: %v", ID, err)
				}
				RecipeStepsAll = append(additionalRecipe.RecipeSteps, RecipeStepsAll...)
				recipe.RecipeSteps = append(additionalRecipe.RecipeSteps, recipe.RecipeSteps...)
				recipe.RecipeStepIds = append(additionalRecipeSteps.ActiveRecipeStepIDs, recipe.RecipeStepIds...)
			}
		}
		recipe.RecipeSteps = filterRecipeIds(recipe.RecipeSteps, recipeConfig.ActiveRecipeStepIDs)
		recipe.RecipeStepIds = recipeConfig.ActiveRecipeStepIDs
	}

	data := map[string]interface{}{
		"Recipe":                  recipe, // Pass the fully assembled recipe from the store
		"CurrentServings":         recipe.Servings,
		"DisplaySystemPreference": "use_metric_default",
		"RecipeStepsAll":          RecipeStepsAll,
		"CurrentYear":             time.Now().Year(),
	}

	// Retrieve the pre-parsed set for "recipe"
	tmplSet, ok := templateSets["recipe"]
	if !ok {
		log.Error("Template set 'recipe' not found")
		http.Error(w, "Internal Server Error: Template set 'recipe' missing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = tmplSet.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Errorf("Error executing recipe page template set for ID %d: %v", recipeID, err)
		return
	}
	log.Infof("Served recipe page for: %s (ID: %d)", recipe.Title, recipe.ID)
}

// TODO: Split this out into separate '/update-' endpoints: 'servings', 'display-units', 'active-steps', 'additional-recipe'
func handleUpdateServings(w http.ResponseWriter, r *http.Request) {
	log.Info("HandleUpdateServingsTrigger: Received request") // Update log message
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // Use correct status code
		return
	}

	// --- Parse the POST form data ---
	err := r.ParseForm()
	if err != nil {
		log.Errorf("HandleUpdateServingsTrigger: Error parsing form: %v", err)
		http.Error(w, "Bad Request: Cannot parse form", http.StatusBadRequest)
		return
	}

	// --- Read 'id' from the parsed form data ---
	recipeIDStr := r.FormValue("id") // Read the 'id' field sent by hx-vals
	log.Debugf("HandleUpdateServingsTrigger: Raw form 'id' value = '%s'", recipeIDStr)

	recipeID, err := strconv.ParseInt(recipeIDStr, 10, 64)
	if err != nil {
		// Updated error messages for clarity
		log.Errorf("HandleUpdateServingsTrigger: Error parsing recipeID from form value '%s': %v", recipeIDStr, err)
		log.Errorf("Request details: %+v", r) // Log request details for more context if needed
		http.Error(w, "Bad Request: Invalid id parameter", http.StatusBadRequest)
		return
	}
	log.Infof("HandleUpdateServingsTrigger: Parsed RecipeID: %d", recipeID)

	servingsStr := r.FormValue("servings") // Get servings (assuming name="servings" in hx-include/hx-vals)
	newServings, err := strconv.Atoi(servingsStr)
	if err != nil || newServings <= 0 {
		log.Warnf("HandleUpdateServingsTrigger: Invalid or missing 'servings' form value: '%s'. Using default/previous might be needed.", servingsStr)
		newServings = 2 // Or fetch default
	}

	unitSystem := r.FormValue("unit-system")
	log.Infof("HandleUpdateServingsTrigger: Parsed unitSystem: %s", unitSystem)

	stepIdStr := r.FormValue("step-id") // Read the 'id' field sent by hx-vals
	log.Debugf("HandleUpdateServingsTrigger: Raw form 'id' value = '%s'", stepIdStr)

	if stepIdStr != "" {
		stepID, err := strconv.ParseInt(stepIdStr, 10, 64)
		if err != nil {
			log.Errorf("HandleUpdateServingsTrigger: Error parsing stepID from form value '%s': %v", stepIdStr, err)
			log.Errorf("Request details: %+v", r) // Log request details for more context if needed
			http.Error(w, "Bad Request: Invalid id parameter", http.StatusBadRequest)
			return
		}
		log.Infof("HandleUpdateServingsTrigger: Parsed stepID: %d", stepID)

		recipe, err := store.GetRecipeByID(recipeID)
		if err != nil {
			log.Errorf("HandleUpdateServingsTrigger: Recipe not found for id '%d': %v", recipeID, err)
			log.Errorf("Request details: %+v", r)
			http.Error(w, "Bad Request: Invalid id parameter", http.StatusBadRequest)
			return
		}

		recipeConfig, exists := store.GetRecipeUserConfig(recipeID)
		if !exists {
			recipeStepsArray := recipe.RecipeStepIds
			removeInt64FromArray(&recipeStepsArray, stepID)
			recipeConfig := model.RecipeUserConfig{
				ActiveRecipeStepIDs: recipeStepsArray,
				AdditionalRecipes:   make(map[int64]model.ActiveRecipeSteps),
			}
			store.SaveRecipeConfiguration(recipeID, recipeConfig)
			log.Infof("No config for recipe %d", recipeID)
		} else {
			recipeStepsArray := recipeConfig.ActiveRecipeStepIDs
			toggleInt64InArray(&recipeStepsArray, stepID)
			recipeConfig.ActiveRecipeStepIDs = recipeStepsArray
			store.SaveRecipeConfiguration(recipeID, recipeConfig)
			log.Infof("Recipe Config: %v", recipeConfig)
		}
	}

	// A bunch of changes need to be made with other parts of the code regarding ActiveRecipeSteps
	additionalRecipeIDstr := r.FormValue("add-id") // Read the 'id' field sent by hx-vals
	if additionalRecipeIDstr != "" {
		additionalRecipeID, err := strconv.Atoi(additionalRecipeIDstr)
		if err != nil {
			log.Errorf("Cannot parse additionalRecipeID from form value '%s': %v", additionalRecipeIDstr, err)
			http.Error(w, "Invalid Recipe ID", http.StatusBadRequest)
			return
		}
		log.Infof("HandleUpdateServingsTrigger: Raw form 'add-id' value = '%d'", additionalRecipeID)

		additionalRecipe, err := store.GetRecipeByID(int64(additionalRecipeID))
		if err != nil {
			log.Errorf("Recipe: %d does not exist: %v", additionalRecipeID, err)
			http.Error(w, "Cannot find recipe", http.StatusBadRequest)
			return
		}

		additionalSteps := model.ActiveRecipeSteps{
			ActiveRecipeStepIDs: additionalRecipe.RecipeStepIds,
		}

		var recipeConfig model.RecipeUserConfig
		recipeConfig, exists := store.GetRecipeUserConfig(recipeID)
		if !exists {
			recipeConfig = model.RecipeUserConfig{
				AdditionalRecipes: make(map[int64]model.ActiveRecipeSteps),
			}
			log.Infof("Creating config for recipe %d", recipeID)
		} else {
			recipeConfig.AdditionalRecipes[int64(additionalRecipeID)] = additionalSteps
		}

		if len(recipeConfig.ActiveRecipeStepIDs) == 0 {
			recipe, err := store.GetRecipeByID(recipeID)
			if err != nil {
				log.Errorf("Cannot find recipe: %v", err)
				http.Error(w, "Cannot find recipe", http.StatusBadRequest)
				return
			}
			recipeConfig.ActiveRecipeStepIDs = append(additionalSteps.ActiveRecipeStepIDs, recipe.RecipeStepIds...)
		} else {
			recipeConfig.ActiveRecipeStepIDs = append(additionalSteps.ActiveRecipeStepIDs, recipeConfig.ActiveRecipeStepIDs...)
		}
		recipeConfig.AdditionalRecipes[int64(additionalRecipeID)] = additionalSteps
		store.SaveRecipeConfiguration(recipeID, recipeConfig)

	}

	log.Infof("HandleUpdateServingsTrigger: Target Servings: %d", newServings)

	// --- Respond with HX-Trigger ---
	w.Header().Set("HX-Trigger", fmt.Sprintf(`{"servingsUpdated": {"newServings": %d, "newSystem": "%s"}}`, newServings, unitSystem))
	w.WriteHeader(http.StatusOK)
}

func handleRenderRecipeSteps(w http.ResponseWriter, r *http.Request) {
	log.Infof("RenderRecipeSteps called")
	idStr := strings.TrimPrefix(r.URL.Path, "/render-recipe-steps/")
	idStr = strings.TrimSuffix(idStr, "/")
	recipeID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Errorf("RenderRecipeSteps: Invalid recipe ID in path '%s': %v", idStr, err)
		http.Error(w, "Invalid Recipe ID", http.StatusBadRequest)
		return
	}
	log.Infof("RenderRecipeSteps: recipeID '%d'", recipeID)

	recipe, err := store.GetRecipeByID(recipeID)
	if err != nil {
		log.Errorf("RenderRecipeSteps: Error getting recipe %d from store: %v", recipeID, err)
		if errors.Is(err, store.ErrNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	RecipeStepsAll := recipe.RecipeSteps

	recipeConfig, exists := store.GetRecipeUserConfig(recipeID)
	if exists {
		additionalRecipes := recipeConfig.AdditionalRecipes
		if additionalRecipes != nil && len(additionalRecipes) > 0 {
			log.Infof("HandleShowRecipe: Additional recipes found: %v", additionalRecipes)
			for ID := range additionalRecipes {
				additionalRecipe, err := store.GetRecipeByID(ID)
				if err != nil {
					log.Errorf("Cannot get recipe %d: %v", ID, err)
				}
				RecipeStepsAll = append(additionalRecipe.RecipeSteps, RecipeStepsAll...)
			}
		}
		getUserRecipeSteps(&recipeConfig, recipe)
		var activeRecipeStepIds []int64

		for _, recipeStep := range recipe.RecipeSteps {
			activeRecipeStepIds = append(activeRecipeStepIds, recipeStep.ID)
		}
		recipe.RecipeStepIds = activeRecipeStepIds
	}

	templateData := map[string]interface{}{
		"Recipe":         recipe,
		"RecipeStepsAll": RecipeStepsAll,
	}

	tmplSet, found := templateSets["select-recipe-steps"]
	if !found {
		log.Error("Template set 'select-recipe-steps' not found")
		http.Error(w, "Internal Server Error: Template missing", http.StatusInternalServerError)
		return
	}

	log.Infof("handleRenderRecipeSteps: TemplateData map: %+v", templateData)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmplSet.ExecuteTemplate(w, "select-recipe-steps", templateData)
	if err != nil {
		log.Errorf("Error executing select-recipe-steps template for recipe %d: %v", recipeID, err)
		return
	}

	log.Infof("RenderRecipeSteps: Sent updated steps list fragment for Recipe ID %d", recipeID)
	w.WriteHeader(http.StatusOK)
}

func handleRenderFullIngredients(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/render-full-ingredients/")
	idStr = strings.TrimSuffix(idStr, "/")
	recipeID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Errorf("RenderFullIngredients: Invalid recipe ID in path '%s': %v", idStr, err)
		http.Error(w, "Invalid Recipe ID", http.StatusBadRequest)
		return
	}

	servingsStr := r.URL.Query().Get("servings")
	unitSystem := r.URL.Query().Get("displaySystem")
	log.Infof("handleRenderFull: Parsed unitSystem: %s", unitSystem)
	targetServings, err := strconv.Atoi(servingsStr)
	if err != nil || targetServings <= 0 {
		tempRecipe, tempErr := store.GetRecipeByID(recipeID)
		if tempErr != nil {
			log.Errorf("RenderFullIngredients: Error getting recipe %d for default servings: %v", recipeID, tempErr)
			http.Error(w, "Recipe not found", http.StatusNotFound)
			return
		}
		targetServings = tempRecipe.Servings
		log.Warnf("RenderFullIngredients: Invalid/missing servings param '%s' for recipe %d. Using default %d.", servingsStr, recipeID, targetServings)
	}

	recipe, err := store.GetRecipeByID(recipeID)
	if err != nil {
		log.Errorf("RenderFullIngredients: Error getting recipe %d from store: %v", recipeID, err)
		if errors.Is(err, store.ErrNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	recipeConfig, exists := store.GetRecipeUserConfig(recipeID)
	if exists {
		getUserRecipeSteps(&recipeConfig, recipe)
	}
	log.Infof("RenderFullIngredients: recipe.RecipeStepIds: %v", recipe.RecipeStepIds)

	templateData := map[string]interface{}{
		"Recipe":                  recipe,
		"CurrentServings":         targetServings,
		"DisplaySystemPreference": unitSystem,
	}

	tmplSet, found := templateSets["ingredient-list"]
	if !found {
		log.Error("Template set 'ingredient-list' not found")
		http.Error(w, "Internal Server Error: Template missing", http.StatusInternalServerError)
		return
	}

	log.Infof("handleRenderZenContent: TemplateData map: %+v", templateData)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmplSet.ExecuteTemplate(w, "ingredient-list", templateData)
	if err != nil {
		log.Errorf("Error executing ingredient-list template for recipe %d: %v", recipeID, err)
		return
	}

	log.Infof("RenderFullIngredients: Sent updated ingredient list fragment for Recipe ID %d with Target Servings %d, and Unit System %s.", recipeID, targetServings, unitSystem)
}

func handleRenderZenContent(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/render-zen-mode-content/")
	idStr = strings.TrimSuffix(idStr, "/")
	recipeID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Errorf("RenderZenContent: Invalid recipe ID in path '%s': %v", idStr, err)
		http.Error(w, "Invalid Recipe ID", http.StatusBadRequest)
		return
	}

	servingsStr := r.URL.Query().Get("servings")
	unitSystem := r.URL.Query().Get("displaySystem")
	log.Infof("handleRenderZen: Parsed unitSystem: %s", unitSystem)
	targetServings, err := strconv.Atoi(servingsStr)
	if err != nil || targetServings <= 0 {
		tempRecipe, tempErr := store.GetRecipeByID(recipeID)
		if tempErr != nil {
			log.Errorf("RenderZenContent: Error getting recipe %d for default servings: %v", recipeID, tempErr)
			http.Error(w, "Recipe not found", http.StatusNotFound)
			return
		}
		targetServings = tempRecipe.Servings
		log.Warnf("RenderZenContent: Invalid/missing servings param '%s' for recipe %d. Using default %d.", servingsStr, recipeID, targetServings)
	}

	recipe, err := store.GetRecipeByID(recipeID)
	if err != nil {
		log.Errorf("RenderZenContent: Error getting recipe %d from store: %v", recipeID, err)
		if errors.Is(err, store.ErrNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	recipeConfig, exists := store.GetRecipeUserConfig(recipeID)
	if exists {
		getUserRecipeSteps(&recipeConfig, recipe)
	}

	templateData := map[string]interface{}{
		"Recipe":                  recipe,
		"CurrentServings":         targetServings,
		"DisplaySystemPreference": unitSystem,
	}

	tmplSet, found := templateSets["zen-mode-content"]
	if !found {
		log.Error("Template set 'zen-mode-content' not found")
		http.Error(w, "Internal Server Error: Template missing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmplSet.ExecuteTemplate(w, "zen-mode-content", templateData)
	if err != nil {
		log.Errorf("Error executing zen-mode-content template for recipe %d: %v", recipeID, err)
		return
	}

	log.Infof("RenderZenContent: Sent updated ingredient list fragment for Recipe ID %d with Target Servings %d.", recipeID, targetServings)
}

func getUserRecipeSteps(recipeConfig *model.RecipeUserConfig, recipe *model.Recipe) {
	log.Infof("getUserRecipeSteps: User Config found: %v", recipeConfig)

	additionalRecipes := recipeConfig.AdditionalRecipes
	if additionalRecipes != nil && len(additionalRecipes) > 0 {
		log.Infof("RenderFullIngredients: Additional recipes found: %v", additionalRecipes)
		for ID := range additionalRecipes {
			additionalRecipe, err := store.GetRecipeByID(ID)
			if err != nil {
				log.Errorf("Cannot get recipe %d: %v", ID, err)
			}
			recipe.RecipeStepIds = recipeConfig.ActiveRecipeStepIDs
			recipe.RecipeSteps = append(additionalRecipe.RecipeSteps, recipe.RecipeSteps...)
		}
	}
	recipe.RecipeSteps = filterRecipeIds(recipe.RecipeSteps, recipeConfig.ActiveRecipeStepIDs)
}

func addInt64ToArray(array *[]int64, value int64) {
	log.Infof("Input array: %v, value: %d", *array, value)
	for _, v := range *array {
		if v == value {
			log.Infof("Output array: %v, value: %d", *array, value)
			return
		}
	}
	*array = append(*array, value)
	log.Infof("Output array: %v, value: %d", *array, value)
}

func removeInt64FromArray(array *[]int64, value int64) {
	log.Infof("Input array: %v, value: %d", *array, value)
	result := []int64{}
	for _, v := range *array {
		if v != value {
			result = append(result, v)
		}
	}
	*array = result
	log.Infof("Output array: %v, value: %d", *array, value)

}

func toggleInt64InArray(array *[]int64, value int64) {
	log.Infof("Input array: %v, value: %d", *array, value)
	for _, v := range *array {
		if v == value {
			removeInt64FromArray(array, value)
			return
		}
	}
	addInt64ToArray(array, value)
	log.Infof("Output array: %v, value: %d", *array, value)
}

func filterRecipeIds(objects []model.RecipeStep, ids []int64) []model.RecipeStep {
	idMap := make(map[int64]bool)
	for _, id := range ids {
		idMap[id] = true
	}

	filtered := []model.RecipeStep{}
	for _, obj := range objects {
		if _, ok := idMap[obj.ID]; ok {
			filtered = append(filtered, obj)
		}
	}
	return filtered
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
		foodItemCopy := ing.FoodItemID
		adjustedIngredients[i] = model.RecipeIngredient{
			FoodItemID:    foodItemCopy,                 // Use the copy
			Quantity:      ing.Quantity * scalingFactor, // Scale the quantity
			SpecifiedUnit: ing.SpecifiedUnit,
			IsOptional:    ing.IsOptional,
			Purpose:       ing.Purpose,
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

func getFoodItem(id int64) model.FoodItem {
	item, found := store.AllFoodItemsCache[id]
	if !found {
		log.Warnf("FoodItem with ID %d not found in cache", id)
		// Return an empty struct to avoid template errors, log the issue.
		return model.FoodItem{}
	}
	return item
}

func getFormDetails(item model.FoodItem, formName string) model.FoodItemFormDetails {
	// Handle potential nil map if FoodItem wasn't found or has no forms
	if item.Forms == nil {
		log.Warnf("FoodItem ID %d has nil Forms map when looking for form '%s'", item.ID, formName)
		return model.FoodItemFormDetails{}
	}
	details, found := item.Forms[formName]
	if !found {
		// Attempt to use DefaultFormName if provided formName is not found
		if item.DefaultFormName != "" && formName != item.DefaultFormName {
			log.Warnf("Form '%s' not found for FoodItem ID %d. Trying default '%s'", formName, item.ID, item.DefaultFormName)
			details, found = item.Forms[item.DefaultFormName]
			if !found {
				log.Errorf("Default form '%s' ALSO not found for FoodItem ID %d", item.DefaultFormName, item.ID)
				return model.FoodItemFormDetails{} // Return empty if default also fails
			}
		} else {
			// If the requested name WAS the default, or default is empty, and it wasn't found.
			log.Errorf("Form '%s' not found for FoodItem ID %d and no default fallback available or default also missing.", formName, item.ID)
			return model.FoodItemFormDetails{} // Return empty
		}
	}
	return details
}

func unitSpace(unit string) string {
	// List of units that should NOT have a preceding space
	noSpaceUnits := map[string]bool{
		"g":  true,
		"ml": true,
		// Add other symbols like °C, °F if needed
	}
	if _, found := noSpaceUnits[unit]; found || unit == "" {
		return "" // No space for these units or empty unit
	}
	return " " // Add space for others (cup, tsp, clove, unit, etc.)
}

func marshal(v interface{}) (template.JS, error) {
	a, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.JS(a), nil // Return as template.JS to prevent over-escaping
}

func containsInt64(slice []int64, value int64) bool {
	log.Infof("ContainsInt64: %v, %d", slice, value)
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
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
	"getFoodItem":      getFoodItem,
	"getFormDetails":   getFormDetails,
	"unitSpace":        unitSpace,
	"formatIngredient": units.FormatIngredientForDisplay,
	"marshal":          marshal,
	"containsInt64":    containsInt64,
}
