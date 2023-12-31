package servicemesh

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

// GetServiceMeshSlb invokes the servicemesh.GetServiceMeshSlb API synchronously
func (client *Client) GetServiceMeshSlb(request *GetServiceMeshSlbRequest) (response *GetServiceMeshSlbResponse, err error) {
	response = CreateGetServiceMeshSlbResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceMeshSlbWithChan invokes the servicemesh.GetServiceMeshSlb API asynchronously
func (client *Client) GetServiceMeshSlbWithChan(request *GetServiceMeshSlbRequest) (<-chan *GetServiceMeshSlbResponse, <-chan error) {
	responseChan := make(chan *GetServiceMeshSlbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceMeshSlb(request)
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

// GetServiceMeshSlbWithCallback invokes the servicemesh.GetServiceMeshSlb API asynchronously
func (client *Client) GetServiceMeshSlbWithCallback(request *GetServiceMeshSlbRequest, callback func(response *GetServiceMeshSlbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceMeshSlbResponse
		var err error
		defer close(result)
		response, err = client.GetServiceMeshSlb(request)
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

// GetServiceMeshSlbRequest is the request struct for api GetServiceMeshSlb
type GetServiceMeshSlbRequest struct {
	*requests.RpcRequest
	ServiceMeshId string `position:"Query" name:"ServiceMeshId"`
}

// GetServiceMeshSlbResponse is the response struct for api GetServiceMeshSlb
type GetServiceMeshSlbResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Data      []SLBInfo `json:"Data" xml:"Data"`
}

// CreateGetServiceMeshSlbRequest creates a request to invoke GetServiceMeshSlb API
func CreateGetServiceMeshSlbRequest() (request *GetServiceMeshSlbRequest) {
	request = &GetServiceMeshSlbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("servicemesh", "2020-01-11", "GetServiceMeshSlb", "servicemesh", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetServiceMeshSlbResponse creates a response to parse from GetServiceMeshSlb response
func CreateGetServiceMeshSlbResponse() (response *GetServiceMeshSlbResponse) {
	response = &GetServiceMeshSlbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
