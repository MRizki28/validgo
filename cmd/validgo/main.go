package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func MakeRequest(name string) {

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

	content := fmt.Sprintf(`package requests

type %sRequest struct {

}
`, className)

	err := os.MkdirAll(fullDir, os.ModePerm)

	if err != nil {
		fmt.Println("failed create directory")
		return
	}

	err = os.WriteFile(
		fullPath,
		[]byte(content),
		0644,
	)

	if err != nil {

		os.RemoveAll(fullDir)

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