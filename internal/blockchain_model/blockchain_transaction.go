package blockchain_model

import (
	ts "blockchain/internal/transaction_model"
	wt "blockchain/internal/wallet_model"
	"blockchain/pkg/utils"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log/slog"
)

func(bc *BlockChain) NewUTXOTransaction(from, to string, amount int, u *UTXOSet)*ts.Transaction {
	//TODO:创建一个通用的交易
	var inputs []ts.TXInput
	var outputs []ts.TXOutput

	wallets, err := wt.NewWallets()
	if err != nil {
		slog.Error("Cannot generate a wallet!")
	}
	wallet_from := wallets.GetWallet(from)
	wallet_to := wallets.GetWallet(to)
	fromPubKeyHash := utils.HashPubKey(wallet_from.PublicKey)
	toPubKeyHash := utils.HashPubKey(wallet_to.PublicKey)
	slog.Debug("from PubKeyHash:", "", string(fromPubKeyHash))
	slog.Debug("to PubKeyHash:", "", string(toPubKeyHash))

	// 找到所有未使用的outputs确保能够支付给定amount
	// acc, validOutputs := bc.FindSpendableOutputs(fromPubKeyHash, amount)
	acc, validOutputs := u.FindSpendableOutputs(fromPubKeyHash, amount)
	slog.Debug("", "acc", acc, "amount", amount)
	if acc < amount {
		slog.Warn("Not enough money to spend!!!")
		fmt.Printf("Not enough money to spend!!!\n")
		return nil
	}

	// 对于每个找到的输出，创建一个引用它的输入
	for txid, outs := range validOutputs{
		txID, _ := hex.DecodeString(txid)

		for _, out := range outs {
			// 钱包的公钥传递给输入
			input := ts.TXInput{Txid:txID, Vout:out, Signature: nil, PubKey: wallet_from.PublicKey}
			inputs = append(inputs, input)
		}
	}

	//  每笔交易一般有1-2个输出，如果找到的outputs钱超过了就会找零
	outputs = append(outputs, *ts.NewTXOutput(amount, to))
	if acc > amount {
		// 找零钱是因为输出不可分割，超过了就另建一个输出
		outputs = append(outputs, *ts.NewTXOutput(acc-amount, from))
	}

	tx := ts.Transaction{ID:nil, VIn:inputs, VOut:outputs}
	tx.SetID()
	slog.Debug("New UTXO Transaction", "ID", fmt.Sprintf("%x", tx.ID))
	// 签名发生在新建UTXO交易中
	bc.SignTransaction(&tx, wallet_from.PrivateKey)
	return &tx
}

func(bc *BlockChain) SignTransaction(tx *ts.Transaction, privKey ecdsa.PrivateKey) {
	//TODO: 为交易签名
	// 先找到先前引用输出所在的交易
	prevTXs := bc.FindMapOfPrevTransactions(tx)
	tx.Sign(privKey, prevTXs)
	slog.Info("sign success for transaction", "ID", fmt.Sprintf("%x", tx.ID))
}

func(bc *BlockChain) VerifyTransaction(tx *ts.Transaction)bool{
	prevTXs := bc.FindMapOfPrevTransactions(tx)
	return tx.Verify(prevTXs)
}

