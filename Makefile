buildcompare:
	go build -o build/compare compare-chain/compare

data/contract.csv:
	mongoexport --db=testnet --collection=contract --fields=_id --type=csv --out=data/contract.csv

data/account.csv:
	mongoexport --db=testnet --collection=account --fields=_id --type=csv --out=data/account.csv

data/erc20.csv:
	mongo --eval 'db.createView("tmpErc20s", "erc20trx", [{ $group: { _id: "$tok" }}])' testnet
	mongoexport --db=testnet --collection=tmpErc20s --fields=_id --type=csv --out=data/erc20.csv

data/erc20acc.csv:
	mongo --eval 'db.erc20trx.aggregate([ { $match: { tty: "ERC20" } }, { $group: { _id: { tok: "$tok", to: "$to" } } }, { $merge: { into: "tmpErc20disponents", whenMatched: "keepExisting" } } ], { allowDiskUse: true });' testnet
	mongoexport --db=testnet --collection=tmpErc20disponents --fields=_id.tok,_id.to --type=csv --out=data/erc20disp.csv
