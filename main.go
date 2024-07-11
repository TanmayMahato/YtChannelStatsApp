package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/lpernett/godotenv"
)

type Usr struct {
	Uname string
}

type Datas struct {
	Kind     string       `json:"kind"`
	Etag     string       `json:"etag"`
	PageInfo PageInfo     `json:"pageInfo"`
	Items    []InsideItem `json:"items"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type InsideItem struct {
	Kind       string     `json:"kind"`
	Etag       string     `json:"etag"`
	Id         string     `json:"id"`
	Statistics Statistics `json:"statistics"`
}

type Statistics struct {
	Subs    string `json:"subscriberCount"`
	Views   string `json:"viewCount"`
	HidSubs bool   `json:"hiddenSubscriberCount"`
	Nov     string `json:"videoCount"`
}

func handler0(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handler1(w http.ResponseWriter, r *http.Request) {
	Lc := r.FormValue("first")

	U1 := Usr{Uname: Lc}
	fmt.Println(U1)

	us1 := U1.Uname

	_, after, _ := strings.Cut(us1, "@")
	before, _, _ := strings.Cut(after, "/")

	req := "https://www.googleapis.com/youtube/v3/channels"
	para1 := "?part=statistics"
	username := "&forHandle=" + before
	para2 := "&key=" + os.Getenv("YtApiKey")

	Url := req + para1 + username + para2

	resp, _ := http.Get(Url)
	var Data Datas

	err := json.NewDecoder(resp.Body).Decode(&Data)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("temp.html")
	tmpl.Execute(w, Data.Items[0].Statistics)

}

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/", handler0)
	http.HandleFunc("/data", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
