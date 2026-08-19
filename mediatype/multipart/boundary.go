package multipart

import (
	"crypto/rand"
)

// RFC 1341 multipart boundary is formally specified by the following BNF:
// boundary := 0*69<bchars> bcharsnospace
// bchars := bcharsnospace / " "
// bcharsnospace := DIGIT / ALPHA / "'" / "(" / ")" / "+" / "_" /
//                  "," / "-" / "." / "/" / ":" / "=" / "?"
// NOTE: The operator "*" preceding an element indicates repetition in ABNF.
//       The form 0*69<element> means that the element can be repeat by 0 time to 69 times.
// NOTE2: This regexp is use "Capturing Group" to extact the boundary string in Content-Type Header

const (
	// 0*69<bchars> bcharsnospace = max 69 chars
	MAX_BOUNDARY_LENGTH = 69
	// custom min boundary length
	MIN_BOUNDARY_LENGTH = 30
)

func generateBoundary() (string, error) {
	// Generate a boundary string that complies with RFC 1341
	// Valid characters are: DIGIT / ALPHA / "'" / "(" / ")" / "+" / "_" / "," / "-" / "." / "/" / ":" / "=" / "?"
	const validChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'()+_,-./:=?"
	var err error

	// Create a random boundary with length between 30 and 69 characters
	var lengthBytes [4]byte
	_, err = rand.Read(lengthBytes[:])
	if err != nil {
		return "", err
	}
	length := MIN_BOUNDARY_LENGTH + int(lengthBytes[0])%(MAX_BOUNDARY_LENGTH-MIN_BOUNDARY_LENGTH)

	boundary := make([]byte, length)
	_, err = rand.Read(boundary)
	if err != nil {
		return "", err
	}
	for i := range boundary {
		boundary[i] = validChars[int(boundary[i])%len(validChars)]
	}

	return string(boundary), nil
}
