package irsem

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/enum"
)

func Verify(mod *ir.Module) error {
	for _, global := range mod.Globals {
		if global.Init == nil && global.Linkage == enum.LinkageNone {
			return fmt.Errorf("global variables without init and have no linkage")
		}
	}

	return nil
}
