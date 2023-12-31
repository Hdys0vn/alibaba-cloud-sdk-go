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

// QueryDeviceDistributeDetail invokes the iot.QueryDeviceDistributeDetail API synchronously
func (client *Client) QueryDeviceDistributeDetail(request *QueryDeviceDistributeDetailRequest) (response *QueryDeviceDistributeDetailResponse, err error) {
	response = CreateQueryDeviceDistributeDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceDistributeDetailWithChan invokes the iot.QueryDeviceDistributeDetail API asynchronously
func (client *Client) QueryDeviceDistributeDetailWithChan(request *QueryDeviceDistributeDetailRequest) (<-chan *QueryDeviceDistributeDetailResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceDistributeDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceDistributeDetail(request)
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

// QueryDeviceDistributeDetailWithCallback invokes the iot.QueryDeviceDistributeDetail API asynchronously
func (client *Client) QueryDeviceDistributeDetailWithCallback(request *QueryDeviceDistributeDetailRequest, callback func(response *QueryDeviceDistributeDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceDistributeDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceDistributeDetail(request)
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

// QueryDeviceDistributeDetailRequest is the request struct for api QueryDeviceDistributeDetail
type QueryDeviceDistributeDetailRequest struct {
	*requests.RpcRequest
	JobId       string `position:"Query" name:"JobId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// QueryDeviceDistributeDetailResponse is the response struct for api QueryDeviceDistributeDetail
type QueryDeviceDistributeDetailResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	File         string `json:"File" xml:"File"`
}

// CreateQueryDeviceDistributeDetailRequest creates a request to invoke QueryDeviceDistributeDetail API
func CreateQueryDeviceDistributeDetailRequest() (request *QueryDeviceDistributeDetailRequest) {
	request = &QueryDeviceDistributeDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceDistributeDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceDistributeDetailResponse creates a response to parse from QueryDeviceDistributeDetail response
func CreateQueryDeviceDistributeDetailResponse() (response *QueryDeviceDistributeDetailResponse) {
	response = &QueryDeviceDistributeDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
