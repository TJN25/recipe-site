package store

import (
	"errors" // For returning errors
	"fmt"    // For potential error messages
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/TJN25/recipe-site/internal/data"  // Import dummy data package (adjust path if needed)
	"github.com/TJN25/recipe-site/internal/model" // Import model package (adjust path if needed)
)

var AllFoodItemsCache map[int64]model.FoodItem
var AllFoodItemNamesCache map[string]int64
var NormalizedFormNameLookup map[string]string
var AllDummyStepsCache map[int64]data.DummyRecipeStep

// setting up the temp storage of 'user' preferences

var stateMutex sync.RWMutex

var userGlobalSettings model.GlobalUserSettings

func InitializeUserState() {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	userGlobalSettings = model.GlobalUserSettings{
		DefaultServings:   2,
		DisplayUnitSystem: "use_original",
		ShowOptions:       false,
	}
	userRecipeConfigs = make(map[int64]model.RecipeUserConfig)
	log.Info("In-memory user state initialized.")
}

var userRecipeConfigs = make(map[int64]model.RecipeUserConfig)

func GetGlobalPreferences() model.GlobalUserSettings {
	stateMutex.RLock()
	defer stateMutex.RUnlock()
	return userGlobalSettings
}

func UpdateGlobalDisplaySystem(newSystem string) {
	stateMutex.RLock()
	defer stateMutex.RUnlock()
	userGlobalSettings.DisplayUnitSystem = newSystem
	log.Infof("Global display system updated to: %s", newSystem)
}

func GetRecipeUserConfig(recipeID int64) (model.RecipeUserConfig, bool) {
	stateMutex.RLock()
	defer stateMutex.RUnlock()
	config, exists := userRecipeConfigs[recipeID]
	return config, exists
}

func SaveRecipeConfiguration(recipeID int64, config model.RecipeUserConfig) {
	stateMutex.Lock()
	defer stateMutex.Unlock()
	userRecipeConfigs[recipeID] = config
	log.Debugf("Saved configuration for recipe %d: %+v", recipeID, config)
}

func GetCurrentServings(recipeID int64) (int, error) {
	config, exists := userRecipeConfigs[recipeID]
	recipe, err := GetRecipeByID(recipeID)
	if err != nil {
		return -1, err
	}
	currentServings := recipe.Servings
	if exists {
		if config.Servings > 0 {
			currentServings = config.Servings
		}
	}

	return currentServings, nil
}

func UpdateServings(recipeID int64, newServings int) {
	recipeConfig, exists := GetRecipeUserConfig(recipeID)
	if !exists {
		recipeConfig := model.RecipeUserConfig{
			Servings:          newServings,
			AdditionalRecipes: make(map[int64]model.ActiveRecipeSteps),
		}
		SaveRecipeConfiguration(recipeID, recipeConfig)
		log.Infof("No config for recipe %d", recipeID)
	} else {
		recipeConfig.Servings = newServings
		SaveRecipeConfiguration(recipeID, recipeConfig)
		log.Infof("Recipe Config: %v", recipeConfig)
	}
	log.Infof("UpdateServings: Target Servings: %d", newServings)
}

func GetBaseServings(recipeID int64) (int, error) {
	recipe, err := GetRecipeByID(recipeID)
	if err != nil {
		return -1, err
	}
	return recipe.Servings, nil
}

func GetUnitSystem() string {
	return "use_metric_default"
}

func UpdateRecipeStep(recipeID int64, recipeStepsArray []int64, stepID int64) {
	recipeConfig, exists := GetRecipeUserConfig(recipeID)
	if !exists {
		removeInt64FromArray(&recipeStepsArray, stepID)
		recipeConfig := model.RecipeUserConfig{
			ActiveRecipeStepIDs: recipeStepsArray,
			AdditionalRecipes:   make(map[int64]model.ActiveRecipeSteps),
		}
		SaveRecipeConfiguration(recipeID, recipeConfig)
		log.Infof("No config for recipe %d", recipeID)
	} else {
		recipeStepsArray := recipeConfig.ActiveRecipeStepIDs
		toggleInt64InArray(&recipeStepsArray, stepID)
		recipeConfig.ActiveRecipeStepIDs = recipeStepsArray
		SaveRecipeConfiguration(recipeID, recipeConfig)
		log.Infof("Recipe Config: %v", recipeConfig)
	}

	log.Infof("UpdateRecipeStep: stepID: %d", stepID)
}

