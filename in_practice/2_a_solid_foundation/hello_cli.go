// package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/urfave/cli"
// )

// func main() {
// 	//create a new application
// 	app := cli.NewApp()
// 	app.Name = "hello_cli"
// 	app.Usage = "Print hello world"
// 	//setup global flag
// 	app.Flags = []cli.Flag{
// 		cli.StringFlag{
// 			Name: "name, n",

// 			Value: "World",
// 			Usage: "Who to say hello to.",
// 		},
// 	}
// 	//define the action to run
// 	app.Action = func(c *cli.Context) error {
// 		name := c.GlobalString("name")
// 		fmt.Printf("Hello %s!\n", name)
// 		return nil
// 	}
// 	//run the application
// 	app.Run(os.Args)
// }
