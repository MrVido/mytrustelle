package util

import (
	"fmt"
	"net/url"
)

// Your application's logo and website URL. Ensure these URLs are publicly accessible for social platforms to fetch.
const (
	appLogoURL   = "https://yourdomain.com/path/to/logo.png"
	websiteURL   = "https://yourdomain.com"
	websiteTitle = "Your Site Name"
)

// GenerateSocialShareURLs prepares URLs for sharing content on multiple social media platforms, including an image and detailed descriptions.
func GenerateSocialShareURLs(contentTitle, contentDescription, contentURL string) map[string]string {
	encodedTitle := url.QueryEscape(contentTitle)
	encodedDescription := url.QueryEscape(contentDescription)
	encodedURL := url.QueryEscape(contentURL)
	encodedImage := url.QueryEscape(appLogoURL)

	return map[string]string{
		"facebook": fmt.Sprintf("https://www.facebook.com/sharer/sharer.php?u=%s", encodedURL),
		"twitter":  fmt.Sprintf("https://twitter.com/intent/tweet?text=%s&url=%s&via=%s", encodedTitle, encodedURL, websiteTitle),
		"linkedin": fmt.Sprintf("https://www.linkedin.com/shareArticle?mini=true&url=%s&title=%s&summary=%s&source=%s", encodedURL, encodedTitle, encodedDescription, websiteURL),
		"pinterest": fmt.Sprintf("https://pinterest.com/pin/create/button/?url=%s&media=%s&description=%s", encodedURL, encodedImage, encodedTitle),
		"whatsapp": fmt.Sprintf("https://api.whatsapp.com/send?text=%s %s", encodedTitle, encodedURL),
		"telegram": fmt.Sprintf("https://telegram.me/share/url?url=%s&text=%s", encodedURL, encodedTitle),
	}
}

// Additional Tips:
// - For SEO: Utilize Open Graph tags (for Facebook) and Twitter Cards in your HTML <head> to define how URLs are displayed when shared. Include tags for the title, type, image, and description.
// - Regularly update your social media strategy to include newer platforms and technologies for sharing content.
// - Monitor how your content is shared and engaged with on social media to refine your approach continuously.
