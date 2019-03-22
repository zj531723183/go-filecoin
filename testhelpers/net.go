package testhelpers

import (
	"context"
	"time"

	ma "gx/ipfs/QmNTCey11oxhb1AxDnQBRHtdhap6Ctud872NjAYPYYXPuc/go-multiaddr"
	"gx/ipfs/QmRhFARzTHcFh8wUxwN5KvyTGq73FLC65EfFAhz8Ng7aGb/go-libp2p-peerstore"
	pstore "gx/ipfs/QmRhFARzTHcFh8wUxwN5KvyTGq73FLC65EfFAhz8Ng7aGb/go-libp2p-peerstore"
	inet "gx/ipfs/QmTGxDz2CjBucFzPNTiWwzQmTWdrBnzqbqrMucDYMsjuPb/go-libp2p-net"
	"gx/ipfs/QmTu65MVbemtUxJEWgsTtzv9Zv9P8rvmqNA4eG9TrTRGYc/go-libp2p-peer"
	smux "gx/ipfs/QmVtV1y2e8W4eQgzsP6qfSpCCZ6zWYE4m6NzJjB7iswwrT/go-stream-muxer"
	"gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	msmux "gx/ipfs/QmabLh8TrJ3emfAoQk5AbqbLTbMyj7XqumMFmAFxa9epo8/go-multistream"
	"gx/ipfs/QmcCk4LZRJPAKuwY9dusFea7LckELZgo5HagErTbm39o38/go-libp2p-interface-connmgr"
	"gx/ipfs/Qmd52WKRSwrBK5gUaJKawryZQ5by6UbNB8KVW2Zy6JtbyW/go-libp2p-host"
)

var _ host.Host = &FakeHost{}

// FakeHost is a test host.Host
type FakeHost struct {
	ConnectImpl func(context.Context, pstore.PeerInfo) error
}

// NewFakeHost constructs a FakeHost with no other parameters needed
func NewFakeHost() host.Host {
	nopfunc := func(_ context.Context, _ peerstore.PeerInfo) error { return nil }
	return &FakeHost{ConnectImpl: nopfunc}
}

// minimal implementation of host.Host interface

func (fh *FakeHost) Addrs() []ma.Multiaddr              { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Close() error                       { panic("not implemented") } // nolint: golint
func (fh *FakeHost) ConnManager() ifconnmgr.ConnManager { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Connect(ctx context.Context, pi pstore.PeerInfo) error { // nolint: golint
	return fh.ConnectImpl(ctx, pi)
}
func (fh *FakeHost) ID() peer.ID                                      { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Network() inet.Network                            { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Mux() *msmux.MultistreamMuxer                     { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Peerstore() pstore.Peerstore                      { panic("not implemented") } // nolint: golint
func (fh *FakeHost) RemoveStreamHandler(protocol.ID)                  { panic("not implemented") } // nolint: golint
func (fh *FakeHost) SetStreamHandler(protocol.ID, inet.StreamHandler) { panic("not implemented") } // nolint: golint
func (fh *FakeHost) SetStreamHandlerMatch(protocol.ID, func(string) bool, inet.StreamHandler) { // nolint: golint
	panic("not implemented")
}

// NewStream is required for the host.Host interface; returns a new FakeStream.
func (fh *FakeHost) NewStream(context.Context, peer.ID, ...protocol.ID) (inet.Stream, error) { // nolint: golint
	return newFakeStream(), nil
}

var _ inet.Dialer = &FakeDialer{}

// FakeDialer is a test inet.Dialer
type FakeDialer struct {
	PeersImpl func() []peer.ID
}

// Minimal implementation of the inet.Dialer interface

// Peers returns a fake inet.Dialer PeersImpl
func (fd *FakeDialer) Peers() []peer.ID {
	return fd.PeersImpl()
}
func (fd *FakeDialer) Peerstore() pstore.Peerstore                          { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) LocalPeer() peer.ID                                   { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) DialPeer(context.Context, peer.ID) (inet.Conn, error) { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) ClosePeer(peer.ID) error                              { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) Connectedness(peer.ID) inet.Connectedness             { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) Conns() []inet.Conn                                   { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) ConnsToPeer(peer.ID) []inet.Conn                      { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) Notify(inet.Notifiee)                                 { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) StopNotify(inet.Notifiee)                             { panic("not implemented") } // nolint: golint

// fakeStream is a test inet.Stream
type fakeStream struct {
	_   smux.Stream
	pid protocol.ID
}

var _ inet.Stream = &fakeStream{}

func newFakeStream() fakeStream { return fakeStream{} }

// Minimal implementation of the inet.Stream interface

func (fs fakeStream) Protocol() protocol.ID              { return fs.pid }            // nolint: golint
func (fs fakeStream) SetProtocol(id protocol.ID)         { fs.pid = id }              // nolint: golint
func (fs fakeStream) Stat() inet.Stat                    { panic("not implemented") } // nolint: golint
func (fs fakeStream) Conn() inet.Conn                    { panic("not implemented") } // nolint: golint
func (fs fakeStream) Write(_ []byte) (int, error)        { return 1, nil }            // nolint: golint
func (fs fakeStream) Read(_ []byte) (int, error)         { return 1, nil }            // nolint: golint
func (fs fakeStream) Close() error                       { return nil }               // nolint: golint
func (fs fakeStream) Reset() error                       { return nil }               // nolint: golint
func (fs fakeStream) SetDeadline(_ time.Time) error      { return nil }               // nolint: golint
func (fs fakeStream) SetReadDeadline(_ time.Time) error  { return nil }               // nolint: golint
func (fs fakeStream) SetWriteDeadline(_ time.Time) error { return nil }               // nolint: golint
