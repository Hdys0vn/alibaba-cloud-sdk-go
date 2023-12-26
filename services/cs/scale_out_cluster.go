package cs

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

// ScaleOutCluster invokes the cs.ScaleOutCluster API synchronously
func (client *Client) ScaleOutCluster(request *ScaleOutClusterRequest) (response *ScaleOutClusterResponse, err error) {
	response = CreateScaleOutClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ScaleOutClusterWithChan invokes the cs.ScaleOutCluster API asynchronously
func (client *Client) ScaleOutClusterWithChan(request *ScaleOutClusterRequest) (<-chan *ScaleOutClusterResponse, <-chan error) {
	responseChan := make(chan *ScaleOutClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScaleOutCluster(request)
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

// ScaleOutClusterWithCallback invokes the cs.ScaleOutCluster API asynchronously
func (client *Client) ScaleOutClusterWithCallback(request *ScaleOutClusterRequest, callback func(response *ScaleOutClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScaleOutClusterResponse
		var err error
		defer close(result)
		response, err = client.ScaleOutCluster(request)
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

// ScaleOutClusterRequest is the request struct for api ScaleOutCluster
type ScaleOutClusterRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

// ScaleOutClusterResponse is the response struct for api ScaleOutCluster
type ScaleOutClusterResponse struct {
	*responses.BaseResponse
	ClusterId string `json:"cluster_id" xml:"cluster_id"`
	TaskId    string `json:"task_id" xml:"task_id"`
	RequestId string `json:"request_id" xml:"request_id"`
}

// CreateScaleOutClusterRequest creates a request to invoke ScaleOutCluster API
func CreateScaleOutClusterRequest() (request *ScaleOutClusterRequest) {
	request = &ScaleOutClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "ScaleOutCluster", "/api/v2/clusters/[ClusterId]", "", "")
	request.Method = requests.POST
	return
}

// CreateScaleOutClusterResponse creates a response to parse from ScaleOutCluster response
func CreateScaleOutClusterResponse() (response *ScaleOutClusterResponse) {
	response = &ScaleOutClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}