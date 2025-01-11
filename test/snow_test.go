package test

import (
	"fmt"
	"testing"
	"ticket/util"
)

func TestSnow(t *testing.T) {
	id := util.GetSnowflakeId()
	fmt.Println(id)
}
