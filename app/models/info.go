package models

const (
	// Major version
	Major = "0"
	// Minor version
	Minor = "0"
	// Patch version
	Patch = "1"
)

// Info - application version information
type Info struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
	Patch string `json:"patch"`
}
