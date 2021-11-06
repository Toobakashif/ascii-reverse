package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)
var rs [9]string
var px string


func main() {
	var banner = "banners/standard.txt"

	if len(os.Args) != 2 || !strings.HasPrefix(os.Args[1], "--reverse=") {
		fmt.Print("Usage: go run . [OPTION]\n\n")
		fmt.Print("EX: go run . something standard --reverse=<fileName>\n")
		os.Exit(0)
	}

	content, err := os.ReadFile(os.Args[1][10:])

	if err != nil {
		log.Fatal(err)
	}

	var r string

	input := strings.Split(string(content), "\n")
	for i := 0; i < len(input)/8; i++ { // Finds how many lines of text there are
		var words [9]string

		for i2, v := range input[i*8 : (i+1)*8] {
			words[i2] = v
		}



		for words != rs {
			var found bool
			for i := 32; i <= 126; i++ {
				var letter [9]string = Text(banner, px+string(rune(i)))
				if Prefix(words, letter) {
					px += string(rune(i))
					rs = Text(banner, px)

					found = true
					break
				}
			}

			if !found { // So it doesnt keep going infinetly on 2 newlines
				break
			}
		}

		// fmt.Println(prefix)
		r += px

		if i != len(input)/8-1 {
			r += "\\n"
		}
	}

	fmt.Println(r)
}
func Text(banner string, text string) [9]string {
	ctt, err := os.ReadFile(banner) // Font file

	if err != nil {
		log.Fatal(err)
	}

	var letters [9]string
	re := regexp.MustCompile("(.+\n){8}")
	for i, v := range text {
		// Loops through the characters in font file
		for ix, vx := range re.FindAllString(string(ctt), -1) {
			if ix == int(v)-32 { // If correct character in font file
				for iy, vy := range strings.Split(vx, "\n") { // Splits font character into lines
					letters[iy] += vy // Appends lines onto array
				}
			}
		}

		if i == len(text)-1 { // If last character
			return letters
		}
	}

	return letters
}

func Prefix(word [9]string, lttr [9]string) bool {
	for i := 0; i < 9; i++ {
		if !strings.HasPrefix(word[i], lttr[i]) {
			return false
		}
	}

	return true
}
