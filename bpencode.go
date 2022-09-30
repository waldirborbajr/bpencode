package bpencode

import (
	"errors"
	"net/url"
)

func init() {
}

func BPEncode(pwd string) (string, error) {
	if len(pwd) == 0 {
		return "", errors.New("empty: empty content")
	}

	pwd, err := url.QueryUnescape(pwd)
	if err != nil {
		return "", err
	}

	return pwd, nil
}
