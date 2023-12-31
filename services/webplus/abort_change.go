package webplus

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

// AbortChange invokes the webplus.AbortChange API synchronously
// api document: https://help.aliyun.com/api/webplus/abortchange.html
func (client *Client) AbortChange(request *AbortChangeRequest) (response *AbortChangeResponse, err error) {
	response = CreateAbortChangeResponse()
	err = client.DoAction(request, response)
	return
}

// AbortChangeWithChan invokes the webplus.AbortChange API asynchronously
// api document: https://help.aliyun.com/api/webplus/abortchange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbortChangeWithChan(request *AbortChangeRequest) (<-chan *AbortChangeResponse, <-chan error) {
	responseChan := make(chan *AbortChangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AbortChange(request)
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

// AbortChangeWithCallback invokes the webplus.AbortChange API asynchronously
// api document: https://help.aliyun.com/api/webplus/abortchange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbortChangeWithCallback(request *AbortChangeRequest, callback func(response *AbortChangeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AbortChangeResponse
		var err error
		defer close(result)
		response, err = client.AbortChange(request)
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

// AbortChangeRequest is the request struct for api AbortChange
type AbortChangeRequest struct {
	*requests.RoaRequest
	ChangeId string `position:"Body" name:"ChangeId"`
}

// AbortChangeResponse is the response struct for api AbortChange
type AbortChangeResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Code      string    `json:"Code" xml:"Code"`
	Message   string    `json:"Message" xml:"Message"`
	EnvChange EnvChange `json:"EnvChange" xml:"EnvChange"`
}

// CreateAbortChangeRequest creates a request to invoke AbortChange API
func CreateAbortChangeRequest() (request *AbortChangeRequest) {
	request = &AbortChangeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "AbortChange", "/pop/v1/wam/change/abort", "", "")
	request.Method = requests.POST
	return
}

// CreateAbortChangeResponse creates a response to parse from AbortChange response
func CreateAbortChangeResponse() (response *AbortChangeResponse) {
	response = &AbortChangeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
