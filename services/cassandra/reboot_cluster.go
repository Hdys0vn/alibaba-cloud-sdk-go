package cassandra

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

// RebootCluster invokes the cassandra.RebootCluster API synchronously
func (client *Client) RebootCluster(request *RebootClusterRequest) (response *RebootClusterResponse, err error) {
	response = CreateRebootClusterResponse()
	err = client.DoAction(request, response)
	return
}

// RebootClusterWithChan invokes the cassandra.RebootCluster API asynchronously
func (client *Client) RebootClusterWithChan(request *RebootClusterRequest) (<-chan *RebootClusterResponse, <-chan error) {
	responseChan := make(chan *RebootClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RebootCluster(request)
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

// RebootClusterWithCallback invokes the cassandra.RebootCluster API asynchronously
func (client *Client) RebootClusterWithCallback(request *RebootClusterRequest, callback func(response *RebootClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RebootClusterResponse
		var err error
		defer close(result)
		response, err = client.RebootCluster(request)
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

// RebootClusterRequest is the request struct for api RebootCluster
type RebootClusterRequest struct {
	*requests.RpcRequest
	DataCenterId string `position:"Query" name:"DataCenterId"`
	ClusterId    string `position:"Query" name:"ClusterId"`
}

// RebootClusterResponse is the response struct for api RebootCluster
type RebootClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRebootClusterRequest creates a request to invoke RebootCluster API
func CreateRebootClusterRequest() (request *RebootClusterRequest) {
	request = &RebootClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "RebootCluster", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRebootClusterResponse creates a response to parse from RebootCluster response
func CreateRebootClusterResponse() (response *RebootClusterResponse) {
	response = &RebootClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
