package types

import (
	"crypto/sha256"
	"fmt"
	"io"
	"time"
)

func Hash(values ...interface{}) string {
	hash := sha256.New224()

	for _, value := range values {
		io.WriteString(hash, fmt.Sprintf("%s", value))
	}

	io.WriteString(hash, fmt.Sprintf("timestamp%d", time.Now().UnixNano()))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
