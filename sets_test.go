package mtg

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/jarcoal/httpmock.v1"
)

func Test_GenerateBooster(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	Convey("When generating a booster", t, func() {
		Convey("If the response is correct", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/PLS/booster",
				httpmock.NewStringResponder(200, `{"cards":[{"name":"Planeswalker's Fury","manaCost":"{2}{R}","cmc":3,"colors":["Red"],"type":"Enchantment","types":["Enchantment"],"rarity":"Rare","set":"PLS","text":"{3}{R}: Target opponent reveals a card at random from his or her hand. Planeswalker's Fury deals damage equal to that card's converted mana cost to that player. Activate this ability only any time you could cast a sorcery.","artist":"Christopher Moeller","number":"70","layout":"normal","multiverseid":25889,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=25889&type=card","rulings":[{"date":"2004-10-04","text":"If the opponent has no cards in hand, then no damage is dealt."}],"foreignNames":[{"name":"Fureur selon l'Arpenteur","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185216&type=card","language":"French","multiverseid":185216},{"name":"Zorn der Weltenwanderer","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183144&type=card","language":"German","multiverseid":183144}],"printings":["PLS"],"originalText":"{3}{R}: Target opponent reveals a card at random from his or her hand. Planeswalker's Fury deals damage equal to that card's converted mana cost to that player. Play this ability only any time you could play a sorcery.","originalType":"Enchantment","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"f55c427829b3876a075fa6edd55145e4dfdae41a"},{"name":"Strafe","manaCost":"{R}","cmc":1,"colors":["Red"],"type":"Sorcery","types":["Sorcery"],"rarity":"Uncommon","set":"PLS","text":"Strafe deals 3 damage to target nonred creature.","flavor":"\"All right, let's light ‘em up\"\n—Gerrard","artist":"Jim Nelson","number":"73","layout":"normal","multiverseid":26288,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26288&type=card","foreignNames":[{"name":"Bombardement","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185219&type=card","language":"French","multiverseid":185219},{"name":"Tiefflugangriff","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183147&type=card","language":"German","multiverseid":183147}],"printings":["PLS"],"originalText":"Strafe deals 3 damage to target nonred creature.","originalType":"Sorcery","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"f9baef2844f0af893995ec7dd8b7d888cb08f95a"},{"name":"Crosis's Catacombs","type":"Land — Lair","types":["Land"],"subtypes":["Lair"],"rarity":"Uncommon","set":"PLS","text":"When Crosis's Catacombs enters the battlefield, sacrifice it unless you return a non-Lair land you control to its owner's hand.\n{T}: Add {U}, {B}, or {R} to your mana pool.","artist":"Edward P. Beard, Jr.","number":"136","layout":"normal","multiverseid":25936,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=25936&type=card","rulings":[{"date":"2004-10-04","text":"You can return a land that is tapped or untapped."},{"date":"2004-10-04","text":"If you don't want to unsummon a land, you can play this card then tap it for mana before the enters the battlefield ability resolves. You may then choose to sacrifice it instead of unsummoning a land."},{"date":"2005-08-01","text":"This land is of type \"Lair\" only; other subtypes have been removed. It is not a basic land."}],"foreignNames":[{"name":"Catacombes de Crosis","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185285&type=card","language":"French","multiverseid":185285},{"name":"Crosis' Katakomben","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183213&type=card","language":"German","multiverseid":183213}],"printings":["PLS"],"originalText":"Crosis's Catacombs is a Lair in addition to its land type.\nWhen Crosis's Catacombs comes into play, sacrifice it unless you return a non-Lair land you control to its owner's hand.\n{T}: Add {U}, {B}, or {R} to your mana pool.","originalType":"Land","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"dfd6c01d1beeea1562f44c29d8cc81839fadbe7f"},{"name":"Warped Devotion","manaCost":"{2}{B}","cmc":3,"colors":["Black"],"type":"Enchantment","types":["Enchantment"],"rarity":"Uncommon","set":"PLS","text":"Whenever a permanent is returned to a player's hand, that player discards a card.","flavor":"\"Before the glory of Yawgmoth, yes, even this makes sense.\"\n—Urza, to Gerrard","artist":"Orizio Daniele","number":"57","layout":"normal","multiverseid":26836,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26836&type=card","rulings":[{"date":"2004-10-04","text":"This card can trigger on itself being returned to a player's hand."}],"foreignNames":[{"name":"Dévotion dévoyée","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185203&type=card","language":"French","multiverseid":185203},{"name":"Verkrümmte Anbetung","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183131&type=card","language":"German","multiverseid":183131}],"printings":["PLS","8ED"],"originalText":"Whenever a permanent is returned to a player's hand, that player discards a card from his or her hand.","originalType":"Enchantment","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"dec949619370828dd5e52add8d3d561c23b00802"},{"name":"Heroic Defiance","manaCost":"{1}{W}","cmc":2,"colors":["White"],"type":"Enchantment — Aura","types":["Enchantment"],"subtypes":["Aura"],"rarity":"Common","set":"PLS","text":"Enchant creature\nEnchanted creature gets +3/+3 unless it shares a color with the most common color among all permanents or a color tied for most common.","flavor":"\"Wear courage as your armor. Wield honor as your blade.\"\n—Gerrard","artist":"Terese Nielsen","number":"6","layout":"normal","multiverseid":26420,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26420&type=card","foreignNames":[{"name":"Résistance héroïque","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185152&type=card","language":"French","multiverseid":185152},{"name":"Heldenhafter Widerstand","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183080&type=card","language":"German","multiverseid":183080}],"printings":["PLS"],"originalText":"Enchanted creature gets +3/+3 unless it shares a color with the most common color among all permanents or a color tied for most common.","originalType":"Enchant Creature","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"eb523a17839793219840c311ff89a910b322ef30"},{"name":"Stormscape Familiar","manaCost":"{1}{U}","cmc":2,"colors":["Blue"],"type":"Creature — Bird","types":["Creature"],"subtypes":["Bird"],"rarity":"Common","set":"PLS","text":"Flying\nWhite spells and black spells you cast cost {1} less to cast.","flavor":"Each Stormscape apprentice must hand-raise and train an owl before attaining the rank of master.","artist":"Heather Hudson","number":"36","power":"1","toughness":"1","layout":"normal","multiverseid":25616,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=25616&type=card","rulings":[{"date":"2004-10-04","text":"If a spell is both white and black, you pay {1} less, not {2} less."},{"date":"2004-10-04","text":"The generic X cost is still considered generic even if there is a requirement that a specific color be used for it. For example, \"only black mana can be spent this way\". This distinction is important for effects which reduce the generic portion of a spell's cost."},{"date":"2004-10-04","text":"This can lower the cost to zero, but not below zero."},{"date":"2004-10-04","text":"The effect is cumulative."},{"date":"2004-10-04","text":"The lower cost is not optional like with some other cost reducers."},{"date":"2004-10-04","text":"Can never affect the colored part of the cost."},{"date":"2004-10-04","text":"If this card is sacrificed to pay part of a spell's cost, the cost reduction still applies."}],"foreignNames":[{"name":"Familier orageosophe","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185182&type=card","language":"French","multiverseid":185182},{"name":"Vertrauter des Sturmpfads","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183110&type=card","language":"German","multiverseid":183110}],"printings":["PLS","TSB"],"originalText":"Flying\nWhite spells and black spells you play cost {1} less to play.","originalType":"Creature — Bird","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Time Spiral Block","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"e968cfcd63cd5d364893af962ad79a5b94e70645"},{"name":"Quirion Explorer","manaCost":"{1}{G}","cmc":2,"colors":["Green"],"type":"Creature — Elf Druid Scout","types":["Creature"],"subtypes":["Elf","Druid","Scout"],"rarity":"Common","set":"PLS","text":"{T}: Add to your mana pool one mana of any color that a land an opponent controls could produce.","flavor":"Fight with a friend at your back, steel in your hands, and magic in your veins.\n—Quirion creed","artist":"Ron Spears","number":"90","power":"1","toughness":"1","layout":"normal","multiverseid":26348,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26348&type=card","rulings":[{"date":"2008-08-01","text":"If the opponent only has lands that produce colorless or no mana, this card's ability can still be activated; it just won't produce any mana."}],"foreignNames":[{"name":"Exploratrice quirionaise","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185237&type=card","language":"French","multiverseid":185237},{"name":"Quirion-Kundschafter","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183165&type=card","language":"German","multiverseid":183165}],"printings":["PLS","C16"],"originalText":"{T}: Add to your mana pool one mana of any color that a land an opponent controls could produce.","originalType":"Creature — Elf","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"1f84492496aa9789243c21ff9aa263b892c2cba0"},{"name":"Sea Snidd","manaCost":"{4}{U}","cmc":5,"colors":["Blue"],"type":"Creature — Beast","types":["Creature"],"subtypes":["Beast"],"rarity":"Common","set":"PLS","text":"{T}: Target land becomes the basic land type of your choice until end of turn.","flavor":"It always has the home-turf advantage.","artist":"Chippy","number":"31","power":"3","toughness":"3","layout":"normal","multiverseid":26362,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26362&type=card","foreignNames":[{"name":"Snide marin","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185177&type=card","language":"French","multiverseid":185177},{"name":"See-Snidd","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183105&type=card","language":"German","multiverseid":183105}],"printings":["PLS"],"originalText":"{T}: Target land's type becomes the basic land type of your choice until end of turn.","originalType":"Creature — Beast","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"7b498cc880fc4e23e36008c138b1b2b39277e22d"},{"name":"Honorable Scout","manaCost":"{W}","cmc":1,"colors":["White"],"type":"Creature — Human Soldier Scout","types":["Creature"],"subtypes":["Human","Soldier","Scout"],"rarity":"Common","set":"PLS","text":"When Honorable Scout enters the battlefield, you gain 2 life for each black and/or red creature target opponent controls.","flavor":"\"I've faced your kind before. This time I'm ready for you.\"","artist":"Mike Ploog","number":"8","power":"1","toughness":"1","layout":"normal","multiverseid":26263,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26263&type=card","rulings":[{"date":"2004-10-04","text":"If a creature is both black and red, you gain 2 life, not 4."}],"foreignNames":[{"name":"Éclaireur honorable","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185154&type=card","language":"French","multiverseid":185154},{"name":"Ehrbarer Späher","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183082&type=card","language":"German","multiverseid":183082}],"printings":["PLS"],"originalText":"When Honorable Scout comes into play, you gain 2 life for each black and/or red creature target opponent controls.","originalType":"Creature — Soldier","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"42599e10b5d6e2fccfad2f59993c257c619e0016"},{"name":"Heroic Defiance","manaCost":"{1}{W}","cmc":2,"colors":["White"],"type":"Enchantment — Aura","types":["Enchantment"],"subtypes":["Aura"],"rarity":"Common","set":"PLS","text":"Enchant creature\nEnchanted creature gets +3/+3 unless it shares a color with the most common color among all permanents or a color tied for most common.","flavor":"\"Wear courage as your armor. Wield honor as your blade.\"\n—Gerrard","artist":"Terese Nielsen","number":"6","layout":"normal","multiverseid":26420,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26420&type=card","foreignNames":[{"name":"Résistance héroïque","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185152&type=card","language":"French","multiverseid":185152},{"name":"Heldenhafter Widerstand","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183080&type=card","language":"German","multiverseid":183080}],"printings":["PLS"],"originalText":"Enchanted creature gets +3/+3 unless it shares a color with the most common color among all permanents or a color tied for most common.","originalType":"Enchant Creature","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"eb523a17839793219840c311ff89a910b322ef30"},{"name":"Malicious Advice","manaCost":"{X}{U}{B}","cmc":2,"colors":["Blue","Black"],"type":"Instant","types":["Instant"],"rarity":"Common","set":"PLS","text":"Tap X target artifacts, creatures, and/or lands. You lose X life.","flavor":"\"Rule through fear, Darigaaz. Only if the dragons feel terror will they truly be your servants.\"\n—Tevesh Szat","artist":"Glen Angus","number":"114","layout":"normal","multiverseid":25870,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=25870&type=card","foreignNames":[{"name":"Conseil malveillant","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185262&type=card","language":"French","multiverseid":185262},{"name":"Boshafter Ratschlag","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183190&type=card","language":"German","multiverseid":183190}],"printings":["PLS"],"originalText":"Tap X target artifacts, creatures, and/or lands. You lose X life.","originalType":"Instant","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"ea330138d494cbf0d577c5cf2af1b4682d10d65d"},{"name":"Falling Timber","manaCost":"{2}{G}","cmc":3,"colors":["Green"],"type":"Instant","types":["Instant"],"rarity":"Common","set":"PLS","text":"Kicker—Sacrifice a land. (You may sacrifice a land in addition to any other costs as you cast this spell.)\nPrevent all combat damage target creature would deal this turn. If Falling Timber was kicked, prevent all combat damage another target creature would deal this turn.","artist":"Eric Peterson","number":"79","layout":"normal","multiverseid":25945,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=25945&type=card","rulings":[{"date":"2004-10-04","text":"You pick a second target only if you choose to kick Falling Timber."}],"foreignNames":[{"name":"Bois abattu","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185226&type=card","language":"French","multiverseid":185226},{"name":"Stürzende Bäume","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183154&type=card","language":"German","multiverseid":183154}],"printings":["PLS"],"originalText":"Kicker—Sacrifice a land. (You may sacrifice a land in addition to any other costs as you play this spell.)\nPrevent all combat damage target creature would deal this turn. If you paid the kicker cost, prevent all combat damage another target creature would deal this turn.","originalType":"Instant","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"86786613184732d87c250b679b8b91ff9d6190e3"},{"name":"Quirion Explorer","manaCost":"{1}{G}","cmc":2,"colors":["Green"],"type":"Creature — Elf Druid Scout","types":["Creature"],"subtypes":["Elf","Druid","Scout"],"rarity":"Common","set":"PLS","text":"{T}: Add to your mana pool one mana of any color that a land an opponent controls could produce.","flavor":"Fight with a friend at your back, steel in your hands, and magic in your veins.\n—Quirion creed","artist":"Ron Spears","number":"90","power":"1","toughness":"1","layout":"normal","multiverseid":26348,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26348&type=card","rulings":[{"date":"2008-08-01","text":"If the opponent only has lands that produce colorless or no mana, this card's ability can still be activated; it just won't produce any mana."}],"foreignNames":[{"name":"Exploratrice quirionaise","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185237&type=card","language":"French","multiverseid":185237},{"name":"Quirion-Kundschafter","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183165&type=card","language":"German","multiverseid":183165}],"printings":["PLS","C16"],"originalText":"{T}: Add to your mana pool one mana of any color that a land an opponent controls could produce.","originalType":"Creature — Elf","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"1f84492496aa9789243c21ff9aa263b892c2cba0"},{"name":"Kavu Recluse","manaCost":"{2}{R}","cmc":3,"colors":["Red"],"type":"Creature — Kavu","types":["Creature"],"subtypes":["Kavu"],"rarity":"Common","set":"PLS","text":"{T}: Target land becomes a Forest until end of turn.","flavor":"Few ever see this particular breed of kavu, which hides in forests of its own making.","artist":"Aaron Boyd","number":"64","power":"2","toughness":"2","layout":"normal","multiverseid":26268,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26268&type=card","foreignNames":[{"name":"Kavru reclus","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185210&type=card","language":"French","multiverseid":185210},{"name":"Einsiedler-Kavu","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183138&type=card","language":"German","multiverseid":183138}],"printings":["PLS"],"originalText":"{T}: Target land becomes a forest until end of turn.","originalType":"Creature — Kavu","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Freeform","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Prismatic","legality":"Legal"},{"format":"Singleton 100","legality":"Legal"},{"format":"Tribal Wars Legacy","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"455560f7ce86542b514072976cf0e178bfa65da9"},{"name":"Maggot Carrier","manaCost":"{B}","cmc":1,"colors":["Black"],"type":"Creature — Zombie","types":["Creature"],"subtypes":["Zombie"],"rarity":"Common","set":"PLS","text":"When Maggot Carrier enters the battlefield, each player loses 1 life.","flavor":"\"The mere sight of our undead allies sickens me. What unholy bargain have you struck?\"\n—Grizzlegom, to Agnate","artist":"Ron Spencer","number":"45","power":"1","toughness":"1","layout":"normal","multiverseid":26266,"imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=26266&type=card","foreignNames":[{"name":"Propagateur de vermine","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=185191&type=card","language":"French","multiverseid":185191},{"name":"Madenwirt","imageUrl":"http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=183119&type=card","language":"German","multiverseid":183119}],"printings":["PLS","8ED"],"originalText":"When Maggot Carrier comes into play, each player loses 1 life.","originalType":"Creature — Zombie","legalities":[{"format":"Commander","legality":"Legal"},{"format":"Invasion Block","legality":"Legal"},{"format":"Legacy","legality":"Legal"},{"format":"Modern","legality":"Legal"},{"format":"Vintage","legality":"Legal"}],"id":"53b1616f17fe6a0205ee0f1525b4171f0e30b2ad"}]}`))

			cards, err := SetCode("PLS").GenerateBooster()
			Convey("There should be no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("The result should contain some cards", func() {
				So(cards, ShouldContainCard, "Planeswalker's Fury")
				So(cards, ShouldContainCard, "Strafe")
				So(cards, ShouldContainCard, "Crosis's Catacombs")
				So(cards, ShouldContainCard, "Warped Devotion")
				So(cards, ShouldContainCard, "Heroic Defiance")
				So(cards, ShouldContainCard, "Stormscape Familiar")
				So(cards, ShouldContainCard, "Honorable Scout")
			})
		})
	})
}

