package identity

import (
	"strings"
	"time"
	"crypto/sha256"
	"encoding/base64"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

var xCode string = ""
var xChallenge string = ""
var mChallenge string = "plain"

func AuthorizationMiddleware(c *fiber.Ctx) error {
	color.Cyan("[DEBUG] Received request at /authorize endpoint")
	clientID := c.Query("client_id")
	redirectURI := c.Query("redirect_uri")
	codeChallenge := c.Query("code_challenge")
	codeChallengeMethod := c.Query("code_challenge_method")

	xChallenge = codeChallenge
	mChallenge = codeChallengeMethod

	color.Blue("[DEBUG] client_id: %s, redirect_uri: %s, code_challenge: %s, code_challenge_method: %s", clientID, redirectURI, codeChallenge, codeChallengeMethod)

	if clientID == "" || redirectURI == "" || codeChallenge == "" || codeChallengeMethod == "" {
		color.Red("[ERROR] Missing required parameters")
		return c.Status(fiber.StatusBadRequest).SendString("Missing required parameters")
	}

	authCode := generateAuthorizationCode() 
	xCode = authCode
	color.Green("[DEBUG] Generated authorization code: %s", authCode)

	color.Yellow("[DEBUG] Redirecting to: %s?code=%s", redirectURI, authCode)
	return c.Redirect(redirectURI + "?code=" + authCode)
}

func TokenMiddleware(c *fiber.Ctx) error {
	color.Cyan("[DEBUG] Received request at /token endpoint")
	code := c.FormValue("code")
	codeVerifier := c.FormValue("code_verifier")

	color.Blue("[DEBUG] code: %s, code_verifier: %s", code, codeVerifier)

	if code == "" || codeVerifier == "" {
		color.Red("[ERROR] Missing required parameters")
		return c.Status(fiber.StatusBadRequest).SendString("Missing required parameters")
	}

	if code != xCode {
		color.Red("[ERROR] Invalid code")
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid code")
	}

	if mChallenge == "plain" && codeVerifier != xChallenge{
		color.Red("[ERROR] Invalid code verifier")
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid code verifier")
	}

	if mChallenge == "S256" &&  hashStringWithSHA256AndBase64URL(codeVerifier) != xChallenge {
			color.Red("[ERROR] Invalid code verifier")
			color.Red("\tExpected: %s", hashStringWithSHA256AndBase64URL(codeVerifier))
			color.Red("\tVerify  : %s", xChallenge)
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid code verifier")
	}

	accessToken := "mockAccessToken123"
	color.Green("[DEBUG] Generated access token: %s", accessToken)

	return c.JSON(fiber.Map{
		"access_token": accessToken,
	})
}

func hashStringWithSHA256AndBase64URL(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	encodedHash := base64.URLEncoding.EncodeToString(hashBytes)
	encodedHash = strings.TrimRight(encodedHash, "=")
	return encodedHash
}

func generateAuthorizationCode() string {
	hasher := sha256.New()
	hasher.Write([]byte(strings.ReplaceAll(color.HiGreenString("%d", time.Now().UnixNano()), " ", "")))
	code := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	code = strings.TrimRight(code, "=")
	return code
}