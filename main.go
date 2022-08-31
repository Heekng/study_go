package main

import (
	"github.com/heekng/study_go/scrapper"
	"github.com/labstack/echo"
	"os"
	"strings"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	//scrapper.Scrape("term")
}
