package nodemanage

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	Block "github.linwenqiang.com/BlockChainStudy/Block"
)

var bcServer chan []Block.Block
var Blockchain []Block.Block

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []Block.Block)

	// create genesis block
	t := time.Now()
	genesisBlock := Block.Block{}
	genesisBlock = Block.Block{0, t.String(), 0, Block.CalculateHash(genesisBlock), "", 2, ""} //创世区块
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "Enter a new Data:")

	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			data, err := strconv.Atoi(scanner.Text())

			if err != nil {
				log.Printf("%v not a number: %v", scanner.Text(), err)
				continue
			}
			newBlock, err := Block.GenerateBlock(Blockchain[len(Blockchain)-1], data)

			if err != nil {
				log.Println(err)
				continue
			}

			if Block.IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
				newBlockchain := append(Blockchain, newBlock)
				Blockchain = Block.ReplaceChain(newBlockchain, Blockchain)
			}

			bcServer <- Blockchain
			io.WriteString(conn, "\n Enter a new Data:")

		}
	}()

	go func() {
		//每隔30秒会同步第一个终端的最新的区块链数据给所有连接的节点
		for {
			time.Sleep(30 * time.Second)
			output, err := json.Marshal(Blockchain)

			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump(Blockchain)
	}
}
