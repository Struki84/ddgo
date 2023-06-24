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
	queryUrl := fmt.Sprintf("https://duckduckgo.com/html/?q=%s", url.QueryEscape(query))

	response, err := http.Get(queryUrl)
	if err != nil {
		log.Printf("get %v error: %s", queryUrl, err)
		return results, err
	}
	
	defer response.Body.Close()
	
	if response.StatusCode != 200 {
        log.Printf("status code error: %d %s", response.StatusCode, response.Status)
		return results, fmt.Errorf("status code error: %d %s", response.StatusCode, response.Status)
    }
	
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("NewDocument Error: %s", err)
		return results, err
	}

	sel := doc.Find(".web-result")

	for i := range sel.Nodes {
		// Break loop once required amount of results are add
		if maxResult == len(results) {
			break
		}
		node := sel.Eq(i)
		titleNode := node.Find(".result__a")
		
		info := node.Find(".result__snippet").Text()
		title := titleNode.Text()
		ref, err := url.QueryUnescape(strings.TrimPrefix(titleNode.Nodes[0].Attr[2].Val, "/l/?kh=-1&uddg="))
		
		if err != nil {
			log.Printf("Error: %s", err)
			return results, err
		}

		results = append(results[:], Result{title, info, ref})

	}

	return results, nil
}
