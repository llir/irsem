package irsem

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/enum"
)

func Verify(mod *ir.Module) error {
	for _, global := range mod.Globals {
		if global.Init == nil && global.Linkage == enum.LinkageNone {
			return fmt.Errorf("global variable %q requires either initializer or (external) linkage", global.Name())
		}
	}

	return nil
}
