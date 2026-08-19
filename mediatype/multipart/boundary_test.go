package multipart

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateBoundary(t *testing.T) {
	const testCount = 1000
	// Valid characters are: DIGIT / ALPHA / "'" / "(" / ")" / "+" / "_" / "," / "-" / "." / "/" / ":" / "=" / "?"
	const validChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'()+_,-./:=?"

	for range testCount {
		boundary, err := generateBoundary()
		require.NoError(t, err)
		boundaryLength := len(boundary)
		require.GreaterOrEqual(t, boundaryLength, MIN_BOUNDARY_LENGTH)
		require.LessOrEqual(t, boundaryLength, MAX_BOUNDARY_LENGTH)
		for _, char := range boundary {
			require.Contains(t, validChars, string(char))
		}
		fmt.Printf("boundary: %s\n", boundary)
	}
}
