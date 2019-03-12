package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	gitti := cli.NewApp()
	gitti.Action = terminalAction

	err := gitti.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func terminalAction(c *cli.Context) error {
	subject := c.Args().Get(0)
	switch subject {
	case "repo":
		log.Println("selected repo")
	case "issue":
		log.Println("selected issue")
	case "user":
		log.Println("selected user")
	default:
		log.Println("wrong selection")
	}

	return nil
}
