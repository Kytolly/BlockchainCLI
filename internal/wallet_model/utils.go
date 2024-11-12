package wallet_model

import (
	st "blockchain/pkg/setting" 
	utils "blockchain/pkg/utils"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log/slog"
	"math/big"
)

func newKeyPair() (ecdsa.PrivateKey, []byte){
	//TODO: 利用椭圆曲线生成私钥-公钥对
	//椭圆曲线的特性决定了这样了这样的对极不可能重复
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err!= nil{
		slog.Error("Wallet Inition Failed: cannot generate key")
	}
	// 公钥是一对坐标的压缩，由私钥生成，无法反推私钥
	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, public
}

func checksum(payload []byte) []byte {
	//TODO： 哈希处理后计算校验和
	addressChecksumLen := st.ChecksumLen
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:addressChecksumLen]
}
func CheckAddress(address string)bool {
	pubKeyHash := utils.Base58Decode([]byte(address))

	ChecksumLen := st.ChecksumLen
	// VersionLen := len([]byte(st.AlgorithmVersion) )
	VersionLen := 1

	actualChecksum := pubKeyHash[len(pubKeyHash)-ChecksumLen:]
	version := pubKeyHash[:VersionLen]
	pubKeyHash = pubKeyHash[VersionLen : len(pubKeyHash)-ChecksumLen]
	targetChecksum := checksum(append(version, pubKeyHash...))

	return bytes.Equal(actualChecksum, targetChecksum) 
}

func(w *Wallet) GobEncode() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	curveName := "P256"
	if err := encoder.Encode(curveName); err!= nil {
		return nil, err
	}
	privateKeyBytes := w.PrivateKey.D.Bytes()
	if err := encoder.Encode(privateKeyBytes); err != nil {
		return nil, err
	} 
	xBytes := w.PrivateKey.PublicKey.X.Bytes()
	yBytes := w.PrivateKey.PublicKey.Y.Bytes()
	if err := encoder.Encode(xBytes); err != nil {
		return nil, err
	}
	if err := encoder.Encode(yBytes); err != nil {
		return nil, err
	} 
	if err := encoder.Encode(w.PublicKey); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (w *Wallet) GobDecode(data []byte) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer) 
	var curveName string
	if err := decoder.Decode(&curveName); err != nil {
		return err
	} 
	var curve elliptic.Curve
	switch curveName {
	case "P256":
		curve = elliptic.P256()
	default:
		return fmt.Errorf("unsupported curve: %s", curveName)
	} 
	var privateKeyBytes []byte
	if err := decoder.Decode(&privateKeyBytes); err != nil {
		return err
	}

	w.PrivateKey.PublicKey.Curve = curve
	w.PrivateKey.D = new(big.Int).SetBytes(privateKeyBytes) 
	var xBytes, yBytes []byte
	if err := decoder.Decode(&xBytes); err != nil {
		return err
	}
	if err := decoder.Decode(&yBytes); err != nil {
		return err
	}

	w.PrivateKey.PublicKey.X = new(big.Int).SetBytes(xBytes)
	w.PrivateKey.PublicKey.Y = new(big.Int).SetBytes(yBytes) 
	if err := decoder.Decode(&w.PublicKey); err != nil {
		return err
	}

	return nil
}