// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains all the wrappers from the core/types package.

package web3go

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	whisper "github.com/ethereum/go-ethereum/whisper/whisperv6"
)

// A Nonce is a 64-bit hash which proves (combined with the mix-hash) that
// a sufficient amount of computation has been carried out on a block.
type Nonce struct {
	nonce types.BlockNonce
}

// GetBytes retrieves the byte representation of the block nonce.
func (n *Nonce) GetBytes() []byte {
	return n.nonce[:]
}

// GetHex retrieves the hex string representation of the block nonce.
func (n *Nonce) GetHex() string {
	return fmt.Sprintf("0x%x", n.nonce[:])
}

// Bloom represents a 256 bit bloom filter.
type Bloom struct {
	bloom types.Bloom
}

// GetBytes retrieves the byte representation of the bloom filter.
func (b *Bloom) GetBytes() []byte {
	return b.bloom[:]
}

// GetHex retrieves the hex string representation of the bloom filter.
func (b *Bloom) GetHex() string {
	return fmt.Sprintf("0x%x", b.bloom[:])
}

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	header *types.Header
}

// NewHeaderFromRLP parses a header from an RLP data dump.
func NewHeaderFromRLP(data []byte) (*Header, error) {
	h := &Header{
		header: new(types.Header),
	}
	if err := rlp.DecodeBytes(common.CopyBytes(data), h.header); err != nil {
		return nil, err
	}
	return h, nil
}

// EncodeRLP encodes a header into an RLP data dump.
func (h *Header) EncodeRLP() ([]byte, error) {
	return rlp.EncodeToBytes(h.header)
}

// NewHeaderFromJSON parses a header from a JSON data dump.
func NewHeaderFromJSON(data string) (*Header, error) {
	h := &Header{
		header: new(types.Header),
	}
	if err := json.Unmarshal([]byte(data), h.header); err != nil {
		return nil, err
	}
	return h, nil
}

// EncodeJSON encodes a header into a JSON data dump.
func (h *Header) EncodeJSON() (string, error) {
	data, err := json.Marshal(h.header)
	return string(data), err
}

// GetParentHash ...
func (h *Header) GetParentHash() *Hash { return &Hash{h.header.ParentHash} }

// GetUncleHash ...
func (h *Header) GetUncleHash() *Hash { return &Hash{h.header.UncleHash} }

// GetCoinbase ...
func (h *Header) GetCoinbase() *Address { return &Address{h.header.Coinbase} }

// GetRoot ...
func (h *Header) GetRoot() *Hash { return &Hash{h.header.Root} }

// GetTxHash ...
func (h *Header) GetTxHash() *Hash { return &Hash{h.header.TxHash} }

// GetReceiptHash ...
func (h *Header) GetReceiptHash() *Hash { return &Hash{h.header.ReceiptHash} }

// GetBloom ...
func (h *Header) GetBloom() *Bloom { return &Bloom{h.header.Bloom} }

// GetDifficulty ...
func (h *Header) GetDifficulty() *BigInt { return &BigInt{h.header.Difficulty} }

// GetNumber ...
func (h *Header) GetNumber() int64 { return h.header.Number.Int64() }

// GetGasLimit ...
func (h *Header) GetGasLimit() int64 { return int64(h.header.GasLimit) }

// GetGasUsed ...
func (h *Header) GetGasUsed() int64 { return int64(h.header.GasUsed) }

// GetTime ...
func (h *Header) GetTime() int64 { return int64(h.header.Time) }

// GetExtra ...
func (h *Header) GetExtra() []byte { return h.header.Extra }

// GetMixDigest ...
func (h *Header) GetMixDigest() *Hash { return &Hash{h.header.MixDigest} }

// GetNonce ...
func (h *Header) GetNonce() *Nonce { return &Nonce{h.header.Nonce} }

// GetHash ...
func (h *Header) GetHash() *Hash { return &Hash{h.header.Hash()} }

