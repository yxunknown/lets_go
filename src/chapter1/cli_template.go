package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	//err := cli.NewApp().Run(os.Args)
	//if err != nil {
	//	log.Fatal(err)
	//}
	app := cli.NewApp()
	app.Name = "Boom"
	app.Usage = "Boom after five tick"
	app.Version = "1.2.0"
	app.Author = "hercats"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "number,n",
			Value: "5",
			Usage: "times for tick",
		},
	}
	app.Action = func(c *cli.Context) error {
		number := c.GlobalInt("number")
		for i := 0; i < number; i++ {
			fmt.Println("tick")
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Boom")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
