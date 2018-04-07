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
