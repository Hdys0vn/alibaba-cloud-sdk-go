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

// ListConfigRuleEvaluationResults invokes the config.ListConfigRuleEvaluationResults API synchronously
func (client *Client) ListConfigRuleEvaluationResults(request *ListConfigRuleEvaluationResultsRequest) (response *ListConfigRuleEvaluationResultsResponse, err error) {
	response = CreateListConfigRuleEvaluationResultsResponse()
	err = client.DoAction(request, response)
	return
}

// ListConfigRuleEvaluationResultsWithChan invokes the config.ListConfigRuleEvaluationResults API asynchronously
func (client *Client) ListConfigRuleEvaluationResultsWithChan(request *ListConfigRuleEvaluationResultsRequest) (<-chan *ListConfigRuleEvaluationResultsResponse, <-chan error) {
	responseChan := make(chan *ListConfigRuleEvaluationResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListConfigRuleEvaluationResults(request)
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

// ListConfigRuleEvaluationResultsWithCallback invokes the config.ListConfigRuleEvaluationResults API asynchronously
func (client *Client) ListConfigRuleEvaluationResultsWithCallback(request *ListConfigRuleEvaluationResultsRequest, callback func(response *ListConfigRuleEvaluationResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListConfigRuleEvaluationResultsResponse
		var err error
		defer close(result)
		response, err = client.ListConfigRuleEvaluationResults(request)
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

// ListConfigRuleEvaluationResultsRequest is the request struct for api ListConfigRuleEvaluationResults
type ListConfigRuleEvaluationResultsRequest struct {
	*requests.RpcRequest
	ConfigRuleId     string           `position:"Query" name:"ConfigRuleId"`
	Regions          string           `position:"Query" name:"Regions"`
	NextToken        string           `position:"Query" name:"NextToken"`
	CompliancePackId string           `position:"Query" name:"CompliancePackId"`
	ComplianceType   string           `position:"Query" name:"ComplianceType"`
	ResourceTypes    string           `position:"Query" name:"ResourceTypes"`
	ResourceGroupIds string           `position:"Query" name:"ResourceGroupIds"`
	MaxResults       requests.Integer `position:"Query" name:"MaxResults"`
}

// ListConfigRuleEvaluationResultsResponse is the response struct for api ListConfigRuleEvaluationResults
type ListConfigRuleEvaluationResultsResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	EvaluationResults EvaluationResults `json:"EvaluationResults" xml:"EvaluationResults"`
}

// CreateListConfigRuleEvaluationResultsRequest creates a request to invoke ListConfigRuleEvaluationResults API
func CreateListConfigRuleEvaluationResultsRequest() (request *ListConfigRuleEvaluationResultsRequest) {
	request = &ListConfigRuleEvaluationResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "ListConfigRuleEvaluationResults", "", "")
	request.Method = requests.GET
	return
}

// CreateListConfigRuleEvaluationResultsResponse creates a response to parse from ListConfigRuleEvaluationResults response
func CreateListConfigRuleEvaluationResultsResponse() (response *ListConfigRuleEvaluationResultsResponse) {
	response = &ListConfigRuleEvaluationResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
