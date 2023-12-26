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

// DeleteAggregateCompliancePacks invokes the config.DeleteAggregateCompliancePacks API synchronously
func (client *Client) DeleteAggregateCompliancePacks(request *DeleteAggregateCompliancePacksRequest) (response *DeleteAggregateCompliancePacksResponse, err error) {
	response = CreateDeleteAggregateCompliancePacksResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAggregateCompliancePacksWithChan invokes the config.DeleteAggregateCompliancePacks API asynchronously
func (client *Client) DeleteAggregateCompliancePacksWithChan(request *DeleteAggregateCompliancePacksRequest) (<-chan *DeleteAggregateCompliancePacksResponse, <-chan error) {
	responseChan := make(chan *DeleteAggregateCompliancePacksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAggregateCompliancePacks(request)
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

// DeleteAggregateCompliancePacksWithCallback invokes the config.DeleteAggregateCompliancePacks API asynchronously
func (client *Client) DeleteAggregateCompliancePacksWithCallback(request *DeleteAggregateCompliancePacksRequest, callback func(response *DeleteAggregateCompliancePacksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAggregateCompliancePacksResponse
		var err error
		defer close(result)
		response, err = client.DeleteAggregateCompliancePacks(request)
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

// DeleteAggregateCompliancePacksRequest is the request struct for api DeleteAggregateCompliancePacks
type DeleteAggregateCompliancePacksRequest struct {
	*requests.RpcRequest
	ClientToken       string           `position:"Body" name:"ClientToken"`
	AggregatorId      string           `position:"Body" name:"AggregatorId"`
	CompliancePackIds string           `position:"Body" name:"CompliancePackIds"`
	DeleteRule        requests.Boolean `position:"Body" name:"DeleteRule"`
}

// DeleteAggregateCompliancePacksResponse is the response struct for api DeleteAggregateCompliancePacks
type DeleteAggregateCompliancePacksResponse struct {
	*responses.BaseResponse
	RequestId                    string                       `json:"RequestId" xml:"RequestId"`
	OperateCompliancePacksResult OperateCompliancePacksResult `json:"OperateCompliancePacksResult" xml:"OperateCompliancePacksResult"`
}

// CreateDeleteAggregateCompliancePacksRequest creates a request to invoke DeleteAggregateCompliancePacks API
func CreateDeleteAggregateCompliancePacksRequest() (request *DeleteAggregateCompliancePacksRequest) {
	request = &DeleteAggregateCompliancePacksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "DeleteAggregateCompliancePacks", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteAggregateCompliancePacksResponse creates a response to parse from DeleteAggregateCompliancePacks response
func CreateDeleteAggregateCompliancePacksResponse() (response *DeleteAggregateCompliancePacksResponse) {
	response = &DeleteAggregateCompliancePacksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
