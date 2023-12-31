package aegis

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

// DescribeSuspTrendStatistics invokes the aegis.DescribeSuspTrendStatistics API synchronously
// api document: https://help.aliyun.com/api/aegis/describesusptrendstatistics.html
func (client *Client) DescribeSuspTrendStatistics(request *DescribeSuspTrendStatisticsRequest) (response *DescribeSuspTrendStatisticsResponse, err error) {
	response = CreateDescribeSuspTrendStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSuspTrendStatisticsWithChan invokes the aegis.DescribeSuspTrendStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesusptrendstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspTrendStatisticsWithChan(request *DescribeSuspTrendStatisticsRequest) (<-chan *DescribeSuspTrendStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeSuspTrendStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSuspTrendStatistics(request)
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

// DescribeSuspTrendStatisticsWithCallback invokes the aegis.DescribeSuspTrendStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesusptrendstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspTrendStatisticsWithCallback(request *DescribeSuspTrendStatisticsRequest, callback func(response *DescribeSuspTrendStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSuspTrendStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSuspTrendStatistics(request)
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

// DescribeSuspTrendStatisticsRequest is the request struct for api DescribeSuspTrendStatistics
type DescribeSuspTrendStatisticsRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeSuspTrendStatisticsResponse is the response struct for api DescribeSuspTrendStatistics
type DescribeSuspTrendStatisticsResponse struct {
	*responses.BaseResponse
	RequestId       string   `json:"RequestId" xml:"RequestId"`
	StartTime       int      `json:"StartTime" xml:"StartTime"`
	Interval        int      `json:"Interval" xml:"Interval"`
	SuspiciousItems []string `json:"SuspiciousItems" xml:"SuspiciousItems"`
}

// CreateDescribeSuspTrendStatisticsRequest creates a request to invoke DescribeSuspTrendStatistics API
func CreateDescribeSuspTrendStatisticsRequest() (request *DescribeSuspTrendStatisticsRequest) {
	request = &DescribeSuspTrendStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeSuspTrendStatistics", "vipaegis", "openAPI")
	return
}

// CreateDescribeSuspTrendStatisticsResponse creates a response to parse from DescribeSuspTrendStatistics response
func CreateDescribeSuspTrendStatisticsResponse() (response *DescribeSuspTrendStatisticsResponse) {
	response = &DescribeSuspTrendStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
