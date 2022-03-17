package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	sha12 "crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

//Sha1 加密
func Sha1(data string) string {
	sha1 := sha12.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

//MD5加密
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//DecryptAES256GCM
// aesKey 密钥
// associatedData 附加数据
//ciphertext 数据密文
//nonce 加密使用的随机串
func DecryptAES256GCM(aesKey, associatedData, nonce, ciphertext string, dest interface{}) error {

	gcm, err := decryptAES256GCM(aesKey, associatedData, nonce, ciphertext)
	if err != nil {
		return err
	}

	err = json.Unmarshal(gcm, &dest)
	return err
}

func decryptAES256GCM(aesKey, associatedData, nonce, ciphertext string) (plaintext []byte, err error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	c, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	dataBytes, err := gcm.Open(nil, []byte(nonce), decodedCiphertext, []byte(associatedData))
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}
