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

// ModifySystemRule invokes the ahas_openapi.ModifySystemRule API synchronously
func (client *Client) ModifySystemRule(request *ModifySystemRuleRequest) (response *ModifySystemRuleResponse, err error) {
	response = CreateModifySystemRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySystemRuleWithChan invokes the ahas_openapi.ModifySystemRule API asynchronously
func (client *Client) ModifySystemRuleWithChan(request *ModifySystemRuleRequest) (<-chan *ModifySystemRuleResponse, <-chan error) {
	responseChan := make(chan *ModifySystemRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySystemRule(request)
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

// ModifySystemRuleWithCallback invokes the ahas_openapi.ModifySystemRule API asynchronously
func (client *Client) ModifySystemRuleWithCallback(request *ModifySystemRuleRequest, callback func(response *ModifySystemRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySystemRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifySystemRule(request)
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

// ModifySystemRuleRequest is the request struct for api ModifySystemRule
type ModifySystemRuleRequest struct {
	*requests.RpcRequest
	Threshold    requests.Float   `position:"Query" name:"Threshold"`
	AhasRegionId string           `position:"Query" name:"AhasRegionId"`
	RuleId       requests.Integer `position:"Query" name:"RuleId"`
}

// ModifySystemRuleResponse is the response struct for api ModifySystemRule
type ModifySystemRuleResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateModifySystemRuleRequest creates a request to invoke ModifySystemRule API
func CreateModifySystemRuleRequest() (request *ModifySystemRuleRequest) {
	request = &ModifySystemRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "ModifySystemRule", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySystemRuleResponse creates a response to parse from ModifySystemRule response
func CreateModifySystemRuleResponse() (response *ModifySystemRuleResponse) {
	response = &ModifySystemRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
