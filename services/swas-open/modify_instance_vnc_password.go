package swas_open

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

// ModifyInstanceVncPassword invokes the swas_open.ModifyInstanceVncPassword API synchronously
func (client *Client) ModifyInstanceVncPassword(request *ModifyInstanceVncPasswordRequest) (response *ModifyInstanceVncPasswordResponse, err error) {
	response = CreateModifyInstanceVncPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceVncPasswordWithChan invokes the swas_open.ModifyInstanceVncPassword API asynchronously
func (client *Client) ModifyInstanceVncPasswordWithChan(request *ModifyInstanceVncPasswordRequest) (<-chan *ModifyInstanceVncPasswordResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceVncPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceVncPassword(request)
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

// ModifyInstanceVncPasswordWithCallback invokes the swas_open.ModifyInstanceVncPassword API asynchronously
func (client *Client) ModifyInstanceVncPasswordWithCallback(request *ModifyInstanceVncPasswordRequest, callback func(response *ModifyInstanceVncPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceVncPasswordResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceVncPassword(request)
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

// ModifyInstanceVncPasswordRequest is the request struct for api ModifyInstanceVncPassword
type ModifyInstanceVncPasswordRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	VncPassword string `position:"Query" name:"VncPassword"`
}

// ModifyInstanceVncPasswordResponse is the response struct for api ModifyInstanceVncPassword
type ModifyInstanceVncPasswordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceVncPasswordRequest creates a request to invoke ModifyInstanceVncPassword API
func CreateModifyInstanceVncPasswordRequest() (request *ModifyInstanceVncPasswordRequest) {
	request = &ModifyInstanceVncPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "ModifyInstanceVncPassword", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceVncPasswordResponse creates a response to parse from ModifyInstanceVncPassword response
func CreateModifyInstanceVncPasswordResponse() (response *ModifyInstanceVncPasswordResponse) {
	response = &ModifyInstanceVncPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
