package handlers

import "fmt"

func GetAllScores(date string) string {
  return fmt.Sprintf("The scoreborad on %s is coming. Thanks for your patience.", date)
}

func GetStandings(date string) string {
  return fmt.Sprintf("The standings on %s is coming. Thanks for your patience.", date)
}
