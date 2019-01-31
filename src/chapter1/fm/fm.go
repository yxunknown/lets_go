package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "fm for movies."
	app.Usage = "find movies and download it."
	app.Version = "1.1.0"
	app.Author = "Hercats <hercats@qq.com>"
	app.Commands = []cli.Command{
		{
			Name:      "search",
			ShortName: "s",
			Usage:     "search movie resources with names.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "name of movie.",
					Value: "",
				},
			},
			Action: func(c *cli.Context) error {
				name := c.String("name")
				search(name)
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func search(name string) {
	fmt.Println("searching movie with name:", name)
	respone, err := http.Get("http://www.btbtdy.net/search/江湖儿女.html")
	if err != nil {
		log.Fatal(err)
	} else {
		for {
			buf := make([]byte, 1024)
			_, err = respone.Body.Read(buf)
			if err == nil {
				fmt.Printf("%q", buf)
			}
		}
	}
}
