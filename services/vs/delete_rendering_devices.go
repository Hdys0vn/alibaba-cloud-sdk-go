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

// DeleteRenderingDevices invokes the vs.DeleteRenderingDevices API synchronously
func (client *Client) DeleteRenderingDevices(request *DeleteRenderingDevicesRequest) (response *DeleteRenderingDevicesResponse, err error) {
	response = CreateDeleteRenderingDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRenderingDevicesWithChan invokes the vs.DeleteRenderingDevices API asynchronously
func (client *Client) DeleteRenderingDevicesWithChan(request *DeleteRenderingDevicesRequest) (<-chan *DeleteRenderingDevicesResponse, <-chan error) {
	responseChan := make(chan *DeleteRenderingDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRenderingDevices(request)
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

// DeleteRenderingDevicesWithCallback invokes the vs.DeleteRenderingDevices API asynchronously
func (client *Client) DeleteRenderingDevicesWithCallback(request *DeleteRenderingDevicesRequest, callback func(response *DeleteRenderingDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRenderingDevicesResponse
		var err error
		defer close(result)
		response, err = client.DeleteRenderingDevices(request)
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

// DeleteRenderingDevicesRequest is the request struct for api DeleteRenderingDevices
type DeleteRenderingDevicesRequest struct {
	*requests.RpcRequest
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds string           `position:"Query" name:"InstanceIds"`
}

// DeleteRenderingDevicesResponse is the response struct for api DeleteRenderingDevices
type DeleteRenderingDevicesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	FailedIds []string `json:"FailedIds" xml:"FailedIds"`
}

// CreateDeleteRenderingDevicesRequest creates a request to invoke DeleteRenderingDevices API
func CreateDeleteRenderingDevicesRequest() (request *DeleteRenderingDevicesRequest) {
	request = &DeleteRenderingDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DeleteRenderingDevices", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteRenderingDevicesResponse creates a response to parse from DeleteRenderingDevices response
func CreateDeleteRenderingDevicesResponse() (response *DeleteRenderingDevicesResponse) {
	response = &DeleteRenderingDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
