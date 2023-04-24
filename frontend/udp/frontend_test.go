package udp_test

import (
	"testing"

	"github.com/whr819987540/chihaya/frontend/udp"
	"github.com/whr819987540/chihaya/middleware"
	"github.com/whr819987540/chihaya/storage"
	_ "github.com/whr819987540/chihaya/storage/memory"
)

func TestStartStopRaceIssue437(t *testing.T) {
	ps, err := storage.NewPeerStore("memory", nil)
	if err != nil {
		t.Fatal(err)
	}
	var responseConfig middleware.ResponseConfig
	lgc := middleware.NewLogic(responseConfig, ps, nil, nil)
	fe, err := udp.NewFrontend(lgc, udp.Config{Addr: "127.0.0.1:0"})
	if err != nil {
		t.Fatal(err)
	}
	errC := fe.Stop()
	if errs := <-errC; len(errs) != 0 {
		t.Fatal(errs[0])
	}
}
