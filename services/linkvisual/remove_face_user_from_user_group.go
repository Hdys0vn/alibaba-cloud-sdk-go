package linkvisual

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

// RemoveFaceUserFromUserGroup invokes the linkvisual.RemoveFaceUserFromUserGroup API synchronously
func (client *Client) RemoveFaceUserFromUserGroup(request *RemoveFaceUserFromUserGroupRequest) (response *RemoveFaceUserFromUserGroupResponse, err error) {
	response = CreateRemoveFaceUserFromUserGroupResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveFaceUserFromUserGroupWithChan invokes the linkvisual.RemoveFaceUserFromUserGroup API asynchronously
func (client *Client) RemoveFaceUserFromUserGroupWithChan(request *RemoveFaceUserFromUserGroupRequest) (<-chan *RemoveFaceUserFromUserGroupResponse, <-chan error) {
	responseChan := make(chan *RemoveFaceUserFromUserGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveFaceUserFromUserGroup(request)
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

// RemoveFaceUserFromUserGroupWithCallback invokes the linkvisual.RemoveFaceUserFromUserGroup API asynchronously
func (client *Client) RemoveFaceUserFromUserGroupWithCallback(request *RemoveFaceUserFromUserGroupRequest, callback func(response *RemoveFaceUserFromUserGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveFaceUserFromUserGroupResponse
		var err error
		defer close(result)
		response, err = client.RemoveFaceUserFromUserGroup(request)
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

// RemoveFaceUserFromUserGroupRequest is the request struct for api RemoveFaceUserFromUserGroup
type RemoveFaceUserFromUserGroupRequest struct {
	*requests.RpcRequest
	IsolationId string `position:"Query" name:"IsolationId"`
	UserId      string `position:"Query" name:"UserId"`
	UserGroupId string `position:"Query" name:"UserGroupId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// RemoveFaceUserFromUserGroupResponse is the response struct for api RemoveFaceUserFromUserGroup
type RemoveFaceUserFromUserGroupResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateRemoveFaceUserFromUserGroupRequest creates a request to invoke RemoveFaceUserFromUserGroup API
func CreateRemoveFaceUserFromUserGroupRequest() (request *RemoveFaceUserFromUserGroupRequest) {
	request = &RemoveFaceUserFromUserGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "RemoveFaceUserFromUserGroup", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveFaceUserFromUserGroupResponse creates a response to parse from RemoveFaceUserFromUserGroup response
func CreateRemoveFaceUserFromUserGroupResponse() (response *RemoveFaceUserFromUserGroupResponse) {
	response = &RemoveFaceUserFromUserGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
