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

// DeleteDeviceProp invokes the iot.DeleteDeviceProp API synchronously
func (client *Client) DeleteDeviceProp(request *DeleteDevicePropRequest) (response *DeleteDevicePropResponse, err error) {
	response = CreateDeleteDevicePropResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDevicePropWithChan invokes the iot.DeleteDeviceProp API asynchronously
func (client *Client) DeleteDevicePropWithChan(request *DeleteDevicePropRequest) (<-chan *DeleteDevicePropResponse, <-chan error) {
	responseChan := make(chan *DeleteDevicePropResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDeviceProp(request)
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

// DeleteDevicePropWithCallback invokes the iot.DeleteDeviceProp API asynchronously
func (client *Client) DeleteDevicePropWithCallback(request *DeleteDevicePropRequest, callback func(response *DeleteDevicePropResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDevicePropResponse
		var err error
		defer close(result)
		response, err = client.DeleteDeviceProp(request)
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

// DeleteDevicePropRequest is the request struct for api DeleteDeviceProp
type DeleteDevicePropRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	PropKey           string `position:"Query" name:"PropKey"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotId             string `position:"Query" name:"IotId"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
	DeviceName        string `position:"Query" name:"DeviceName"`
}

// DeleteDevicePropResponse is the response struct for api DeleteDeviceProp
type DeleteDevicePropResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteDevicePropRequest creates a request to invoke DeleteDeviceProp API
func CreateDeleteDevicePropRequest() (request *DeleteDevicePropRequest) {
	request = &DeleteDevicePropRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteDeviceProp", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDevicePropResponse creates a response to parse from DeleteDeviceProp response
func CreateDeleteDevicePropResponse() (response *DeleteDevicePropResponse) {
	response = &DeleteDevicePropResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
