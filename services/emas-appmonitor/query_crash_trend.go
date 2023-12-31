package emas_appmonitor

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

// QueryCrashTrend invokes the emas_appmonitor.QueryCrashTrend API synchronously
func (client *Client) QueryCrashTrend(request *QueryCrashTrendRequest) (response *QueryCrashTrendResponse, err error) {
	response = CreateQueryCrashTrendResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCrashTrendWithChan invokes the emas_appmonitor.QueryCrashTrend API asynchronously
func (client *Client) QueryCrashTrendWithChan(request *QueryCrashTrendRequest) (<-chan *QueryCrashTrendResponse, <-chan error) {
	responseChan := make(chan *QueryCrashTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCrashTrend(request)
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

// QueryCrashTrendWithCallback invokes the emas_appmonitor.QueryCrashTrend API asynchronously
func (client *Client) QueryCrashTrendWithCallback(request *QueryCrashTrendRequest, callback func(response *QueryCrashTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCrashTrendResponse
		var err error
		defer close(result)
		response, err = client.QueryCrashTrend(request)
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

// QueryCrashTrendRequest is the request struct for api QueryCrashTrend
type QueryCrashTrendRequest struct {
	*requests.RpcRequest
	AppVersionStrategy string           `position:"Body" name:"AppVersionStrategy"`
	StartTime          requests.Integer `position:"Body" name:"StartTime"`
	IntervalMinutes    requests.Integer `position:"Body" name:"IntervalMinutes"`
	UniqueAppId        string           `position:"Body" name:"UniqueAppId"`
	CrashStatType      string           `position:"Body" name:"CrashStatType"`
	EndTime            requests.Integer `position:"Body" name:"EndTime"`
	AppVersion         *[]string        `position:"Body" name:"AppVersion"  type:"Repeated"`
	ErrorType          *[]string        `position:"Body" name:"ErrorType"  type:"Repeated"`
	ErrorCategory      string           `position:"Body" name:"ErrorCategory"`
}

// QueryCrashTrendResponse is the response struct for api QueryCrashTrend
type QueryCrashTrendResponse struct {
	*responses.BaseResponse
	RequestId        string             `json:"RequestId" xml:"RequestId"`
	MetricResultList []MetricResultItem `json:"MetricResultList" xml:"MetricResultList"`
}

// CreateQueryCrashTrendRequest creates a request to invoke QueryCrashTrend API
func CreateQueryCrashTrendRequest() (request *QueryCrashTrendRequest) {
	request = &QueryCrashTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emas-appmonitor", "2019-06-11", "QueryCrashTrend", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryCrashTrendResponse creates a response to parse from QueryCrashTrend response
func CreateQueryCrashTrendResponse() (response *QueryCrashTrendResponse) {
	response = &QueryCrashTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
