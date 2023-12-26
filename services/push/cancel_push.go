package push

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

// CancelPush invokes the push.CancelPush API synchronously
func (client *Client) CancelPush(request *CancelPushRequest) (response *CancelPushResponse, err error) {
	response = CreateCancelPushResponse()
	err = client.DoAction(request, response)
	return
}

// CancelPushWithChan invokes the push.CancelPush API asynchronously
func (client *Client) CancelPushWithChan(request *CancelPushRequest) (<-chan *CancelPushResponse, <-chan error) {
	responseChan := make(chan *CancelPushResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelPush(request)
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

// CancelPushWithCallback invokes the push.CancelPush API asynchronously
func (client *Client) CancelPushWithCallback(request *CancelPushRequest, callback func(response *CancelPushResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelPushResponse
		var err error
		defer close(result)
		response, err = client.CancelPush(request)
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

// CancelPushRequest is the request struct for api CancelPush
type CancelPushRequest struct {
	*requests.RpcRequest
	MessageId requests.Integer `position:"Query" name:"MessageId"`
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
}

// CancelPushResponse is the response struct for api CancelPush
type CancelPushResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelPushRequest creates a request to invoke CancelPush API
func CreateCancelPushRequest() (request *CancelPushRequest) {
	request = &CancelPushRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "CancelPush", "", "")
	request.Method = requests.POST
	return
}

// CreateCancelPushResponse creates a response to parse from CancelPush response
func CreateCancelPushResponse() (response *CancelPushResponse) {
	response = &CancelPushResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}