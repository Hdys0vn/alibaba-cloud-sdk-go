package aliyuncvc

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

// ListDeviceIp invokes the aliyuncvc.ListDeviceIp API synchronously
func (client *Client) ListDeviceIp(request *ListDeviceIpRequest) (response *ListDeviceIpResponse, err error) {
	response = CreateListDeviceIpResponse()
	err = client.DoAction(request, response)
	return
}

// ListDeviceIpWithChan invokes the aliyuncvc.ListDeviceIp API asynchronously
func (client *Client) ListDeviceIpWithChan(request *ListDeviceIpRequest) (<-chan *ListDeviceIpResponse, <-chan error) {
	responseChan := make(chan *ListDeviceIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDeviceIp(request)
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

// ListDeviceIpWithCallback invokes the aliyuncvc.ListDeviceIp API asynchronously
func (client *Client) ListDeviceIpWithCallback(request *ListDeviceIpRequest, callback func(response *ListDeviceIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDeviceIpResponse
		var err error
		defer close(result)
		response, err = client.ListDeviceIp(request)
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

// ListDeviceIpRequest is the request struct for api ListDeviceIp
type ListDeviceIpRequest struct {
	*requests.RpcRequest
	GroupId string `position:"Body" name:"GroupId"`
	SN      string `position:"Body" name:"SN"`
}

// ListDeviceIpResponse is the response struct for api ListDeviceIp
type ListDeviceIpResponse struct {
	*responses.BaseResponse
	ErrorCode int           `json:"ErrorCode" xml:"ErrorCode"`
	Message   string        `json:"Message" xml:"Message"`
	Success   bool          `json:"Success" xml:"Success"`
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Devices   []DevicesItem `json:"Devices" xml:"Devices"`
}

// CreateListDeviceIpRequest creates a request to invoke ListDeviceIp API
func CreateListDeviceIpRequest() (request *ListDeviceIpRequest) {
	request = &ListDeviceIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "ListDeviceIp", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDeviceIpResponse creates a response to parse from ListDeviceIp response
func CreateListDeviceIpResponse() (response *ListDeviceIpResponse) {
	response = &ListDeviceIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
