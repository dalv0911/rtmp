package rtmp

import "github.com/dalv0911/rtmp/server"

// NewServer returns a new RMTP server, listening on the given bind string. If
// an error was encountered in binding to that address, it will be returned,
// along with a nil server.
//
// For more information, see package github.com/dalv0911/rtmp/server.
func NewServer(bind string) (*server.Server, error) {
	return server.New(bind)
}
