package r_kvstore

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

// ModifyDedicatedUserCluster invokes the r_kvstore.ModifyDedicatedUserCluster API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifydedicatedusercluster.html
func (client *Client) ModifyDedicatedUserCluster(request *ModifyDedicatedUserClusterRequest) (response *ModifyDedicatedUserClusterResponse, err error) {
	response = CreateModifyDedicatedUserClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDedicatedUserClusterWithChan invokes the r_kvstore.ModifyDedicatedUserCluster API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifydedicatedusercluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedUserClusterWithChan(request *ModifyDedicatedUserClusterRequest) (<-chan *ModifyDedicatedUserClusterResponse, <-chan error) {
	responseChan := make(chan *ModifyDedicatedUserClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDedicatedUserCluster(request)
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

// ModifyDedicatedUserClusterWithCallback invokes the r_kvstore.ModifyDedicatedUserCluster API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifydedicatedusercluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedUserClusterWithCallback(request *ModifyDedicatedUserClusterRequest, callback func(response *ModifyDedicatedUserClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDedicatedUserClusterResponse
		var err error
		defer close(result)
		response, err = client.ModifyDedicatedUserCluster(request)
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

// ModifyDedicatedUserClusterRequest is the request struct for api ModifyDedicatedUserCluster
type ModifyDedicatedUserClusterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterName               string           `position:"Query" name:"ClusterName"`
	SecurityToken             string           `position:"Query" name:"SecurityToken"`
	Engine                    string           `position:"Query" name:"Engine"`
	DiskOverAllocationRatio   requests.Integer `position:"Query" name:"DiskOverAllocationRatio"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	ClusterId                 string           `position:"Query" name:"ClusterId"`
	MemoryOverAllocationRatio requests.Integer `position:"Query" name:"MemoryOverAllocationRatio"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	HostReplacePolicy         string           `position:"Query" name:"HostReplacePolicy"`
	AllocationPolicy          string           `position:"Query" name:"AllocationPolicy"`
	ZoneId                    string           `position:"Query" name:"ZoneId"`
	CpuOverAllocationRatio    requests.Integer `position:"Query" name:"CpuOverAllocationRatio"`
}

// ModifyDedicatedUserClusterResponse is the response struct for api ModifyDedicatedUserCluster
type ModifyDedicatedUserClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDedicatedUserClusterRequest creates a request to invoke ModifyDedicatedUserCluster API
func CreateModifyDedicatedUserClusterRequest() (request *ModifyDedicatedUserClusterRequest) {
	request = &ModifyDedicatedUserClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyDedicatedUserCluster", "redisa", "openAPI")
	return
}

// CreateModifyDedicatedUserClusterResponse creates a response to parse from ModifyDedicatedUserCluster response
func CreateModifyDedicatedUserClusterResponse() (response *ModifyDedicatedUserClusterResponse) {
	response = &ModifyDedicatedUserClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
