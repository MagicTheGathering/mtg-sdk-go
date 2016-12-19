package mtg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type setColumn string

var (
	// SetName is the name of the set
	SetName = setColumn("name")
	// SetBlock is the block the set is in
	SetBlock = setColumn("block")
)

// SetCode representing one specific Set of cards
type SetCode string

// BoosterContent represent one or more types of cards within a booster
type BoosterContent []string

// Set stores information about a mtg-set
type Set struct {
	// The code name of the set
	SetCode `json:"code"`

	// The name of the set
	Name string `json:"name"`
	// The block the set is in
	Block string `json:"block"`
	// The code that Gatherer uses for the set. Only present if different than ‘code’
	GathererCode string `json:"gathererCode"`
	// An old style code used by some Magic software. Only present if different than 'gathererCode’ and 'code’
	OldCode string `json:"oldCode"`
	// The code that magiccards.info uses for the set. Only present if magiccards.info has this set
	MagicCardsInfoCode string `json:"magicCardsInfoCode"`
	// When the set was released (YYYY-MM-DD). For promo sets, the date the first card was released.
	ReleaseDate string `json:"releaseDate"`
	// The type of border on the cards, either “white”, “black” or “silver”
	Border string `json:"border"`
	// Type of set. One of: “core”, “expansion”, “reprint”, “box”, “un”, “from the vault”, “premium deck”, “duel deck”, “starter”, “commander”, “planechase”, “archenemy”, “promo”, “vanguard”, “masters”
	Expansion string `json:"expansion"`
	// Present and set to true if the set was only released online
	OnlineOnly bool `json:"onlineOnly"`
	// Booster contents for this set
	Booster []BoosterContent `json:"booster"`
}

// SetQuery is in Interface to query sets
type SetQuery interface {
	// Where filters the given column by the given value
	Where(col setColumn, qry string) SetQuery

	// Copy creates a copy of the SetQuery.
	Copy() SetQuery
	// All returns alls Sets which match the query
	All() ([]*Set, error)
	// Page returns the Sets of the given page and the total count of sets which match the query.
	// The default PageSize is 500. See also PageS
	Page(pageNum int) (sets []*Set, totalSetCount int, err error)
	// PageS returns the Sets of the given page and page size. It also returns the total count of sets
	// which match the query.
	PageS(pageNum int, pageSize int) (sets []*Set, totalSetCount int, err error)
}

type setQuery map[string]string

// GenerateBooster returns a slice of cards which contains cards like a booster of the given set.
func (sc SetCode) GenerateBooster() ([]*Card, error) {
	cards, _, err := fetchCards(fmt.Sprintf("%ssets/%s/booster", queryUrl, sc))
	return cards, err
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (bc *BoosterContent) UnmarshalJSON(data []byte) error {
	var s string
	var sc []string
	if err := json.Unmarshal(data, &s); err == nil {
		*bc = []string{s}
		return nil
	} else if err = json.Unmarshal(data, &sc); err == nil {
		*bc = sc
		return nil
	}
	return fmt.Errorf("Unexpected booster content. Got %q", string(data))
}

// String returns the string representation of the BoosterContent
func (bc *BoosterContent) String() string {
	s := ""
	for i, c := range *bc {
		if i > 0 {
			s += "|"
		}
		s += c
	}
	return s
}

// String returns the string representation for the Set
func (s *Set) String() string {
	return fmt.Sprintf("%s (%s)", s.Name, s.SetCode)
}

// NewSetQuery returns a new SetQuery
func NewSetQuery() SetQuery {
	return make(setQuery)
}

// Fetch returns the Set of the given SetCode.
func (sc SetCode) Fetch() (*Set, error) {
	sets, _, err := fetchSets(fmt.Sprintf("%ssets/%s", queryUrl, sc))
	if err != nil {
		return nil, err
	}
	if len(sets) != 1 {
		return nil, fmt.Errorf("Set %q not found", string(sc))
	}
	return sets[0], nil
}

func fetchSets(url string) ([]*Set, http.Header, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	if err := checkError(resp); err != nil {
		return nil, nil, err
	}

	sr := new(struct {
		Sets []*Set `json:"sets"`
		Set  *Set   `json:"set"`
	})
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&sr)
	if err != nil {
		return nil, nil, err
	}
	if sr.Set != nil {
		return []*Set{sr.Set}, resp.Header, nil
	}
	return sr.Sets, resp.Header, nil
}

// All returns alls Sets which match the query
func (q setQuery) All() ([]*Set, error) {
	var allSets []*Set

	queryVals := make(url.Values)
	for k, v := range q {
		queryVals.Set(k, v)
	}
	nextUrl := queryUrl + "sets?" + queryVals.Encode()
	for nextUrl != "" {
		sets, header, err := fetchSets(nextUrl)
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

		allSets = append(allSets, sets...)
	}
	return allSets, nil
}

// Page returns the Sets of the given page and the total count of sets which match the query.
// The default PageSize is 500. See also PageS
func (q setQuery) Page(pageNum int) (sets []*Set, totalSetCount int, err error) {
	return q.PageS(pageNum, 500)
}

// PageS returns the Sets of the given page and page size. It also returns the total count of sets
// which match the query.
func (q setQuery) PageS(pageNum int, pageSize int) (sets []*Set, totalSetCount int, err error) {
	sets = nil
	totalSetCount = 0
	err = nil

	queryVals := make(url.Values)
	for k, v := range q {
		queryVals.Set(k, v)
	}

	queryVals.Set("page", strconv.Itoa(pageNum))
	queryVals.Set("pageSize", strconv.Itoa(pageSize))

	url := queryUrl + "sets?" + queryVals.Encode()
	sets, header, err := fetchSets(url)
	if err != nil {
		return nil, 0, err
	}
	totalSetCount = len(sets)
	if totals, ok := header["Total-Count"]; ok && len(totals) > 0 {
		if totalSetCount, err = strconv.Atoi(totals[0]); err != nil {
			return nil, 0, err
		}
	}
	return sets, totalSetCount, nil
}

// Copy creates a copy of the SetQuery.
func (q setQuery) Copy() SetQuery {
	r := make(setQuery)
	for k, v := range q {
		r[k] = v
	}
	return r
}

func (q setQuery) Where(col setColumn, qry string) SetQuery {
	q[string(col)] = qry
	return q
}
