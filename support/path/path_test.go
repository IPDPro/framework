package path

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	tests := map[string]struct {
		a        []string
		expected string
	}{
		"no args": {
			a:        []string{},
			expected: "app",
		},
		"single arg": {
			a:        []string{"config/goravel.go"},
			expected: filepath.Join("app", "config", "goravel.go"),
		},
		"multi arg": {
			a:        []string{"config/goravel.go", "database/migrations"},
			expected: filepath.Join("app", "config", "goravel.go", "database", "migrations"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := App(test.a...)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestBase(t *testing.T) {
	tests := map[string]struct {
		a        []string
		expected string
	}{
		"no args": {
			a:        []string{},
			expected: "",
		},
		"single arg": {
			a:        []string{"config/goravel.go"},
			expected: filepath.Join("config", "goravel.go"),
		},
		"multi arg": {
			a:        []string{"config/goravel.go", "database/migrations"},
			expected: filepath.Join("config", "goravel.go", "database", "migrations"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Base(test.a...)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestConfig(t *testing.T) {
	tests := map[string]struct {
		a        []string
		expected string
	}{
		"no args": {
			a:        []string{},
			expected: "config",
		},
		"single arg": {
			a:        []string{"goravel.go"},
			expected: filepath.Join("config", "goravel.go"),
		},
		"multi arg": {
			a:        []string{"goravel.go", "database/migrations"},
			expected: filepath.Join("config", "goravel.go", "database", "migrations"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Config(test.a...)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestDatabase(t *testing.T) {
	tests := map[string]struct {
		a        []string
		expected string
	}{
		"no args": {
			a:        []string{},
			expected: "database",
		},
		"single arg": {
			a:        []string{"migrations"},
			expected: filepath.Join("database", "migrations"),
		},
		"multi arg": {
			a:        []string{"migrations", ".gitignore"},
			expected: filepath.Join("database", "migrations", ".gitignore"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Database(test.a...)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestStorage(t *testing.T) {
	tests := map[string]struct {
		a        []string
		expected string
	}{
		"no args": {
			a:        []string{},
			expected: "storage",
		},
		"single arg": {
			a:        []string{"test"},
			expected: filepath.Join("storage", "test"),
		},
		"multi arg": {
			a:        []string{"test", ".gitignore"},
			expected: filepath.Join("storage", "test", ".gitignore"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Storage(test.a...)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestLang(t *testing.T) {
	tests := map[string]struct {
		a        []string
		expected string
	}{
		"no args": {
			a:        []string{},
			expected: "lang",
		},
		"single arg": {
			a:        []string{"test"},
			expected: filepath.Join("lang", "test"),
		},
		"multi arg": {
			a:        []string{"test", ".gitignore"},
			expected: filepath.Join("lang", "test", ".gitignore"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Lang(test.a...)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestPublic(t *testing.T) {
	tests := map[string]struct {
		a        []string
		expected string
	}{
		"no args": {
			a:        []string{},
			expected: "public",
		},
		"single arg": {
			a:        []string{"test"},
			expected: filepath.Join("public", "test"),
		},
		"multi arg": {
			a:        []string{"test", ".gitignore"},
			expected: filepath.Join("public", "test", ".gitignore"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Public(test.a...)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestExecutable(t *testing.T) {
	path, err := os.Getwd()
	assert.NoError(t, err)

	tests := map[string]struct {
		a        []string
		expected string
	}{
		"no args": {
			a:        []string{},
			expected: path,
		},
		"single arg": {
			a:        []string{"test"},
			expected: filepath.Join(path, "test"),
		},
		"multi arg": {
			a:        []string{"test", ".gitignore"},
			expected: filepath.Join(path, "test", ".gitignore"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Executable(test.a...)

			assert.Equal(t, test.expected, actual)
		})
	}
}
