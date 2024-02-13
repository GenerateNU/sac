package scraper

import (
	"net/http"
)

func RetrieveClubs() {
	resp, _ := http.Get("https://neu.campuslabs.com/engage/api/discovery/search/organizations?orderBy%5B0%5D=UpperName%20asc&top=10&filter=&query=&skip=0")
	print(resp)
}
