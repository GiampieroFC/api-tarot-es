package main

import (
	"fmt"
	"os"

	"github.com/GiampieroFC/db-tarot-v1/server"
)

func main() {
	fmt.Println("Hola, Mundo.")

	// cards := scraper.NewCards().
	// cards.CreateTable(db.GetPool())
	// cards.FillTabla()

	// arr := scraper.Start()

	// for _, v := range arr {
	// 	cards.FillTabla(db.GetPool(), &v)
	// }

	// fmt.Println(string(bytes))

	port := os.Getenv("PORT")

	server.ToServe(port)

	// fmt.Println(cards)
}
