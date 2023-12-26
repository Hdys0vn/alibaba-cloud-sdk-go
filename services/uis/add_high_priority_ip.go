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

// AddHighPriorityIp invokes the uis.AddHighPriorityIp API synchronously
// api document: https://help.aliyun.com/api/uis/addhighpriorityip.html
func (client *Client) AddHighPriorityIp(request *AddHighPriorityIpRequest) (response *AddHighPriorityIpResponse, err error) {
	response = CreateAddHighPriorityIpResponse()
	err = client.DoAction(request, response)
	return
}

// AddHighPriorityIpWithChan invokes the uis.AddHighPriorityIp API asynchronously
// api document: https://help.aliyun.com/api/uis/addhighpriorityip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddHighPriorityIpWithChan(request *AddHighPriorityIpRequest) (<-chan *AddHighPriorityIpResponse, <-chan error) {
	responseChan := make(chan *AddHighPriorityIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddHighPriorityIp(request)
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

// AddHighPriorityIpWithCallback invokes the uis.AddHighPriorityIp API asynchronously
// api document: https://help.aliyun.com/api/uis/addhighpriorityip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddHighPriorityIpWithCallback(request *AddHighPriorityIpRequest, callback func(response *AddHighPriorityIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddHighPriorityIpResponse
		var err error
		defer close(result)
		response, err = client.AddHighPriorityIp(request)
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

// AddHighPriorityIpRequest is the request struct for api AddHighPriorityIp
type AddHighPriorityIpRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	HighPriorityIp       string           `position:"Query" name:"HighPriorityIp"`
	UisId                string           `position:"Query" name:"UisId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// AddHighPriorityIpResponse is the response struct for api AddHighPriorityIp
type AddHighPriorityIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddHighPriorityIpRequest creates a request to invoke AddHighPriorityIp API
func CreateAddHighPriorityIpRequest() (request *AddHighPriorityIpRequest) {
	request = &AddHighPriorityIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "AddHighPriorityIp", "uis", "openAPI")
	return
}

// CreateAddHighPriorityIpResponse creates a response to parse from AddHighPriorityIp response
func CreateAddHighPriorityIpResponse() (response *AddHighPriorityIpResponse) {
	response = &AddHighPriorityIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
