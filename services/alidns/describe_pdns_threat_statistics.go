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

// DescribePdnsThreatStatistics invokes the alidns.DescribePdnsThreatStatistics API synchronously
func (client *Client) DescribePdnsThreatStatistics(request *DescribePdnsThreatStatisticsRequest) (response *DescribePdnsThreatStatisticsResponse, err error) {
	response = CreateDescribePdnsThreatStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePdnsThreatStatisticsWithChan invokes the alidns.DescribePdnsThreatStatistics API asynchronously
func (client *Client) DescribePdnsThreatStatisticsWithChan(request *DescribePdnsThreatStatisticsRequest) (<-chan *DescribePdnsThreatStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribePdnsThreatStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePdnsThreatStatistics(request)
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

// DescribePdnsThreatStatisticsWithCallback invokes the alidns.DescribePdnsThreatStatistics API asynchronously
func (client *Client) DescribePdnsThreatStatisticsWithCallback(request *DescribePdnsThreatStatisticsRequest, callback func(response *DescribePdnsThreatStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePdnsThreatStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribePdnsThreatStatistics(request)
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

// DescribePdnsThreatStatisticsRequest is the request struct for api DescribePdnsThreatStatistics
type DescribePdnsThreatStatisticsRequest struct {
	*requests.RpcRequest
	Type           string           `position:"Query" name:"Type"`
	StartDate      string           `position:"Query" name:"StartDate"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	ThreatType     string           `position:"Query" name:"ThreatType"`
	Lang           string           `position:"Query" name:"Lang"`
	Direction      string           `position:"Query" name:"Direction"`
	ThreatSourceIp string           `position:"Query" name:"ThreatSourceIp"`
	DomainName     string           `position:"Query" name:"DomainName"`
	OrderBy        string           `position:"Query" name:"OrderBy"`
	EndDate        string           `position:"Query" name:"EndDate"`
	SubDomain      string           `position:"Query" name:"SubDomain"`
	ThreatLevel    string           `position:"Query" name:"ThreatLevel"`
}

// DescribePdnsThreatStatisticsResponse is the response struct for api DescribePdnsThreatStatistics
type DescribePdnsThreatStatisticsResponse struct {
	*responses.BaseResponse
	TotalCount int64           `json:"TotalCount" xml:"TotalCount"`
	PageSize   int64           `json:"PageSize" xml:"PageSize"`
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	PageNumber int64           `json:"PageNumber" xml:"PageNumber"`
	Data       []StatisticItem `json:"Data" xml:"Data"`
}

// CreateDescribePdnsThreatStatisticsRequest creates a request to invoke DescribePdnsThreatStatistics API
func CreateDescribePdnsThreatStatisticsRequest() (request *DescribePdnsThreatStatisticsRequest) {
	request = &DescribePdnsThreatStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribePdnsThreatStatistics", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePdnsThreatStatisticsResponse creates a response to parse from DescribePdnsThreatStatistics response
func CreateDescribePdnsThreatStatisticsResponse() (response *DescribePdnsThreatStatisticsResponse) {
	response = &DescribePdnsThreatStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
