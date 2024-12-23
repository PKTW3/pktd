// Copyright (c) 2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

//go:build !generate
// +build !generate

package rpchelp

import "github.com/pkt-cash/pktd/btcjson"

// Common return types.
var (
	returnsBool     = []interface{}{(*bool)(nil)}
	returnsNumber   = []interface{}{(*float64)(nil)}
	returnsString   = []interface{}{(*string)(nil)}
	returnsLTRArray = []interface{}{(*[]btcjson.ListTransactionsResult)(nil)}
)

// Methods contains all methods and result types that help is generated for,
// for every locale.
var Methods = []struct {
	Method      string
	ResultTypes []interface{}
}{
	{"addmultisigaddress", returnsString},
	{"createmultisig", []interface{}{(*btcjson.CreateMultiSigResult)(nil)}},
	{"createtransaction", returnsString},
	{"getaddressbalances", []interface{}{(*[]btcjson.GetAddressBalancesResult)(nil)}},
	{"setnetworkstewardvote", []interface{}{(*btcjson.SetNetworkStewardVoteResult)(nil)}},
	{"getnetworkstewardvote", []interface{}{(*btcjson.GetNetworkStewardVoteResult)(nil)}},
	{"resync", nil},
	{"stopresync", returnsString},
	{"addp2shscript", returnsString},
	{"dumpprivkey", returnsString},
	{"getbalance", append(returnsNumber, returnsNumber[0])},
	{"getbestblockhash", returnsString},
	{"getblockcount", returnsNumber},
	{"getinfo", []interface{}{(*btcjson.InfoWalletResult)(nil)}},
	{"getnewaddress", returnsString},
	{"getreceivedbyaddress", returnsNumber},
	{"gettransaction", []interface{}{(*btcjson.GetTransactionResult)(nil)}},
	{"getwalletseed", returnsString},
	{"getsecret", returnsString},
	{"help", append(returnsString, returnsString[0])},
	{"importprivkey", nil},
	{"listlockunspent", []interface{}{(*[]btcjson.TransactionInput)(nil)}},
	{"listreceivedbyaddress", []interface{}{(*[]btcjson.ListReceivedByAddressResult)(nil)}},
	{"listsinceblock", []interface{}{(*btcjson.ListSinceBlockResult)(nil)}},
	{"listtransactions", returnsLTRArray},
	{"listunspent", []interface{}{(*btcjson.ListUnspentResult)(nil)}},
	{"lockunspent", returnsBool},
	{"sendfrom", returnsString},
	{"sendmany", returnsString},
	{"sendtoaddress", returnsString},
	{"sendvote", returnsString},
	{"settxfee", returnsBool},
	{"signmessage", returnsString},
	{"signrawtransaction", []interface{}{(*btcjson.SignRawTransactionResult)(nil)}},
	{"validateaddress", []interface{}{(*btcjson.ValidateAddressWalletResult)(nil)}},
	{"verifymessage", returnsBool},
	{"walletlock", nil},
	{"walletpassphrase", nil},
	{"walletpassphrasechange", nil},
	{"walletmempool", []interface{}{(*btcjson.WalletMempoolRes)(nil)}},
	{"exportwatchingwallet", returnsString},
	{"getbestblock", []interface{}{(*btcjson.GetBestBlockResult)(nil)}},
	{"getunconfirmedbalance", returnsNumber},
	{"listaddresstransactions", returnsLTRArray},
	{"listalltransactions", returnsLTRArray},
	{"walletislocked", returnsBool},
}

// HelpDescs contains the locale-specific help strings along with the locale.
var HelpDescs = []struct {
	Locale   string // Actual locale, e.g. en_US
	GoLocale string // Locale used in Go names, e.g. EnUS
	Descs    map[string]string
}{
	{"en_US", "EnUS", helpDescsEnUS}, // helpdescs_en_US.go
}
