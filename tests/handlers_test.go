package handlersTest

import (
  "testing"

  "github.com/jizh6311/nba_statistics_cli/handlers"
  "github.com/stretchr/testify/assert"
)

func TestGetAllScores(t *testing.T) {
  allScores := handlers.GetAllScores("20191220")
  assert.Contains(t, allScores, "20191220", "The score results should contain the date")
}

func TestGetStandings(t *testing.T) {
  standings := handlers.GetStandings("20191220")
  assert.Contains(t, standings, "20191220", "The standing results should contain the date")
}
