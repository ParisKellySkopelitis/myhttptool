package client

import (
	"myhttptool/mockdata"
	"testing"
)

// TestGetMD5Page tests the GetMD5Page function from the client package
func TestGetMD5Page(t *testing.T) {
	testCases := []struct {
		testDescription string
		url             string
		outputUrl       string
		function        func(*testing.T, RequestedPage, string)
	}{
		{
			testDescription: "Test Case: Success with a correct URL",
			url:             "http://reddit.com/funny",
			outputUrl:       "http://reddit.com/funny",
			function:        Success,
		},
		{
			testDescription: "Test Case: Sucess without http in url",
			url:             "twitter.com",
			outputUrl:       "http://twitter.com",
			function:        Success,
		},
	}

	client := ClientAPI{
		Client: &mockdata.ClientMock{},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testDescription, func(t *testing.T) {
			mockPage := client.GetMD5Page(testCase.url)
			testCase.function(t, mockPage, testCase.outputUrl)
		})
	}
}

func Success(t *testing.T, mockPage RequestedPage, expectedUrl string) {
	// Check Url gets converted (http://)
	if expectedUrl != mockPage.url {
		t.Errorf("wrong url. URL expected %s, URL received %s", expectedUrl, mockPage.url)
	} else {
		//Check if hashing is correct
		expectedHash := mockdata.ReturnMD5HashOfBody(expectedUrl)
		if expectedHash != mockPage.md5Hash {
			t.Errorf("wrong hash. hash expected %s, URL received %s", expectedHash, mockPage.md5Hash)
		}

	}
}
