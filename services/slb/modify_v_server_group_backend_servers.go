package slb

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

// ModifyVServerGroupBackendServers invokes the slb.ModifyVServerGroupBackendServers API synchronously
func (client *Client) ModifyVServerGroupBackendServers(request *ModifyVServerGroupBackendServersRequest) (response *ModifyVServerGroupBackendServersResponse, err error) {
	response = CreateModifyVServerGroupBackendServersResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVServerGroupBackendServersWithChan invokes the slb.ModifyVServerGroupBackendServers API asynchronously
func (client *Client) ModifyVServerGroupBackendServersWithChan(request *ModifyVServerGroupBackendServersRequest) (<-chan *ModifyVServerGroupBackendServersResponse, <-chan error) {
	responseChan := make(chan *ModifyVServerGroupBackendServersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVServerGroupBackendServers(request)
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

// ModifyVServerGroupBackendServersWithCallback invokes the slb.ModifyVServerGroupBackendServers API asynchronously
func (client *Client) ModifyVServerGroupBackendServersWithCallback(request *ModifyVServerGroupBackendServersRequest, callback func(response *ModifyVServerGroupBackendServersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVServerGroupBackendServersResponse
		var err error
		defer close(result)
		response, err = client.ModifyVServerGroupBackendServers(request)
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

// ModifyVServerGroupBackendServersRequest is the request struct for api ModifyVServerGroupBackendServers
type ModifyVServerGroupBackendServersRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VServerGroupId       string           `position:"Query" name:"VServerGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	NewBackendServers    string           `position:"Query" name:"NewBackendServers"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	OldBackendServers    string           `position:"Query" name:"OldBackendServers"`
}

// ModifyVServerGroupBackendServersResponse is the response struct for api ModifyVServerGroupBackendServers
type ModifyVServerGroupBackendServersResponse struct {
	*responses.BaseResponse
	VServerGroupId string                                           `json:"VServerGroupId" xml:"VServerGroupId"`
	RequestId      string                                           `json:"RequestId" xml:"RequestId"`
	BackendServers BackendServersInModifyVServerGroupBackendServers `json:"BackendServers" xml:"BackendServers"`
}

// CreateModifyVServerGroupBackendServersRequest creates a request to invoke ModifyVServerGroupBackendServers API
func CreateModifyVServerGroupBackendServersRequest() (request *ModifyVServerGroupBackendServersRequest) {
	request = &ModifyVServerGroupBackendServersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "ModifyVServerGroupBackendServers", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVServerGroupBackendServersResponse creates a response to parse from ModifyVServerGroupBackendServers response
func CreateModifyVServerGroupBackendServersResponse() (response *ModifyVServerGroupBackendServersResponse) {
	response = &ModifyVServerGroupBackendServersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
