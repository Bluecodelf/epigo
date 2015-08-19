package epigo

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func (client *Client) Authenticate(login string, password string) (err error) {
	data := url.Values{
		"login":    {login},
		"password": {password},
	}
	payload := data.Encode()

	request, _ := http.NewRequest("POST", client.host,
		strings.NewReader(payload))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	var resp *http.Response
	client.http.Jar, _ = cookiejar.New(nil)
	resp, err = client.http.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("HTTP Error " + resp.Status)
	}
	return
}
