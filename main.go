package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Inisialisasi collector
	c := colly.NewCollector()

	// Slice untuk menyimpan daftar proyek
	var projects []string

	// Buat regex untuk menghapus "Project #X."
	re := regexp.MustCompile(`Project #\d+\.`)

	// Menangkap hanya elemen <h2> yang relevan
	c.OnHTML("h2", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)                 // Menghapus spasi berlebih
		text = strings.ReplaceAll(text, "⭐", "")         // Menghapus simbol "⭐"
		text = re.ReplaceAllString(text, "")             // Menghapus "Project #X."
		text = strings.TrimSpace(text)                   // Trim lagi setelah replace

		// Tambahkan ke daftar proyek jika tidak kosong
		if text != "" {
			projects = append(projects, text)
		}
	})

	// Menangani error
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	// Kunjungi website target
	url := "https://zerotomastery.io/blog/golang-practice-projects/"
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	// Menampilkan hasil yang lebih rapi
	fmt.Println("\nDaftar Proyek yang Dapat Dibuat Sebagai Portfolio Golang:")
	for i, project := range projects {
		fmt.Printf("%d. %s\n", i+1, project)
	}
}
