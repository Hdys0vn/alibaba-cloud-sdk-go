package ecs

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

// DeleteSnapshot invokes the ecs.DeleteSnapshot API synchronously
func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (response *DeleteSnapshotResponse, err error) {
	response = CreateDeleteSnapshotResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSnapshotWithChan invokes the ecs.DeleteSnapshot API asynchronously
func (client *Client) DeleteSnapshotWithChan(request *DeleteSnapshotRequest) (<-chan *DeleteSnapshotResponse, <-chan error) {
	responseChan := make(chan *DeleteSnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSnapshot(request)
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

// DeleteSnapshotWithCallback invokes the ecs.DeleteSnapshot API asynchronously
func (client *Client) DeleteSnapshotWithCallback(request *DeleteSnapshotRequest, callback func(response *DeleteSnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSnapshotResponse
		var err error
		defer close(result)
		response, err = client.DeleteSnapshot(request)
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

// DeleteSnapshotRequest is the request struct for api DeleteSnapshot
type DeleteSnapshotRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string           `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Force                requests.Boolean `position:"Query" name:"Force"`
}

// DeleteSnapshotResponse is the response struct for api DeleteSnapshot
type DeleteSnapshotResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSnapshotRequest creates a request to invoke DeleteSnapshot API
func CreateDeleteSnapshotRequest() (request *DeleteSnapshotRequest) {
	request = &DeleteSnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeleteSnapshot", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteSnapshotResponse creates a response to parse from DeleteSnapshot response
func CreateDeleteSnapshotResponse() (response *DeleteSnapshotResponse) {
	response = &DeleteSnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
