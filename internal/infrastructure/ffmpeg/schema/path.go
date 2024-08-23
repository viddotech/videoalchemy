package schema

import (
	"fmt"
	"net/url"
	"os"
)

type FFMPEGSource string

// IsOSPath checks if the FFMPEGSource is a valid file system path
func (p FFMPEGSource) IsOSPath() bool {
	if _, err := os.Stat(string(p)); err == nil {
		return true
	}
	return false
}

// IsHTTPPath checks if the FFMPEGSource is a valid URL
func (p FFMPEGSource) IsHTTPPath() bool {
	_, err := url.ParseRequestURI(string(p))
	return err == nil
}

// Validate checks if the FFMPEGSource is either a valid OS path or a valid HTTP path
func (p FFMPEGSource) Validate() error {
	if p.IsOSPath() {
		return nil
	}
	if p.IsHTTPPath() {
		return nil
	}
	return fmt.Errorf("invalid path: %s", p)
}
