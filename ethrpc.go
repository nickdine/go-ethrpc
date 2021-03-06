/*
Package ethrpc implements RPC methods to interact with the ethreum node geth.
*/
package ethrpc

import (
	"errors"

	"github.com/KeisukeYamashita/jsonrpc"
)

/*
RPCClient ...
*/
type RPCClient struct {
	*jsonrpc.RPCClient
}

/*
RPCer ...
*/
type RPCer interface {
	GetBlockNumber() (string, error)
}

/*
NewRPCClient creates JSONRPC clients for your bitcoin node.
*/
func NewRPCClient(endpoint string) *RPCClient {
	c := new(RPCClient)
	c.RPCClient = jsonrpc.NewRPCClient(endpoint)
	return c
}

/*
GetBlockNumber gets the most resent block height
*/
func (c *RPCClient) GetBlockNumber() (string, error) {
	resp, err := c.RPCClient.Call("eth_blockNumber")
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var heightHex string
	resp.GetObject(&heightHex)
	return heightHex, nil
}
