package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	outputFile    = "public/sitemap.xml"
	productsAPI   = "/api/collections/products/records?perPage=500"
	categoriesAPI = "/api/collections/categories/records?perPage=500"
)

// Structs for config
type Config struct {
	Domains struct {
		Nuxt string `json:"nuxt"`
		PocketBase  string `json:"pocketbase"`
	} `json:"domains"`
}

// PocketBase response types
type PocketBaseResponse[T any] struct {
	Items []T `json:"items"`
}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Category struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Sitemap structs
type Url struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod,omitempty"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
}

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

// Fetch function for PocketBase
func fetch[T any](baseURL, endpoint string, target *[]T) error {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data PocketBaseResponse[T]
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	*target = data.Items
	return nil
}

// Slugify helper
func slugify(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	s = re.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

// Read config file
func loadConfig(path string) (Config, error) {
	var config Config
	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return config, err
	}
	return config, nil
}

// Main logic
func main() {
	// Load config
	config, err := loadConfig("app/pocketstore.json")
	if err != nil {
		log.Fatalf("❌ Failed to read config: %v", err)
	}

	baseURL := "https://"+config.Domains.Nuxt
	pocketbaseURL := "https://"+config.Domains.PocketBase

	fmt.Println("Using config:")
	fmt.Println("Frontend:", baseURL)
	fmt.Println("Backend:", pocketbaseURL)

	var products []Product
	var categories []Category

	fmt.Println("Fetching data from PocketBase...")

	if err := fetch(pocketbaseURL, productsAPI, &products); err != nil {
		log.Fatalf("Failed to fetch products: %v", err)
	}

	if err := fetch(pocketbaseURL, categoriesAPI, &categories); err != nil {
		log.Fatalf("Failed to fetch categories: %v", err)
	}

	urlset := UrlSet{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}

	// Homepage
	urlset.Urls = append(urlset.Urls, Url{
		Loc:        baseURL,
		LastMod:    time.Now().Format("2006-01-02"),
		ChangeFreq: "weekly",
		Priority:   "1.0",
	})

	// Categories
	for _, c := range categories {
		url := fmt.Sprintf("%s/de/category/%s.html", baseURL, c.Slug)
		urlset.Urls = append(urlset.Urls, Url{
			Loc:        url,
			LastMod:    time.Now().Format("2006-01-02"),
			ChangeFreq: "weekly",
			Priority:   "0.8",
		})
	}

	// Products
	for _, p := range products {
		slug := slugify(p.Name)
		url := fmt.Sprintf("%s/de/product/%s.html", baseURL, slug)
		urlset.Urls = append(urlset.Urls, Url{
			Loc:        url,
			LastMod:    time.Now().Format("2006-01-02"),
			ChangeFreq: "weekly",
			Priority:   "0.7",
		})
	}

	// Write sitemap
	if err := os.MkdirAll("public", 0755); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create sitemap file: %v", err)
	}
	defer file.Close()

	file.WriteString(xml.Header)
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	if err := encoder.Encode(urlset); err != nil {
		log.Fatalf("Failed to write XML: %v", err)
	}

	fmt.Printf("✅ Sitemap generated successfully at %s\n", outputFile)
}
