package iot

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

// QueryJobStatistics invokes the iot.QueryJobStatistics API synchronously
func (client *Client) QueryJobStatistics(request *QueryJobStatisticsRequest) (response *QueryJobStatisticsResponse, err error) {
	response = CreateQueryJobStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryJobStatisticsWithChan invokes the iot.QueryJobStatistics API asynchronously
func (client *Client) QueryJobStatisticsWithChan(request *QueryJobStatisticsRequest) (<-chan *QueryJobStatisticsResponse, <-chan error) {
	responseChan := make(chan *QueryJobStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryJobStatistics(request)
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

// QueryJobStatisticsWithCallback invokes the iot.QueryJobStatistics API asynchronously
func (client *Client) QueryJobStatisticsWithCallback(request *QueryJobStatisticsRequest, callback func(response *QueryJobStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryJobStatisticsResponse
		var err error
		defer close(result)
		response, err = client.QueryJobStatistics(request)
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

// QueryJobStatisticsRequest is the request struct for api QueryJobStatistics
type QueryJobStatisticsRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	JobId             string `position:"Query" name:"JobId"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
}

// QueryJobStatisticsResponse is the response struct for api QueryJobStatistics
type QueryJobStatisticsResponse struct {
	*responses.BaseResponse
	RequestId    string                   `json:"RequestId" xml:"RequestId"`
	Success      bool                     `json:"Success" xml:"Success"`
	Code         string                   `json:"Code" xml:"Code"`
	ErrorMessage string                   `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryJobStatistics `json:"Data" xml:"Data"`
}

// CreateQueryJobStatisticsRequest creates a request to invoke QueryJobStatistics API
func CreateQueryJobStatisticsRequest() (request *QueryJobStatisticsRequest) {
	request = &QueryJobStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryJobStatistics", "", "")
	request.Method = requests.GET
	return
}

// CreateQueryJobStatisticsResponse creates a response to parse from QueryJobStatistics response
func CreateQueryJobStatisticsResponse() (response *QueryJobStatisticsResponse) {
	response = &QueryJobStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
