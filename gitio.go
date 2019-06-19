/*
Package gitio shortens github urls using the git.io service.
*/
package gitio

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Shorten a long github url.
func Shorten(longurl string) (string, error) {
	resp, err := http.PostForm(`https://git.io/create`, url.Values{`url`: {longurl}})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Expected 200 response, got: %d", resp.StatusCode)
	}

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response from gitio: %v", err)
	}

	return fmt.Sprintf(`https://git.io/%s`, text), nil
}
