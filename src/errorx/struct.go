package errorx

type EX struct {
	ClientIP    string `json:"clientIp,omitempty"`
	Method      string `json:"method,omitempty"`
	URL         string `json:"url,omitempty"`
	StatusCode  int    `json:"statusCode,omitempty"`
	Message     string `json:"message,omitempty"`
	Respons     any    `json:"response,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
	ErrorParent error  `json:"errorParent,omitempty"`
}

func (s *EX) Error() string {
	if s.Message == `` {
		return `Unknown`
	}
	return s.Message
}
