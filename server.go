/**
 * This is a self-contained example application
 * which shows the basic components behind pnthr
 * and allows you to create your own secure transport layer
 * using this code as a boilerplate.
 */

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/**
	 * GET or POST /server
	 *
	 * All requests will come through the /server route
	 * Each request should have a public key header and payload body
	 */
	http.HandleFunc("/server", root)

	log.Println("Listening for connections...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	/**
	 * Public Key
	 *
	 * If we don't have an api key in the header, then we can't fulfill this request
	 *
	 * Normally we would lookup and retrieve our private key from a database
	 * using the public key to find the record, but for now, we'll just make certain
	 * that a public key is passed and use a fixed private key and initialization vector
	 */
	publicKey := r.Header.Get("publicKey")
	privateKey := []byte("53d07a9c3566360002020001")
	initializationVector := []byte("53a1c59f62393700")[:aes.BlockSize]

	if len(publicKey) == 0 {
		Responder(w, r, 412, "The 'publicKey' request header was not found")
		return
	}

	/**
	 * Received Payload
	 *
	 * Ensure that we received a payload and that
	 * the payload is encrypted with the private key
	 */
	payload, err := ioutil.ReadAll(r.Body)

	if err != nil || string(payload) == "" {
		Responder(w, r, 422, "The request body was empty or unprocessable")
		return
	}

	fmt.Println("\n\nPayload: ", string(payload))

	/**
	 * Encrypt the insecure payload
	 *
	 * We allocate the bytes array that represent both the encrypted and decrypted data
	 */
	encrypted := make([]byte, len(payload))
	err = encrypt(encrypted, payload, privateKey, initializationVector)

	if err != nil {
		panic(err)
	}

	fmt.Println("\nEncrypted Payload: ", encode(encrypted))

	/**
	 * Decrypt the encrypted payload
	 */
	decrypted := make([]byte, len(payload))
	err = decrypt(decrypted, encrypted, privateKey, initializationVector)

	if err != nil {
		panic(err)
	}

	fmt.Println("\nDecrypted Payload: ", string(decrypted))

	// Return the encrypted data to the sender
	fmt.Fprintf(w, encode(encrypted))
}

/**
 * HTTP Error Handler
 *
 * Set the http status code and provide an error message
 */
func Responder(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	log.Println(message)
	fmt.Fprintf(w, message)
}

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func encrypt(dst, src, key, iv []byte) error {
	aesBlockEncryptor, err := aes.NewCipher([]byte(key))

	if err != nil {
		return err
	}

	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
	aesEncrypter.XORKeyStream(dst, src)

	return nil
}

func decrypt(dst, src, key, iv []byte) error {
	aesBlockEncryptor, err := aes.NewCipher([]byte(key))

	if err != nil {
		return err
	}

	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
	aesEncrypter.XORKeyStream(dst, src)

	return nil
}
