package core
/**
 *区块链结构体，以数组的形式存放区块，理论上可以建立无数个区块
 */
type Blcokchain struct {
	Blcok []*Block
}
/**
 *添加区块函数
 */
func (bc *Blcokchain)AddBlock(data string){
	//获取区块链中最后一个区块
	preBlock := bc.Blcok[len(bc.Blcok)-1]
	//创建新区块
	newBlock := CreateNewBlock(preBlock.Hash, data)
	//将新区块追加如入区块中
	bc.Blcok = append(bc.Blcok, newBlock)
}
/**
 * 创建区块链，并创建创世函数
 */
func CreateNewBlcokchain() *Blcokchain{
	return  &Blcokchain{[]*Block{CreateGenesisBlcok()}}
}
