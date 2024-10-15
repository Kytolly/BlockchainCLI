package wallet_model

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log/slog"
	st "blockchain/pkg/setting"
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