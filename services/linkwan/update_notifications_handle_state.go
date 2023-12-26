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

// UpdateNotificationsHandleState invokes the linkwan.UpdateNotificationsHandleState API synchronously
func (client *Client) UpdateNotificationsHandleState(request *UpdateNotificationsHandleStateRequest) (response *UpdateNotificationsHandleStateResponse, err error) {
	response = CreateUpdateNotificationsHandleStateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateNotificationsHandleStateWithChan invokes the linkwan.UpdateNotificationsHandleState API asynchronously
func (client *Client) UpdateNotificationsHandleStateWithChan(request *UpdateNotificationsHandleStateRequest) (<-chan *UpdateNotificationsHandleStateResponse, <-chan error) {
	responseChan := make(chan *UpdateNotificationsHandleStateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateNotificationsHandleState(request)
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

// UpdateNotificationsHandleStateWithCallback invokes the linkwan.UpdateNotificationsHandleState API asynchronously
func (client *Client) UpdateNotificationsHandleStateWithCallback(request *UpdateNotificationsHandleStateRequest, callback func(response *UpdateNotificationsHandleStateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateNotificationsHandleStateResponse
		var err error
		defer close(result)
		response, err = client.UpdateNotificationsHandleState(request)
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

// UpdateNotificationsHandleStateRequest is the request struct for api UpdateNotificationsHandleState
type UpdateNotificationsHandleStateRequest struct {
	*requests.RpcRequest
	TargetHandleState string    `position:"Query" name:"TargetHandleState"`
	ApiProduct        string    `position:"Body" name:"ApiProduct"`
	ApiRevision       string    `position:"Body" name:"ApiRevision"`
	NotificationId    *[]string `position:"Query" name:"NotificationId"  type:"Repeated"`
}

// UpdateNotificationsHandleStateResponse is the response struct for api UpdateNotificationsHandleState
type UpdateNotificationsHandleStateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateNotificationsHandleStateRequest creates a request to invoke UpdateNotificationsHandleState API
func CreateUpdateNotificationsHandleStateRequest() (request *UpdateNotificationsHandleStateRequest) {
	request = &UpdateNotificationsHandleStateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "UpdateNotificationsHandleState", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateNotificationsHandleStateResponse creates a response to parse from UpdateNotificationsHandleState response
func CreateUpdateNotificationsHandleStateResponse() (response *UpdateNotificationsHandleStateResponse) {
	response = &UpdateNotificationsHandleStateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}