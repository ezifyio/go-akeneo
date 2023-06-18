package goakeneo

import "time"

const (
	defaultHTTPTimeout       = 10 * time.Second
	defaultAccept            = "application/json"
	defaultContentType       = "application/json"
	defaultUploadContentType = "multipart/form-data"
	defaultUserAgent         = "go-akeneo v1.0.0"
	defaultRateLimit         = 5 // 5 requests per second
	defaultVersion           = AkeneoPimVersion6
	defaultRetry             = 2
	defaultRetryWaitTime     = 3 * time.Second
	defaultRetryMaxWaitTime  = 30 * time.Second
)

const (
	// AkeneoPimVersion4 is the version 4 of Akeneo PIM
	AkeneoPimVersion4 = iota + 4
	// AkeneoPimVersion5 is the version 5 of Akeneo PIM
	AkeneoPimVersion5
	// AkeneoPimVersion6 is the version 6 of Akeneo PIM
	AkeneoPimVersion6
	// AkeneoPimVersion7 is the version 7 of Akeneo PIM
	AkeneoPimVersion7
)

var (
	pimVersionMap = map[int]string{
		AkeneoPimVersion7: "7.0",
		AkeneoPimVersion6: "6.0",
		AkeneoPimVersion5: "5.0",
		AkeneoPimVersion4: "4.0",
	}
)
