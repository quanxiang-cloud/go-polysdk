package polysdk

import (
	"net"
	"net/http"
	"time"

	"github.com/quanxiang-cloud/go-polysdk/internal/signature"

	"github.com/quanxiang-cloud/go-polysdk/internal/config"
)

const (
	httpTimeout      = 10
	httpMaxIdleConns = 3
)

// NewPolyClient create a ploy client from config file
func NewPolyClient(cfg *config.PolyConfig) (*PolyClient, error) {
	sign, err := signature.NewSignerFromKey(cfg.Key.SecretKey)
	if err != nil {
		return nil, err
	}

	r := &PolyClient{
		remoteURL:   cfg.RemoteURL,
		accessKeyID: cfg.Key.AccessKeyID,
		sign:        sign,
		httpClient: http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					deadline := time.Now().Add(time.Second * httpTimeout)
					c, err := net.DialTimeout(netw, addr, time.Second*httpTimeout)
					if err != nil {
						return nil, err
					}
					c.SetDeadline(deadline)
					return c, nil
				},
				MaxIdleConns:      httpMaxIdleConns,
				DisableKeepAlives: false,
			},
		},
	}
	return r, nil
}

// PolyClient is a client for polyapi
type PolyClient struct {
	timeAdjust  int64 // adjust time clock with server
	remoteURL   string
	accessKeyID string
	sign        signature.Signer
	httpClient  http.Client
}
