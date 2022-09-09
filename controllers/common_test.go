package controllers

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()

	if withTemplates {
		r.Static("/static", "../static")
		r.LoadHTMLGlob("../static/templates/*")
	}
	return r
}

func payloadValid() string {
	params := url.Values{}
	params.Add("email", "testing1234@testing.com")
	params.Add("password", "testing1234")
	return params.Encode()
}
func payloadInvalid() string {
	params := url.Values{}
	params.Add("email", "testing")
	params.Add("password", "testing123")
	return params.Encode()
}