// Headers represents a slice of headers.
type Headers struct{ headers []*types.Header }

// Size returns the number of headers in the slice.
func (h *Headers) Size() int {
	return len(h.headers)
}

// Get returns the header at the given index from the slice.
func (h *Headers) Get(index int) (header *Header, _ error) {
	if index < 0 || index >= len(h.headers) {
		return nil, errors.New("index out of bounds")
	}
	return &Header{h.headers[index]}, nil
}

// Block represents an entire block in the Ethereum blockchain.
type Block struct {
	block *types.Block
}

// NewBlockFromRLP parses a block from an RLP data dump.
func NewBlockFromRLP(data []byte) (*Block, error) {
	b := &Block{
		block: new(types.Block),
	}
	if err := rlp.DecodeBytes(common.CopyBytes(data), b.block); err != nil {
		return nil, err
	}
	return b, nil
}

// EncodeRLP encodes a block into an RLP data dump.
func (b *Block) EncodeRLP() ([]byte, error) {
	return rlp.EncodeToBytes(b.block)
}

// NewBlockFromJSON parses a block from a JSON data dump.
func NewBlockFromJSON(data string) (*Block, error) {
	b := &Block{
		block: new(types.Block),
	}
	if err := json.Unmarshal([]byte(data), b.block); err != nil {
		return nil, err
	}
	return b, nil
}

// EncodeJSON encodes a block into a JSON data dump.
func (b *Block) EncodeJSON() (string, error) {
	data, err := json.Marshal(b.block)
	return string(data), err
}

// GetParentHash ...
func (b *Block) GetParentHash() *Hash { return &Hash{b.block.ParentHash()} }

// GetUncleHash ...
func (b *Block) GetUncleHash() *Hash { return &Hash{b.block.UncleHash()} }

// GetCoinbase ...
func (b *Block) GetCoinbase() *Address { return &Address{b.block.Coinbase()} }

// GetRoot ...
func (b *Block) GetRoot() *Hash { return &Hash{b.block.Root()} }

// GetTxHash ...
func (b *Block) GetTxHash() *Hash { return &Hash{b.block.TxHash()} }

// GetReceiptHash ...
func (b *Block) GetReceiptHash() *Hash { return &Hash{b.block.ReceiptHash()} }

// GetBloom ...
func (b *Block) GetBloom() *Bloom { return &Bloom{b.block.Bloom()} }

// GetDifficulty ...
func (b *Block) GetDifficulty() *BigInt { return &BigInt{b.block.Difficulty()} }

// GetNumber ...
func (b *Block) GetNumber() int64 { return b.block.Number().Int64() }

// GetGasLimit ...
func (b *Block) GetGasLimit() int64 { return int64(b.block.GasLimit()) }

// GetGasUsed ...
func (b *Block) GetGasUsed() int64 { return int64(b.block.GasUsed()) }

// GetTime ...
func (b *Block) GetTime() int64 { return int64(b.block.Time()) }

// GetExtra ...
func (b *Block) GetExtra() []byte { return b.block.Extra() }

// GetMixDigest ...
func (b *Block) GetMixDigest() *Hash { return &Hash{b.block.MixDigest()} }

// GetNonce ...
func (b *Block) GetNonce() int64 { return int64(b.block.Nonce()) }

// GetHash ...
func (b *Block) GetHash() *Hash { return &Hash{b.block.Hash()} }

// GetHeader ...
func (b *Block) GetHeader() *Header { return &Header{b.block.Header()} }

// GetUncles ...
func (b *Block) GetUncles() *Headers { return &Headers{b.block.Uncles()} }

// GetTransactions ...
func (b *Block) GetTransactions() *Transactions { return &Transactions{b.block.Transactions()} }

// GetTransaction ...
func (b *Block) GetTransaction(hash *Hash) *Transaction {
	return &Transaction{b.block.Transaction(hash.hash)}
}

