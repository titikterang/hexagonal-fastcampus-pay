package common

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
)

func ExtractUserIDFromHeader(ctx context.Context) (string, bool) {
	tr, _ := transport.FromServerContext(ctx)
	header := tr.RequestHeader()
	uID := header.Get("X-User")
	if uID == "" {
		return "", false
	}

	return uID, true
}
