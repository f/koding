package kitepinger

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/koding/kite"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKitePinger(tt *testing.T) {
	Convey("It should implement Pinger", tt, func() {
		var _ Pinger = &KitePinger{}
	})
}

func TestKitePingerPing(tt *testing.T) {
	Convey("Given a KitePinger", tt, func() {
		r := kite.New("remote", "0.0.0")
		r.Config.DisableAuthentication = true
		ts := httptest.NewServer(r)

		var shouldPingError bool
		r.HandleFunc("kite.ping", func(req *kite.Request) (interface{}, error) {
			if shouldPingError {
				return nil, errors.New("shouldPingError=true")
			}

			return "pong", nil
		})

		Convey("When it can communicate with remote", func() {
			l := kite.New("local", "0.0.0").NewClient(fmt.Sprintf("%s/kite", ts.URL))
			So(l.Dial(), ShouldBeNil)
			p := NewKitePinger(l)

			Convey("It should return success", func() {
				So(p.Ping(), ShouldEqual, Success)
			})
		})

		Convey("When the kite is not dialed", func() {
			l := kite.New("local", "0.0.0").NewClient(fmt.Sprintf("%s/kite", ts.URL))
			p := NewKitePinger(l)

			Convey("It should return failure", func() {
				So(p.Ping(), ShouldEqual, Failure)
			})
		})

		Convey("When the remote kite returns an error", func() {
			shouldPingError = true
			l := kite.New("local", "0.0.0").NewClient(fmt.Sprintf("%s/kite", ts.URL))
			So(l.Dial(), ShouldBeNil)
			p := NewKitePinger(l)

			Convey("It should return failure", func() {
				So(p.Ping(), ShouldEqual, Failure)
			})
		})

		Convey("When the kite url is not responding", func() {
			l := kite.New("local", "0.0.0").NewClient(fmt.Sprintf("%s/kite", ts.URL))
			So(l.Dial(), ShouldBeNil)
			ts.CloseClientConnections()
			ts.Close()
			p := NewKitePinger(l)

			Convey("It should return failure", func() {
				So(p.Ping(), ShouldEqual, Failure)
			})
		})
	})
}
