//go:build linux && !android

package sweetcookie

import (
	"path/filepath"
	"testing"
)

func TestFirefoxRoots_CommunityLinux(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(home, "xdg"))

	cases := map[Browser][]string{
		BrowserFirefox: {
			filepath.Join(home, "snap", "firefox", "common", ".mozilla", "firefox"),
			filepath.Join(home, ".var", "app", "org.mozilla.firefox", ".mozilla", "firefox"),
			filepath.Join(home, "xdg", "mozilla", "firefox"),
			filepath.Join(home, ".mozilla", "firefox"),
		},
		BrowserZen: {
			filepath.Join(home, ".zen"),
			filepath.Join(home, ".var", "app", "app.zen_browser.zen", ".zen"),
		},
		BrowserFloorp: {
			filepath.Join(home, ".mozilla", "floorp"),
			filepath.Join(home, ".floorp"),
			filepath.Join(home, ".var", "app", "one.ablaze.floorp", ".floorp"),
		},
		BrowserWaterfox: {
			filepath.Join(home, ".waterfox"),
			filepath.Join(home, ".var", "app", "net.waterfox.waterfox", ".waterfox"),
		},
		BrowserLibreWolf: {
			filepath.Join(home, ".librewolf"),
			filepath.Join(home, ".var", "app", "io.gitlab.librewolf-community", ".librewolf"),
		},
	}
	for browser, want := range cases {
		got := firefoxRoots(browser)
		if len(got) != len(want) {
			t.Fatalf("%s: want %v got %v", browser, want, got)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("%s: want %v got %v", browser, want, got)
			}
		}
	}
}
