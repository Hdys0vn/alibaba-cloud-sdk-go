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

// DescribeRecordStatistics invokes the alidns.DescribeRecordStatistics API synchronously
func (client *Client) DescribeRecordStatistics(request *DescribeRecordStatisticsRequest) (response *DescribeRecordStatisticsResponse, err error) {
	response = CreateDescribeRecordStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecordStatisticsWithChan invokes the alidns.DescribeRecordStatistics API asynchronously
func (client *Client) DescribeRecordStatisticsWithChan(request *DescribeRecordStatisticsRequest) (<-chan *DescribeRecordStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeRecordStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecordStatistics(request)
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

// DescribeRecordStatisticsWithCallback invokes the alidns.DescribeRecordStatistics API asynchronously
func (client *Client) DescribeRecordStatisticsWithCallback(request *DescribeRecordStatisticsRequest, callback func(response *DescribeRecordStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecordStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecordStatistics(request)
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

// DescribeRecordStatisticsRequest is the request struct for api DescribeRecordStatistics
type DescribeRecordStatisticsRequest struct {
	*requests.RpcRequest
	Rr           string `position:"Query" name:"Rr"`
	DomainName   string `position:"Query" name:"DomainName"`
	StartDate    string `position:"Query" name:"StartDate"`
	EndDate      string `position:"Query" name:"EndDate"`
	DomainType   string `position:"Query" name:"DomainType"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeRecordStatisticsResponse is the response struct for api DescribeRecordStatistics
type DescribeRecordStatisticsResponse struct {
	*responses.BaseResponse
	RequestId  string                               `json:"RequestId" xml:"RequestId"`
	Statistics StatisticsInDescribeRecordStatistics `json:"Statistics" xml:"Statistics"`
}

// CreateDescribeRecordStatisticsRequest creates a request to invoke DescribeRecordStatistics API
func CreateDescribeRecordStatisticsRequest() (request *DescribeRecordStatisticsRequest) {
	request = &DescribeRecordStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeRecordStatistics", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRecordStatisticsResponse creates a response to parse from DescribeRecordStatistics response
func CreateDescribeRecordStatisticsResponse() (response *DescribeRecordStatisticsResponse) {
	response = &DescribeRecordStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
