package alidns

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

// UpdateAppKeyState invokes the alidns.UpdateAppKeyState API synchronously
func (client *Client) UpdateAppKeyState(request *UpdateAppKeyStateRequest) (response *UpdateAppKeyStateResponse, err error) {
	response = CreateUpdateAppKeyStateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAppKeyStateWithChan invokes the alidns.UpdateAppKeyState API asynchronously
func (client *Client) UpdateAppKeyStateWithChan(request *UpdateAppKeyStateRequest) (<-chan *UpdateAppKeyStateResponse, <-chan error) {
	responseChan := make(chan *UpdateAppKeyStateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAppKeyState(request)
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

// UpdateAppKeyStateWithCallback invokes the alidns.UpdateAppKeyState API asynchronously
func (client *Client) UpdateAppKeyStateWithCallback(request *UpdateAppKeyStateRequest, callback func(response *UpdateAppKeyStateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAppKeyStateResponse
		var err error
		defer close(result)
		response, err = client.UpdateAppKeyState(request)
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

// UpdateAppKeyStateRequest is the request struct for api UpdateAppKeyState
type UpdateAppKeyStateRequest struct {
	*requests.RpcRequest
	AppKeyId string `position:"Query" name:"AppKeyId"`
	State    string `position:"Query" name:"State"`
	Lang     string `position:"Query" name:"Lang"`
}

// UpdateAppKeyStateResponse is the response struct for api UpdateAppKeyState
type UpdateAppKeyStateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateAppKeyStateRequest creates a request to invoke UpdateAppKeyState API
func CreateUpdateAppKeyStateRequest() (request *UpdateAppKeyStateRequest) {
	request = &UpdateAppKeyStateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateAppKeyState", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateAppKeyStateResponse creates a response to parse from UpdateAppKeyState response
func CreateUpdateAppKeyStateResponse() (response *UpdateAppKeyStateResponse) {
	response = &UpdateAppKeyStateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
