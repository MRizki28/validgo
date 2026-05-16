package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("usage: validgo make:request User/Test")
		return
	}

	command := os.Args[1]
	name := os.Args[2]

	switch command {

	case "make:request":
		makeRequest(name)

	default:
		fmt.Println("command not found")
	}
}

func makeRequest(name string) {

	basePath := getRequestPath()

	name = strings.ReplaceAll(name, "\\", "/")

	parts := strings.Split(name, "/")

	className := parts[len(parts)-1]

	fileName := toSnakeCase(className) + "_request.go"

	var folders []string

	if len(parts) > 1 {

		for _, part := range parts[:len(parts)-1] {
			folders = append(folders, toSnakeCase(part))
		}
	}

	fullDir := filepath.Join(
		append([]string{basePath}, folders...)...,
	)

	fullPath := filepath.Join(fullDir, fileName)

	packageName := "requests"

	if len(folders) > 0 {
		packageName = folders[len(folders)-1]
	}

	content := fmt.Sprintf(`package %s

type %sRequest struct {

}
`, packageName, className)

	err := os.MkdirAll(fullDir, os.ModePerm)

	if err != nil {
		fmt.Println("failed create directory")
		return
	}

	_, err = os.Stat(fullPath)

	if err == nil {
		fmt.Println("request already exists")
		return
	}

	err = os.WriteFile(
		fullPath,
		[]byte(content),
		0644,
	)

	if err != nil {
		fmt.Println("failed create request")
		return
	}

	fmt.Println("request created:", fullPath)
}

func getRequestPath() string {

	_, err := os.Stat("app")

	if err == nil {
		return "app/requests"
	}

	return "requests"
}

func toSnakeCase(str string) string {

	var result strings.Builder

	for i, r := range str {

		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}

		result.WriteRune(r)
	}

	return strings.ToLower(result.String())
}