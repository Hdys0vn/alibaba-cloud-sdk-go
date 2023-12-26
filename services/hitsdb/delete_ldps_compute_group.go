package hitsdb

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

// DeleteLdpsComputeGroup invokes the hitsdb.DeleteLdpsComputeGroup API synchronously
func (client *Client) DeleteLdpsComputeGroup(request *DeleteLdpsComputeGroupRequest) (response *DeleteLdpsComputeGroupResponse, err error) {
	response = CreateDeleteLdpsComputeGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLdpsComputeGroupWithChan invokes the hitsdb.DeleteLdpsComputeGroup API asynchronously
func (client *Client) DeleteLdpsComputeGroupWithChan(request *DeleteLdpsComputeGroupRequest) (<-chan *DeleteLdpsComputeGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteLdpsComputeGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLdpsComputeGroup(request)
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

// DeleteLdpsComputeGroupWithCallback invokes the hitsdb.DeleteLdpsComputeGroup API asynchronously
func (client *Client) DeleteLdpsComputeGroupWithCallback(request *DeleteLdpsComputeGroupRequest, callback func(response *DeleteLdpsComputeGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLdpsComputeGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteLdpsComputeGroup(request)
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

// DeleteLdpsComputeGroupRequest is the request struct for api DeleteLdpsComputeGroup
type DeleteLdpsComputeGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	GroupName            string           `position:"Query" name:"GroupName"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DeleteLdpsComputeGroupResponse is the response struct for api DeleteLdpsComputeGroup
type DeleteLdpsComputeGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLdpsComputeGroupRequest creates a request to invoke DeleteLdpsComputeGroup API
func CreateDeleteLdpsComputeGroupRequest() (request *DeleteLdpsComputeGroupRequest) {
	request = &DeleteLdpsComputeGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2020-06-15", "DeleteLdpsComputeGroup", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLdpsComputeGroupResponse creates a response to parse from DeleteLdpsComputeGroup response
func CreateDeleteLdpsComputeGroupResponse() (response *DeleteLdpsComputeGroupResponse) {
	response = &DeleteLdpsComputeGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
