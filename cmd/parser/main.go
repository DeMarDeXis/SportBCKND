package main

import (
	"github.com/DeMarDeXis/VProj/internal/lib/customjsonexp"
	"github.com/DeMarDeXis/VProj/internal/parser/nhl"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/export"
	"log"
	"os"
	"strings"
)

const (
	nhlNameAbbr = "https://en.wikipedia.org/wiki/Wikipedia:WikiProject_Ice_Hockey/NHL_team_abbreviations"
)

func main() {
	log.SetOutput(os.Stdout)
	exporter, err := customjsonexp.NewCustomJSONExporter("./jsondata/NHL.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start")

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{nhlNameAbbr},
		ParseFunc: nhl.NHLNameParse,
		Exporters: []export.Exporter{exporter},
	}).Start()

}

func logValue(value string, exists ...bool) string {
	if len(exists) > 0 && !exists[0] {
		return "empty"
	}
	if value == "" {
		return "empty"
	}
	return value
}

func printElementHTML(s *goquery.Selection) {
	html, err := s.Html()
	if err != nil {
		log.Printf("Error getting HTML: %v", err)
		return
	}
	log.Println("Element HTML:")
	log.Println(html)
}

func saveHTMLToFile(doc *goquery.Document, filename string) error {
	html, err := doc.Html()
	if err != nil {
		return err
	}

	// Format the HTML for better readability
	formatted := strings.Replace(html, "><", ">\n<", -1)

	return os.WriteFile(filename, []byte(formatted), 0644)
}
