package domain_intl

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// LookupTmchNotice invokes the domain_intl.LookupTmchNotice API synchronously
// api document: https://help.aliyun.com/api/domain-intl/lookuptmchnotice.html
func (client *Client) LookupTmchNotice(request *LookupTmchNoticeRequest) (response *LookupTmchNoticeResponse, err error) {
	response = CreateLookupTmchNoticeResponse()
	err = client.DoAction(request, response)
	return
}

// LookupTmchNoticeWithChan invokes the domain_intl.LookupTmchNotice API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/lookuptmchnotice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) LookupTmchNoticeWithChan(request *LookupTmchNoticeRequest) (<-chan *LookupTmchNoticeResponse, <-chan error) {
	responseChan := make(chan *LookupTmchNoticeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LookupTmchNotice(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// LookupTmchNoticeWithCallback invokes the domain_intl.LookupTmchNotice API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/lookuptmchnotice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) LookupTmchNoticeWithCallback(request *LookupTmchNoticeRequest, callback func(response *LookupTmchNoticeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LookupTmchNoticeResponse
		var err error
		defer close(result)
		response, err = client.LookupTmchNotice(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// LookupTmchNoticeRequest is the request struct for api LookupTmchNotice
type LookupTmchNoticeRequest struct {
	*requests.RpcRequest
	ClaimKey     string `position:"Query" name:"ClaimKey"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// LookupTmchNoticeResponse is the response struct for api LookupTmchNotice
type LookupTmchNoticeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        int    `json:"Id" xml:"Id"`
	NotBefore string `json:"NotBefore" xml:"NotBefore"`
	NotAfter  string `json:"NotAfter" xml:"NotAfter"`
	Label     string `json:"Label" xml:"Label"`
	Claims    Claims `json:"Claims" xml:"Claims"`
}

// CreateLookupTmchNoticeRequest creates a request to invoke LookupTmchNotice API
func CreateLookupTmchNoticeRequest() (request *LookupTmchNoticeRequest) {
	request = &LookupTmchNoticeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "LookupTmchNotice", "domain", "openAPI")
	return
}

// CreateLookupTmchNoticeResponse creates a response to parse from LookupTmchNotice response
func CreateLookupTmchNoticeResponse() (response *LookupTmchNoticeResponse) {
	response = &LookupTmchNoticeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
