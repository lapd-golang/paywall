package wall

import "net/url"

func HasQuota(UserIdentifier string, URL *url.URL) bool {
	// Check if the user has enough quota
	return false
}

func PostRequest(UserIdentifier string, URL *url.URL) bool {
	// After Request
	// Quota = Quota - 1
	return true
}

func BootStrap(UserIdentifier string, URL *url.URL) bool {
	// Add URL to database
	return true
}

func ChargeQuota(UserIdentifier string, URL *url.URL) bool {
	return true
}
