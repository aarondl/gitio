/*
Package gitio shortens github urls using the git.io service.
*/
package gitio

import (
	"fmt"
	"net/http"
	"net/url"
)

// Shorten a long github url.
func Shorten(longurl string) (string, error) {
	resp, err := http.PostForm("http://git.io", url.Values{`url`: {longurl}})
	if err != nil {
		return "", nil
	}
	if resp.StatusCode != 201 {
		return "", fmt.Errorf("Expected 201 response, got: %d", resp.StatusCode)
	}
	return resp.Header.Get("Location"), nil
}
