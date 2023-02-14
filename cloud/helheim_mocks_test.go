// Code generated by git.sr.ht/~nelsam/hel/v4. DO NOT EDIT.
//
// This file contains mocks generated by hel. Do not edit this code by
// hand unless you *really* know what you're doing. Expect any changes
// made manually to be overwritten the next time hel regenerates this
// file.

package cloud_test

import (
	"context"
	"time"

	"git.sr.ht/~nelsam/hel/v4/vegr"
	"github.com/earthly/cloud-api/logstream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type mockGRPCConn struct {
	t            vegr.T
	timeout      time.Duration
	InvokeCalled chan bool
	InvokeInput  struct {
		Ctx         chan context.Context
		Method      chan string
		Args, Reply chan any
		Opts        chan []grpc.CallOption
	}
	InvokeOutput struct {
		Ret0 chan error
	}
	NewStreamCalled chan bool
	NewStreamInput  struct {
		Ctx    chan context.Context
		Desc   chan *grpc.StreamDesc
		Method chan string
		Opts   chan []grpc.CallOption
	}
	NewStreamOutput struct {
		Ret0 chan grpc.ClientStream
		Ret1 chan error
	}
}

func newMockGRPCConn(t vegr.T, timeout time.Duration) *mockGRPCConn {
	m := &mockGRPCConn{t: t, timeout: timeout}
	m.InvokeCalled = make(chan bool, 100)
	m.InvokeInput.Ctx = make(chan context.Context, 100)
	m.InvokeInput.Method = make(chan string, 100)
	m.InvokeInput.Args = make(chan any, 100)
	m.InvokeInput.Reply = make(chan any, 100)
	m.InvokeInput.Opts = make(chan []grpc.CallOption, 100)
	m.InvokeOutput.Ret0 = make(chan error, 100)
	m.NewStreamCalled = make(chan bool, 100)
	m.NewStreamInput.Ctx = make(chan context.Context, 100)
	m.NewStreamInput.Desc = make(chan *grpc.StreamDesc, 100)
	m.NewStreamInput.Method = make(chan string, 100)
	m.NewStreamInput.Opts = make(chan []grpc.CallOption, 100)
	m.NewStreamOutput.Ret0 = make(chan grpc.ClientStream, 100)
	m.NewStreamOutput.Ret1 = make(chan error, 100)
	return m
}
func (m *mockGRPCConn) Invoke(ctx context.Context, method string, args, reply any, opts ...grpc.CallOption) (ret0 error) {
	m.t.Helper()
	m.InvokeCalled <- true
	m.InvokeInput.Ctx <- ctx
	m.InvokeInput.Method <- method
	m.InvokeInput.Args <- args
	m.InvokeInput.Reply <- reply
	m.InvokeInput.Opts <- opts
	vegr.PopulateReturns(m.t, "Invoke", m.timeout, m.InvokeOutput, &ret0)
	return ret0
}
func (m *mockGRPCConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (ret0 grpc.ClientStream, ret1 error) {
	m.t.Helper()
	m.NewStreamCalled <- true
	m.NewStreamInput.Ctx <- ctx
	m.NewStreamInput.Desc <- desc
	m.NewStreamInput.Method <- method
	m.NewStreamInput.Opts <- opts
	vegr.PopulateReturns(m.t, "NewStream", m.timeout, m.NewStreamOutput, &ret0, &ret1)
	return ret0, ret1
}

type mockDeltas struct {
	t          vegr.T
	timeout    time.Duration
	NextCalled chan bool
	NextInput  struct {
		Ctx chan context.Context
	}
	NextOutput struct {
		Ret0 chan []*logstream.Delta
		Ret1 chan error
	}
}

func newMockDeltas(t vegr.T, timeout time.Duration) *mockDeltas {
	m := &mockDeltas{t: t, timeout: timeout}
	m.NextCalled = make(chan bool, 100)
	m.NextInput.Ctx = make(chan context.Context, 100)
	m.NextOutput.Ret0 = make(chan []*logstream.Delta, 100)
	m.NextOutput.Ret1 = make(chan error, 100)
	return m
}
func (m *mockDeltas) Next(ctx context.Context) (ret0 []*logstream.Delta, ret1 error) {
	m.t.Helper()
	m.NextCalled <- true
	m.NextInput.Ctx <- ctx
	vegr.PopulateReturns(m.t, "Next", m.timeout, m.NextOutput, &ret0, &ret1)
	return ret0, ret1
}

type mockContext struct {
	t              vegr.T
	timeout        time.Duration
	DeadlineCalled chan bool
	DeadlineOutput struct {
		Deadline chan time.Time
		Ok       chan bool
	}
	DoneCalled chan bool
	DoneOutput struct {
		Ret0 chan (<-chan struct{})
	}
	ErrCalled chan bool
	ErrOutput struct {
		Ret0 chan error
	}
	ValueCalled chan bool
	ValueInput  struct {
		Key chan any
	}
	ValueOutput struct {
		Ret0 chan any
	}
}

func newMockContext(t vegr.T, timeout time.Duration) *mockContext {
	m := &mockContext{t: t, timeout: timeout}
	m.DeadlineCalled = make(chan bool, 100)
	m.DeadlineOutput.Deadline = make(chan time.Time, 100)
	m.DeadlineOutput.Ok = make(chan bool, 100)
	m.DoneCalled = make(chan bool, 100)
	m.DoneOutput.Ret0 = make(chan (<-chan struct{}), 100)
	m.ErrCalled = make(chan bool, 100)
	m.ErrOutput.Ret0 = make(chan error, 100)
	m.ValueCalled = make(chan bool, 100)
	m.ValueInput.Key = make(chan any, 100)
	m.ValueOutput.Ret0 = make(chan any, 100)
	return m
}
func (m *mockContext) Deadline() (deadline time.Time, ok bool) {
	m.t.Helper()
	m.DeadlineCalled <- true
	vegr.PopulateReturns(m.t, "Deadline", m.timeout, m.DeadlineOutput, &deadline, &ok)
	return deadline, ok
}
func (m *mockContext) Done() (ret0 <-chan struct{}) {
	m.t.Helper()
	m.DoneCalled <- true
	vegr.PopulateReturns(m.t, "Done", m.timeout, m.DoneOutput, &ret0)
	return ret0
}
func (m *mockContext) Err() (ret0 error) {
	m.t.Helper()
	m.ErrCalled <- true
	vegr.PopulateReturns(m.t, "Err", m.timeout, m.ErrOutput, &ret0)
	return ret0
}
func (m *mockContext) Value(key any) (ret0 any) {
	m.t.Helper()
	m.ValueCalled <- true
	m.ValueInput.Key <- key
	vegr.PopulateReturns(m.t, "Value", m.timeout, m.ValueOutput, &ret0)
	return ret0
}

type mockClientStream struct {
	t            vegr.T
	timeout      time.Duration
	HeaderCalled chan bool
	HeaderOutput struct {
		Ret0 chan metadata.MD
		Ret1 chan error
	}
	TrailerCalled chan bool
	TrailerOutput struct {
		Ret0 chan metadata.MD
	}
	CloseSendCalled chan bool
	CloseSendOutput struct {
		Ret0 chan error
	}
	ContextCalled chan bool
	ContextOutput struct {
		Ret0 chan context.Context
	}
	SendMsgCalled chan bool
	SendMsgInput  struct {
		M chan interface{}
	}
	SendMsgOutput struct {
		Ret0 chan error
	}
	RecvMsgCalled chan bool
	RecvMsgInput  struct {
		M chan interface{}
	}
	RecvMsgOutput struct {
		Ret0 chan error
	}
}

func newMockClientStream(t vegr.T, timeout time.Duration) *mockClientStream {
	m := &mockClientStream{t: t, timeout: timeout}
	m.HeaderCalled = make(chan bool, 100)
	m.HeaderOutput.Ret0 = make(chan metadata.MD, 100)
	m.HeaderOutput.Ret1 = make(chan error, 100)
	m.TrailerCalled = make(chan bool, 100)
	m.TrailerOutput.Ret0 = make(chan metadata.MD, 100)
	m.CloseSendCalled = make(chan bool, 100)
	m.CloseSendOutput.Ret0 = make(chan error, 100)
	m.ContextCalled = make(chan bool, 100)
	m.ContextOutput.Ret0 = make(chan context.Context, 100)
	m.SendMsgCalled = make(chan bool, 100)
	m.SendMsgInput.M = make(chan interface{}, 100)
	m.SendMsgOutput.Ret0 = make(chan error, 100)
	m.RecvMsgCalled = make(chan bool, 100)
	m.RecvMsgInput.M = make(chan interface{}, 100)
	m.RecvMsgOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockClientStream) Header() (ret0 metadata.MD, ret1 error) {
	m.t.Helper()
	m.HeaderCalled <- true
	vegr.PopulateReturns(m.t, "Header", m.timeout, m.HeaderOutput, &ret0, &ret1)
	return ret0, ret1
}
func (m *mockClientStream) Trailer() (ret0 metadata.MD) {
	m.t.Helper()
	m.TrailerCalled <- true
	vegr.PopulateReturns(m.t, "Trailer", m.timeout, m.TrailerOutput, &ret0)
	return ret0
}
func (m *mockClientStream) CloseSend() (ret0 error) {
	m.t.Helper()
	m.CloseSendCalled <- true
	vegr.PopulateReturns(m.t, "CloseSend", m.timeout, m.CloseSendOutput, &ret0)
	return ret0
}
func (m *mockClientStream) Context() (ret0 context.Context) {
	m.t.Helper()
	m.ContextCalled <- true
	vegr.PopulateReturns(m.t, "Context", m.timeout, m.ContextOutput, &ret0)
	return ret0
}
func (m *mockClientStream) SendMsg(m_ interface{}) (ret0 error) {
	m.t.Helper()
	m.SendMsgCalled <- true
	m.SendMsgInput.M <- m_
	vegr.PopulateReturns(m.t, "SendMsg", m.timeout, m.SendMsgOutput, &ret0)
	return ret0
}
func (m *mockClientStream) RecvMsg(m_ interface{}) (ret0 error) {
	m.t.Helper()
	m.RecvMsgCalled <- true
	m.RecvMsgInput.M <- m_
	vegr.PopulateReturns(m.t, "RecvMsg", m.timeout, m.RecvMsgOutput, &ret0)
	return ret0
}
