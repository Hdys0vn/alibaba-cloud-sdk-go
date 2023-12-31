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

// RemoveClusterFromServiceMesh invokes the servicemesh.RemoveClusterFromServiceMesh API synchronously
func (client *Client) RemoveClusterFromServiceMesh(request *RemoveClusterFromServiceMeshRequest) (response *RemoveClusterFromServiceMeshResponse, err error) {
	response = CreateRemoveClusterFromServiceMeshResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveClusterFromServiceMeshWithChan invokes the servicemesh.RemoveClusterFromServiceMesh API asynchronously
func (client *Client) RemoveClusterFromServiceMeshWithChan(request *RemoveClusterFromServiceMeshRequest) (<-chan *RemoveClusterFromServiceMeshResponse, <-chan error) {
	responseChan := make(chan *RemoveClusterFromServiceMeshResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveClusterFromServiceMesh(request)
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

// RemoveClusterFromServiceMeshWithCallback invokes the servicemesh.RemoveClusterFromServiceMesh API asynchronously
func (client *Client) RemoveClusterFromServiceMeshWithCallback(request *RemoveClusterFromServiceMeshRequest, callback func(response *RemoveClusterFromServiceMeshResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveClusterFromServiceMeshResponse
		var err error
		defer close(result)
		response, err = client.RemoveClusterFromServiceMesh(request)
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

// RemoveClusterFromServiceMeshRequest is the request struct for api RemoveClusterFromServiceMesh
type RemoveClusterFromServiceMeshRequest struct {
	*requests.RpcRequest
	ClusterId     string `position:"Body" name:"ClusterId"`
	ServiceMeshId string `position:"Body" name:"ServiceMeshId"`
}

// RemoveClusterFromServiceMeshResponse is the response struct for api RemoveClusterFromServiceMesh
type RemoveClusterFromServiceMeshResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateRemoveClusterFromServiceMeshRequest creates a request to invoke RemoveClusterFromServiceMesh API
func CreateRemoveClusterFromServiceMeshRequest() (request *RemoveClusterFromServiceMeshRequest) {
	request = &RemoveClusterFromServiceMeshRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("servicemesh", "2020-01-11", "RemoveClusterFromServiceMesh", "servicemesh", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveClusterFromServiceMeshResponse creates a response to parse from RemoveClusterFromServiceMesh response
func CreateRemoveClusterFromServiceMeshResponse() (response *RemoveClusterFromServiceMeshResponse) {
	response = &RemoveClusterFromServiceMeshResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
