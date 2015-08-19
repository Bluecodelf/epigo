package epigo

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	host string
	http http.Client
}

func (client *Client) GetData(url string) (data []byte, err error) {
	var resp *http.Response
	resp, err = client.http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("HTTP Error " + resp.Status)
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	return
}
