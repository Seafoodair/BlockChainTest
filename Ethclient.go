package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func Connect(host string) (*ethclient.Client, error) {
	ctx, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	conn := ethclient.NewClient(ctx)
	return conn, nil
}