func InitializeCaches() {
	log.Info("Initializing store caches...")
	AllFoodItemsCache = make(map[int64]model.FoodItem)
	AllFoodItemNamesCache = make(map[string]int64)
	// Ensure data.DummyFoodItems uses the NEW model.FoodItem structure
	for _, item := range data.DummyFoodItems {
		AllFoodItemsCache[item.ID] = item
		normalizedKey := NormalizeName(item.Name)
		AllFoodItemNamesCache[normalizedKey] = item.ID
		log.Debugf("Added to name cache: Key='%s', ID=%d", normalizedKey, item.ID)
	}
	log.Infof("Cached %d FoodItems", len(AllFoodItemsCache))
	log.Infof("Cached %d FoodItemsNames", len(AllFoodItemNamesCache))

	NormalizedFormNameLookup = make(map[string]string)
	for foodItemID, foodItem := range AllFoodItemsCache {
		if foodItem.Forms == nil {
			log.Warnf("FoodItem ID %d ('%s') has nil Forms map during form cache init.", foodItemID, foodItem.Name)
			continue // Skip items with no forms map
		}
		for originalFormKey := range foodItem.Forms {
			normalizedKey := NormalizeName(originalFormKey) // Normalize the key
			if normalizedKey == "" {
				log.Warnf("FoodItem ID %d ('%s') form key '%s' resulted in empty normalized key. Skipping.", foodItemID, foodItem.Name, originalFormKey)
				continue
			}

			if existingOriginalKey, exists := NormalizedFormNameLookup[normalizedKey]; exists {
				log.Errorf("Normalized Form Name Collision: Normalized key '%s' produced by both '%s' and '%s' (FoodItem ID %d). Keeping first ('%s'). Check FoodItem definitions.",
					normalizedKey, existingOriginalKey, originalFormKey, foodItemID, existingOriginalKey)
				continue
			} else {
				NormalizedFormNameLookup[normalizedKey] = originalFormKey
				log.Debugf("Added to form name cache: Key='%s', Value='%s' (from FoodItem %d)", normalizedKey, originalFormKey, foodItemID)
			}
		}
	}
	log.Infof("Cached %d unique normalized form names", len(NormalizedFormNameLookup))

	// Init Dummy Steps
	AllDummyStepsCache = make(map[int64]data.DummyRecipeStep)
	// Ensure data.DummyRecipeSteps is the slice of raw dummy step data
	for _, step := range data.DummyRecipeSteps {
		// Add validation if needed: Check for duplicate step IDs?
		AllDummyStepsCache[step.ID] = step
	}
	log.Infof("Cached %d DummyRecipeSteps", len(AllDummyStepsCache))
	if log.GetLevel() >= log.DebugLevel {
		count := 0
		log.Debug("Sample DummyRecipeStep Cache Entries:")
		for id, step := range AllDummyStepsCache {
			log.Debugf("  - ID: %d, Title: '%s', RecipeID: %d", id, step.Title, step.RecipeID) // Added RecipeID for context
			count++
			if count >= 5 { // Limit logging to first 5 entries
				log.Debug("  - ... (logging first 5 entries only)")
				break
			}
		}
		if count == 0 {
			log.Debug("  - Cache is empty!")
		}
	}
}

func findFoodItem(name string) (model.FoodItem, bool) {
	log.Debugf("findFoodItem: %s", name)
	normalizedSearchName := NormalizeName(name)
	log.Debugf("findFoodItem: normalized %s", normalizedSearchName)
	for _, item := range data.DummyFoodItems {
		if NormalizeName(item.Name) == normalizedSearchName {
			return item, true
		}
	}
	log.Warnf("findFoodItem: did not find '%s'", normalizedSearchName)
	return model.FoodItem{}, false
}

func findEquipment(id int64) (model.Equipment, bool) {
	for _, item := range data.DummyEquipment {
		if item.ID == id {
			return item, true
		}
	}
	return model.Equipment{}, false
}

func findTag(id int64) (model.Tag, bool) {
	for _, item := range data.DummyTags {
		if item.ID == id {
			return item, true
		}
	}
	return model.Tag{}, false
}

var ErrNotFound = errors.New("resource not found")

