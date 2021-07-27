package memballast

import (
	"os"
	"strconv"
)

const env = "GO_MEMBALLAST"

var ballast []byte

func init() {
	s, ok := os.LookupEnv(env)
	if !ok {
		return
	}
  
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
  
	ballast = make([]byte, i)
}
