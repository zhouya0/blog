package errcode

var (
	Success = NewError(0, "Success")
	ServerError = NewError(10000000, "Server internal error")
	InvalidParams = NewError(10000001, "Params invalid")
	NotFound = NewError(10000002, "Not found")
	UnauthorizedAuthNotExist = NewError(10000003, "Authorized failed, can't find AppKey and AppSecret")
	UnauthorizedTokenError = NewError(10000004, "Authorized failed, token error")
	UnauthorizedTokenTimeOut = NewError(10000005, "Authorized failed, token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "Authorized failed, token generate failed")
	TooManyRequests = NewError(10000007, "Too many requests")
)
