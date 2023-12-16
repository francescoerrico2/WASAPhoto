package api

import (
	"net/http"
	"strings"
)

func validIdentifier(identifier string) bool {
	var trimmedId = strings.TrimSpace(identifier)
	return len(identifier) >= 3 && len(identifier) <= 16 && trimmedId != "" && !strings.ContainsAny(trimmedId, "?_")
}
func extractBearer(authorization string) string {
	var tokens = strings.Split(authorization, " ")
	if len(tokens) == 2 {
		return strings.Trim(tokens[1], " ")
	}
	return ""
}

func validateRequestingUser(identifier string, bearerToken string) int {
	if isNotLogged(bearerToken) {
		return http.StatusForbidden
	}
	if identifier != bearerToken {
		return http.StatusUnauthorized
	}
	return 0
}

func isNotLogged(auth string) bool {

	return auth == ""
}
