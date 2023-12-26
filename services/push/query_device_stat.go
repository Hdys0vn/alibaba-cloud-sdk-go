package push

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

// QueryDeviceStat invokes the push.QueryDeviceStat API synchronously
func (client *Client) QueryDeviceStat(request *QueryDeviceStatRequest) (response *QueryDeviceStatResponse, err error) {
	response = CreateQueryDeviceStatResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceStatWithChan invokes the push.QueryDeviceStat API asynchronously
func (client *Client) QueryDeviceStatWithChan(request *QueryDeviceStatRequest) (<-chan *QueryDeviceStatResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceStatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceStat(request)
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

// QueryDeviceStatWithCallback invokes the push.QueryDeviceStat API asynchronously
func (client *Client) QueryDeviceStatWithCallback(request *QueryDeviceStatRequest, callback func(response *QueryDeviceStatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceStatResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceStat(request)
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

// QueryDeviceStatRequest is the request struct for api QueryDeviceStat
type QueryDeviceStatRequest struct {
	*requests.RpcRequest
	EndTime    string           `position:"Query" name:"EndTime"`
	StartTime  string           `position:"Query" name:"StartTime"`
	DeviceType string           `position:"Query" name:"DeviceType"`
	AppKey     requests.Integer `position:"Query" name:"AppKey"`
	QueryType  string           `position:"Query" name:"QueryType"`
}

// QueryDeviceStatResponse is the response struct for api QueryDeviceStat
type QueryDeviceStatResponse struct {
	*responses.BaseResponse
	RequestId      string                          `json:"RequestId" xml:"RequestId"`
	AppDeviceStats AppDeviceStatsInQueryDeviceStat `json:"AppDeviceStats" xml:"AppDeviceStats"`
}

// CreateQueryDeviceStatRequest creates a request to invoke QueryDeviceStat API
func CreateQueryDeviceStatRequest() (request *QueryDeviceStatRequest) {
	request = &QueryDeviceStatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryDeviceStat", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceStatResponse creates a response to parse from QueryDeviceStat response
func CreateQueryDeviceStatResponse() (response *QueryDeviceStatResponse) {
	response = &QueryDeviceStatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
