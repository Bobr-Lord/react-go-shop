package hash

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	pass := "password"
	hash, err := HashPassword(pass)
	if err != nil || hash == "" {
		t.Errorf("HashPassword() error = %v", err)
		return
	}
	t.Logf("HashPassword() hash = %v", hash)

}
