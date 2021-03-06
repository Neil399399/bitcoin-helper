package bitcoin

// VoutDetail is bitcoin transaction detail
type VoutDetail struct {
	BlockNumber   int64
	BlockHash     string
	Txid          string
	Address       []string
	Category      string
	Amount        float64
	Vout          uint32
	Confirmations uint64
	Time          int64
	LockTime      uint32
	Blocktime     int64
	Memo          string
}

// Utxo is bitcoin Unspent Transaction Output
type Utxo struct {
	Address     string
	TxID        string
	OutputIndex uint32
	Script      []byte
	Satoshis    int64
}

// SignInfo information of sign key
type SignInfo struct {
	CoinType string
	Network  string
	KeyID    string
	ChildIdx string
}
