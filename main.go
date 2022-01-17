package main

import (
	"encoding/base64"
	"net/http"
)

const Gif = "\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\xFF\x00\xFF\xFF\xFF\x00\x00\x00\x2C\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02\x44\x01\x00\x3B\x00"

const Png = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR4nGP6zwAAAgcBApocMXEAAAAASUVORK5CYII="

func main() {
	http.HandleFunc("/pixel.gif", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/gif")
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		w.Write([]byte(Gif))
	})

	http.HandleFunc("/pixel.png", func(w http.ResponseWriter, r *http.Request) {
		dat, _ := base64.StdEncoding.DecodeString(Png)
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		w.Write(dat)
	})
	http.ListenAndServe(":9800", nil)
}
