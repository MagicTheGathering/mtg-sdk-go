package mtg

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/jarcoal/httpmock"
)

func Test_GetTypes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	Convey("Fetching all types", t, func() {
		Convey("If the response is correct", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/types",
				httpmock.NewStringResponder(200, `{"types":["Artifact","Conspiracy","Creature","Enchantment","Instant","Land","Phenomenon","Plane","Planeswalker","Scheme","Sorcery","Tribal","Vanguard"]}`))

			types, err := GetTypes()
			So(err, ShouldBeNil)
			So(types, ShouldContain, "Enchantment")
		})
		Convey("If we have network issues or such things", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/types",
				httpmock.NewErrorResponder(errors.New("Network Issues")))
			types, err := GetTypes()

			So(types, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("If we get a error from the server", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/types",
				httpmock.NewStringResponder(500, `{"status": "500", "error":"Internal server error"}`))

			Convey("There should be an error", func() {
				_, err := GetTypes()
				So(err, ShouldNotBeNil)

				Convey("and it should be a ServerError", func() {
					_, isServerError := err.(ServerError)
					So(isServerError, ShouldBeTrue)
				})
			})
		})
		Convey("If we get invalid json", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/types",
				httpmock.NewStringResponder(200, `{"types":"YES"}`))
			Convey("There should be an error", func() {
				_, err := GetTypes()
				So(err, ShouldNotBeNil)

				Convey("and it should be no ServerError", func() {
					_, isServerError := err.(ServerError)
					So(isServerError, ShouldBeFalse)
				})
			})
		})
	})
}

func Test_GetSuperTypes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	Convey("Fetching all supertypes", t, func() {
		Convey("If the response is correct", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/supertypes",
				httpmock.NewStringResponder(200, `{"supertypes":["Basic","Legendary","Ongoing","Snow","World"]}`))

			types, err := GetSuperTypes()
			So(err, ShouldBeNil)
			So(types, ShouldContain, "Snow")
		})
		Convey("If we have network issues or such things", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/supertypes",
				httpmock.NewErrorResponder(errors.New("Network Issues")))
			types, err := GetSuperTypes()

			So(types, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("If we get a error from the server", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/supertypes",
				httpmock.NewStringResponder(500, `{"status": "500", "error":"Internal server error"}`))

			_, err := GetSuperTypes()
			So(err, ShouldNotBeNil)

			_, isServerError := err.(ServerError)
			So(isServerError, ShouldBeTrue)
		})
		Convey("If we get invalid json", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/supertypes",
				httpmock.NewStringResponder(200, `{"supertypes":"YES"}`))

			_, err := GetSuperTypes()
			So(err, ShouldNotBeNil)

			_, isServerError := err.(ServerError)
			So(isServerError, ShouldBeFalse)
		})
	})
}

