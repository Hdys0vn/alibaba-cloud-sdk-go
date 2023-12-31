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

// DescribeDohDomainStatisticsSummary invokes the alidns.DescribeDohDomainStatisticsSummary API synchronously
func (client *Client) DescribeDohDomainStatisticsSummary(request *DescribeDohDomainStatisticsSummaryRequest) (response *DescribeDohDomainStatisticsSummaryResponse, err error) {
	response = CreateDescribeDohDomainStatisticsSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDohDomainStatisticsSummaryWithChan invokes the alidns.DescribeDohDomainStatisticsSummary API asynchronously
func (client *Client) DescribeDohDomainStatisticsSummaryWithChan(request *DescribeDohDomainStatisticsSummaryRequest) (<-chan *DescribeDohDomainStatisticsSummaryResponse, <-chan error) {
	responseChan := make(chan *DescribeDohDomainStatisticsSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDohDomainStatisticsSummary(request)
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

// DescribeDohDomainStatisticsSummaryWithCallback invokes the alidns.DescribeDohDomainStatisticsSummary API asynchronously
func (client *Client) DescribeDohDomainStatisticsSummaryWithCallback(request *DescribeDohDomainStatisticsSummaryRequest, callback func(response *DescribeDohDomainStatisticsSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDohDomainStatisticsSummaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeDohDomainStatisticsSummary(request)
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

// DescribeDohDomainStatisticsSummaryRequest is the request struct for api DescribeDohDomainStatisticsSummary
type DescribeDohDomainStatisticsSummaryRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OrderBy    string           `position:"Query" name:"OrderBy"`
	StartDate  string           `position:"Query" name:"StartDate"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	EndDate    string           `position:"Query" name:"EndDate"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Lang       string           `position:"Query" name:"Lang"`
	Direction  string           `position:"Query" name:"Direction"`
}

// DescribeDohDomainStatisticsSummaryResponse is the response struct for api DescribeDohDomainStatisticsSummary
type DescribeDohDomainStatisticsSummaryResponse struct {
	*responses.BaseResponse
	PageSize   int         `json:"PageSize" xml:"PageSize"`
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	PageNumber int         `json:"PageNumber" xml:"PageNumber"`
	TotalPages int         `json:"TotalPages" xml:"TotalPages"`
	TotalItems int         `json:"TotalItems" xml:"TotalItems"`
	Statistics []Statistic `json:"Statistics" xml:"Statistics"`
}

// CreateDescribeDohDomainStatisticsSummaryRequest creates a request to invoke DescribeDohDomainStatisticsSummary API
func CreateDescribeDohDomainStatisticsSummaryRequest() (request *DescribeDohDomainStatisticsSummaryRequest) {
	request = &DescribeDohDomainStatisticsSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDohDomainStatisticsSummary", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDohDomainStatisticsSummaryResponse creates a response to parse from DescribeDohDomainStatisticsSummary response
func CreateDescribeDohDomainStatisticsSummaryResponse() (response *DescribeDohDomainStatisticsSummaryResponse) {
	response = &DescribeDohDomainStatisticsSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
