package errors

const (
	SUCCESS             = 0
	FAILURE             = 1
	AuthorizationError  = 410
	NotFound            = 411
	NotLogin            = 412
	NotTimeout          = 413
	InvalidParameter    = 10000
	UserDoesNotExist    = 10001
	ServerError         = 10101
	TooManyRequests     = 10102
	SocialAddError      = 20100
	SocialListError     = 20101
	SocialAllReadyExist = 20102
	SocialUpdateError   = 20103
	SocialNameNotExist  = 20104
)

type ErrorText struct {
	Language string
}

func NewErrorText(language string) *ErrorText {
	return &ErrorText{
		Language: language,
	}
}

func (et *ErrorText) Text(code int) (str string) {
	var ok bool
	switch et.Language {
	case "zh_CN":
		str, ok = zhCNText[code]
	case "en":
		str, ok = enUSText[code]
	default:
		str, ok = zhCNText[code]
	}
	if !ok {
		return "unknown error"
	}
	return
}
