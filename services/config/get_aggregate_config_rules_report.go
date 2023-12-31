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

// GetAggregateConfigRulesReport invokes the config.GetAggregateConfigRulesReport API synchronously
func (client *Client) GetAggregateConfigRulesReport(request *GetAggregateConfigRulesReportRequest) (response *GetAggregateConfigRulesReportResponse, err error) {
	response = CreateGetAggregateConfigRulesReportResponse()
	err = client.DoAction(request, response)
	return
}

// GetAggregateConfigRulesReportWithChan invokes the config.GetAggregateConfigRulesReport API asynchronously
func (client *Client) GetAggregateConfigRulesReportWithChan(request *GetAggregateConfigRulesReportRequest) (<-chan *GetAggregateConfigRulesReportResponse, <-chan error) {
	responseChan := make(chan *GetAggregateConfigRulesReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAggregateConfigRulesReport(request)
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

// GetAggregateConfigRulesReportWithCallback invokes the config.GetAggregateConfigRulesReport API asynchronously
func (client *Client) GetAggregateConfigRulesReportWithCallback(request *GetAggregateConfigRulesReportRequest, callback func(response *GetAggregateConfigRulesReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAggregateConfigRulesReportResponse
		var err error
		defer close(result)
		response, err = client.GetAggregateConfigRulesReport(request)
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

// GetAggregateConfigRulesReportRequest is the request struct for api GetAggregateConfigRulesReport
type GetAggregateConfigRulesReportRequest struct {
	*requests.RpcRequest
	ReportId     string `position:"Query" name:"ReportId"`
	AggregatorId string `position:"Query" name:"AggregatorId"`
}

// GetAggregateConfigRulesReportResponse is the response struct for api GetAggregateConfigRulesReport
type GetAggregateConfigRulesReportResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	ConfigRulesReport ConfigRulesReport `json:"ConfigRulesReport" xml:"ConfigRulesReport"`
}

// CreateGetAggregateConfigRulesReportRequest creates a request to invoke GetAggregateConfigRulesReport API
func CreateGetAggregateConfigRulesReportRequest() (request *GetAggregateConfigRulesReportRequest) {
	request = &GetAggregateConfigRulesReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GetAggregateConfigRulesReport", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAggregateConfigRulesReportResponse creates a response to parse from GetAggregateConfigRulesReport response
func CreateGetAggregateConfigRulesReportResponse() (response *GetAggregateConfigRulesReportResponse) {
	response = &GetAggregateConfigRulesReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
