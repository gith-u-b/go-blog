package util

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/xerrors"
)

func EncodeMD5(value string) (string, error) {
	m := md5.New()
	_, err := m.Write([]byte(value))
	if err != nil {
		return "", xerrors.Errorf("EncodeMD5 failed: %v", err)
	}

	return hex.EncodeToString(m.Sum(nil)), nil
}
