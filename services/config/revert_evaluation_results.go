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

// RevertEvaluationResults invokes the config.RevertEvaluationResults API synchronously
func (client *Client) RevertEvaluationResults(request *RevertEvaluationResultsRequest) (response *RevertEvaluationResultsResponse, err error) {
	response = CreateRevertEvaluationResultsResponse()
	err = client.DoAction(request, response)
	return
}

// RevertEvaluationResultsWithChan invokes the config.RevertEvaluationResults API asynchronously
func (client *Client) RevertEvaluationResultsWithChan(request *RevertEvaluationResultsRequest) (<-chan *RevertEvaluationResultsResponse, <-chan error) {
	responseChan := make(chan *RevertEvaluationResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevertEvaluationResults(request)
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

// RevertEvaluationResultsWithCallback invokes the config.RevertEvaluationResults API asynchronously
func (client *Client) RevertEvaluationResultsWithCallback(request *RevertEvaluationResultsRequest, callback func(response *RevertEvaluationResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevertEvaluationResultsResponse
		var err error
		defer close(result)
		response, err = client.RevertEvaluationResults(request)
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

// RevertEvaluationResultsRequest is the request struct for api RevertEvaluationResults
type RevertEvaluationResultsRequest struct {
	*requests.RpcRequest
	ConfigRuleId string                              `position:"Body" name:"ConfigRuleId"`
	Resources    *[]RevertEvaluationResultsResources `position:"Body" name:"Resources"  type:"Json"`
}

// RevertEvaluationResultsResources is a repeated param struct in RevertEvaluationResultsRequest
type RevertEvaluationResultsResources struct {
	ResourceId        string `name:"ResourceId"`
	ResourceAccountId string `name:"ResourceAccountId"`
	Region            string `name:"Region"`
	ResourceType      string `name:"ResourceType"`
}

// RevertEvaluationResultsResponse is the response struct for api RevertEvaluationResults
type RevertEvaluationResultsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRevertEvaluationResultsRequest creates a request to invoke RevertEvaluationResults API
func CreateRevertEvaluationResultsRequest() (request *RevertEvaluationResultsRequest) {
	request = &RevertEvaluationResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "RevertEvaluationResults", "", "")
	request.Method = requests.POST
	return
}

// CreateRevertEvaluationResultsResponse creates a response to parse from RevertEvaluationResults response
func CreateRevertEvaluationResultsResponse() (response *RevertEvaluationResultsResponse) {
	response = &RevertEvaluationResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
