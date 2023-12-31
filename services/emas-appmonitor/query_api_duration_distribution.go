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

// QueryApiDurationDistribution invokes the emas_appmonitor.QueryApiDurationDistribution API synchronously
func (client *Client) QueryApiDurationDistribution(request *QueryApiDurationDistributionRequest) (response *QueryApiDurationDistributionResponse, err error) {
	response = CreateQueryApiDurationDistributionResponse()
	err = client.DoAction(request, response)
	return
}

// QueryApiDurationDistributionWithChan invokes the emas_appmonitor.QueryApiDurationDistribution API asynchronously
func (client *Client) QueryApiDurationDistributionWithChan(request *QueryApiDurationDistributionRequest) (<-chan *QueryApiDurationDistributionResponse, <-chan error) {
	responseChan := make(chan *QueryApiDurationDistributionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryApiDurationDistribution(request)
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

// QueryApiDurationDistributionWithCallback invokes the emas_appmonitor.QueryApiDurationDistribution API asynchronously
func (client *Client) QueryApiDurationDistributionWithCallback(request *QueryApiDurationDistributionRequest, callback func(response *QueryApiDurationDistributionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryApiDurationDistributionResponse
		var err error
		defer close(result)
		response, err = client.QueryApiDurationDistribution(request)
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

// QueryApiDurationDistributionRequest is the request struct for api QueryApiDurationDistribution
type QueryApiDurationDistributionRequest struct {
	*requests.RpcRequest
	AppVersionStrategy string           `position:"Body" name:"AppVersionStrategy"`
	StartTime          requests.Integer `position:"Body" name:"StartTime"`
	StatType           string           `position:"Body" name:"StatType"`
	IntervalMinutes    requests.Integer `position:"Body" name:"IntervalMinutes"`
	UniqueAppId        string           `position:"Body" name:"UniqueAppId"`
	Ip                 *[]string        `position:"Body" name:"Ip"  type:"Repeated"`
	EndTime            requests.Integer `position:"Body" name:"EndTime"`
	AppVersion         *[]string        `position:"Body" name:"AppVersion"  type:"Repeated"`
	UrlPath            *[]string        `position:"Body" name:"UrlPath"  type:"Repeated"`
	Domain             string           `position:"Body" name:"Domain"`
}

// QueryApiDurationDistributionResponse is the response struct for api QueryApiDurationDistribution
type QueryApiDurationDistributionResponse struct {
	*responses.BaseResponse
	RequestId        string             `json:"RequestId" xml:"RequestId"`
	MetricResultList []MetricResultItem `json:"MetricResultList" xml:"MetricResultList"`
}

// CreateQueryApiDurationDistributionRequest creates a request to invoke QueryApiDurationDistribution API
func CreateQueryApiDurationDistributionRequest() (request *QueryApiDurationDistributionRequest) {
	request = &QueryApiDurationDistributionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emas-appmonitor", "2019-06-11", "QueryApiDurationDistribution", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryApiDurationDistributionResponse creates a response to parse from QueryApiDurationDistribution response
func CreateQueryApiDurationDistributionResponse() (response *QueryApiDurationDistributionResponse) {
	response = &QueryApiDurationDistributionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
