package login

import (
	"github.com/romaingallez/scform_api/internals/models"
)

func LoginWithPassword(page models.Page) (models.Page, error) {

	// 	page.Page.MustElementByJS(`document.querySelector("#MainContent_LoginUser_UserName")`).MustInput(page.User.Username)
	page.Page.MustElementByJS(`() => document.querySelector("#MainContent_LoginUser_UserName")`).MustInput(page.User.Username)
	page.Page.MustElementByJS(`() => document.querySelector("#MainContent_LoginUser_Password")`).MustInput(page.User.Password)
	page.Page.MustEval(`() => LoginBt()`)

	page.Page.MustWaitDOMStable()

	cookies := page.Page.Browser().MustGetCookies()

	page.User.Cookies = cookies
	// WriteJsonCookies("cookies.json", cookies)
	return page, nil
}
