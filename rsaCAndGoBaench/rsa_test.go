package rsaCAndGoBaench

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	rsaC "github.com/dgkang/rsa/rsa"
	"io/ioutil"
	"testing"
)

/*

由于调用了C语言的库，需要在有C环境中才能正确执行；

比如Linux Unix 系统中可以不用在意C环境（自带C的所有执行库）并执行本测试。

在windows上不要配置C的环境，这里不展开说明，请自行百度 “window 上 golang 运行C语言”
*/

/*
	压测C语言rsa解密
*/
func BenchmarkRsaCDecrypt(b *testing.B) {
	str := "123456"
	for i := 0; i < b.N; i++ {
		encryptStr := RsaCEncrypt([]byte(str))
		_ = RsaCDecrypt(encryptStr)
	}
}

/*
	压测C语言rsa加密
*/
func BenchmarkRsaCEncrypt(b *testing.B) {
	str := "123456"
	for i := 0; i < b.N; i++ {
		_ = RsaCEncrypt([]byte(str))
	}
}

/*
	压测GO语言rsa解密
*/
func BenchmarkRsaGoDecrypt(b *testing.B) {
	str := "123456"
	for i := 0; i < b.N; i++ {
		encryptStr := RsaGoEncrypt([]byte(str))
		_ = RsaGoDecrypt(encryptStr)
	}
}

/*
	压测GO语言rsa加密
*/
func BenchmarkRsaGoEncrypt(b *testing.B) {
	str := "123456"
	for i := 0; i < b.N; i++ {
		_ = RsaGoEncrypt([]byte(str))
	}
}

//test C语言解密
func TestRsaCDecrypt(t *testing.T) {
	str := "123456"
	encryptStr := RsaCEncrypt([]byte(str))
	t.Log(encryptStr)
	decrypt := RsaCDecrypt(encryptStr)
	t.Log(string(decrypt))
}

//test C语言加密
func TestRsaCEncrypt(t *testing.T) {
	str := "123456"
	encryptStr := RsaCEncrypt([]byte(str))
	t.Log(encryptStr)
}

//test go语言解密
func TestRsaGoDecrypt(t *testing.T) {
	str := "123456"
	encryptStr := RsaGoEncrypt([]byte(str))
	t.Log(encryptStr)
	decrypt := RsaGoDecrypt(encryptStr)
	t.Log(string(decrypt))
}

//test go语言加密
func TestRsaGoEncrypt(t *testing.T) {
	str := "123456"
	encryptStr := RsaGoEncrypt([]byte(str))
	t.Log(encryptStr)
}

// C语言解密
func RsaCDecrypt(from []byte) []byte {
	decrypt, err := rsaC.PublicDecrypt(from, "./public.pem", rsaC.RSA_PKCS1_PADDING)
	if err != nil {
		panic(err)
	}
	return decrypt
}

// C语言加密
func RsaCEncrypt(from []byte) []byte {
	decrypt, err := rsaC.PrivateEncrypt(from, "./private.pem", rsaC.RSA_PKCS1_PADDING, "")
	if err != nil {
		panic(err)
	}
	return decrypt
}

var (
	publicKey, privateKey []byte
)

func init() {
	var err error
	publicKey, err = ioutil.ReadFile("./public.pem")
	if err != nil {
		panic(err)
	}
	privateKey, err = ioutil.ReadFile("./private.pem")
	if err != nil {
		panic(err)
	}
}

// Go语言解密
func RsaGoDecrypt(from []byte) []byte {
	decode, _ := pem.Decode(privateKey)
	key, err := x509.ParsePKCS1PrivateKey(decode.Bytes)
	if err != nil {
		panic(err)
	}
	bytes, err := rsa.DecryptPKCS1v15(rand.Reader, key, from)
	if err != nil {
		panic(err)
	}
	return bytes
}

// Go语言加密
func RsaGoEncrypt(from []byte) []byte {
	decode, _ := pem.Decode(publicKey)
	key, err := x509.ParsePKIXPublicKey(decode.Bytes)
	if err != nil {
		println(err)
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		println("get public key fail ...", ok)
	}
	bytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, from)
	if err != nil {
		println(err)
	}
	return bytes
}
