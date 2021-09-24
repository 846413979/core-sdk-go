package types

import (
	"google.golang.org/grpc"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
)

type TxManager interface {
	TmQuery
	SendBatch(msgs Msgs, baseTx BaseTx) ([]ResultTx, error)
	BuildAndSend(msg []Msg, baseTx BaseTx) (ResultTx, error)
	BuildAndSign(msg []Msg, baseTx BaseTx) ([]byte, error)
	BuildTxHash(msg []Msg, baseTx BaseTx) (string, error)
	BuildAndSendWithAccount(addr string, accountNumber, sequence uint64, msg []Msg, baseTx BaseTx) (ResultTx, error)
}

type Queries interface {
	StoreQuery
	AccountQuery
	TmQuery
}

type GRPCClient interface {
	GenConn() (*grpc.ClientConn, error)
}

type ParamQuery interface {
	QueryParams(module string, res Response) error
}

type StoreQuery interface {
	QueryWithResponse(path string, data interface{}, result Response) error
	Query(path string, data interface{}) ([]byte, error)
	QueryStore(key HexBytes, storeName string, height int64, prove bool) (abci.ResponseQuery, error)
}

type AccountQuery interface {
	QueryAccount(address string) (BaseAccount, error)
	QueryAddress(name, password string) (AccAddress, error)
}

type TmQuery interface {
	QueryTx(hash string) (ResultQueryTx, error)
	QueryTxs(builder *EventQueryBuilder, page, size *int) (ResultSearchTxs, error)
	QueryBlock(height int64) (BlockDetail, error)
}

type TokenManager interface {
	QueryToken(denom string) (Token, error)
	ToMinCoin(coin ...DecCoin) (Coins, error)
	ToMainCoin(coin ...Coin) (DecCoins, error)
}

type Logger interface {
	Logger() log.Logger
	SetLogger(log.Logger)
}

type WSClient interface {
	SubscribeNewBlock(builder *EventQueryBuilder, handler EventNewBlockHandler) (Subscription, error)
	SubscribeTx(builder *EventQueryBuilder, handler EventTxHandler) (Subscription, error)
	SubscribeNewBlockHeader(handler EventNewBlockHeaderHandler) (Subscription, error)
	SubscribeValidatorSetUpdates(handler EventValidatorSetUpdatesHandler) (Subscription, error)
	Unsubscribe(subscription Subscription) error
}

type TmClient interface {
	ABCIClient
	SignClient
	WSClient
	StatusClient
	NetworkClient
}

type BaseClient interface {
	TokenManager
	TxManager
	Queries
	TmClient
	Logger
	GRPCClient
	KeyManager
}
