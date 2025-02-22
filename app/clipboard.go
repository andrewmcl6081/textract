package app

import (
	"fmt"
	"log"

	"golang.design/x/clipboard"
)

func InitClipboard() {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		log.Fatalln("Clipboard package not ready for use:", err)
	}

	log.Println("Clipboard initialized.")
}

func Copy(s *State) {
	if s.imageBuffer.Len() != 0 {
		extractedText, err := Ocr(s)
		if err != nil {
			log.Println("Failed to OCR text:", err)
			return
		}

		if extractedText == "" {
			log.Println("OCR returned empty text.")
			return nil
		}

		clipboard.Write(clipboard.FmtText, []byte(extractedText))
		log.Println("Text coppied to clipboard.")
		return nil
	} else {
		log.Println("Image buffer empty. Skipping clipboard copy.")
	}
}