func GetRecipeByID(id int64) (*model.Recipe, error) {
	var coreRecipeData *struct {
		ID            int64
		RecipeStepIds []int64
		Title         string
		Description   string
		Servings      int
		Notes         string
		ImagePath     string
	}

	foundCore := false
	for i := range data.DummyRecipes {
		if data.DummyRecipes[i].ID == id {
			coreRecipeData = &struct {
				ID            int64
				RecipeStepIds []int64
				Title         string
				Description   string
				Servings      int
				Notes         string
				ImagePath     string
			}{
				ID:            data.DummyRecipes[i].ID,
				Title:         data.DummyRecipes[i].Title,
				Description:   data.DummyRecipes[i].Description,
				Servings:      data.DummyRecipes[i].Servings,
				Notes:         data.DummyRecipes[i].Notes,
				ImagePath:     data.DummyRecipes[i].ImagePath,
				RecipeStepIds: data.DummyRecipes[i].RecipeStepIds,
			}
			foundCore = true
			break
		}
	}
	if !foundCore || coreRecipeData == nil {
		return nil, fmt.Errorf("recipe with ID %d not found or missing RecipeStepIds in DummyRecipes: %w", id, ErrNotFound)
	}

	recipe := &model.Recipe{
		ID:            coreRecipeData.ID,
		RecipeStepIds: coreRecipeData.RecipeStepIds,
		Title:         coreRecipeData.Title,
		Description:   coreRecipeData.Description,
		Servings:      coreRecipeData.Servings,
		Notes:         coreRecipeData.Notes,
		ImagePath:     coreRecipeData.ImagePath,
		RecipeSteps:   make([]model.RecipeStep, 0, len(coreRecipeData.RecipeStepIds)), // Initialize slices
		Tags:          []model.Tag{},
	}

	for _, stepID := range coreRecipeData.RecipeStepIds {
		dummyStep, found := AllDummyStepsCache[stepID]
		if !found {
			log.Warnf("Store: RecipeStep ID %d referenced by Recipe ID %d not found in cache. Skipping step.", stepID, recipe.ID)
			continue
		}

		recipeStep := model.RecipeStep{
			ID:           dummyStep.ID,
			StepOrder:    dummyStep.StepOrder,
			Title:        dummyStep.Title,
			Description:  dummyStep.Description,
			Notes:        dummyStep.Notes,
			Ingredients:  []model.RecipeIngredient{},
			MethodSteps:  dummyStep.MethodSteps,
			Equipment:    []model.Equipment{},
			BaseServings: 2,
		}

		for _, ingRef := range dummyStep.Ingredients {
			// 1. Lookup FoodItem
			foodItem, found := findFoodItem(ingRef.FoodItemName)
			if !found {
				log.Warnf("Store: FoodItem Name %s not found referenced in RecipeStep ID %d. Skipping ingredient.", ingRef.FoodItemName, dummyStep.ID)
				continue
			}

			// 2. Basic Validation (Optional but good)
			if foodItem.DefaultFormName == "" {
				log.Warnf("Store: FoodItem ID %d ('%s') has empty DefaultFormName. Data might be incomplete.", foodItem.ID, foodItem.Name)
				// We can still proceed, DefaultFormName is mostly for substitutions now
			}

			targetFormName := ingRef.FormName // Start with the name from the reference
			if targetFormName == "" {
				log.Tracef("Store Init: No FormName specified for '%s' in Step %d and FoodItem %d ('%s').",
					ingRef.FoodItemName, dummyStep.ID, foodItem.ID, foodItem.Name)
				// If no form was specified in the reference, use the FoodItem's default
				targetFormName = foodItem.DefaultFormName
				if targetFormName == "" {
					// If the FoodItem itself ALSO lacks a default, we have a problem
					log.Errorf("Store Init: No FormName specified for '%s' in Step %d and FoodItem %d ('%s') also has no DefaultFormName. Skipping.",
						ingRef.FoodItemName, dummyStep.ID, foodItem.ID, foodItem.Name)
					continue // Skip this ingredient
				}
				log.Tracef("Store Init: No FormName for '%s' in Step %d. Using default '%s'.", ingRef.FoodItemName, dummyStep.ID, targetFormName)
			}

			// Now, validate that the determined targetFormName exists in the Forms map
			normalisedTargetFormName := NormalizeName(targetFormName)
			lookupFormName, ok := NormalizedFormNameLookup[normalisedTargetFormName]
			if !ok {
				log.Errorf("Target FormName '%s' for '%s' in Step %d not found in FormNameLookup. Trying original name.", targetFormName, ingRef.FoodItemName, dummyStep.StepOrder)
				lookupFormName = targetFormName
			}
			_, formExists := foodItem.Forms[lookupFormName]
			if !formExists {
				// The intended form (either specified or default) doesn't exist in the definition
				log.Errorf("Store Init: Target FormName '%s' for '%s' in Step %d not found in FoodItem %d Forms map. Check data definitions. Skipping ingredient.",
					targetFormName, ingRef.FoodItemName, dummyStep.ID, foodItem.ID)
				continue // Skip this ingredient - cannot proceed without valid form details
			}
			// 3. Create the RecipeIngredient - Directly recording source data
			newIngredient := model.RecipeIngredient{
				FoodItemID:    foodItem.ID,
				FoodItemName:  ingRef.FoodItemName,
				FormName:      targetFormName,
				Quantity:      ingRef.Quantity, // Store quantity as given
				SpecifiedUnit: ingRef.Unit,     // <<< STORE THE UNIT FROM THE SOURCE
				IsOptional:    ingRef.IsOptional,
				Purpose:       ingRef.Purpose,
			}

			// 4. Append
			recipeStep.Ingredients = append(recipeStep.Ingredients, newIngredient)
		}

		// for _, equipID := range dummyStep.EquipmentIDs {
		// 	equipment, found := findEquipment(equipID)
		// 	if !found {
		// 		log.Warnf("Warning: Equipment ID %d not found for RecipeStep ID %d\n", equipID, dummyStep.ID)
		// 		continue // Skip if not found
		// 	}
		// 	recipeStep.Equipment = append(recipeStep.Equipment, equipment)
		// }

		// Add the fully assembled step to the recipe
		recipe.RecipeSteps = append(recipe.RecipeSteps, recipeStep)

	}
	for _, link := range data.DummyRecipeTags {
		if link.RecipeID == recipe.ID {
			tag, found := findTag(link.TagID)
			if !found {
				fmt.Printf("Warning: Tag ID %d not found for Recipe ID %d\n", link.TagID, recipe.ID)
				continue // Skip if not found
			}
			recipe.Tags = append(recipe.Tags, tag)
		}
	}

	// 4. Calculate derived totals
	// recipe.CalculateTotals() // Call the method on the assembled recipe

	return recipe, nil

}

