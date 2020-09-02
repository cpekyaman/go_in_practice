package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("templates/file_upload.html"))
}

func main() {
	http.HandleFunc("/", uploadHandler)
	http.ListenAndServe(":8080", nil)
}

func uploadHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		t.Execute(resp, nil)
	} else {
		err := req.ParseMultipartForm(16 << 20)
		if err != nil {
			fmt.Fprint(resp, err)
			return
		}

		data := req.MultipartForm
		files := data.File["files"]

		if len(files) == 0 {
			fmt.Println("No Files")
			fmt.Fprint(resp, "No Files Received")
			return
		}

		for _, fh := range files {
			fmt.Printf("Uploading file %s\n", fh.Filename)

			f, err := fh.Open()
			defer f.Close()
			if err != nil {
				fmt.Fprint(resp, err)
				return
			}

			out, err := os.Create("files/" + fh.Filename)
			defer out.Close()
			if err != nil {
				fmt.Fprint(resp, err)
				return
			}

			_, err = io.Copy(out, f)
			if err != nil {
				fmt.Fprint(resp, err)
				return
			}
		}

		fmt.Fprint(resp, "Upload Complete")

	}
}
