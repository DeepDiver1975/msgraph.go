// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FileHashType undocumented
type FileHashType int

const (
	// FileHashTypeVUnknown undocumented
	FileHashTypeVUnknown FileHashType = 0
	// FileHashTypeVSha1 undocumented
	FileHashTypeVSha1 FileHashType = 1
	// FileHashTypeVSha256 undocumented
	FileHashTypeVSha256 FileHashType = 2
	// FileHashTypeVMd5 undocumented
	FileHashTypeVMd5 FileHashType = 3
	// FileHashTypeVAuthenticodeHash256 undocumented
	FileHashTypeVAuthenticodeHash256 FileHashType = 4
	// FileHashTypeVLsHash undocumented
	FileHashTypeVLsHash FileHashType = 5
	// FileHashTypeVCtph undocumented
	FileHashTypeVCtph FileHashType = 6
	// FileHashTypeVUnknownFutureValue undocumented
	FileHashTypeVUnknownFutureValue FileHashType = 127
)

// FileHashTypePUnknown returns a pointer to FileHashTypeVUnknown
func FileHashTypePUnknown() *FileHashType {
	v := FileHashTypeVUnknown
	return &v
}

// FileHashTypePSha1 returns a pointer to FileHashTypeVSha1
func FileHashTypePSha1() *FileHashType {
	v := FileHashTypeVSha1
	return &v
}

// FileHashTypePSha256 returns a pointer to FileHashTypeVSha256
func FileHashTypePSha256() *FileHashType {
	v := FileHashTypeVSha256
	return &v
}

// FileHashTypePMd5 returns a pointer to FileHashTypeVMd5
func FileHashTypePMd5() *FileHashType {
	v := FileHashTypeVMd5
	return &v
}

// FileHashTypePAuthenticodeHash256 returns a pointer to FileHashTypeVAuthenticodeHash256
func FileHashTypePAuthenticodeHash256() *FileHashType {
	v := FileHashTypeVAuthenticodeHash256
	return &v
}

// FileHashTypePLsHash returns a pointer to FileHashTypeVLsHash
func FileHashTypePLsHash() *FileHashType {
	v := FileHashTypeVLsHash
	return &v
}

// FileHashTypePCtph returns a pointer to FileHashTypeVCtph
func FileHashTypePCtph() *FileHashType {
	v := FileHashTypeVCtph
	return &v
}

// FileHashTypePUnknownFutureValue returns a pointer to FileHashTypeVUnknownFutureValue
func FileHashTypePUnknownFutureValue() *FileHashType {
	v := FileHashTypeVUnknownFutureValue
	return &v
}
