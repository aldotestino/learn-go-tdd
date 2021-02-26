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

	fmt.Println("Lets's play poker")
	fmt.Println("Type {name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
