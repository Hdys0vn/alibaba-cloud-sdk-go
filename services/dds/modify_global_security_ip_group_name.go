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

// ModifyGlobalSecurityIPGroupName invokes the dds.ModifyGlobalSecurityIPGroupName API synchronously
func (client *Client) ModifyGlobalSecurityIPGroupName(request *ModifyGlobalSecurityIPGroupNameRequest) (response *ModifyGlobalSecurityIPGroupNameResponse, err error) {
	response = CreateModifyGlobalSecurityIPGroupNameResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGlobalSecurityIPGroupNameWithChan invokes the dds.ModifyGlobalSecurityIPGroupName API asynchronously
func (client *Client) ModifyGlobalSecurityIPGroupNameWithChan(request *ModifyGlobalSecurityIPGroupNameRequest) (<-chan *ModifyGlobalSecurityIPGroupNameResponse, <-chan error) {
	responseChan := make(chan *ModifyGlobalSecurityIPGroupNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGlobalSecurityIPGroupName(request)
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

// ModifyGlobalSecurityIPGroupNameWithCallback invokes the dds.ModifyGlobalSecurityIPGroupName API asynchronously
func (client *Client) ModifyGlobalSecurityIPGroupNameWithCallback(request *ModifyGlobalSecurityIPGroupNameRequest, callback func(response *ModifyGlobalSecurityIPGroupNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGlobalSecurityIPGroupNameResponse
		var err error
		defer close(result)
		response, err = client.ModifyGlobalSecurityIPGroupName(request)
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

// ModifyGlobalSecurityIPGroupNameRequest is the request struct for api ModifyGlobalSecurityIPGroupName
type ModifyGlobalSecurityIPGroupNameRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	GlobalSecurityGroupId string           `position:"Query" name:"GlobalSecurityGroupId"`
	SecurityToken         string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	GlobalIgName          string           `position:"Query" name:"GlobalIgName"`
}

// ModifyGlobalSecurityIPGroupNameResponse is the response struct for api ModifyGlobalSecurityIPGroupName
type ModifyGlobalSecurityIPGroupNameResponse struct {
	*responses.BaseResponse
	RequestId             string                      `json:"RequestId" xml:"RequestId"`
	GlobalSecurityIPGroup []GlobalSecurityIPGroupItem `json:"GlobalSecurityIPGroup" xml:"GlobalSecurityIPGroup"`
}

// CreateModifyGlobalSecurityIPGroupNameRequest creates a request to invoke ModifyGlobalSecurityIPGroupName API
func CreateModifyGlobalSecurityIPGroupNameRequest() (request *ModifyGlobalSecurityIPGroupNameRequest) {
	request = &ModifyGlobalSecurityIPGroupNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyGlobalSecurityIPGroupName", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyGlobalSecurityIPGroupNameResponse creates a response to parse from ModifyGlobalSecurityIPGroupName response
func CreateModifyGlobalSecurityIPGroupNameResponse() (response *ModifyGlobalSecurityIPGroupNameResponse) {
	response = &ModifyGlobalSecurityIPGroupNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
