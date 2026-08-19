package pool

import (
	"context"
	"net"
	"sync"
)

// ... existing code ...

func (p *ConnPool) newConn(ctx context.Context, pooled bool) (*Conn, error) {
	cn, err := p.dial(ctx)
	if err != nil {
		return nil, err
	}

	if p.OnConnect != nil {
		if err := p.OnConnect(ctx, cn); err != nil {
			_ = cn.Close()
			return nil, err
		}
	}

	return cn, nil
}

// ... existing code ...