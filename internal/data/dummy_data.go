package data

import (
	"github.com/TJN25/recipe-site/internal/model"
)

var placeholderNutrition = model.Nutrition{}
var placeHolderPrice = float32(1.0)

var DummyFoodItems = []model.FoodItem{
	{
		ID:                 1,
		Name:               "Chicken",
		FormComparisonUnit: "g",
		DefaultFormName:    "Chicken Breast",
		Forms: map[string]model.FoodItemFormDetails{
			"Chicken Breast":  {Unit: "g", ConversionScaleFactor: 1.0},
			"Chicken Thigh":   {Unit: "g", ConversionScaleFactor: 1.0},
			"Mexican Chicken": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 2,
		Name:               "Garlic",
		FormComparisonUnit: "g",
		DefaultFormName:    "Garlic Clove",
		Forms: map[string]model.FoodItemFormDetails{
			"Garlic Clove":  {Unit: "clove", ConversionScaleFactor: 5.0},
			"Garlic Paste":  {Unit: "ml", ConversionScaleFactor: 1.0},
			"Garlic Powder": {Unit: "ml", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 3,
		Name:               "Ginger",
		FormComparisonUnit: "g",
		DefaultFormName:    "Fresh Ginger Root",
		Forms: map[string]model.FoodItemFormDetails{
			"Fresh Ginger Root": {Unit: "g", ConversionScaleFactor: 1.0}, // often recorded as cm ?
			"Ginger Paste":      {Unit: "ml", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 4,
		Name:               "Onion",
		FormComparisonUnit: "g",
		DefaultFormName:    "Brown Onion",
		Forms: map[string]model.FoodItemFormDetails{
			"Brown Onion":  {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 150.0},
			"Red Onion":    {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 150.0},
			"Shallot":      {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 150.0}, // placeholder
			"Spring Onion": {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 150.0}, // placeholder
			"Leek Onion":   {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 150.0}, // placeholder
		},
	},
	{
		ID:                 5,
		Name:               "Cheese",
		FormComparisonUnit: "g",
		DefaultFormName:    "Cheddar Cheese",
		Forms: map[string]model.FoodItemFormDetails{
			"Cheddar Cheese":    {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 1.8},
			"Everday Cheese":    {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 1.8},
			"Cheese Slice":      {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 10},
			"Parmesan Cheese":   {Unit: "g", ConversionScaleFactor: 1.0},
			"Mozzarella Cheese": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 6,
		Name:               "Tomato",
		FormComparisonUnit: "g",
		DefaultFormName:    "Fresh Tomato",
		Forms: map[string]model.FoodItemFormDetails{
			"Fresh Tomato":           {Unit: "unit", ConversionScaleFactor: 120.0},
			"Canned Tomates (whole)": {Unit: "unit", ConversionScaleFactor: 400.0},
			"Canned Tomates (diced)": {Unit: "unit", ConversionScaleFactor: 400.0},
		},
	},
	{
		ID:                 7,
		Name:               "Sugar",
		FormComparisonUnit: "g",
		DefaultFormName:    "Raw Sugar",
		Forms: map[string]model.FoodItemFormDetails{
			"Raw Sugar":    {Unit: "g", ConversionScaleFactor: 1.0},
			"Brown Sugar":  {Unit: "g", ConversionScaleFactor: 1.0},
			"Caster Sugar": {Unit: "g", ConversionScaleFactor: 1.0},
			"Icing Sugar":  {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 8,
		Name:               "Rice",
		FormComparisonUnit: "g",
		DefaultFormName:    "Basmati Rice",
		Forms: map[string]model.FoodItemFormDetails{
			"Basmati Rice": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 9,
		Name:               "Salt",
		FormComparisonUnit: "g",
		DefaultFormName:    "Kosher Salt",
		Forms: map[string]model.FoodItemFormDetails{
			"Kosher Salt":    {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 0.83},
			"Flaky Sea Salt": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 0.83},
		},
	},
	{
		ID:                 10,
		Name:               "Pepper",
		FormComparisonUnit: "g",
		DefaultFormName:    "Ground Black Pepper",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground Black Pepper": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 2.0},
			"Ground White Pepper": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 2.0},
		},
	},
	{
		ID:                 11,
		Name:               "Pasta",
		FormComparisonUnit: "g",
		DefaultFormName:    "Fettuccine",
		Forms: map[string]model.FoodItemFormDetails{
			"Fettuccine": {Unit: "g", ConversionScaleFactor: 1.0},
			"Macaroni":   {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 1.1},
			"Spaghetti":  {Unit: "g", ConversionScaleFactor: 1.0},
			"Lasagna":    {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 12,
		Name:               "Cumin",
		FormComparisonUnit: "g",
		DefaultFormName:    "Ground Cumin",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground Cumin": {Unit: "g", ConversionScaleFactor: 1.0},
			"Cumin Seed":   {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 13,
		Name:               "Coriander Spice",
		FormComparisonUnit: "g",
		DefaultFormName:    "Ground Coriander",
		Forms: map[string]model.FoodItemFormDetails{
			"Ground Coriander": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 14,
		Name:               "Mustard",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Mustard Powder",
		Forms: map[string]model.FoodItemFormDetails{
			"Mustard Powder":   {Unit: "ml", ConversionScaleFactor: 1.0, ToGrams: 0.6, ToMl: 1.0},
			"American Mustard": {Unit: "ml", ConversionScaleFactor: 5.0, ToGrams: 1.05, ToMl: 1.0},
			"Dijon Mustard":    {Unit: "ml", ConversionScaleFactor: 5.0, ToGrams: 1.05, ToMl: 1.0},
			"Honey Mustard":    {Unit: "ml", ConversionScaleFactor: 5.0, ToGrams: 1.05, ToMl: 1.0},
		},
	},
	{
		ID:                 15,
		Name:               "Milk",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Full Fat Milk",
		Forms: map[string]model.FoodItemFormDetails{
			"Full Fat Milk": {Unit: "ml", ConversionScaleFactor: 1.0, ToGrams: 1.03, ToMl: 1.0},
			"Trim Milk":     {Unit: "ml", ConversionScaleFactor: 1.0, ToGrams: 1.03, ToMl: 1.0},
		},
	},
	{
		ID:                 16,
		Name:               "Smoked Paprika",
		FormComparisonUnit: "g",
		DefaultFormName:    "Smoked Paprika",
		Forms: map[string]model.FoodItemFormDetails{
			"Smoked Paprika": {Unit: "ml", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 2.0},
		},
	},
	{
		ID:                 17,
		Name:               "Flour",
		FormComparisonUnit: "g",
		DefaultFormName:    "All Purpose Flour",
		Forms: map[string]model.FoodItemFormDetails{
			"All Purpose Flour": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 1.8},
			"Bread Flour":       {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 1.8},
		},
	},
	{
		ID:                 18,
		Name:               "Butter",
		FormComparisonUnit: "g",
		DefaultFormName:    "Butter",
		Forms: map[string]model.FoodItemFormDetails{
			"Butter": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 1.04},
		},
	},
	{
		ID:                 19,
		Name:               "Chili",
		FormComparisonUnit: "unit",
		DefaultFormName:    "Chili",
		Forms: map[string]model.FoodItemFormDetails{
			"Chili":          {Unit: "unit", ConversionScaleFactor: 1.0},
			"Bird-Eye Chili": {Unit: "unit", ConversionScaleFactor: 1.0},
			"Kashmiri Chili": {Unit: "unit", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 20,
		Name:               "Soy Sauce",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Soy Sauce",
		Forms: map[string]model.FoodItemFormDetails{
			"Soy Sauce": {Unit: "ml", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 21,
		Name:               "Oil",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Canola Oil",
		Forms: map[string]model.FoodItemFormDetails{
			"Canola Oil":  {Unit: "ml", ConversionScaleFactor: 1.0},
			"Sesame Oil":  {Unit: "ml", ConversionScaleFactor: 1.0},
			"Avocado Oil": {Unit: "ml", ConversionScaleFactor: 1.0},
			"Olive Oil":   {Unit: "ml", ConversionScaleFactor: 1.0},
			"Chili Oil":   {Unit: "ml", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 22,
		Name:               "Sichuan Pepper",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Sichuan Pepper",
		Forms: map[string]model.FoodItemFormDetails{
			"Sichuan Pepper": {Unit: "ml", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 23,
		Name:               "Allspice",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Allspice",
		Forms: map[string]model.FoodItemFormDetails{
			"Allspice": {Unit: "ml", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 24,
		Name:               "Egg",
		FormComparisonUnit: "unit",
		DefaultFormName:    "Egg",
		Forms: map[string]model.FoodItemFormDetails{
			"Egg": {Unit: "unit", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 25,
		Name:               "Cornflour",
		FormComparisonUnit: "g",
		DefaultFormName:    "Cornflour",
		Forms: map[string]model.FoodItemFormDetails{
			"Cornflour": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 26,
		Name:               "Cabbage",
		FormComparisonUnit: "g",
		DefaultFormName:    "Cabbage",
		Forms: map[string]model.FoodItemFormDetails{
			"Cabbage":       {Unit: "g", ConversionScaleFactor: 1.0},
			"Red Cabbage":   {Unit: "g", ConversionScaleFactor: 1.0},
			"White Cabbage": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 27,
		Name:               "Carrot",
		FormComparisonUnit: "g",
		DefaultFormName:    "Carrot",
		Forms: map[string]model.FoodItemFormDetails{
			"Carrot": {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 50},
		},
	},
	{
		ID:                 28,
		Name:               "Cucumber",
		FormComparisonUnit: "g",
		DefaultFormName:    "Cucumber",
		Forms: map[string]model.FoodItemFormDetails{
			"Cucumber": {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 100},
		},
	},
	{
		ID:                 29,
		Name:               "Peanut",
		FormComparisonUnit: "g",
		DefaultFormName:    "Roasted Peanut",
		Forms: map[string]model.FoodItemFormDetails{
			"Roasted Peanut": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 30,
		Name:               "Mayonnaise",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Mayonnaise",
		Forms: map[string]model.FoodItemFormDetails{
			"Mayonnaise":        {Unit: "ml", ConversionScaleFactor: 1.0, ToGrams: 1.0},
			"Kewpie Mayonnaise": {Unit: "ml", ConversionScaleFactor: 1.0, ToGrams: 1.0},
		},
	},
	{
		ID:                 31,
		Name:               "Sriracha Sauce",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Sriracha Sauce",
		Forms: map[string]model.FoodItemFormDetails{
			"Sriracha Sauce": {Unit: "ml", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 32,
		Name:               "Honey",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Honey",
		Forms: map[string]model.FoodItemFormDetails{
			"Honey": {Unit: "ml", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 33,
		Name:               "Lemon",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Lemon Juice",
		Forms: map[string]model.FoodItemFormDetails{
			"Lemon Juice": {Unit: "ml", ConversionScaleFactor: 1.0},
			"Lemon":       {Unit: "unit", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 34,
		Name:               "Bun",
		FormComparisonUnit: "unit",
		DefaultFormName:    "Bun",
		Forms: map[string]model.FoodItemFormDetails{
			"Bao Bun":     {Unit: "unit", ConversionScaleFactor: 1.0},
			"Brioche Bun": {Unit: "unit", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 35,
		Name:               "Green Bean",
		FormComparisonUnit: "g",
		DefaultFormName:    "Green Bean",
		Forms: map[string]model.FoodItemFormDetails{
			"Green Bean": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 36,
		Name:               "Mushroom",
		FormComparisonUnit: "g",
		DefaultFormName:    "Mushroom",
		Forms: map[string]model.FoodItemFormDetails{
			"Mushroom":            {Unit: "g", ConversionScaleFactor: 1.0},
			"Button Mushroom":     {Unit: "g", ConversionScaleFactor: 1.0},
			"Portabello Mushroom": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 37,
		Name:               "Capsicum",
		FormComparisonUnit: "unit",
		DefaultFormName:    "Capsicum",
		Forms: map[string]model.FoodItemFormDetails{
			"Capsicum":         {Unit: "unit", ConversionScaleFactor: 1.0},
			"Red Capsicum":     {Unit: "unit", ConversionScaleFactor: 1.0},
			"Yellow Capsicum":  {Unit: "unit", ConversionScaleFactor: 1.0},
			"Green Capsicum":   {Unit: "unit", ConversionScaleFactor: 1.0},
			"Roasted Capsicum": {Unit: "unit", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 38,
		Name:               "Spinach",
		FormComparisonUnit: "g",
		DefaultFormName:    "Baby Spinach",
		Forms: map[string]model.FoodItemFormDetails{
			"Baby Spinach": {Unit: "g", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 39,
		Name:               "Water",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Water",
		Forms: map[string]model.FoodItemFormDetails{
			"Water": {Unit: "ml", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 1.0},
		},
	},
	{
		ID:                 40,
		Name:               "Noodle",
		FormComparisonUnit: "unit",
		DefaultFormName:    "Noodle",
		Forms: map[string]model.FoodItemFormDetails{
			"Egg Noodle": {Unit: "unit", ConversionScaleFactor: 1.0},
		},
	},
	{
		ID:                 41,
		Name:               "Mince",
		FormComparisonUnit: "g",
		DefaultFormName:    "Mince",
		Forms: map[string]model.FoodItemFormDetails{
			"Mince": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0},
		},
	},
	{
		ID:                 42,
		Name:               "Pickles",
		FormComparisonUnit: "g",
		DefaultFormName:    "Pickles",
		Forms: map[string]model.FoodItemFormDetails{
			"Pickles": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0},
		},
	},
	{
		ID:                 43,
		Name:               "Tomato Sauce",
		FormComparisonUnit: "ml",
		DefaultFormName:    "Tomato Sauce",
		Forms: map[string]model.FoodItemFormDetails{
			"Tomato Sauce": {Unit: "ml", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 1.0},
		},
	},
	{
		ID:                 44,
		Name:               "Lettuce",
		FormComparisonUnit: "g",
		DefaultFormName:    "Lettuce leaf",
		Forms: map[string]model.FoodItemFormDetails{
			"Lettuce leaf": {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 20.0},
		},
	},
	{
		ID:                 45,
		Name:               "Yeast",
		FormComparisonUnit: "g",
		DefaultFormName:    "Active Dry Yeast",
		Forms: map[string]model.FoodItemFormDetails{
			"Active Dry Yeast": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 1.25},
		},
	},
	{
		ID:                 46,
		Name:               "Italian Herb",
		FormComparisonUnit: "g",
		DefaultFormName:    "Italian Herb",
		Forms: map[string]model.FoodItemFormDetails{
			"Italian Herb": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 5},
		},
	},
	{
		ID:                 47,
		Name:               "Potato",
		FormComparisonUnit: "g",
		DefaultFormName:    "Potato",
		Forms: map[string]model.FoodItemFormDetails{
			"Potato": {Unit: "unit", ConversionScaleFactor: 1.0, ToGrams: 200.0},
		},
	},
	{
		ID:                 48,
		Name:               "Cayenne Pepper",
		FormComparisonUnit: "g",
		DefaultFormName:    "Cayenne Pepper",
		Forms: map[string]model.FoodItemFormDetails{
			"Cayenne Pepper": {Unit: "g", ConversionScaleFactor: 1.0, ToGrams: 1.0, ToMl: 2.0},
		},
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
		Title:         "Classic Mac and Cheese",
		Description:   "A simple, comforting, and irresistibly cheesy homemade classic mac and cheese.",
		Servings:      2,
		Notes:         "Let rest 5 minutes before serving. Uses full fat milk for creaminess. Sauce should look slightly too saucy before baking; it will thicken.",
		ImagePath:     "img/mac-and-cheese.jpg",
	},
	{
		ID:            2,
		RecipeStepIds: []int64{1801, 1802, 1803, 1804, 1805},
		Title:         "Spiced Chicken Bao Buns with Sriracha Mayo",
		Description:   "Crispy fried spiced chicken pieces served in soft bao buns with coleslaw, cucumber, peanuts, and a tangy sriracha mayo.",
		Servings:      3,
		Notes:         "Marinate chicken for 15-20 mins. Fry chicken in batches. Steam buns just before serving. Assemble just before eating for best texture contrast. Enhancements: Salt cucumber slices briefly before use. Toast peanuts before crushing.",
		ImagePath:     "img/bao-buns.jpg", // Assign an appropriate image path
	},
	{
		ID:            3,
		RecipeStepIds: []int64{1801, 1902, 1803, 1903, 1905, 1906},
		Title:         "Crispy Fried Chicken with Vegetable & Egg Noodle Stir-fry",
		Description:   "Crispy fried chicken pieces served over a flavorful stir-fry of vegetables and egg noodles.",
		Servings:      2,
		Notes:         "Wok Hei: Preheating the wok properly over high heat is crucial for achieving \"wok hei\" - the characteristic smoky flavour of good stir-fries.\nWok Frying: Be mindful when deep-frying/shallow-frying in a wok. The sloped sides mean oil depth varies. Keep pieces moving and adjust heat to prevent burning. Use a wok spatula or spider strainer for removal.\nStir-fry Motion: Use a scooping, tossing motion to move ingredients constantly, ensuring even cooking and preventing sticking. Add sauce around the perimeter to allow it to heat and reduce slightly before coating ingredients.\nSpeed: Wok cooking is fast. Have everything prepped and ready next to the stove before you start heating the wok.",
		ImagePath:     "img/noodles.jpg", // Placeholder path
	},
	{
		ID:            4,
		RecipeStepIds: []int64{2001, 2002, 2003, 2004},
		Title:         "Cheesebuger with Smash Patty",
		Description:   "Enjoy a delicious homemade classic smash cheeseburger featuring thin, crispy-edged patties, perfectly melted cheese, and your favourite toppings on a toasted bun.",
		Servings:      2,
		Notes:         "To help melt the cheese, add a splash of water to the frypan and a metal bowl (or lid) on top of the patties which will let them . ",
		ImagePath:     "img/cheeseburger.jpg", // Placeholder path
	},
	{
		ID:            5,
		RecipeStepIds: []int64{5001, 5002, 5003, 5004, 5005, 5006},
		Title:         "Pizza Dough",
		Description:   "A reliable recipe for crafting delicious, chewy pizza dough ready in about 4-5 hours, perfect for a same-day pizza night.",
		Servings:      2,
		Notes:         "This recipe uses 75% hydration dough for a tender crumb. Warm water (40-46°C) is crucial for activating the yeast. Sugar feeds the yeast and helps with browning. Salt enhances flavor and controls fermentation. Olive oil adds tenderness and flavor.\n\nTo help the sauce thicken slightly and coat ingredients better, add it around the perimeter of the pan during stir-frying.\n\nOvernight Fermentation Option:\n- Reduce yeast to 1-2g (~1/4-1/3 tsp).\n- Complete only 1-2 stretch and folds.\n- Limit room temperature time to max 1 hour before refrigerating.\n- Start around 3pm for cooking at 7pm the next day.\n- Handle overproofed dough gently; it may be better suited for focaccia-style.\n- Remove cold dough from fridge 45-60 mins before cooking to warm up.",
		ImagePath:     "img/pizza.jpg",
	},
	{
		ID:            6,
		RecipeStepIds: []int64{5001, 6002, 5003, 5004, 6005, 6006},                                                                                                                                                                                                                                                                                                                                                                                                                                                                       // Updated RecipeStep IDs
		Title:         "Classic Focaccia",                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // More descriptive title
		Description:   "A simple and delicious classic focaccia with a tender interior, crispy crust, and fragrant olive oil and sea salt finish.",                                                                                                                                                                                                                                                                                                                                                                                       // Added description
		Servings:      2,                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 // Assumes makes a batch equivalent to 2 pizzas
		Notes:         "This recipe starts with a 75% hydration dough for a soft, open crumb. Baking it in a pan helps achieve the characteristic shape and texture. Be generous with the olive oil before baking for best results.\n\nHandling Overproofed Dough: If your pizza dough recipe (using reduced yeast for overnight fermentation) overproofs and becomes too delicate to stretch thinly for pizza, it can be successfully transformed into focaccia. Handle it gently and proceed with the shaping and baking steps below.", // Added relevant notes
		ImagePath:     "img/focaccia.jpg",
	},
	{
		ID:            7,
		RecipeStepIds: []int64{7001, 7002, 7003, 7004},
		Title:         "Agria Potato Wedges",
		Description:   "Crispy and fluffy Agria potato wedges seasoned with a savoury blend of paprika, garlic, and herbs, perfect as a side dish or snack.",
		Servings:      2,                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   // Assuming 2 medium potatoes serves 2 as a side
		Notes:         "Using Agria potatoes yields a fluffy interior and crisp exterior. Leaving the skin on adds texture and nutrients. Soaking removes excess starch for crispiness. Patting dry is essential for crispy wedges. Avocado oil has a high smoke point suitable for high heat. For extra crispy wedges, parboil the potatoes for 5 minutes before seasoning and baking.\n\nEnhancements:\n- Texture: Parboil potatoes for 5 mins before seasoning and baking.\n- Flavor: Add 1 tbsp finely grated parmesan during the last 5 mins of cooking.\n- Dipping Sauce: Serve with Greek yogurt + pinch smoked paprika + minced garlic.\n\nUser Note: Adding ½ tsp cayenne and using canola oil also works well, and cook time may vary (e.g., ~25 mins in oven).", // Consolidated notes
		ImagePath:     "img/wedges.jpg",                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    // Suggested image path
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
	Preparation  string
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
			{FoodItemName: "Chili", FormName: "Bird Eye Chili", Quantity: 1, Unit: "unit", Purpose: "heat"},          // Dried Bird's Eye Chili
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
			{FoodItemName: "Cabbage", Quantity: 100, Unit: "g", Purpose: "slaw base"},                                   // Cabbage, shredded
			{FoodItemName: "Carrot", Quantity: 1, Unit: "unit", Purpose: "slaw color"},                                  // Carrot, grated
			{FoodItemName: "Cucumber", Quantity: 0.5, Unit: "unit", Purpose: "freshness"},                               // Cucumber (1/2), thinly sliced
			{FoodItemName: "Peanut", FormName: "Roasted Peanut", Quantity: 25, Unit: "g", Purpose: "garnish texture"},   // Roasted Peanuts (2 Tbsp), crushed
			{FoodItemName: "Mayonnaise", FormName: "Kewpie Mayonnaise", Quantity: 45, Unit: "ml", Purpose: "mayo base"}, // Kewpie Mayo (3 Tbsp)
			{FoodItemName: "Sriracha Sauce", Quantity: 10, Unit: "ml", Purpose: "mayo heat"},                            // Sriracha Sauce (1-2 tsp)
			{FoodItemName: "Honey", Quantity: 2.5, Unit: "ml", Purpose: "mayo balance"},                                 // Honey (1/2 tsp)
			{FoodItemName: "Lemon", FormName: "Lemon Juice", Quantity: 1, Unit: "ml", Purpose: "mayo brightness"},       // Lemon Juice (Few drops)
			{FoodItemName: "Oil", FormName: "Sesame Oil", Quantity: 1, Unit: "ml", Purpose: "mayo aroma"},               // Sesame Oil (1/4 tsp / drop)
			{FoodItemName: "Onion", FormName: "Spring Onion", Quantity: 20, Unit: "g", Purpose: "garnish"},              // Spring Onion, thinly sliced
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
		Title:       "Fry Chicken",
		Description: "Fry marinated chicken until crispy.",
		Ingredients: []dummyIngredientRef{
			// Marinated Chicken from Step 1
			{FoodItemName: "Oil", Quantity: 45, Unit: "ml", Purpose: "frying"}, // Canola Oil (3 Tbsp)
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Heat canola oil in a wok or large frypan over medium-high heat."},
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
		Title:       "Steam Bao Buns",
		Description: "Steam bao buns.",
		Ingredients: []dummyIngredientRef{
			// Marinated Chicken from Step 1
			{FoodItemName: "Bun", FormName: "Bao Buns", Quantity: 4, Unit: "unit", Purpose: "vessel"}, // Bao Buns (adjust number as needed)
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
	{
		ID:          2001,
		RecipeID:    4,
		StepOrder:   1,
		Title:       "Make Burger Sauce",
		Description: "",
		Ingredients: []dummyIngredientRef{
			// Oil for stir-fry is listed as needed in method step 4
			{FoodItemName: "Pickles", Quantity: 10, Unit: "g", Purpose: "sour"},
			{FoodItemName: "Onion", FormName: "Red Onion", Quantity: 10, Unit: "g", Purpose: "aromatic/bite"},
			{FoodItemName: "Mayonnaise", FormName: "Kewpie Mayonnaise", Quantity: 10, Unit: "g", Purpose: "umami/fat"},
			{FoodItemName: "Mustard", FormName: "American Mustard", Quantity: 10, Unit: "g", Purpose: "sour"},
			{FoodItemName: "Tomato Sauce", Quantity: 20, Unit: "g", Purpose: "sweet"},
			{FoodItemName: "Smoked paprika", Quantity: 2, Unit: "dash", Purpose: "aromtic/smoke"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Combine all ingredients for the burger sauce in a bowl."},
			{StepNumber: 2, Instruction: "Mix well and taste, adjusting seasoning as needed. Set aside."},
		},
	},
	{
		ID:          2002,
		RecipeID:    4,
		StepOrder:   2,
		Title:       "Make burger patties",
		Description: "",
		Ingredients: []dummyIngredientRef{
			// Oil for stir-fry is listed as needed in method step 4
			{FoodItemName: "Mince", Quantity: 280, Unit: "g", Purpose: "main"},
			{FoodItemName: "Cheese", FormName: "Cheese Slice", Quantity: 2, Unit: "unit", Purpose: "fat/creamy"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Preheat frypan (cast iron works best) over medium-high heat"},
			{StepNumber: 2, Instruction: "Divide the ground mince into two (per burger) equal-sized balls."},
			{StepNumber: 3, Instruction: "Carefully place the meat balls in the hot frypan, leaving some space between them. Let them sear for a few seconds to prevent sticking."},
			{StepNumber: 4, Instruction: "Flip the meat balls over and use a sturdy, flat spatula to firmly smash each ball down into a thin patty. Apply good pressure."},
			{StepNumber: 4, Instruction: "Let them cook undisturbed for 1-2 minutes until the edges are crispy and browned."},
			{StepNumber: 5, Instruction: "Flip the patties. Let them further cook."},
			{StepNumber: 6, Instruction: "Place a slice of cheese on top of each patty. Stack one cheeseburger patty on top of the other in the frypan"},
			{StepNumber: 7, Instruction: "Cover (using the water/steam method from the notes if desired) and cook for another 30 seconds to 1 minute, or until the cheese is fully melted and the patties are cooked through."},
			{StepNumber: 8, Instruction: "Remove the patties from the heat once the cheese is melted."},
		},
	},
	{
		ID:          2003,
		RecipeID:    4,
		StepOrder:   3,
		Title:       "Toast buns",
		Description: "",
		Ingredients: []dummyIngredientRef{
			// Oil for stir-fry is listed as needed in method step 4
			{FoodItemName: "Bun", FormName: "Brioche Bun", Quantity: 2, Unit: "unit", Purpose: ""},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Add a bun to the preheated frypan."},
			{StepNumber: 2, Instruction: "Toast for 1-2 minutes until golden brown and slightly crispy."},
		},
	},
	{
		ID:          2004,
		RecipeID:    4,
		StepOrder:   4,
		Title:       "Assemble",
		Description: "",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Lettuce", Quantity: 2, Unit: "unit", Purpose: ""},
			{FoodItemName: "Cheese", FormName: "Cheese Slice", Quantity: 2, Unit: "unit", Purpose: "fat/creamy"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Finely slice your lettuce and prepare any other desired toppings."},
			{StepNumber: 2, Instruction: "Add sauce to top and botton of the burger bun"},
			{StepNumber: 3, Instruction: "Spread a generous amount of the prepared burger sauce on the cut sides of both the top and bottom buns."},
			{StepNumber: 4, Instruction: "Place the stacked cheeseburger patties on the bottom bun."},
			{StepNumber: 5, Instruction: "Top with sliced lettuce and any other desired toppings."},
			{StepNumber: 6, Instruction: "Add the top bun. Serve and enjoy!"},
		},
	},
	{
		ID:          5001,
		RecipeID:    5,
		StepOrder:   1,
		Title:       "Activate Yeast",
		Description: "",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Water", Quantity: 210, Unit: "ml", Purpose: "hydration"},
			{FoodItemName: "Yeast", FormName: "Active Dry Yeast", Quantity: 5, Unit: "g", Purpose: "leavening"},
			{FoodItemName: "Sugar", FormName: "Raw Sugar", Quantity: 5, Unit: "g", Purpose: "yeast food, browning"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In a medium bowl, mix the warm water (40-46°C) with the active dry yeast and raw sugar."},
			{StepNumber: 2, Instruction: "Let the mixture stand for 10-15 minutes until it is foamy and bubbly. This confirms the yeast is active."},
		},
	},
	{
		ID:          5002,
		RecipeID:    5,
		StepOrder:   2,
		Title:       "Make the Dough",
		Description: "",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Flour", FormName: "Bread Flour", Quantity: 280, Unit: "g", Purpose: "structure, gluten development"},
			{FoodItemName: "Salt", FormName: "Kosher Salt", Quantity: 5, Unit: "g", Purpose: "flavor, fermentation control"},
			{FoodItemName: "Oil", FormName: "Olive Oil", Quantity: 8, Unit: "g", Purpose: "tenderness, flavor"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In a separate larger bowl, whisk together the bread flour and salt."},
			{StepNumber: 2, Instruction: "Pour the activated yeast mixture and the olive oil into the bowl with the flour."},
			{StepNumber: 3, Instruction: "Mix with a fork or spatula until there is no dry flour remaining and a shaggy dough forms."},
			{StepNumber: 4, Instruction: "Cover the bowl and let the dough rest for 10 minutes."},
		},
	},
	{
		ID:          5003,
		RecipeID:    5,
		StepOrder:   3,
		Title:       "Develop Gluten (Stretch and Folds)",
		Description: "",
		Ingredients: []dummyIngredientRef{}, // No new ingredients added in this step
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Perform the first set of stretch and folds: wet your hands slightly, grab an edge of the dough, stretch it upwards, and fold it back over the center. Repeat this process 3-4 times around the dough."},
			{StepNumber: 2, Instruction: "Cover the bowl and let the dough rest for 30 minutes."},
			{StepNumber: 3, Instruction: "Perform the second set of stretch and folds. Cover and rest for 30 minutes."},
			{StepNumber: 4, Instruction: "Perform the third set of stretch and folds. Cover and rest for 30 minutes."},
			{StepNumber: 5, Instruction: "Perform a final stretch and fold if needed (the dough should be noticeably smoother and stronger). Cover and rest for the final 30 minutes."},
		},
	},
	{
		ID:          5004,
		RecipeID:    5,
		StepOrder:   4,
		Title:       "Divide and Shape Dough",
		Description: "",
		Ingredients: []dummyIngredientRef{}, // No new ingredients added in this step
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Turn the dough out onto a lightly floured surface."},
			{StepNumber: 2, Instruction: "The total dough weight should be approximately 490g. Divide the dough into two equal portions, roughly 245g each."},
			{StepNumber: 3, Instruction: "Shape each portion into a tight ball using the edge-folding technique (tuck the edges underneath the dough ball to create surface tension)."},
			{StepNumber: 4, Instruction: "Place the dough balls on a lightly floured surface, cover loosely with plastic wrap or a damp towel, and let them rest at room temperature until you are ready to stretch and cook them."},
		},
	},
	{
		ID:          5005,
		RecipeID:    5,
		StepOrder:   5,
		Title:       "Prepare for Cooking",
		Description: "",
		Ingredients: []dummyIngredientRef{}, // No new ingredients added in this step
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "While the dough completes its final rise, prepare all your pizza toppings."},
			{StepNumber: 2, Instruction: "Set up your cold cast iron pan and lightly coat it with cooking spray."},
			{StepNumber: 3, Instruction: "Gently stretch one dough ball into a 12-14 inch circle. If the dough resists stretching, cover it and let it rest for 5-10 minutes before trying again."},
			{StepNumber: 4, Instruction: "Carefully transfer the stretched dough to the prepared cold cast iron pan."},
		},
	},
	{
		ID:          5006,
		RecipeID:    5,
		StepOrder:   6,
		Title:       "Cook the Pizza",
		Description: "",
		Ingredients: []dummyIngredientRef{}, // No new ingredients added in this step
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Top the dough in the pan with your sauce, cheese, and desired ingredients."},
			{StepNumber: 2, Instruction: "Preheat your oven to 250°C (Fan Bake)."},
			{StepNumber: 3, Instruction: "Place the pan on a medium-high heat stovetop burner and cook for 5-7 minutes. Watch for slight bubbling on the top surface of the dough."},
			{StepNumber: 4, Instruction: "Carefully transfer the pan from the stovetop to the preheated oven."},
			{StepNumber: 5, Instruction: "Bake for an initial 10-12 minutes. After 10 minutes, check for doneness: the cheese should be fully melted and starting to brown, the crust edge should be golden brown, and the bottom should be crispy (lift with a spatula to check)."},
			{StepNumber: 6, Instruction: "If needed, return the pizza to the oven for additional 2-minute intervals until desired doneness is reached."},
			{StepNumber: 7, Instruction: "Once cooked, carefully remove the pan from the oven. Let the pizza rest for 2-3 minutes before slicing and serving."},
		},
	},
	{
		ID:          6002,
		RecipeID:    6,
		StepOrder:   2,
		Title:       "Make the Dough",
		Description: "",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Flour", FormName: "Bread Flour", Quantity: 280, Unit: "g", Purpose: "structure, gluten development"},
			{FoodItemName: "Salt", FormName: "Kosher Salt", Quantity: 5, Unit: "g", Purpose: "flavor, fermentation control"},
			{FoodItemName: "Oil", FormName: "Olive Oil", Quantity: 15, Unit: "g", Purpose: "tenderness, flavor, rich flavour"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In a separate larger bowl, whisk together the bread flour and salt."},
			{StepNumber: 2, Instruction: "Pour the activated yeast mixture and the olive oil into the bowl with the flour."},
			{StepNumber: 3, Instruction: "Mix with a fork or spatula until there is no dry flour remaining and a shaggy dough forms."},
			{StepNumber: 4, Instruction: "Cover the bowl and let the dough rest for 10 minutes."},
		},
	},
	{
		ID:          6005, // New step for Focaccia Shaping and Preparation
		RecipeID:    6,
		StepOrder:   5,
		Title:       "Shape and Prepare for Baking",
		Description: "",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Oil", FormName: "Olive Oil", Quantity: 0, Unit: "dash", Purpose: "topping, texture"},
			{FoodItemName: "Salt", FormName: "Flaky Sea Salt", Quantity: 0, Unit: "pinch", Purpose: "topping, flavor"},
			{FoodItemName: "Italian Herbs", Quantity: 0, Unit: "pinch", IsOptional: true, Purpose: "topping, aromatic, garnish"}, // Added optional Italian Herbs
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Take the dough portion(s) and gently press or stretch it into your desired shape and thickness (about 1/2 inch thick) in a well-oiled baking pan or tray."},
			{StepNumber: 2, Instruction: "Cover the pan loosely and let the dough rise at room temperature for about 15 minutes while you preheat the oven."},
			{StepNumber: 3, Instruction: "Using your fingertips, gently but firmly press deep dimples all over the surface of the dough."},
			{StepNumber: 4, Instruction: "Generously drizzle the surface of the dough with olive oil."},
			{StepNumber: 5, Instruction: "Sprinkle the top with flaky sea salt and any other desired toppings (like fresh rosemary)."},
		},
	},
	{
		ID:          6006, // New step for Focaccia Baking
		RecipeID:    6,
		StepOrder:   6,
		Title:       "Bake and Finish",
		Description: "",
		Ingredients: []dummyIngredientRef{}, // No new ingredients added in this step
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Preheat your oven to 200°C (Fan Bake)."},
			{StepNumber: 2, Instruction: "Place the pan in the preheated oven and bake for 15-18 minutes, or until the focaccia is golden brown on top and cooked through."},
			{StepNumber: 3, Instruction: "Carefully remove the pan from the oven."},
			{StepNumber: 4, Instruction: "For a softer crust, immediately transfer the hot focaccia to a wire rack and loosely wrap it in a clean tea towel for a few minutes to trap steam."},
			{StepNumber: 5, Instruction: "Cut the focaccia into slices or squares."},
			{StepNumber: 6, Instruction: "Serve warm, perhaps with extra olive oil for dipping, and enjoy!"},
		},
	},
	{
		ID:          7001, // Placeholder ID
		RecipeID:    7,
		StepOrder:   1,
		Title:       "Prepare the Potatoes",
		Description: "Prepare the potatoes to ensure maximum crispiness and flavour.",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Potato", Quantity: 2, Unit: "unit", Purpose: "starchy, fluffy interior, crisp exterior"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Wash and scrub the Agria potatoes thoroughly, leaving the skin on for extra texture and nutrients."},
			{StepNumber: 2, Instruction: "Cut each potato in half lengthwise, then cut each half into 3-4 wedges."},
			{StepNumber: 3, Instruction: "Place the cut wedges in a bowl of cold water for 5 minutes to remove excess starch."},
			{StepNumber: 4, Instruction: "Drain the wedges and pat them dry thoroughly with paper towels. This is a crucial step for achieving crispiness."},
		},
	},
	{
		ID:          7002, // Placeholder ID
		RecipeID:    7,
		StepOrder:   2,
		Title:       "Make the Seasoning Mix",
		Description: "Combine the oil and spices to create the flavour coating for the wedges.",
		Ingredients: []dummyIngredientRef{
			{FoodItemName: "Oil", FormName: "Avocado Oil", Quantity: 2, Unit: "tbsp", Purpose: "high heat tolerance, neutral flavor"},
			{FoodItemName: "Smoked Paprika", Quantity: 1, Unit: "tsp", Purpose: "smoky flavor, color"},
			{FoodItemName: "Garlic", FormName: "Garlic Powder", Quantity: 0.5, Unit: "tsp", Purpose: "aromatic, flavor"},
			{FoodItemName: "Salt", FormName: "Kosher Salt", Quantity: 0.5, Unit: "tsp", Purpose: "seasoning"},
			{FoodItemName: "Pepper", FormName: "Ground Black Pepper", Quantity: 0.25, Unit: "tsp", Purpose: "heat, seasoning"},
			{FoodItemName: "Italian Herbs", Quantity: 1, Unit: "tsp", Purpose: "aromatic, flavor complexity"},
			{FoodItemName: "Cayenne Pepper", Quantity: 0.5, Unit: "tsp", IsOptional: true, Purpose: "heat"},
		},
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "In a large bowl, mix together the avocado oil, smoked paprika, garlic powder, salt, black pepper, and mixed herbs until well combined."},
		},
	},
	{
		ID:          7003, // Placeholder ID
		RecipeID:    7,
		StepOrder:   3,
		Title:       "Season and Arrange Wedges",
		Description: "Coat the dried potato wedges evenly with the seasoning mix and arrange for cooking.",
		Ingredients: []dummyIngredientRef{}, // Ingredients are from the previous step
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Add the dried potato wedges to the bowl with the seasoning mixture."},
			{StepNumber: 2, Instruction: "Toss the wedges until they are evenly coated with the oil and spices."},
			{StepNumber: 3, Instruction: "Arrange the seasoned wedges in a single layer on a baking tray (preferably on a frying rack) or in your air fryer basket."},
		},
	},
	{
		ID:          7004, // Placeholder ID
		RecipeID:    7,
		StepOrder:   4,
		Title:       "Cook the Wedges",
		Description: "Cook the wedges until golden brown and crispy, using either an air fryer or oven.",
		Ingredients: []dummyIngredientRef{}, // Ingredients are from previous steps
		MethodSteps: []model.MethodStep{
			{StepNumber: 1, Instruction: "Preheat your air fryer to 200°C or oven to 230°C on a fan forced setting."},
			{StepNumber: 2, Instruction: "Cook the wedges for 25-30 minutes, flipping halfway through if using an oven tray without a rack or if needed for even cooking in the air fryer."},
			{StepNumber: 3, Instruction: "The wedges are done when they are golden brown and crispy on the outside and fluffy on the inside."},
			{StepNumber: 4, Instruction: "Remove from the heat and let them cool slightly before serving."},
		},
	},
}

var DummyRecipeTags = []struct {
	RecipeID int64
	TagID    int64
}{}
