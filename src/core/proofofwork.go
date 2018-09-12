package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
)

var(
	maxNonce = math.MaxInt64
)

const targetBits =20

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target,uint(256 - targetBits))

	pow := &ProofOfWork{b,target}
	return pow
}

func (pow *ProofOfWork)prepareData(nonce int) string  {
	data := string(pow.block.Timestamp) + string(pow.block.Index) + pow.block.PreBlockHash +string(nonce) +string(targetBits)
	return data
}
func (pow *ProofOfWork) Run()(int, string){
	var hash string
	var hash_bytes [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n",pow.block.data)
	for nonce < maxNonce  {
		data := pow.prepareData(nonce)

		hash_bytes = sha256.Sum256([]byte(data))
		hash = hex.EncodeToString(hash_bytes[:])
		fmt.Printf("\r%x", hash_bytes)
		nonce ++
	}
	fmt.Println("ok")
	return nonce, hash
}
