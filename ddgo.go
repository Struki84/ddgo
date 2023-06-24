package ddgo

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Result holds the returned query data
type Result struct {
	Title string
	Info  string
	Ref   string
}

// Requests the query and puts the results into an array
func Query(query string, maxResult int) ([]Result, error) {
	results := []Result{}
	url := fmt.Sprintf("https://duckduckgo.com/html/?q=%s", url.QueryEscape(query))

	response, err := http.Get(url)
	if err != nil {
		log.Printf("Get %v Error: %s", url, err)
		return results, err
	}
	defer response.Body.Close()
	
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Pr
	}

	sel := doc.Find(".web-result")

	for i := range sel.Nodes {
		// Break loop once required amount of results are add
		if it == len(results) {
			break
		}

		single := sel.Eq(i)
		titleNode := single.Find(".result__a")
		info := single.Find(".result__snippet").Text()
		title := titleNode.Text()
		ref, _ := url.QueryUnescape(strings.TrimPrefix(titleNode.Nodes[0].Attr[2].Val, "/l/?kh=-1&uddg="))

		results = append(results[:], Qresult{title, info, ref})

	}

	// Return array of results and formated query used to get the results
	return results, qf
}
