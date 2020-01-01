package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)


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
		for _, tag := range article.Tags {
			fmt.Println(tag.Name)
		}
	}

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
