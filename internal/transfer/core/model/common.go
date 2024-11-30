package model

import "time"

type EventType int

const (
	UserStatusActive  = "A"
	UserStatusSuspend = "S"

	// redis key
	DefaultIdempotenceTTL      = 1 * time.Hour
	IdempotenceTransferHandler = "trf:api:%s:%s"         // "trf:api:[account_no]:[id]" - TTL
	IdempotenceMoneyReply      = "trf:money:reply:%s:%s" // "trf:money:reply:[account_no]:[id]"
	IdempotenceBankingReply    = "trf:bank:reply:%s:%s"  // "trf:bank:reply:[account_no]:[id]"

	EventTypeSubmitTransfer EventType = iota
	EventTypeMoneyReply
	EventTypeBankReply

	// config key for topic name
	MoneyTopicKey   = "money_service"
	BankingTopicKey = "banking_service"
)
