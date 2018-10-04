package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/charly3pins/piglatin"
)

// Message is the object received via POST
type Message struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/translate", translate)

	log.Println("Server up and running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func translate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	m := Message{}
	if err = json.Unmarshal(body, &m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	result := processText(m.Text)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, result)
}

func processText(text string) string {
	defaultHandler := piglatin.DefaultHandler{}
	quHandler := piglatin.QuHandler{Next: &defaultHandler}
	vowelHandler := piglatin.VowelHandler{Next: &quHandler}
	hyphenHandler := piglatin.HyphenHandler{Next: &vowelHandler}
	symbolsHandler := piglatin.SymbolsHandler{Next: &hyphenHandler}

	myText := piglatin.NewText(text)
	translation := []string{}
	for _, word := range myText {
		word = symbolsHandler.Translate(word)
		translation = append(translation, word.GetTranslated())
	}

	return strings.Join(translation, " ")
}
