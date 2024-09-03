package services

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"v3/api"
	"v3/config"

	"golang.org/x/oauth2"
)

type TeamCompareService struct {
	client *http.Client
}

func NewTeamCompareService() *TeamCompareService {
	return &TeamCompareService{}
}

func (s *TeamCompareService) GetWinningMatchupsLeague(token any) ([]api.Team, error) {

	return nil, nil

}

type WinningMatchups struct {
	WinningMatchupCount int
	WinningMatchupTeams map[string]int
}

type Category struct {
	CategoryId       int
	StatisticTeamMap map[string]float64
	Average          float64
}

func (s *TeamCompareService) GetTeams(token any) []string {
	oauthClient := config.OAuthConfig.Client(context.Background(), token.(*oauth2.Token))
	url := fmt.Sprintf("https://fantasysports.yahooapis.com/fantasy/v2/league/%s/teams/", os.Getenv("LEAGUE_KEY"))
	r, err := oauthClient.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(r.Body)

	var l api.FantasyContent

	err = xml.Unmarshal(bytes, &l)
	var teamres []string
	for _, team := range l.League.Teams {
		teamres = append(teamres, team.Name)
	}

	return teamres
}

func (s *TeamCompareService) GetCategoryLeaders(token any) []Category {

	oauthClient := config.OAuthConfig.Client(context.Background(), token.(*oauth2.Token))
	url := fmt.Sprintf("https://fantasysports.yahooapis.com/fantasy/v2/league/%s/teams/stats;type=week;week=20", os.Getenv("LEAGUE_KEY"))
	r, err := oauthClient.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(r.Body)

	var l api.FantasyContent
	err = xml.Unmarshal(bytes, &l)

	statLength := len(l.League.Teams[0].TeamStatistics[0].Stats)
	var res []Category
	for index := range statLength {
		StatisticMap := make(map[string]float64)
		var CatId int
		average := 0.0
		for _, team := range l.League.Teams {
			floatValue := api.ConvertFractionToDecimal(team.TeamStatistics[0].Stats[index].Value)
			average += floatValue
			StatisticMap[team.Name] = floatValue
			CatId = team.TeamStatistics[0].Stats[index].StatId
		}
		average /= float64(statLength)
		c := Category{CategoryId: CatId, StatisticTeamMap: StatisticMap, Average: average}
		res = append(res, c)
	}

	return res
}

func (s *TeamCompareService) GetData(token any) map[string]WinningMatchups {

	oauthClient := config.OAuthConfig.Client(context.Background(), token.(*oauth2.Token))
	url := fmt.Sprintf("https://fantasysports.yahooapis.com/fantasy/v2/league/%s/teams/stats;type=week;week=20", os.Getenv("LEAGUE_KEY"))
	r, err := oauthClient.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(r.Body)

	var l api.FantasyContent
	err = xml.Unmarshal(bytes, &l)
	m := make(map[string]WinningMatchups)
	statLength := len(l.League.Teams[0].TeamStatistics[0].Stats)
	for _, team := range l.League.Teams {
		winningMatchupTeam := make(map[string]int)
		winningMatchupCount := 0

		for _, teamcomp := range l.League.Teams {
			winningStats := 0
			if teamcomp.Name == team.Name {
				continue
			}

			for index := range statLength {

				stat1 := team.TeamStatistics[0].Stats[index].Value
				stat2 := teamcomp.TeamStatistics[0].Stats[index].Value
				var floatVal1 float64
				var floatVal2 float64

				floatVal1 = api.ConvertFractionToDecimal(stat1)
				floatVal2 = api.ConvertFractionToDecimal(stat2)

				log.Println(team.TeamStatistics[0].Stats[index].Value, teamcomp.TeamStatistics[0].Stats[index].Value)
				if floatVal1 > floatVal2 {
					log.Printf("Team: %s beats Team: %s since %f > %f", team.Name, teamcomp.Name, floatVal1, floatVal2)
					winningStats += 1
				}

			}
			if winningStats >= 5 {
				winningMatchupCount += 1
				winningMatchupTeam[teamcomp.Name] = winningStats
			}

		}
		mu := WinningMatchups{WinningMatchupCount: winningMatchupCount, WinningMatchupTeams: winningMatchupTeam}
		m[team.Name] = mu
	}

	log.Println(m)

	if err != nil {
		log.Fatal(err)
	}

	return m
}
