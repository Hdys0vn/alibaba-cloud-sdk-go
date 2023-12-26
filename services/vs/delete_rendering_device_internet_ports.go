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

// DeleteRenderingDeviceInternetPorts invokes the vs.DeleteRenderingDeviceInternetPorts API synchronously
func (client *Client) DeleteRenderingDeviceInternetPorts(request *DeleteRenderingDeviceInternetPortsRequest) (response *DeleteRenderingDeviceInternetPortsResponse, err error) {
	response = CreateDeleteRenderingDeviceInternetPortsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRenderingDeviceInternetPortsWithChan invokes the vs.DeleteRenderingDeviceInternetPorts API asynchronously
func (client *Client) DeleteRenderingDeviceInternetPortsWithChan(request *DeleteRenderingDeviceInternetPortsRequest) (<-chan *DeleteRenderingDeviceInternetPortsResponse, <-chan error) {
	responseChan := make(chan *DeleteRenderingDeviceInternetPortsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRenderingDeviceInternetPorts(request)
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

// DeleteRenderingDeviceInternetPortsWithCallback invokes the vs.DeleteRenderingDeviceInternetPorts API asynchronously
func (client *Client) DeleteRenderingDeviceInternetPortsWithCallback(request *DeleteRenderingDeviceInternetPortsRequest, callback func(response *DeleteRenderingDeviceInternetPortsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRenderingDeviceInternetPortsResponse
		var err error
		defer close(result)
		response, err = client.DeleteRenderingDeviceInternetPorts(request)
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

// DeleteRenderingDeviceInternetPortsRequest is the request struct for api DeleteRenderingDeviceInternetPorts
type DeleteRenderingDeviceInternetPortsRequest struct {
	*requests.RpcRequest
	ISP          string           `position:"Query" name:"ISP"`
	ShowLog      string           `position:"Query" name:"ShowLog"`
	IpProtocol   string           `position:"Query" name:"IpProtocol"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds  string           `position:"Query" name:"InstanceIds"`
	InternalPort string           `position:"Query" name:"InternalPort"`
}

// DeleteRenderingDeviceInternetPortsResponse is the response struct for api DeleteRenderingDeviceInternetPorts
type DeleteRenderingDeviceInternetPortsResponse struct {
	*responses.BaseResponse
	RequestId   string                                          `json:"RequestId" xml:"RequestId"`
	InstanceIds InstanceIdsInDeleteRenderingDeviceInternetPorts `json:"InstanceIds" xml:"InstanceIds"`
}

// CreateDeleteRenderingDeviceInternetPortsRequest creates a request to invoke DeleteRenderingDeviceInternetPorts API
func CreateDeleteRenderingDeviceInternetPortsRequest() (request *DeleteRenderingDeviceInternetPortsRequest) {
	request = &DeleteRenderingDeviceInternetPortsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DeleteRenderingDeviceInternetPorts", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteRenderingDeviceInternetPortsResponse creates a response to parse from DeleteRenderingDeviceInternetPorts response
func CreateDeleteRenderingDeviceInternetPortsResponse() (response *DeleteRenderingDeviceInternetPortsResponse) {
	response = &DeleteRenderingDeviceInternetPortsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
