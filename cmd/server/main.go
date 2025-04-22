package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type Ingredient struct {
	Name     string
	Quantity float32
	Unit     string
}

type RecipeData struct {
	Title       string
	Description string
	Servings    int
	Ingredients []Ingredient
	Method      []string
	Notes       string
}

var templates *template.Template

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
	dummyRecipe := RecipeData{
		Title:       "Simple Example Recipe",
		Description: "A basic recipe structure to demonstrate Go templates and quantity adjustment.",
		Servings:    2, // This recipe is originally for 2 servings
		Ingredients: []Ingredient{
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
	baseRecipe := RecipeData{
		// Title:       "Simple Example Recipe", // Not needed for calculation
		// Description: "...",                  // Not needed
		Servings: 2, // Base servings
		Ingredients: []Ingredient{
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
	adjustedIngredients := make([]Ingredient, len(baseRecipe.Ingredients))
	for i, ing := range baseRecipe.Ingredients {
		adjustedIngredients[i] = Ingredient{
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
