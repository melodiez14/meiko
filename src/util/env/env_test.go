package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		env      string
		expected string
	}{
		{
			env:      "",
			expected: "development",
		},
		{
			env:      "development",
			expected: "development",
		},
		{
			env:      "production",
			expected: "production",
		},
		{
			env:      "staging",
			expected: "staging",
		},
		{
			env:      "uptoyou",
			expected: "development",
		},
	}
	for _, val := range cases {
		os.Setenv("LCENV", val.env)
		if got := Get(); got != val.expected {
			t.Errorf("Get() = %v, expected %v", got, val.expected)
		}
	}
}

func TestIsProduction(t *testing.T) {
	cases := []struct {
		env      string
		expected bool
	}{
		{
			env:      "",
			expected: false,
		},
		{
			env:      "development",
			expected: false,
		},
		{
			env:      "production",
			expected: true,
		},
		{
			env:      "uptoyou",
			expected: false,
		},
		{
			env:      "staging",
			expected: false,
		},
	}
	for _, val := range cases {
		os.Setenv("LCENV", val.env)
		if got := IsProduction(); got != val.expected {
			t.Errorf("Get() = %v, expected %v", got, val.expected)
		}
	}
}

func TestIsDevelopment(t *testing.T) {
	cases := []struct {
		env      string
		expected bool
	}{
		{
			env:      "",
			expected: true,
		},
		{
			env:      "development",
			expected: true,
		},
		{
			env:      "production",
			expected: false,
		},
		{
			env:      "uptoyou",
			expected: true,
		},
		{
			env:      "staging",
			expected: false,
		},
	}
	for _, val := range cases {
		os.Setenv("LCENV", val.env)
		if got := IsDevelopment(); got != val.expected {
			t.Errorf("Get() = %v, expected %v", got, val.expected)
		}
	}
}

func TestIsStaging(t *testing.T) {
	cases := []struct {
		env      string
		expected bool
	}{
		{
			env:      "",
			expected: false,
		},
		{
			env:      "development",
			expected: false,
		},
		{
			env:      "production",
			expected: false,
		},
		{
			env:      "uptoyou",
			expected: false,
		},
		{
			env:      "staging",
			expected: true,
		},
	}
	for _, val := range cases {
		os.Setenv("LCENV", val.env)
		if got := IsStaging(); got != val.expected {
			t.Errorf("Get() = %v, expected %v", got, val.expected)
		}
	}
}
