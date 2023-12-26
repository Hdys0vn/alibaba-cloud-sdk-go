package uis

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

// ModifyHighPriorityIp invokes the uis.ModifyHighPriorityIp API synchronously
// api document: https://help.aliyun.com/api/uis/modifyhighpriorityip.html
func (client *Client) ModifyHighPriorityIp(request *ModifyHighPriorityIpRequest) (response *ModifyHighPriorityIpResponse, err error) {
	response = CreateModifyHighPriorityIpResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyHighPriorityIpWithChan invokes the uis.ModifyHighPriorityIp API asynchronously
// api document: https://help.aliyun.com/api/uis/modifyhighpriorityip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHighPriorityIpWithChan(request *ModifyHighPriorityIpRequest) (<-chan *ModifyHighPriorityIpResponse, <-chan error) {
	responseChan := make(chan *ModifyHighPriorityIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyHighPriorityIp(request)
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

// ModifyHighPriorityIpWithCallback invokes the uis.ModifyHighPriorityIp API asynchronously
// api document: https://help.aliyun.com/api/uis/modifyhighpriorityip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHighPriorityIpWithCallback(request *ModifyHighPriorityIpRequest, callback func(response *ModifyHighPriorityIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyHighPriorityIpResponse
		var err error
		defer close(result)
		response, err = client.ModifyHighPriorityIp(request)
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

// ModifyHighPriorityIpRequest is the request struct for api ModifyHighPriorityIp
type ModifyHighPriorityIpRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	HighPriorityIp       string           `position:"Query" name:"HighPriorityIp"`
	UisId                string           `position:"Query" name:"UisId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyHighPriorityIpResponse is the response struct for api ModifyHighPriorityIp
type ModifyHighPriorityIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyHighPriorityIpRequest creates a request to invoke ModifyHighPriorityIp API
func CreateModifyHighPriorityIpRequest() (request *ModifyHighPriorityIpRequest) {
	request = &ModifyHighPriorityIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "ModifyHighPriorityIp", "uis", "openAPI")
	return
}

// CreateModifyHighPriorityIpResponse creates a response to parse from ModifyHighPriorityIp response
func CreateModifyHighPriorityIpResponse() (response *ModifyHighPriorityIpResponse) {
	response = &ModifyHighPriorityIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
