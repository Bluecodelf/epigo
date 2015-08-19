package user

import (
    "errors"
    "strings"
    "net/url"
    "net/http"
    "net/http/cookiejar"
)

type User struct {
    username string
}

type UserAuthenticator struct {
    user User
    client http.Client
}

func (p *UserAuthenticator) Authenticate(host string, pass string) (err error) {
    data := url.Values {
        "login": {p.user.username},
        "password": {pass},
    }
    payload := data.Encode()

    request, _ := http.NewRequest("POST", host, strings.NewReader(payload))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    var resp *http.Response
    p.client.Jar, _ = cookiejar.New(nil);
    resp, err = p.client.Do(request)
    if err != nil {
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        err = errors.New("HTTP Error " + resp.Status)
    }
    return
}

func (p *UserAuthenticator) AuthorizeRequest(
    request *http.Request) (*http.Response, error) {
    return p.client.Do(request)
}
