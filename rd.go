package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/kennygrant/sanitize"
	"github.com/mgutz/ansi"
	"github.com/robinjmurphy/go-readability-api/readability"
)

func usage() {
	fmt.Println("Usage: rd <url>")
	flag.PrintDefaults()
	os.Exit(1)
}

func printMissingKeyMessage() {
	fmt.Println("Ensure that READABILITY_PARSER_API_KEY is set. ")
	fmt.Println("See https://github.com/robinjmurphy/rd#installation.")
	os.Exit(1)
}

func printError(err error) {
	fmt.Printf(ansi.Color("× %s\n", "red"), err.Error())
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		usage()
	}
	url := args[0]
	key := os.Getenv("READABILITY_PARSER_API_KEY")
	if key == "" {
		printMissingKeyMessage()
	}
	client := readability.NewClient("", "", key)
	parser := client.NewParserClient()
	article, resp, err := parser.Parse(url)
	if resp != nil && resp.StatusCode == 400 {
		usage()
	}
	if err != nil {
		printError(err)
	}
	content := sanitize.HTML(article.Content)
	// remove vertical whitespace
	r := regexp.MustCompile("[\n]{2,}")
	content = r.ReplaceAllString(content, "")
	// remove horizontal whitespace
	content = strings.Replace(content, "  ", "", -1)
	fmt.Println(article.Title + "\n")
	fmt.Println(strings.Replace(strings.TrimSpace(content), "\n", "\n\n", -1))
}
