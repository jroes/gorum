package models

import (
	"testing"
)

import (
	"github.com/jroes/gorum/models"
)

func TestUserHasPassword(t *testing.T) {
	user := models.NewUser("test@example.com", "password")

	if err := user.HasPassword("password"); nil != err {
		t.Errorf("Failed to recognize valid password for user, %v.\n", err)
	}
}
