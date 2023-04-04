package client

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// ClientAPI is the base struct for the http calls
type ClientAPI struct {
	Client Client
}

// Client is the interface so that we can also mock for testing
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

// RequestPage is the base struct containing the url and the md5Hash
type RequestedPage struct {
	url     string
	md5Hash string
}

// String is a custom string function to print the url and hash of the requestedPage
func (r RequestedPage) ToString() string {
	return fmt.Sprintf("%s %s", r.url, r.md5Hash)
}

// makeHTTPRquest does the actual HTTP request
func (c ClientAPI) makeHTTPRequest(url string) ([]byte, error) {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(request)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// getMD5Hash is a simple MD5Hashing function
func getMD5Hash(bytes []byte) string {
	hash := md5.Sum(bytes)
	return hex.EncodeToString(hash[:])
}

// GetMD5Page makes the request and returns the url after sanitisation and the hash as a RequestedPage struct
func (c ClientAPI) GetMD5Page(url string) RequestedPage {

	if !strings.Contains(url, "http://") {
		url = "http://" + url
	}

	res, err := c.makeHTTPRequest(url)
	if err != nil {
		log.Fatalf("request failed, %v", err)
	}
	md5Hash := getMD5Hash(res)
	return RequestedPage{url: url, md5Hash: md5Hash}
}
