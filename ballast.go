package memballast

import (
	"os"
	"strconv"
)

const envBytes = "MEMBALLAST_BYTE"

var ballast []byte

func init() {
	s, ok := os.LookupEnv(envBytes)
	if !ok {
		return
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	ballast = make([]byte, i)
}
