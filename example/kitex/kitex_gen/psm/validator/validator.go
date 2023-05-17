// Code generated by Kitex v0.5.2. DO NOT EDIT.

package validator

import (
	psm "a/b/c/kitex_gen/psm"
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return validatorServiceInfo
}

var validatorServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Validator"
	handlerType := (*psm.Validator)(nil)
	methods := map[string]kitex.MethodInfo{
		"Method1": kitex.NewMethodInfo(method1Handler, newMethod1Args, newMethod1Result, false),
	}
	extra := map[string]interface{}{
		"PackageName": "psm",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func method1Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(psm.IntValidate)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(psm.Validator).Method1(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *Method1Args:
		success, err := handler.(psm.Validator).Method1(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*Method1Result)
		realResult.Success = success
	}
	return nil
}
func newMethod1Args() interface{} {
	return &Method1Args{}
}

func newMethod1Result() interface{} {
	return &Method1Result{}
}

type Method1Args struct {
	Req *psm.IntValidate
}

func (p *Method1Args) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(psm.IntValidate)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *Method1Args) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *Method1Args) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *Method1Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in Method1Args")
	}
	return proto.Marshal(p.Req)
}

func (p *Method1Args) Unmarshal(in []byte) error {
	msg := new(psm.IntValidate)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var Method1Args_Req_DEFAULT *psm.IntValidate

func (p *Method1Args) GetReq() *psm.IntValidate {
	if !p.IsSetReq() {
		return Method1Args_Req_DEFAULT
	}
	return p.Req
}

func (p *Method1Args) IsSetReq() bool {
	return p.Req != nil
}

func (p *Method1Args) GetFirstArgument() interface{} {
	return p.Req
}

type Method1Result struct {
	Success *psm.IntValidate
}

var Method1Result_Success_DEFAULT *psm.IntValidate

func (p *Method1Result) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(psm.IntValidate)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *Method1Result) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *Method1Result) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *Method1Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in Method1Result")
	}
	return proto.Marshal(p.Success)
}

func (p *Method1Result) Unmarshal(in []byte) error {
	msg := new(psm.IntValidate)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *Method1Result) GetSuccess() *psm.IntValidate {
	if !p.IsSetSuccess() {
		return Method1Result_Success_DEFAULT
	}
	return p.Success
}

func (p *Method1Result) SetSuccess(x interface{}) {
	p.Success = x.(*psm.IntValidate)
}

func (p *Method1Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *Method1Result) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Method1(ctx context.Context, Req *psm.IntValidate) (r *psm.IntValidate, err error) {
	var _args Method1Args
	_args.Req = Req
	var _result Method1Result
	if err = p.c.Call(ctx, "Method1", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
