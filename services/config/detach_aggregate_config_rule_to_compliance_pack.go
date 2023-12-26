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

// DetachAggregateConfigRuleToCompliancePack invokes the config.DetachAggregateConfigRuleToCompliancePack API synchronously
func (client *Client) DetachAggregateConfigRuleToCompliancePack(request *DetachAggregateConfigRuleToCompliancePackRequest) (response *DetachAggregateConfigRuleToCompliancePackResponse, err error) {
	response = CreateDetachAggregateConfigRuleToCompliancePackResponse()
	err = client.DoAction(request, response)
	return
}

// DetachAggregateConfigRuleToCompliancePackWithChan invokes the config.DetachAggregateConfigRuleToCompliancePack API asynchronously
func (client *Client) DetachAggregateConfigRuleToCompliancePackWithChan(request *DetachAggregateConfigRuleToCompliancePackRequest) (<-chan *DetachAggregateConfigRuleToCompliancePackResponse, <-chan error) {
	responseChan := make(chan *DetachAggregateConfigRuleToCompliancePackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachAggregateConfigRuleToCompliancePack(request)
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

// DetachAggregateConfigRuleToCompliancePackWithCallback invokes the config.DetachAggregateConfigRuleToCompliancePack API asynchronously
func (client *Client) DetachAggregateConfigRuleToCompliancePackWithCallback(request *DetachAggregateConfigRuleToCompliancePackRequest, callback func(response *DetachAggregateConfigRuleToCompliancePackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachAggregateConfigRuleToCompliancePackResponse
		var err error
		defer close(result)
		response, err = client.DetachAggregateConfigRuleToCompliancePack(request)
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

// DetachAggregateConfigRuleToCompliancePackRequest is the request struct for api DetachAggregateConfigRuleToCompliancePack
type DetachAggregateConfigRuleToCompliancePackRequest struct {
	*requests.RpcRequest
	ConfigRuleIds    string `position:"Query" name:"ConfigRuleIds"`
	AggregatorId     string `position:"Query" name:"AggregatorId"`
	CompliancePackId string `position:"Query" name:"CompliancePackId"`
}

// DetachAggregateConfigRuleToCompliancePackResponse is the response struct for api DetachAggregateConfigRuleToCompliancePack
type DetachAggregateConfigRuleToCompliancePackResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	OperateRuleResult OperateRuleResult `json:"OperateRuleResult" xml:"OperateRuleResult"`
}

// CreateDetachAggregateConfigRuleToCompliancePackRequest creates a request to invoke DetachAggregateConfigRuleToCompliancePack API
func CreateDetachAggregateConfigRuleToCompliancePackRequest() (request *DetachAggregateConfigRuleToCompliancePackRequest) {
	request = &DetachAggregateConfigRuleToCompliancePackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "DetachAggregateConfigRuleToCompliancePack", "", "")
	request.Method = requests.POST
	return
}

// CreateDetachAggregateConfigRuleToCompliancePackResponse creates a response to parse from DetachAggregateConfigRuleToCompliancePack response
func CreateDetachAggregateConfigRuleToCompliancePackResponse() (response *DetachAggregateConfigRuleToCompliancePackResponse) {
	response = &DetachAggregateConfigRuleToCompliancePackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
