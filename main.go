package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	templ := template.Must(template.ParseGlob("template/*.html"))

	v := map[string]interface{}{
		"Header": "Hi there",
		"Footer": "Welcome to Nepu",
	}
	fmt.Println("\n===foo===")
	templ.ExecuteTemplate(os.Stdout, "foo.html", v)

	fmt.Println("\n===bar===")
	templ.ExecuteTemplate(os.Stdout, "bar.html", v)
}
