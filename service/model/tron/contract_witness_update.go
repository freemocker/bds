package tron

type WitnessUpdateContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	UpdateUrl       string `xorm:"update_url varchar(1024) notnull"`
}

func (c WitnessUpdateContract) TableName() string {
	return tableName("contract_witness_update")
}
