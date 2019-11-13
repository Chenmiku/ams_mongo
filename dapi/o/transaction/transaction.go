package transaction

import (
	"ams_system/dapi/x/mlog"
	"ams_system/dapi/o/transaction_input"
	"ams_system/dapi/o/transaction_output"
	"db/mgo"
)

var objectTransactionLog = mlog.NewTagLog("object_transaction")

//Transaction
type Transaction struct {
	mgo.BaseModel          `bson:",inline"`
	Hash                string `bson:"hash,omitempty" json:"hash"`
	TotalExchanged               int `bson:"total_exchanged,omitempty" json:"total_exchanged"`
	Fees         int    `bson:"fees,omitempty" json:"fess"`
	Size              int    `bson:"size,omitempty" json:"size"`
	Version                int    `bson:"version,omitempty" json:"version"`
	DoubleSpend     bool    `bson:"double_spend,omitempty" json:"double_spend"`
	BlockHash           string    `bson:"block_hash,omitempty" json:"block_hash"`
	BlockIndex               int `bson:"block_index,omitempty" json:"block_index`
	BlockHeight   int    `bson:"block_height,omitempty" json:"block_height"`
	TotalBlock int    `bson:"total_block,omitempty" json:"total_block"`
	ConfirmedTime       int64    `bson:"confirmed_time,omitempty" json:"confirmed_time"`
	InputsTransaction       int    `bson:"inputs_transaction,omitempty" json:"inputs_transaction"`
	OutputsTransaction       int    `bson:"outputs_transaction,omitempty" json:"outputs_transaction"`
	Inputs					[]transaction_input.TransactionInput		
	Outputs					[]transaction_output.TransactionOutput		
	IsCoinBase       bool    `bson:"is_coinbase,omitempty" json:"is_coinbase"`
	Addresses       []string    `bson:"addresses,omitempty" json:"addresses"`
	ToSign       []string    `bson:"to_sign,omitempty" json:"to_sign"`
	Signatures       []string    `bson:"signatures,omitempty" json:"signatures"`
	PublicKeys       []string    `bson:"public_keys,omitempty" json:"public_keys"`
	CTime                  int64  `bson:"ctime,omitempty" json:"ctime"` // Create Time
}

func NewCleanTransaction() interface{} {
	return &Transaction{}
}
