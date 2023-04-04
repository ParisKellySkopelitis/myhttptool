package mockdata

import (
	"bytes"
	"io"
	"net/http"
)

// ClientMock struct to mimic real request.
// (Stackoverflow helped...) https://stackoverflow.com/questions/43240970/how-to-mock-http-client-do-method
type ClientMock struct{}

func (mock *ClientMock) Do(req *http.Request) (*http.Response, error) {
	content := returnBody(req.URL.String())
	return &http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(content))),
	}, nil
}

func ReturnMD5HashOfBody(url string) string {
	var hash string

	//md5 of the body text from the returnBody()
	switch url {
	case "http://reddit.com/funny":
		hash = "c4ce32288a0b2ba40f87e7258cdacf70"
	case "http://twitter.com":
		hash = "e09c80c42fda55f9d992e59ca6b3307d"
	case "http://youtube.com":
		hash = "4720bc99c310cb289ed3832e5bff9532"
	}
	return hash
}

func returnBody(url string) string {
	var body string
	switch url {
	case "http://reddit.com/funny":
		body = "old memes from 2014"
	case "http://twitter.com":
		body = "aaaaaaaaaa"
	case "http://youtube.com":
		body = "whats up guys! loud intro!!"
	}
	return body
}
