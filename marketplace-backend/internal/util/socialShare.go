package util

import (
	"net/url"
)

// GenerateSocialShareURLs prepares URLs for sharing content on social media.
func GenerateSocialShareURLs(contentTitle, contentURL string) map[string]string {
	encodedTitle := url.QueryEscape(contentTitle)
	encodedURL := url.QueryEscape(contentURL)

	return map[string]string{
		"facebook": "https://www.facebook.com/sharer/sharer.php?u=" + encodedURL,
		"twitter":  "https://twitter.com/intent/tweet?text=" + encodedTitle + "&url=" + encodedURL,
		"linkedin": "https://www.linkedin.com/shareArticle?mini=true&url=" + encodedURL + "&title=" + encodedTitle,
		// Add more platforms as needed
	}
}
