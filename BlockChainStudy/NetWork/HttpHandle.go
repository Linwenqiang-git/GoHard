package httphandle

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	Block "github.linwenqiang.com/BlockChainStudy/Block"
	
)

//post请求返回数据
type Message struct {
	Data int
}

//全局变量
var Blockchain []Block.Block
var mutex = &sync.Mutex{}

//运行server
func Run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//每次启动添加一个创世区块
	go func() {
		t := time.Now()
		genesisBlock := Block.Block{}
		genesisBlock = Block.Block{0, t.String(), 0, Block.CalculateHash(genesisBlock), "", 2, ""} //创世区块
		spew.Dump(genesisBlock)
		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock() //格式化 structs 和 slices
	}()

	mux := MakeMuxRouter()
	httpAddr := os.Getenv("PORT")
	log.Println("Listening on ", os.Getenv("PORT"))
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

//处理http请求
func MakeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", HandleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", HandleWriteBlock).Methods("POST")
	return muxRouter
}

func HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		RespondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	newBlock, err := Block.GenerateBlock(Blockchain[len(Blockchain)-1], m.Data)
	if err != nil {
		RespondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if Block.IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		newBlockchain := append(Blockchain, newBlock)
		log.Printf("成功添加一个新的区块链")
		Blockchain = Block.ReplaceChain(newBlockchain, Blockchain)
		spew.Dump(Blockchain)
	}

	RespondWithJSON(w, r, http.StatusCreated, newBlock)
}

//统一的 JSON数据 返回方法
func RespondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
