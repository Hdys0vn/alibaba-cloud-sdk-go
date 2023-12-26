package clickhouse

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

// ModifyDBCluster invokes the clickhouse.ModifyDBCluster API synchronously
func (client *Client) ModifyDBCluster(request *ModifyDBClusterRequest) (response *ModifyDBClusterResponse, err error) {
	response = CreateModifyDBClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterWithChan invokes the clickhouse.ModifyDBCluster API asynchronously
func (client *Client) ModifyDBClusterWithChan(request *ModifyDBClusterRequest) (<-chan *ModifyDBClusterResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBCluster(request)
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

// ModifyDBClusterWithCallback invokes the clickhouse.ModifyDBCluster API asynchronously
func (client *Client) ModifyDBClusterWithCallback(request *ModifyDBClusterRequest, callback func(response *ModifyDBClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBCluster(request)
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

// ModifyDBClusterRequest is the request struct for api ModifyDBCluster
type ModifyDBClusterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBClusterClass       string           `position:"Query" name:"DBClusterClass"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBNodeGroupCount     string           `position:"Query" name:"DBNodeGroupCount"`
	DBNodeStorage        string           `position:"Query" name:"DBNodeStorage"`
}

// ModifyDBClusterResponse is the response struct for api ModifyDBCluster
type ModifyDBClusterResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	DBCluster DBCluster `json:"DBCluster" xml:"DBCluster"`
}

// CreateModifyDBClusterRequest creates a request to invoke ModifyDBCluster API
func CreateModifyDBClusterRequest() (request *ModifyDBClusterRequest) {
	request = &ModifyDBClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "ModifyDBCluster", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterResponse creates a response to parse from ModifyDBCluster response
func CreateModifyDBClusterResponse() (response *ModifyDBClusterResponse) {
	response = &ModifyDBClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}