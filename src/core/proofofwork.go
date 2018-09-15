package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	//定义工作量证明最大值为64位整数的最大值
	maxNonce = math.MaxInt64
)
//目标
const targetBits = 20
//定义工作量结构体
type ProofOfWork struct {
	//被计算的区块
	block *Block
	//目标，对区块计算要满足这个目标
	target *big.Int
}
/**
 * 创建工作量证明函数
 */
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//对整数target前面的比特位进行移位操作，因为target是1,移位后前面的字节变为0
	target.Lsh(target, uint(256 - targetBits))
	//创建工作量证明并返回
	pow := &ProofOfWork{b,target}
	return  pow
}
/**
 *运行工作量证明的过程
 */
func (pow *ProofOfWork)Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	//提示，显示这个区块值
	fmt.Printf("Mining the block containing\"%s\"\n ", pow.block.Data)
	//通过循环反复计算
	for nonce < maxNonce {
		//获取拼装好的byte数组格式的数据块
		data := pow.prepareData(nonce)
		//对数据块进行加密返回byte字节数组
		hash = sha256.Sum256(data)
		//显示这个hash值
		fmt.Printf("\r%x", hash)
		//将hash转换为hash整数
		hashInt.SetBytes(hash[:])
		//使用这个hash整数与target进行对比
		if hashInt.Cmp(pow.target) == -1{
			break
		}else {
			nonce ++
		}
	}
	fmt.Printf("\n\n")
 	return nonce, hash[:]
}

//将各个数据拼凑大数据返回byte数组
func (pow *ProofOfWork)prepareData(nonce int) []byte{
	data := bytes.Join(
		[][]byte{
			pow.block.PreBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			//nonce为可变值，所以循环中每轮的hash都不同
			IntToHex(int64(nonce)),
		},
		[]byte{},
		)
	return data
}

func (pow *ProofOfWork)Validate() bool{
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValidate := hashInt.Cmp(pow.target) == -1

	return isValidate
}

