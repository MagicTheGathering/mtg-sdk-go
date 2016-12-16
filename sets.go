package mtg

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SetCode string

type BoosterContent []string

type Set struct {
	// The name of the set
	Name string `json:"name"`
	// The block the set is in
	Block string `json:"block"`
	// The code name of the set
	Code SetCode `json:"code"`
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
	Booster []*BoosterContent `json:"booster"`
}

type SetQuery interface {
	// The name of the set
	Name(qry string) SetQuery
	// The block the set is in
	Block(qry string) SetQuery

	// Creates a copy of this query
	Copy() SetQuery
	// Fetches all sets matching the current query
	All() ([]*Set, error)
	// Fetches the given page of sets.
	Page(pageNum int) (sets []*Set, totalSetCount int, err error)
	// Fetches one page of sets with a given page size
	PageS(pageNum int, pageSize int) (sets []*Set, totalSetCount int, err error)
}

type setQuery map[string]string

func (bc *BoosterContent) UnmarshalJSON(data []byte) error {
	var s string
	var sc []string
	if err := json.Unmarshal(data, &s); err == nil {
		*bc = []string{s}
	} else if err = json.Unmarshal(data, &sc); err == nil {
		*bc = sc
	} else {
		return err
	}
	return nil
}

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

func NewSetQuery() SetQuery {
	return make(setQuery)
}

func (q setQuery) fetch(url string) ([]*Set, http.Header, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

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
	} else {
		return sr.Sets, resp.Header, nil
	}
}

func (q setQuery) All() ([]*Set, error) {
	var allSets []*Set

	queryVals := make(url.Values)
	for k, v := range q {
		queryVals.Set(k, v)
	}
	nextUrl := queryUrl + "sets?" + queryVals.Encode()
	for nextUrl != "" {
		sets, header, err := q.fetch(nextUrl)
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

func (q setQuery) Page(pageNum int) (sets []*Set, totalSetCount int, err error) {
	return q.PageS(pageNum, 500)
}

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
	sets, header, err := q.fetch(url)
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

func (q setQuery) Copy() SetQuery {
	r := make(setQuery)
	for k, v := range q {
		r[k] = v
	}
	return r
}

func (q setQuery) Name(qry string) SetQuery {
	q["name"] = qry
	return q
}

func (q setQuery) Block(qry string) SetQuery {
	q["block"] = qry
	return q
}
