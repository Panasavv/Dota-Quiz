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

type Helpers struct {
	Fifty  bool
	Phone  bool
	Double bool
}
