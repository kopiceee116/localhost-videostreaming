package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func listazd(basepath string) []string {
	var movielist = []string{}
	files, err := os.ReadDir(basepath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		movielist = append(movielist, fmt.Sprint(file.Name()))
	}
	return movielist
}

func main() {

	const (
		bufferSize int16  = 8192 // Adjust buffer size as needed
		moviePpath string = "./movies"
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>koppany csinalta</h1> <hr>")
		movielist := listazd(moviePpath)
		for _, movie := range movielist {
			fmt.Fprintf(w, "<a href='movies/%s'style='font-size: xx-large;'> %s </a> <br>\n", movie, movie)
		}
		return
	})

	http.HandleFunc("/movies/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path
		requestedFilePath := filepath.Join(moviePpath, requestedFile[8:]) // remove slashmovies from beginning

		videoFile, err := os.Open(requestedFilePath)
		if err != nil {
			http.Error(w, "Could not open video file", http.StatusInternalServerError)
			return
		}
		defer videoFile.Close()

		stat, err := videoFile.Stat()
		if err != nil {
			http.Error(w, "Could not get video file information", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "video/mp4")
		http.ServeContent(w, r, "", stat.ModTime(), videoFile)
	})

	log.Println("go to localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
