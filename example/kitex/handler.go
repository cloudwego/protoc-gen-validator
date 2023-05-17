package main

import (
	psm "a/b/c/kitex_gen/psm"
	"context"
)

// ValidatorImpl implements the last service interface defined in the IDL.
type ValidatorImpl struct{}

// Method1 implements the ValidatorImpl interface.
func (s *ValidatorImpl) Method1(ctx context.Context, req *psm.IntValidate) (resp *psm.IntValidate, err error) {
	// TODO: Your code here...
	return
}
