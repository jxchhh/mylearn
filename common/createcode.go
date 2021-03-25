package common

import (
"fmt"
"math/rand"
"time"
)

func Createcode()  string{
	return fmt.Sprintf("%04v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000))
}

