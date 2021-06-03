package scanner

import (
	"fmt"
	"strings"
	"github.com/anyswap/CrossChain-Bridge/cmd/utils"
	"github.com/anyswap/CrossChain-Bridge/log"
	"github.com/weijun-sh/gethscan/params"
	"github.com/weijun-sh/gethscan/mongodb"
	"github.com/urfave/cli/v2"
)

var (
	// ScanSwapCommand scan swaps on eth like blockchain
	ScanMongodbCommand = &cli.Command{
		Action:    scanMongodb,
		Name:      "scanmongodb",
		Usage:     "scan mongodb",
		ArgsUsage: " ",
		Description: `
scan mongodb
`,
		Flags: []cli.Flag{
			utils.ConfigFileFlag,
			 &cli.StringFlag{
                     Name:  "chain",
                     Value: "FSN",
                     Usage: `chain selected`,
              },
			//utils.ChainFlag,
			//utils.TableFlag,
			//utils.KeyFlag,
		},
	}
)

var (
	table string
	chain string
	key string
)

func scanMongodb(ctx *cli.Context) error {
	params.LoadConfig(utils.GetConfigFilePath(ctx))
	params.Chain = strings.ToUpper(ctx.String("chain"))
	fmt.Printf("params.Chain: %v\n", params.Chain)

	InitMongodb()
	findBlocks()
	findTransactions()

	return nil
}

// InitMongodb init mongodb by config
func InitMongodb() {
	log.Info("InitMongodb")
        dbConfig := params.GetScanConfig()
        mongodb.MongoServerInit([]string{dbConfig.DBURL}, dbConfig.DBName, dbConfig.UserName, dbConfig.Password)
}

// FindAllTokenAccounts find accounts
func findBlocks() () {
	mongodb.FindBlocks()
}

// FindAllTokenAccounts find accounts
func findTransactions() () {
	mongodb.FindTransactions()
}

