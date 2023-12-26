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

// DeleteHighPriorityIp invokes the uis.DeleteHighPriorityIp API synchronously
// api document: https://help.aliyun.com/api/uis/deletehighpriorityip.html
func (client *Client) DeleteHighPriorityIp(request *DeleteHighPriorityIpRequest) (response *DeleteHighPriorityIpResponse, err error) {
	response = CreateDeleteHighPriorityIpResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteHighPriorityIpWithChan invokes the uis.DeleteHighPriorityIp API asynchronously
// api document: https://help.aliyun.com/api/uis/deletehighpriorityip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHighPriorityIpWithChan(request *DeleteHighPriorityIpRequest) (<-chan *DeleteHighPriorityIpResponse, <-chan error) {
	responseChan := make(chan *DeleteHighPriorityIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHighPriorityIp(request)
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

// DeleteHighPriorityIpWithCallback invokes the uis.DeleteHighPriorityIp API asynchronously
// api document: https://help.aliyun.com/api/uis/deletehighpriorityip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHighPriorityIpWithCallback(request *DeleteHighPriorityIpRequest, callback func(response *DeleteHighPriorityIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHighPriorityIpResponse
		var err error
		defer close(result)
		response, err = client.DeleteHighPriorityIp(request)
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

// DeleteHighPriorityIpRequest is the request struct for api DeleteHighPriorityIp
type DeleteHighPriorityIpRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	HighPriorityIp       string           `position:"Query" name:"HighPriorityIp"`
	UisId                string           `position:"Query" name:"UisId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteHighPriorityIpResponse is the response struct for api DeleteHighPriorityIp
type DeleteHighPriorityIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteHighPriorityIpRequest creates a request to invoke DeleteHighPriorityIp API
func CreateDeleteHighPriorityIpRequest() (request *DeleteHighPriorityIpRequest) {
	request = &DeleteHighPriorityIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "DeleteHighPriorityIp", "uis", "openAPI")
	return
}

// CreateDeleteHighPriorityIpResponse creates a response to parse from DeleteHighPriorityIp response
func CreateDeleteHighPriorityIpResponse() (response *DeleteHighPriorityIpResponse) {
	response = &DeleteHighPriorityIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
