package mtg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Date which can be unmarshalled from json
type Date time.Time

// ServerError is an error implementation for server messages.
type ServerError struct {
	// Status code given by the server
	Status string `json:"status"`
	// Message given by the server
	Message string `json:"error"`
}

// Error implements the error interface
func (se ServerError) Error() string {
	return se.Message
}

// Id interface for different card id types such as MultiverseId or CardId
type Id interface {
	Fetch() (*Card, error)
}

// MultiverseId which can be used to fetch the card by its id
type MultiverseId string

// CardId which can be used to fetch the card by its id
type CardId string

// Ruling contains additional rule information about the card.
type Ruling struct {
	// Date the information was released.
	Date Date `json:"date"`
	// Text of the ruling hint.
	Text string `json:"text"`
}

// ForeignCardName represents the name of the card in an other language
type ForeignCardName struct {
	// Name is the name of the card in the given language
	Name string `json:"name"`
	// Language of the ForeignCardName
	Language string `json:"language"`
	// MultiverseId of the ForeignCardName (might be 0)
	MultiverseId uint `json:"multiverseid"`
}

// Legality stores information about legality notices for a specific format.
type Legality struct {
	// Format, such as Commander, Standard, Legacy, etc.
	Format string `json:"format"`
	// Legality for the given format such as Legal, Banned or Restricted.
	Legality string `json:"legality"`
}

// Card stores information about one single card.
type Card struct {
	// The card name. For split, double-faced and flip cards, just the name of one side of the card. Basically each ‘sub-card’ has its own record.
	Name string `json:"name"`
	// Only used for split, flip and dual cards. Will contain all the names on this card, front or back.
	Names []string `json:"names"`
	// The mana cost of this card. Consists of one or more mana symbols. (use cmc and colors to query)
	ManaCost string `json:"manaCost"`
	// Converted mana cost. Always a number.
	CMC float64 `json:"cmc"`
	// The card colors. Usually this is derived from the casting cost, but some cards are special (like the back of dual sided cards and Ghostfire).
	Colors []string `json:"colors"`
	// The card colors by color code. [“Red”, “Blue”] becomes [“R”, “U”]
	ColorIdentity []string `json:"colorIdentity"`
	// The card type. This is the type you would see on the card if printed today. Note: The dash is a UTF8 'long dash’ as per the MTG rules
	Type string `json:"type"`
	// The types of the card. These appear to the left of the dash in a card type. Example values: Instant, Sorcery, Artifact, Creature, Enchantment, Land, Planeswalker
	Types []string `json:"types"`
	// The supertypes of the card. These appear to the far left of the card type. Example values: Basic, Legendary, Snow, World, Ongoing
	Supertypes []string `json:"supertypes"`
	// The subtypes of the card. These appear to the right of the dash in a card type. Usually each word is its own subtype. Example values: Trap, Arcane, Equipment, Aura, Human, Rat, Squirrel, etc.
	Subtypes []string `json:"subtypes"`
	// The rarity of the card. Examples: Common, Uncommon, Rare, Mythic Rare, Special, Basic Land
	Rarity string `json:"rarity"`
	// The set the card belongs to (set code).
	Set SetCode `json:"set"`
	// The set the card belongs to.
	SetName string `json:"setName"`
	// The oracle text of the card. May contain mana symbols and other symbols.
	Text string `json:"text"`
	// The flavor text of the card.
	Flavor string `json:"flavor"`
	// The artist of the card. This may not match what is on the card as MTGJSON corrects many card misprints.
	Artist string `json:"artist"`
	// The card number. This is printed at the bottom-center of the card in small text. This is a string, not an integer, because some cards have letters in their numbers.
	Number string `json:"number"`
	// The power of the card. This is only present for creatures. This is a string, not an integer, because some cards have powers like: “1+*”
	Power string `json:"power"`
	// The toughness of the card. This is only present for creatures. This is a string, not an integer, because some cards have toughness like: “1+*”
	Toughness string `json:"toughness"`
	// The loyalty of the card. This is only present for planeswalkers.
	Loyalty string `json:"loyalty"`
	// The card layout. Possible values: normal, split, flip, double-faced, token, plane, scheme, phenomenon, leveler, vanguard
	Layout string `json:"layout"`
	// The multiverseid of the card on Wizard’s Gatherer web page. Cards from sets that do not exist on Gatherer will NOT have a multiverseid. Sets not on Gatherer are: ATH, ITP, DKM, RQS, DPA and all sets with a 4 letter code that starts with a lowercase 'p’.
	MultiverseId MultiverseId `json:"multiverseid"`
	// If a card has alternate art (for example, 4 different Forests, or the 2 Brothers Yamazaki) then each other variation’s multiverseid will be listed here, NOT including the current card’s multiverseid.
	Variations []string `json:"variations"`
	// The image url for a card. Only exists if the card has a multiverse id.
	ImageUrl string `json:"imageUrl"`
	// The watermark on the card. Note: Split cards don’t currently have this field set, despite having a watermark on each side of the split card.
	Watermark string `json:"watermark"`
	// If the border for this specific card is DIFFERENT than the border specified in the top level set JSON, then it will be specified here. (Example: Unglued has silver borders, except for the lands which are black bordered)
	Border string `json:"border"`
	// If this card was a timeshifted card in the set.
	Timeshifted bool `json:"timeshifted"`
	// Maximum hand size modifier. Only exists for Vanguard cards.
	Hand int `json:"hand"`
	// Starting life total modifier. Only exists for Vanguard cards.
	Life int `json:"life"`
	// Set to true if this card is reserved by Wizards Official Reprint Policy
	Reserved bool `json:"reserved"`
	// The date this card was released. This is only set for promo cards. The date may not be accurate to an exact day and month, thus only a partial date may be set (YYYY-MM-DD or YYYY-MM or YYYY). Some promo cards do not have a known release date.
	ReleaseDate Date `json:"releaseDate"`
	// Set to true if this card was only released as part of a core box set. These are technically part of the core sets and are tournament legal despite not being available in boosters.
	Starter bool `json:"starter"`
	// The rulings for the card.
	Rulings []*Ruling `json:"rulings"`
	// Foreign language names for the card, if this card in this set was printed in another language. An array of objects, each object having 'language’, 'name’ and 'multiverseid’ keys. Not available for all sets.
	ForeignNames []ForeignCardName `json:"foreignNames"`
	// The sets that this card was printed in, expressed as an array of set codes.
	Printings []SetCode `json:"printings"`
	// The original text on the card at the time it was printed. This field is not available for promo cards.
	OriginalText string `json:"originalText"`
	// The original type on the card at the time it was printed. This field is not available for promo cards.
	OriginalType string `json:"originalType"`
	// A unique id for this card. It is made up by doing an SHA1 hash of setCode + cardName + cardImageName
	Id CardId `json:"id"`
	// For promo cards, this is where this card was originally obtained. For box sets that are theme decks, this is which theme deck the card is from.
	Source string `json:"source"`
	// Which formats this card is legal, restricted or banned in. An array of objects, each object having 'format’ and 'legality’.
	Legalities []Legality `json:"legalities"`
}

