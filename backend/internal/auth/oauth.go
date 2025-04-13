// Contém configuração do cliente OAuth2 com os dados fornecidos pelo Google (Client ID, Secret, etc)
package auth

import (
	"os"

	"golang.org/x/oauth2"        // fornece toda a estrutura do protocolo OAuth2 (cria configs, gerar URLs de login, trocar tokens etc).
	"golang.org/x/oauth2/google" // traz valores padrão específicos do Google (como endpoints para login, troca de token e obtenção de perfil).
)

var GoogleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}
