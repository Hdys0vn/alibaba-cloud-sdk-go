package mpaas

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

// RevokePushMessage invokes the mpaas.RevokePushMessage API synchronously
func (client *Client) RevokePushMessage(request *RevokePushMessageRequest) (response *RevokePushMessageResponse, err error) {
	response = CreateRevokePushMessageResponse()
	err = client.DoAction(request, response)
	return
}

// RevokePushMessageWithChan invokes the mpaas.RevokePushMessage API asynchronously
func (client *Client) RevokePushMessageWithChan(request *RevokePushMessageRequest) (<-chan *RevokePushMessageResponse, <-chan error) {
	responseChan := make(chan *RevokePushMessageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokePushMessage(request)
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

// RevokePushMessageWithCallback invokes the mpaas.RevokePushMessage API asynchronously
func (client *Client) RevokePushMessageWithCallback(request *RevokePushMessageRequest, callback func(response *RevokePushMessageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokePushMessageResponse
		var err error
		defer close(result)
		response, err = client.RevokePushMessage(request)
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

// RevokePushMessageRequest is the request struct for api RevokePushMessage
type RevokePushMessageRequest struct {
	*requests.RpcRequest
	TargetId    string `position:"Body" name:"TargetId"`
	MessageId   string `position:"Body" name:"MessageId"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// RevokePushMessageResponse is the response struct for api RevokePushMessage
type RevokePushMessageResponse struct {
	*responses.BaseResponse
	ResultMessage string     `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string     `json:"ResultCode" xml:"ResultCode"`
	RequestId     string     `json:"RequestId" xml:"RequestId"`
	PushResult    PushResult `json:"PushResult" xml:"PushResult"`
}

// CreateRevokePushMessageRequest creates a request to invoke RevokePushMessage API
func CreateRevokePushMessageRequest() (request *RevokePushMessageRequest) {
	request = &RevokePushMessageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "RevokePushMessage", "", "")
	request.Method = requests.POST
	return
}

// CreateRevokePushMessageResponse creates a response to parse from RevokePushMessage response
func CreateRevokePushMessageResponse() (response *RevokePushMessageResponse) {
	response = &RevokePushMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}