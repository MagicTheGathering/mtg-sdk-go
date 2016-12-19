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

type cardColumn string

var (
	// The card name. For split, double-faced and flip cards, just the name of one side of the card. Basically each ‘sub-card’ has its own record.
	Card_Name = cardColumn("name")
	// The card layout. Possible values: normal, split, flip, double-faced, token, plane, scheme, phenomenon, leveler, vanguard
	Card_Layout = cardColumn("layout")
	// Converted mana cost. Always a number.
	Card_CMC = cardColumn("cmc")
	// The card colors. Usually this is derived from the casting cost, but some cards are special (like the back of dual sided cards and Ghostfire).
	Card_Colors = cardColumn("colors")
	// The card colors by color code. [“Red”, “Blue”] becomes [“R”, “U”]
	Card_ColorIdentity = cardColumn("colorIdentity")
	// The card type. This is the type you would see on the card if printed today. Note: The dash is a UTF8 'long dash’ as per the MTG rules
	Card_Type = cardColumn("type")
	// The supertypes of the card. These appear to the far left of the card type. Example values: Basic, Legendary, Snow, World, Ongoing
	Card_Supertypes = cardColumn("supertypes")
	// The types of the card. These appear to the left of the dash in a card type. Example values: Instant, Sorcery, Artifact, Creature, Enchantment, Land, Planeswalker
	Card_Types = cardColumn("types")
	// The subtypes of the card. These appear to the right of the dash in a card type. Usually each word is its own subtype. Example values: Trap, Arcane, Equipment, Aura, Human, Rat, Squirrel, etc.
	Card_Subtypes = cardColumn("subtypes")
	// The rarity of the card. Examples: Common, Uncommon, Rare, Mythic Rare, Special, Basic Land
	Card_Rarity = cardColumn("rarity")
	// The set the card belongs to (set code).
	Card_Set = cardColumn("set")
	// The set the card belongs to.
	Card_SetName = cardColumn("setName")
	// The oracle text of the card. May contain mana symbols and other symbols.
	Card_Text = cardColumn("text")
	// The flavor text of the card.
	Card_Flavor = cardColumn("flavor")
	// The artist of the card. This may not match what is on the card as MTGJSON corrects many card misprints.
	Card_Artist = cardColumn("artist")
	// The card number. This is printed at the bottom-center of the card in small text. This is a string, not an integer, because some cards have letters in their numbers.
	Card_Number = cardColumn("number")
	// The power of the card. This is only present for creatures. This is a string, not an integer, because some cards have powers like: “1+*”
	Card_Power = cardColumn("power")
	// The toughness of the card. This is only present for creatures. This is a string, not an integer, because some cards have toughness like: “1+*”
	Card_Toughness = cardColumn("toughness")
	// The loyalty of the card. This is only present for planeswalkers.
	Card_Loyalty = cardColumn("loyalty")
	// The name of a card in a foreign language it was printed in
	Card_ForeignName = cardColumn("foreignName")
	// The language the card is printed in. Use this parameter when searching by foreignName
	Card_Language = cardColumn("language")
	// The game format, such as Commander, Standard, Legacy, etc. (when used, legality defaults to Legal unless supplied)
	Card_GameFormat = cardColumn("gameFormat")
	// The legality of the card for a given format, such as Legal, Banned or Restricted.
	Card_Legality = cardColumn("legality")
)

// Query interface can be used to query multiple cards by their properties
type Query interface {
	// Where filters the given column by the given value
	Where(column cardColumn, qry string) Query
	// Sorts the query results by the given column
	OrderBy(column cardColumn) Query

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

// NewQuery creates a new Query to fetch cards
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

	if err := checkError(resp); err != nil {
		return nil, nil, err
	}
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

func (q query) Where(column cardColumn, qry string) Query {
	q[string(column)] = qry
	return q
}

func (q query) OrderBy(column cardColumn) Query {
	q["orderBy"] = string(column)
	return q
}
