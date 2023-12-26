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

// CompleteContinuouslyPush invokes the push.CompleteContinuouslyPush API synchronously
func (client *Client) CompleteContinuouslyPush(request *CompleteContinuouslyPushRequest) (response *CompleteContinuouslyPushResponse, err error) {
	response = CreateCompleteContinuouslyPushResponse()
	err = client.DoAction(request, response)
	return
}

// CompleteContinuouslyPushWithChan invokes the push.CompleteContinuouslyPush API asynchronously
func (client *Client) CompleteContinuouslyPushWithChan(request *CompleteContinuouslyPushRequest) (<-chan *CompleteContinuouslyPushResponse, <-chan error) {
	responseChan := make(chan *CompleteContinuouslyPushResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CompleteContinuouslyPush(request)
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

// CompleteContinuouslyPushWithCallback invokes the push.CompleteContinuouslyPush API asynchronously
func (client *Client) CompleteContinuouslyPushWithCallback(request *CompleteContinuouslyPushRequest, callback func(response *CompleteContinuouslyPushResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CompleteContinuouslyPushResponse
		var err error
		defer close(result)
		response, err = client.CompleteContinuouslyPush(request)
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

// CompleteContinuouslyPushRequest is the request struct for api CompleteContinuouslyPush
type CompleteContinuouslyPushRequest struct {
	*requests.RpcRequest
	MessageId string           `position:"Query" name:"MessageId"`
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
}

// CompleteContinuouslyPushResponse is the response struct for api CompleteContinuouslyPush
type CompleteContinuouslyPushResponse struct {
	*responses.BaseResponse
	MessageId string `json:"MessageId" xml:"MessageId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCompleteContinuouslyPushRequest creates a request to invoke CompleteContinuouslyPush API
func CreateCompleteContinuouslyPushRequest() (request *CompleteContinuouslyPushRequest) {
	request = &CompleteContinuouslyPushRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "CompleteContinuouslyPush", "", "")
	request.Method = requests.POST
	return
}

// CreateCompleteContinuouslyPushResponse creates a response to parse from CompleteContinuouslyPush response
func CreateCompleteContinuouslyPushResponse() (response *CompleteContinuouslyPushResponse) {
	response = &CompleteContinuouslyPushResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
