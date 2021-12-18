package irsem

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func Verify(mod *ir.Module) error {
	for _, global := range mod.Globals {
		_, ok := global.ContentType.(*types.PointerType)
		if !ok {
			return fmt.Errorf("global variables cannot use non-pointer type as content type")
		}
	}

	return nil
}
