buildcompare:
	go build -o build/compare compare-chain/compare

data/contract.csv:
	mongoexport --db=testnet --collection=contract --fields=_id --type=csv --out=data/contract.csv

data/account.csv:
	mongoexport --db=testnet --collection=account --fields=_id --type=csv --out=data/account.csv