func Test_BoosterContentString(t *testing.T) {
	Convey("When converting a BoosterContent to a string", t, func() {
		Convey("A single type should be the type itself", func() {
			bc := BoosterContent{"Common"}
			So(bc.String(), ShouldEqual, "Common")
		})
		Convey("If there are multiple possible types they should be split by |", func() {
			bc := BoosterContent{"Common", "Rare"}
			So(bc.String(), ShouldEqual, "Common|Rare")
		})
	})
}

func Test_SetQuery(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	Convey("With a new SetQuery", t, func() {
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=Planeshift&page=1&pageSize=500",
			NewStringResponderWithHeader(200, `{"sets":[{"code":"PLS","name":"Planeshift","type":"expansion","border":"black","booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"releaseDate":"2001-02-05","gathererCode":"PS","magicCardsInfoCode":"ps","block":"Invasion"}]}`,
				map[string]string{
					"Total-Count": "1337",
				}))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/PLS",
			httpmock.NewStringResponder(200, `{"set":{"code":"PLS","name":"Planeshift","type":"expansion","border":"black","booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"releaseDate":"2001-02-05","gathererCode":"PS","magicCardsInfoCode":"ps","block":"Invasion"}}`))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/FOO_BAR",
			httpmock.NewStringResponder(200, `{"sets":[]}`))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/network_issue",
			httpmock.NewErrorResponder(errors.New("Network Issue")))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/server_issue",
			httpmock.NewStringResponder(500, `{"status": "500", "error":"Internal server error"}`))
		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/invalid_json",
			httpmock.NewStringResponder(200, `{"sets":`))
		qry := NewSetQuery()

		Convey("When searching by name", func() {
			qry = qry.Where(SetName, "Planeshift")

			Convey("a copy should make no difference", func() {
				cpy := qry.Copy()
				So(cpy, ShouldResemble, qry)
				So(cpy, ShouldNotEqual, qry)
			})

			sets, totalCount, err := qry.Page(1)

			So(err, ShouldBeNil)
			So(totalCount, ShouldEqual, 1337)
			So(sets, ShouldHaveLength, 1)

			set := sets[0]
			So(set.Name, ShouldEqual, "Planeshift")
			So(set.String(), ShouldEqual, "Planeshift (PLS)")

			Convey("In case of errors", func() {
				Convey("Invalid json should be reported", func() {
					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=Planeshift&page=1&pageSize=500",
						httpmock.NewStringResponder(200, `{"sets":`))

					_, _, err := qry.Page(1)
					So(err, ShouldNotBeNil)
				})

				Convey("If Total-Count is not a number", func() {
					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=Planeshift&page=1&pageSize=500",
						NewStringResponderWithHeader(200, `{"sets":[{"code":"PLS","name":"Planeshift","type":"expansion","border":"black","booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"releaseDate":"2001-02-05","gathererCode":"PS","magicCardsInfoCode":"ps","block":"Invasion"}]}`,
							map[string]string{
								"Total-Count": "two",
							}))
					_, _, err := qry.Page(1)
					So(err, ShouldNotBeNil)
				})
			})

			Convey("fetching the same set by its id should result in the same values", func() {
				other, err := set.SetCode.Fetch()
				So(err, ShouldBeNil)
				So(other, ShouldResemble, set)
			})
			Convey("fetching an invalid setcode should return an error", func() {
				_, err := SetCode("FOO_BAR").Fetch()
				So(err, ShouldNotBeNil)
			})
			Convey("when we have network issues, there should also be an error", func() {
				_, err := SetCode("network_issue").Fetch()
				So(err, ShouldNotBeNil)
			})
			Convey("when the server reports an error we should get a ServerError", func() {
				_, err := SetCode("server_issue").Fetch()
				So(err, ShouldNotBeNil)
				_, isServerError := err.(ServerError)
				So(isServerError, ShouldBeTrue)
			})
			Convey("when the server sends invalid json there should be an error", func() {
				_, err := SetCode("invalid_json").Fetch()
				So(err, ShouldNotBeNil)
			})

			Convey("with paging", func() {
				qry = NewSetQuery().Where(SetName, "n")

				httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=n",
					NewStringResponderWithHeader(200, `{"sets":[{"code":"LEA","name":"Limited Edition Alpha","type":"core","border":"black","mkm_id":1,"booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"mkm_name":"Alpha","releaseDate":"1993-08-05","gathererCode":"1E","magicCardsInfoCode":"al"}]}`,
						map[string]string{
							"Link": `<https://api.magicthegathering.io/v1/sets?name=n&page=2>; rel="last", <https://api.magicthegathering.io/v1/sets?name=n&page=2>; rel="next"`,
						}))

				httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=n&page=2",
					httpmock.NewStringResponder(200, `{"sets":[{"code":"LEB","name":"Limited Edition Beta","type":"core","border":"black","mkm_id":2,"booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"mkm_name":"Beta","releaseDate":"1993-10-01","gathererCode":"2E","magicCardsInfoCode":"be"}]}`))

				cards, err := qry.All()
				So(err, ShouldBeNil)
				So(cards, ShouldHaveLength, 2)

				Convey("If one of the following pages cause problems they should be reported", func() {
					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=n&page=2",
						httpmock.NewErrorResponder(errors.New("Network Issue")))
					_, err := qry.All()
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
