package mtg

import (
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	queryUrl = "https://api.magicthegathering.io/v1/"
)

var (
	linkRE = regexp.MustCompile(`<(.*)>; rel="(.*)"`)
)

type Query interface {
	// The card name. For split, double-faced and flip cards, just the name of one side of the card. Basically each ‘sub-card’ has its own record.
	Name(qry string) Query
	// The card layout. Possible values: normal, split, flip, double-faced, token, plane, scheme, phenomenon, leveler, vanguard
	Layout(qry string) Query
	// Converted mana cost. Always a number.
	CMC(qry string) Query
	// The card colors. Usually this is derived from the casting cost, but some cards are special (like the back of dual sided cards and Ghostfire).
	Colors(qry string) Query
	// The card colors by color code. [“Red”, “Blue”] becomes [“R”, “U”]
	ColorIdentity(qry string) Query
	// The card type. This is the type you would see on the card if printed today. Note: The dash is a UTF8 'long dash’ as per the MTG rules
	Type(qry string) Query
	// The supertypes of the card. These appear to the far left of the card type. Example values: Basic, Legendary, Snow, World, Ongoing
	Supertypes(qry string) Query
	// The types of the card. These appear to the left of the dash in a card type. Example values: Instant, Sorcery, Artifact, Creature, Enchantment, Land, Planeswalker
	Types(qry string) Query
	// The subtypes of the card. These appear to the right of the dash in a card type. Usually each word is its own subtype. Example values: Trap, Arcane, Equipment, Aura, Human, Rat, Squirrel, etc.
	Subtypes(qry string) Query
	// The rarity of the card. Examples: Common, Uncommon, Rare, Mythic Rare, Special, Basic Land
	Rarity(qry string) Query
	// The set the card belongs to (set code).
	Set(qry string) Query
	// The set the card belongs to.
	SetName(qry string) Query
	// The oracle text of the card. May contain mana symbols and other symbols.
	Text(qry string) Query
	// The flavor text of the card.
	Flavor(qry string) Query
	// The artist of the card. This may not match what is on the card as MTGJSON corrects many card misprints.
	Artist(qry string) Query
	// The card number. This is printed at the bottom-center of the card in small text. This is a string, not an integer, because some cards have letters in their numbers.
	Number(qry string) Query
	// The power of the card. This is only present for creatures. This is a string, not an integer, because some cards have powers like: “1+*”
	Power(qry string) Query
	// The toughness of the card. This is only present for creatures. This is a string, not an integer, because some cards have toughness like: “1+*”
	Toughness(qry string) Query
	// The loyalty of the card. This is only present for planeswalkers.
	Loyalty(qry string) Query
	// The name of a card in a foreign language it was printed in
	ForeignName(qry string) Query

	// The language the card is printed in. Use this parameter when searching by foreignName
	Language(qry string) Query
	// The game format, such as Commander, Standard, Legacy, etc. (when used, legality defaults to Legal unless supplied)
	GameFormat(qry string) Query
	// The legality of the card for a given format, such as Legal, Banned or Restricted.
	Legality(qry string) Query

	// Creates a copy of this query
	Copy() Query

	// Fetches all cards matching the current query
	All() ([]*Card, error)

	// Fetches the given page of cards.
	Page(pageNum int) (cards []*Card, totalCardCount int, err error)
	// Fetches one page of cards with a given page size
	PageS(pageNum int, pageSize int) (cards []*Card, totalCardCount int, err error)
	// Fetches some random cards
	Random(count int) ([]*Card, error)
}

func NewQuery() Query {
	return make(query)
}

type query map[string]string

func fetchCards(url string) ([]*Card, http.Header, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	bdy := resp.Body
	defer bdy.Close()
	cards, err := decodeCards(bdy)
	if err != nil {
		return nil, nil, err
	}
	return cards, resp.Header, nil
}

