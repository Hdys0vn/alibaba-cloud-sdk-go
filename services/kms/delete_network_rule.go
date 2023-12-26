package kms

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

// DeleteNetworkRule invokes the kms.DeleteNetworkRule API synchronously
func (client *Client) DeleteNetworkRule(request *DeleteNetworkRuleRequest) (response *DeleteNetworkRuleResponse, err error) {
	response = CreateDeleteNetworkRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNetworkRuleWithChan invokes the kms.DeleteNetworkRule API asynchronously
func (client *Client) DeleteNetworkRuleWithChan(request *DeleteNetworkRuleRequest) (<-chan *DeleteNetworkRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteNetworkRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNetworkRule(request)
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

// DeleteNetworkRuleWithCallback invokes the kms.DeleteNetworkRule API asynchronously
func (client *Client) DeleteNetworkRuleWithCallback(request *DeleteNetworkRuleRequest, callback func(response *DeleteNetworkRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNetworkRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteNetworkRule(request)
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

// DeleteNetworkRuleRequest is the request struct for api DeleteNetworkRule
type DeleteNetworkRuleRequest struct {
	*requests.RpcRequest
	Name string `position:"Query" name:"Name"`
}

// DeleteNetworkRuleResponse is the response struct for api DeleteNetworkRule
type DeleteNetworkRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteNetworkRuleRequest creates a request to invoke DeleteNetworkRule API
func CreateDeleteNetworkRuleRequest() (request *DeleteNetworkRuleRequest) {
	request = &DeleteNetworkRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "DeleteNetworkRule", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteNetworkRuleResponse creates a response to parse from DeleteNetworkRule response
func CreateDeleteNetworkRuleResponse() (response *DeleteNetworkRuleResponse) {
	response = &DeleteNetworkRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
