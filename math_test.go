package testrepo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRand(t *testing.T) {
	for i := 0; i < 100; i++ {
		require.NotEqual(t, Rand(), Rand())
	}
}
