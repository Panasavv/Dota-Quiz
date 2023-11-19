package questions

import (
	"encoding/json"
	"fmt"
	"interfaces"
	"io"
	"math/rand"
	"os"
)

func GetQuestion(s string, onepointers int, twopointers int, threepointers int) (interfaces.AllPoints, error) {
	questionsReturned := interfaces.AllPoints{}
	jsonFile, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Succesfully opened json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	for i := 0; i < onepointers; i++ {

		var onePointers interfaces.OnePointers
		json.Unmarshal(byteValue, &onePointers)
		k := rand.Intn(8)
		questionsReturned.OnePointers = append(questionsReturned.OnePointers, onePointers.OnePointers[k])
	}
	for i := 0; i < twopointers; i++ {

		var twoPointers interfaces.TwoPointers
		json.Unmarshal(byteValue, &twoPointers)
		k := rand.Intn(8)
		questionsReturned.TwoPointers = append(questionsReturned.TwoPointers, twoPointers.TwoPointers[k])
	}
	for i := 0; i < threepointers; i++ {

		var threePointers interfaces.ThreePointers
		json.Unmarshal(byteValue, &threePointers)
		k := rand.Intn(8)
		questionsReturned.ThreePointers = append(questionsReturned.ThreePointers, threePointers.ThreePointers[k])
	}
	return questionsReturned, nil
}
