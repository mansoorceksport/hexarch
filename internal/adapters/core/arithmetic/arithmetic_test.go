package arithmetic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdapter_Addition(t *testing.T) {
	arith := NewAdapter()
	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got %v", nil, err)
	}
	require.Equal(t, answer, int32(2))
}

func TestAdapter_Subtraction(t *testing.T) {
	arith := NewAdapter()
	answer, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got %v", nil, err)
	}
	require.Equal(t, answer, int32(0))
}

func TestAdapter_Multiplication(t *testing.T) {
	arith := NewAdapter()
	answer, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got %v", nil, err)
	}
	require.Equal(t, answer, int32(1))
}

func TestAdapter_Division(t *testing.T) {
	arith := NewAdapter()
	answer, err := arith.Division(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got %v", nil, err)
	}
	require.Equal(t, answer, int32(1))
}
