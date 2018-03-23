package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Pdf struct {
	Page  *Page
	Token string `json:"token"`
	State int    `json:"state"`
	URL   string `json:"url"`
}

func TempPdfPath(token string) string {
	path := tmpDir + token + ".pdf"
	return path
}

var PdfList = make(map[string]Pdf)
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const tokenLength = 12

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (pdf *Pdf) InitToken() {
	b := make([]rune, tokenLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	pdf.Token = string(b)
}

func DownloadPageBody(pipeline chan Pdf, pdf Pdf) error {
	page := pdf.Page

	fmt.Println("[Server]", pdf.Token, "[1]", "Downloading page from url:", page.URL)

	response, err := http.Get(page.URL)
	if err != nil {
		log.Fatal(err)
		return err
	} else {
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		page.Body = body

		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("[Server]", pdf.Token, "[1]", "Download Successful")
	}
	pipeline <- pdf
	return nil
}

func SavePageToFile(pipeline chan Pdf) error {
	pdf := <-pipeline
	page := pdf.Page
	fmt.Println("[Server]", pdf.Token, "[2]", "Saving page to path:", page.FilePath)

	err := ioutil.WriteFile(page.FilePath, page.Body, 0644)

	if err != nil {
		log.Fatal("[Server]", pdf.Token, "[2]", "Failed to write file", page.FilePath, err)
		return err
	} else {
		fmt.Println("[Server]", pdf.Token, "[2]", "Page saved successfuly to a file")
	}
	pipeline <- pdf
	return nil
}

func RenderAndSavePdf(pipeline chan Pdf) error {
	pdf := <-pipeline
	page := pdf.Page
	fmt.Println("[Server]", pdf.Token, "[3]", "Rendering PDF to internal buffer")

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPage(page.FilePath))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
		return err
	}
	pdfPath := TempPdfPath(pdf.Token)
	fmt.Println("[Server]", pdf.Token, "[3]", "Writing PDF to file:", pdfPath)
	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("[Server]", pdf.Token, "[3]", "Saved and rendered with success")
	pipeline <- pdf
	return nil
}

func UploadPdfToS3(pipeline chan Pdf) error {
	pdf := <-pipeline
	// pdfPath := TempPdfPath(pdf.Token)
	fmt.Println("[Server]", pdf.Token, "[4]", "Uploading PDF to S3")

	time.Sleep(time.Second * 1)
	// err := SendToS3(pdfPath)

	// if err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }
	fmt.Println("[Server]", pdf.Token, "[4]", "Uploaded PDF with success")

	<-pipeline
	return nil
}
