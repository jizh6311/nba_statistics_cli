package handlers

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"

  "github.com/olekukonko/tablewriter"
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

func GetStandingsFromJSON(json string) [][]string {
  eastStandingsList := gojsonq.New().FromString(json).From("league.standard.conference.east")
  var eastTeamList [][]string
  for i := 1; i <= eastStandingsList.Count(); i++ {
    var teamArr []string
    teamMap1 := eastStandingsList.Nth(i).(map[string]interface{})
    teamSitesOnly := teamMap1["teamSitesOnly"].(map[string]interface{})

    teamArr = append(
      teamArr,
      teamSitesOnly["teamCode"].(string),
      teamMap1["win"].(string),
      teamMap1["loss"].(string),
    )

    eastTeamList = append(eastTeamList, teamArr)
  }

  return eastTeamList
}

func GetAllScores(date string) string {
  return fmt.Sprintf("The scoreborad on %s is coming. Thanks for your patience.", date)
}

func GetStandings(date string) *tablewriter.Table {
  standingsResponse := GetHttpResponse("https://data.nba.net/10s/prod/v1/current/standings_conference.json")
  standingsList := GetStandingsFromJSON(standingsResponse)

  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Team", "Win", "Loss"})

  for _, teamInfo := range standingsList {
    table.Append(teamInfo)
  }

  return table
}