// Transaction represents a single Ethereum transaction.
type Transaction struct {
	tx *types.Transaction
}

// NewTransaction creates a new transaction with the given properties.
func NewTransaction(nonce int64, to *Address, amount *BigInt, gasLimit int64, gasPrice *BigInt, data []byte) *Transaction {
	return &Transaction{types.NewTransaction(uint64(nonce), to.address, amount.bigint, uint64(gasLimit), gasPrice.bigint, common.CopyBytes(data))}
}

// NewTransactionFromRLP parses a transaction from an RLP data dump.
func NewTransactionFromRLP(data []byte) (*Transaction, error) {
	tx := &Transaction{
		tx: new(types.Transaction),
	}
	if err := rlp.DecodeBytes(common.CopyBytes(data), tx.tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// EncodeRLP encodes a transaction into an RLP data dump.
func (tx *Transaction) EncodeRLP() ([]byte, error) {
	return rlp.EncodeToBytes(tx.tx)
}

// NewTransactionFromJSON parses a transaction from a JSON data dump.
func NewTransactionFromJSON(data string) (*Transaction, error) {
	tx := &Transaction{
		tx: new(types.Transaction),
	}
	if err := json.Unmarshal([]byte(data), tx.tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// EncodeJSON encodes a transaction into a JSON data dump.
func (tx *Transaction) EncodeJSON() (string, error) {
	data, err := json.Marshal(tx.tx)
	return string(data), err
}

// GetData ...
func (tx *Transaction) GetData() []byte { return tx.tx.Data() }

// GetGas ...
func (tx *Transaction) GetGas() int64 { return int64(tx.tx.Gas()) }

// GetGasPrice ...
func (tx *Transaction) GetGasPrice() *BigInt { return &BigInt{tx.tx.GasPrice()} }

// GetValue ...
func (tx *Transaction) GetValue() *BigInt { return &BigInt{tx.tx.Value()} }

// GetNonce ...
func (tx *Transaction) GetNonce() int64 { return int64(tx.tx.Nonce()) }

// GetHash ...
func (tx *Transaction) GetHash() *Hash { return &Hash{tx.tx.Hash()} }

// GetCost ...
func (tx *Transaction) GetCost() *BigInt { return &BigInt{tx.tx.Cost()} }

// GetSigHash ...
// Deprecated: GetSigHash cannot know which signer to use.
func (tx *Transaction) GetSigHash() *Hash { return &Hash{types.HomesteadSigner{}.Hash(tx.tx)} }

// GetFrom ...
// Deprecated: use EthereumClient.TransactionSender
func (tx *Transaction) GetFrom(chainID *BigInt) (address *Address, _ error) {
	var signer types.Signer = types.HomesteadSigner{}
	if chainID != nil {
		signer = types.NewEIP155Signer(chainID.bigint)
	}
	from, err := types.Sender(signer, tx.tx)
	return &Address{from}, err
}

// GetTo ...
func (tx *Transaction) GetTo() *Address {
	if to := tx.tx.To(); to != nil {
		return &Address{*to}
	}
	return nil
}

// WithSignature ...
func (tx *Transaction) WithSignature(sig []byte, chainID *BigInt) (signedTx *Transaction, _ error) {
	var signer types.Signer = types.HomesteadSigner{}
	if chainID != nil {
		signer = types.NewEIP155Signer(chainID.bigint)
	}
	rawTx, err := tx.tx.WithSignature(signer, common.CopyBytes(sig))
	return &Transaction{rawTx}, err
}

// Transactions represents a slice of transactions.
type Transactions struct{ txs types.Transactions }

// Size returns the number of transactions in the slice.
func (txs *Transactions) Size() int {
	return len(txs.txs)
}

// Get returns the transaction at the given index from the slice.
func (txs *Transactions) Get(index int) (tx *Transaction, _ error) {
	if index < 0 || index >= len(txs.txs) {
		return nil, errors.New("index out of bounds")
	}
	return &Transaction{txs.txs[index]}, nil
}

// Receipt represents the results of a transaction.
type Receipt struct {
	receipt *types.Receipt
}

// NewReceiptFromRLP parses a transaction receipt from an RLP data dump.
func NewReceiptFromRLP(data []byte) (*Receipt, error) {
	r := &Receipt{
		receipt: new(types.Receipt),
	}
	if err := rlp.DecodeBytes(common.CopyBytes(data), r.receipt); err != nil {
		return nil, err
	}
	return r, nil
}

// EncodeRLP encodes a transaction receipt into an RLP data dump.
func (r *Receipt) EncodeRLP() ([]byte, error) {
	return rlp.EncodeToBytes(r.receipt)
}

// NewReceiptFromJSON parses a transaction receipt from a JSON data dump.
func NewReceiptFromJSON(data string) (*Receipt, error) {
	r := &Receipt{
		receipt: new(types.Receipt),
	}
	if err := json.Unmarshal([]byte(data), r.receipt); err != nil {
		return nil, err
	}
	return r, nil
}

// EncodeJSON encodes a transaction receipt into a JSON data dump.
func (r *Receipt) EncodeJSON() (string, error) {
	data, err := rlp.EncodeToBytes(r.receipt)
	return string(data), err
}

// GetStatus ...
func (r *Receipt) GetStatus() int { return int(r.receipt.Status) }

// GetPostState ...
func (r *Receipt) GetPostState() []byte { return r.receipt.PostState }

// GetCumulativeGasUsed ...
func (r *Receipt) GetCumulativeGasUsed() int64 { return int64(r.receipt.CumulativeGasUsed) }

// GetBloom ...
func (r *Receipt) GetBloom() *Bloom { return &Bloom{r.receipt.Bloom} }

// GetLogs ...
func (r *Receipt) GetLogs() *Logs { return &Logs{r.receipt.Logs} }

// GetTxHash ...
func (r *Receipt) GetTxHash() *Hash { return &Hash{r.receipt.TxHash} }

// GetContractAddress ...
func (r *Receipt) GetContractAddress() *Address { return &Address{r.receipt.ContractAddress} }

// GetGasUsed ...
func (r *Receipt) GetGasUsed() int64 { return int64(r.receipt.GasUsed) }

// Info represents a diagnostic information about the whisper node.
type Info struct {
	info *whisper.Info
}

// NewMessage represents a new whisper message that is posted through the RPC.
type NewMessage struct {
	newMessage *whisper.NewMessage
}

// NewNewMessage ...
func NewNewMessage() *NewMessage {
	nm := &NewMessage{
		newMessage: new(whisper.NewMessage),
	}
	return nm
}

// GetSymKeyID ...
func (nm *NewMessage) GetSymKeyID() string { return nm.newMessage.SymKeyID }

// SetSymKeyID ...
func (nm *NewMessage) SetSymKeyID(symKeyID string) { nm.newMessage.SymKeyID = symKeyID }

// GetPublicKey ...
func (nm *NewMessage) GetPublicKey() []byte { return nm.newMessage.PublicKey }

// SetPublicKey ...
func (nm *NewMessage) SetPublicKey(publicKey []byte) {
	nm.newMessage.PublicKey = common.CopyBytes(publicKey)
}

// GetSig ...
func (nm *NewMessage) GetSig() string { return nm.newMessage.Sig }

// SetSig ...
func (nm *NewMessage) SetSig(sig string) { nm.newMessage.Sig = sig }

// GetTTL ...
func (nm *NewMessage) GetTTL() int64 { return int64(nm.newMessage.TTL) }

// SetTTL ...
func (nm *NewMessage) SetTTL(ttl int64) { nm.newMessage.TTL = uint32(ttl) }

// GetPayload ...
func (nm *NewMessage) GetPayload() []byte { return nm.newMessage.Payload }

// SetPayload ...
func (nm *NewMessage) SetPayload(payload []byte) { nm.newMessage.Payload = common.CopyBytes(payload) }

// GetPowTime ...
func (nm *NewMessage) GetPowTime() int64 { return int64(nm.newMessage.PowTime) }

// SetPowTime ...
func (nm *NewMessage) SetPowTime(powTime int64) { nm.newMessage.PowTime = uint32(powTime) }

// GetPowTarget ...
func (nm *NewMessage) GetPowTarget() float64 { return nm.newMessage.PowTarget }

// SetPowTarget ...
func (nm *NewMessage) SetPowTarget(powTarget float64) { nm.newMessage.PowTarget = powTarget }

// GetTargetPeer ...
func (nm *NewMessage) GetTargetPeer() string { return nm.newMessage.TargetPeer }

// SetTargetPeer ...
func (nm *NewMessage) SetTargetPeer(targetPeer string) { nm.newMessage.TargetPeer = targetPeer }

// GetTopic ...
func (nm *NewMessage) GetTopic() []byte { return nm.newMessage.Topic[:] }

// SetTopic ...
func (nm *NewMessage) SetTopic(topic []byte) { nm.newMessage.Topic = whisper.BytesToTopic(topic) }

// Message represents a whisper message.
type Message struct {
	message *whisper.Message
}

// GetSig ...
func (m *Message) GetSig() []byte { return m.message.Sig }

// GetTTL ...
func (m *Message) GetTTL() int64 { return int64(m.message.TTL) }

// GetTimestamp ...
func (m *Message) GetTimestamp() int64 { return int64(m.message.Timestamp) }

// GetPayload ...
func (m *Message) GetPayload() []byte { return m.message.Payload }

// GetPoW ...
func (m *Message) GetPoW() float64 { return m.message.PoW }

// GetHash ...
func (m *Message) GetHash() []byte { return m.message.Hash }

// GetDst ...
func (m *Message) GetDst() []byte { return m.message.Dst }

// Messages represents an array of messages.
type Messages struct {
	messages []*whisper.Message
}

// Size returns the number of messages in the slice.
func (m *Messages) Size() int {
	return len(m.messages)
}

// Get returns the message at the given index from the slice.
func (m *Messages) Get(index int) (message *Message, _ error) {
	if index < 0 || index >= len(m.messages) {
		return nil, errors.New("index out of bounds")
	}
	return &Message{m.messages[index]}, nil
}

// Criteria holds various filter options for inbound messages.
type Criteria struct {
	criteria *whisper.Criteria
}

// NewCriteria ...
func NewCriteria(topic []byte) *Criteria {
	c := &Criteria{
		criteria: new(whisper.Criteria),
	}
	encodedTopic := whisper.BytesToTopic(topic)
	c.criteria.Topics = []whisper.TopicType{encodedTopic}
	return c
}

// GetSymKeyID ...
func (c *Criteria) GetSymKeyID() string { return c.criteria.SymKeyID }

// SetSymKeyID ...
func (c *Criteria) SetSymKeyID(symKeyID string) { c.criteria.SymKeyID = symKeyID }

// GetPrivateKeyID ...
func (c *Criteria) GetPrivateKeyID() string { return c.criteria.PrivateKeyID }

// SetPrivateKeyID ...
func (c *Criteria) SetPrivateKeyID(privateKeyID string) { c.criteria.PrivateKeyID = privateKeyID }

// GetSig ...
func (c *Criteria) GetSig() []byte { return c.criteria.Sig }

// SetSig ...
func (c *Criteria) SetSig(sig []byte) { c.criteria.Sig = common.CopyBytes(sig) }

// GetMinPow ...
func (c *Criteria) GetMinPow() float64 { return c.criteria.MinPow }

// SetMinPow ...
func (c *Criteria) SetMinPow(pow float64) { c.criteria.MinPow = pow }
