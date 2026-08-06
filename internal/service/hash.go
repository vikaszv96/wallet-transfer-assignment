package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// requestHash fingerprints the parts of a transfer request that must not
// change between retries sharing the same idempotency key. If a client sends
// a different payload under a key it already used, that's a client bug, not
// a legitimate retry — the hash mismatch is what catches it.
func requestHash(in CreateTransferInput) string {
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s|%s|%d", in.FromWalletID, in.ToWalletID, in.Amount)))
	return hex.EncodeToString(sum[:])
}
