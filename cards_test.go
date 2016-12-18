package mtg

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/jarcoal/httpmock.v1"
)

func Test_FetchCards(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	Convey("Fetching cards by id", t, func() {
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards/417594",
			httpmock.NewStringResponder(200, `{"card":{"name":"Master Trinketeer","manaCost":"{2}{W}","cmc":3,"colors":["White"],"colorIdentity":["W"],"type":"Creature — Dwarf Artificer","types":["Creature"],"subtypes":["Dwarf","Artificer"],"rarity":"Rare","set":"KLD","setName":"Kaladesh","text":"Servos and Thopters you control get +1/+1.\n{3}{W}: Create a 1/1 colorless Servo artifact creature token.","flavor":"\"Let us never forget the joy that lies at the heart of invention.\"","artist":"Matt Stewart","number":"21","power":"3","toughness":"2","layout":"normal","multiverseid":417594,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=417594&type=card","rulings":[{"date":"2016-09-20","text":"A creature that is both a Servo and a Thopter gets +1/+1, not +2/+2."}],"foreignNames":[{"name":"琐物大师","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=417858&type=card","language":"Chinese Simplified","multiverseid":417858},{"name":"瑣物大師","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418122&type=card","language":"Chinese Traditional","multiverseid":418122},{"name":"Maître bibeloteur","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418650&type=card","language":"French","multiverseid":418650},{"name":"Meisterhafter Feinmechaniker","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418386&type=card","language":"German","multiverseid":418386},{"name":"Maestro Fabbricagingilli","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418914&type=card","language":"Italian","multiverseid":418914},{"name":"小物作りの達人","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419178&type=card","language":"Japanese","multiverseid":419178},{"name":"장신구 공예의 대가","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419442&type=card","language":"Korean","multiverseid":419442},{"name":"Bugigangueiro Mestre","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419706&type=card","language":"Portuguese (Brazil)","multiverseid":419706},{"name":"Мастер Диковинок","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419970&type=card","language":"Russian","multiverseid":419970},{"name":"Maestro artilugista","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=420234&type=card","language":"Spanish","multiverseid":420234}],"printings":["KLD"],"originalText":"Servos and Thopters you control get +1/+1.\n{3}{W}: Create a 1/1 colorless Servo artifact creature token.","originalType":"Creature — Dwarf Artificer","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Kaladesh Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Standard","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"b246420195aba6e1ef8e8e20f73c40a1130444fc"}}`))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards/0",
			httpmock.NewStringResponder(404, `{"status":"404","error":"Not Found"}`))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards/1",
			httpmock.NewErrorResponder(errors.New("request failed")))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards/invalidJson",
			httpmock.NewStringResponder(200, `{"card":{"name":1}}`))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards/noCardsInResponse",
			httpmock.NewStringResponder(200, `{"card":null}`))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards/noErrorMsg",
			httpmock.NewStringResponder(500, ``))

		Convey("Fetching a MultiverseId", func() {
			card, err := MultiverseId(417594).Fetch()

			So(card, ShouldNotBeNil)
			So(err, ShouldBeNil)

			So(card.Name, ShouldEqual, "Master Trinketeer")
			So(card.CMC, ShouldEqual, 3)
			So(card.Power, ShouldEqual, "3")
			So(card.Rulings, ShouldNotBeEmpty)
			So(card.ForeignNames, ShouldNotBeEmpty)
			So(card.Variations, ShouldBeEmpty)
		})

		Convey("Fetching an invalid MultiverseId", func() {
			card, err := MultiverseId(0).Fetch()
			So(card, ShouldBeNil)
			So(err, ShouldNotBeNil)

			_, ok := err.(ServerError)
			So(ok, ShouldBeTrue)
		})

		Convey("Fetching a CardId", func() {

		})

		Convey("On Server errors", func() {
			card, err := CardId("invalidJson").Fetch()
			So(card, ShouldBeNil)
			So(err, ShouldNotBeNil)

			card, err = MultiverseId(1).Fetch()
			So(card, ShouldBeNil)
			So(err, ShouldNotBeNil)

			card, err = CardId("noErrorMsg").Fetch()
			So(card, ShouldBeNil)
			So(err, ShouldNotBeNil)

			_, ok := err.(ServerError)
			So(ok, ShouldBeFalse)

			card, err = CardId("noCardsInResponse").Fetch()
			So(card, ShouldBeNil)
			So(err, ShouldNotBeNil)

		})
	})
}
