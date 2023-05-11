package server

import (
	"encoding/json"
	"fmt"

	// "strconv"

	// "html"
	"log"
	"net/http"

	"github.com/GiampieroFC/db-tarot-v1/helpers"
	"github.com/GiampieroFC/db-tarot-v1/scraper"
)

var (
	db *scraper.DataBase
)

func isBadMethod(w http.ResponseWriter, r *http.Request) bool {
	if err := helpers.MethodChecker(r.Method, helpers.Get); err != "" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(err))
		return true
	}
	return false
}

func init() {
	db = scraper.NewDB()
	db.Connect()

}

func ToServe(port string) {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/", main)
	mux.HandleFunc("/api/v1/all", getAll)
	mux.HandleFunc("/api/v1/number", getByNumber)
	mux.HandleFunc("/api/v1/suit", getByPalo)
	mux.HandleFunc("/api/v1/minor", getMenor)
	mux.HandleFunc("/api/v1/major", getMayor)
	mux.HandleFunc("/api/v1/random", getRandom)

	fmt.Printf("Listening and serving at http://localhost:%s/api/v1 ", port)

	http.ListenAndServe(":"+port, mux)
}

func main(w http.ResponseWriter, r *http.Request) {

	if isBadMethod(w, r) {
		return
	}

	const msg = "Endpoints:\n/all/\n/number/\n/suit/\n/minor/\n/major/\n/random/"
	w.Write([]byte(msg))
}

func getAll(w http.ResponseWriter, r *http.Request) {

	if isBadMethod(w, r) {
		return
	}

	cards, err := db.GetAll()
	if err != nil {
		log.Fatalf("db.getAll() main: %v", err)
	}

	bytes, _ := json.Marshal(cards)
	w.Write(bytes)
	// fmt.Fprintf(w, string(bytes))
}
func getMenor(w http.ResponseWriter, r *http.Request) {

	if isBadMethod(w, r) {
		return
	}

	cards, err := db.GetMenor()
	if err != nil {
		log.Fatalf("db.GetMenor() main: %v", err)
	}

	bytes, _ := json.Marshal(cards)
	w.Write(bytes)
	// fmt.Fprintf(w, string(bytes))
}
func getMayor(w http.ResponseWriter, r *http.Request) {

	if isBadMethod(w, r) {
		return
	}

	cards, err := db.GetMayor()
	if err != nil {
		log.Fatalf("db.GetMayor() main: %v", err)
	}

	bytes, _ := json.Marshal(cards)
	w.Write(bytes)
	// fmt.Fprintf(w, string(bytes))
}
func getByNumber(w http.ResponseWriter, r *http.Request) {

	if isBadMethod(w, r) {
		return
	}

	var n string

	if r.URL.Query().Get("n") == "" {
		n = "22"
	} else {
		n = r.URL.Query().Get("n")
	}

	cards, err := db.GetByNumber(n)

	if err != nil {
		log.Fatalf("db.GetByName(r.URL.Query().Get(name)): %v", err)
	}

	bytes, _ := json.Marshal(cards)
	w.Write(bytes)

	// fmt.Fprintf(w, string(bytes))
}
func getRandom(w http.ResponseWriter, r *http.Request) {

	if isBadMethod(w, r) {
		return
	}

	var ran string

	if r.URL.Query().Get("r") == "" {
		ran = "1"
	} else {
		ran = r.URL.Query().Get("r")
	}

	cards, err := db.GetRandom(ran)

	if err != nil {
		log.Fatalf("db.GetByName(r.URL.Query().Get(name)): %v", err)
	}

	bytes, _ := json.Marshal(cards)
	w.Write(bytes)

	// fmt.Fprintf(w, string(bytes))
}
func getByPalo(w http.ResponseWriter, r *http.Request) {

	if isBadMethod(w, r) {
		return
	}

	// var p string

	// if r.URL.Query().Get("p") == "" {
	// 	p = ""
	// } else {
	p := r.URL.Query().Get("p")
	// }

	cards, err := db.GetByPalo(p)

	if err != nil {
		log.Fatalf("db.GetByName(r.URL.Query().Get(name)): %v", err)
	}

	bytes, _ := json.Marshal(cards)
	w.Write(bytes)

	// fmt.Fprintf(w, string(bytes))
}
