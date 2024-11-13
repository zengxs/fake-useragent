package fakeuseragent

import "slices"

type UserAgent struct {
	Value string   `yaml:"ua"`
	Tags  []string `yaml:"tags"`
}

func (ua UserAgent) String() string {
	return ua.Value
}

func (ua UserAgent) HasTag(tag UserAgentTag) bool {
	return slices.Contains(ua.Tags, string(tag))
}

type UserAgentTag string

var (
	// platforms
	TagDesktop UserAgentTag = "desktop"
	TagMobile  UserAgentTag = "mobile"
	TagTablet  UserAgentTag = "tablet"
	// browsers
	TagChrome  UserAgentTag = "chrome"
	TagChroium UserAgentTag = "chromium"
	TagSafari  UserAgentTag = "safari"
	TagFirefox UserAgentTag = "firefox"
	TagMSEdge  UserAgentTag = "msedge"
	TagMSIE    UserAgentTag = "msie"
	TagBrave   UserAgentTag = "brave"
	TagYandex  UserAgentTag = "yandex"
	// operating systems
	TagWindows  UserAgentTag = "windows"
	TagMacOS    UserAgentTag = "macos"
	TagLinux    UserAgentTag = "linux"
	TagAndroid  UserAgentTag = "android"
	TagIOS      UserAgentTag = "ios"
	TagChromeOS UserAgentTag = "chromeos"
)
