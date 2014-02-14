package main

import (
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	MEANINGS_COUNT = 3
)

func loadUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	t, err := charset.TranslatorFrom("windows-1251")
	if err != nil {
		return "", err
	}
	_, translatedBytes, err := t.Translate(bodyBytes, true)
	if err != nil {
		return "", err
	}

	return string(translatedBytes), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run multitran.go <word>")
		return
	}
	arg_word := os.Args[1]

	x, err := loadUrl("http://www.multitran.ru/c/m.exe?l1=1&l2=2&s=" + arg_word)
	if err == nil {
		ind := strings.Index(x, `title="Общая лексика"`)
		if ind == -1 {
			return
		}
		tail := x[ind:]
		ind = strings.Index(tail, "<td>")
		if ind == -1 {
			return
		}
		tail = tail[ind:]
		ind = strings.Index(tail, "</td>")
		if ind == -1 {
			return
		}
		tail = tail[:ind]

		re := regexp.MustCompile(`<[^>]+>`)
		tail = re.ReplaceAllString(tail, "")
		words := strings.Split(tail, "; ")
		fmt.Print(arg_word, "\t")
		for i, w := range words {
			fmt.Print(w, ";")
			if i == MEANINGS_COUNT-1 {
				break
			}
		}
		fmt.Println()
	}
}
