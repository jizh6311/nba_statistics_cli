package handlersTest

import (
  "io/ioutil"
  "testing"

  "github.com/jarcoal/httpmock"
  "github.com/jizh6311/nba_statistics_cli/handlers"
  "github.com/stretchr/testify/assert"
)

func TestGetHttpResponse(t *testing.T) {
  httpmock.Activate()
  defer httpmock.DeactivateAndReset()

  httpmock.RegisterResponder("GET", `=~^https://data.nba.net/10s/prod/v1`,
    httpmock.NewStringResponder(200, `{"id": 1, "name": "This is for testing"}`))

  responseString := handlers.GetHttpResponse("https://data.nba.net/10s/prod/v1/20191222/scoreboard.json")

  assert.Contains(t, responseString, "This is for testing", "The response string should match")
  assert.Contains(t, responseString, "\"id\": 1", "The response string should match")
}

func TestGetStandingsFromJSON(t *testing.T) {
  data, err := ioutil.ReadFile("./data/standings.json")
  if err != nil {
    t.Errorf("Cannot read file ./data/standings.json")
  }
  standings := handlers.GetStandingsFromJSON(string(data))

  assert.ElementsMatch(t, standings[0], [3]string{"bucks", "10", "3"}, "The standings do not match bucks")
  assert.ElementsMatch(t, standings[1], [3]string{"lakers", "11", "2"}, "The standings do not match lakers")
}

func TestGetAllScores(t *testing.T) {
  allScores := handlers.GetAllScores("20191220")
  assert.Contains(t, allScores, "20191220", "The score results should contain the date")
}

func TestGetStandings(t *testing.T) {
  // TODO: Mock data for standings
}
