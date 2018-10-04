# piglatin
[![Go Report Card](https://goreportcard.com/badge/github.com/charly3pins/piglatin)](https://goreportcard.com/report/github.com/charly3pins/piglatin) [![GoDoc](https://godoc.org/github.com/charly3pins/piglatin?status.svg)](https://godoc.org/github.com/charly3pins/piglatin)

## Pig Latin Language
The rules to translate from English to Pig Latin are the following:
- Words that start with a vowel (A, E, I, O, U) simply have "WAY" appended to the end of the word;
- Words that start with a consonant have all consonant letters up to the first vowel moved to the end of the word (as opposed to just the first consonant letter), and "AY" is appended.
    - In this context a ‘y' in the middle of the word is counted as a vowel;
    - In this context if the word starts with 'qu’ we consider that to be a single consonant;
- Translations should respect upper/lower case formatting;
- Hyphenated words are treated as two words;
- Words may consist of alphabetic characters only (A-Z and a-z);
- All punctuation, numerals, symbols and whitespace are not modified;
- Let's assume that there are no contractions in the English text;

## How it works
First of all you need to clone the project:
```
git clone https://github.com/charly3pins/piglatin.git
```
The translation service is accessible via 2 ways. A command-line servive and the API interface. 

### Command line
This part is for testing the service without up&running the API. You can provide the text via `flag` in your terminal and the response will be the text translated.
To generate the binary you need to type:
```
go build -o demo ./cmd/demo
```
If you want to translate the text `Roses are red, violets are blue` you need to type:
```
./demo -text "Roses are red, violets are blue"
```
The text you should obtain is `Osesray areway edray, ioletsvay areway ueblay`.

### API
If you want to up the service via API you can do it using `docker-compose`. Simply type the following command:
```
docker-compose up
```

The server is listening port `:8080` and it's waiting for a POST request with the following format:
```
{
	"text": "Roses are red, violets are blue"
}
```

You can make a call using `curl`:
```
curl -d '{"text":"Roses are red, violets are blue"}' -H "Content-Type: application/json" -X POST http://localhost:8080/translate
```

## Test & Coverage
Taking advantage of the tools that `Go` provides for testing, the project also contains a dummy server with nginx that runs the `go test` and `go tool cover` commands to generate the html output with the corresponding coverage.
```
make docker_coverage
```
Once all is ready visit `http://localhost:8081/`