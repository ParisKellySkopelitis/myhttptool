# MYHTTPTOOL - gets md5 of http response of websites

A tool that gives you the md5 of a htttp response of a url. Supports parallel requests

## Installation
- Requirements: Golang 1.20
To build the tool run the command `go build`

## Usage
To run use the command `./myhttptool <url>` eg `./myhttptool http://google.com`
The tool also supports parallel calls with the optional parallel flag eg `.myhttptool -parallel 2 http://google.com http://twitter.com`, parallel is 2 since there are two requests. The maximum supported parallel calls is 10.

## Testing

To test the tool use the command `go test myhttptool/client -v`

## Possible Improvements

More unit testing on the worker and to test parallel input flag. Test Failures for client.