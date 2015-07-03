# rd

> Get the plain-text content of any article on the web using the [Readability](https://readability.com/) parser

## Installation

```
go get github.com/robinjmurphy/rd
```

Export your Readability Parser API key:

```
export READABILITY_PARSER_API_KEY=...
```

## Usage

```bash
rd http://www.bbc.co.uk/news/technology-33228149
# Taylor forces Apple to listen
#
# Apple is not a company famed for listening. After all, it prides itself on
# knowing what consumers want before they do, so why should it care what they
# think? All the more surprising then, that it should have listened to one angry
# customer, a Ms T Swift of Beverly Hills, California.
# ...
```
