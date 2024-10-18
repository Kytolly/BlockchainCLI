clear
cd ..
cd cmd/service
go build --o ../../bin/blockchain_go main.go
cd ../../bin

./blockchain_go createwallet
./blockchain_go createwallet
./blockchain_go createwallet
# Welcome to the blockchain CLI!
# New address: Vv2VG1aZJ5g8DR7Gqnr6RZRptqP3NzSx1i78YF
# Welcome to the blockchain CLI!
# New address: Vv2VGAuNkzDWxChKghXLMuaqan6wrTBYXoJbKG
# Welcome to the blockchain CLI!
# New address: Vv2VGKAQs1Wocu9StHMLiypU59gPM7kDN5hzzb

./blockchain_go new --address  Vv2VG1aZJ5g8DR7Gqnr6RZRptqP3NzSx1i78YF
./blockchain_go getbalance --address Vv2VG1aZJ5g8DR7Gqnr6RZRptqP3NzSx1i78YF
./blockchain_go getbalance --address Vv2VGKAQs1Wocu9StHMLiypU59gPM7kDN5hzzb
./blockchain_go getbalance --address Vv2VGAuNkzDWxChKghXLMuaqan6wrTBYXoJbKG

./blockchain_go send --from Vv2VG1aZJ5g8DR7Gqnr6RZRptqP3NzSx1i78YF --to Vv2VGKAQs1Wocu9StHMLiypU59gPM7kDN5hzzb --amount 5
./blockchain_go send --from Vv2VG1aZJ5g8DR7Gqnr6RZRptqP3NzSx1i78YF --to Vv2VGAuNkzDWxChKghXLMuaqan6wrTBYXoJbKG --amount 1