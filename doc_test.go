package mtg_test

import (
	"github.com/MagicTheGathering/mtg-sdk-go"
	"log"
)

func ExampleQuery_all() {
	log.Println("Fetching all cards with CMC >= 16")
	cards, err := mtg.NewQuery().Where(mtg.CardCMC, "gte16").All()
	if err != nil {
		log.Panic(err)
	}
	for _, card := range cards {
		log.Println(card)
	}
}

func ExampleQuery_page() {
	log.Println("fetch first page (100 cards in total)")

	cards, totalCards, err := mtg.NewQuery().Where(mtg.CardColors, "green|red").Page(1)
	if err != nil {
		log.Panic(err)
	}

	log.Println("There are", totalCards, "green or red cards")
	for _, card := range cards {
		log.Println(card)
	}
}

func ExampleQuery_pageS() {
	log.Println("Fetch Page 2 with a page size of 5")

	cards, totalCards, err := mtg.NewQuery().Where(mtg.CardColors, "white").PageS(2, 5)
	if err != nil {
		log.Panic(err)
	}

	log.Println("There are", totalCards, "white cards")
	for _, card := range cards {
		log.Println(card)
	}
}

func ExampleId_fetch() {
	fetchCardID := func(cID mtg.Id) {
		// cID could either be a CardId or a MultiverseId
		card, err := cID.Fetch()
		if err != nil {
			log.Panic(err)
		}
		log.Println(card)
	}

	log.Println("Fetching one Card with a given multiverseId")
	fetchCardID(mtg.MultiverseId(73947))

	log.Println("Fetching one Card with a given cardId")
	fetchCardID(mtg.CardId("9d91ef4896ab4c1a5611d4d06971fc8026dd2f3f"))
}

func ExampleQuery_random() {
	// Fetch 2 random red rare cards
	cards, err := mtg.NewQuery().Where(mtg.CardRarity, "rare").Where(mtg.CardColors, "red").Random(2)
	if err != nil {
		log.Panic(err)
	}
	for _, c := range cards {
		log.Println(c)
	}
}

func ExampleSetQuery_all() {
	sets, err := mtg.NewSetQuery().Where(mtg.SetName, "khans").All()
	if err != nil {
		log.Panic(err)
	}

	for _, set := range sets {
		log.Println(set)
	}
}

func ExampleSetCode_fetch() {
	set, err := mtg.SetCode("KTK").Fetch()
	if err != nil {
		log.Panic(err)
	}
	log.Println(set)
}

func ExampleSetCode_generateBooster() {
	cards, err := mtg.SetCode("KTK").GenerateBooster()
	if err != nil {
		log.Panic(err)
	}
	for _, c := range cards {
		log.Println(c)
	}
}
