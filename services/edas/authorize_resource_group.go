package edas

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

// AuthorizeResourceGroup invokes the edas.AuthorizeResourceGroup API synchronously
func (client *Client) AuthorizeResourceGroup(request *AuthorizeResourceGroupRequest) (response *AuthorizeResourceGroupResponse, err error) {
	response = CreateAuthorizeResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// AuthorizeResourceGroupWithChan invokes the edas.AuthorizeResourceGroup API asynchronously
func (client *Client) AuthorizeResourceGroupWithChan(request *AuthorizeResourceGroupRequest) (<-chan *AuthorizeResourceGroupResponse, <-chan error) {
	responseChan := make(chan *AuthorizeResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AuthorizeResourceGroup(request)
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

// AuthorizeResourceGroupWithCallback invokes the edas.AuthorizeResourceGroup API asynchronously
func (client *Client) AuthorizeResourceGroupWithCallback(request *AuthorizeResourceGroupRequest, callback func(response *AuthorizeResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AuthorizeResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.AuthorizeResourceGroup(request)
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

// AuthorizeResourceGroupRequest is the request struct for api AuthorizeResourceGroup
type AuthorizeResourceGroupRequest struct {
	*requests.RoaRequest
	ResourceGroupIds string `position:"Query" name:"ResourceGroupIds"`
	TargetUserId     string `position:"Query" name:"TargetUserId"`
}

// AuthorizeResourceGroupResponse is the response struct for api AuthorizeResourceGroup
type AuthorizeResourceGroupResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAuthorizeResourceGroupRequest creates a request to invoke AuthorizeResourceGroup API
func CreateAuthorizeResourceGroupRequest() (request *AuthorizeResourceGroupRequest) {
	request = &AuthorizeResourceGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "AuthorizeResourceGroup", "/pop/v5/account/authorize_res_group", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAuthorizeResourceGroupResponse creates a response to parse from AuthorizeResourceGroup response
func CreateAuthorizeResourceGroupResponse() (response *AuthorizeResourceGroupResponse) {
	response = &AuthorizeResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
