package login

import (
	"encoding/json"
	"log"
	"os"

	"github.com/romaingallez/scform_api/internals/models"

	"github.com/go-rod/rod/lib/proto"
)

func LoginWithCookies(page models.Page) (models.Page, error) {

	// test if page.User.Cookies is empty
	if len(page.User.Cookies) == 0 {
		return page, nil
	}
	cookies := page.User.Cookies

	// // test if cookies.json exists
	// _, err := os.Stat("cookies.json")
	// if err != nil {
	// 	return page, nil
	// }

	// cookies := ReadJsonCookies("cookies.json")

	for _, cookie := range cookies {
		page.Page.Browser().MustSetCookies(cookie)
	}

	// reload page
	page.Page.MustReload().MustWaitDOMStable()

	// utils.Countdown(5, time.Second)

	return page, nil

}

func ReadJsonCookies(path string) []*proto.NetworkCookie {

	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var cookies []*proto.NetworkCookie
	err = json.Unmarshal(jsonBytes, &cookies)
	if err != nil {
		log.Fatal(err)
	}

	return cookies
}

func WriteJsonCookies(path string, cookies []*proto.NetworkCookie) {

	jsonBytes, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path, jsonBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
