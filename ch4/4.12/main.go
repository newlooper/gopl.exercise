package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopl.exercise/ch4/4.12/xkcd"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

const urlPrefix = "https://xkcd.com/"

const tpl = `
==================================================
Title: {{.Title}}
URL: {{.Num | getUrl}}
Year: {{.Year}}
Month: {{.Month}}
Day: {{.Day}}
==================================================
`

var index = flag.Int("i", 0, "the comic index")

func main() {

	flag.Parse()
	if *index <= 0 {
		log.Fatalf("error: %s\n", "invalid index")
	}

	url := urlPrefix + strconv.Itoa(*index) + "/info.0.json"
	resp, err := http.Get(url)
	if err != nil {
		resp.Body.Close()
		log.Fatalf("request error: %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("http error code: %d\n", resp.StatusCode)
	}

	var info = xkcd.Info{}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		resp.Body.Close()
		log.Fatalf("decode error: %v\n", err)
	}

	resp.Body.Close()

	report := template.Must(template.New("info").
		Funcs(template.FuncMap{"getUrl": getUrl}).
		Parse(tpl))
	report.Execute(os.Stdout, info)
}

func getUrl(i int) string {
	return urlPrefix + fmt.Sprintf("%d", i)
}
