package hbase

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

// DeleteMultiZoneCluster invokes the hbase.DeleteMultiZoneCluster API synchronously
func (client *Client) DeleteMultiZoneCluster(request *DeleteMultiZoneClusterRequest) (response *DeleteMultiZoneClusterResponse, err error) {
	response = CreateDeleteMultiZoneClusterResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMultiZoneClusterWithChan invokes the hbase.DeleteMultiZoneCluster API asynchronously
func (client *Client) DeleteMultiZoneClusterWithChan(request *DeleteMultiZoneClusterRequest) (<-chan *DeleteMultiZoneClusterResponse, <-chan error) {
	responseChan := make(chan *DeleteMultiZoneClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMultiZoneCluster(request)
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

// DeleteMultiZoneClusterWithCallback invokes the hbase.DeleteMultiZoneCluster API asynchronously
func (client *Client) DeleteMultiZoneClusterWithCallback(request *DeleteMultiZoneClusterRequest, callback func(response *DeleteMultiZoneClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMultiZoneClusterResponse
		var err error
		defer close(result)
		response, err = client.DeleteMultiZoneCluster(request)
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

// DeleteMultiZoneClusterRequest is the request struct for api DeleteMultiZoneCluster
type DeleteMultiZoneClusterRequest struct {
	*requests.RpcRequest
	ImmediateDeleteFlag requests.Boolean `position:"Query" name:"ImmediateDeleteFlag"`
	ClusterId           string           `position:"Query" name:"ClusterId"`
}

// DeleteMultiZoneClusterResponse is the response struct for api DeleteMultiZoneCluster
type DeleteMultiZoneClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteMultiZoneClusterRequest creates a request to invoke DeleteMultiZoneCluster API
func CreateDeleteMultiZoneClusterRequest() (request *DeleteMultiZoneClusterRequest) {
	request = &DeleteMultiZoneClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DeleteMultiZoneCluster", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteMultiZoneClusterResponse creates a response to parse from DeleteMultiZoneCluster response
func CreateDeleteMultiZoneClusterResponse() (response *DeleteMultiZoneClusterResponse) {
	response = &DeleteMultiZoneClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
