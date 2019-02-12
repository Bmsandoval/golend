package registration

import (
	"crypto/rand"
	"lendr/internal/models/lender"
	"log"
	"net/http"
	"os"
)



func EnrollmentUrl(w http.ResponseWriter, r *http.Request) {
	slackClientID := os.Getenv("SLACK_CLIENT_ID")
	slackState := lender.CreateProspect(SecureRandomAlphaString)
	url := "https://slack.com/oauth/authorize?" +
		"client_id=" + slackClientID + "&" +
		"scope=chat:write:bot bot commands" + "&" +
		"state=" + slackState
	http.Redirect(w, r, url, 302)
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52 possibilities
	letterIdxBits = 6                                                      // 6 bits to represent 64 possibilities / indexes
	letterIdxMask = 1<<letterIdxBits - 1                                   // All 1-bits, as many as letterIdxBits
)

func SecureRandomAlphaString(length int) string {

	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}
