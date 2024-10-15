./main new -address Ivan
./main getbalance -address Ivan
./main send -from Ivan -to Pedro -amount 6
./main getbalance -address Ivan
./main getbalance -address Pedro
./main send -from Pedro -to Helen -amount 2
./main send -from Ivan -to Helen -amount 2
./main send -from Helen -to Rachel -amount 3
./main getbalance -address Ivan
./main getbalance -address Pedro
./main getbalance -address Helen
./main getbalance -address Rachel
./main send -from Pedro -to Ivan -amount 5
./main getbalance -address Pedro
./main getbalance -address Ivan