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

// ReprotectDiskReplicaGroup invokes the ebs.ReprotectDiskReplicaGroup API synchronously
func (client *Client) ReprotectDiskReplicaGroup(request *ReprotectDiskReplicaGroupRequest) (response *ReprotectDiskReplicaGroupResponse, err error) {
	response = CreateReprotectDiskReplicaGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ReprotectDiskReplicaGroupWithChan invokes the ebs.ReprotectDiskReplicaGroup API asynchronously
func (client *Client) ReprotectDiskReplicaGroupWithChan(request *ReprotectDiskReplicaGroupRequest) (<-chan *ReprotectDiskReplicaGroupResponse, <-chan error) {
	responseChan := make(chan *ReprotectDiskReplicaGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReprotectDiskReplicaGroup(request)
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

// ReprotectDiskReplicaGroupWithCallback invokes the ebs.ReprotectDiskReplicaGroup API asynchronously
func (client *Client) ReprotectDiskReplicaGroupWithCallback(request *ReprotectDiskReplicaGroupRequest, callback func(response *ReprotectDiskReplicaGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReprotectDiskReplicaGroupResponse
		var err error
		defer close(result)
		response, err = client.ReprotectDiskReplicaGroup(request)
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

// ReprotectDiskReplicaGroupRequest is the request struct for api ReprotectDiskReplicaGroup
type ReprotectDiskReplicaGroupRequest struct {
	*requests.RpcRequest
	ClientToken    string `position:"Query" name:"ClientToken"`
	SourceZoneId   string `position:"Query" name:"SourceZoneId"`
	SourceRegionId string `position:"Query" name:"SourceRegionId"`
	ReplicaGroupId string `position:"Query" name:"ReplicaGroupId"`
}

// ReprotectDiskReplicaGroupResponse is the response struct for api ReprotectDiskReplicaGroup
type ReprotectDiskReplicaGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReprotectDiskReplicaGroupRequest creates a request to invoke ReprotectDiskReplicaGroup API
func CreateReprotectDiskReplicaGroupRequest() (request *ReprotectDiskReplicaGroupRequest) {
	request = &ReprotectDiskReplicaGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ebs", "2021-07-30", "ReprotectDiskReplicaGroup", "ebs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReprotectDiskReplicaGroupResponse creates a response to parse from ReprotectDiskReplicaGroup response
func CreateReprotectDiskReplicaGroupResponse() (response *ReprotectDiskReplicaGroupResponse) {
	response = &ReprotectDiskReplicaGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}