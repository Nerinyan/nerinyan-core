package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type Auth struct {
	username string
	password string
	Token    struct {
		TokenType    string `json:"token_type"`
		ExpiresIn    int64  `json:"expires_in"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ParsedToken  struct {
			Aud    string   `json:"aud"`
			Jti    string   `json:"jti"`
			Iat    float64  `json:"iat"`
			Nbf    float64  `json:"nbf"`
			Exp    float64  `json:"exp"`
			Sub    string   `json:"sub"`
			Scopes []string `json:"scopes"`
		} `json:"-"`
	} `json:"token"`
}

func LoginWithAuth(username string, password string) (auth Auth, err error) {
	auth.username = username
	auth.password = password
	err = auth.Login()
	return
}
func (a *Auth) Login() (err error) {
	url := "https://osu.ppy.sh/oauth/token"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("client_id", "5")
	_ = writer.WriteField("client_secret", "FGc9GAtyHzeQDshWP5Ah7dega8hJACAJpQtw6OXk")
	_ = writer.WriteField("scope", "*")

	_ = writer.WriteField("username", a.username)
	_ = writer.WriteField("password", a.password)
	_ = writer.WriteField("grant_type", "password")

	err = writer.Close()
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &a.Token)
	return
}
func (a *Auth) Refresh() (err error) {
	url := "https://osu.ppy.sh/oauth/token"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("client_id", "5")
	_ = writer.WriteField("client_secret", "FGc9GAtyHzeQDshWP5Ah7dega8hJACAJpQtw6OXk")
	_ = writer.WriteField("scope", "*")

	_ = writer.WriteField("grant_type", "refresh_token")
	_ = writer.WriteField("refresh_token", a.Token.RefreshToken)

	err = writer.Close()
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", a.Token.TokenType+" "+a.Token.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	return json.Unmarshal(body, &a.Token)
}

func (a *Auth) ExpiredAt() (secondsRemainingUntilExpiration int) {
	s := strings.Split(a.Token.AccessToken, ".")
	if len(s) != 3 {
		return
	}

	decodeString, err := base64.RawStdEncoding.DecodeString(s[1])
	if err != nil {
		decodeString, err = base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			decodeString, err = base64.RawURLEncoding.DecodeString(s[1])
			if err != nil {
				decodeString, err = base64.URLEncoding.DecodeString(s[1])
				if err != nil {
					return
				}
			}
		}
	}

	if err = json.Unmarshal(decodeString, &a.Token.ParsedToken); err != nil {
		return
	}
	return int(int64(a.Token.ParsedToken.Exp) - time.Now().Unix())
}
