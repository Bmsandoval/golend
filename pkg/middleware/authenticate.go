package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
 * Function: Valid(http.HandlerFunc, string)
 * Purpose: wrap an http.HandlerFunc as Middleware
 * Returns: 401 Unauthorized if invalid
 */
func ValidateSlackRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slackSigningSecret := os.Getenv("SLACK_SIGNING_SECRET")
		// split request to preserve body
		if IsRequestValid(r, slackSigningSecret) {
			next.ServeHTTP(w, r)
		} else {
			// Return failure for invalidated requests
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized Request"))
			return
		}
	})
}

/**
 * Function: IsRequestValid(*http.Request, string)
 * Purpose: Generates a Slack-style hash from the request
 * Returns: false if generated hash doesn't match hash in request
 */
func IsRequestValid(r *http.Request, slackSigningSecret string) bool {
	// *************************
	// Set variables
	// *************************
	var requestTimeout float64 = 300 // milliseconds
	slackVersion := "v0"
	// Get request data
	slackRequestTimestamp := r.Header.Get("X-Slack-Request-Timestamp")
	slackSignature := r.Header.Get("X-Slack-Signature")
	// *************************
	// Invalidate older messages
	// *************************
	// ParseFloat was acting weird with the decimal value
	splitTimestamp := strings.Split(slackRequestTimestamp, ".")[0]
	requestTimestamp, _ := strconv.ParseInt(splitTimestamp, 10, 64)
	// invalid if older than requestTimeout
	if math.Abs(float64(time.Now().Unix()-requestTimestamp)) > requestTimeout {
		log.Println("Old request received from IP " + r.RemoteAddr)
		return false
	}
	// *************************
	// Generate hash
	// *************************
	// Slack apparently only hashes the body
	requestData, _ := ioutil.ReadAll(r.Body)
	// Put the body back and share with others
	r.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))

	signatureBaseString := slackVersion + ":" + slackRequestTimestamp + ":" + string(requestData)
	// Generate a hash of the request, everything must be in bytes
	h := hmac.New(sha256.New, []byte(slackSigningSecret))
	h.Write([]byte(signatureBaseString))
	expectedSignature := slackVersion + "=" + hex.EncodeToString(h.Sum(nil))
	// *************************
	// Invalidate if no match
	// *************************
	if expectedSignature == slackSignature {
		return true
	} else {
		log.Println("Invalid hash received from IP " + r.RemoteAddr)
		return false
	}
}
