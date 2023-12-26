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

// UpdateRoamingJoinPermissionEnablingState invokes the linkwan.UpdateRoamingJoinPermissionEnablingState API synchronously
func (client *Client) UpdateRoamingJoinPermissionEnablingState(request *UpdateRoamingJoinPermissionEnablingStateRequest) (response *UpdateRoamingJoinPermissionEnablingStateResponse, err error) {
	response = CreateUpdateRoamingJoinPermissionEnablingStateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRoamingJoinPermissionEnablingStateWithChan invokes the linkwan.UpdateRoamingJoinPermissionEnablingState API asynchronously
func (client *Client) UpdateRoamingJoinPermissionEnablingStateWithChan(request *UpdateRoamingJoinPermissionEnablingStateRequest) (<-chan *UpdateRoamingJoinPermissionEnablingStateResponse, <-chan error) {
	responseChan := make(chan *UpdateRoamingJoinPermissionEnablingStateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRoamingJoinPermissionEnablingState(request)
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

// UpdateRoamingJoinPermissionEnablingStateWithCallback invokes the linkwan.UpdateRoamingJoinPermissionEnablingState API asynchronously
func (client *Client) UpdateRoamingJoinPermissionEnablingStateWithCallback(request *UpdateRoamingJoinPermissionEnablingStateRequest, callback func(response *UpdateRoamingJoinPermissionEnablingStateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRoamingJoinPermissionEnablingStateResponse
		var err error
		defer close(result)
		response, err = client.UpdateRoamingJoinPermissionEnablingState(request)
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

// UpdateRoamingJoinPermissionEnablingStateRequest is the request struct for api UpdateRoamingJoinPermissionEnablingState
type UpdateRoamingJoinPermissionEnablingStateRequest struct {
	*requests.RpcRequest
	JoinPermissionId string           `position:"Query" name:"JoinPermissionId"`
	Enabled          requests.Boolean `position:"Query" name:"Enabled"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
}

// UpdateRoamingJoinPermissionEnablingStateResponse is the response struct for api UpdateRoamingJoinPermissionEnablingState
type UpdateRoamingJoinPermissionEnablingStateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateRoamingJoinPermissionEnablingStateRequest creates a request to invoke UpdateRoamingJoinPermissionEnablingState API
func CreateUpdateRoamingJoinPermissionEnablingStateRequest() (request *UpdateRoamingJoinPermissionEnablingStateRequest) {
	request = &UpdateRoamingJoinPermissionEnablingStateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "UpdateRoamingJoinPermissionEnablingState", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateRoamingJoinPermissionEnablingStateResponse creates a response to parse from UpdateRoamingJoinPermissionEnablingState response
func CreateUpdateRoamingJoinPermissionEnablingStateResponse() (response *UpdateRoamingJoinPermissionEnablingStateResponse) {
	response = &UpdateRoamingJoinPermissionEnablingStateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
