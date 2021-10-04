/**
  @File：osGetenv_test.go
  @Author：TT
  @Time：2021/10/4 21:09
*/
package test

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	home := os.Getenv("VERSION")

	fmt.Println(home)
}
