package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)

type TemplateProfile struct {
	Name string `json:"name"`
	URL string `json:"url"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	e := echo.New()
	e.POST("/templateProfile", executeTemplate)
	e.Logger.Fatal(e.Start(":8000"))

	fmt.Println("Hello")

}

func runCollection(fileName string) {
	out, err := exec.Command("newman", "run", fileName).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func executeTemplate(c echo.Context) error {
	templateProfile := new(TemplateProfile)

	if err := c.Bind(templateProfile); err != nil {
		return err
	}

	populateTemplate(*templateProfile, "user.json")

	return c.JSON(http.StatusOK, templateProfile)
}

func populateTemplate(profile TemplateProfile, fileName string) {
	tpl, err := template.ParseFiles(fileName)

	if err != nil {
		//FIXME
		fmt.Println(err)
		log.Fatalln(err)
	}

	collectionFile, err := os.Create(fileName)
	tpl.Execute(collectionFile, profile)

	runCollection(fileName)

}
