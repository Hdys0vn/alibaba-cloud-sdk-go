package foas

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

// DestroyCluster invokes the foas.DestroyCluster API synchronously
func (client *Client) DestroyCluster(request *DestroyClusterRequest) (response *DestroyClusterResponse, err error) {
	response = CreateDestroyClusterResponse()
	err = client.DoAction(request, response)
	return
}

// DestroyClusterWithChan invokes the foas.DestroyCluster API asynchronously
func (client *Client) DestroyClusterWithChan(request *DestroyClusterRequest) (<-chan *DestroyClusterResponse, <-chan error) {
	responseChan := make(chan *DestroyClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DestroyCluster(request)
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

// DestroyClusterWithCallback invokes the foas.DestroyCluster API asynchronously
func (client *Client) DestroyClusterWithCallback(request *DestroyClusterRequest, callback func(response *DestroyClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DestroyClusterResponse
		var err error
		defer close(result)
		response, err = client.DestroyCluster(request)
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

// DestroyClusterRequest is the request struct for api DestroyCluster
type DestroyClusterRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"clusterId"`
}

// DestroyClusterResponse is the response struct for api DestroyCluster
type DestroyClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDestroyClusterRequest creates a request to invoke DestroyCluster API
func CreateDestroyClusterRequest() (request *DestroyClusterRequest) {
	request = &DestroyClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "DestroyCluster", "/api/v2/clusters/[clusterId]", "foas", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDestroyClusterResponse creates a response to parse from DestroyCluster response
func CreateDestroyClusterResponse() (response *DestroyClusterResponse) {
	response = &DestroyClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
