package im

import "github.com/lphhh/douyin/open/context"

// IM .
type IM struct {
	*context.Context
}

// NewIM .
func NewIM(context *context.Context) *IM {
	IM := new(IM)
	IM.Context = context
	return IM
}
