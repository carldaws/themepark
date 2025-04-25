package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed themes/*
var themesFS embed.FS

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: themepark <command> [arguments]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "use":
		if len(os.Args) < 3 {
			fmt.Println("Usage: themepark use <theme>")
			os.Exit(1)
		}
		themeName := os.Args[2]
		err := useTheme(themeName)
		if err != nil {
			fmt.Println("Theme", themeName, "not found! Try `themepark list`?")
			os.Exit(1)
		}
	case "list":
		err := listThemes()
		if err != nil {
			fmt.Println("Error listing themes:", err)
			os.Exit(1)
		}
	case "where":
		if len(os.Args) < 3 {
			fmt.Println("Usage: themepark where <target>")
			os.Exit(1)
		}
		target := os.Args[2]
		err := whereTarget(target)
		if err != nil {
			fmt.Println("Error finding target:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}

func useTheme(themeName string) error {
	themeData, err := loadTheme(themeName)
	if err != nil {
		return err
	}

	templateContent, err := templatesFS.ReadFile("templates/ghostty.tmpl")
	if err != nil {
		return err
	}

	tmpl, err := template.New("ghostty").Parse(string(templateContent))
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	outputDir := filepath.Join(homeDir, ".themepark")
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return err
	}

	outputPath := filepath.Join(outputDir, "ghostty.conf")

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = tmpl.Execute(outFile, themeData)
	if err != nil {
		return err
	}

	fmt.Println("✅ Theme switched!")

	return nil
}

func loadTheme(themeName string) (map[string]string, error) {
	themeData := make(map[string]string)

	file, err := themesFS.ReadFile("themes/" + themeName + ".json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &themeData)
	if err != nil {
		return nil, err
	}

	return themeData, nil
}

func listThemes() error {
	entries, err := fs.ReadDir(themesFS, "themes")
	if err != nil {
		return err
	}

	fmt.Println("Available themes:")
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if strings.HasSuffix(name, ".json") {
			name = strings.TrimSuffix(name, ".json")
		}
		fmt.Println("-", name)
	}

	return nil
}

func whereTarget(target string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	switch target {
	case "ghostty":
		outputPath := filepath.Join(homeDir, ".themepark", "ghostty.conf")
		fmt.Println(outputPath)
	default:
		fmt.Println("Unknown target:", target)
		os.Exit(1)
	}

	return nil
}
