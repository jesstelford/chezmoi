package chezmoi

import (
	"bytes"
	"os"
	"runtime"
)

// An EntryStateType is an entry state type.
type EntryStateType string

// Entry state types.
const (
	EntryStateTypeDir     EntryStateType = "dir"
	EntryStateTypeFile    EntryStateType = "file"
	EntryStateTypeSymlink EntryStateType = "symlink"
	EntryStateTypeRemove  EntryStateType = "remove"
	EntryStateTypeScript  EntryStateType = "script"
)

// An EntryState represents the state of an entry. A nil EntryState is
// equivalent to EntryStateTypeAbsent.
type EntryState struct {
	Type           EntryStateType `json:"type" toml:"type" yaml:"type"`
	Mode           os.FileMode    `json:"mode,omitempty" toml:"mode,omitempty" yaml:"mode,omitempty"`
	ContentsSHA256 HexBytes       `json:"contentsSHA256,omitempty" toml:"contentsSHA256,omitempty" yaml:"contentsSHA256,omitempty"`
	contents       []byte
}

// Contents returns s's contents, if available.
func (s *EntryState) Contents() []byte {
	return s.contents
}

// Equal returns true if s is equal to other.
func (s *EntryState) Equal(other *EntryState) bool {
	if s.Type != other.Type {
		return false
	}
	if runtime.GOOS != "windows" && s.Mode.Perm() != other.Mode.Perm() {
		return false
	}
	return bytes.Equal(s.ContentsSHA256, other.ContentsSHA256)
}

// Equivalent returns true if s is equivalent to other.
func (s *EntryState) Equivalent(other *EntryState) bool {
	switch {
	case s == nil:
		return other == nil || other.Type == EntryStateTypeRemove
	case other == nil:
		return s.Type == EntryStateTypeRemove
	default:
		return s.Equal(other)
	}
}
