package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

type MovieMeta struct {
	Page       string  `json:"page"`
	Total      int     `json:"total"`
	PerPage    int     `json:"per_page"`
	TotalPages int     `json:"total_pages"`
	Data       []Movie `json:"data"`
}

type Movie struct {
	Title string `json:"title"`
}

//func main() {
//	bar := null.Int{}
//
//	fmt.Println("user member %v", bar.Valid)
//	fmt.Println(getMovieTitles("spiderman"))
//}

func getMovieTitles(substr string) []string {
	firstPageUrl := fmt.Sprintf("https://jsonmock.hackerrank.com/api/movies/search/?Title=%s&page=%d", substr, 1)
	fmt.Println("url", firstPageUrl)
	req, err := http.NewRequest(http.MethodGet, firstPageUrl, nil)
	if err != nil {
		fmt.Println("err creating http req", err)
	}
	client := http.Client{}
	firstResp, err := client.Do(req)
	if err != nil {
		fmt.Println("err sending http req for first page", err)
	}
	firstPage := new(MovieMeta)
	if err := json.NewDecoder(firstResp.Body).Decode(firstPage); err != nil {
		fmt.Println("err decoding response body for first page", err)
	}
	firstResp.Body.Close()

	fmt.Println("first page", firstPage.Data)

	fmt.Println("first page size", firstPage.Page)

	itemLen := firstPage.Total
	pageSize := firstPage.TotalPages
	perPage := firstPage.PerPage
	movieTitles := make([]string, itemLen)

	for i := 1; i <= pageSize; i++ {
		url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/movies/search/?Title=%s&page=%d", substr, i)
		fmt.Println("url", url)
		req, err = http.NewRequest(http.MethodGet, url, nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("err sending http req", err)
		}

		movieData := new(MovieMeta)
		if err := json.NewDecoder(resp.Body).Decode(movieData); err != nil {
			fmt.Println("err decoding response body", err)
		}
		for j, movie := range movieData.Data {
			// offset is pagesize + itemSize
			offset := (i-1)*perPage + j
			movieTitles[offset] = movie.Title
		}
		resp.Body.Close()
	}

	sort.Strings(movieTitles)
	return movieTitles

}

//func getMovieTitles(title string) []string {
//	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/movies/search/?Title=%s", title)
//	req, err := http.NewRequest(http.MethodGet, url, nil)
//	if err != nil {
//		fmt.Println("err", err)
//	}
//	client := http.Client{}
//	totalResp, err := client.Do(req)
//
//	meta := new(MovieMeta)
//	if err := json.NewDecoder(totalResp.Body).Decode(meta); err != nil {
//		fmt.Println("err", err)
//	}
//	totalResp.Body.Close()
//
//	pages := meta.TotalPages
//	titles := []string{}
//
//	for i := 1; i <= pages; i++ {
//		query := fmt.Sprintf("%s&page=%d", url, i)
//		subReq, err := http.NewRequest(http.MethodGet, query, nil)
//		if err != nil {
//			fmt.Println("err sub querying", err)
//		}
//		subResp, err := client.Do(subReq)
//		if err := json.NewDecoder(subResp.Body).Decode(meta); err != nil {
//			fmt.Println("err decoding sub response", err)
//		}
//		subResp.Body.Close()
//		for _, movie := range meta.Data {
//			titles = append(titles, movie.Title)
//		}
//	}
//	return titles
//}
func bindServer() {
	r := mux.NewRouter()
	r.HandleFunc("/foo", fooHandler)
	r.HandleFunc("/", homeHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:9200", nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 - not found")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is my home")
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yo bro!!!")
}
