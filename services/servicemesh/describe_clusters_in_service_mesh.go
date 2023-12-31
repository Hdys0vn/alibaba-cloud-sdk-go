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

// DescribeClustersInServiceMesh invokes the servicemesh.DescribeClustersInServiceMesh API synchronously
func (client *Client) DescribeClustersInServiceMesh(request *DescribeClustersInServiceMeshRequest) (response *DescribeClustersInServiceMeshResponse, err error) {
	response = CreateDescribeClustersInServiceMeshResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClustersInServiceMeshWithChan invokes the servicemesh.DescribeClustersInServiceMesh API asynchronously
func (client *Client) DescribeClustersInServiceMeshWithChan(request *DescribeClustersInServiceMeshRequest) (<-chan *DescribeClustersInServiceMeshResponse, <-chan error) {
	responseChan := make(chan *DescribeClustersInServiceMeshResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClustersInServiceMesh(request)
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

// DescribeClustersInServiceMeshWithCallback invokes the servicemesh.DescribeClustersInServiceMesh API asynchronously
func (client *Client) DescribeClustersInServiceMeshWithCallback(request *DescribeClustersInServiceMeshRequest, callback func(response *DescribeClustersInServiceMeshResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClustersInServiceMeshResponse
		var err error
		defer close(result)
		response, err = client.DescribeClustersInServiceMesh(request)
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

// DescribeClustersInServiceMeshRequest is the request struct for api DescribeClustersInServiceMesh
type DescribeClustersInServiceMeshRequest struct {
	*requests.RpcRequest
	ServiceMeshId string `position:"Query" name:"ServiceMeshId"`
}

// DescribeClustersInServiceMeshResponse is the response struct for api DescribeClustersInServiceMesh
type DescribeClustersInServiceMeshResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Clusters  []Cluster `json:"Clusters" xml:"Clusters"`
}

// CreateDescribeClustersInServiceMeshRequest creates a request to invoke DescribeClustersInServiceMesh API
func CreateDescribeClustersInServiceMeshRequest() (request *DescribeClustersInServiceMeshRequest) {
	request = &DescribeClustersInServiceMeshRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("servicemesh", "2020-01-11", "DescribeClustersInServiceMesh", "servicemesh", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClustersInServiceMeshResponse creates a response to parse from DescribeClustersInServiceMesh response
func CreateDescribeClustersInServiceMeshResponse() (response *DescribeClustersInServiceMeshResponse) {
	response = &DescribeClustersInServiceMeshResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
