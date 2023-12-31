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

// ModifyDegradeRule invokes the ahas_openapi.ModifyDegradeRule API synchronously
func (client *Client) ModifyDegradeRule(request *ModifyDegradeRuleRequest) (response *ModifyDegradeRuleResponse, err error) {
	response = CreateModifyDegradeRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDegradeRuleWithChan invokes the ahas_openapi.ModifyDegradeRule API asynchronously
func (client *Client) ModifyDegradeRuleWithChan(request *ModifyDegradeRuleRequest) (<-chan *ModifyDegradeRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyDegradeRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDegradeRule(request)
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

// ModifyDegradeRuleWithCallback invokes the ahas_openapi.ModifyDegradeRule API asynchronously
func (client *Client) ModifyDegradeRuleWithCallback(request *ModifyDegradeRuleRequest, callback func(response *ModifyDegradeRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDegradeRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyDegradeRule(request)
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

// ModifyDegradeRuleRequest is the request struct for api ModifyDegradeRule
type ModifyDegradeRuleRequest struct {
	*requests.RpcRequest
	RecoveryTimeoutMs         requests.Integer `position:"Query" name:"RecoveryTimeoutMs"`
	Threshold                 requests.Float   `position:"Query" name:"Threshold"`
	AhasRegionId              string           `position:"Query" name:"AhasRegionId"`
	HalfOpenBaseAmountPerStep requests.Integer `position:"Query" name:"HalfOpenBaseAmountPerStep"`
	StatDurationMs            requests.Integer `position:"Query" name:"StatDurationMs"`
	MinRequestAmount          requests.Integer `position:"Query" name:"MinRequestAmount"`
	HalfOpenRecoveryStepNum   requests.Integer `position:"Query" name:"HalfOpenRecoveryStepNum"`
	Strategy                  requests.Integer `position:"Query" name:"Strategy"`
	RuleId                    requests.Integer `position:"Query" name:"RuleId"`
	SlowRtMs                  requests.Integer `position:"Query" name:"SlowRtMs"`
}

// ModifyDegradeRuleResponse is the response struct for api ModifyDegradeRule
type ModifyDegradeRuleResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateModifyDegradeRuleRequest creates a request to invoke ModifyDegradeRule API
func CreateModifyDegradeRuleRequest() (request *ModifyDegradeRuleRequest) {
	request = &ModifyDegradeRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "ModifyDegradeRule", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDegradeRuleResponse creates a response to parse from ModifyDegradeRule response
func CreateModifyDegradeRuleResponse() (response *ModifyDegradeRuleResponse) {
	response = &ModifyDegradeRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
