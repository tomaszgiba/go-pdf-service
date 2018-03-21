package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL      string
	FilePath string
	Body     []byte
}

func (page *Page) SaveToFile(c chan *Page) {
	fmt.Println("[Server]", "Saving page to path:", page.FilePath)

	err := ioutil.WriteFile(page.FilePath, page.Body, 0644)

	if err != nil {
		log.Fatal(err)
	} else {
		c <- page
		fmt.Println("[Server]", "Page saved successfuly to a file")
	}
}

func (page *Page) DownloadBody(c chan *Page) {
	fmt.Println("[Server]", "Downloading page from url:", page.URL)
	response, err := http.Get(page.URL)

	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		page.Body = body

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("[Server]", "Download Successful")
			c <- page
		}
	}
}
