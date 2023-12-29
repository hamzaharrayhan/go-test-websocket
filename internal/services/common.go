package services

import (
	"jti-test/internal/helpers/pages"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func HandleLogin(c *fiber.Ctx, oauthConf *oauth2.Config, oauthStateString string) error {
	authURL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Println("Parse:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing authorization URL")
	}

	params := url.Values{}
	params.Add("client_id", oauthConf.ClientID)
	params.Add("scope", strings.Join(oauthConf.Scopes, " "))
	params.Add("redirect_uri", oauthConf.RedirectURL)
	params.Add("response_type", "code")
	params.Add("state", oauthStateString)
	authURL.RawQuery = params.Encode()

	log.Println("Redirecting to:", authURL.String())
	return c.Redirect(authURL.String(), fiber.StatusTemporaryRedirect)
}

func HandleMain(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	c.Status(http.StatusOK)
	_, err := c.Response().BodyWriter().Write([]byte(pages.IndexPage))
	return err
}

func HandleFormPage(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	c.Status(http.StatusOK)
	_, err := c.Response().BodyWriter().Write([]byte(pages.FormPage))
	return err
}

func HandleDashboardPage(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	c.Status(http.StatusOK)
	_, err := c.Response().BodyWriter().Write([]byte(pages.DashboardPage))
	return err
}

func HandleUpdatePage(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	c.Status(http.StatusOK)
	_, err := c.Response().BodyWriter().Write([]byte(pages.UpdatePage))
	return err
}
