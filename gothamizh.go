package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	thamizh   string = "தமிழ்"
	gothamizh string = "go-" + thamizh
)

var (
	vowels = map[string]string{
		"00":  "\u0bcd",
		"1e":  "\u0bbf",
		"0a":  "\u0bbe",
		"11e": "\u0bc0",
		"1u":  "\u0bc1",
		"2u":  "\u0bc2",
		"4a":  "\u0bc7",
		"11o": "\u0bcb",
		"3a":  "\u0bc6",
		"1o":  "\u0bca",
		"3o":  "\u0bcc",
		"4o":  "\u0bd7",
		"1i":  "\u0bc8",
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
		return mei[in], false
	}

	if vowels[in] != "" {
		return vowels[in], true
	}

	return "", false
}

func prompt() {
	fmt.Print(gothamizh + " >>")
}

func fabstring(in []string) string {
	var out string
	for _, s := range in {
		out += s
	}
	return out
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

func transliteratetamil(tokens []string) []string {
	var tamil []string
	for _, token := range tokens {
		token = strings.ToLower(token)
		i := 1
		s := 0
		var strTkns []string
		//fmt.Println("Debug => token " + token)
		for i <= len(token) {
			str, isvowel := checkmatch(token[s:i])
			if str != "" {
				//			fmt.Print("Debug000 => " + str)

				//			fmt.Print("Debug001 => " + strText)
				s = i
				i++
				if isvowel {
					// Take a copy of last item of the token
					sLast := strTkns[len(strTkns)-1]
					// Remove the item
					strTkns = strTkns[:len(strTkns)-1]
					// Attach it with str to form a perfect unicode
					smei := sLast + str
					strTkns = append(strTkns, smei)
				} else {
					strTkns = append(strTkns, str)
				}
			} else {
				i++
			}
		}
		//	fmt.Println("Debug002 => " + strText)
		strText := fabstring(strTkns)
		tamil = append(tamil, strText)
	}
	//fmt.Println("=== Debug ")
	//for _, t := range tamil {
	//	fmt.Println(t)
	//}
	return tamil
}

func renderhelp() {
	fmt.Println("Sample help")
}

func consoleMode() {
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
			texts := transliteratetamil(input)
			for _, output := range texts {
				fmt.Print(output)
				fmt.Print(" ")
			}
			fmt.Println()
			prompt()
		}
	}
}

func htmlhandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func helphandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "help.html")
}

func transliterateHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Debug===> Entering trans handler")
	r.ParseForm()
	text := r.Form["text"][0]

	input := strings.Split(text, " ")
	texts := transliteratetamil(input)
	var wout string
	for _, output := range texts {
		wout = wout + output + " "
	}

	//fmt.Println(r.Form["text"])
	fmt.Fprintf(w, wout)
}

func daemonMode() {
	http.HandleFunc("/", htmlhandler)
	http.HandleFunc("/trans", transliterateHandler)
	http.HandleFunc("/help", helphandler)
	fmt.Println("Server listening on port 8080....")
	http.ListenAndServe(":8080", nil)
}

func main() {
	if len(os.Args) < 2 {
		renderhelp()
	} else {
		switch args := os.Args[1]; args {
		case "-d":
			daemonMode()
		case "-c":
			consoleMode()
		case "-h":
			fallthrough
		case "--help":
			fallthrough
		default:
			renderhelp()
		}
	}
}
