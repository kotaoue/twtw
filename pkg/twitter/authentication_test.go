package twitter_test

import (
	"testing"

	"github.com/kotaoue/twtw/pkg/twitter"
	"github.com/stretchr/testify/assert"
)

func Test_AuthNonce(t *testing.T) {
	obj := twitter.NewAuthentication()
	got := obj.AuthNonce()
	assert.NotEmpty(t, got)
}
