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

// GetAggregateCompliancePack invokes the config.GetAggregateCompliancePack API synchronously
func (client *Client) GetAggregateCompliancePack(request *GetAggregateCompliancePackRequest) (response *GetAggregateCompliancePackResponse, err error) {
	response = CreateGetAggregateCompliancePackResponse()
	err = client.DoAction(request, response)
	return
}

// GetAggregateCompliancePackWithChan invokes the config.GetAggregateCompliancePack API asynchronously
func (client *Client) GetAggregateCompliancePackWithChan(request *GetAggregateCompliancePackRequest) (<-chan *GetAggregateCompliancePackResponse, <-chan error) {
	responseChan := make(chan *GetAggregateCompliancePackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAggregateCompliancePack(request)
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

// GetAggregateCompliancePackWithCallback invokes the config.GetAggregateCompliancePack API asynchronously
func (client *Client) GetAggregateCompliancePackWithCallback(request *GetAggregateCompliancePackRequest, callback func(response *GetAggregateCompliancePackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAggregateCompliancePackResponse
		var err error
		defer close(result)
		response, err = client.GetAggregateCompliancePack(request)
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

// GetAggregateCompliancePackRequest is the request struct for api GetAggregateCompliancePack
type GetAggregateCompliancePackRequest struct {
	*requests.RpcRequest
	AggregatorId     string `position:"Query" name:"AggregatorId"`
	CompliancePackId string `position:"Query" name:"CompliancePackId"`
}

// GetAggregateCompliancePackResponse is the response struct for api GetAggregateCompliancePack
type GetAggregateCompliancePackResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	CompliancePack CompliancePack `json:"CompliancePack" xml:"CompliancePack"`
}

// CreateGetAggregateCompliancePackRequest creates a request to invoke GetAggregateCompliancePack API
func CreateGetAggregateCompliancePackRequest() (request *GetAggregateCompliancePackRequest) {
	request = &GetAggregateCompliancePackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GetAggregateCompliancePack", "", "")
	request.Method = requests.GET
	return
}

// CreateGetAggregateCompliancePackResponse creates a response to parse from GetAggregateCompliancePack response
func CreateGetAggregateCompliancePackResponse() (response *GetAggregateCompliancePackResponse) {
	response = &GetAggregateCompliancePackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
