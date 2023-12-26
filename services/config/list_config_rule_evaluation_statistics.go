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

// ListConfigRuleEvaluationStatistics invokes the config.ListConfigRuleEvaluationStatistics API synchronously
func (client *Client) ListConfigRuleEvaluationStatistics(request *ListConfigRuleEvaluationStatisticsRequest) (response *ListConfigRuleEvaluationStatisticsResponse, err error) {
	response = CreateListConfigRuleEvaluationStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// ListConfigRuleEvaluationStatisticsWithChan invokes the config.ListConfigRuleEvaluationStatistics API asynchronously
func (client *Client) ListConfigRuleEvaluationStatisticsWithChan(request *ListConfigRuleEvaluationStatisticsRequest) (<-chan *ListConfigRuleEvaluationStatisticsResponse, <-chan error) {
	responseChan := make(chan *ListConfigRuleEvaluationStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListConfigRuleEvaluationStatistics(request)
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

// ListConfigRuleEvaluationStatisticsWithCallback invokes the config.ListConfigRuleEvaluationStatistics API asynchronously
func (client *Client) ListConfigRuleEvaluationStatisticsWithCallback(request *ListConfigRuleEvaluationStatisticsRequest, callback func(response *ListConfigRuleEvaluationStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListConfigRuleEvaluationStatisticsResponse
		var err error
		defer close(result)
		response, err = client.ListConfigRuleEvaluationStatistics(request)
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

// ListConfigRuleEvaluationStatisticsRequest is the request struct for api ListConfigRuleEvaluationStatistics
type ListConfigRuleEvaluationStatisticsRequest struct {
	*requests.RpcRequest
}

// ListConfigRuleEvaluationStatisticsResponse is the response struct for api ListConfigRuleEvaluationStatistics
type ListConfigRuleEvaluationStatisticsResponse struct {
	*responses.BaseResponse
	RequestId         string                  `json:"RequestId" xml:"RequestId"`
	EvaluationResults []EvaluationResultsItem `json:"EvaluationResults" xml:"EvaluationResults"`
}

// CreateListConfigRuleEvaluationStatisticsRequest creates a request to invoke ListConfigRuleEvaluationStatistics API
func CreateListConfigRuleEvaluationStatisticsRequest() (request *ListConfigRuleEvaluationStatisticsRequest) {
	request = &ListConfigRuleEvaluationStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "ListConfigRuleEvaluationStatistics", "", "")
	request.Method = requests.POST
	return
}

// CreateListConfigRuleEvaluationStatisticsResponse creates a response to parse from ListConfigRuleEvaluationStatistics response
func CreateListConfigRuleEvaluationStatisticsResponse() (response *ListConfigRuleEvaluationStatisticsResponse) {
	response = &ListConfigRuleEvaluationStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
