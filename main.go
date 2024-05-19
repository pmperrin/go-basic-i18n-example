package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	basicI18n "github.com/pmperrin/go-basic-i18n/i18n"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	i18nInstance, err := basicI18n.InitI18n("./translate", "messages", "fr")
	if err != nil {
		fmt.Println(err.Error())
	}

	r.GET("/", func(c *gin.Context) {

		userLang, _ := c.GetQuery("lang")
		lang, err := i18nInstance.GetLang(userLang)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(lang.GetText("signup.title"))
		fmt.Println(lang.GetText("signup.button"))

		c.HTML(http.StatusOK, "index.html", gin.H{
			"lang": lang,
		})
	})

	r.Run(":8080")
}
