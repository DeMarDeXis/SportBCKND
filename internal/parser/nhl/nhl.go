package nhl

import (
	"github.com/DeMarDeXis/VProj/internal/model/parsModel"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"strings"
)

func NHLNameParse(g *geziyor.Geziyor, r *client.Response) {
	count := 0
	r.HTMLDoc.Find("li").Each(func(i int, s *goquery.Selection) {
		if count >= 32 {
			return
		}

		abbr := strings.TrimSpace(strings.Split(s.Text(), "–")[0])
		if len(abbr) <= 3 {
			teamName := s.Find("a").Text()

			teams := parsModel.Team{
				Name: teamName,
				Abbr: abbr,
			}

			g.Exports <- teams
			count++
		}
	})
}
