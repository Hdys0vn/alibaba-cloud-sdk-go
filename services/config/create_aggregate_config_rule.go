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

// CreateAggregateConfigRule invokes the config.CreateAggregateConfigRule API synchronously
func (client *Client) CreateAggregateConfigRule(request *CreateAggregateConfigRuleRequest) (response *CreateAggregateConfigRuleResponse, err error) {
	response = CreateCreateAggregateConfigRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAggregateConfigRuleWithChan invokes the config.CreateAggregateConfigRule API asynchronously
func (client *Client) CreateAggregateConfigRuleWithChan(request *CreateAggregateConfigRuleRequest) (<-chan *CreateAggregateConfigRuleResponse, <-chan error) {
	responseChan := make(chan *CreateAggregateConfigRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAggregateConfigRule(request)
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

// CreateAggregateConfigRuleWithCallback invokes the config.CreateAggregateConfigRule API asynchronously
func (client *Client) CreateAggregateConfigRuleWithCallback(request *CreateAggregateConfigRuleRequest, callback func(response *CreateAggregateConfigRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAggregateConfigRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateAggregateConfigRule(request)
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

// CreateAggregateConfigRuleRequest is the request struct for api CreateAggregateConfigRule
type CreateAggregateConfigRuleRequest struct {
	*requests.RpcRequest
	TagKeyScope               string           `position:"Body" name:"TagKeyScope"`
	ClientToken               string           `position:"Body" name:"ClientToken"`
	ResourceTypesScope        *[]string        `position:"Body" name:"ResourceTypesScope"  type:"Repeated"`
	Description               string           `position:"Body" name:"Description"`
	AggregatorId              string           `position:"Body" name:"AggregatorId"`
	ConfigRuleTriggerTypes    string           `position:"Body" name:"ConfigRuleTriggerTypes"`
	SourceIdentifier          string           `position:"Body" name:"SourceIdentifier"`
	TagValueScope             string           `position:"Body" name:"TagValueScope"`
	ExcludeAccountIdsScope    string           `position:"Body" name:"ExcludeAccountIdsScope"`
	RegionIdsScope            string           `position:"Body" name:"RegionIdsScope"`
	ExcludeFolderIdsScope     string           `position:"Body" name:"ExcludeFolderIdsScope"`
	RiskLevel                 requests.Integer `position:"Body" name:"RiskLevel"`
	SourceOwner               string           `position:"Body" name:"SourceOwner"`
	ResourceGroupIdsScope     string           `position:"Body" name:"ResourceGroupIdsScope"`
	InputParameters           string           `position:"Body" name:"InputParameters"`
	ConfigRuleName            string           `position:"Body" name:"ConfigRuleName"`
	TagKeyLogicScope          string           `position:"Body" name:"TagKeyLogicScope"`
	MaximumExecutionFrequency string           `position:"Body" name:"MaximumExecutionFrequency"`
	FolderIdsScope            string           `position:"Body" name:"FolderIdsScope"`
	ExcludeResourceIdsScope   string           `position:"Body" name:"ExcludeResourceIdsScope"`
	Conditions                string           `position:"Body" name:"Conditions"`
}

// CreateAggregateConfigRuleResponse is the response struct for api CreateAggregateConfigRule
type CreateAggregateConfigRuleResponse struct {
	*responses.BaseResponse
	ConfigRuleId string `json:"ConfigRuleId" xml:"ConfigRuleId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateAggregateConfigRuleRequest creates a request to invoke CreateAggregateConfigRule API
func CreateCreateAggregateConfigRuleRequest() (request *CreateAggregateConfigRuleRequest) {
	request = &CreateAggregateConfigRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "CreateAggregateConfigRule", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateAggregateConfigRuleResponse creates a response to parse from CreateAggregateConfigRule response
func CreateCreateAggregateConfigRuleResponse() (response *CreateAggregateConfigRuleResponse) {
	response = &CreateAggregateConfigRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
