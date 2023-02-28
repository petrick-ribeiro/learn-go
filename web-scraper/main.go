package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	dt := time.Now()

	// Enable the domain.
	c := colly.NewCollector(
		colly.AllowedDomains("weather.com"),
	)

	// Get Weather infos.
	c.OnHTML(".CurrentConditions--primary--2DOqs", func(h *colly.HTMLElement) {
		weather := h.ChildText(".CurrentConditions--tempValue--MHmYY")
		phrase := h.ChildText(".CurrentConditions--phraseValue--mZC_p")

		fmt.Println("")
		fmt.Println("🌤️", weather)
		fmt.Println(phrase)
		fmt.Println("Curitiba -", dt.Format("02/01/2006 15:04"))
	})

	// Get the HTTP status code.
	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Status Code:", r.StatusCode)
	// 	fmt.Println("Status Code:", r.Request.URL)
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting...")
	// })

	// Define the url and acess.
	c.Visit("https://weather.com/pt-BR/clima/hoje/l/cd456e246b710e10cd019303ee89dd7486d87c3f89d1b01b0c96e0929fe4b296")
}
