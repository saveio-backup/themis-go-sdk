module github.com/saveio/themis-go-sdk

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.15
	github.com/gorilla/websocket v1.4.2
	github.com/itchyny/base58-go v0.1.0
	github.com/saveio/themis v1.0.175-0.20221129100116-4008d4b1ba7e
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc
)

replace github.com/saveio/themis => ../themis