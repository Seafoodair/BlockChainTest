package main

import "context"

//产生私钥，从字符串还原
// from hex  从16进制进行转化
privKey, err := crypto.HexToECDSA("your private key");
// from other
privKey, err := crypto.ToECDSA(DecodeStringToBytes("your private key"))
if err != nil {
	fmt.Println(err)
}else {
	// do something
}
//也可以 重新生成密钥
privKey, err := crypto.GenerateKey()

if err != nil {
	fmt.Println(err)
} else {
	// do something
}
//取得公钥和以太坊地址
publicKey :=  privKey.PublicKey
address := crypto.PubkeyToAddress(publicKey).Hex()

//签名交易
//主要可分三步 1.产生事务
amount := big.NewInt(1)
gasLimit := uint64(90000)
gasPrice := big.NewInt(0)
data := []byte{}
tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data) //产生事务
//产生Signer 的相关事务 （这里我认为是个随机数或者签名的方法）
// EIP155 signer
// signer := types.NewEIP155Signer(big.NewInt(4))
signer := types.HomesteadSigner{}

//用私钥签名
signedTx, _ := types.SignTx(tx, signer, privKey)
//新增一個 SendRawTransaction 函數來取得區塊編號
func (ec *Client) SendRawTransaction(ctx context.Context, tx *types.Transaction) (common.Hash, error) {
	var txHash common.Hash
	if data, err := rlp.EncodeToBytes(tx);err != nil {
		return txHash, err
	} else {
		err := ec.rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", common.ToHex(data))
		return txHash, err
	}
}



//发送交易
amount := big.NewInt(1)
gasLimit := uint64(90000)
gasPrice := big.NewInt(0)
data := []byte{}
tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
// EIP155 signer
// signer := types.NewEIP155Signer(big.NewInt(4))
signer := types.HomesteadSigner{}
signedTx, _ := types.SignTx(tx, signer, privKey)
// client.EthClient.SendTransaction(context.TODO(), signedTx)
txHash, err := client.SendRawTransaction(context.TODO(), signedTx)
// do something to txHash