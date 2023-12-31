package linkwan

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

// UpdateOwnedLocalJoinPermissionEnablingState invokes the linkwan.UpdateOwnedLocalJoinPermissionEnablingState API synchronously
func (client *Client) UpdateOwnedLocalJoinPermissionEnablingState(request *UpdateOwnedLocalJoinPermissionEnablingStateRequest) (response *UpdateOwnedLocalJoinPermissionEnablingStateResponse, err error) {
	response = CreateUpdateOwnedLocalJoinPermissionEnablingStateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateOwnedLocalJoinPermissionEnablingStateWithChan invokes the linkwan.UpdateOwnedLocalJoinPermissionEnablingState API asynchronously
func (client *Client) UpdateOwnedLocalJoinPermissionEnablingStateWithChan(request *UpdateOwnedLocalJoinPermissionEnablingStateRequest) (<-chan *UpdateOwnedLocalJoinPermissionEnablingStateResponse, <-chan error) {
	responseChan := make(chan *UpdateOwnedLocalJoinPermissionEnablingStateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateOwnedLocalJoinPermissionEnablingState(request)
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

// UpdateOwnedLocalJoinPermissionEnablingStateWithCallback invokes the linkwan.UpdateOwnedLocalJoinPermissionEnablingState API asynchronously
func (client *Client) UpdateOwnedLocalJoinPermissionEnablingStateWithCallback(request *UpdateOwnedLocalJoinPermissionEnablingStateRequest, callback func(response *UpdateOwnedLocalJoinPermissionEnablingStateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateOwnedLocalJoinPermissionEnablingStateResponse
		var err error
		defer close(result)
		response, err = client.UpdateOwnedLocalJoinPermissionEnablingState(request)
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

// UpdateOwnedLocalJoinPermissionEnablingStateRequest is the request struct for api UpdateOwnedLocalJoinPermissionEnablingState
type UpdateOwnedLocalJoinPermissionEnablingStateRequest struct {
	*requests.RpcRequest
	JoinPermissionId string           `position:"Query" name:"JoinPermissionId"`
	Enabled          requests.Boolean `position:"Query" name:"Enabled"`
	IotInstanceId    string           `position:"Query" name:"IotInstanceId"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
}

// UpdateOwnedLocalJoinPermissionEnablingStateResponse is the response struct for api UpdateOwnedLocalJoinPermissionEnablingState
type UpdateOwnedLocalJoinPermissionEnablingStateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateOwnedLocalJoinPermissionEnablingStateRequest creates a request to invoke UpdateOwnedLocalJoinPermissionEnablingState API
func CreateUpdateOwnedLocalJoinPermissionEnablingStateRequest() (request *UpdateOwnedLocalJoinPermissionEnablingStateRequest) {
	request = &UpdateOwnedLocalJoinPermissionEnablingStateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "UpdateOwnedLocalJoinPermissionEnablingState", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateOwnedLocalJoinPermissionEnablingStateResponse creates a response to parse from UpdateOwnedLocalJoinPermissionEnablingState response
func CreateUpdateOwnedLocalJoinPermissionEnablingStateResponse() (response *UpdateOwnedLocalJoinPermissionEnablingStateResponse) {
	response = &UpdateOwnedLocalJoinPermissionEnablingStateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
