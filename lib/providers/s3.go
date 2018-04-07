package providers

import (
	"time"
)

const s3Region = "eu-central-1"
const s3Bucket = "sandbox75982"

func SendToS3(filePath string) error {
	time.Sleep(2 * time.Second)
	return nil
	// // Create a single AWS session (we can re use this if we're uploading many files)
	// s, err := session.NewSession(&aws.Config{Region: aws.String(s3Region)})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Open the file for use
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// // Get file size and read the file content into a buffer
	// fileInfo, _ := file.Stat()
	// var size int64 = fileInfo.Size()
	// buffer := make([]byte, size)
	// file.Read(buffer)

	// // Config settings: this is where you choose the bucket, filename, content-type etc.
	// // of the file you're uploading.
	// _, err = s3.New(s).PutObject(&s3.PutObjectInput{
	// 	Bucket:               aws.String(s3Bucket),
	// 	Key:                  aws.String(filePath),
	// 	ACL:                  aws.String("private"),
	// 	Body:                 bytes.NewReader(buffer),
	// 	ContentLength:        aws.Int64(size),
	// 	ContentType:          aws.String(http.DetectContentType(buffer)),
	// 	ContentDisposition:   aws.String("attachment"),
	// 	ServerSideEncryption: aws.String("AES256"),
	// })
	// return err
}