func GetRecipes() []model.Recipe {
	recipes := make([]model.Recipe, 0, len(data.DummyRecipes))
	for _, coreData := range data.DummyRecipes {
		recipes = append(recipes, model.Recipe{
			ID:          coreData.ID,
			Title:       coreData.Title,
			Description: coreData.Description,
		})
	}
	return recipes
}

func NormalizeName(name string) string {
	// 1. Trim leading/trailing whitespace
	processedName := strings.TrimSpace(name)
	// 2. Convert to lowercase
	processedName = strings.ToLower(processedName)
	// 3. Remove hyphens
	processedName = strings.ReplaceAll(processedName, "-", "")
	processedName = strings.ReplaceAll(processedName, " ", "")
	// 4. Remove trailing 's' (basic plural handling)
	// Be careful: this might incorrectly change "pasta" to "pata", "tapas" to "tapa" etc.
	// Consider only removing 's' if preceded by certain letters, or use a more robust stemmer later.
	// For now, basic removal:
	if len(processedName) > 1 && strings.HasSuffix(processedName, "s") {
		// Only remove if it's not just "s"
		// Might want to add exceptions for words like "pasta", "tapas", "molasses" if they occur
		exceptions := map[string]bool{"pasta": true, "tapas": true, "molasses": true}
		if !exceptions[processedName] {
			processedName = processedName[:len(processedName)-1]
		}
	}
	// 5. Optional: Condense multiple spaces?
	// processedName = regexp.MustCompile(`\s+`).ReplaceAllString(processedName, " ")

	return processedName
}

func addInt64ToArray(array *[]int64, value int64) {
	log.Infof("Add: Input array: %v, value: %d", *array, value)
	for _, v := range *array {
		if v == value {
			log.Infof("Output array: %v, value: %d", *array, value)
			return
		}
	}
	*array = append(*array, value)
	log.Infof("Add: Output array: %v, value: %d", *array, value)
}

func removeInt64FromArray(array *[]int64, value int64) {
	log.Infof("Remove: Input array: %v, value: %d", *array, value)
	result := []int64{}
	for _, v := range *array {
		if v != value {
			result = append(result, v)
		}
	}
	*array = result
	log.Infof("Remove: Output array: %v, value: %d", *array, value)

}

func toggleInt64InArray(array *[]int64, value int64) {
	log.Infof("Toggle: Input array: %v, value: %d", *array, value)
	for _, v := range *array {
		if v == value {
			removeInt64FromArray(array, value)
			return
		}
	}
	addInt64ToArray(array, value)
	log.Infof("Toggle: Output array: %v, value: %d", *array, value)
}
