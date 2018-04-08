package model

type Page struct {
	URL      string
	FilePath string
	Body     []byte
}

const tmpDir = "tmp/"

func TempFilePath(token string) string {
	path := tmpDir + token

	return path
}

func (page *Page) Init(url string, body []byte) {
	page.URL = url
	page.Body = body
}
