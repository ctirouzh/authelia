package utils

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

// URLPathFullClean returns a URL path with the query parameters appended (full path) with the path portion parsed
// through path.Clean given a *url.URL.
func URLPathFullClean(u *url.URL) (output string) {
	lengthPath := len(u.Path)
	lengthQuery := len(u.RawQuery)
	appendForwardSlash := lengthPath > 1 && u.Path[lengthPath-1] == '/'

	switch {
	case lengthPath == 1 && lengthQuery == 0:
		return u.Path
	case lengthPath == 1:
		return path.Clean(u.Path) + "?" + u.RawQuery
	case lengthQuery != 0 && appendForwardSlash:
		return path.Clean(u.Path) + "/?" + u.RawQuery
	case lengthQuery != 0:
		return path.Clean(u.Path) + "?" + u.RawQuery
	case appendForwardSlash:
		return path.Clean(u.Path) + "/"
	default:
		return path.Clean(u.Path)
	}
}

// URLDomainHasSuffix determines whether the uri has a suffix of the domain value.
func URLDomainHasSuffix(uri url.URL, domain string) bool {
	if uri.Scheme != https {
		return false
	}

	if uri.Hostname() == domain {
		return true
	}

	if strings.HasSuffix(uri.Hostname(), period+domain) {
		return true
	}

	return false
}

// IsRedirectionURISafe determines whether the URI is safe to be redirected to.
func IsRedirectionURISafe(uri, protectedDomain string) (safe bool, err error) {
	var parsedURI *url.URL

	if parsedURI, err = url.ParseRequestURI(uri); err != nil {
		return false, fmt.Errorf("failed to parse URI '%s': %w", uri, err)
	}

	return parsedURI != nil && URLDomainHasSuffix(*parsedURI, protectedDomain), nil
}
