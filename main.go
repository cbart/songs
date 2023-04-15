package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func replaceSharpWithSuperscript(text string) string {
	superscriptSharp := "♯"
	return strings.ReplaceAll(text, "#", superscriptSharp)
}

func main() {
	inputFile := "song_lyrics.txt"
	outputFile := "song_lyrics.pdf"

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("DejaVu", "", "DejaVuSans.ttf")
	pdf.AddUTF8Font("DejaVu", "B", "DejaVuSans-Bold.ttf")

	pdf.AddPage()
	pdf.SetFont("DejaVu", "", 12)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		if len(parts) > 1 {
			lyrics := strings.TrimSpace(parts[0])
			chords := strings.TrimSpace(parts[1])

			chords = replaceSharpWithSuperscript(chords)

			pdf.CellFormat(140, 10, lyrics, "", 0, "L", false, 0, "")
			pdf.CellFormat(50, 10, chords, "", 1, "R", false, 0, "")
		} else {
			pdf.CellFormat(0, 10, line, "", 1, "", false, 0, "")
		}
	}

	err = pdf.OutputFileAndClose(outputFile)
	if err != nil {
		fmt.Println("Error creating PDF:", err)
		return
	}
	fmt.Println("PDF created:", outputFile)
}
