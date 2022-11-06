package urlx

import "net/url"

// SignedURL is a shared-key HMAC wrapped URL.
type SignedURL struct {
	uri    url.URL
	key    []byte
	signed bool
}

// NewSignedURL creates a new copy of a URL that can be signed with a shared key.
//
// The key is not copied and must not be changed after the URL is signed.
func NewSignedURL(uri *url.URL, key []byte) *SignedURL {
	return &SignedURL{
		uri: *uri,
		key: key,
	}
}
