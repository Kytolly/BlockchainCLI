# 中心结点创建钱包和一个新区块链:3000
export NODE_ID=3000
./blockchain createwallet 
# centernode:  1GVRd2fuWopKryR2CVcJbUXwsKqc9RoVqp
./blockchain new --address 1GVRd2fuWopKryR2CVcJbUXwsKqc9RoVqp
cp chain_3000.db genesis.db 
./blockchain getbalance --address 1GVRd2fuWopKryR2CVcJbUXwsKqc9RoVqp

# 余额50

# 钱包结点创建三个钱包:3001
export NODE_ID=3001
./blockchain createwallet
./blockchain createwallet
./blockchain createwallet
# wallet1:  14fMyMpdQrLJujnudwWzjS4rumrUktJUvi
# wallet2:  1EbPutZHr9aC4MAZREe6Vpnq42f23FEY21
# wallet3:  1MYVMSGfrDtkjryX3KpzNBKEjHXUmVwjYh
cp genesis.db chain_3001.db
./blockchain getbalance --address 14fMyMpdQrLJujnudwWzjS4rumrUktJUvi
./blockchain getbalance --address 1EbPutZHr9aC4MAZREe6Vpnq42f23FEY21
./blockchain getbalance --address 1MYVMSGfrDtkjryX3KpzNBKEjHXUmVwjYh

# 0,0,0

# 发送硬币到钱包结点:3000
./blockchain send --from 1GVRd2fuWopKryR2CVcJbUXwsKqc9RoVqp --to 14fMyMpdQrLJujnudwWzjS4rumrUktJUvi --amount 10 --mine
./blockchain send --from 1GVRd2fuWopKryR2CVcJbUXwsKqc9RoVqp --to 1EbPutZHr9aC4MAZREe6Vpnq42f23FEY21 --amount 10 --mine 
./blockchain start

# 3001
./blockchain start

# 检查余额：3001
./blockchain getbalance --address 14fMyMpdQrLJujnudwWzjS4rumrUktJUvi
./blockchain getbalance --address 1EbPutZHr9aC4MAZREe6Vpnq42f23FEY21

# 检查余额：3000
./blockchain getbalance --address 1GVRd2fuWopKryR2CVcJbUXwsKqc9RoVqp


# 创建新矿工结点:3002
export NODE_ID=3002
cp genesis.db chain_3002.db
./blockchain createwallet 
./blockchain createwallet 
# wallet4:      1BRCUMw35fkuzUqjLjFLX9KzrzWsSB3MUk
# mineraddr:    1BkqQ2iLeDyFy1NuUpfiXCsHJPpv9Wxdat
./blockchain getbalance --address 1BRCUMw35fkuzUqjLjFLX9KzrzWsSB3MUk
./blockchain getbalance --address 1BkqQ2iLeDyFy1NuUpfiXCsHJPpv9Wxdat

# 0,0
./blockchain start --miner 1BkqQ2iLeDyFy1NuUpfiXCsHJPpv9Wxdat


# 发送硬币：3001
./blockchain send --from 14fMyMpdQrLJujnudwWzjS4rumrUktJUvi --to 1MYVMSGfrDtkjryX3KpzNBKEjHXUmVwjYh --amount 1
./blockchain send --from 1EbPutZHr9aC4MAZREe6Vpnq42f23FEY21 --to 1BRCUMw35fkuzUqjLjFLX9KzrzWsSB3MUk --amount 2


# 矿工检查区块：3002
./blockchain print

# 检查中心结点：3000
./blockchain getbalance --address 1GVRd2fuWopKryR2CVcJbUXwsKqc9RoVqp
./blockchain print

# 启动钱包结点，下载挖出来的区块:3001
./blockchain start

# 一切测试结束，查看对应结点的钱包余额
./blockchain getbalance --address 1GVRd2fuWopKryR2CVcJbUXwsKqc9RoVqp
./blockchain getbalance --address 14fMyMpdQrLJujnudwWzjS4rumrUktJUvi
./blockchain getbalance --address 1EbPutZHr9aC4MAZREe6Vpnq42f23FEY21
./blockchain getbalance --address 1MYVMSGfrDtkjryX3KpzNBKEjHXUmVwjYh
./blockchain getbalance --address 1BRCUMw35fkuzUqjLjFLX9KzrzWsSB3MUk
./blockchain getbalance --address 1BkqQ2iLeDyFy1NuUpfiXCsHJPpv9Wxdat