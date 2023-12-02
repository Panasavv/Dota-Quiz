package interfaces

type OnePointers struct {
	OnePointers []Question `json:"onepointers"`
}

type TwoPointers struct {
	TwoPointers []Question `json:"twopointers"`
}

type ThreePointers struct {
	ThreePointers []Question `json:"threepointers"`
}

type Question struct {
	Number   string `json:"number"`
	Text     string `json:"text"`
	Answer   string `json:"answer"`
	Points   string `json:"points"`
	Fifty    string `json:"fifty"`
	Qtype    string `json:"qtype"`
	Comments string `json:"comments"`
	IsPlayed bool   `json:"isplayed"`
}

type AllQuestions struct {
	Category1 AllPoints
	Category2 AllPoints
	Category3 AllPoints
	Category4 AllPoints
	Category5 AllPoints
	Category6 AllPoints
	Category7 AllPoints
	Category8 AllPoints
}

type AllPoints struct {
	OnePointers   []Question
	TwoPointers   []Question
	ThreePointers []Question
}

type QCategory struct {
	Category1 Category
	Category2 Category
	Category3 Category
	Category4 Category
	Category5 Category
	Category6 Category
	Category7 Category
	Category8 Category
}

type Category struct {
	IsPicked      bool
	OnePointers   int
	TwoPointers   int
	ThreePointers int
}

type User struct {
	Name      string
	Points    int
	Lifelines Helpers
}

type Game struct {
	Date               string       `json:"date"`
	GameID             string       `json:"gameID"`
	Participants       []User       `json:"participants"`
	ActivePlayer       User         `json:"activeplayer"`
	StartingQuestions  AllQuestions `json:"startingQuestions"`
	QuestionPicked     Question     `json:"questionsPicked"`
	QuestionsRemaining AllQuestions `json:"questionsRemaining"`
	TotalQuestions     int          `json:"totalquestions"`
	Winner             User         `json:"winner"`
}
type Helpers struct {
	Fifty  bool
	Phone  bool
	Double bool
}

type QuestionResponse struct {
	Category1 CategoryResponse `json:"category1"`
	Category2 CategoryResponse `json:"category2"`
	Category3 CategoryResponse `json:"category3"`
	Category4 CategoryResponse `json:"category4"`
	Category5 CategoryResponse `json:"category5"`
	Category6 CategoryResponse `json:"category6"`
	Category7 CategoryResponse `json:"category7"`
	Category8 CategoryResponse `json:"category8"`
}

type CategoryResponse struct {
	CatString string `json:"catstring"`
	CatCount  int    `json:"catcount"`
}

type UniResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
	Status  string `json:"status"`
}

var QuestionsPlayed = QCategory{
	Category1: Category{
		IsPicked:      true,
		OnePointers:   1,
		TwoPointers:   1,
		ThreePointers: 1,
	},
	Category2: Category{
		IsPicked:      true,
		OnePointers:   0,
		TwoPointers:   2,
		ThreePointers: 0,
	},
	Category3: Category{
		IsPicked:      true,
		OnePointers:   0,
		TwoPointers:   0,
		ThreePointers: 2,
	},
	Category4: Category{
		IsPicked:      true,
		OnePointers:   1,
		TwoPointers:   1,
		ThreePointers: 1,
	},
	Category5: Category{
		IsPicked:      true,
		OnePointers:   1,
		TwoPointers:   1,
		ThreePointers: 1,
	},
	Category6: Category{
		IsPicked:      true,
		OnePointers:   3,
		TwoPointers:   0,
		ThreePointers: 0,
	},
	Category7: Category{
		IsPicked:      true,
		OnePointers:   0,
		TwoPointers:   0,
		ThreePointers: 2,
	},
	Category8: Category{
		IsPicked:      true,
		OnePointers:   0,
		TwoPointers:   2,
		ThreePointers: 0,
	},
}

var ImagesTBC1 = [17]string{
	//"absorb",
	"annul",
	//"call of the herd",
	//"clone",
	"deathmark",
	//"enlightened tutor",
	"fevered visions",
	"gatekeeper of malakir",
	//"goblin welder",
	"grapple with the past",
	"hashep oasis",
	"lightning mauler",
	"lotleth troll",
	//"opt",
	"rampaging baloths",
	"remove soul",
	"ricochet trap",
	"spark elemental",
	"sunblast angel",
	"titanic ultimatum",
	"troll ascetic",
	//"unsummon",
	"wargate",
	"wee dragonauts",
	//"wirewood symbiote",
}

var ImagesTBC2 = [9]string{
	"Academy Rector",
	//"azors gateway",
	"bee sting",
	"elvish piper",
	//"cancel",
	"harrow",
	//"deathcap cultivator",
	//"echoing calm",
	"icy manipulator",
	//"feather the redeemed",
	//"gnarlid pack",
	//"honored hierarch",
	//"iroas god of victory",
	"memory lapse",
	//"mage ring network",
	//"make a stand",
	//"merciless eviction",
	//"regenerate",
	//"shimian specter",
	//"terror",
	//"unmake",
	"predict",
	//"yavimaya enchantress",
	"time stretch",
	"zap",
}

var ImagesTBC3 = [21]string{
	//"deluge",
	//"insist",
	//"lightning dart",
	//"winters grasp",
	"beacon of immortality",
	"borderland ranger",
	"deathsprout",
	"dimir aqueduct",
	"doomskar titan",
	"early frost",
	"favor of the overbeing",
	"gibbering descent",
	"hamlet captain",
	"heartwood storyteller",
	"jedit ojanen of efrava",
	"magus of the mirror",
	"mirror mockery",
	"pulse tracker",
	"recross the paths",
	"revitalize",
	"scythe tiger",
	"tempest of light",
	"viviens arkbow",
	"silumgar assassin",
	"zombie goliath",
}
