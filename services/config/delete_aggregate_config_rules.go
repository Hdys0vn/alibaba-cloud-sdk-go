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

// DeleteAggregateConfigRules invokes the config.DeleteAggregateConfigRules API synchronously
func (client *Client) DeleteAggregateConfigRules(request *DeleteAggregateConfigRulesRequest) (response *DeleteAggregateConfigRulesResponse, err error) {
	response = CreateDeleteAggregateConfigRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAggregateConfigRulesWithChan invokes the config.DeleteAggregateConfigRules API asynchronously
func (client *Client) DeleteAggregateConfigRulesWithChan(request *DeleteAggregateConfigRulesRequest) (<-chan *DeleteAggregateConfigRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteAggregateConfigRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAggregateConfigRules(request)
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

// DeleteAggregateConfigRulesWithCallback invokes the config.DeleteAggregateConfigRules API asynchronously
func (client *Client) DeleteAggregateConfigRulesWithCallback(request *DeleteAggregateConfigRulesRequest, callback func(response *DeleteAggregateConfigRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAggregateConfigRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteAggregateConfigRules(request)
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

// DeleteAggregateConfigRulesRequest is the request struct for api DeleteAggregateConfigRules
type DeleteAggregateConfigRulesRequest struct {
	*requests.RpcRequest
	ConfigRuleIds string `position:"Query" name:"ConfigRuleIds"`
	AggregatorId  string `position:"Query" name:"AggregatorId"`
}

// DeleteAggregateConfigRulesResponse is the response struct for api DeleteAggregateConfigRules
type DeleteAggregateConfigRulesResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	OperateRuleResult OperateRuleResult `json:"OperateRuleResult" xml:"OperateRuleResult"`
}

// CreateDeleteAggregateConfigRulesRequest creates a request to invoke DeleteAggregateConfigRules API
func CreateDeleteAggregateConfigRulesRequest() (request *DeleteAggregateConfigRulesRequest) {
	request = &DeleteAggregateConfigRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "DeleteAggregateConfigRules", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteAggregateConfigRulesResponse creates a response to parse from DeleteAggregateConfigRules response
func CreateDeleteAggregateConfigRulesResponse() (response *DeleteAggregateConfigRulesResponse) {
	response = &DeleteAggregateConfigRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
