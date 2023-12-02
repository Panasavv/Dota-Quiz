package questions

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"interfaces"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func PreviewImage(c *gin.Context) {
	// Assuming you have an image URL or path, you can fetch the image
	id, ok := c.GetQuery("id")
	if !ok {
		unires := interfaces.UniResponse[string]{
			Message: "id is missing",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}
	imagePath := "./Questions/card images/3/" + id

	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
		return
	}

	// Determine the content type based on the file extension
	contentType := http.DetectContentType(imageBytes)

	// Set the appropriate headers for image response
	c.Header("Content-Type", contentType)
	c.Header("Content-Length", fmt.Sprint(len(imageBytes)))

	// Serve the image content
	c.Data(http.StatusOK, contentType, imageBytes)
}

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

func CropImages(imgName string) {
	originalImageFile, err := os.Open(string("./Questions/card images/3/" + imgName + ".jpg"))
	if err != nil {
		panic(err)
	}
	defer originalImageFile.Close()

	originalImage, err := jpeg.Decode(originalImageFile)
	if err != nil {
		panic(err)
	}
	bounds := originalImage.Bounds()
	width := bounds.Dx()
	cropSize := image.Rect(0, 0, width/2+85, width/2+30)
	//cropSize = cropSize.Add(image.Point{30, 35}) //old border
	cropSize = cropSize.Add(image.Point{25, 43}) //new border
	croppedImage := originalImage.(SubImager).SubImage(cropSize)
	croppedImage2 := resize.Resize(800, 0, croppedImage, resize.Lanczos3)

	croppedImageFile, err := os.Create(string("./Questions/card images/3/" + imgName + "_question.jpg"))
	if err != nil {
		panic(err)
	}
	defer croppedImageFile.Close()
	if err := jpeg.Encode(croppedImageFile, croppedImage2, nil); err != nil {
		panic(err)
	}
}
