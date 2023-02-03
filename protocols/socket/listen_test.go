package socket

import (
	"net"
	"testing"
)

func TestListenTCP(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	if ln.Addr().Network() != "tcp" {
		t.Fatalf("expected network 'tcp', got %q", ln.Addr().Network())
	}
	if ln.Addr().String() == "" {
		t.Fatal("expected non-empty address")
	}
}

func TestListenUDP(t *testing.T) {
	ln, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	defer ln.Close()

	if ln.LocalAddr().Network() != "udp" {
		t.Fatalf("expected network 'udp', got %q", ln.LocalAddr().Network())
	}
	if ln.LocalAddr().String() == "" {
		t.Fatal("expected non-empty address")
	}
}
