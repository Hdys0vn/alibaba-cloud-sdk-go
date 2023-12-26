package smc

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

// ModifySourceServerAttribute invokes the smc.ModifySourceServerAttribute API synchronously
func (client *Client) ModifySourceServerAttribute(request *ModifySourceServerAttributeRequest) (response *ModifySourceServerAttributeResponse, err error) {
	response = CreateModifySourceServerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySourceServerAttributeWithChan invokes the smc.ModifySourceServerAttribute API asynchronously
func (client *Client) ModifySourceServerAttributeWithChan(request *ModifySourceServerAttributeRequest) (<-chan *ModifySourceServerAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifySourceServerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySourceServerAttribute(request)
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

// ModifySourceServerAttributeWithCallback invokes the smc.ModifySourceServerAttribute API asynchronously
func (client *Client) ModifySourceServerAttributeWithCallback(request *ModifySourceServerAttributeRequest, callback func(response *ModifySourceServerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySourceServerAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifySourceServerAttribute(request)
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

// ModifySourceServerAttributeRequest is the request struct for api ModifySourceServerAttribute
type ModifySourceServerAttributeRequest struct {
	*requests.RpcRequest
	Description          string           `position:"Query" name:"Description"`
	SourceId             string           `position:"Query" name:"SourceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifySourceServerAttributeResponse is the response struct for api ModifySourceServerAttribute
type ModifySourceServerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySourceServerAttributeRequest creates a request to invoke ModifySourceServerAttribute API
func CreateModifySourceServerAttributeRequest() (request *ModifySourceServerAttributeRequest) {
	request = &ModifySourceServerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("smc", "2019-06-01", "ModifySourceServerAttribute", "smc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySourceServerAttributeResponse creates a response to parse from ModifySourceServerAttribute response
func CreateModifySourceServerAttributeResponse() (response *ModifySourceServerAttributeResponse) {
	response = &ModifySourceServerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
