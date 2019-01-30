package cmd

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	retCode := m.Run()

	os.Exit(retCode)
}
