package decode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Bird struct {
	Wings int
	Tail  bool
}

type Fish struct {
	Scales int
	Teeth  bool
}

type Animal struct {
	Bird *Bird
	Fish *Fish
}

func TestOneOf(t *testing.T) {
	var data = []byte(`{"Wings":2,"Tail":true}`)

	animal := Animal{}

	err := OneOf(data, &animal)
	if err != nil {
		t.Fatalf("Failed to marshal animal: %v", err)
	}

	require.Equal(t, &Bird{Wings: 2, Tail: true}, animal.Bird)
	require.Nil(t, animal.Fish)

	animal = Animal{}
	data = []byte(`{"Scales":2,"Teeth":true}`)
	err = OneOf(data, &animal)
	if err != nil {
		t.Fatalf("Failed to marshal animal: %v", err)
	}

	require.Equal(t, &Fish{Scales: 2, Teeth: true}, animal.Fish)
}
