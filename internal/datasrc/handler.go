package datasrc

import (
	"context"
	"sync"
	"time"
)

// DSrcer defines the interface of a handler data
type DSrcer interface {
	Close() error                       // 断开mqtt连接
	SendData(string, interface{}) error // 发送上行数据
	DataDownChan() chan DataDownPayload // 返回订阅到的消息数据channel
	IsConnected() bool
}

// SyncMSG for interactive with gateway
type SyncMSG interface {
	StartAndWaitRet(mID string, timeOut time.Duration) (any, error)
	HandleReceive(mID string, payload any)
}

type ctxCancel struct {
	mu     *sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
	rsp    chan interface{}
}

// MsgInteractive ...
type MsgInteractive struct {
	mccs map[string]ctxCancel
}

// NewMsgInteractive .....
func NewMsgInteractive() SyncMSG {
	return &MsgInteractive{
		mccs: make(map[string]ctxCancel),
	}
}

// StartAndWaitRet ...
func (mi *MsgInteractive) StartAndWaitRet(mID string, timeOut time.Duration) (any, error) {
	mcc, ok := mi.mccs[mID]
	if !ok {
		tctx, tcancel := context.WithTimeout(context.Background(), timeOut)
		mcc = ctxCancel{
			mu:     &sync.Mutex{},
			ctx:    tctx,
			cancel: tcancel,
			rsp:    make(chan any),
		}
		mi.mccs[mID] = mcc
	}
	mcc.mu.Lock()
	defer mcc.cancel()
	defer func() {
		mcc.mu.Unlock()
		close(mcc.rsp)
		delete(mi.mccs, mID)
	}()
	for {
		select {
		case <-mcc.ctx.Done():
			return nil, mcc.ctx.Err()
		case r := <-mcc.rsp:
			return r, nil
		}
	}
}

// HandleReceive ....
func (mi *MsgInteractive) HandleReceive(mID string, payload any) {
	if mcc, ok := mi.mccs[mID]; ok {
		select {
		case <-mcc.ctx.Done():
			return
		default:
			mcc.rsp <- payload
		}
	}
}
