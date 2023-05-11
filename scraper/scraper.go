package scraper

import (
	// "fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var arrCard []modelCard

func Start() []modelCard {
	return getLinks("https://www.7tarot.es/cartas/")
}

// getLinks get the links of each card, receive the main link.
func getLinks(url string) []modelCard {

	log.Print("empezamos GetLinks\n")

	l := colly.NewCollector()

	l.OnHTML("#autres > div:nth-child(1) > ul:nth-child(2) > li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		urlImg := strings.Split(url, "arcanos")[0]
		nombre := e.ChildAttr("img", "alt")
		imagen := urlImg + e.ChildAttr("img", "src")
		palo := strings.Split(nombre, " de ")[len(strings.Split(nombre, " de "))-1]
		numero := strings.Split(strings.Split(e.ChildAttr("img", "src"), "-")[0], "/")[2]
		var arcano string
		n, _ := strconv.Atoi(numero)
		if n <= 22 {
			arcano = "Mayor"
		}
		if n > 22 {
			arcano = "Menor"
		}

		pag := urlImg + link
		// fmt.Println("----------------------------------------------------")
		// log.Print("♦ New card:", numero)

		// fmt.Printf("👉nombre: %v\n", nombre)
		// fmt.Printf("len(nombre): %v\n", len(nombre))
		// fmt.Printf("len(palo): %v\n", len(strings.Split(nombre, " de ")))
		// fmt.Printf("👉palo: %v\n", palo)
		// fmt.Printf("👉imagen: %v\n", imagen)

		descripcion, Sig_al_derecho, Sig_al_revés, Inter_al_derecho, Inter_al_revés := getInfoArc(pag)

		// if n == 22 {
		// 	n = 0
		// }
		if strings.Contains(strings.ToLower(palo), "el") || strings.Contains(strings.ToLower(palo), "la") {
			palo = ""
		}
		// if !strings.Contains(palo, "de") {
		// 	palo = ""
		// }

		card := modelCard{
			"DEFAULT",
			&arcano,
			int16(n),
			&palo,
			&nombre,
			&descripcion,
			&Sig_al_derecho,
			&Sig_al_revés,
			&Inter_al_revés,
			&Inter_al_derecho,
			&imagen,
		}

		arrCard = append(arrCard, card)

		// fmt.Println("----------------------------------------------------")

		// fmt.Println("palo: ", palo)
		// fmt.Println("nombre: ", nombre)
		// fmt.Println("imagen: ", imagen)

	})

	l.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	l.OnResponse(func(r *colly.Response) {
		// log.Println("Visited", r.Request.URL)
	})

	l.OnScraped(func(r *colly.Response) {
		// fmt.Printf("arrCard: %v\n", arrCard[50])
		// log.Println("Finished", r.Request.URL)
	})

	l.Visit(url)
	return arrCard

}

// getInfoArc get info of the card, receive the link of specific card.
func getInfoArc(url string) (descripcion, Sig_al_derecho, Sig_al_revés, Inter_al_derecho, Inter_al_revés string) {

	iM := colly.NewCollector()

	iM.OnHTML("body.arcane", func(e *colly.HTMLElement) {
		descripcion = e.ChildText("div.bloc-content > div.noir")
		Sig_al_derecho = e.ChildText("div.bloc-content ul li p.liste:nth-child(2)")
		Sig_al_revés = e.ChildText("div.bloc-content ul li p.liste:nth-child(4)")
		Inter_al_derecho = e.ChildText("#interpretation > div:nth-child(1) > div:nth-child(2) > div:nth-child(3)")
		Inter_al_revés = e.ChildText("div.bloc-content:nth-child(3) > div:nth-child(3)")
		// fmt.Printf("👉descripcion: %v\n", descripcion)
		// fmt.Printf("👉Sig_al_derecho: %v\n", Sig_al_derecho)
		// fmt.Printf("👉Sig_al_revés: %v\n", Sig_al_revés)
		// fmt.Printf("👉Inter_al_derecho: %v\n", Inter_al_derecho)
		// fmt.Printf("👉Inter_al_revés: %v\n", Inter_al_revés)
	})

	iM.OnError(func(_ *colly.Response, err error) {
		// log.Println("Something went wrong:", err)
	})

	iM.OnResponse(func(r *colly.Response) {
		// log.Println("Visited", r.Request.URL)
	})

	iM.OnScraped(func(r *colly.Response) {
		// log.Println("Finished", r.Request.URL)
	})

	iM.Visit(url)

	return descripcion, Sig_al_derecho, Sig_al_revés, Inter_al_derecho, Inter_al_revés
}
