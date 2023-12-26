package ltl

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

// ListDevice invokes the ltl.ListDevice API synchronously
func (client *Client) ListDevice(request *ListDeviceRequest) (response *ListDeviceResponse, err error) {
	response = CreateListDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// ListDeviceWithChan invokes the ltl.ListDevice API asynchronously
func (client *Client) ListDeviceWithChan(request *ListDeviceRequest) (<-chan *ListDeviceResponse, <-chan error) {
	responseChan := make(chan *ListDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDevice(request)
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

// ListDeviceWithCallback invokes the ltl.ListDevice API asynchronously
func (client *Client) ListDeviceWithCallback(request *ListDeviceRequest, callback func(response *ListDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDeviceResponse
		var err error
		defer close(result)
		response, err = client.ListDevice(request)
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

// ListDeviceRequest is the request struct for api ListDevice
type ListDeviceRequest struct {
	*requests.RpcRequest
	IotId         string           `position:"Query" name:"IotId"`
	Size          requests.Integer `position:"Query" name:"Size"`
	Num           requests.Integer `position:"Query" name:"Num"`
	ApiVersion    string           `position:"Query" name:"ApiVersion"`
	DeviceGroupId string           `position:"Query" name:"DeviceGroupId"`
	BizChainId    string           `position:"Query" name:"BizChainId"`
}

// ListDeviceResponse is the response struct for api ListDevice
type ListDeviceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListDeviceRequest creates a request to invoke ListDevice API
func CreateListDeviceRequest() (request *ListDeviceRequest) {
	request = &ListDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "ListDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateListDeviceResponse creates a response to parse from ListDevice response
func CreateListDeviceResponse() (response *ListDeviceResponse) {
	response = &ListDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
