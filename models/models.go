package models

import "github.com/satori/go.uuid"

func createObfuscatedID() string {
	obfuscatedID := uuid.NewV4()
	return obfuscatedID.String()
}
