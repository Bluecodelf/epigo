package epigo

import (
    "net/http"
)

type Client struct {
    host string
    http http.Client
}
