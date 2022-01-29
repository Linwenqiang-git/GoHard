package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index      int    //标识区块在区块链里的位置
	Timestamp  string //生成区块的时间戳
	Data       int    //要写入区块的数据
	Hash       string //整个区块数据 SHA256 的哈希值
	PrevHash   string //上一个区块的哈希值
	Difficulty int    //定义PoW的解答难度，即哈希值要有多少个0开头才能算是解答成功。
	Nonce      string //临时存放hash值
}

var difficulty int = 1

//计算区块哈希值
func CalculateHash(block Block) string {
	//将必要数据以区块的形式拼接起来
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.Data) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//模拟生成新的区块
func GenerateBlock(oldBlock Block, Data int) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	//newBlock.Hash = CalculateHash(newBlock)
	newBlock.Difficulty = difficulty

	//pow共识机制
	//这里不断循环，找到符合条件的数字
	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex

		if !IsHashValid(CalculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(CalculateHash(newBlock), " do more work!!")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(CalculateHash(newBlock), " work done!!")
			newBlock.Hash = CalculateHash(newBlock)
			break
		}
	}
	return newBlock, nil
}

//判断新生成的区块是否有效 -> 通过比对新旧区块的 Index, 哈希值来决定新生成的区块是否 Valid。
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

//验证哈希值是否以difficulty定义的个数的0开头。
func IsHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	//判断这个字符串是否有指定的前缀
	return strings.HasPrefix(hash, prefix)
}

//区块链会发生分叉，我们取更长的链作为正确的链:
func ReplaceChain(newBlocks []Block, globalChainBlocks []Block) []Block {
	if len(newBlocks) > len(globalChainBlocks) {
		return newBlocks
	}
	return globalChainBlocks
}
