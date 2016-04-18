package main

import (
	"fmt"
	"io"
	"os"

	"bufio"

	"github.com/christer79/cmdLine/dateformatter"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = os.Getenv("VERSION")
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "date-format, D",
			Value: "2006-01-02 15:04:05",
			Usage: "Specify a date format for use",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "reformat",
			Usage: "Find date string in text file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input-date-format, d",
					Value: "2006-Jan-02 15:04:05",
					Usage: "Specify a date format in original file",
				},
			},
			Action: func(c *cli.Context) {
				var str string
				var err error
				reader := bufio.NewReader(os.Stdin)

				for {
					str, err = reader.ReadString('\n')
					if err == io.EOF {
						break
					}
					date := dateformatter.GetDate(str, c.String("input-date-format"))
					fmt.Printf("%s %s", date.Format(c.GlobalString("date-format")), str)
				}
			},
		},
	}

	app.Run(os.Args)
}
