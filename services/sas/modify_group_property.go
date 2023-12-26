package sas

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

// ModifyGroupProperty invokes the sas.ModifyGroupProperty API synchronously
func (client *Client) ModifyGroupProperty(request *ModifyGroupPropertyRequest) (response *ModifyGroupPropertyResponse, err error) {
	response = CreateModifyGroupPropertyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGroupPropertyWithChan invokes the sas.ModifyGroupProperty API asynchronously
func (client *Client) ModifyGroupPropertyWithChan(request *ModifyGroupPropertyRequest) (<-chan *ModifyGroupPropertyResponse, <-chan error) {
	responseChan := make(chan *ModifyGroupPropertyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGroupProperty(request)
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

// ModifyGroupPropertyWithCallback invokes the sas.ModifyGroupProperty API asynchronously
func (client *Client) ModifyGroupPropertyWithCallback(request *ModifyGroupPropertyRequest, callback func(response *ModifyGroupPropertyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGroupPropertyResponse
		var err error
		defer close(result)
		response, err = client.ModifyGroupProperty(request)
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

// ModifyGroupPropertyRequest is the request struct for api ModifyGroupProperty
type ModifyGroupPropertyRequest struct {
	*requests.RpcRequest
	Data     string `position:"Query" name:"Data"`
	SourceIp string `position:"Query" name:"SourceIp"`
}

// ModifyGroupPropertyResponse is the response struct for api ModifyGroupProperty
type ModifyGroupPropertyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyGroupPropertyRequest creates a request to invoke ModifyGroupProperty API
func CreateModifyGroupPropertyRequest() (request *ModifyGroupPropertyRequest) {
	request = &ModifyGroupPropertyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyGroupProperty", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyGroupPropertyResponse creates a response to parse from ModifyGroupProperty response
func CreateModifyGroupPropertyResponse() (response *ModifyGroupPropertyResponse) {
	response = &ModifyGroupPropertyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
