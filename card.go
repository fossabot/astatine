package astatine

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Card struct {
	ID         string `json:"id"`
	*Notes     `json:"-"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	formatFunc func() error
}

func NewCard() *Card {
	return &Card{
		ID:         newID(),
		Notes:      NewNotes(),
		formatFunc: func() error { return nil },
	}
}

func NewLanguageCard() *Card {
	card := NewCard()
	if len(card.Notes.Notes) == 0 {
		card.Notes.Notes = append(card.Notes.Notes, NewNote())
	}
	note := card.Notes.Notes[0]
	note.Add(NewField("srcLang", ""))
	note.Add(NewField("srcPhrase", ""))
	note.Add(NewField("srcIPA", ""))

	note.Add(NewField("tgtLang", ""))
	note.Add(NewField("tgtPhrase", ""))
	note.Add(NewField("tgtIPA", ""))

	card.formatFunc = func() error {
		var formatStr = `%s      ---      (%s)      ---      %s`
		card.Question = fmt.Sprintf(
			formatStr,
			note.Fields.Get("srcPhrase").GetValue(),
			note.Fields.Get("srcIPA").GetValue(),
			note.Fields.Get("srcLang").GetValue(),
		)
		card.Answer = fmt.Sprintf(
			formatStr,
			note.Fields.Get("tgtLang").GetValue(),
			note.Fields.Get("tgtIPA").GetValue(),
			note.Fields.Get("tgtPhrase").GetValue(),
		)
		return nil
	}
	return card
}

func (c *Card) WithFunc(formatFunc func() error) *Card {
	c.formatFunc = formatFunc
	return c
}

func (c *Card) Apply() {
	if err := c.formatFunc(); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}

func (c *Card) With(notes *Notes) *Card {
	defer c.Apply()
	c.Notes = notes
	return c
}

func (c *Card) Add(note *Note) *Card {
	defer c.Apply()
	c.Notes.Notes = append(c.Notes.Notes, note)
	c.Apply()
	return c
}

func (c *Card) String() string {
	return toString(c)
}

func (c *Card) SetField(key, value string) {
	defer c.Apply()
	note := c.Notes.Notes[0]
	field := note.Fields.Get(key)
	if field == nil {
		return
	}
	field.Value = value

	c.Apply()
}
