package xkcd

import (
	"html/template"
)

type Info struct {
	Title      template.HTML // suppress escape HTML
	Num        int
	Img        string
	Transcript string
	Year       string
	Month      string
	Day        string
	SafeTitle  string `json:safe_title`
	Alt        string
	Link       string
	News       string
}
