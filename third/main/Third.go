package main

import (
	"github.com/zhuxiujia/GoMybatisMall/third/com/third/fs"
	"net/http"
	"os"
	"time"
)

func main() {
	go func() {
		//this is client code! not run in server. like go, java,js
		time.Sleep(2 * time.Second)

		var f, e = os.Open("F:/image/fdsaf.jpg")
		defer f.Close()
		if e != nil {
			println(e.Error())
		}
		s, e := fs.PostGoFastDfsFile("http://127.0.0.1:8000/upload", "new8.jpg", f)
		if e != nil {
			println(e.Error())
		}
		println(s)
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
	var c = fs.UploadHandler("http://localhost:8080/upload", func(result string, e error) {
		println(result)
		if e != nil {
			println(e.Error())
		}
	})
	http.HandleFunc("/upload", c)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
