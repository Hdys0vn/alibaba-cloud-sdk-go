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

// QueryDeviceProp invokes the iot.QueryDeviceProp API synchronously
func (client *Client) QueryDeviceProp(request *QueryDevicePropRequest) (response *QueryDevicePropResponse, err error) {
	response = CreateQueryDevicePropResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDevicePropWithChan invokes the iot.QueryDeviceProp API asynchronously
func (client *Client) QueryDevicePropWithChan(request *QueryDevicePropRequest) (<-chan *QueryDevicePropResponse, <-chan error) {
	responseChan := make(chan *QueryDevicePropResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceProp(request)
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

// QueryDevicePropWithCallback invokes the iot.QueryDeviceProp API asynchronously
func (client *Client) QueryDevicePropWithCallback(request *QueryDevicePropRequest, callback func(response *QueryDevicePropResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDevicePropResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceProp(request)
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

// QueryDevicePropRequest is the request struct for api QueryDeviceProp
type QueryDevicePropRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotId             string `position:"Query" name:"IotId"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
	DeviceName        string `position:"Query" name:"DeviceName"`
}

// QueryDevicePropResponse is the response struct for api QueryDeviceProp
type QueryDevicePropResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Props        string `json:"Props" xml:"Props"`
}

// CreateQueryDevicePropRequest creates a request to invoke QueryDeviceProp API
func CreateQueryDevicePropRequest() (request *QueryDevicePropRequest) {
	request = &QueryDevicePropRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceProp", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDevicePropResponse creates a response to parse from QueryDeviceProp response
func CreateQueryDevicePropResponse() (response *QueryDevicePropResponse) {
	response = &QueryDevicePropResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
