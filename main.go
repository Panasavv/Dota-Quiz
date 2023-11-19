package main

import (
	"game"
	"interfaces"
)

func main() {
	questionsPlayed := &interfaces.QCategory{
		Category1: interfaces.Category{
			IsPicked:      true,
			OnePointers:   0,
			TwoPointers:   2,
			ThreePointers: 3,
		},
		Category2: interfaces.Category{
			IsPicked:      true,
			OnePointers:   0,
			TwoPointers:   2,
			ThreePointers: 3,
		},
		Category3: interfaces.Category{
			IsPicked:      true,
			OnePointers:   1,
			TwoPointers:   0,
			ThreePointers: 1,
		},
		Category4: interfaces.Category{
			IsPicked:      false,
			OnePointers:   0,
			TwoPointers:   2,
			ThreePointers: 3,
		},
		Category5: interfaces.Category{
			IsPicked:      false,
			OnePointers:   0,
			TwoPointers:   2,
			ThreePointers: 3,
		},
		Category6: interfaces.Category{
			IsPicked:      false,
			OnePointers:   0,
			TwoPointers:   2,
			ThreePointers: 3,
		},
		Category7: interfaces.Category{
			IsPicked:      false,
			OnePointers:   0,
			TwoPointers:   2,
			ThreePointers: 3,
		},
		Category8: interfaces.Category{
			IsPicked:      false,
			OnePointers:   0,
			TwoPointers:   2,
			ThreePointers: 3,
		},
	}
	game.StartGame(*questionsPlayed)

}
