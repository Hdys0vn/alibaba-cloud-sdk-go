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

// PushUnBind invokes the mpaas.PushUnBind API synchronously
func (client *Client) PushUnBind(request *PushUnBindRequest) (response *PushUnBindResponse, err error) {
	response = CreatePushUnBindResponse()
	err = client.DoAction(request, response)
	return
}

// PushUnBindWithChan invokes the mpaas.PushUnBind API asynchronously
func (client *Client) PushUnBindWithChan(request *PushUnBindRequest) (<-chan *PushUnBindResponse, <-chan error) {
	responseChan := make(chan *PushUnBindResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushUnBind(request)
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

// PushUnBindWithCallback invokes the mpaas.PushUnBind API asynchronously
func (client *Client) PushUnBindWithCallback(request *PushUnBindRequest, callback func(response *PushUnBindResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushUnBindResponse
		var err error
		defer close(result)
		response, err = client.PushUnBind(request)
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

// PushUnBindRequest is the request struct for api PushUnBind
type PushUnBindRequest struct {
	*requests.RpcRequest
	UserId        string `position:"Body" name:"UserId"`
	DeliveryToken string `position:"Body" name:"DeliveryToken"`
	AppId         string `position:"Body" name:"AppId"`
	WorkspaceId   string `position:"Body" name:"WorkspaceId"`
}

// PushUnBindResponse is the response struct for api PushUnBind
type PushUnBindResponse struct {
	*responses.BaseResponse
	ResultMessage string     `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string     `json:"ResultCode" xml:"ResultCode"`
	RequestId     string     `json:"RequestId" xml:"RequestId"`
	PushResult    PushResult `json:"PushResult" xml:"PushResult"`
}

// CreatePushUnBindRequest creates a request to invoke PushUnBind API
func CreatePushUnBindRequest() (request *PushUnBindRequest) {
	request = &PushUnBindRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "PushUnBind", "", "")
	request.Method = requests.POST
	return
}

// CreatePushUnBindResponse creates a response to parse from PushUnBind response
func CreatePushUnBindResponse() (response *PushUnBindResponse) {
	response = &PushUnBindResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}