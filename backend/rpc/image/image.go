// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package image

import (
	"bytes"
	"reflect"
	"database/sql/driver"
	"errors"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type ResponseState int64
const (
  ResponseState_StateOk ResponseState = 0
  ResponseState_StateError ResponseState = 1
  ResponseState_StateEmpty ResponseState = 2
)

func (p ResponseState) String() string {
  switch p {
  case ResponseState_StateOk: return "StateOk"
  case ResponseState_StateError: return "StateError"
  case ResponseState_StateEmpty: return "StateEmpty"
  }
  return "<UNSET>"
}

func ResponseStateFromString(s string) (ResponseState, error) {
  switch s {
  case "StateOk": return ResponseState_StateOk, nil 
  case "StateError": return ResponseState_StateError, nil 
  case "StateEmpty": return ResponseState_StateEmpty, nil 
  }
  return ResponseState(0), fmt.Errorf("not a valid ResponseState string")
}


func ResponseStatePtr(v ResponseState) *ResponseState { return &v }

func (p ResponseState) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *ResponseState) UnmarshalText(text []byte) error {
q, err := ResponseStateFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *ResponseState) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = ResponseState(v)
return nil
}

func (p * ResponseState) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
// Attributes:
//  - ImgPath
//  - App
//  - ProjectType
type Request struct {
  ImgPath string `thrift:"imgPath,1" db:"imgPath" json:"imgPath"`
  App string `thrift:"app,2" db:"app" json:"app"`
  ProjectType int32 `thrift:"projectType,3" db:"projectType" json:"projectType"`
}

func NewRequest() *Request {
  return &Request{}
}


func (p *Request) GetImgPath() string {
  return p.ImgPath
}

func (p *Request) GetApp() string {
  return p.App
}

func (p *Request) GetProjectType() int32 {
  return p.ProjectType
}
func (p *Request) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Request)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ImgPath = v
}
  return nil
}

func (p *Request)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.App = v
}
  return nil
}

func (p *Request)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.ProjectType = v
}
  return nil
}

func (p *Request) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Request"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Request) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("imgPath", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:imgPath: ", p), err) }
  if err := oprot.WriteString(string(p.ImgPath)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.imgPath (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:imgPath: ", p), err) }
  return err
}

func (p *Request) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("app", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:app: ", p), err) }
  if err := oprot.WriteString(string(p.App)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.app (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:app: ", p), err) }
  return err
}

func (p *Request) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("projectType", thrift.I32, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:projectType: ", p), err) }
  if err := oprot.WriteI32(int32(p.ProjectType)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.projectType (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:projectType: ", p), err) }
  return err
}

func (p *Request) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Request(%+v)", *p)
}

// Attributes:
//  - BookID
//  - State
type Response struct {
  BookID string `thrift:"bookID,1" db:"bookID" json:"bookID"`
  State ResponseState `thrift:"state,2" db:"state" json:"state"`
}

func NewResponse() *Response {
  return &Response{}
}


func (p *Response) GetBookID() string {
  return p.BookID
}

func (p *Response) GetState() ResponseState {
  return p.State
}
func (p *Response) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Response)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.BookID = v
}
  return nil
}

func (p *Response)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := ResponseState(v)
  p.State = temp
}
  return nil
}

func (p *Response) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Response"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Response) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("bookID", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:bookID: ", p), err) }
  if err := oprot.WriteString(string(p.BookID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.bookID (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:bookID: ", p), err) }
  return err
}

func (p *Response) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("state", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:state: ", p), err) }
  if err := oprot.WriteI32(int32(p.State)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.state (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:state: ", p), err) }
  return err
}

func (p *Response) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Response(%+v)", *p)
}

type ImageService interface {
  // Parameters:
  //  - Request
  GetIdentify(ctx context.Context, request *Request) (r *Response, err error)
}

type ImageServiceClient struct {
  c thrift.TClient
}

// Deprecated: Use NewImageService instead
func NewImageServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ImageServiceClient {
  return &ImageServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

// Deprecated: Use NewImageService instead
func NewImageServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ImageServiceClient {
  return &ImageServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewImageServiceClient(c thrift.TClient) *ImageServiceClient {
  return &ImageServiceClient{
    c: c,
  }
}

// Parameters:
//  - Request
func (p *ImageServiceClient) GetIdentify(ctx context.Context, request *Request) (r *Response, err error) {
  var _args0 ImageServiceGetIdentifyArgs
  _args0.Request = request
  var _result1 ImageServiceGetIdentifyResult
  if err = p.c.Call(ctx, "getIdentify", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type ImageServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler ImageService
}

func (p *ImageServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *ImageServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *ImageServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewImageServiceProcessor(handler ImageService) *ImageServiceProcessor {

  self2 := &ImageServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["getIdentify"] = &imageServiceProcessorGetIdentify{handler:handler}
return self2
}

func (p *ImageServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x3

}

type imageServiceProcessorGetIdentify struct {
  handler ImageService
}

func (p *imageServiceProcessorGetIdentify) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ImageServiceGetIdentifyArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getIdentify", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := ImageServiceGetIdentifyResult{}
var retval *Response
  var err2 error
  if retval, err2 = p.handler.GetIdentify(ctx, args.Request); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getIdentify: " + err2.Error())
    oprot.WriteMessageBegin("getIdentify", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("getIdentify", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type ImageServiceGetIdentifyArgs struct {
  Request *Request `thrift:"request,1" db:"request" json:"request"`
}

func NewImageServiceGetIdentifyArgs() *ImageServiceGetIdentifyArgs {
  return &ImageServiceGetIdentifyArgs{}
}

var ImageServiceGetIdentifyArgs_Request_DEFAULT *Request
func (p *ImageServiceGetIdentifyArgs) GetRequest() *Request {
  if !p.IsSetRequest() {
    return ImageServiceGetIdentifyArgs_Request_DEFAULT
  }
return p.Request
}
func (p *ImageServiceGetIdentifyArgs) IsSetRequest() bool {
  return p.Request != nil
}

func (p *ImageServiceGetIdentifyArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ImageServiceGetIdentifyArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.Request = &Request{}
  if err := p.Request.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
  }
  return nil
}

func (p *ImageServiceGetIdentifyArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getIdentify_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ImageServiceGetIdentifyArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err) }
  if err := p.Request.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err) }
  return err
}

func (p *ImageServiceGetIdentifyArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ImageServiceGetIdentifyArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ImageServiceGetIdentifyResult struct {
  Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewImageServiceGetIdentifyResult() *ImageServiceGetIdentifyResult {
  return &ImageServiceGetIdentifyResult{}
}

var ImageServiceGetIdentifyResult_Success_DEFAULT *Response
func (p *ImageServiceGetIdentifyResult) GetSuccess() *Response {
  if !p.IsSetSuccess() {
    return ImageServiceGetIdentifyResult_Success_DEFAULT
  }
return p.Success
}
func (p *ImageServiceGetIdentifyResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ImageServiceGetIdentifyResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ImageServiceGetIdentifyResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &Response{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *ImageServiceGetIdentifyResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getIdentify_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ImageServiceGetIdentifyResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *ImageServiceGetIdentifyResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ImageServiceGetIdentifyResult(%+v)", *p)
}


