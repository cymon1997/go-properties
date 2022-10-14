package usecase

import (
	"context"

	"github.com/cymon1997/go-properties/core/inbound/http/spec"
)

func (u *usecaseImpl) Version(_ context.Context) (*spec.VersionResponse, error) {
	return &spec.VersionResponse{
		Version: "1.0.0",
	}, nil
}
