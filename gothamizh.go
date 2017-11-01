package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	thamizh   string = "தமிழ்"
	gothamizh string = "go-" + thamizh
)

var (
	vowels = map[string]string{
		"00": "\u0bcd",
		"1e": "\u0bbf",
		"0a": "\u0bbe",
		"2e": "\u0bc0",
		"1u": "\u0bc1",
		"2u": "\u0bc2",
		"3a": "\u0bc7",
		"4a": "\u0bcb",
		"1o": "\u0bc6",
		"2o": "\u0bca",
		"3o": "\u0bcc",
		"4o": "\u0bd7",
	}

	mei = map[string]string{
		"ka":   "\u0b95",
		"nga":  "\u0b99",
		"ca":   "\u0b9a",
		"nya":  "\u0b9e",
		"ja":   "\u0b9c",
		"tta":  "\u0b9f",
		"nna":  "\u0ba3",
		"ta":   "\u0ba4",
		"na":   "\u0ba8",
		"nnna": "\u0ba9",
		"pa":   "\u0baa",
		"ma":   "\u0bae",
		"ya":   "\u0baf",
		"ra":   "\u0bb0",
		"rra":  "\u0bb1",
		"la":   "\u0bb2",
		"lla":  "\u0bb3",
		"llla": "\u0bb4",
		"va":   "\u0bb5",
	}

	uyir = map[string]string{
		"a":  "\u0b85",
		"2a": "\u0b86",
		"i":  "\u0b87",
		"2i": "\u0b88",
		"u":  "\u0b89",
		"2u": "\u0b8a",
		"e":  "\u0b8e",
		"2e": "\u0b8f",
		"3i": "\u0b90",
		"o":  "\u0b92",
		"2o": "\u0b93",
	}
)

func checkmatch(in string) (string, bool) {
	//fmt.Println("Debug : input string to checkmatch() " + in)
	if uyir[in] != "" {
		return uyir[in], false
	}

	if mei[in] != "" {
		return mei[in], true
	}

	if vowels[in] != "" {
		return vowels[in], false
	}

	return "", false
}

func prompt() {
	fmt.Print(gothamizh + " >>")
}

/*
Algo
-----
*1) Range list of tokens
*2) On each token, take every character and check for Uyir or Mei
*3) If not, check for Mei with vowels
*4) Start with single char, look for above 3 case, if not successful, take char + next char and proceed
* until you have a match
*/

func transliteratetamil(tokens []string) {
	for _, token := range tokens {
		i := 1
		s := 0
		//fmt.Println("Debug => token " + token)
		for i <= len(token) {
			str, ismei := checkmatch(token[s:i])
			if str != "" {
				if ismei {
					var str2 string
					if i+3 < len(token) {
						str2, _ = checkmatch(token[i+1 : i+3])
						i = i + 3
					} else {
						str2, _ = checkmatch(token[i:])
					}
					//fmt.Println(str)
					//fmt.Println(str2)
					//fmt.Println(str + str2)
					fmt.Print(str + str2)
					s = i
				} else {
					fmt.Print(str)
					s = i
					i++
				}
			} else {
				i++
			}
		}
		fmt.Print(" ")
	}
	fmt.Println("")
}

func renderhelp() {
	fmt.Println("Sample help")
}

func main() {
	prompt()
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		scan := reader.Text()
		switch scan {
		case "exit":
			fallthrough
		case "quit":
			fallthrough
		case "shutdown":
			os.Exit(0)
		case "help":
			fallthrough
		case "h":
			renderhelp()
			prompt()
		default:
			input := strings.Split(reader.Text(), " ")
			transliteratetamil(input)
			prompt()
		}
	}
}
