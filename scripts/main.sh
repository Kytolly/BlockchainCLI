#！/bin/bash

echo "开始测试..."
cd /home/kytolly/Project/gowork/blockchain/cmd/service
go build
./service printchain
./service addblock --data "Send 1 BTC to Ivan"
./service addblock --data "Pay 0.31337 BTC for a coffee"
./service printchain
echo "测试结束!"