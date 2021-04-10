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

	Convey("Fetching a planeswalker", t, func() {
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards/419003",
			httpmock.NewStringResponder(200, `{"cards":[{"name":"Chandra, Torch of Defiance","manaCost":"{2}{R}{R}","cmc":4.0,"colors":["Red"],"colorIdentity":["R"],"type":"Planeswalker — Chandra","types":["Planeswalker"],"subtypes":["Chandra"],"rarity":"Mythic Rare","set":"KLD","setName":"Kaladesh","text":"+1: Exile the top card of your library. You may cast that card. If you don't, Chandra, Torch of Defiance deals 2 damage to each opponent.\n+1: Add {R}{R} to your mana pool.\n−3: Chandra, Torch of Defiance deals 4 damage to target creature.\n−7: You get an emblem with \"Whenever you cast a spell, this emblem deals 5 damage to target creature or player.\"","artist":"Magali Villeneuve","number":"110","layout":"normal","multiverseid":417683,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=417683&type=card","loyalty":"4","rulings":[{"date":"2016-09-20","text":"An effect that instructs you to \"cast\" a card doesn't allow you to play lands. If the card exiled with Chandra's first ability is a land card, you can't play it and Chandra deals 2 damage to each opponent."},{"date":"2016-09-20","text":"If you cast the exiled card, you do so as part of the resolution of Chandra's ability. You can't wait to cast it later in the turn. Timing permissions based on the card's type are ignored, but other restrictions (such as \"Cast [this card] only during combat\") are not."},{"date":"2016-09-20","text":"You pay the costs for the exiled card if you cast it. You may pay alternative costs such as emerge rather than the card's mana cost."},{"date":"2016-09-20","text":"Loyalty abilities can't be mana abilities. Chandra's second ability uses the stack and can be countered or otherwise responded to. Like all loyalty abilities, it can be activated only once per turn, during your main phase, when the stack is empty, and only if no other loyalty abilities of the planeswalker have been activated this turn."},{"date":"2016-09-20","text":"The emblem created by Chandra's last ability is colorless. The damage it deals is from a colorless source."},{"date":"2016-09-20","text":"Chandra's emblem's ability resolves before the spell that caused it to trigger."},{"date":"2016-09-20","text":"In a Two-Headed Giant game, Chandra's first ability causes 4 damage total to be dealt to the opposing team."}],"foreignNames":[{"name":"反抗烈炬茜卓","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=417947&type=card","language":"Chinese Simplified","multiverseid":417947},{"name":"反抗烈炬茜卓","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418211&type=card","language":"Chinese Traditional","multiverseid":418211},{"name":"Chandra, torche de la défiance","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418739&type=card","language":"French","multiverseid":418739},{"name":"Chandra, Fackel des Widerstands","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418475&type=card","language":"German","multiverseid":418475},{"name":"Chandra, Fiamma di Sfida","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419003&type=card","language":"Italian","multiverseid":419003},{"name":"反逆の先導者、チャンドラ","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419267&type=card","language":"Japanese","multiverseid":419267},{"name":"저항의 횃불 찬드라","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419531&type=card","language":"Korean","multiverseid":419531},{"name":"Chandra, Chama da Rebeldia","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419795&type=card","language":"Portuguese (Brazil)","multiverseid":419795},{"name":"Чандра, Факел Непокорности","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=420059&type=card","language":"Russian","multiverseid":420059},{"name":"Chandra, aurora de la rebeldía","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=420323&type=card","language":"Spanish","multiverseid":420323}],"printings":["KLD"],"originalText":"+1: Exile the top card of your library. You may cast that card. If you don't, Chandra, Torch of Defiance deals 2 damage to each opponent.\n+1: Add {R}{R} to your mana pool.\n−3: Chandra, Torch of Defiance deals 4 damage to target creature.\n−7: You get an emblem with \"Whenever you cast a spell, this emblem deals 5 damage to target creature or player.\"","originalType":"Planeswalker — Chandra","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Kaladesh Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Standard","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"0ef97e4324dbdcc0eaedda8f4301f68f3567d2ca"}]}`))

		card, err := MultiverseId("419003").Fetch()

		So(err, ShouldBeNil)
		So(card, ShouldNotBeNil)

		So(card.Name, ShouldEqual, "Chandra, Torch of Defiance")
		So(card.Loyalty, ShouldEqual, "4")
		So(card.CMC, ShouldEqual, 4.0)
		So(card.Rulings, ShouldNotBeEmpty)
		So(card.ForeignNames, ShouldNotBeEmpty)
		So(card.Variations, ShouldBeEmpty)
	})

	Convey("Fetching cards by id", t, func() {
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards/417594",
			httpmock.NewStringResponder(200, `{"card":{"name":"Master Trinketeer","manaCost":"{2}{W}","cmc":3.0,"colors":["White"],"colorIdentity":["W"],"type":"Creature — Dwarf Artificer","types":["Creature"],"subtypes":["Dwarf","Artificer"],"rarity":"Rare","set":"KLD","setName":"Kaladesh","text":"Servos and Thopters you control get +1/+1.\n{3}{W}: Create a 1/1 colorless Servo artifact creature token.","flavor":"\"Let us never forget the joy that lies at the heart of invention.\"","artist":"Matt Stewart","number":"21","power":"3","toughness":"2","layout":"normal","multiverseid":417594,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=417594&type=card","rulings":[{"date":"2016-09-20","text":"A creature that is both a Servo and a Thopter gets +1/+1, not +2/+2."}],"foreignNames":[{"name":"琐物大师","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=417858&type=card","language":"Chinese Simplified","multiverseid":417858},{"name":"瑣物大師","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418122&type=card","language":"Chinese Traditional","multiverseid":418122},{"name":"Maître bibeloteur","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418650&type=card","language":"French","multiverseid":418650},{"name":"Meisterhafter Feinmechaniker","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418386&type=card","language":"German","multiverseid":418386},{"name":"Maestro Fabbricagingilli","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=418914&type=card","language":"Italian","multiverseid":418914},{"name":"小物作りの達人","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419178&type=card","language":"Japanese","multiverseid":419178},{"name":"장신구 공예의 대가","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419442&type=card","language":"Korean","multiverseid":419442},{"name":"Bugigangueiro Mestre","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419706&type=card","language":"Portuguese (Brazil)","multiverseid":419706},{"name":"Мастер Диковинок","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=419970&type=card","language":"Russian","multiverseid":419970},{"name":"Maestro artilugista","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=420234&type=card","language":"Spanish","multiverseid":420234}],"printings":["KLD"],"originalText":"Servos and Thopters you control get +1/+1.\n{3}{W}: Create a 1/1 colorless Servo artifact creature token.","originalType":"Creature — Dwarf Artificer","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Kaladesh Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Standard","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"b246420195aba6e1ef8e8e20f73c40a1130444fc"}}`))
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
			card, err := MultiverseId("417594").Fetch()

			So(card, ShouldNotBeNil)
			So(err, ShouldBeNil)

			So(card.Name, ShouldEqual, "Master Trinketeer")
			So(card.CMC, ShouldEqual, 3.0)
			So(card.Power, ShouldEqual, "3")
			So(card.Rulings, ShouldNotBeEmpty)
			So(card.ForeignNames, ShouldNotBeEmpty)
			So(card.Variations, ShouldBeEmpty)
		})

		Convey("Fetching an invalid MultiverseId", func() {
			card, err := MultiverseId("0").Fetch()
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

			card, err = MultiverseId("1").Fetch()
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
