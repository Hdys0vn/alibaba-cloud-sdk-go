package dds

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

// ModifyGlobalSecurityIPGroup invokes the dds.ModifyGlobalSecurityIPGroup API synchronously
func (client *Client) ModifyGlobalSecurityIPGroup(request *ModifyGlobalSecurityIPGroupRequest) (response *ModifyGlobalSecurityIPGroupResponse, err error) {
	response = CreateModifyGlobalSecurityIPGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGlobalSecurityIPGroupWithChan invokes the dds.ModifyGlobalSecurityIPGroup API asynchronously
func (client *Client) ModifyGlobalSecurityIPGroupWithChan(request *ModifyGlobalSecurityIPGroupRequest) (<-chan *ModifyGlobalSecurityIPGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyGlobalSecurityIPGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGlobalSecurityIPGroup(request)
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

// ModifyGlobalSecurityIPGroupWithCallback invokes the dds.ModifyGlobalSecurityIPGroup API asynchronously
func (client *Client) ModifyGlobalSecurityIPGroupWithCallback(request *ModifyGlobalSecurityIPGroupRequest, callback func(response *ModifyGlobalSecurityIPGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGlobalSecurityIPGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyGlobalSecurityIPGroup(request)
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

// ModifyGlobalSecurityIPGroupRequest is the request struct for api ModifyGlobalSecurityIPGroup
type ModifyGlobalSecurityIPGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	GIpList               string           `position:"Query" name:"GIpList"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	GlobalSecurityGroupId string           `position:"Query" name:"GlobalSecurityGroupId"`
	SecurityToken         string           `position:"Query" name:"SecurityToken"`
	SecurityIPType        string           `position:"Query" name:"SecurityIPType"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	WhitelistNetType      string           `position:"Query" name:"WhitelistNetType"`
	GlobalIgName          string           `position:"Query" name:"GlobalIgName"`
}

// ModifyGlobalSecurityIPGroupResponse is the response struct for api ModifyGlobalSecurityIPGroup
type ModifyGlobalSecurityIPGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyGlobalSecurityIPGroupRequest creates a request to invoke ModifyGlobalSecurityIPGroup API
func CreateModifyGlobalSecurityIPGroupRequest() (request *ModifyGlobalSecurityIPGroupRequest) {
	request = &ModifyGlobalSecurityIPGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyGlobalSecurityIPGroup", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyGlobalSecurityIPGroupResponse creates a response to parse from ModifyGlobalSecurityIPGroup response
func CreateModifyGlobalSecurityIPGroupResponse() (response *ModifyGlobalSecurityIPGroupResponse) {
	response = &ModifyGlobalSecurityIPGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
