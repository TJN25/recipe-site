package store

import (
	"errors" // For returning errors
	"fmt"    // For potential error messages

	log "github.com/sirupsen/logrus"

	"github.com/TJN25/recipe-site/internal/data"  // Import dummy data package (adjust path if needed)
	"github.com/TJN25/recipe-site/internal/model" // Import model package (adjust path if needed)
)

var AllFoodItemsCache map[int64]model.FoodItem
var AllDummyStepsCache map[int64]data.DummyRecipeStep

func InitializeCaches() {
	AllFoodItemsCache = make(map[int64]model.FoodItem)
	// Ensure data.DummyFoodItems uses the NEW model.FoodItem structure
	for _, item := range data.DummyFoodItems {
		AllFoodItemsCache[item.ID] = item
	}
	log.Infof("Cached %d FoodItems", len(AllFoodItemsCache))

	// Init Dummy Steps
	AllDummyStepsCache = make(map[int64]data.DummyRecipeStep)
	// Ensure data.DummyRecipeSteps is the slice of raw dummy step data
	for _, step := range data.DummyRecipeSteps {
		// Add validation if needed: Check for duplicate step IDs?
		AllDummyStepsCache[step.ID] = step
	}
	log.Infof("Cached %d DummyRecipeSteps", len(AllDummyStepsCache))
}

func findFoodItem(id int64) (model.FoodItem, bool) {
	for _, item := range data.DummyFoodItems {
		if item.ID == id {
			return item, true
		}
	}
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
	for i := range data.DummyRecipes {
		if data.DummyRecipes[i].ID == id {
			coreRecipeData = &data.DummyRecipes[i]
			break
		}
	}
	if coreRecipeData == nil {
		return nil, fmt.Errorf("recipe with ID %d not found in DummyRecipes: %w", id, ErrNotFound)
	}

	recipe := &model.Recipe{
		ID:            coreRecipeData.ID,
		RecipeStepIds: coreRecipeData.RecipeStepIds,
		Title:         coreRecipeData.Title,
		Description:   coreRecipeData.Description,
		Servings:      coreRecipeData.Servings,
		Notes:         coreRecipeData.Notes,
		ImagePath:     coreRecipeData.ImagePath,
		RecipeSteps:   []model.RecipeStep{}, // Initialize slices
		Tags:          []model.Tag{},
	}

	for _, dummyStep := range data.DummyRecipeSteps {
		if recipeContainsStep(dummyStep.ID, recipe.RecipeStepIds) {
			recipeStep := model.RecipeStep{
				ID:          dummyStep.ID,
				StepOrder:   dummyStep.StepOrder,
				Title:       dummyStep.Title,
				Description: dummyStep.Description,
				Notes:       dummyStep.Notes,
				Ingredients: []model.RecipeIngredient{},
				MethodSteps: dummyStep.MethodSteps,
				Equipment:   []model.Equipment{},
			}

			for _, ingRef := range dummyStep.Ingredients {
				// 1. Lookup FoodItem
				foodItem, found := findFoodItem(ingRef.FoodItemID)
				if !found {
					log.Warnf("Store: FoodItem ID %d not found referenced in RecipeStep ID %d. Skipping ingredient.", ingRef.FoodItemID, dummyStep.ID)
					continue
				}

				// 2. Basic Validation (Optional but good)
				if foodItem.DefaultFormName == "" {
					log.Warnf("Store: FoodItem ID %d ('%s') has empty DefaultFormName. Data might be incomplete.", foodItem.ID, foodItem.Name)
					// We can still proceed, DefaultFormName is mostly for substitutions now
				}

				// 3. Create the RecipeIngredient - Directly recording source data
				newIngredient := model.RecipeIngredient{
					FoodItemID: foodItem.ID,
					// Store the default form name for reference / potential initial display choice
					FormName:      foodItem.DefaultFormName,
					Quantity:      ingRef.Quantity, // Store quantity as given
					SpecifiedUnit: ingRef.Unit,     // <<< STORE THE UNIT FROM THE SOURCE
					IsOptional:    ingRef.IsOptional,
					Purpose:       ingRef.Purpose,
				}

				// 4. Append
				recipeStep.Ingredients = append(recipeStep.Ingredients, newIngredient)
			}

			for _, equipID := range dummyStep.EquipmentIDs {
				equipment, found := findEquipment(equipID)
				if !found {
					log.Warnf("Warning: Equipment ID %d not found for RecipeStep ID %d\n", equipID, dummyStep.ID)
					continue // Skip if not found
				}
				recipeStep.Equipment = append(recipeStep.Equipment, equipment)
			}

			// Add the fully assembled step to the recipe
			recipe.RecipeSteps = append(recipe.RecipeSteps, recipeStep)
		}
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

func recipeContainsStep(id int64, recipeStepIds []int64) bool {
	for _, step := range recipeStepIds {
		if step == id {
			return true
		}
	}
	return false
}

// Use the following to find the correct form for the recipe, and ensure units are correct (along with quantity)
// func FindFoodItemForm(unit string, forms map[string]model.FoodItemFormDetails) (string, string) // find the FormName based on the unit
// // Not sure if we need to do this
// func SetFoodItemQuantity(q float64, unit string, form model.FoodItemFormDetails) (float64, string) // take the relevant Form, the current unit and the current quantity and set the new quantity and unit based on the default for that form
