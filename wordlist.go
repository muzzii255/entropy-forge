package entropyforge

import (
	_ "embed"
	"encoding/json"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
)

// Embed the wordlist file directly into the binary
//go:embed words.json
var wordlistData []byte

var caser cases.Caser
var wordList map[string]string

func init() {
	// Load wordlist from embedded data instead of reading file
	if err := json.Unmarshal(wordlistData, &wordList); err != nil {
		log.Fatalf("Failed to load embedded wordlist: %v", err)
	}
	
	caser = cases.Title(language.Und)
}