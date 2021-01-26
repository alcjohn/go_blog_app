package html_functions

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func getAssetPaths() map[string]map[string]interface{} {
	jsonFile, err := os.Open("./public/assets/manifest.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

func Vitejs() template.HTML {
	isProduction := os.Getenv("APP_ENV") == "prod"
	if isProduction {
		return ""
	}
	return template.HTML(`<script type="module" src="http://localhost:3000/@vite/client"></script`)
}

func LinkTag(name string) template.HTML {
	isProduction := os.Getenv("APP_ENV") == "prod"
	if !isProduction {
		return ""
	}
	asset := getAssetPaths()
	return template.HTML(fmt.Sprintf(`<link rel="stylesheet" media="screen" href="/public/assets/%s"/>`, asset[name+".css"]["file"]))
}

func ScriptTag(name string) template.HTML {
	isProduction := os.Getenv("APP_ENV") == "prod"
	script := []string{}
	if isProduction {
		asset := getAssetPaths()
		imports := asset[name+".js"]["imports"].([]interface{})
		fmt.Println(imports)
		for _, imp := range imports {
			script = append(script, fmt.Sprintf(`<link rel="modulepreload" href="/public/assets/%s">`, imp))
		}
		script = append(script, fmt.Sprintf(`<script defer type="module" src="/public/assets/%s"></script>`, asset[name+".js"]["file"]))
	} else {
		script = append(script, fmt.Sprintf(`<script defer type="module" src="http://localhost:3000/%s.ts"></script>`, name))
	}
	return template.HTML(strings.Join(script, "\n"))
}
