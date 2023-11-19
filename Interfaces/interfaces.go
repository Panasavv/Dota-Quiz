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

/*
	type QuestionMap struct {
		Category1 QuestionMapPoints
		Category2 QuestionMapPoints
		Category3 QuestionMapPoints
		Category4 QuestionMapPoints
		Category5 QuestionMapPoints
		Category6 QuestionMapPoints
		Category7 QuestionMapPoints
		Category8 QuestionMapPoints
	}

	type QuestionMapPoints struct {
		OnePointers   []bool
		TwoPointers   []bool
		ThreePointers []bool
	}
*/
type Helpers struct {
	Fifty  bool
	Phone  bool
	Double bool
}
