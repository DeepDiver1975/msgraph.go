// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// BrowserSyncSetting undocumented
type BrowserSyncSetting int

const (
	// BrowserSyncSettingVNotConfigured undocumented
	BrowserSyncSettingVNotConfigured BrowserSyncSetting = 0
	// BrowserSyncSettingVBlockedWithUserOverride undocumented
	BrowserSyncSettingVBlockedWithUserOverride BrowserSyncSetting = 1
	// BrowserSyncSettingVBlocked undocumented
	BrowserSyncSettingVBlocked BrowserSyncSetting = 2
)

// BrowserSyncSettingPNotConfigured returns a pointer to BrowserSyncSettingVNotConfigured
func BrowserSyncSettingPNotConfigured() *BrowserSyncSetting {
	v := BrowserSyncSettingVNotConfigured
	return &v
}

// BrowserSyncSettingPBlockedWithUserOverride returns a pointer to BrowserSyncSettingVBlockedWithUserOverride
func BrowserSyncSettingPBlockedWithUserOverride() *BrowserSyncSetting {
	v := BrowserSyncSettingVBlockedWithUserOverride
	return &v
}

// BrowserSyncSettingPBlocked returns a pointer to BrowserSyncSettingVBlocked
func BrowserSyncSettingPBlocked() *BrowserSyncSetting {
	v := BrowserSyncSettingVBlocked
	return &v
}
