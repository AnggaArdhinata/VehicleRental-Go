package main

import (
	"log"
	"os"

	"github.com/AnggaArdhinata/backend-go/src/commands"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	if err := commands.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

}
