package templating

import (
	"html/template"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const (
	templateDir = "web/template"
	layoutsDir  = "layouts"
	partialsDir = "partials"
)

func InitTemplates(funcMap template.FuncMap) map[string]*template.Template {
	log.Println("Initializing template sets...")
	sets := make(map[string]*template.Template)

	baseFile := filepath.Join(templateDir, layoutsDir, "base.html")

	// Internal helper to parse a single template set
	parseSet := func(name string, files []string, isRootedAtBase bool) {
		log.Printf("Parsing template set '%s': %v", name, files)

		var rootTemplateName string
		if isRootedAtBase {
			rootTemplateName = filepath.Base(baseFile) // e.g., "base.html"
		} else {
			if len(files) == 0 {
				log.Fatalf("No files provided for template set '%s'", name)
				return
			}
			rootTemplateName = filepath.Base(files[0])
		}

		tmpl := template.Must(template.New(rootTemplateName).
			Funcs(funcMap).
			ParseFiles(files...))

		sets[name] = tmpl
		log.Printf("Stored template set: '%s'", name)
	}

	// 1. Index Page (uses base.html)
	parseSet("index", []string{
		baseFile,
		filepath.Join(templateDir, "index.html"),
	}, true) // true because it's rooted at base.html

	// 2. Recipe Page (uses base.html and includes partials)
	parseSet("recipe", []string{
		baseFile,
		filepath.Join(templateDir, "recipe_page.html"),
		filepath.Join(templateDir, partialsDir, "ingredients.html"),
		filepath.Join(templateDir, partialsDir, "methods.html"),
		filepath.Join(templateDir, partialsDir, "select-recipe-steps.html"),
		filepath.Join(templateDir, partialsDir, "zen_mode_content.html"),
		filepath.Join(templateDir, partialsDir, "tabs.html"),
		filepath.Join(templateDir, partialsDir, "view_options_panel.html"),
		filepath.Join(templateDir, partialsDir, "toggle_switch.html"),
		filepath.Join(templateDir, partialsDir, "full_look_content.html"),
		filepath.Join(templateDir, partialsDir, "banner_description.html"),
		filepath.Join(templateDir, partialsDir, "select_steps_section.html"),
		filepath.Join(templateDir, partialsDir, "ingredients_section.html"),
		filepath.Join(templateDir, partialsDir, "methods_section.html"),
		filepath.Join(templateDir, partialsDir, "notes_section.html"),
		filepath.Join(templateDir, partialsDir, "zen_mode_content_wrapper.html"),
	}, true) // true, rooted at base.html

	// 3. Standalone Partials (for HTMX swaps, etc.)
	// These are not rooted at base.html; their root is their own file name.
	parseSet("ingredient-list", []string{
		filepath.Join(templateDir, partialsDir, "ingredients.html"),
	}, false) // false, root is "ingredients.html"

	parseSet("methods-list", []string{
		filepath.Join(templateDir, partialsDir, "methods.html"),
	}, false) // false, root is "methods.html"

	parseSet("select-recipe-steps", []string{
		filepath.Join(templateDir, partialsDir, "select-recipe-steps.html"),
	}, false) // false, root is "select-recipe-steps.html"

	parseSet("zen-mode-content", []string{
		filepath.Join(templateDir, partialsDir, "zen_mode_content.html"),
	}, false) // false, root is "zen_mode_content.html"

	log.Println("All template sets initialized.")
	return sets
}
