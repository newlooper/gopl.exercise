package main

import (
	"github.com/olekukonko/tablewriter"
	"gopl.io/ch4/github"
	"log"
	"os"
	"strconv"
	"time"
)

const hoursOfMonth = 24 * 31
const hoursOfYear = 24 * 365

func hoursAgo(t time.Time) int {
	return int(time.Since(t).Hours())
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var inMonth, inYear, outYear []*github.Issue
	for _, item := range result.Items {
		if hoursAgo(item.CreatedAt) < hoursOfMonth {
			inMonth = append(inMonth, item)
		} else if hoursAgo(item.CreatedAt) < hoursOfYear {
			inYear = append(inYear, item)
		} else {
			outYear = append(outYear, item)
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Time", "Number", "User", "created", "title"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)

	for _, item := range inMonth {
		table.Append([]string{strconv.Itoa(len(inMonth)) + " In Month", strconv.Itoa(item.Number), item.User.Login, item.CreatedAt.String(), item.Title})
	}

	for _, item := range inYear {
		table.Append([]string{strconv.Itoa(len(inYear)) + " In Year", strconv.Itoa(item.Number), item.User.Login, item.CreatedAt.String(), item.Title})
	}

	for _, item := range outYear {
		table.Append([]string{strconv.Itoa(len(outYear)) + " Out Year", strconv.Itoa(item.Number), item.User.Login, item.CreatedAt.String(), item.Title})
	}

	table.Render()
}

/*
go run . repo:golang/go is:open json decoder

+-------------+--------+------------------+-------------------------------+--------------------------------------+
|    TIME     | NUMBER |       USER       |            CREATED            |                TITLE                 |
+-------------+--------+------------------+-------------------------------+--------------------------------------+
| 3 In Month  |  40982 | Segflow          | 2020-08-22 21:07:03 +0000 UTC | encoding/json: use different         |
|             |        |                  |                               | error type for unknown field         |
|             |        |                  |                               | if they are disallowed               |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  41144 | alvaroaleman     | 2020-08-31 14:27:19 +0000 UTC | encoding/json: Unmarshaler           |
|             |        |                  |                               | breaks DisallowUnknownFields         |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  40983 | Segflow          | 2020-08-22 21:13:48 +0000 UTC | encoding/json: return a              |
|             |        |                  |                               | different error type for             |
|             |        |                  |                               | unknown field if they are            |
|             |        |                  |                               | disallowed                           |
+-------------+--------+------------------+-------------------------------+--------------------------------------+
| 6 In Year   |  34647 | babolivier       | 2019-10-01 21:42:29 +0000 UTC | encoding/json: fix byte              |
|             |        |                  |                               | counter increments when using        |
|             |        |                  |                               | decoder.Token()                      |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  40128 | rogpeppe         | 2020-07-09 07:58:19 +0000 UTC | proposal: encoding/json:             |
|             |        |                  |                               | garbage-free reading of tokens       |
+             +--------+                  +-------------------------------+--------------------------------------+
|             |  40127 |                  | 2020-07-09 07:52:47 +0000 UTC | proposal: encoding/json: add         |
|             |        |                  |                               | Encoder.EncodeToken method           |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  36225 | dsnet            | 2019-12-19 22:26:12 +0000 UTC | encoding/json: the                   |
|             |        |                  |                               | Decoder.Decode API lends             |
|             |        |                  |                               | itself to misuse                     |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  34543 | maxatome         | 2019-09-25 22:13:24 +0000 UTC | encoding/json: Unmarshal             |
|             |        |                  |                               | & json.(*Decoder).Token              |
|             |        |                  |                               | report different values for          |
|             |        |                  |                               | SyntaxError.Offset for the           |
|             |        |                  |                               | same input                           |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  34564 | mdempsky         | 2019-09-27 00:48:51 +0000 UTC | go/internal/gcimporter: single       |
|             |        |                  |                               | source of truth for decoder          |
|             |        |                  |                               | logic                                |
+-------------+--------+------------------+-------------------------------+--------------------------------------+
| 21 Out Year |  33416 | bserdar          | 2019-08-01 19:40:12 +0000 UTC | encoding/json: This CL adds          |
|             |        |                  |                               | Decoder.InternKeys                   |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  12001 | lukescott        | 2015-08-03 19:14:17 +0000 UTC | encoding/json:                       |
|             |        |                  |                               | Marshaler/Unmarshaler not            |
|             |        |                  |                               | stream friendly                      |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  29035 | jaswdr           | 2018-11-30 11:21:31 +0000 UTC | proposal: encoding/json: add         |
|             |        |                  |                               | error var to compare  the            |
|             |        |                  |                               | returned error when using            |
|             |        |                  |                               | json.Decoder.DisallowUnknownFields() |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  11046 | kurin            | 2015-06-03 19:25:08 +0000 UTC | encoding/json: Decoder               |
|             |        |                  |                               | internally buffers full input        |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |   5901 | rsc              | 2013-07-17 16:39:15 +0000 UTC | encoding/json: allow                 |
|             |        |                  |                               | per-Encoder/per-Decoder              |
|             |        |                  |                               | registration of                      |
|             |        |                  |                               | marshal/unmarshal functions          |
+             +--------+                  +-------------------------------+--------------------------------------+
|             |  32779 |                  | 2019-06-25 21:08:30 +0000 UTC | encoding/json: memoize strings       |
|             |        |                  |                               | during decode                        |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  28923 | mvdan            | 2018-11-22 13:50:18 +0000 UTC | encoding/json: speed up the          |
|             |        |                  |                               | decoding scanner                     |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  31701 | lr1980           | 2019-04-26 20:50:17 +0000 UTC | encoding/json: second decode         |
|             |        |                  |                               | after error impossible               |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  14750 | cyberphone       | 2016-03-10 13:04:44 +0000 UTC | encoding/json: parser ignores        |
|             |        |                  |                               | the case of member names             |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  16212 | josharian        | 2016-06-29 16:07:36 +0000 UTC | encoding/json: do all reflect        |
|             |        |                  |                               | work before decoding                 |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |   6647 | btracey          | 2013-10-23 17:19:48 +0000 UTC | x/tools/cmd/godoc: display           |
|             |        |                  |                               | type kind of each named type         |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  28143 | arp242           | 2018-10-11 07:08:25 +0000 UTC | proposal: encoding/json: add         |
|             |        |                  |                               | "readonly" tag                       |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  30301 | zelch            | 2019-02-18 19:49:27 +0000 UTC | encoding/xml: option to treat        |
|             |        |                  |                               | unknown fields as an error           |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  33854 | Qhesz            | 2019-08-27 00:20:25 +0000 UTC | encoding/json: unmarshal             |
|             |        |                  |                               | option to treat omitted fields       |
|             |        |                  |                               | as null                              |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  26946 | deuill           | 2018-08-12 18:19:01 +0000 UTC | encoding/json: clarify what          |
|             |        |                  |                               | happens when unmarshaling into       |
|             |        |                  |                               | a non-empty interface{}              |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  22752 | buyology         | 2017-11-15 23:46:13 +0000 UTC | proposal: encoding/json: add         |
|             |        |                  |                               | access to the underlying data        |
|             |        |                  |                               | causing UnmarshalTypeError           |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  33835 | Qhesz            | 2019-08-26 10:27:12 +0000 UTC | encoding/json: unmarshalling         |
|             |        |                  |                               | null into non-nullable golang        |
|             |        |                  |                               | types                                |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  22480 | blixt            | 2017-10-28 20:20:06 +0000 UTC | proposal: encoding/json: add         |
|             |        |                  |                               | omitnil option                       |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  28189 | adnsv            | 2018-10-13 16:22:53 +0000 UTC | encoding/json: confusing             |
|             |        |                  |                               | errors when unmarshaling             |
|             |        |                  |                               | custom types                         |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |  27179 | lavalamp         | 2018-08-23 18:21:32 +0000 UTC | encoding/json: no way to             |
|             |        |                  |                               | preserve the order of map keys       |
+             +--------+------------------+-------------------------------+--------------------------------------+
|             |   7872 | extemporalgenome | 2014-04-26 17:47:25 +0000 UTC | encoding/json: Encoder               |
|             |        |                  |                               | internally buffers full output       |
+-------------+--------+------------------+-------------------------------+--------------------------------------+

*/