// UnmarshalJSON implements the json.Unmarshaler interface. The Date is expected to be either YYYY, YYYY-MM or YYYY-MM-DD
func (d *Date) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	layouts := []string{
		"2006-01-02", "2006-01", "2006",
	}
	err = nil
	var t time.Time
	for _, layout := range layouts {
		t, err = time.Parse(layout, s)
		if err == nil {
			*d = Date(t)
			return nil
		}
	}
	return fmt.Errorf("%q is no valid date", s)
}

// String returns the string representation of the card. Containing the cardname and the id
func (c *Card) String() string {
	return fmt.Sprintf("%s (%s)", c.Name, c.Id)
}

type cardResponse struct {
	Card  *Card   `json:"card"`
	Cards []*Card `json:"cards"`
}

func decodeCards(reader io.Reader) ([]*Card, error) {
	cr := new(cardResponse)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&cr)
	if err != nil {
		return nil, err
	}
	if cr.Card != nil {
		return []*Card{cr.Card}, nil
	}
	return cr.Cards, nil
}

func checkError(r *http.Response) error {
	if r.StatusCode == 200 {
		return nil
	}

	var se ServerError

	if err := json.NewDecoder(r.Body).Decode(&se); err != nil {
		return errors.New(r.Status)
	}
	return se
}

func fetchCardById(str string) (*Card, error) {
	resp, err := http.Get(fmt.Sprintf("%scards/%s", queryUrl, str))
	if err != nil {
		return nil, err
	}
	bdy := resp.Body
	defer bdy.Close()

	if err := checkError(resp); err != nil {
		return nil, err
	}
	cards, err := decodeCards(bdy)
	if err != nil {
		return nil, err
	}
	if len(cards) != 1 {
		return nil, fmt.Errorf("Card with Id %s not found", str)
	}
	return cards[0], nil
}

// Fetch returns the card represented by the MutliverseId
func (mID MultiverseId) Fetch() (*Card, error) {
	return fetchCardById(fmt.Sprintf("%v", mID))
}

// Fetch returns the card represented by the CardId
func (id CardId) Fetch() (*Card, error) {
	return fetchCardById(string(id))
}
