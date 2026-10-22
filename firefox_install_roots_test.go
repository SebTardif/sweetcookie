package sweetcookie

import (
	"path/filepath"
	"testing"
)

func TestFirefoxInstallRootsIncludesSnapFlatpakAndXDG(t *testing.T) {
	home := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(home, "xdg"))
	got := firefoxInstallRoots(home)
	want := []string{
		filepath.Join(home, "snap", "firefox", "common", ".mozilla", "firefox"),
		filepath.Join(home, ".var", "app", "org.mozilla.firefox", ".mozilla", "firefox"),
		filepath.Join(home, "xdg", "mozilla", "firefox"),
		filepath.Join(home, ".mozilla", "firefox"),
	}
	if len(got) != len(want) {
		t.Fatalf("got %v want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v want %v", got, want)
		}
	}
}
