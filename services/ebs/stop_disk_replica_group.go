package ebs

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

// StopDiskReplicaGroup invokes the ebs.StopDiskReplicaGroup API synchronously
func (client *Client) StopDiskReplicaGroup(request *StopDiskReplicaGroupRequest) (response *StopDiskReplicaGroupResponse, err error) {
	response = CreateStopDiskReplicaGroupResponse()
	err = client.DoAction(request, response)
	return
}

// StopDiskReplicaGroupWithChan invokes the ebs.StopDiskReplicaGroup API asynchronously
func (client *Client) StopDiskReplicaGroupWithChan(request *StopDiskReplicaGroupRequest) (<-chan *StopDiskReplicaGroupResponse, <-chan error) {
	responseChan := make(chan *StopDiskReplicaGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopDiskReplicaGroup(request)
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

// StopDiskReplicaGroupWithCallback invokes the ebs.StopDiskReplicaGroup API asynchronously
func (client *Client) StopDiskReplicaGroupWithCallback(request *StopDiskReplicaGroupRequest, callback func(response *StopDiskReplicaGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopDiskReplicaGroupResponse
		var err error
		defer close(result)
		response, err = client.StopDiskReplicaGroup(request)
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

// StopDiskReplicaGroupRequest is the request struct for api StopDiskReplicaGroup
type StopDiskReplicaGroupRequest struct {
	*requests.RpcRequest
	ClientToken    string `position:"Query" name:"ClientToken"`
	ReplicaGroupId string `position:"Query" name:"ReplicaGroupId"`
}

// StopDiskReplicaGroupResponse is the response struct for api StopDiskReplicaGroup
type StopDiskReplicaGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopDiskReplicaGroupRequest creates a request to invoke StopDiskReplicaGroup API
func CreateStopDiskReplicaGroupRequest() (request *StopDiskReplicaGroupRequest) {
	request = &StopDiskReplicaGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ebs", "2021-07-30", "StopDiskReplicaGroup", "ebs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopDiskReplicaGroupResponse creates a response to parse from StopDiskReplicaGroup response
func CreateStopDiskReplicaGroupResponse() (response *StopDiskReplicaGroupResponse) {
	response = &StopDiskReplicaGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
