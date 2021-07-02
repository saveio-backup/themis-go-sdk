module github.com/saveio/themis-go-sdk

go 1.14

replace github.com/saveio/themis => ../themis

require (
	github.com/gorilla/websocket v1.4.2
	github.com/itchyny/base58-go v0.1.0
	github.com/saveio/themis v0.0.0-20210702080134-9d051264abe3
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc
)
