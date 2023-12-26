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

// QueryApiErrorTrend invokes the emas_appmonitor.QueryApiErrorTrend API synchronously
func (client *Client) QueryApiErrorTrend(request *QueryApiErrorTrendRequest) (response *QueryApiErrorTrendResponse, err error) {
	response = CreateQueryApiErrorTrendResponse()
	err = client.DoAction(request, response)
	return
}

// QueryApiErrorTrendWithChan invokes the emas_appmonitor.QueryApiErrorTrend API asynchronously
func (client *Client) QueryApiErrorTrendWithChan(request *QueryApiErrorTrendRequest) (<-chan *QueryApiErrorTrendResponse, <-chan error) {
	responseChan := make(chan *QueryApiErrorTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryApiErrorTrend(request)
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

// QueryApiErrorTrendWithCallback invokes the emas_appmonitor.QueryApiErrorTrend API asynchronously
func (client *Client) QueryApiErrorTrendWithCallback(request *QueryApiErrorTrendRequest, callback func(response *QueryApiErrorTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryApiErrorTrendResponse
		var err error
		defer close(result)
		response, err = client.QueryApiErrorTrend(request)
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

// QueryApiErrorTrendRequest is the request struct for api QueryApiErrorTrend
type QueryApiErrorTrendRequest struct {
	*requests.RpcRequest
	AppVersionStrategy string           `position:"Body" name:"AppVersionStrategy"`
	StartTime          requests.Integer `position:"Body" name:"StartTime"`
	IntervalMinutes    requests.Integer `position:"Body" name:"IntervalMinutes"`
	UniqueAppId        string           `position:"Body" name:"UniqueAppId"`
	Ip                 *[]string        `position:"Body" name:"Ip"  type:"Repeated"`
	EndTime            requests.Integer `position:"Body" name:"EndTime"`
	AppVersion         *[]string        `position:"Body" name:"AppVersion"  type:"Repeated"`
	UrlPath            *[]string        `position:"Body" name:"UrlPath"  type:"Repeated"`
	Domain             string           `position:"Body" name:"Domain"`
}

// QueryApiErrorTrendResponse is the response struct for api QueryApiErrorTrend
type QueryApiErrorTrendResponse struct {
	*responses.BaseResponse
	RequestId        string             `json:"RequestId" xml:"RequestId"`
	MetricResultList []MetricResultItem `json:"MetricResultList" xml:"MetricResultList"`
}

// CreateQueryApiErrorTrendRequest creates a request to invoke QueryApiErrorTrend API
func CreateQueryApiErrorTrendRequest() (request *QueryApiErrorTrendRequest) {
	request = &QueryApiErrorTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emas-appmonitor", "2019-06-11", "QueryApiErrorTrend", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryApiErrorTrendResponse creates a response to parse from QueryApiErrorTrend response
func CreateQueryApiErrorTrendResponse() (response *QueryApiErrorTrendResponse) {
	response = &QueryApiErrorTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
