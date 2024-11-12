package utils

import ( 
	"crypto/sha256"
	"log/slog" 
	st "blockchain/pkg/setting"
	"golang.org/x/crypto/ripemd160"
)

func HashPubKey(pubKey []byte) []byte {
	// TODO: 利用两次hash算法对公钥进行hash
	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		slog.Error("failed in hashing pubkey!")
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}

func GetPubKeyHashInAddress(address []byte) []byte {
	// TODO: 从地址中提取公钥哈希值
	// 完整的未编码压缩的地址包括版本号，公钥哈希值，校验和,需要去掉版本号和若干位校验和
	pubKeyHash := Base58Decode(address)
	ChecksumLen := st.ChecksumLen
	// VersionLen := len([]byte(st.Version) )
	VersionLen := 1
	return pubKeyHash[VersionLen:len(pubKeyHash)-ChecksumLen]
}
