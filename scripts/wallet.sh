cd cmd/service
go build --o ../../bin/blockchain_go main.go
cd ../../bin

./blockchain_go createwallet
# New address: 1Vv2VGBWSG3m68nbF2bUWFBmVi63VMGxh7t9nER

./blockchain_go createwallet
# New address: 1Vv2VG7PwwoCSmCoKTZ8yyur2p18kRP7NqNNLR8

./blockchain_go createwallet
# New address: 1Vv2VG2XZcUR4moGso3QjiiEyTQMXrEE2YsZaAa

./blockchain_go new --address 1Vv2VGBWSG3m68nbF2bUWFBmVi63VMGxh7t9nER
0000005420fbfdafa00c093f56e033903ba43599fa7cd9df40458e373eee724d

Done!

./blockchain_go getbalance --address 1Vv2VGDbu2WGxJJ3xQrEJQenuyyvSiBXXjx4n1x
Balance of 1Vv2VGDbu2WGxJJ3xQrEJQenuyyvSiBXXjx4n1x:10

./blockchain_go getbalance --address 1Vv2VG7PwwoCSmCoKTZ8yyur2p18kRP7NqNNLR8
Balance of 1Vv2VGDbu2WGxJJ3xQrEJQenuyyvSiBXXjx4n1x:0


./blockchain_go send --from 1Vv2VG7PwwoCSmCoKTZ8yyur2p18kRP7NqNNLR8 --to 1Vv2VGDbu2WGxJJ3xQrEJQenuyyvSiBXXjx4n1x --amount 5
2017/09/12 13:08:56 ERROR: Not enough funds

./blockchain_go send --from 1Vv2VGDbu2WGxJJ3xQrEJQenuyyvSiBXXjx4n1x --to 1Vv2VG7PwwoCSmCoKTZ8yyur2p18kRP7NqNNLR8 --amount 6
00000019afa909094193f64ca06e9039849709f5948fbac56cae7b1b8f0ff162

Success!

./blockchain_go getbalance --address 13Uu7B1vDP4ViXqHFsWtbraM3EfQ3UkWXt
Balance of '13Uu7B1vDP4ViXqHFsWtbraM3EfQ3UkWXt': 4

./blockchain_go getbalance --address 15pUhCbtrGh3JUx5iHnXjfpyHyTgawvG5h
Balance of '15pUhCbtrGh3JUx5iHnXjfpyHyTgawvG5h': 6

./blockchain_go getbalance --address 1Lhqun1E9zZZhodiTqxfPQBcwr1CVDV2sy
Balance of '1Lhqun1E9zZZhodiTqxfPQBcwr1CVDV2sy': 0






./blockchain_go createwallet
Your new address: 1AmVdDvvQ977oVCpUqz7zAPUEiXKrX5avR

./blockchain_go createwallet
Your new address: 1NE86r4Esjf53EL7fR86CsfTZpNN42Sfab

./blockchain_go createblockchain --address 1AmVdDvvQ977oVCpUqz7zAPUEiXKrX5avR
000000122348da06c19e5c513710340f4c307d884385da948a205655c6a9d008

Done!

./blockchain_go send --from 1AmVdDvvQ977oVCpUqz7zAPUEiXKrX5avR --to 1NE86r4Esjf53EL7fR86CsfTZpNN42Sfab --amount 6
0000000f3dbb0ab6d56c4e4b9f7479afe8d5a5dad4d2a8823345a1a16cf3347b

Success!

./blockchain_go getbalance --address 1AmVdDvvQ977oVCpUqz7zAPUEiXKrX5avR
Balance of '1AmVdDvvQ977oVCpUqz7zAPUEiXKrX5avR': 4

./blockchain_go getbalance --address 1NE86r4Esjf53EL7fR86CsfTZpNN42Sfab
Balance of '1NE86r4Esjf53EL7fR86CsfTZpNN42Sfab': 6