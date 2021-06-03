# gethscan

scan mongodb (block and transactions) table blocknumber

## building

```shell
make
```

this will generate a binary file `./build/bin/gethscan`,
and an example config file of `scanmongodb` subcommand [config-example.toml](https://github.com/weijun-sh/market-mongodb-check/blob/master/params/config-example.toml)

## help

#### gethscan

```shell
./build/bin/gethscan -h
```

```text
NAME:
   gethscan - scan eth like blockchain

USAGE:
   gethscan [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   scanmongodb  scan mongodb
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbosity value  0:panic, 1:fatal, 2:error, 3:warn, 4:info, 5:debug, 6:trace (default: 4)
   --json             output log in json format (default: false)
   --color            output log in color text format (default: true)
   --help, -h         show help (default: false)
   --version, -v      print the version (default: false)
```

#### gethscan scanmongodb

```shell
./build/bin/gethscan scanmongodb -h
```

```text
NAME:
   gethscan scanmongodb - scan mongodb

USAGE:
   gethscan scanmongodb [command options]

DESCRIPTION:
   scan mongodb

OPTIONS:
   --chain value             chain selected (default: FSN, others: BSC/FTM/HT)
   --config value, -c value  Specify config file
   --help, -h                show help (default: false)
```
