package zyxxq

import (
	"net/http"
	"time"
)

type Message struct {
	PubHash  string
	EditHash string
	Content  string
	Date     time.Time
	Modified time.Time
}

func generatePubHash(content string, date time.Time) string {
	hashContent := []byte(content)
	bytes, _ := time.GobEncode(date)
	hasher := sha1.New()
	hasher.Write(content)
	hasher.Write(date)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func generateEditHash(pubHash string) {
	rand.Seed(90245)
	r := rand.Int63n(math.MaxInt64)
	hasher := sha1.New()
	hasher.Write(pubHash)
	hasher.Write(r)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/resources/", resourceHandler)
	http.HandleFunc("/post", postHandler)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	content := r.FormValue("content")
	if content == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	date := time.Now()
	pubHash := generatePubHash(content, date)
	editHash := generateEditHash(pubHash)
	m := Message{
		PubHash:  pubHash,
		EditHash: editHash,
		Content:  content,
		Date:     date,
		Modified: date,
	}
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "resources/index.html")
}
