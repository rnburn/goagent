package internal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyLookupSuccess(t *testing.T) {
	tCases := []struct {
		contentType  string
		shouldRecord bool
	}{
		{"text/plain", false},
		{"application/json", true},
		{"application/x-www-form-urlencoded", true},
	}

	for _, tCase := range tCases {
		h := http.Header{}
		h.Add("Content-Type", tCase.contentType)
		assert.Equal(t, tCase.shouldRecord, ShouldRecordBodyOfContentType(h))
	}
}
