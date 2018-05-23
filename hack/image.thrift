namespace cpp image 
namespace go  image 

enum ResponseState {
    StateOk = 0,
    StateError = 1,
    StateEmpty = 2
}

struct Request {
    1: string imgPath,
    2: string app,
    3: i32 projectType
}

struct Response {
    1: i32 bookID = 0,
    2: ResponseState state
}

service ImageService {
    Response getIdentify(1: Request request);
}
