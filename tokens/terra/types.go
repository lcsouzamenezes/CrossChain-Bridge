package terra

import (
	"time"
)

// GetBlockResult get block result
type GetBlockResult struct {
	Block *Block `json:"block"`
}

// Block block
type Block struct {
	Header *Header `json:"header"`
}

// Header header
type Header struct {
	ChainID string    `json:"chain_id"`
	Height  string    `json:"height"`
	Time    time.Time `json:"time"`
}

// GetTxResult gettx result
type GetTxResult struct {
	Tx         Tx         `json:"tx"`
	TxResponse TxResponse `json:"tx_response"`
}

// Tx tx
type Tx struct {
}

// TxResponse tx response
type TxResponse struct {
}

// BroadcastTxRequest broadcat tx request
type BroadcastTxRequest struct {
	TxBytes string `json:"tx_bytes"`
	Mode    string `json:"mode"`
}

// BroadcastTxResult broadcast tx result
type BroadcastTxResult struct {
	TxResponse BroadcastTxResponse `json:"tx_response"`
}

// BroadcastTxResponse broadcast tx response
type BroadcastTxResponse struct {
	TxHash string `json:"txhash"`
	Code   int64  `json:"code"`
}

// SimulateRequest simulate request
type SimulateRequest struct {
}

// SimulateResponse simulate responce
type SimulateResponse struct {
}

// GetBaseAccountResult get base account result
type GetBaseAccountResult struct {
	Account *BaseAccount `json:"account"`
}

// BaseAccount base account
type BaseAccount struct {
	TypeURL       string  `json:"@type"`
	Address       string  `json:"address"`
	Pubkey        *Pubkey `json:"pub_key,omitempty"`
	AccountNumber string  `json:"account_number"`
	Sequence      string  `json:"sequence"`
	Value         string  `json:"value,omitempty"`
}

// Pubkey public key
type Pubkey struct {
	TypeURL string `json:"@type"`
	Key     string `json:"key"`
}
