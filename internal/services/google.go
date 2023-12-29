package services

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:9090/callback-gl",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	googleOauthStateString = ""
)

func InitialzeGoogleOauth() {
	googleOauthConfig.ClientID = viper.GetString("google.clientID")
	googleOauthConfig.ClientSecret = viper.GetString("google.clientSecret")
	googleOauthStateString = viper.GetString("oauthStateString")
}

func HandleGoogleLogin(c *fiber.Ctx) error {
	return HandleLogin(c, &googleOauthConfig, googleOauthStateString)
}

func CallBackFromGoogle(c *fiber.Ctx) error {
	state := c.Query("state")
	log.Println(state)
	if state != googleOauthStateString {
		log.Println("invalid oauth state, expected " + googleOauthStateString + ", got " + state + "\n")
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	code := c.Query("code")
	log.Println(code)

	if code == "" {
		log.Println("Code not found..")
		return c.Send([]byte("Code Not Found to provide AccessToken..\n"))
	}

	token, err := googleOauthConfig.Exchange(context.TODO(), code)
	if err != nil {
		log.Println("googleOauthConfig.Exchange() failed with " + err.Error() + "\n")
		return c.Status(fiber.StatusInternalServerError).SendString("Error exchanging code for token")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		log.Println("Get: " + err.Error() + "\n")
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll: " + err.Error() + "\n")
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	log.Println("parseResponseBody: " + string(response) + "\n")

	return c.Redirect("/form")
}
