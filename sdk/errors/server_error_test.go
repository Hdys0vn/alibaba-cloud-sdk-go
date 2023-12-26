package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerError(t *testing.T) {
	e := NewServerError(400, "content", "comment")
	assert.NotNil(t, e)
	serverError, ok := e.(*ServerError)
	assert.True(t, ok)
	assert.Nil(t, serverError.OriginError())
	assert.Equal(t, 400, serverError.HttpStatus())
	assert.Equal(t, "", serverError.RequestId())
	assert.Equal(t, "", serverError.ErrorCode())
	assert.Equal(t, "", serverError.Recommend())
	assert.Equal(t, "", serverError.HostId())
	assert.Equal(t, "comment", serverError.Comment())
	assert.Equal(t, "content", serverError.Message())
	assert.Equal(t, "SDK.ServerError\nErrorCode: \nRecommend: comment\nRequestId: \nMessage: content\nRespHeaders: map[]", serverError.Error())
}

func TestServerErrorWithContent(t *testing.T) {
	e := NewServerError(400, `{"RequestId":"request id","HostId":"host id","Code":"InvalidAK","Recommend":"recommend","Message":"message"}`, "comment")
	assert.NotNil(t, e)
	serverError, ok := e.(*ServerError)
	assert.True(t, ok)
	assert.Nil(t, serverError.OriginError())
	assert.Equal(t, 400, serverError.HttpStatus())
	assert.Equal(t, "request id", serverError.RequestId())
	assert.Equal(t, "host id", serverError.HostId())
	assert.Equal(t, "InvalidAK", serverError.ErrorCode())
	assert.Equal(t, "recommend", serverError.Recommend())
	assert.Equal(t, "comment", serverError.Comment())
	assert.Equal(t, "message", serverError.Message())
	assert.Equal(t, "SDK.ServerError\nErrorCode: InvalidAK\nRecommend: commentrecommend\nRequestId: request id\nMessage: message\nRespHeaders: map[]", serverError.Error())

	e = NewServerError(400, `{"RequestId":"request id","HostId":"host id","Code":"InvalidAK","Recommend":"recommend","Message":"message","AccessDeniedDetail":{}}`, "comment")
	assert.NotNil(t, e)
	serverError, ok = e.(*ServerError)
	assert.True(t, ok)
	assert.Equal(t, "SDK.ServerError\nErrorCode: InvalidAK\nRecommend: commentrecommend\nRequestId: request id\nMessage: message\nRespHeaders: map[]\nAccessDeniedDetail: map[]", serverError.Error())
	e = NewServerError(400, `{"RequestId":"request id","HostId":"host id","Code":"InvalidAK","Recommend":"recommend","Message":"message","AccessDeniedDetail":{"AuthAction":"ram:ListUsers","UserId":123}}`, "comment")
	assert.NotNil(t, e)
	serverError, ok = e.(*ServerError)
	assert.True(t, ok)
	accessDeniedDetail := serverError.AccessDeniedDetail()
	assert.Equal(t, "ram:ListUsers", accessDeniedDetail["AuthAction"])
	assert.Equal(t, float64(123), accessDeniedDetail["UserId"])
	assert.Equal(t, "SDK.ServerError\nErrorCode: InvalidAK\nRecommend: commentrecommend\nRequestId: request id\nMessage: message\nRespHeaders: map[]\nAccessDeniedDetail: map[AuthAction:ram:ListUsers UserId:%!s(float64=123)]", serverError.Error())
}

func TestWrapServerError(t *testing.T) {
	err := NewServerError(400, `{"Code":"SignatureDoesNotMatch","Message":"Specified signature is not matched with our calculation. server string to sign is:hehe"}`, "comment")
	se, ok := err.(*ServerError)
	assert.True(t, ok)
	assert.Equal(t, "SignatureDoesNotMatch", se.ErrorCode())
	m := make(map[string]string)
	m["StringToSign"] = "not match"
	WrapServerError(se, m)
	assert.Equal(t, "This may be a bug with the SDK and we hope you can submit this question in the github issue(https://github.com/hdys0vn/alibaba-cloud-sdk-go/issues), thanks very much", se.Recommend())

	err = NewServerError(400, `{"Code":"SignatureDoesNotMatch","Message":"Specified signature is not matched with our calculation. server string to sign is:match"}`, "comment")
	se, ok = err.(*ServerError)
	assert.True(t, ok)
	assert.Equal(t, "SignatureDoesNotMatch", se.ErrorCode())
	m = make(map[string]string)
	m["StringToSign"] = "match"
	WrapServerError(se, m)
	assert.Equal(t, "InvalidAccessKeySecret: Please check you AccessKeySecret", se.Recommend())

	err = NewServerError(400, `{"Code":"Other"}`, "comment")
	se, ok = err.(*ServerError)
	assert.True(t, ok)
	assert.Equal(t, "Other", se.ErrorCode())
	m = make(map[string]string)
	WrapServerError(se, m)
	assert.Equal(t, "", se.Recommend())

	err = NewServerError(400, `{"Code":"Other","AccessDeniedDetail":{}}`, "comment")
	se, ok = err.(*ServerError)
	assert.True(t, ok)
	assert.Equal(t, "Other", se.ErrorCode())
	m = make(map[string]string)
	WrapServerError(se, m)
	assert.Equal(t, "", se.Recommend())
	assert.Equal(t, "SDK.ServerError\nErrorCode: Other\nRecommend: comment\nRequestId: \nMessage: {\"Code\":\"Other\",\"AccessDeniedDetail\":{}}\nRespHeaders: map[]\nAccessDeniedDetail: map[]", se.Error())
	err = NewServerError(400, `{"Code":"Other","AccessDeniedDetail":{"AuthAction":"ram:ListUsers","UserId":123}}`, "comment")
	se, ok = err.(*ServerError)
	assert.True(t, ok)
	assert.Equal(t, "Other", se.ErrorCode())
	m = make(map[string]string)
	WrapServerError(se, m)
	assert.Equal(t, "", se.Recommend())
	assert.Equal(t, "SDK.ServerError\nErrorCode: Other\nRecommend: comment\nRequestId: \nMessage: {\"Code\":\"Other\",\"AccessDeniedDetail\":{\"AuthAction\":\"ram:ListUsers\",\"UserId\":123}}\nRespHeaders: map[]\nAccessDeniedDetail: map[AuthAction:ram:ListUsers UserId:%!s(float64=123)]", se.Error())
}
