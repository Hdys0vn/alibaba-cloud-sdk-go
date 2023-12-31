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

// DescribeRecordResolveStatisticsSummary invokes the alidns.DescribeRecordResolveStatisticsSummary API synchronously
func (client *Client) DescribeRecordResolveStatisticsSummary(request *DescribeRecordResolveStatisticsSummaryRequest) (response *DescribeRecordResolveStatisticsSummaryResponse, err error) {
	response = CreateDescribeRecordResolveStatisticsSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecordResolveStatisticsSummaryWithChan invokes the alidns.DescribeRecordResolveStatisticsSummary API asynchronously
func (client *Client) DescribeRecordResolveStatisticsSummaryWithChan(request *DescribeRecordResolveStatisticsSummaryRequest) (<-chan *DescribeRecordResolveStatisticsSummaryResponse, <-chan error) {
	responseChan := make(chan *DescribeRecordResolveStatisticsSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecordResolveStatisticsSummary(request)
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

// DescribeRecordResolveStatisticsSummaryWithCallback invokes the alidns.DescribeRecordResolveStatisticsSummary API asynchronously
func (client *Client) DescribeRecordResolveStatisticsSummaryWithCallback(request *DescribeRecordResolveStatisticsSummaryRequest, callback func(response *DescribeRecordResolveStatisticsSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecordResolveStatisticsSummaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecordResolveStatisticsSummary(request)
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

// DescribeRecordResolveStatisticsSummaryRequest is the request struct for api DescribeRecordResolveStatisticsSummary
type DescribeRecordResolveStatisticsSummaryRequest struct {
	*requests.RpcRequest
	Threshold    requests.Integer `position:"Query" name:"Threshold"`
	StartDate    string           `position:"Query" name:"StartDate"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	DomainType   string           `position:"Query" name:"DomainType"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Keyword      string           `position:"Query" name:"Keyword"`
	Lang         string           `position:"Query" name:"Lang"`
	Direction    string           `position:"Query" name:"Direction"`
	DomainName   string           `position:"Query" name:"DomainName"`
	EndDate      string           `position:"Query" name:"EndDate"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	SearchMode   string           `position:"Query" name:"SearchMode"`
}

// DescribeRecordResolveStatisticsSummaryResponse is the response struct for api DescribeRecordResolveStatisticsSummary
type DescribeRecordResolveStatisticsSummaryResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	PageSize   int         `json:"PageSize" xml:"PageSize"`
	PageNumber int         `json:"PageNumber" xml:"PageNumber"`
	TotalPages int         `json:"TotalPages" xml:"TotalPages"`
	TotalItems int         `json:"TotalItems" xml:"TotalItems"`
	Statistics []Statistic `json:"Statistics" xml:"Statistics"`
}

// CreateDescribeRecordResolveStatisticsSummaryRequest creates a request to invoke DescribeRecordResolveStatisticsSummary API
func CreateDescribeRecordResolveStatisticsSummaryRequest() (request *DescribeRecordResolveStatisticsSummaryRequest) {
	request = &DescribeRecordResolveStatisticsSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeRecordResolveStatisticsSummary", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRecordResolveStatisticsSummaryResponse creates a response to parse from DescribeRecordResolveStatisticsSummary response
func CreateDescribeRecordResolveStatisticsSummaryResponse() (response *DescribeRecordResolveStatisticsSummaryResponse) {
	response = &DescribeRecordResolveStatisticsSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
