package http

import (
	"net/http"
	"strings"
)

// contentTypeAllowList is the list of allowed content types in lowercase
var contentTypeAllowListLowerCase = []string{
	"application/json",
	"application/x-www-form-urlencoded",
}

// shouldRecordBodyOfContentType checks if the body is meant
// to be recorded based on the content-type and the fact that body is
// not streamed.
func shouldRecordBodyOfContentType(h http.Header) bool {
	var contentTypeValues = h["Content-Type"] // "Content-Type" is the canonical key

	// we iterate all the values as userland code add the headers in the inverse order,
	// e.g.
	// ```
	//    header.Add("content-type", "charset=utf-8")
	//    header.Add("content-type", "application/json")
	// ```
	for _, contentTypeValue := range contentTypeValues {
		for _, contentTypeAllowedLowerCase := range contentTypeAllowListLowerCase {
			// userland code can set joint headers directly instead of adding
			// them for example
			//
			// ```
			//   header.Set("content-type", "application/json; charset=utf-8")
			// ```
			//
			// instead of
			//
			// ```
			//    header.Add("content-type", "application/json")
			//    header.Add("content-type", "charset=utf-8")
			// ```
			// hence we need to inspect it by using contains.
			if strings.Contains(strings.ToLower(contentTypeValue), contentTypeAllowedLowerCase) {
				return true
			}
		}
	}
	return false
}
