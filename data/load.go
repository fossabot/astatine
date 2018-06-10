package data

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"log"
	"strings"

	"github.com/johnmcdnl/astatine"
)

const (
	Countries = `./data/countries.csv`
)

func Load(file string) *astatine.Deck {
	deck := astatine.NewDeck()
	data := readCSV(file)
	for _, row := range data {
		question := strings.Trim(row[1], " ")
		answer := strings.Trim(row[0], " ")

		card := astatine.NewLanguageCard()
		card.SetField("srcLang", "ru")
		card.SetField("tgtLang", "en")

		card.SetField("srcPhrase", question)
		card.SetField("tgtPhrase", answer)

		card.SetField("srcIPA", "")
		card.SetField("tgtIPA", "")

		deck.Add(card)

	}
	deck.Cards.Shuffle()
	return deck
}

func readCSV(file string) [][]string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(bytes.NewReader(b))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}
