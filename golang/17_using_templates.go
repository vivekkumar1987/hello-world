package main

import ("fmt"
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"html/template"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

type NewsAppPage struct {
	Title string
	News map[string]NewsMap
}

func newsAgghandler(w http.ResponseWriter, r *http.Request)  {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations{
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	p := NewsAppPage{Title: "Amazing news agg", News: news_map}
	// From IntelliJ the following throws "no such file or directory"
	t, err := template.ParseFiles("newsaggtemplate.html")
	fmt.Println(err)
	fmt.Println(t.Execute(w,p))
}

func index_handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>Whoa, Go is neat!</h1>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/agg/", newsAgghandler)
	http.ListenAndServe(":8000", nil)

}
