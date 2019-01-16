package cmd

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	before()
	retCode := m.Run()

	os.Exit(retCode)
}
