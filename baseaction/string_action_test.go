package baseaction

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShowConstString(t *testing.T) {
	ShowConstString()
	fmt.Println("Success!")
}

func TestShowForEach(t *testing.T) {
}

func TestUpdateStr(t *testing.T) {
	assert.Equal(t, "zh*ngsan", UpdateStr1("*"))
	assert.Equal(t, "李三", UpdateStr2("三"))
}
