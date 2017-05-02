// Package env contains all env related data
package env

import "os"

// This constant is an identity of environment variable used in server
const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

// Get is used to get current environment from "LCENV" variables on the OS
func Get() string {
	env := os.Getenv("LCENV")
	if env == "" {
		env = EnvDevelopment
	}

	if env != EnvDevelopment && env != EnvProduction {
		env = EnvDevelopment
	}
	return env
}

// IsProduction is used to check whether current environment is on "production" or not
func IsProduction() bool {
	if Get() == EnvProduction {
		return true
	}
	return false
}

// IsDevelopment is used to check whether current environment is on "development" or not
func IsDevelopment() bool {
	if Get() == EnvDevelopment {
		return true
	}
	return false
}
