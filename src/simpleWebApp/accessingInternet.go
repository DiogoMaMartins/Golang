package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)


type SiteMapIndex struct {
	Locations []Location `xml:"url"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return  fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://oldtowntec.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}


/*func main() {
	resp, _ := http.Get("https://oldtowntec.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
}*/
