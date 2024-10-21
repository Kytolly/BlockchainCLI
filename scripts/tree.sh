rm /home/kytolly/Project/gowork/blockchain/bin/chain.db
rm /home/kytolly/Project/gowork/blockchain/bin/wallet.db
rm /home/kytolly/Project/gowork/blockchain/log/chain.log
go build -o /home/kytolly/Project/gowork/blockchain/bin/blockchain_go /home/kytolly/Project/gowork/blockchain/cmd/service/main.go

cd /home/kytolly/Project/gowork/blockchain/bin

clear
./blockchain_go createwallet
./blockchain_go createwallet
./blockchain_go createwallet

# Welcome to the blockchain CLI!
# New address: 3BruC7Up8ReBoiCu5tVoMwMyxY6Hgh7yoDYo46Jx
# Welcome to the blockchain CLI!
# New address: 3BruC7UhggE9VXFroHZLEHsqX5oRTWAw1NzpjVcL
# Welcome to the blockchain CLI!
# New address: 3BruC7UhHucGe6Cw2gHz3hMimCRJNiUua9FGuqdU


./blockchain_go new --address 3BruC7Up8ReBoiCu5tVoMwMyxY6Hgh7yoDYo46Jx
# ./blockchain_go new --address 3BruC7UhggE9VXFroHZLEHsqX5oRTWAw1NzpjVcL
# ./blockchain_go new --address 3BruC7UhHucGe6Cw2gHz3hMimCRJNiUua9FGuqdU

./blockchain_go getbalance --address 3BruC7Up8ReBoiCu5tVoMwMyxY6Hgh7yoDYo46Jx 
./blockchain_go getbalance --address 3BruC7UhggE9VXFroHZLEHsqX5oRTWAw1NzpjVcL
./blockchain_go getbalance --address 3BruC7UhHucGe6Cw2gHz3hMimCRJNiUua9FGuqdU

rm /home/kytolly/Project/gowork/blockchain/log/chain.log


./blockchain_go send --from 3BruC7Up8ReBoiCu5tVoMwMyxY6Hgh7yoDYo46Jx  --to 3BruC7UhggE9VXFroHZLEHsqX5oRTWAw1NzpjVcL --amount 3
./blockchain_go getbalance --address 3BruC7Up8ReBoiCu5tVoMwMyxY6Hgh7yoDYo46Jx
./blockchain_go getbalance --address 3BruC7UhggE9VXFroHZLEHsqX5oRTWAw1NzpjVcL

./blockchain_go send --from 3BruC7Up8ReBoiCu5tVoMwMyxY6Hgh7yoDYo46Jx  --to 3BruC7UhHucGe6Cw2gHz3hMimCRJNiUua9FGuqdU --amount 2
./blockchain_go getbalance --address 3BruC7Up8ReBoiCu5tVoMwMyxY6Hgh7yoDYo46Jx 
./blockchain_go getbalance --address 3BruC7UhHucGe6Cw2gHz3hMimCRJNiUua9FGuqdU