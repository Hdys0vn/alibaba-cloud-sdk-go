package ahas_openapi

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

// DisableFlowRule invokes the ahas_openapi.DisableFlowRule API synchronously
func (client *Client) DisableFlowRule(request *DisableFlowRuleRequest) (response *DisableFlowRuleResponse, err error) {
	response = CreateDisableFlowRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DisableFlowRuleWithChan invokes the ahas_openapi.DisableFlowRule API asynchronously
func (client *Client) DisableFlowRuleWithChan(request *DisableFlowRuleRequest) (<-chan *DisableFlowRuleResponse, <-chan error) {
	responseChan := make(chan *DisableFlowRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableFlowRule(request)
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

// DisableFlowRuleWithCallback invokes the ahas_openapi.DisableFlowRule API asynchronously
func (client *Client) DisableFlowRuleWithCallback(request *DisableFlowRuleRequest, callback func(response *DisableFlowRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableFlowRuleResponse
		var err error
		defer close(result)
		response, err = client.DisableFlowRule(request)
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

// DisableFlowRuleRequest is the request struct for api DisableFlowRule
type DisableFlowRuleRequest struct {
	*requests.RpcRequest
	AhasRegionId string           `position:"Query" name:"AhasRegionId"`
	RuleId       requests.Integer `position:"Query" name:"RuleId"`
}

// DisableFlowRuleResponse is the response struct for api DisableFlowRule
type DisableFlowRuleResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDisableFlowRuleRequest creates a request to invoke DisableFlowRule API
func CreateDisableFlowRuleRequest() (request *DisableFlowRuleRequest) {
	request = &DisableFlowRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "DisableFlowRule", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableFlowRuleResponse creates a response to parse from DisableFlowRule response
func CreateDisableFlowRuleResponse() (response *DisableFlowRuleResponse) {
	response = &DisableFlowRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
