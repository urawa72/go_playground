package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type Article []struct {
	RenderedBody  string    `json:"rendered_body"`
	Body          string    `json:"body"`
	Coediting     bool      `json:"coediting"`
	CommentsCount int       `json:"comments_count"`
	CreatedAt     time.Time `json:"created_at"`
	Group         struct {
		CreatedAt time.Time `json:"created_at"`
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Private   bool      `json:"private"`
		UpdatedAt time.Time `json:"updated_at"`
		URLName   string    `json:"url_name"`
	} `json:"group"`
	ID             string `json:"id"`
	LikesCount     int    `json:"likes_count"`
	Private        bool   `json:"private"`
	ReactionsCount int    `json:"reactions_count"`
	Tags           []struct {
		Name     string   `json:"name"`
		Versions []string `json:"versions"`
	} `json:"tags"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	User      struct {
		Description       string `json:"description"`
		FacebookID        string `json:"facebook_id"`
		FolloweesCount    int    `json:"followees_count"`
		FollowersCount    int    `json:"followers_count"`
		GithubLoginName   string `json:"github_login_name"`
		ID                string `json:"id"`
		ItemsCount        int    `json:"items_count"`
		LinkedinID        string `json:"linkedin_id"`
		Location          string `json:"location"`
		Name              string `json:"name"`
		Organization      string `json:"organization"`
		PermanentID       int    `json:"permanent_id"`
		ProfileImageURL   string `json:"profile_image_url"`
		TeamOnly          bool   `json:"team_only"`
		TwitterScreenName string `json:"twitter_screen_name"`
		WebsiteURL        string `json:"website_url"`
	} `json:"user"`
	PageViewsCount int `json:"page_views_count"`
}

func newRequest() (*http.Request, error) {
	url := "https://qiita.com/api/v2/items?page=1&per_page=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return req, err
}

func getResponse() (*http.Response, error) {
	req, err := newRequest()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error http status: %d", resp.StatusCode)
	}

	return resp, err
}

var articles Article

func encodeJSON() (Article, error) {
	resp, err := getResponse()
	if err != nil {
		return nil, err
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(byteArray, &articles); err != nil {
		log.Fatalf("Error %v", err)
	}

	return articles, err
}

func qiita(c *cli.Context) error {
	articles, err := encodeJSON()
	if err != nil {
		return err
	}

	for _, article := range articles {
		fmt.Println(article.Title)
		fmt.Println(article.URL)
	}

	// jsonArticles, err := json.Marshal(articles)
	// out := new(bytes.Buffer)
	// json.Indent(out, jsonArticles, "", "\t")
	// fmt.Println(out.String())

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "qiitago"
	app.Usage = "test"
	app.Version = "0.0.1"
	app.Action = qiita
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
