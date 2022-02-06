package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
	"log"
)

type EnsOperator struct {
	ethclient *ethclient.Client
}

func NewEnsOperator() *EnsOperator {
	ethClient, err := ethclient.Dial("https://eth-mainnet.alchemyapi.io/v2/Oqj-uq0ARWn6Nh0Quxj5NNqS-Q5ST-H1")
	if err != nil {
		log.Fatal(err)
	}
	return &EnsOperator{ethclient: ethClient}
}

func (ensOperator EnsOperator) ResolveName(ensName string) (common.Address, error) {
	return ens.Resolve(ensOperator.ethclient, ensName)
}

func (ensOperator EnsOperator) ResolveAddress(address string) (string, error) {
	addressString := common.HexToAddress(address)
	return ens.ReverseResolve(ensOperator.ethclient, addressString) // todo if address does not have ens name, it returns empty string - handle this correctly
}

func (ensOperator EnsOperator) GetAvatarByName(ensName string) (string, error) {
	resolver, err := ens.NewResolver(ensOperator.ethclient, ensName)
	if err != nil {
		return "", err
	}
	return resolver.Text("avatar")
}

func (ensOperator EnsOperator) GetAvatarByAddress(address string) (string, error) {
	name, err := ensOperator.ResolveAddress(address)
	if err != nil {
		return "", err
	}
	return ensOperator.GetAvatarByName(name)
}
