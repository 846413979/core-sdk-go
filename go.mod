module github.com/irisnet/core-sdk-go

go 1.16

require (
	github.com/99designs/keyring v1.1.6
	github.com/tharsis/ethermint v0.8.1
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/bluele/gcache v0.0.2
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/go-bip39 v1.0.0
	github.com/dvsekhvalnov/jose2go v0.0.0-20201001154944-b09cfaf05951
	github.com/ethereum/go-ethereum v1.10.16
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.4.2
	github.com/magiconair/properties v1.8.5
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mtibben/percent v0.2.1
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.8.0
	github.com/prometheus/common v0.23.0
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.11
	github.com/tendermint/tm-db v0.6.4
	github.com/tjfoc/gmsm v1.4.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
	github.com/tharsis/ethermint v0.8.1 => github.com/bianjieai/ethermint v0.8.2-0.20220211020007-9ec25dde74d4
)
