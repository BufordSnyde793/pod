// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	context "context"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	mock "github.com/stretchr/testify/mock"

	multiaddr "github.com/multiformats/go-multiaddr"

	network "github.com/libp2p/go-libp2p-core/network"

	peer "github.com/libp2p/go-libp2p-core/peer"
)

// ConnMock is an autogenerated mock type for the Conn type
type ConnMock struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ConnMock) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetStreams provides a mock function with given fields:
func (_m *ConnMock) GetStreams() []network.Stream {
	ret := _m.Called()

	var r0 []network.Stream
	if rf, ok := ret.Get(0).(func() []network.Stream); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]network.Stream)
		}
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *ConnMock) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LocalMultiaddr provides a mock function with given fields:
func (_m *ConnMock) LocalMultiaddr() multiaddr.Multiaddr {
	ret := _m.Called()

	var r0 multiaddr.Multiaddr
	if rf, ok := ret.Get(0).(func() multiaddr.Multiaddr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(multiaddr.Multiaddr)
		}
	}

	return r0
}

// LocalPeer provides a mock function with given fields:
func (_m *ConnMock) LocalPeer() peer.ID {
	ret := _m.Called()

	var r0 peer.ID
	if rf, ok := ret.Get(0).(func() peer.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(peer.ID)
	}

	return r0
}

// LocalPrivateKey provides a mock function with given fields:
func (_m *ConnMock) LocalPrivateKey() crypto.PrivKey {
	ret := _m.Called()

	var r0 crypto.PrivKey
	if rf, ok := ret.Get(0).(func() crypto.PrivKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivKey)
		}
	}

	return r0
}

// NewStream provides a mock function with given fields: _a0
func (_m *ConnMock) NewStream(_a0 context.Context) (network.Stream, error) {
	ret := _m.Called(_a0)

	var r0 network.Stream
	if rf, ok := ret.Get(0).(func(context.Context) network.Stream); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(network.Stream)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteMultiaddr provides a mock function with given fields:
func (_m *ConnMock) RemoteMultiaddr() multiaddr.Multiaddr {
	ret := _m.Called()

	var r0 multiaddr.Multiaddr
	if rf, ok := ret.Get(0).(func() multiaddr.Multiaddr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(multiaddr.Multiaddr)
		}
	}

	return r0
}

// RemotePeer provides a mock function with given fields:
func (_m *ConnMock) RemotePeer() peer.ID {
	ret := _m.Called()

	var r0 peer.ID
	if rf, ok := ret.Get(0).(func() peer.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(peer.ID)
	}

	return r0
}

// RemotePublicKey provides a mock function with given fields:
func (_m *ConnMock) RemotePublicKey() crypto.PubKey {
	ret := _m.Called()

	var r0 crypto.PubKey
	if rf, ok := ret.Get(0).(func() crypto.PubKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PubKey)
		}
	}

	return r0
}

// Scope provides a mock function with given fields:
func (_m *ConnMock) Scope() network.ConnScope {
	ret := _m.Called()

	var r0 network.ConnScope
	if rf, ok := ret.Get(0).(func() network.ConnScope); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(network.ConnScope)
		}
	}

	return r0
}

// Stat provides a mock function with given fields:
func (_m *ConnMock) Stat() network.ConnStats {
	ret := _m.Called()

	var r0 network.ConnStats
	if rf, ok := ret.Get(0).(func() network.ConnStats); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(network.ConnStats)
	}

	return r0
}

type NewConnMockT interface {
	mock.TestingT
	Cleanup(func())
}

// NewConnMock creates a new instance of ConnMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConnMock(t NewConnMockT) *ConnMock {
	mock := &ConnMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}