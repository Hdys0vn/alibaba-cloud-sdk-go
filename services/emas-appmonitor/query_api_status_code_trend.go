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

// QueryApiStatusCodeTrend invokes the emas_appmonitor.QueryApiStatusCodeTrend API synchronously
func (client *Client) QueryApiStatusCodeTrend(request *QueryApiStatusCodeTrendRequest) (response *QueryApiStatusCodeTrendResponse, err error) {
	response = CreateQueryApiStatusCodeTrendResponse()
	err = client.DoAction(request, response)
	return
}

// QueryApiStatusCodeTrendWithChan invokes the emas_appmonitor.QueryApiStatusCodeTrend API asynchronously
func (client *Client) QueryApiStatusCodeTrendWithChan(request *QueryApiStatusCodeTrendRequest) (<-chan *QueryApiStatusCodeTrendResponse, <-chan error) {
	responseChan := make(chan *QueryApiStatusCodeTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryApiStatusCodeTrend(request)
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

// QueryApiStatusCodeTrendWithCallback invokes the emas_appmonitor.QueryApiStatusCodeTrend API asynchronously
func (client *Client) QueryApiStatusCodeTrendWithCallback(request *QueryApiStatusCodeTrendRequest, callback func(response *QueryApiStatusCodeTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryApiStatusCodeTrendResponse
		var err error
		defer close(result)
		response, err = client.QueryApiStatusCodeTrend(request)
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

// QueryApiStatusCodeTrendRequest is the request struct for api QueryApiStatusCodeTrend
type QueryApiStatusCodeTrendRequest struct {
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

// QueryApiStatusCodeTrendResponse is the response struct for api QueryApiStatusCodeTrend
type QueryApiStatusCodeTrendResponse struct {
	*responses.BaseResponse
	RequestId        string             `json:"RequestId" xml:"RequestId"`
	MetricResultList []MetricResultItem `json:"MetricResultList" xml:"MetricResultList"`
}

// CreateQueryApiStatusCodeTrendRequest creates a request to invoke QueryApiStatusCodeTrend API
func CreateQueryApiStatusCodeTrendRequest() (request *QueryApiStatusCodeTrendRequest) {
	request = &QueryApiStatusCodeTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emas-appmonitor", "2019-06-11", "QueryApiStatusCodeTrend", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryApiStatusCodeTrendResponse creates a response to parse from QueryApiStatusCodeTrend response
func CreateQueryApiStatusCodeTrendResponse() (response *QueryApiStatusCodeTrendResponse) {
	response = &QueryApiStatusCodeTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
