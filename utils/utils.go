package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"io/ioutil"
	rand2 "math/rand"
	"net"
	"net/mail"
	"net/smtp"
	"os"
	"time"
)

type DbConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Sslmode  string `json:"sslmode"`
}

type Sender struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type hmac struct {
	Secret string `json:"hmac-secret"`
}

type jwtSecret struct {
	Secret string `json:"jwt-secret"`
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

func SendEmail(email string, data string) error {
	path, err := os.Getwd()

	if err != nil {
		return err
	}

	file, err := os.Open(path + "\\keys.json")

	if err != nil {
		return err
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	var sender Sender

	err = json.Unmarshal(byteValue, &sender)

	if err != nil {
		return err
	}

	from := mail.Address{"", sender.Email}
	to := mail.Address{"", email}
	subj := "Registration"
	body := "To complete registration go to \n http://localhost:8080/register?query=mutation+_{completeRegister(hash:\"" +
		data + "\"){successful}}"

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "smtp.gmail.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", sender.Email, sender.Password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil

}

func ParseHMACSecret(path string) string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	hmac := new(hmac)

	err = json.Unmarshal(byteValue, hmac)

	if err != nil {
		panic(err)
	}

	return hmac.Secret

}

func ParseJwtSecret(path string) []byte {
	file, _ := os.Open(path)

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	var jwtS jwtSecret

	err = json.Unmarshal(byteValue, &jwtS)

	if err != nil {
		panic(err)
	}

	return []byte(jwtS.Secret)
}

func CreateToken(secret []byte, email string, userType string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["type"] = userType

	token.Claims = claims

	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func ValidateToken(tokenHeader string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		path, err := os.Getwd()

		if err != nil {
			return nil, err
		}

		return ParseJwtSecret(path + "\\keys.json"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("authentication failed")
	}

	return claims, nil
}
