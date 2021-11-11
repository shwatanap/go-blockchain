package repository

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-blockchain/domain/entity"
	"go-blockchain/domain/repository"
)

type blockRepository struct{}

func NewBlockRepository() repository.BlockRepository {
	return &blockRepository{}
}

func (br *blockRepository) PreviousHash(b *entity.Block) [32]byte {
	return b.PreviousHash
}

func (br *blockRepository) Nonce(b *entity.Block) int {
	return b.Nonce
}

func (br *blockRepository) Transaction(b *entity.Block) []*entity.Transaction {
	return b.Transactions
}

func (br *blockRepository) Print(b *entity.Block) {
	fmt.Printf("timestamp       %d\n", b.Timestamp)
	fmt.Printf("nonce           %d\n", b.Nonce)
	fmt.Printf("previous_hash   %x\n", b.PreviousHash)
	// for _, t := range b.Transactions {
	// 	// repository.TransactionRepository.Print(t)
	// }
}

func (br *blockRepository) Hash(b *entity.Block) [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (br *blockRepository) MarshalJSON(b *entity.Block) ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64                 `json:"timestamp"`
		Nonce        int                   `json:"nonce"`
		PreviousHash string                `json:"previous_hash"`
		Transactions []*entity.Transaction `json:"transactions"`
	}{
		Timestamp:    b.Timestamp,
		Nonce:        b.Nonce,
		PreviousHash: fmt.Sprintf("%x", b.PreviousHash),
		Transactions: b.Transactions,
	})
}

func (br *blockRepository) UnmarshalJSON(b *entity.Block, data []byte) error {
	var previousHash string
	v := &struct {
		Timestamp    *int64                 `json:"timestamp"`
		Nonce        *int                   `json:"nonce"`
		PreviousHash *string                `json:"previous_hash"`
		Transactions *[]*entity.Transaction `json:"transactions"`
	}{
		Timestamp:    &b.Timestamp,
		Nonce:        &b.Nonce,
		PreviousHash: &previousHash,
		Transactions: &b.Transactions,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	ph, _ := hex.DecodeString(*v.PreviousHash)
	copy(b.PreviousHash[:], ph[:32])
	return nil
}
