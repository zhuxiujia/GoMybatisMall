package fs

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func PostGoFastDfsFile(url string, fileName string, file io.Reader) (string, error) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	fileWriter, _ := bodyWriter.CreateFormFile("file", fileName)
	io.Copy(fileWriter, file)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, e := http.Post(url, contentType, bodyBuffer)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if e != nil {
		return "", e
	}
	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}
	resp_body, e := ioutil.ReadAll(resp.Body)
	return string(resp_body), e
}

func UploadHandler(goFastDfsUrl string, callback func(result string, e error)) func(w http.ResponseWriter, r *http.Request) {
	var hand = func(w http.ResponseWriter, r *http.Request) {
		reader, err := r.MultipartReader()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			log.Printf("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())
			if part.FileName() == "" { // this is FormData
				data, _ := ioutil.ReadAll(part)
				log.Printf("FormData=[%s]\n", string(data))
			} else { // This is FileData
				var tempPath = "./upload/temp/"
				os.MkdirAll(tempPath, os.ModeDir)
				var filePath = tempPath + part.FileName()
				dst, e := os.Create(filePath)
				if e != nil {
					callback("", e)
					return
				}
				io.Copy(dst, part)

				part.Close()
				dst.Close()
				println("upload:" + dst.Name())

				f, e := os.Open(filePath)
				if f != nil {
					defer func() {
						f.Close()
						println("remove:" + f.Name())
						os.Remove(f.Name())
					}()
				}
				if e != nil {
					callback("", e)
					return
				}
				var result, error = PostGoFastDfsFile(goFastDfsUrl, part.FileName(), f)
				if callback != nil {
					callback(result, error)
					return
				}
			}
		}
	}
	return hand
}
