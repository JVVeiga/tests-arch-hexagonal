package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandlerJsonError(t *testing.T) {
	msg := "Test Json Error"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"Test Json Error"}`), result)
}
