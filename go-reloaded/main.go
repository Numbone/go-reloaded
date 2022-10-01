package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Error: need 2 file")
		return
	}
	if os.Args[1] != "sample.txt" {
		fmt.Println("You need to write to sample.txt ")
		return
	}
	if os.Args[2] != "result.txt" {
		fmt.Println("You need to write to result.txt ")
		return
	}
	sampleText, err := os.ReadFile(args[0])
	if len(sampleText) == 0 {
		fmt.Println("There is nothing to do")
		return
	}
	if err != nil {
		fmt.Println("Error:wrong parameter file")
	}

	words := strings.Split(string(sampleText), " ")

	fmt.Println(string(sampleText))
	Vowels(words)
	b := LastApos(Punctuations(FirstApos(MainFunc(BrackShit(SpaceFirst(words))))))
	text := strings.Split(string(sampleText), " ")
	c := ReturnBrackShit(b, text)
	middleware := strings.Join(c, " ")
	res := os.WriteFile(args[1], []byte(middleware), 0644)
	if res != nil {
		panic(res)
	}
	fmt.Println(c)
}

func Hex(hex string) string {
	n, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(n)
}

func Bin(bin string) string {
	n, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(n)
}

func Vowels(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}

	for i, word := range s {
		for _, letter := range vowels {
			if word == "a" && string(s[i+1][0]) == letter {
				s[i] = "an"
			} else if word == "A" && string(s[i+1][0]) == letter {
				s[i] = "An"
			}
		}
	}
	return s
}

func removeBeetweenSpace(s []string) []string {
	var res []string
	for _, word := range s {
		if word != "" {
			res = append(res, word)
		}
	}
	return res
}

func SpaceFirst(s []string) []string {
	for s[0] == "" {
		s = s[1:]
		SpaceFirst(s)
	}

	for s[len(s)-1] == "" {
		s = s[:len(s)-1]
		SpaceFirst(s)
	}

	return s
}

func Punctuations(s []string) []string {
	for s[0] == "" {
		s = s[1:]
		SpaceFirst(s)
	}

	for s[len(s)-1] == "" {
		s = s[:len(s)-1]
		SpaceFirst(s)
	}

	puncs := []string{",", ".", "!", "?", ":", ";"}

	for i, word := range s {
		// for _, except := range exception {
		if string(word) == "..." && i != 0 {

			s[i-1] += word
			s = append(s[:i], s[i+1:]...)
		}
	}
	for i, word := range s {
		if string(word) == "!?" && i != 0 {

			fmt.Println(s[i])
			s[i-1] += word
			s = append(s[:i], s[i+1:]...)
		}
	}
	for i, word := range s {
		for _, punc := range puncs {
			if (string(word[0])) == punc && (string(word[len(word)-1])) == punc && i > 0 {
				// if (string(word[0])) == punc && s[0] != punc && s[len(s)-1] != punc && (string(word[len(word)-1])) == punc && i > 0  {
				s[i-1] += word

				s = append(s[:i], s[i+1:]...)
			} else if s[0] == punc && s[len(s)-1] == punc && s[len(s)-1] == " " {
				return s
			}
		}
	}
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] && i != 0 {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	for i, word := range s {
		for _, p := range puncs {
			if string(word[0]) == p && string(word[len(word)-1]) != p && s[0] != p && i != 0 {

				s[i-1] += p
				s[i] = word[1:]
			}
		}
	}
	for _, word := range s {
		if s[len(s)-1] != " " && word != "..." {
			return s
		}
	}

	return removeBeetweenSpace(s)
}

func Apos(s []string) []string {
	count := 0
	for i, word := range s {
		if word == "‘" && count == 0 {
			count += 1
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}

	if s[len(s)-1] == "‘" && count == 1 {
		// s[len(s)-2] = s[len(s)-2] + s[len(s)-1]
		// s = s[:len(s)-1]
		fmt.Println(s[len(s)-2])
	}
	return s
}

func FirstApos(s []string) []string {
	count := 0
	for i, word := range s {
		if word == "'" && count == 0 {
			count += 1
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func LastApos(s []string) []string {
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i][0]) == "'" {
			s[i-1] = s[i-1] + string(s[i])
			s = append(s[:i], s[i+1:]...)
		}
		break
	}
	return s
}

func BrackShit(words []string) []string {
	if words[0][0] == 34 && len(words) > 1 {
		words[0] = words[0][1:]
	}
	b := (words[len(words)-1])
	if words[len(words)-1][len(b)-1] == 34 && len(words) > 1 {
		words[len(words)-1] = words[len(words)-1][:len(b)-1]
	}
	return words
}

func ReturnBrackShit(b, words []string) []string {
	words = SpaceFirst(words)
	if words[0][0] == 34 && len(words[0]) > 1 {
		b[0] = string(rune(34)) + b[0]
	}
	a := (words[len(words)-1])
	if words[len(words)-1][len(a)-1] == 34 && len(words) > 1 {
		b[len(b)-1] = b[len(b)-1] + string(rune(34))
	}
	return b
}

func MainFunc(words []string) []string {
	return HexBin(Up(Low(Cap(LowNum(CapNum(UpNum(words)))))))
}

func Up(words []string) []string {
	for i, word := range words {
		if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func Low(words []string) []string {
	for i, word := range words {
		if word == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func Cap(words []string) []string {
	for i, word := range words {
		if word == "(cap)" {
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func UpNum(words []string) []string {
	for i, word := range words {
		if word == "(up," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])

			number, _ := strconv.Atoi(string(b))

			for j := 1; j <= number; j++ {
				if number > i-1 {
					number = i
				}
				words[i-j] = strings.ToUpper(words[i-j])
			}

			words = append(words[:i], words[i+2:]...)
		}
	}
	return words
}

func LowNum(words []string) []string {
	for i, word := range words {
		if word == "(low," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))

			for j := 1; j <= number; j++ {
				if number > i-1 {
					number = i
				}
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		}
	}

	return words
}

func CapNum(words []string) []string {
	for i, word := range words {
		if word == "(cap," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))

			for j := 1; j <= number; j++ {
				if number > i-1 {
					number = i
				}
				words[i-j] = strings.Title(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	return words
}

func HexBin(words []string) []string {
	for i, word := range words {
		if word == "(hex)" {
			words[i-1] = Hex(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(bin)" {
			words[i-1] = Bin(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
