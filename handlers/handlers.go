package handlers

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "strings"

  "github.com/thedevsaddam/gojsonq"
)

func GetHttpResponse(url string) string {
  resp, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()
  responseData, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatal(err)
  }

  responseString := string(responseData)

  return responseString
}

func GetStandingsFromJSON(json string) []string {
  eastStandingsList := gojsonq.New().FromString(json).From("league.standard.conference.east")
  var eastTeamList []string
  for i := 1; i <= eastStandingsList.Count(); i++ {
    teamMap1 := eastStandingsList.Nth(i).(map[string]interface{})
    teamMap2 := teamMap1["teamSitesOnly"].(map[string]interface{})
    eastTeamList = append(eastTeamList, teamMap2["teamCode"].(string))
  }

  return eastTeamList
}

func GetAllScores(date string) string {
  return fmt.Sprintf("The scoreborad on %s is coming. Thanks for your patience.", date)
}

func GetStandings(date string) string {
  standingsResponse := GetHttpResponse("https://data.nba.net/10s/prod/v1/current/standings_conference.json")
  standingsList := GetStandingsFromJSON(standingsResponse)
  return fmt.Sprintf(strings.Join(standingsList, ","))
}
