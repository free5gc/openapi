package encode

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
	animal := Animal{
		Bird: &Bird{Wings: 2, Tail: true},
	}

	data, err := OneOf(&animal)
	if err != nil {
		t.Fatalf("Failed to marshal animal: %v", err)
	}

	require.Equal(t, []byte(`{"Wings":2,"Tail":true}`), data)

	animal = Animal{
		Fish: &Fish{Scales: 2, Teeth: true},
	}
	data, err = OneOf(&animal)
	if err != nil {
		t.Fatalf("Failed to marshal animal: %v", err)
	}

	require.Equal(t, []byte(`{"Scales":2,"Teeth":true}`), data)
}
