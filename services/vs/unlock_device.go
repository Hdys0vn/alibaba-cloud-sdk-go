package vs

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

// UnlockDevice invokes the vs.UnlockDevice API synchronously
func (client *Client) UnlockDevice(request *UnlockDeviceRequest) (response *UnlockDeviceResponse, err error) {
	response = CreateUnlockDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// UnlockDeviceWithChan invokes the vs.UnlockDevice API asynchronously
func (client *Client) UnlockDeviceWithChan(request *UnlockDeviceRequest) (<-chan *UnlockDeviceResponse, <-chan error) {
	responseChan := make(chan *UnlockDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnlockDevice(request)
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

// UnlockDeviceWithCallback invokes the vs.UnlockDevice API asynchronously
func (client *Client) UnlockDeviceWithCallback(request *UnlockDeviceRequest, callback func(response *UnlockDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnlockDeviceResponse
		var err error
		defer close(result)
		response, err = client.UnlockDevice(request)
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

// UnlockDeviceRequest is the request struct for api UnlockDevice
type UnlockDeviceRequest struct {
	*requests.RpcRequest
	Id      string           `position:"Query" name:"Id"`
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// UnlockDeviceResponse is the response struct for api UnlockDevice
type UnlockDeviceResponse struct {
	*responses.BaseResponse
	Id        string `json:"Id" xml:"Id"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnlockDeviceRequest creates a request to invoke UnlockDevice API
func CreateUnlockDeviceRequest() (request *UnlockDeviceRequest) {
	request = &UnlockDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "UnlockDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateUnlockDeviceResponse creates a response to parse from UnlockDevice response
func CreateUnlockDeviceResponse() (response *UnlockDeviceResponse) {
	response = &UnlockDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
