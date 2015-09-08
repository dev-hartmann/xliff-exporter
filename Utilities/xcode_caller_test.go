package utilities

import "testing"

func TestNewXcodeCaller(t *testing.T) {

	emptyCfg := &Config{}
	xcc := NewXcodeCaller(emptyCfg)

	if xcc == nil {
		t.Errorf("XcodeCaller should not be nil, but was")
	}
}

func TestXcodeCallerChangeToProjectDirNoError(t *testing.T) {
	emptyCfg := &Config{}
	xcc := NewXcodeCaller(emptyCfg)

	if _, err := xcc.ChangeToProjectDir("pwd"); err != nil {
		t.Error("An error occured during change dir:%v", err.Error())
	}
}

func TestXcodeCallerChangeToProjectDirError(t *testing.T) {
	emptyCfg := &Config{}
	xcc := NewXcodeCaller(emptyCfg)

	path, err := xcc.ChangeToProjectDir("test/test/test")

	if err == nil {
		t.Errorf("An error should occur when trying to change dir to:%v", path)
	}
}
