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

// DescribeDomainStatistics invokes the alidns.DescribeDomainStatistics API synchronously
func (client *Client) DescribeDomainStatistics(request *DescribeDomainStatisticsRequest) (response *DescribeDomainStatisticsResponse, err error) {
	response = CreateDescribeDomainStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainStatisticsWithChan invokes the alidns.DescribeDomainStatistics API asynchronously
func (client *Client) DescribeDomainStatisticsWithChan(request *DescribeDomainStatisticsRequest) (<-chan *DescribeDomainStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainStatistics(request)
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

// DescribeDomainStatisticsWithCallback invokes the alidns.DescribeDomainStatistics API asynchronously
func (client *Client) DescribeDomainStatisticsWithCallback(request *DescribeDomainStatisticsRequest, callback func(response *DescribeDomainStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainStatistics(request)
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

// DescribeDomainStatisticsRequest is the request struct for api DescribeDomainStatistics
type DescribeDomainStatisticsRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	StartDate    string `position:"Query" name:"StartDate"`
	EndDate      string `position:"Query" name:"EndDate"`
	DomainType   string `position:"Query" name:"DomainType"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeDomainStatisticsResponse is the response struct for api DescribeDomainStatistics
type DescribeDomainStatisticsResponse struct {
	*responses.BaseResponse
	RequestId  string                               `json:"RequestId" xml:"RequestId"`
	Statistics StatisticsInDescribeDomainStatistics `json:"Statistics" xml:"Statistics"`
}

// CreateDescribeDomainStatisticsRequest creates a request to invoke DescribeDomainStatistics API
func CreateDescribeDomainStatisticsRequest() (request *DescribeDomainStatisticsRequest) {
	request = &DescribeDomainStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainStatistics", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainStatisticsResponse creates a response to parse from DescribeDomainStatistics response
func CreateDescribeDomainStatisticsResponse() (response *DescribeDomainStatisticsResponse) {
	response = &DescribeDomainStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
