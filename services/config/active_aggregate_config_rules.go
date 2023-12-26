package config

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

// ActiveAggregateConfigRules invokes the config.ActiveAggregateConfigRules API synchronously
func (client *Client) ActiveAggregateConfigRules(request *ActiveAggregateConfigRulesRequest) (response *ActiveAggregateConfigRulesResponse, err error) {
	response = CreateActiveAggregateConfigRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ActiveAggregateConfigRulesWithChan invokes the config.ActiveAggregateConfigRules API asynchronously
func (client *Client) ActiveAggregateConfigRulesWithChan(request *ActiveAggregateConfigRulesRequest) (<-chan *ActiveAggregateConfigRulesResponse, <-chan error) {
	responseChan := make(chan *ActiveAggregateConfigRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ActiveAggregateConfigRules(request)
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

// ActiveAggregateConfigRulesWithCallback invokes the config.ActiveAggregateConfigRules API asynchronously
func (client *Client) ActiveAggregateConfigRulesWithCallback(request *ActiveAggregateConfigRulesRequest, callback func(response *ActiveAggregateConfigRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ActiveAggregateConfigRulesResponse
		var err error
		defer close(result)
		response, err = client.ActiveAggregateConfigRules(request)
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

// ActiveAggregateConfigRulesRequest is the request struct for api ActiveAggregateConfigRules
type ActiveAggregateConfigRulesRequest struct {
	*requests.RpcRequest
	ConfigRuleIds string `position:"Query" name:"ConfigRuleIds"`
	AggregatorId  string `position:"Query" name:"AggregatorId"`
}

// ActiveAggregateConfigRulesResponse is the response struct for api ActiveAggregateConfigRules
type ActiveAggregateConfigRulesResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	OperateRuleResult OperateRuleResult `json:"OperateRuleResult" xml:"OperateRuleResult"`
}

// CreateActiveAggregateConfigRulesRequest creates a request to invoke ActiveAggregateConfigRules API
func CreateActiveAggregateConfigRulesRequest() (request *ActiveAggregateConfigRulesRequest) {
	request = &ActiveAggregateConfigRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "ActiveAggregateConfigRules", "", "")
	request.Method = requests.POST
	return
}

// CreateActiveAggregateConfigRulesResponse creates a response to parse from ActiveAggregateConfigRules response
func CreateActiveAggregateConfigRulesResponse() (response *ActiveAggregateConfigRulesResponse) {
	response = &ActiveAggregateConfigRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
