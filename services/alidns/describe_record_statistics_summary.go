package alidns

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

// DescribeRecordStatisticsSummary invokes the alidns.DescribeRecordStatisticsSummary API synchronously
func (client *Client) DescribeRecordStatisticsSummary(request *DescribeRecordStatisticsSummaryRequest) (response *DescribeRecordStatisticsSummaryResponse, err error) {
	response = CreateDescribeRecordStatisticsSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecordStatisticsSummaryWithChan invokes the alidns.DescribeRecordStatisticsSummary API asynchronously
func (client *Client) DescribeRecordStatisticsSummaryWithChan(request *DescribeRecordStatisticsSummaryRequest) (<-chan *DescribeRecordStatisticsSummaryResponse, <-chan error) {
	responseChan := make(chan *DescribeRecordStatisticsSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecordStatisticsSummary(request)
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

// DescribeRecordStatisticsSummaryWithCallback invokes the alidns.DescribeRecordStatisticsSummary API asynchronously
func (client *Client) DescribeRecordStatisticsSummaryWithCallback(request *DescribeRecordStatisticsSummaryRequest, callback func(response *DescribeRecordStatisticsSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecordStatisticsSummaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecordStatisticsSummary(request)
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

// DescribeRecordStatisticsSummaryRequest is the request struct for api DescribeRecordStatisticsSummary
type DescribeRecordStatisticsSummaryRequest struct {
	*requests.RpcRequest
	Threshold    requests.Integer `position:"Query" name:"Threshold"`
	StartDate    string           `position:"Query" name:"StartDate"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	DomainType   string           `position:"Query" name:"DomainType"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	Keyword      string           `position:"Query" name:"Keyword"`
	Direction    string           `position:"Query" name:"Direction"`
	DomainName   string           `position:"Query" name:"DomainName"`
	OrderBy      string           `position:"Query" name:"OrderBy"`
	EndDate      string           `position:"Query" name:"EndDate"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	SearchMode   string           `position:"Query" name:"SearchMode"`
}

// DescribeRecordStatisticsSummaryResponse is the response struct for api DescribeRecordStatisticsSummary
type DescribeRecordStatisticsSummaryResponse struct {
	*responses.BaseResponse
	PageSize   int                                         `json:"PageSize" xml:"PageSize"`
	RequestId  string                                      `json:"RequestId" xml:"RequestId"`
	PageNumber int                                         `json:"PageNumber" xml:"PageNumber"`
	TotalPages int                                         `json:"TotalPages" xml:"TotalPages"`
	TotalItems int                                         `json:"TotalItems" xml:"TotalItems"`
	Statistics StatisticsInDescribeRecordStatisticsSummary `json:"Statistics" xml:"Statistics"`
}

// CreateDescribeRecordStatisticsSummaryRequest creates a request to invoke DescribeRecordStatisticsSummary API
func CreateDescribeRecordStatisticsSummaryRequest() (request *DescribeRecordStatisticsSummaryRequest) {
	request = &DescribeRecordStatisticsSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeRecordStatisticsSummary", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRecordStatisticsSummaryResponse creates a response to parse from DescribeRecordStatisticsSummary response
func CreateDescribeRecordStatisticsSummaryResponse() (response *DescribeRecordStatisticsSummaryResponse) {
	response = &DescribeRecordStatisticsSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
