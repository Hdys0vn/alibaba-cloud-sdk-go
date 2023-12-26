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

// StartAggregateRemediation invokes the config.StartAggregateRemediation API synchronously
func (client *Client) StartAggregateRemediation(request *StartAggregateRemediationRequest) (response *StartAggregateRemediationResponse, err error) {
	response = CreateStartAggregateRemediationResponse()
	err = client.DoAction(request, response)
	return
}

// StartAggregateRemediationWithChan invokes the config.StartAggregateRemediation API asynchronously
func (client *Client) StartAggregateRemediationWithChan(request *StartAggregateRemediationRequest) (<-chan *StartAggregateRemediationResponse, <-chan error) {
	responseChan := make(chan *StartAggregateRemediationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartAggregateRemediation(request)
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

// StartAggregateRemediationWithCallback invokes the config.StartAggregateRemediation API asynchronously
func (client *Client) StartAggregateRemediationWithCallback(request *StartAggregateRemediationRequest, callback func(response *StartAggregateRemediationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartAggregateRemediationResponse
		var err error
		defer close(result)
		response, err = client.StartAggregateRemediation(request)
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

// StartAggregateRemediationRequest is the request struct for api StartAggregateRemediation
type StartAggregateRemediationRequest struct {
	*requests.RpcRequest
	ConfigRuleId      string           `position:"Query" name:"ConfigRuleId"`
	ResourceOwnerId   string           `position:"Query" name:"ResourceOwnerId"`
	AggregatorId      string           `position:"Query" name:"AggregatorId"`
	ResourceAccountId requests.Integer `position:"Query" name:"ResourceAccountId"`
	ResourceRegionId  string           `position:"Query" name:"ResourceRegionId"`
	ResourceId        string           `position:"Query" name:"ResourceId"`
	ResourceType      string           `position:"Query" name:"ResourceType"`
}

// StartAggregateRemediationResponse is the response struct for api StartAggregateRemediation
type StartAggregateRemediationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateStartAggregateRemediationRequest creates a request to invoke StartAggregateRemediation API
func CreateStartAggregateRemediationRequest() (request *StartAggregateRemediationRequest) {
	request = &StartAggregateRemediationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "StartAggregateRemediation", "", "")
	request.Method = requests.POST
	return
}

// CreateStartAggregateRemediationResponse creates a response to parse from StartAggregateRemediation response
func CreateStartAggregateRemediationResponse() (response *StartAggregateRemediationResponse) {
	response = &StartAggregateRemediationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
