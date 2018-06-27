/**
 * Autogenerated by Thrift Compiler (0.11.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#include "image_types.h"

#include <algorithm>
#include <ostream>

#include <thrift/TToString.h>

namespace image {

int _kResponseStateValues[] = {
  ResponseState::StateOk,
  ResponseState::StateError,
  ResponseState::StateEmpty
};
const char* _kResponseStateNames[] = {
  "StateOk",
  "StateError",
  "StateEmpty"
};
const std::map<int, const char*> _ResponseState_VALUES_TO_NAMES(::apache::thrift::TEnumIterator(3, _kResponseStateValues, _kResponseStateNames), ::apache::thrift::TEnumIterator(-1, NULL, NULL));

std::ostream& operator<<(std::ostream& out, const ResponseState::type& val) {
  std::map<int, const char*>::const_iterator it = _ResponseState_VALUES_TO_NAMES.find(val);
  if (it != _ResponseState_VALUES_TO_NAMES.end()) {
    out << it->second;
  } else {
    out << static_cast<int>(val);
  }
  return out;
}


Request::~Request() throw() {
}


void Request::__set_imgPath(const std::string& val) {
  this->imgPath = val;
}

void Request::__set_app(const std::string& val) {
  this->app = val;
}

void Request::__set_projectType(const int32_t val) {
  this->projectType = val;
}
std::ostream& operator<<(std::ostream& out, const Request& obj)
{
  obj.printTo(out);
  return out;
}


uint32_t Request::read(::apache::thrift::protocol::TProtocol* iprot) {

  ::apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->imgPath);
          this->__isset.imgPath = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->app);
          this->__isset.app = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 3:
        if (ftype == ::apache::thrift::protocol::T_I32) {
          xfer += iprot->readI32(this->projectType);
          this->__isset.projectType = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t Request::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("Request");

  xfer += oprot->writeFieldBegin("imgPath", ::apache::thrift::protocol::T_STRING, 1);
  xfer += oprot->writeString(this->imgPath);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("app", ::apache::thrift::protocol::T_STRING, 2);
  xfer += oprot->writeString(this->app);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("projectType", ::apache::thrift::protocol::T_I32, 3);
  xfer += oprot->writeI32(this->projectType);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(Request &a, Request &b) {
  using ::std::swap;
  swap(a.imgPath, b.imgPath);
  swap(a.app, b.app);
  swap(a.projectType, b.projectType);
  swap(a.__isset, b.__isset);
}

Request::Request(const Request& other0) {
  imgPath = other0.imgPath;
  app = other0.app;
  projectType = other0.projectType;
  __isset = other0.__isset;
}
Request& Request::operator=(const Request& other1) {
  imgPath = other1.imgPath;
  app = other1.app;
  projectType = other1.projectType;
  __isset = other1.__isset;
  return *this;
}
void Request::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "Request(";
  out << "imgPath=" << to_string(imgPath);
  out << ", " << "app=" << to_string(app);
  out << ", " << "projectType=" << to_string(projectType);
  out << ")";
}


Response::~Response() throw() {
}


void Response::__set_bookID(const std::string& val) {
  this->bookID = val;
}

void Response::__set_state(const ResponseState::type val) {
  this->state = val;
}
std::ostream& operator<<(std::ostream& out, const Response& obj)
{
  obj.printTo(out);
  return out;
}


uint32_t Response::read(::apache::thrift::protocol::TProtocol* iprot) {

  ::apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->bookID);
          this->__isset.bookID = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_I32) {
          int32_t ecast2;
          xfer += iprot->readI32(ecast2);
          this->state = (ResponseState::type)ecast2;
          this->__isset.state = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t Response::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("Response");

  xfer += oprot->writeFieldBegin("bookID", ::apache::thrift::protocol::T_STRING, 1);
  xfer += oprot->writeString(this->bookID);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("state", ::apache::thrift::protocol::T_I32, 2);
  xfer += oprot->writeI32((int32_t)this->state);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(Response &a, Response &b) {
  using ::std::swap;
  swap(a.bookID, b.bookID);
  swap(a.state, b.state);
  swap(a.__isset, b.__isset);
}

Response::Response(const Response& other3) {
  bookID = other3.bookID;
  state = other3.state;
  __isset = other3.__isset;
}
Response& Response::operator=(const Response& other4) {
  bookID = other4.bookID;
  state = other4.state;
  __isset = other4.__isset;
  return *this;
}
void Response::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "Response(";
  out << "bookID=" << to_string(bookID);
  out << ", " << "state=" << to_string(state);
  out << ")";
}

} // namespace
