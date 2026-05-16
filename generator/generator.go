package generator

import (
	"fmt"
	"os"
	"strings"
)

func MakeRequest(name string) {

	fileName := toSnakeCase(name) + "_request.go"

	content := fmt.Sprintf(`package requests

type %sRequest struct {

}
`, name)

	requestPath := getRequestPath()

	os.MkdirAll(requestPath, os.ModePerm)

	err := os.WriteFile(
		requestPath+"/"+fileName,
		[]byte(content),
		0644,
	)

	if err != nil {
		fmt.Println("failed create request")
		return
	}

	fmt.Println("request created")
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
