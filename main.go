package main

import (
  "fmt"
  "log"
  "os"
  "sort"
  "github.com/jizh6311/nba_statistics_cli/handlers"
  "github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Name:  "nba_statistics_cli",
    Usage: "This CLI helps you fetch updated NBA statistics",
  }

  app.Commands = []*cli.Command{
    {
      Name:  "score-board",
      Usage: "Get today's scoreboard",
      Action: func(c *cli.Context) error {
      fmt.Printf(handlers.GetAllScores("20191220"))
        return nil
      },
    },
    {
      Name:  "standings",
      Usage: "Get current standings",
      Action: func(c *cli.Context) error {
      fmt.Printf(handlers.GetStandings("20191220"))
        return nil
      },
    },
  }

  sort.Sort(cli.CommandsByName(app.Commands))

  err := app.Run(os.Args)
    if err != nil {
      log.Fatal(err)
    }
}
