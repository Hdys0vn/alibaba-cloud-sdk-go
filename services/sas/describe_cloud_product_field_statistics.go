package sas

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

// DescribeCloudProductFieldStatistics invokes the sas.DescribeCloudProductFieldStatistics API synchronously
func (client *Client) DescribeCloudProductFieldStatistics(request *DescribeCloudProductFieldStatisticsRequest) (response *DescribeCloudProductFieldStatisticsResponse, err error) {
	response = CreateDescribeCloudProductFieldStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudProductFieldStatisticsWithChan invokes the sas.DescribeCloudProductFieldStatistics API asynchronously
func (client *Client) DescribeCloudProductFieldStatisticsWithChan(request *DescribeCloudProductFieldStatisticsRequest) (<-chan *DescribeCloudProductFieldStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudProductFieldStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudProductFieldStatistics(request)
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

// DescribeCloudProductFieldStatisticsWithCallback invokes the sas.DescribeCloudProductFieldStatistics API asynchronously
func (client *Client) DescribeCloudProductFieldStatisticsWithCallback(request *DescribeCloudProductFieldStatisticsRequest, callback func(response *DescribeCloudProductFieldStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudProductFieldStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudProductFieldStatistics(request)
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

// DescribeCloudProductFieldStatisticsRequest is the request struct for api DescribeCloudProductFieldStatistics
type DescribeCloudProductFieldStatisticsRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeCloudProductFieldStatisticsResponse is the response struct for api DescribeCloudProductFieldStatistics
type DescribeCloudProductFieldStatisticsResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	GroupedFields GroupedFields `json:"GroupedFields" xml:"GroupedFields"`
}

// CreateDescribeCloudProductFieldStatisticsRequest creates a request to invoke DescribeCloudProductFieldStatistics API
func CreateDescribeCloudProductFieldStatisticsRequest() (request *DescribeCloudProductFieldStatisticsRequest) {
	request = &DescribeCloudProductFieldStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeCloudProductFieldStatistics", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudProductFieldStatisticsResponse creates a response to parse from DescribeCloudProductFieldStatistics response
func CreateDescribeCloudProductFieldStatisticsResponse() (response *DescribeCloudProductFieldStatisticsResponse) {
	response = &DescribeCloudProductFieldStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
