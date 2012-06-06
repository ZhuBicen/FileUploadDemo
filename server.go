package main

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	"io"
	"log"
	"net/http"
	//"strings"
	//"path/filepath"
	"os"
)

var G_COUNT = 0

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("HomeHandler")
	file, err := os.Open("app.html")
	if err != nil{
		log.Println("ERR=", err)
		return
	}
	io.Copy(w, file)
	return
}

// hello world, the web server
func UploadHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("\nReceive %d request\n\n", G_COUNT)
	G_COUNT++

	//join the filename to the upload dir
	// saveToFilePath := filepath.Join("d:/TMP", saveToFilename)

	// osFile, err := os.Create(saveToFilePath)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer osFile.Close()

	// count, err := io.Copy(osFile, formFile)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("ALLOW: %s SAVE: %s (%d)\n", req.RemoteAddr, saveToFilename, count)
	w.Write([]byte("Upload Complete for"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/upload", UploadHandler)
	//r.HandleFunc("/{sessionId:[0-9]+}", SessionHandler)
	//r.HandleFunc(`/{sessionId:[0-9]+}/{fileName:.*\.html}`, SessionHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.Handle("/", r)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
