package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/tomaszgiba/go-pdf-service/lib"
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

func (pdf *Pdf) Init() {
	pdf.InitToken()
}

func (pdf *Pdf) InitToken() {
	b := make([]rune, tokenLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	pdf.Token = string(b)
}

func (pdf *Pdf) Finalize() {
	pdf.State = 1
}

func DownloadPageBody(pdf *Pdf) error {
	page := pdf.Page

	fmt.Println("[Server]", pdf.Token, "[1]", "Downloading page from url:", page.URL)

	response, err := http.Get(page.URL)
	if err != nil {
		log.Fatal("[Server]", pdf.Token, "[1]", err)
		return err
	} else {
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		page.Body = body

		if err != nil {
			log.Fatal("[Server]", pdf.Token, "[1]", err)
			return err
		}
		fmt.Println("[Server]", pdf.Token, "[1]", "Download Successful")
	}
	return nil
}

func SavePageToFile(pdf *Pdf) error {
	page := pdf.Page
	fmt.Println("[Server]", pdf.Token, "[2]", "Saving page to path:", page.FilePath)

	err := ioutil.WriteFile(page.FilePath, page.Body, 0644)

	if err != nil {
		log.Fatal("[Server]", pdf.Token, "[2]", "Failed to write file", page.FilePath, err)
		return err
	} else {
		fmt.Println("[Server]", pdf.Token, "[2]", "Page saved successfuly to a file")
	}
	return nil
}

func RenderAndSavePdf(pdf *Pdf) error {
	page := pdf.Page
	fmt.Println("[Server]", pdf.Token, "[3]", "Rendering PDF to internal buffer")

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal("[Server]", pdf.Token, "[3]", err)
		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(page.Body)))

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
		log.Fatal("[Server]", pdf.Token, "[3]", err)
		return err
	}

	fmt.Println("[Server]", pdf.Token, "[3]", "Saved and rendered with success")
	return nil
}

func UploadPdfToS3(pdf *Pdf) error {
	pdfPath := TempPdfPath(pdf.Token)
	fmt.Println("[Server]", pdf.Token, "[4]", "Uploading PDF to S3")

	err := lib.SendToS3(pdfPath)

	if err != nil {
		log.Fatal("[Server]", pdf.Token, "[4]", err)
		return err
	}
	fmt.Println("[Server]", pdf.Token, "[4]", "Uploaded PDF with success")

	return nil
}