func (q query) All() ([]*Card, error) {
	var allCards []*Card

	queryVals := make(url.Values)
	for k, v := range q {
		queryVals.Set(k, v)
	}
	nextUrl := queryUrl + "cards?" + queryVals.Encode()
	for nextUrl != "" {
		cards, header, err := fetchCards(nextUrl)
		if err != nil {
			return nil, err
		}

		nextUrl = ""

		if linkH, ok := header["Link"]; ok {
			parts := strings.Split(linkH[0], ",")
			for _, link := range parts {
				match := linkRE.FindStringSubmatch(link)
				if match != nil {
					if match[2] == "next" {
						nextUrl = match[1]
					}
				}
			}
		}

		allCards = append(allCards, cards...)
	}
	return allCards, nil
}

func (q query) Page(pageNum int) (cards []*Card, totalCardCount int, err error) {
	return q.PageS(pageNum, 100)
}

func (q query) PageS(pageNum int, pageSize int) (cards []*Card, totalCardCount int, err error) {
	cards = nil
	totalCardCount = 0
	err = nil

	queryVals := make(url.Values)
	for k, v := range q {
		queryVals.Set(k, v)
	}

	queryVals.Set("page", strconv.Itoa(pageNum))
	queryVals.Set("pageSize", strconv.Itoa(pageSize))

	url := queryUrl + "cards?" + queryVals.Encode()
	cards, header, err := fetchCards(url)
	if err != nil {
		return nil, 0, err
	}
	totalCardCount = len(cards)
	if totals, ok := header["Total-Count"]; ok && len(totals) > 0 {
		if totalCardCount, err = strconv.Atoi(totals[0]); err != nil {
			return nil, 0, err
		}
	}
	return cards, totalCardCount, nil
}

func (q query) Random(count int) ([]*Card, error) {
	queryVals := make(url.Values)
	for k, v := range q {
		queryVals.Set(k, v)
	}

	queryVals.Set("random", "true")
	queryVals.Set("pageSize", strconv.Itoa(count))

	url := queryUrl + "cards?" + queryVals.Encode()
	cards, _, err := fetchCards(url)
	return cards, err
}

func (q query) Copy() Query {
	r := make(query)
	for k, v := range q {
		r[k] = v
	}
	return r
}

func (q query) Name(qry string) Query {
	q["name"] = qry
	return q
}

func (q query) Layout(qry string) Query {
	q["layout"] = qry
	return q
}

func (q query) CMC(qry string) Query {
	q["cmc"] = qry
	return q
}

func (q query) Colors(qry string) Query {
	q["colors"] = qry
	return q
}

func (q query) ColorIdentity(qry string) Query {
	q["colorIdentity"] = qry
	return q
}

func (q query) Type(qry string) Query {
	q["type"] = qry
	return q
}

func (q query) Supertypes(qry string) Query {
	q["supertypes"] = qry
	return q
}

func (q query) Types(qry string) Query {
	q["types"] = qry
	return q
}

func (q query) Subtypes(qry string) Query {
	q["subtypes"] = qry
	return q
}

func (q query) Rarity(qry string) Query {
	q["rarity"] = qry
	return q
}
func (q query) Set(qry string) Query {
	q["set"] = qry
	return q
}
func (q query) SetName(qry string) Query {
	q["setName"] = qry
	return q
}

func (q query) Text(qry string) Query {
	q["text"] = qry
	return q
}
func (q query) Flavor(qry string) Query {
	q["flavor"] = qry
	return q
}
func (q query) Artist(qry string) Query {
	q["artist"] = qry
	return q
}
func (q query) Number(qry string) Query {
	q["number"] = qry
	return q
}
func (q query) Power(qry string) Query {
	q["power"] = qry
	return q
}
func (q query) Toughness(qry string) Query {
	q["toughness"] = qry
	return q
}
func (q query) Loyalty(qry string) Query {
	q["loyalty"] = qry
	return q
}
func (q query) ForeignName(qry string) Query {
	q["foreignName"] = qry
	return q
}
func (q query) Language(qry string) Query {
	q["language"] = qry
	return q
}
func (q query) GameFormat(qry string) Query {
	q["gameFormat"] = qry
	return q
}
func (q query) Legality(qry string) Query {
	q["legality"] = qry
	return q
}
