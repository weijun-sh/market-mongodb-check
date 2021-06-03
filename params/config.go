package params

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/anyswap/CrossChain-Bridge/common"
	"github.com/anyswap/CrossChain-Bridge/log"
)

// swap tx types
const (
	TxSwapin   = "swapin"
	TxSwapout  = "swapout"
	TxSwapout2 = "swapout2"
)

var (
	configFile string
	Chain string
	scanConfig = &ScanConfig{}
)

// ScanConfig scan config
type ScanConfig struct {
	FSN  *MongodbConfig
	BSC  *MongodbConfig
	HT  *MongodbConfig
	FTM  *MongodbConfig
}

// MongoDBConfig mongodb config
type MongoDBConfig struct {
        DBURL      string
        DBName     string
        UserName   string `json:"-"`
        Password   string `json:"-"`
	Enable     bool
	BlockChain string
}

// TokenConfig token config
type MongodbConfig struct {
	DBURL string
	DBName string
	UserName string
	Password string `toml:",omitempty" json:",omitempty"`
}

// GetScanConfig get scan config
func GetScanConfig() *MongodbConfig {
	switch Chain {
	case "FSN":
		return scanConfig.FSN
	case "BSC":
		return scanConfig.BSC
	case "FTM":
		return scanConfig.FTM
	case "HT":
		return scanConfig.HT
	default:
		return scanConfig.FSN
	}
}

// LoadConfig load config
func LoadConfig(filePath string) *ScanConfig {
	log.Println("LoadConfig Config file is", filePath)
	if !common.FileExist(filePath) {
		log.Fatalf("LoadConfig error: config file '%v' not exist", filePath)
	}

	config := &ScanConfig{}
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		log.Fatalf("LoadConfig error (toml DecodeFile): %v", err)
	}

	var bs []byte
	if log.JSONFormat {
		bs, _ = json.Marshal(config)
	} else {
		bs, _ = json.MarshalIndent(config, "", "  ")
	}
	log.Println("LoadConfig finished.", string(bs))

	scanConfig = config

	configFile = filePath // init config file path
	return scanConfig
}

