package vpcpeer

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

// ModifyVpcPeerConnection invokes the vpcpeer.ModifyVpcPeerConnection API synchronously
func (client *Client) ModifyVpcPeerConnection(request *ModifyVpcPeerConnectionRequest) (response *ModifyVpcPeerConnectionResponse, err error) {
	response = CreateModifyVpcPeerConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpcPeerConnectionWithChan invokes the vpcpeer.ModifyVpcPeerConnection API asynchronously
func (client *Client) ModifyVpcPeerConnectionWithChan(request *ModifyVpcPeerConnectionRequest) (<-chan *ModifyVpcPeerConnectionResponse, <-chan error) {
	responseChan := make(chan *ModifyVpcPeerConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpcPeerConnection(request)
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

// ModifyVpcPeerConnectionWithCallback invokes the vpcpeer.ModifyVpcPeerConnection API asynchronously
func (client *Client) ModifyVpcPeerConnectionWithCallback(request *ModifyVpcPeerConnectionRequest, callback func(response *ModifyVpcPeerConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpcPeerConnectionResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpcPeerConnection(request)
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

// ModifyVpcPeerConnectionRequest is the request struct for api ModifyVpcPeerConnection
type ModifyVpcPeerConnectionRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Body" name:"ClientToken"`
	Channel     string           `position:"Body" name:"Channel"`
	Description string           `position:"Body" name:"Description"`
	DryRun      requests.Boolean `position:"Body" name:"DryRun"`
	Bandwidth   requests.Integer `position:"Body" name:"Bandwidth"`
	InstanceId  string           `position:"Body" name:"InstanceId"`
	Name        string           `position:"Body" name:"Name"`
}

// ModifyVpcPeerConnectionResponse is the response struct for api ModifyVpcPeerConnection
type ModifyVpcPeerConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVpcPeerConnectionRequest creates a request to invoke ModifyVpcPeerConnection API
func CreateModifyVpcPeerConnectionRequest() (request *ModifyVpcPeerConnectionRequest) {
	request = &ModifyVpcPeerConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VpcPeer", "2022-01-01", "ModifyVpcPeerConnection", "vpcpeer", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpcPeerConnectionResponse creates a response to parse from ModifyVpcPeerConnection response
func CreateModifyVpcPeerConnectionResponse() (response *ModifyVpcPeerConnectionResponse) {
	response = &ModifyVpcPeerConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
