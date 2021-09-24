package transfer

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

// expose transfer module api for user
type Client interface {
	sdk.Module

	Transfer(request TransferRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)

	QueryDenomTrace(request QueryDenomTraceRequest) (QueryDenomTraceResponse, error)
	QueryDenomTraces(request QueryDenomTracesRequest) (QueryDenomTracesResponse, error)
}

type TransferRequest struct {
	SourcePort       string   `json:"source_port"`
	SourceChannel    string   `json:"source_channel"`
	Token            sdk.Coin `json:"token"`
	Sender           string   `json:"sender"`
	Receiver         string   `json:"receiver"`
	TimeoutHeight    Height   `json:"timeout_height"`
	TimeoutTimestamp uint64   `json:"timeout_timestamp"`
}
