package main

import (
	"game"
	"interfaces"
)

func main() {
	questionsPlayed := &interfaces.QCategory{
		Category1: interfaces.Category{
			IsPicked:      true,
			OnePointers:   1,
			TwoPointers:   1,
			ThreePointers: 1,
		},
		Category2: interfaces.Category{
			IsPicked:      true,
			OnePointers:   0,
			TwoPointers:   2,
			ThreePointers: 0,
		},
		Category3: interfaces.Category{
			IsPicked:      true,
			OnePointers:   0,
			TwoPointers:   0,
			ThreePointers: 2,
		},
		Category4: interfaces.Category{
			IsPicked:      true,
			OnePointers:   1,
			TwoPointers:   1,
			ThreePointers: 1,
		},
		Category5: interfaces.Category{
			IsPicked:      true,
			OnePointers:   1,
			TwoPointers:   1,
			ThreePointers: 1,
		},
		Category6: interfaces.Category{
			IsPicked:      true,
			OnePointers:   3,
			TwoPointers:   0,
			ThreePointers: 0,
		},
		Category7: interfaces.Category{
			IsPicked:      true,
			OnePointers:   0,
			TwoPointers:   0,
			ThreePointers: 2,
		},
		Category8: interfaces.Category{
			IsPicked:      true,
			OnePointers:   0,
			TwoPointers:   2,
			ThreePointers: 0,
		},
	}
	game.StartGame(*questionsPlayed)

}
