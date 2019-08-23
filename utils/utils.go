package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	rand2 "math/rand"
	"os"
	"time"
)

type DbConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Sslmode  string `json:"sslmode"`
}

func ParseDbConfig(path string) (*DbConfig, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	bytevalue, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var dbConfig DbConfig

	err = json.Unmarshal(bytevalue, &dbConfig)

	if err != nil {
		return nil, err
	}

	return &dbConfig, nil
}

func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func createHash(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))

	return hex.EncodeToString(hash.Sum(nil))
}

func GetRandomString(length int) string {
	rand2.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-=/.,:;")
	var byteStr []byte
	for i := 0; i < length; i++ {
		randInd := rand2.Intn(len(letters))
		byteStr = append(byteStr, byte(letters[randInd]))
	}

	return string(byteStr)
}
