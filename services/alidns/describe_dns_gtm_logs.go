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

// DescribeDnsGtmLogs invokes the alidns.DescribeDnsGtmLogs API synchronously
func (client *Client) DescribeDnsGtmLogs(request *DescribeDnsGtmLogsRequest) (response *DescribeDnsGtmLogsResponse, err error) {
	response = CreateDescribeDnsGtmLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDnsGtmLogsWithChan invokes the alidns.DescribeDnsGtmLogs API asynchronously
func (client *Client) DescribeDnsGtmLogsWithChan(request *DescribeDnsGtmLogsRequest) (<-chan *DescribeDnsGtmLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeDnsGtmLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDnsGtmLogs(request)
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

// DescribeDnsGtmLogsWithCallback invokes the alidns.DescribeDnsGtmLogs API asynchronously
func (client *Client) DescribeDnsGtmLogsWithCallback(request *DescribeDnsGtmLogsRequest, callback func(response *DescribeDnsGtmLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDnsGtmLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDnsGtmLogs(request)
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

// DescribeDnsGtmLogsRequest is the request struct for api DescribeDnsGtmLogs
type DescribeDnsGtmLogsRequest struct {
	*requests.RpcRequest
	StartTimestamp requests.Integer `position:"Query" name:"StartTimestamp"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	EndTimestamp   requests.Integer `position:"Query" name:"EndTimestamp"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	UserClientIp   string           `position:"Query" name:"UserClientIp"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	Lang           string           `position:"Query" name:"Lang"`
	Keyword        string           `position:"Query" name:"Keyword"`
}

// DescribeDnsGtmLogsResponse is the response struct for api DescribeDnsGtmLogs
type DescribeDnsGtmLogsResponse struct {
	*responses.BaseResponse
	PageSize   int                      `json:"PageSize" xml:"PageSize"`
	RequestId  string                   `json:"RequestId" xml:"RequestId"`
	PageNumber int                      `json:"PageNumber" xml:"PageNumber"`
	TotalPages int                      `json:"TotalPages" xml:"TotalPages"`
	TotalItems int                      `json:"TotalItems" xml:"TotalItems"`
	Logs       LogsInDescribeDnsGtmLogs `json:"Logs" xml:"Logs"`
}

// CreateDescribeDnsGtmLogsRequest creates a request to invoke DescribeDnsGtmLogs API
func CreateDescribeDnsGtmLogsRequest() (request *DescribeDnsGtmLogsRequest) {
	request = &DescribeDnsGtmLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsGtmLogs", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDnsGtmLogsResponse creates a response to parse from DescribeDnsGtmLogs response
func CreateDescribeDnsGtmLogsResponse() (response *DescribeDnsGtmLogsResponse) {
	response = &DescribeDnsGtmLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
