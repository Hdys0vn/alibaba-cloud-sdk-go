package cc5g

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

// DeleteGroupAuthorizationRule invokes the cc5g.DeleteGroupAuthorizationRule API synchronously
func (client *Client) DeleteGroupAuthorizationRule(request *DeleteGroupAuthorizationRuleRequest) (response *DeleteGroupAuthorizationRuleResponse, err error) {
	response = CreateDeleteGroupAuthorizationRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGroupAuthorizationRuleWithChan invokes the cc5g.DeleteGroupAuthorizationRule API asynchronously
func (client *Client) DeleteGroupAuthorizationRuleWithChan(request *DeleteGroupAuthorizationRuleRequest) (<-chan *DeleteGroupAuthorizationRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteGroupAuthorizationRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGroupAuthorizationRule(request)
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

// DeleteGroupAuthorizationRuleWithCallback invokes the cc5g.DeleteGroupAuthorizationRule API asynchronously
func (client *Client) DeleteGroupAuthorizationRuleWithCallback(request *DeleteGroupAuthorizationRuleRequest, callback func(response *DeleteGroupAuthorizationRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGroupAuthorizationRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteGroupAuthorizationRule(request)
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

// DeleteGroupAuthorizationRuleRequest is the request struct for api DeleteGroupAuthorizationRule
type DeleteGroupAuthorizationRuleRequest struct {
	*requests.RpcRequest
	WirelessCloudConnectorGroupId string           `position:"Query" name:"WirelessCloudConnectorGroupId"`
	DryRun                        requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
	AuthorizationRuleId           string           `position:"Query" name:"AuthorizationRuleId"`
}

// DeleteGroupAuthorizationRuleResponse is the response struct for api DeleteGroupAuthorizationRule
type DeleteGroupAuthorizationRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteGroupAuthorizationRuleRequest creates a request to invoke DeleteGroupAuthorizationRule API
func CreateDeleteGroupAuthorizationRuleRequest() (request *DeleteGroupAuthorizationRuleRequest) {
	request = &DeleteGroupAuthorizationRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "DeleteGroupAuthorizationRule", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGroupAuthorizationRuleResponse creates a response to parse from DeleteGroupAuthorizationRule response
func CreateDeleteGroupAuthorizationRuleResponse() (response *DeleteGroupAuthorizationRuleResponse) {
	response = &DeleteGroupAuthorizationRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}