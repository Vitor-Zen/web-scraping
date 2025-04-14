// Funções das rotas
package auth

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	log.Println("ClientID usado:", GoogleOAuthConfig.ClientID)
	url := GoogleOAuthConfig.AuthCodeURL("random-state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}
