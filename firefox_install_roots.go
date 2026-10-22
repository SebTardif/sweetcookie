package sweetcookie

import (
	"os"
	"path/filepath"
)

func firefoxInstallRoots(home string) []string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = filepath.Join(home, ".config")
	}
	return []string{
		filepath.Join(home, "snap", "firefox", "common", ".mozilla", "firefox"),
		filepath.Join(home, ".var", "app", "org.mozilla.firefox", ".mozilla", "firefox"),
		filepath.Join(configHome, "mozilla", "firefox"),
		filepath.Join(home, ".mozilla", "firefox"),
	}
}
