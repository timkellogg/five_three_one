package models

import "testing"

func TestInfo(t *testing.T) {
	if Major != "0" {
		t.Errorf("Major version should be 0 but is %s", Major)
	}

	if Minor != "0" {
		t.Errorf("Minor version should be 0 but is %s", Minor)
	}

	if Patch != "1" {
		t.Errorf("Patch version should be 1 but is %s", Patch)
	}
}
