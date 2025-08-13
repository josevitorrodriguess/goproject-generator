package cmd

import (
	"fmt"
	"os"

	"github.com/josevitorrodriguess/goproject-generator/internal/generator"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gpg",
	Short: "Project Generator in Go with the Chi Router",
	Long:  "CLI tool to generate standardized Go project structures using the Chi Router",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func init(){
	rootCmd.AddCommand(createCmd)
}

var (
	port   int
	module string
)

var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new Go project",
	Long:  "Create a new Go project with a standardized structure using the Chi router",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		config := generator.ProjectConfig{
			Name:   projectName,
			Module: module,
			Port:   port,
		}

		if config.Module == "" {
			config.Module = projectName
		}

		fmt.Printf("Creating project: %s\n", projectName)

		if err := generator.CreateProject(config); err != nil {
			fmt.Printf("Error creating project: %v\n", err)
			return
		}

		fmt.Printf("✅ Project %s created successfully!\n", projectName)
		fmt.Printf("📁 Navigate to the directory: cd %s\n", projectName)
		fmt.Printf("🚀 Run: go run cmd/%s/main.go\n", projectName)
	},
}

func init() {
	createCmd.Flags().IntVarP(&port, "port", "p", 3000, "Server port")
	createCmd.Flags().StringVarP(&module, "module", "m", "", "Go module name")
}

