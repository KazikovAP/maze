package srandom

import (
	"crypto/rand"
	"math/big"
)

func Intn(maxVal int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(maxVal)))
	if err != nil {
		return 0
	}

	return int(nBig.Int64())
}
