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
