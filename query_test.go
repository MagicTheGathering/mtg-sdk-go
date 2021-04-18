package mtg

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/jarcoal/httpmock.v1"
)

func Test_Querys(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	Convey("When executing queries", t, func() {
		qry := NewQuery()
		So(qry, ShouldNotBeNil)

		Convey("and query for red rares", func() {
			qry = qry.Where(CardColors, "red").Where(CardRarity, "rare").OrderBy(CardCMC)

			Convey("a copy of the query should make no difference", func() {
				other := qry.Copy()

				So(other, ShouldResemble, qry)
			})

			Convey("when fetching two random cards", func() {
				httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&pageSize=2&random=true&rarity=rare",
					httpmock.NewStringResponder(200, `{"cards":[{"name":"Karplusan Yeti","manaCost":"{3}{R}{R}","cmc":5.0,"colors":["Red"],"colorIdentity":["R"],"type":"Creature — Yeti","types":["Creature"],"subtypes":["Yeti"],"rarity":"Rare","set":"ICE","setName":"Ice Age","text":"{T}: Karplusan Yeti deals damage equal to its power to target creature. That creature deals damage equal to its power to Karplusan Yeti.","flavor":"\"What's that smell?\"\n—Perena Deepcutter,\nDwarven Armorer","artist":"Quinton Hoover","power":"3","toughness":"3","layout":"normal","multiverseid":2633,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=2633&type=card","rulings":[{"date":"2004-10-04","text":"Giving either creature first strike does not affect the ability."},{"date":"2004-10-04","text":"If this leaves the battlefield before its activated ability resolves, it will still deal damage to the targeted creature. On the other hand, if the targeted creature leaves the battlefield before the ability resolves, the ability will be countered and no damage will be dealt."}],"printings":["ICE","9ED"],"originalText":"{T}: Karplusan Yeti deals an amount of damage equal to its power to target creature. That creature deals an amount of damage equal to its power to Karplusan Yeti.","originalType":"Summon — Yeti","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Ice Age Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"441ccf2c2ad92d25852284238859ef5ed556a1fe"},{"name":"Flowstone Overseer","manaCost":"{2}{R}{R}{R}","cmc":5,"colors":["Red"],"colorIdentity":["R"],"type":"Creature — Beast","types":["Creature"],"subtypes":["Beast"],"rarity":"Rare","set":"NMS","setName":"Nemesis","text":"{R}{R}: Target creature gets +1/-1 until end of turn.","flavor":"The rebels couldn't see where the roar was coming from. Then they saw it was coming from everywhere.","artist":"Andrew Goldhawk","number":"82","power":"4","toughness":"4","layout":"normal","multiverseid":21351,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=21351&type=card","printings":["NMS"],"originalText":"{R}{R}: Target creature gets +1/-1 until end of turn.","originalType":"Creature — Beast","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Masques Block","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"802f0aeb631bbc77b1d6dbe780f421152370a683"}]}`))

				cards, err := qry.Random(2)

				Convey("There should be no error", func() {
					So(err, ShouldBeNil)
				})
				Convey("it should return two cards", func() {
					So(cards, ShouldHaveLength, 2)
					So(cards, ShouldContainCard, "Karplusan Yeti")
					So(cards, ShouldContainCard, "Flowstone Overseer")
				})
			})

			Convey("when fetching a single page", func() {
				httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&page=1&pageSize=100&rarity=rare",
					NewStringResponderWithHeader(200, `{"cards":[{"name":"Pact of the Titan","manaCost":"{0}","cmc":0.0,"colors":["Red"],"colorIdentity":["R"],"type":"Instant","types":["Instant"],"rarity":"Rare","set":"FUT","setName":"Future Sight","text":"Create a 4/4 red Giant creature token.\nAt the beginning of your next upkeep, pay {4}{R}. If you don't, you lose the game.","artist":"Raymond Swanland","number":"103","layout":"normal","multiverseid":130638,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=130638&type=card","rulings":[{"date":"2013-06-07","text":"Although originally printed with a characteristic-defining ability that defined its color, this card now has a color indicator. This color indicator can't be affected by text-changing effects (such as the one created by Crystal Spray), although color-changing effects can still overwrite it."}],"foreignNames":[{"name":"泰坦条约","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=144089&type=card","language":"Chinese Simplified","multiverseid":144089},{"name":"Pacte du titan","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=145693&type=card","language":"French","multiverseid":145693},{"name":"Pakt mit dem Titanen","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=143954&type=card","language":"German","multiverseid":143954},{"name":"Patto del Titano","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=145594&type=card","language":"Italian","multiverseid":145594},{"name":"タイタンの契約","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=145792&type=card","language":"Japanese","multiverseid":145792},{"name":"Pacto do Titã","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=144188&type=card","language":"Portuguese (Brazil)","multiverseid":144188},{"name":"Договор Титана","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=144487&type=card","language":"Russian","multiverseid":144487},{"name":"Pacto del titán","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=144388&type=card","language":"Spanish","multiverseid":144388}],"printings":["FUT"],"originalText":"Pact of the Titan is red.\nPut a 4/4 red Giant creature token into play.\nAt the beginning of your next upkeep, pay {4}{R}. If you don't, you lose the game.","originalType":"Instant","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Time Spiral Block","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"61992dde29e8f808bc6742ac71db85a72f222fc4"},{"name":"Earthquake","manaCost":"{X}{R}","cmc":1,"colors":["Red"],"colorIdentity":["R"],"type":"Sorcery","types":["Sorcery"],"rarity":"Rare","set":"LEA","setName":"Limited Edition Alpha","text":"Earthquake deals X damage to each creature without flying and each player.","artist":"Dan Frazier","layout":"normal","multiverseid":194,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=194&type=card","rulings":[{"date":"2004-10-04","text":"Whether or not a creature is without flying is only checked on resolution."}],"printings":["LEA","LEB","2ED","CED","CEI","3ED","4ED","5ED","POR","PO2","6ED","7ED","M10","CMD","C15"],"originalText":"Does X damage to each player and each non-flying creature in play.","originalType":"Sorcery","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"0657fac72175ae3cae88198983258f99061ab982"}]}`,
						map[string]string{
							"Total-Count": "1337",
						}))
				cards, totalCards, err := qry.Page(1)

				Convey("there should be no error", func() {
					So(err, ShouldBeNil)
				})
				Convey("the total card count should be read from the header", func() {
					So(totalCards, ShouldEqual, 1337)

					Convey("but if the header was invalid, it should return an error", func() {
						httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&page=1&pageSize=100&rarity=rare",
							NewStringResponderWithHeader(200, `{"cards":[{"name":"Pact of the Titan","manaCost":"{0}","cmc":0.0,"colors":["Red"],"colorIdentity":["R"],"type":"Instant","types":["Instant"],"rarity":"Rare","set":"FUT","setName":"Future Sight","text":"Create a 4/4 red Giant creature token.\nAt the beginning of your next upkeep, pay {4}{R}. If you don't, you lose the game.","artist":"Raymond Swanland","number":"103","layout":"normal","multiverseid":130638,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=130638&type=card","rulings":[{"date":"2013-06-07","text":"Although originally printed with a characteristic-defining ability that defined its color, this card now has a color indicator. This color indicator can't be affected by text-changing effects (such as the one created by Crystal Spray), although color-changing effects can still overwrite it."}],"foreignNames":[{"name":"泰坦条约","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=144089&type=card","language":"Chinese Simplified","multiverseid":144089},{"name":"Pacte du titan","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=145693&type=card","language":"French","multiverseid":145693},{"name":"Pakt mit dem Titanen","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=143954&type=card","language":"German","multiverseid":143954},{"name":"Patto del Titano","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=145594&type=card","language":"Italian","multiverseid":145594},{"name":"タイタンの契約","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=145792&type=card","language":"Japanese","multiverseid":145792},{"name":"Pacto do Titã","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=144188&type=card","language":"Portuguese (Brazil)","multiverseid":144188},{"name":"Договор Титана","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=144487&type=card","language":"Russian","multiverseid":144487},{"name":"Pacto del titán","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=144388&type=card","language":"Spanish","multiverseid":144388}],"printings":["FUT"],"originalText":"Pact of the Titan is red.\nPut a 4/4 red Giant creature token into play.\nAt the beginning of your next upkeep, pay {4}{R}. If you don't, you lose the game.","originalType":"Instant","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Time Spiral Block","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"61992dde29e8f808bc6742ac71db85a72f222fc4"},{"name":"Earthquake","manaCost":"{X}{R}","cmc":1,"colors":["Red"],"colorIdentity":["R"],"type":"Sorcery","types":["Sorcery"],"rarity":"Rare","set":"LEA","setName":"Limited Edition Alpha","text":"Earthquake deals X damage to each creature without flying and each player.","artist":"Dan Frazier","layout":"normal","multiverseid":194,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=194&type=card","rulings":[{"date":"2004-10-04","text":"Whether or not a creature is without flying is only checked on resolution."}],"printings":["LEA","LEB","2ED","CED","CEI","3ED","4ED","5ED","POR","PO2","6ED","7ED","M10","CMD","C15"],"originalText":"Does X damage to each player and each non-flying creature in play.","originalType":"Sorcery","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"0657fac72175ae3cae88198983258f99061ab982"}]}`,
								map[string]string{
									"Total-Count": "two",
								}))
						_, _, err := qry.Page(1)
						So(err, ShouldNotBeNil)
					})
				})
				Convey("and the result should contain all send cards", func() {
					So(cards, ShouldContainCard, "Pact of the Titan")
					So(cards, ShouldContainCard, "Earthquake")
				})

				Convey("invalid json should return an error", func() {
					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&page=1&pageSize=100&rarity=rare",
						NewStringResponderWithHeader(200, `{"cards":}`,
							map[string]string{
								"Total-Count": "1337",
							}))
					_, _, err := qry.Page(1)
					So(err, ShouldNotBeNil)
				})

				Convey("Network issues should return an error", func() {
					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&page=1&pageSize=100&rarity=rare",
						httpmock.NewErrorResponder(errors.New("Network Issue")))
					_, _, err := qry.Page(1)
					So(err, ShouldNotBeNil)
				})

				Convey("If the server returns an error it should be returned", func() {
					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&page=1&pageSize=100&rarity=rare",
						httpmock.NewStringResponder(404, `{"status": "404", "error": "Not found"}`))
					_, _, err := qry.Page(1)
					So(err, ShouldNotBeNil)

					_, isServerError := err.(ServerError)
					So(isServerError, ShouldBeTrue)
				})
			})

			Convey("when fetching all cards", func() {
				httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&rarity=rare",
					NewStringResponderWithHeader(200, `{"cards":[{"name":"Karplusan Yeti","manaCost":"{3}{R}{R}","cmc":5.0,"colors":["Red"],"colorIdentity":["R"],"type":"Creature — Yeti","types":["Creature"],"subtypes":["Yeti"],"rarity":"Rare","set":"ICE","setName":"Ice Age","text":"{T}: Karplusan Yeti deals damage equal to its power to target creature. That creature deals damage equal to its power to Karplusan Yeti.","flavor":"\"What's that smell?\"\n—Perena Deepcutter,\nDwarven Armorer","artist":"Quinton Hoover","power":"3","toughness":"3","layout":"normal","multiverseid":2633,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=2633&type=card","rulings":[{"date":"2004-10-04","text":"Giving either creature first strike does not affect the ability."},{"date":"2004-10-04","text":"If this leaves the battlefield before its activated ability resolves, it will still deal damage to the targeted creature. On the other hand, if the targeted creature leaves the battlefield before the ability resolves, the ability will be countered and no damage will be dealt."}],"printings":["ICE","9ED"],"originalText":"{T}: Karplusan Yeti deals an amount of damage equal to its power to target creature. That creature deals an amount of damage equal to its power to Karplusan Yeti.","originalType":"Summon — Yeti","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Ice Age Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"441ccf2c2ad92d25852284238859ef5ed556a1fe"},{"name":"Flowstone Overseer","manaCost":"{2}{R}{R}{R}","cmc":5.0,"colors":["Red"],"colorIdentity":["R"],"type":"Creature — Beast","types":["Creature"],"subtypes":["Beast"],"rarity":"Rare","set":"NMS","setName":"Nemesis","text":"{R}{R}: Target creature gets +1/-1 until end of turn.","flavor":"The rebels couldn't see where the roar was coming from. Then they saw it was coming from everywhere.","artist":"Andrew Goldhawk","number":"82","power":"4","toughness":"4","layout":"normal","multiverseid":21351,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=21351&type=card","printings":["NMS"],"originalText":"{R}{R}: Target creature gets +1/-1 until end of turn.","originalType":"Creature — Beast","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Masques Block","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"802f0aeb631bbc77b1d6dbe780f421152370a683"}]}`,
						map[string]string{
							"Link": `<https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&rarity=rare&page=2>; rel="last", <https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&rarity=rare&page=2>; rel="next"`,
						}))
				httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&page=2&rarity=rare",
					httpmock.NewStringResponder(200, `{"cards":[{"name":"Earthquake","manaCost":"{X}{R}","cmc":1.0,"colors":["Red"],"colorIdentity":["R"],"type":"Sorcery","types":["Sorcery"],"rarity":"Rare","set":"LEA","setName":"Limited Edition Alpha","text":"Earthquake deals X damage to each creature without flying and each player.","artist":"Dan Frazier","layout":"normal","multiverseid":194,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=194&type=card","rulings":[{"date":"2004-10-04","text":"Whether or not a creature is without flying is only checked on resolution."}],"printings":["LEA","LEB","2ED","CED","CEI","3ED","4ED","5ED","POR","PO2","6ED","7ED","M10","CMD","C15"],"originalText":"Does X damage to each player and each non-flying creature in play.","originalType":"Sorcery","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"0657fac72175ae3cae88198983258f99061ab982"},{"name":"Earthquake","manaCost":"{X}{R}","cmc":1.0,"colors":["Red"],"colorIdentity":["R"],"type":"Sorcery","types":["Sorcery"],"rarity":"Rare","set":"LEB","setName":"Limited Edition Beta","text":"Earthquake deals X damage to each creature without flying and each player.","artist":"Dan Frazier","layout":"normal","multiverseid":489,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=489&type=card","rulings":[{"date":"2004-10-04","text":"Whether or not a creature is without flying is only checked on resolution."}],"printings":["LEA","LEB","2ED","CED","CEI","3ED","4ED","5ED","POR","PO2","6ED","7ED","M10","CMD","C15"],"originalText":"Does X damage to each player and each non-flying creature in play.","originalType":"Sorcery","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"adcbbe0e1679f760d5428d3bbad879d6abf8807d"}]}`))

				cards, err := qry.All()

				Convey("there should be no error", func() {
					So(err, ShouldBeNil)
				})
				Convey("there should be 4 cards", func() {
					So(cards, ShouldHaveLength, 4)
					So(cards, ShouldContainCard, "Karplusan Yeti")
					So(cards, ShouldContainCard, "Flowstone Overseer")
					So(cards, ShouldContainCard, "Earthquake")
				})

				Convey("if there is an error on the second request it should be reported", func() {
					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/cards?colors=red&orderBy=cmc&rarity=rare&page=2",
						httpmock.NewErrorResponder(errors.New("Network issue")))
					_, err := qry.All()

					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
