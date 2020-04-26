package test

import (
	"testing"

	"github.com/solrac97gr/twitter-fake/database"
)

func testDbconection(t *testing.T) {
	actual := true
	isConected := database.ConnectionOK()
	if actual != isConected {
		t.Fail()
	}
}
