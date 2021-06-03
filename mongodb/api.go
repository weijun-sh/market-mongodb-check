package mongodb

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/anyswap/ANYToken-distribution/log"
	"gopkg.in/mgo.v2"
	//"github.com/davecgh/go-spew/spew"
)

const (
	retryDBCount    = 3
	retryDBInterval = 1 * time.Second
)

// TryDoTimes try do again if meet error
func TryDoTimes(name string, f func() error) (err error) {
	for i := 0; i < retryDBCount; i++ {
		err = f()
		if err == nil || mgo.IsDup(err) {
			return nil
		}
		time.Sleep(retryDBInterval)
	}
	log.Warn("[mongodb] TryDoTimes", "name", name, "times", retryDBCount, "err", err)
	return err
}

type block struct {
    Id string
    Number uint64
    Hash string
    //ParentHash string
    //Nonce string
    //Miner string
    //Difficulty string
    //GasLimit uint64
    //GasUsed uint64
    //Timestamp uint64
}

// --------------- find ---------------------------------
func FindBlocks() (uint64, error) {
	var res block
	err := collectionBlocks.Find(nil).Sort("-number").One(&res)
	if err != nil {
		return uint64(0), err
	}
	fmt.Printf("       | Block(number): %v\n", res.Number)
	return res.Number, nil
}

type transaction struct {
	Id string
	Hash string
	BlockNumber int64
	//BlockHash string
	//TransactionIndex uint64
	//From string
	//Nonce uint64
	//Timestamp uint64
}

func FindTransactions() (uint64, error) {
	var res interface{}
	var r transaction
	err := collectionTransactions.Find(nil).Sort("-blockNumber").One(&res)
	if err != nil {
		return uint64(0), err
	}
	c, err := json.Marshal(res)
    if err != nil {
        fmt.Println(err)
		return uint64(0), err
    }
    err = json.Unmarshal(c, &r)
    if err != nil {
        fmt.Println(err)
		return uint64(0), err
    }
	fmt.Printf("       | Txs (bnumber): %v\n", r.BlockNumber)
	return uint64(r.BlockNumber), nil
}

