-- File: schema.sql

-- Drop existing tables in reverse order of creation dependency
-- (Use IF EXISTS to avoid errors if they don't exist yet)
DROP TABLE IF EXISTS recipe_tags CASCADE;
DROP TABLE IF EXISTS recipe_equipment CASCADE;
DROP TABLE IF EXISTS recipe_ingredients CASCADE;
DROP TABLE IF EXISTS method_steps CASCADE;
DROP TABLE IF EXISTS recipes CASCADE;
DROP TABLE IF EXISTS tags CASCADE;
DROP TABLE IF EXISTS equipment CASCADE;
DROP TABLE IF EXISTS food_items CASCADE;

-- Represents generic food items and their base properties
CREATE TABLE food_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    base_unit VARCHAR(50) NOT NULL, -- Unit for price/nutrition (e.g., 'g', 'ml', 'whole')
    price_per_base_unit NUMERIC(10, 2) DEFAULT 0.00, -- Price per base_unit
    carbs_grams REAL DEFAULT 0.0,
    fat_grams REAL DEFAULT 0.0,
    protein_grams REAL DEFAULT 0.0,
    calories_kcal REAL DEFAULT 0.0
);

-- Represents types of equipment
CREATE TABLE equipment (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    type VARCHAR(100),             -- e.g., 'Cookware', 'Utensil', 'Appliance'
    cleaning_difficulty SMALLINT DEFAULT 0 -- Corresponds to CleaningDifficulty enum (0=None, 1=Easy, 2=Medium, 3=Hard)
);

-- Represents tags for categorizing recipes
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL, -- e.g., 'cuisine', 'duration', 'season'
    UNIQUE(name, type)         -- Ensure tag names are unique within a type
);

-- Core recipe table
CREATE TABLE recipes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    servings INT DEFAULT 1 CHECK (servings > 0), -- Ensure servings is positive
    notes TEXT,
    image_path VARCHAR(255) -- Path/URL to the recipe image (nullable)
    -- created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Optional audit field
    -- updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP  -- Optional audit field
);

-- Method steps for each recipe
CREATE TABLE method_steps (
    id SERIAL PRIMARY KEY,
    recipe_id INT NOT NULL REFERENCES recipes(id) ON DELETE CASCADE, -- If recipe deleted, delete steps
    step_number INT NOT NULL,
    instruction TEXT NOT NULL,
    stage VARCHAR(100),            -- Optional grouping (e.g., 'Make Dough')
    prep_time_minutes INT DEFAULT 0,
    cook_time_minutes INT DEFAULT 0,
    UNIQUE(recipe_id, step_number) -- Ensure step numbers are unique per recipe
);

-- Join table: Links recipes to the food items they require (Many-to-Many)
CREATE TABLE recipe_ingredients (
    recipe_id INT NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
    food_item_id INT NOT NULL REFERENCES food_items(id) ON DELETE RESTRICT, -- Don't delete food item if used
    quantity REAL NOT NULL CHECK (quantity > 0),
    unit VARCHAR(50) NOT NULL, -- Unit used in the recipe (e.g., 'cup', 'tbsp')
    is_optional BOOLEAN DEFAULT FALSE NOT NULL, -- Added
    purpose TEXT,                               -- Added (use TEXT for potentially longer descriptions)
    PRIMARY KEY (recipe_id, food_item_id) -- Can only list a food item once per recipe
);

-- Join table: Links recipes to the equipment needed (Many-to-Many)
CREATE TABLE recipe_equipment (
    recipe_id INT NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
    equipment_id INT NOT NULL REFERENCES equipment(id) ON DELETE RESTRICT, -- Don't delete equipment if used
    PRIMARY KEY (recipe_id, equipment_id)
);

-- Join table: Links recipes to tags (Many-to-Many)
CREATE TABLE recipe_tags (
    recipe_id INT NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
    tag_id INT NOT NULL REFERENCES tags(id) ON DELETE RESTRICT, -- Don't delete tag if used
    PRIMARY KEY (recipe_id, tag_id)
);

-- Optional: Indexes can be added later for performance if needed

