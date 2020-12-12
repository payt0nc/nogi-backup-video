package bilibili

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Generated by https://quicktype.io

type SearchResult struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
	Data    Data   `json:"data"`
}

type Data struct {
	List           List           `json:"list"`
	Page           Page           `json:"page"`
	EpisodicButton EpisodicButton `json:"episodic_button"`
}

type EpisodicButton struct {
	Text string `json:"text"`
	URI  string `json:"uri"`
}

type List struct {
	Tlist map[string]Tlist `json:"tlist"`
	Vlist []Vlist          `json:"vlist"`
}

type Tlist struct {
	Tid   int64  `json:"tid"`
	Count int64  `json:"count"`
	Name  string `json:"name"`
}

type Vlist struct {
	Comment      int64  `json:"comment"`
	Typeid       int64  `json:"typeid"`
	Play         int64  `json:"play"`
	Pic          string `json:"pic"`
	Subtitle     string `json:"subtitle"`
	Description  string `json:"description"`
	Copyright    string `json:"copyright"`
	Title        string `json:"title"`
	Review       int64  `json:"review"`
	Author       string `json:"author"`
	Mid          int64  `json:"mid"`
	Created      int64  `json:"created"`
	Length       string `json:"length"`
	VideoReview  int64  `json:"video_review"`
	Aid          int64  `json:"aid"`
	Bvid         string `json:"bvid"`
	HideClick    bool   `json:"hide_click"`
	IsPay        int64  `json:"is_pay"`
	IsUnionVideo int64  `json:"is_union_video"`
	IsSteinsGate int64  `json:"is_steins_gate"`
}

type Page struct {
	Pn    int64 `json:"pn"`
	PS    int64 `json:"ps"`
	Count int64 `json:"count"`
}

func SearchVideo(keyword string, mid, pageNumber int) (SearchResult, error) {
	var sr SearchResult
	url := fmt.Sprintf("https://api.bilibili.com/x/space/arc/search?mid=%d&ps=30&tid=0&pn=%d&keyword=%s&order=pubdate&jsonp=jsonp", mid, pageNumber, keyword)
	res, err := http.Get(url)
	if err != nil {
		return sr, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		return sr, err
	}
	return sr, nil
}