package generator

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
)

type ProjectConfig struct {
	Name   string
	Module string
	Port   int
}

func CreateProject(config ProjectConfig) error {
	if err := os.MkdirAll(config.Name, 0755); err != nil {
		return fmt.Errorf("Failed to create project directory: %v", err)
	}

	dirs := []string{
		filepath.Join(config.Name, "cmd", config.Name),
		filepath.Join(config.Name, "internal", "api"),
		filepath.Join(config.Name, "internal", "jsonutils"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("Failed to create project directory: %v", err)
		}
	}

	if err := createGoMod(config); err != nil {
		return err
	}

	if err := createFromTemplates(config); err != nil {
		return err
	}

	if err := installDependencies(config.Name); err != nil {
		return err
	}

	return nil
}

func createGoMod(config ProjectConfig) error {
	goModPath := filepath.Join(config.Name, "go.mod")

	tmpl := `module {{.Module}}

	go 1.24.3
`

	return executeTemplate(tmpl, goModPath, config)
}

func createFromTemplates(config ProjectConfig) error {
	templates := map[string]string{
		filepath.Join(config.Name, "cmd", config.Name, "main.go"):            mainGoTemplate,
		filepath.Join(config.Name, "internal", "api", "api.go"):              apiGoTemplate,
		filepath.Join(config.Name, "internal", "api", "router.go"):           routerGoTemplate,
		filepath.Join(config.Name, "internal", "jsonutils", "json_utils.go"): jsonutilsGoTemplate,
	}

	for path, tmpl := range templates {
		if err := executeTemplate(tmpl, path, config); err != nil {
			return err
		}
	}

	return nil
}

func installDependencies(projectName string) error {
	cmd := exec.Command("go", "get", "github.com/go-chi/chi/v5",
		"go", "get", "github.com/go-chi/chi/",
		"go", "get", "github.com/go-chi/cors",
		"go", "get", "github.com/joho/godotenv")

	cmd.Dir = projectName

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erro ao instalar dependências: %v", err)
	}

	return nil
}
func executeTemplate(tmplStr, outputPath string, config ProjectConfig) error {
	tmpl, err := template.New("template").Parse(tmplStr)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", outputPath, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, config); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	return nil
}
