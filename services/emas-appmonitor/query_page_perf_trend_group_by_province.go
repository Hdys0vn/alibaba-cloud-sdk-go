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

// QueryPagePerfTrendGroupByProvince invokes the emas_appmonitor.QueryPagePerfTrendGroupByProvince API synchronously
func (client *Client) QueryPagePerfTrendGroupByProvince(request *QueryPagePerfTrendGroupByProvinceRequest) (response *QueryPagePerfTrendGroupByProvinceResponse, err error) {
	response = CreateQueryPagePerfTrendGroupByProvinceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPagePerfTrendGroupByProvinceWithChan invokes the emas_appmonitor.QueryPagePerfTrendGroupByProvince API asynchronously
func (client *Client) QueryPagePerfTrendGroupByProvinceWithChan(request *QueryPagePerfTrendGroupByProvinceRequest) (<-chan *QueryPagePerfTrendGroupByProvinceResponse, <-chan error) {
	responseChan := make(chan *QueryPagePerfTrendGroupByProvinceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPagePerfTrendGroupByProvince(request)
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

// QueryPagePerfTrendGroupByProvinceWithCallback invokes the emas_appmonitor.QueryPagePerfTrendGroupByProvince API asynchronously
func (client *Client) QueryPagePerfTrendGroupByProvinceWithCallback(request *QueryPagePerfTrendGroupByProvinceRequest, callback func(response *QueryPagePerfTrendGroupByProvinceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPagePerfTrendGroupByProvinceResponse
		var err error
		defer close(result)
		response, err = client.QueryPagePerfTrendGroupByProvince(request)
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

// QueryPagePerfTrendGroupByProvinceRequest is the request struct for api QueryPagePerfTrendGroupByProvince
type QueryPagePerfTrendGroupByProvinceRequest struct {
	*requests.RpcRequest
	MetricType         string           `position:"Body" name:"MetricType"`
	AppVersionStrategy string           `position:"Body" name:"AppVersionStrategy"`
	StartTime          requests.Integer `position:"Body" name:"StartTime"`
	DeviceLevel        string           `position:"Body" name:"DeviceLevel"`
	StatType           string           `position:"Body" name:"StatType"`
	IntervalMinutes    requests.Integer `position:"Body" name:"IntervalMinutes"`
	UniqueAppId        string           `position:"Body" name:"UniqueAppId"`
	EndTime            requests.Integer `position:"Body" name:"EndTime"`
	AppVersion         *[]string        `position:"Body" name:"AppVersion"  type:"Repeated"`
	GroupByDistrict    requests.Boolean `position:"Body" name:"GroupByDistrict"`
}

// QueryPagePerfTrendGroupByProvinceResponse is the response struct for api QueryPagePerfTrendGroupByProvince
type QueryPagePerfTrendGroupByProvinceResponse struct {
	*responses.BaseResponse
	RequestId        string             `json:"RequestId" xml:"RequestId"`
	MetricResultList []MetricResultItem `json:"MetricResultList" xml:"MetricResultList"`
}

// CreateQueryPagePerfTrendGroupByProvinceRequest creates a request to invoke QueryPagePerfTrendGroupByProvince API
func CreateQueryPagePerfTrendGroupByProvinceRequest() (request *QueryPagePerfTrendGroupByProvinceRequest) {
	request = &QueryPagePerfTrendGroupByProvinceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emas-appmonitor", "2019-06-11", "QueryPagePerfTrendGroupByProvince", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryPagePerfTrendGroupByProvinceResponse creates a response to parse from QueryPagePerfTrendGroupByProvince response
func CreateQueryPagePerfTrendGroupByProvinceResponse() (response *QueryPagePerfTrendGroupByProvinceResponse) {
	response = &QueryPagePerfTrendGroupByProvinceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
