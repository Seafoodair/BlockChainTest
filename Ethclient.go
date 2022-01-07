package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"time"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct { //定义客户端结构
	rpcClient *rpc.Client       //支持rpc
	EthClient *ethclient.Client //支持普通的客户端
}

//官方提供的client 接口
func Connectoffice(host string) (*ethclient.Client, error) {
	ctx, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	conn := ethclient.NewClient(ctx)
	return conn, nil
}

// Connect creates a client that uses the given host.
func Connect(host string) (*Client, error) { //连接函数加上了rpc
	rpcClient, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	ethClient := ethclient.NewClient(rpcClient)
	return &Client{rpcClient, ethClient}, nil
}

// GetBlockNumber returns the block number.获取当前区块的区块号
func (ec *Client) GetBlockNumber(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := ec.rpcClient.CallContext(ctx, &result, "eth_blockNumber")
	return (*big.Int)(&result), err
}

// Message is a fully derived transaction and implements core.Message  交易类型，包括隐私数据
type Message struct {
	To       *common.Address `json:"to"`
	From     common.Address  `json:"from"`
	Value    string          `json:"value"`
	GasLimit string          `json:"gas"`
	GasPrice string          `json:"gasPrice"`
	Data     []byte          `json:"data"`
}

// NewMessage returns the message.获取到要送出的交易的struct
func NewMessage(from common.Address, to *common.Address, value *big.Int, gasLimit *big.Int, gasPrice *big.Int, data []byte) Message {
	return Message{
		From:     from,
		To:       to,
		Value:    toHexInt(value),
		GasLimit: toHexInt(gasLimit),
		GasPrice: toHexInt(gasPrice),
		Data:     data,
	}
}

// SendTransaction injects a transaction into the pending pool for execution. 重写一个SendTransaction 函数
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (ec *Client) SendTransaction(ctx context.Context, tx *Message) error {
	err := ec.rpcClient.CallContext(ctx, nil, "eth_sendTransaction", tx)
	return err
}

//取的交易編號就可以確認交易是否完成。 改名为SendTransactionnum
// SendTransaction injects a transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (ec *Client) SendTransactionnum(ctx context.Context, tx *Message) (common.Hash, error) {
	var txHash common.Hash
	err := ec.rpcClient.CallContext(ctx, &txHash, "eth_sendTransaction", tx)
	return txHash, err
}

func main() {
	//用功能获取区块链的区块
	client, err := Connect("http://localhost:8545")
	if err != nil {
		fmt.Println(err.Error())
	}
	blockNumber, err := client.GetBlockNumber(context.TODO())

	//发送一笔交易，原先的client客户端有sendRawTransaction（发送原始的交易）方法

	//发送一笔交易
	from := common.HexToAddress("Your from address") //这个地址需要更改
	to := common.HexToAddress("Your to address")     //这个地址需要更改
	amount := big.NewInt(1)
	gasLimit := big.NewInt(90000)
	gasPrice := big.NewInt(0)
	data := []byte{}
	message := eth.NewMessage(from, &to, amount, gasLimit, gasPrice, data)
	fmt.Println(message)
	err = client.SendTransaction(context.TODO(), &message)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Transaction has been sent")

	//取得交易资料
	// tx types.Transaction
	// isPending boolean
	tx, isPending, _ := client.EthClient.TransactionByHash(context.TODO(), txHash)

	//并发确认 交易是否完成
	// remember to use *types.Receipt
	receiptChan := make(chan *types.Receipt)
	// check transaction receipt
	go func() {
		fmt.Printf("Check transaction: %s\n", txHash.String())
		for {
			receipt, _ := client.EthClient.TransactionReceipt(context.TODO(), txHash)
			if receipt != nil {
				receiptChan <- receipt
				break
			} else {
				fmt.Println("Retry after i second")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	receipt := <-receiptChan
	fmt.Printf("Transaction status: %v\n", receipt.Status)
}
