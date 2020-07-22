package logic

import (
	"context"
)

// PublishInterface interface
type PublishInterface interface {
}

// Publish struct
type Publish struct {
	Context context.Context
	Self    PublishInterface
}

// NewPublishLogic function for create Publish logic
func NewPublishLogic(ctx context.Context) PublishInterface {
	x := Publish{Context: ctx}
	x.Self = &x
	return &x
}
