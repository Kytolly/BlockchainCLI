package server_model

import (
	ts "blockchain/internal/transaction_model"
)

var mempool = make(map[string]ts.Transaction) 