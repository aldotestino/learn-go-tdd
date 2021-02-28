package main

import (
	poker "app"
	"fmt"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Lets's play poker")
	fmt.Println("Type {name} wins to record a win")

	cli.PlayPoker()
}
