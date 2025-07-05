package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)



func CalculateFileHash(filePath string) (string, error) {
	f,err:= os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func CalculateNodeHash(nodeName string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(nodeName))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
