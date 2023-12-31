package arms

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

// CreateOrUpdateContactGroup invokes the arms.CreateOrUpdateContactGroup API synchronously
func (client *Client) CreateOrUpdateContactGroup(request *CreateOrUpdateContactGroupRequest) (response *CreateOrUpdateContactGroupResponse, err error) {
	response = CreateCreateOrUpdateContactGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrUpdateContactGroupWithChan invokes the arms.CreateOrUpdateContactGroup API asynchronously
func (client *Client) CreateOrUpdateContactGroupWithChan(request *CreateOrUpdateContactGroupRequest) (<-chan *CreateOrUpdateContactGroupResponse, <-chan error) {
	responseChan := make(chan *CreateOrUpdateContactGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrUpdateContactGroup(request)
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

// CreateOrUpdateContactGroupWithCallback invokes the arms.CreateOrUpdateContactGroup API asynchronously
func (client *Client) CreateOrUpdateContactGroupWithCallback(request *CreateOrUpdateContactGroupRequest, callback func(response *CreateOrUpdateContactGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrUpdateContactGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateOrUpdateContactGroup(request)
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

// CreateOrUpdateContactGroupRequest is the request struct for api CreateOrUpdateContactGroup
type CreateOrUpdateContactGroupRequest struct {
	*requests.RpcRequest
	ContactGroupId   requests.Integer `position:"Body" name:"ContactGroupId"`
	ContactGroupName string           `position:"Body" name:"ContactGroupName"`
	ProxyUserId      string           `position:"Body" name:"ProxyUserId"`
	ContactIds       string           `position:"Body" name:"ContactIds"`
}

// CreateOrUpdateContactGroupResponse is the response struct for api CreateOrUpdateContactGroup
type CreateOrUpdateContactGroupResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	AlertContactGroup AlertContactGroup `json:"AlertContactGroup" xml:"AlertContactGroup"`
}

// CreateCreateOrUpdateContactGroupRequest creates a request to invoke CreateOrUpdateContactGroup API
func CreateCreateOrUpdateContactGroupRequest() (request *CreateOrUpdateContactGroupRequest) {
	request = &CreateOrUpdateContactGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreateOrUpdateContactGroup", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateOrUpdateContactGroupResponse creates a response to parse from CreateOrUpdateContactGroup response
func CreateCreateOrUpdateContactGroupResponse() (response *CreateOrUpdateContactGroupResponse) {
	response = &CreateOrUpdateContactGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
