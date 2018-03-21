package lib

import "math/rand"

type Pdf struct {
	Page  *Page
	Token string `json:"token"`
	State int    `json:"state"`
	URL   string `json:"url"`
}

var PdfList = make(map[string]Pdf)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const tokenLength = 32

func (pdf *Pdf) InitToken() {
	b := make([]byte, tokenLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	pdf.Token = string(b)
}

func (pdf *Pdf) renderPdf(c chan int) error {

	return nil
}
