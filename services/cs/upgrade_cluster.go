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

// UpgradeCluster invokes the cs.UpgradeCluster API synchronously
func (client *Client) UpgradeCluster(request *UpgradeClusterRequest) (response *UpgradeClusterResponse, err error) {
	response = CreateUpgradeClusterResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeClusterWithChan invokes the cs.UpgradeCluster API asynchronously
func (client *Client) UpgradeClusterWithChan(request *UpgradeClusterRequest) (<-chan *UpgradeClusterResponse, <-chan error) {
	responseChan := make(chan *UpgradeClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeCluster(request)
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

// UpgradeClusterWithCallback invokes the cs.UpgradeCluster API asynchronously
func (client *Client) UpgradeClusterWithCallback(request *UpgradeClusterRequest, callback func(response *UpgradeClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeClusterResponse
		var err error
		defer close(result)
		response, err = client.UpgradeCluster(request)
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

// UpgradeClusterRequest is the request struct for api UpgradeCluster
type UpgradeClusterRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

// UpgradeClusterResponse is the response struct for api UpgradeCluster
type UpgradeClusterResponse struct {
	*responses.BaseResponse
}

// CreateUpgradeClusterRequest creates a request to invoke UpgradeCluster API
func CreateUpgradeClusterRequest() (request *UpgradeClusterRequest) {
	request = &UpgradeClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "UpgradeCluster", "/api/v2/clusters/[ClusterId]/upgrade", "", "")
	request.Method = requests.POST
	return
}

// CreateUpgradeClusterResponse creates a response to parse from UpgradeCluster response
func CreateUpgradeClusterResponse() (response *UpgradeClusterResponse) {
	response = &UpgradeClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
