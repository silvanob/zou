package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Image struct {
	Filename string
	Data     []byte
}

func generateFilename() []byte {
	rand.Seed(time.Now().UnixNano())
	token := make([]byte, 6)
	rand.Read(token)
	return token
}

func (image *Image) save() error {
	return ioutil.WriteFile(image.Filename, image.Data, 0600)
}

func loadImage(title string) (*Image, error) {
	filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Image{Filename: filename, Data: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadImage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div><%s</div>", p.Filename, p.Data)
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println(r.Body)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		filename := "5.png"
		image := Image{Filename: filename, Data: body}
		err = image.save()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	fmt.Println(generateFilename())
	http.HandleFunc("/post", post)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
