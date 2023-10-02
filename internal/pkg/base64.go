package pkg

import (
	"encoding/base64"
	"errors"
)

func Base64Encode(data string) (*string, error) {
	if len(data) == 0 {
		return nil, errors.New("empty string passed")
	}

	encodedText := base64.StdEncoding.EncodeToString([]byte(data))

	return &encodedText, nil

}

func Base64Decode(encodedData string) (*string, error) {
	if len(encodedData) == 0 {
		return nil, errors.New("empty string passed")
	}

	decodedText, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, err
	}
	str := string(decodedText)
	return &str, nil
}
