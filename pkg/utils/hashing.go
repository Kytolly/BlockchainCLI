package utils

import ( 
	"crypto/sha256"
	"log/slog" 
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