func Test_GetSubTypes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	Convey("Fetching all subtypes", t, func() {
		Convey("If the response is correct", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/subtypes",
				httpmock.NewStringResponder(200, `{"subtypes":["Advisor","Aetherborn","Ajani","Alara","Ally","Angel","Antelope","Ape","Arcane","Archer","Archon","Arkhos","Arlinn","Artificer","Ashiok","Assassin","Assembly-Worker","Atog","Aura","Aurochs","Avatar","Azgol","Baddest,","Badger","Barbarian","Basilisk","Bat","Bear","Beast","Beeble","Belenon","Berserker","Biggest,","Bird","Boar","Bolas","Bolas’s Meditation Realm","Bringer","Brushwagg","Bureaucrat","Camel","Carrier","Cat","Centaur","Cephalid","Chandra","Chicken","Child","Chimera","Clamfolk","Cleric","Cockatrice","Construct","Cow","Crab","Crocodile","Curse","Cyclops","Dack","Daretti","Dauthi","Demon","Desert","Designer","Devil","Dinosaur","Djinn","Dominaria","Domri","Donkey","Dovin","Dragon","Drake","Dreadnought","Drone","Druid","Dryad","Dwarf","Efreet","Egg","Elder","Eldrazi","Elemental","Elephant","Elf","Elk","Elspeth","Elves","Equilor","Equipment","Ergamon","Etiquette","Eye","Fabacin","Faerie","Ferret","Fish","Flagbearer","Forest","Fortification","Fox","Freyalise","Frog","Fungus","Gamer","Gargoyle","Garruk","Gate","Giant","Gideon","Gnome","Goat","Goblin","Goblins","God","Golem","Gorgon","Gremlin","Griffin","Gus","Hag","Harpy","Hellion","Hero","Hippo","Hippogriff","Homarid","Homunculus","Horror","Horse","Hound","Human","Hydra","Hyena","Igpay","Illusion","Imp","Incarnation","Innistrad","Insect","Iquatana","Ir","Island","Jace","Jellyfish","Juggernaut","Kaldheim","Kamigawa","Karn","Kavu","Kaya","Kephalai","Kiora","Kirin","Kithkin","Knight","Kobold","Kolbahan","Kor","Koth","Kraken","Kyneth","Lady","Lair","Lamia","Lammasu","Leech","Legend","Leviathan","Lhurgoyf","Licid","Liliana","Lizard","Locus","Lord","Lorwyn","Manticore","Masticore","Mercadia","Mercenary","Merfolk","Metathran","Mime","Mine","Minion","Minotaur","Mirrodin","Moag","Mole","Monger","Mongoose","Mongseng","Monk","Monkey","Moonfolk","Mountain","Mummy","Muraganda","Mutant","Myr","Mystic","Naga","Nahiri","Narset","Nastiest,","Nautilus","Nephilim","New Phyrexia","Nightmare","Nightstalker","Ninja","Nissa","Nixilis","Noggle","Nomad","Nymph","Octopus","Ogre","Ooze","Orc","Orgg","Ouphe","Ox","Oyster","Paratrooper","Pegasus","Pest","Phelddagrif","Phoenix","Phyrexia","Pilot","Pirate","Plains","Plant","Power-Plant","Praetor","Processor","Proper","Rabbit","Rabiah","Ral","Rat","Rath","Ravnica","Rebel","Reflection","Regatha","Rhino","Rigger","Rogue","Sable","Saheeli","Salamander","Samurai","Saproling","Sarkhan","Satyr","Scarecrow","Scion","Scorpion","Scout","Segovia","Serpent","Serra’s Realm","Shade","Shadowmoor","Shaman","Shandalar","Shapeshifter","Sheep","Ship","Shrine","Siren","Skeleton","Slith","Sliver","Slug","Snake","Soldier","Soltari","Sorin","Spawn","Specter","Spellshaper","Sphinx","Spider","Spike","Spirit","Sponge","Squid","Squirrel","Starfish","Surrakar","Swamp","Tamiyo","Teferi","Tezzeret","Thalakos","The","Thopter","Thrull","Tibalt","Tower","Townsfolk","Trap","Treefolk","Troll","Turtle","Ugin","Ulgrotha","Unicorn","Urza’s","Valla","Vampire","Vedalken","Vehicle","Venser","Viashino","Volver","Vraska","Vryn","Waiter","Wall","Warrior","Weird","Werewolf","Whale","Wildfire","Wizard","Wolf","Wolverine","Wombat","Worm","Wraith","Wurm","Xenagos","Xerex","Yeti","Zendikar","Zombie","Zubera"]}`))

			types, err := GetSubTypes()
			So(err, ShouldBeNil)
			So(types, ShouldContain, "Mole")
			So(types, ShouldContain, "Angel")
		})
		Convey("If we have network issues or such things", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/subtypes",
				httpmock.NewErrorResponder(errors.New("Network Issues")))
			types, err := GetSubTypes()

			So(types, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("If we get a error from the server", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/subtypes",
				httpmock.NewStringResponder(500, `{"status": "500", "error":"Internal server error"}`))

			_, err := GetSubTypes()
			So(err, ShouldNotBeNil)

			_, isServerError := err.(ServerError)
			So(isServerError, ShouldBeTrue)
		})
		Convey("If we get invalid json", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/subtypes",
				httpmock.NewStringResponder(200, `{"subtypes":"YES"}`))

			_, err := GetSubTypes()
			So(err, ShouldNotBeNil)

			_, isServerError := err.(ServerError)
			So(isServerError, ShouldBeFalse)
		})
	})
}

func Test_GetFormats(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	Convey("Fetching all formats", t, func() {
		Convey("If the response is correct", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/formats",
				httpmock.NewStringResponder(200, `{"formats":["Battle for Zendikar Block","Classic","Commander","Freeform","Ice Age Block","Innistrad Block","Invasion Block","Kaladesh Block","Kamigawa Block","Khans of Tarkir Block","Legacy","Lorwyn-Shadowmoor Block","Masques Block","Mirage Block","Mirrodin Block","Modern","Odyssey Block","Onslaught Block","Prismatic","Ravnica Block","Return to Ravnica Block","Scars of Mirrodin Block","Shadows over Innistrad Block","Shards of Alara Block","Singleton 100","Standard","Tempest Block","Theros Block","Time Spiral Block","Tribal Wars Legacy","Un-Sets","Urza Block","Vintage","Zendikar Block"]}`))

			types, err := GetFormats()
			So(err, ShouldBeNil)
			So(types, ShouldContain, "Singleton 100")
		})
		Convey("If we have network issues or such things", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/formats",
				httpmock.NewErrorResponder(errors.New("Network Issues")))
			types, err := GetFormats()

			So(types, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("If we get a error from the server", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/formats",
				httpmock.NewStringResponder(500, `{"status": "500", "error":"Internal server error"}`))

			_, err := GetFormats()
			So(err, ShouldNotBeNil)

			_, isServerError := err.(ServerError)
			So(isServerError, ShouldBeTrue)
		})
		Convey("If we get invalid json", func() {
			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/formats",
				httpmock.NewStringResponder(200, `{"formats":"YES"}`))

			_, err := GetFormats()
			So(err, ShouldNotBeNil)

			_, isServerError := err.(ServerError)
			So(isServerError, ShouldBeFalse)
		})
	})
}
