package irsem

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyGlobal(t *testing.T) {
	mod := ir.NewModule()
	mod.NewGlobal("bad global with non-pointer type", types.I8)

	assert.NotNil(t, Verify(mod))
}
