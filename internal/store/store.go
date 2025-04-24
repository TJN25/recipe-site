package store

import (
	"errors" // For returning errors
	"fmt"    // For potential error messages

	log "github.com/sirupsen/logrus"

	"github.com/TJN25/recipe-site/internal/data"  // Import dummy data package (adjust path if needed)
	"github.com/TJN25/recipe-site/internal/model" // Import model package (adjust path if needed)
)

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
		ID          int64
		Title       string
		Description string
		Servings    int
		Notes       string
		ImagePath   string
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
		ID:          coreRecipeData.ID,
		Title:       coreRecipeData.Title,
		Description: coreRecipeData.Description,
		Servings:    coreRecipeData.Servings,
		Notes:       coreRecipeData.Notes,
		ImagePath:   coreRecipeData.ImagePath,
		RecipeSteps: []model.RecipeStep{}, // Initialize slices
		Tags:        []model.Tag{},
	}

	for _, dummyStep := range data.DummyRecipeSteps {
		if dummyStep.RecipeID == id {
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
				foodItem, found := findFoodItem(ingRef.FoodItemID)
				if !found {
					log.Warnf("Warning: FoodItem ID %d not found for RecipeStep ID %d\n", ingRef.FoodItemID, dummyStep.ID)
					continue // Skip this ingredient if not found
				}

				recipeStep.Ingredients = append(recipeStep.Ingredients, model.RecipeIngredient{
					FoodItem:   foodItem,
					Quantity:   ingRef.Quantity,
					Unit:       ingRef.Unit,
					IsOptional: ingRef.IsOptional,
					Purpose:    ingRef.Purpose,
				})
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
	recipe.CalculateTotals() // Call the method on the assembled recipe

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
