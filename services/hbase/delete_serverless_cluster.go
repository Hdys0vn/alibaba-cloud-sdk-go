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

// DeleteServerlessCluster invokes the hbase.DeleteServerlessCluster API synchronously
func (client *Client) DeleteServerlessCluster(request *DeleteServerlessClusterRequest) (response *DeleteServerlessClusterResponse, err error) {
	response = CreateDeleteServerlessClusterResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteServerlessClusterWithChan invokes the hbase.DeleteServerlessCluster API asynchronously
func (client *Client) DeleteServerlessClusterWithChan(request *DeleteServerlessClusterRequest) (<-chan *DeleteServerlessClusterResponse, <-chan error) {
	responseChan := make(chan *DeleteServerlessClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteServerlessCluster(request)
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

// DeleteServerlessClusterWithCallback invokes the hbase.DeleteServerlessCluster API asynchronously
func (client *Client) DeleteServerlessClusterWithCallback(request *DeleteServerlessClusterRequest, callback func(response *DeleteServerlessClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteServerlessClusterResponse
		var err error
		defer close(result)
		response, err = client.DeleteServerlessCluster(request)
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

// DeleteServerlessClusterRequest is the request struct for api DeleteServerlessCluster
type DeleteServerlessClusterRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
	ZoneId    string `position:"Query" name:"ZoneId"`
}

// DeleteServerlessClusterResponse is the response struct for api DeleteServerlessCluster
type DeleteServerlessClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteServerlessClusterRequest creates a request to invoke DeleteServerlessCluster API
func CreateDeleteServerlessClusterRequest() (request *DeleteServerlessClusterRequest) {
	request = &DeleteServerlessClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DeleteServerlessCluster", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteServerlessClusterResponse creates a response to parse from DeleteServerlessCluster response
func CreateDeleteServerlessClusterResponse() (response *DeleteServerlessClusterResponse) {
	response = &DeleteServerlessClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
