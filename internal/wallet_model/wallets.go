package wallet_model

import (
	st "blockchain/pkg/setting"
	"bytes"
	"crypto/elliptic"
	"encoding/gob" 
	"log/slog"
	"os"
)

type Wallets struct {
	Wallets map[string]*Wallet // 采用地址字符串代表钱包
}

var walletFile = st.WalletFile

func NewWallets(nodeID string) (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	err := wallets.LoadFromFile(nodeID)
	return &wallets, err
}
 
func(ws *Wallets) CreateWallet()string {
	wallet := NewWallet()
	address := string(wallet.GetAddress())
	ws.Wallets[address] = wallet
	slog.Info("New address generated, adding into Wallets...")
	return address
}

func(ws *Wallets) GetAddress()[]string {
	// TODO: 找到所有键值：全体地址
	var addresses []string

	for address:= range ws.Wallets {
		addresses = append(addresses, address)
	}
	return addresses
}


func(ws *Wallets) GetWallet(address string)Wallet{
	// TODO: 根据钱包地址获得钱包
	return *ws.Wallets[address]
}
func(ws *Wallets) LoadFromFile(nodeID string) error {
	// TODO: 从文件中导入钱包 
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}
	fileContent, err := os.ReadFile(walletFile)
	if err != nil {
		slog.Error(err.Error())
	}
	ws.Wallets = Deserialize(fileContent).Wallets
	slog.Info("Wallets successfully loaded from file")
	return nil
}

func(ws *Wallets) SaveToFile() {
	err := os.WriteFile(walletFile, ws.serialize(), 0644)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("New Wallet has been save into file")
}

func(ws *Wallets) serialize() []byte {
	// TODO: 序列化钱包集
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		slog.Error(err.Error())
	}
	return content.Bytes()
}
func Deserialize(seq []byte) Wallets{
	// TODO：反序列化钱包集
	var wallets Wallets
	// gob.Register(elliptic.P256())
	// gob.Register("p256")
	decoder := gob.NewDecoder(bytes.NewReader(seq))
	err := decoder.Decode(&wallets)
	if err != nil {
		panic(err)
	}
	return wallets
}