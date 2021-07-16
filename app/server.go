package main

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/qeesung/image2ascii/convert"
)

func saveToAscii(filename string) string {

	convertOptions := convert.DefaultOptions
	convertOptions.FitScreen = true
	convertOptions.Colored = false
	converter := convert.NewImageConverter()

	// Write ASCII file with uuid filename
	id := uuid.New()
	asciiFile, err := os.Create("./ascii/" + id.String())
	if err != nil {
		return "failed"
	}
	_, err = asciiFile.WriteString(converter.ImageFile2ASCIIString(filename, &convertOptions))
	if err != nil {
		return "failed"
	}

	// Return identifier
	return id.String()

}

func imageGetHandler(c *fiber.Ctx) error {

	// Read ASCII File directory
	// NOTE: Assumes valid directory!
	files, err := ioutil.ReadDir("./ascii")
	if err != nil {
		c.SendStatus(500)
		return c.SendString("500 Internal Server Error\n")
	}

	// Store ASCII file name strings
	var out []string
	for _, file := range files {
		out = append(out, file.Name())
	}

	return c.SendString(strings.Join(out, "\n") + "\n")

}

func imageIdGetHandler(c *fiber.Ctx) error {

	// Read ASCII PNG by uuid
	file, err := ioutil.ReadFile("./ascii/" + c.Params("id"))
	if err != nil {
		c.SendStatus(404)
		return c.SendString("404 Not Found")
	}

	return c.SendStream(bytes.NewReader([]byte(file)))

}

func imagePostHandler(c *fiber.Ctx) error {

	// Decode PNG POST body
	// NOTE: Assumes binary data!
	img, err := png.Decode(bytes.NewReader(c.Body()))
	if err != nil {
		c.SendStatus(415)
		return c.SendString("415 Unsupported Media Type\n")
	}

	// Save/overwrite a temporary image file to convert to ASCII
	// NOTE: Assumes valid upload directory!
	tempFile := "./uploads/tempFile.png"
	out, err := os.Create(tempFile)
	if err != nil {
		c.SendStatus(500)
		return c.SendString("500 Internal Server Error\n")
	}
	err = png.Encode(out, img)
	if err != nil {
		c.SendStatus(500)
		return c.SendString("500 Internal Server Error\n")
	}

	// Convert and save temp image to ASCII
	id := saveToAscii(tempFile)

	return c.SendString("ASCII ID: " + id + "\n")

}

func main() {

	app := fiber.New()

	app.Get("/images", imageGetHandler)
	app.Get("/images/:id", imageIdGetHandler)
	app.Post("/images", imagePostHandler)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
