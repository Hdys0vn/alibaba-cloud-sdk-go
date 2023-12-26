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

// DeleteCluster invokes the vs.DeleteCluster API synchronously
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
	response = CreateDeleteClusterResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteClusterWithChan invokes the vs.DeleteCluster API asynchronously
func (client *Client) DeleteClusterWithChan(request *DeleteClusterRequest) (<-chan *DeleteClusterResponse, <-chan error) {
	responseChan := make(chan *DeleteClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCluster(request)
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

// DeleteClusterWithCallback invokes the vs.DeleteCluster API asynchronously
func (client *Client) DeleteClusterWithCallback(request *DeleteClusterRequest, callback func(response *DeleteClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteClusterResponse
		var err error
		defer close(result)
		response, err = client.DeleteCluster(request)
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

// DeleteClusterRequest is the request struct for api DeleteCluster
type DeleteClusterRequest struct {
	*requests.RpcRequest
	ShowLog   string           `position:"Query" name:"ShowLog"`
	ClusterId string           `position:"Query" name:"ClusterId"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteClusterResponse is the response struct for api DeleteCluster
type DeleteClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteClusterRequest creates a request to invoke DeleteCluster API
func CreateDeleteClusterRequest() (request *DeleteClusterRequest) {
	request = &DeleteClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DeleteCluster", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteClusterResponse creates a response to parse from DeleteCluster response
func CreateDeleteClusterResponse() (response *DeleteClusterResponse) {
	response = &DeleteClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
