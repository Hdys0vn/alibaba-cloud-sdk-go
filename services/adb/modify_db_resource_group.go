package adb

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

// ModifyDBResourceGroup invokes the adb.ModifyDBResourceGroup API synchronously
func (client *Client) ModifyDBResourceGroup(request *ModifyDBResourceGroupRequest) (response *ModifyDBResourceGroupResponse, err error) {
	response = CreateModifyDBResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBResourceGroupWithChan invokes the adb.ModifyDBResourceGroup API asynchronously
func (client *Client) ModifyDBResourceGroupWithChan(request *ModifyDBResourceGroupRequest) (<-chan *ModifyDBResourceGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyDBResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBResourceGroup(request)
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

// ModifyDBResourceGroupWithCallback invokes the adb.ModifyDBResourceGroup API asynchronously
func (client *Client) ModifyDBResourceGroupWithCallback(request *ModifyDBResourceGroupRequest, callback func(response *ModifyDBResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBResourceGroup(request)
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

// ModifyDBResourceGroupRequest is the request struct for api ModifyDBResourceGroup
type ModifyDBResourceGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NodeNum              requests.Integer `position:"Query" name:"NodeNum"`
	GroupType            string           `position:"Query" name:"GroupType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	GroupName            string           `position:"Query" name:"GroupName"`
}

// ModifyDBResourceGroupResponse is the response struct for api ModifyDBResourceGroup
type ModifyDBResourceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBResourceGroupRequest creates a request to invoke ModifyDBResourceGroup API
func CreateModifyDBResourceGroupRequest() (request *ModifyDBResourceGroupRequest) {
	request = &ModifyDBResourceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "ModifyDBResourceGroup", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBResourceGroupResponse creates a response to parse from ModifyDBResourceGroup response
func CreateModifyDBResourceGroupResponse() (response *ModifyDBResourceGroupResponse) {
	response = &ModifyDBResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
