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

// GetAppMonthlyDeviceCount invokes the emas_appmonitor.GetAppMonthlyDeviceCount API synchronously
func (client *Client) GetAppMonthlyDeviceCount(request *GetAppMonthlyDeviceCountRequest) (response *GetAppMonthlyDeviceCountResponse, err error) {
	response = CreateGetAppMonthlyDeviceCountResponse()
	err = client.DoAction(request, response)
	return
}

// GetAppMonthlyDeviceCountWithChan invokes the emas_appmonitor.GetAppMonthlyDeviceCount API asynchronously
func (client *Client) GetAppMonthlyDeviceCountWithChan(request *GetAppMonthlyDeviceCountRequest) (<-chan *GetAppMonthlyDeviceCountResponse, <-chan error) {
	responseChan := make(chan *GetAppMonthlyDeviceCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAppMonthlyDeviceCount(request)
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

// GetAppMonthlyDeviceCountWithCallback invokes the emas_appmonitor.GetAppMonthlyDeviceCount API asynchronously
func (client *Client) GetAppMonthlyDeviceCountWithCallback(request *GetAppMonthlyDeviceCountRequest, callback func(response *GetAppMonthlyDeviceCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAppMonthlyDeviceCountResponse
		var err error
		defer close(result)
		response, err = client.GetAppMonthlyDeviceCount(request)
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

// GetAppMonthlyDeviceCountRequest is the request struct for api GetAppMonthlyDeviceCount
type GetAppMonthlyDeviceCountRequest struct {
	*requests.RpcRequest
	UniqueAppId string           `position:"Body" name:"UniqueAppId"`
	FromDateMs  requests.Integer `position:"Body" name:"FromDateMs"`
	Service     string           `position:"Body" name:"Service"`
	UntilDateMs requests.Integer `position:"Body" name:"UntilDateMs"`
}

// GetAppMonthlyDeviceCountResponse is the response struct for api GetAppMonthlyDeviceCount
type GetAppMonthlyDeviceCountResponse struct {
	*responses.BaseResponse
	RequestId       string            `json:"RequestId" xml:"RequestId"`
	DeviceCountList []DeviceCountItem `json:"DeviceCountList" xml:"DeviceCountList"`
}

// CreateGetAppMonthlyDeviceCountRequest creates a request to invoke GetAppMonthlyDeviceCount API
func CreateGetAppMonthlyDeviceCountRequest() (request *GetAppMonthlyDeviceCountRequest) {
	request = &GetAppMonthlyDeviceCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emas-appmonitor", "2019-06-11", "GetAppMonthlyDeviceCount", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAppMonthlyDeviceCountResponse creates a response to parse from GetAppMonthlyDeviceCount response
func CreateGetAppMonthlyDeviceCountResponse() (response *GetAppMonthlyDeviceCountResponse) {
	response = &GetAppMonthlyDeviceCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
