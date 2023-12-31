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

// OperateRenderingDevices invokes the vs.OperateRenderingDevices API synchronously
func (client *Client) OperateRenderingDevices(request *OperateRenderingDevicesRequest) (response *OperateRenderingDevicesResponse, err error) {
	response = CreateOperateRenderingDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// OperateRenderingDevicesWithChan invokes the vs.OperateRenderingDevices API asynchronously
func (client *Client) OperateRenderingDevicesWithChan(request *OperateRenderingDevicesRequest) (<-chan *OperateRenderingDevicesResponse, <-chan error) {
	responseChan := make(chan *OperateRenderingDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OperateRenderingDevices(request)
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

// OperateRenderingDevicesWithCallback invokes the vs.OperateRenderingDevices API asynchronously
func (client *Client) OperateRenderingDevicesWithCallback(request *OperateRenderingDevicesRequest, callback func(response *OperateRenderingDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OperateRenderingDevicesResponse
		var err error
		defer close(result)
		response, err = client.OperateRenderingDevices(request)
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

// OperateRenderingDevicesRequest is the request struct for api OperateRenderingDevices
type OperateRenderingDevicesRequest struct {
	*requests.RpcRequest
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds string           `position:"Query" name:"InstanceIds"`
	Operation   string           `position:"Query" name:"Operation"`
	PodId       string           `position:"Query" name:"PodId"`
}

// OperateRenderingDevicesResponse is the response struct for api OperateRenderingDevices
type OperateRenderingDevicesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	FailedIds []string `json:"FailedIds" xml:"FailedIds"`
}

// CreateOperateRenderingDevicesRequest creates a request to invoke OperateRenderingDevices API
func CreateOperateRenderingDevicesRequest() (request *OperateRenderingDevicesRequest) {
	request = &OperateRenderingDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "OperateRenderingDevices", "", "")
	request.Method = requests.POST
	return
}

// CreateOperateRenderingDevicesResponse creates a response to parse from OperateRenderingDevices response
func CreateOperateRenderingDevicesResponse() (response *OperateRenderingDevicesResponse) {
	response = &OperateRenderingDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
