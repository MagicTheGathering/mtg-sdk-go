package main

import (
	"github.com/boombuler/magicthegathering.io"
	"log"
)

func exampleFetchAllCards() {
	log.Println("Fetching all cards with CMC >= 16")
	cards, errors := mtg.NewQuery().CMC("gte16").All()

	for card := range cards {
		log.Println(card)
	}

	if err := <-errors; err != nil {
		log.Panic(err)
	}
}

func exampleFetchCardPage() {
	log.Println("fetch first page (100 cards in total)")

	cards, totalCards, err := mtg.NewQuery().Colors("green|red").Page(1)
	if err != nil {
		log.Panic(err)
	}

	log.Println("There are", totalCards, "green or red cards")
	for _, card := range cards {
		log.Println(card)
	}
}

func exampleFetchCardPageWithPageSize() {
	log.Println("Fetch Page 2 with a page size of 5")

	cards, totalCards, err := mtg.NewQuery().Colors("white").PageS(2, 5)
	if err != nil {
		log.Panic(err)
	}

	log.Println("There are", totalCards, "white cards")
	for _, card := range cards {
		log.Println(card)
	}
}

func fetchCardID(cID mtg.Id) {
	// cID could either be a CardId or a MultiverseId
	card, err := cID.Fetch()
	if err != nil {
		log.Panic(err)
	}
	log.Println(card)
}

func exampleFetchCardByIDs() {
	log.Println("Fetching one Card with a given multiverseId")
	fetchCardID(mtg.MultiverseId(73947))

	log.Println("Fetching one Card with a given cardId")
	fetchCardID(mtg.CardId("9d91ef4896ab4c1a5611d4d06971fc8026dd2f3f"))
}

func exampleFetchRandomCard() {
	// Fetch 2 random red rare cards
	cards, err := mtg.NewQuery().Rarity("rare").Colors("red").Random(2)
	if err != nil {
		log.Panic(err)
	}
	for _, c := range cards {
		log.Println(c)
	}
}

func main() {
	exampleFetchRandomCard()
	exampleFetchAllCards()
	exampleFetchCardByIDs()
	exampleFetchCardPageWithPageSize()
	exampleFetchCardPage()
}
