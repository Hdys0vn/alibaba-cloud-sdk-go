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

// UnbindJoinPermissionFromNodeGroup invokes the linkwan.UnbindJoinPermissionFromNodeGroup API synchronously
func (client *Client) UnbindJoinPermissionFromNodeGroup(request *UnbindJoinPermissionFromNodeGroupRequest) (response *UnbindJoinPermissionFromNodeGroupResponse, err error) {
	response = CreateUnbindJoinPermissionFromNodeGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindJoinPermissionFromNodeGroupWithChan invokes the linkwan.UnbindJoinPermissionFromNodeGroup API asynchronously
func (client *Client) UnbindJoinPermissionFromNodeGroupWithChan(request *UnbindJoinPermissionFromNodeGroupRequest) (<-chan *UnbindJoinPermissionFromNodeGroupResponse, <-chan error) {
	responseChan := make(chan *UnbindJoinPermissionFromNodeGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindJoinPermissionFromNodeGroup(request)
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

// UnbindJoinPermissionFromNodeGroupWithCallback invokes the linkwan.UnbindJoinPermissionFromNodeGroup API asynchronously
func (client *Client) UnbindJoinPermissionFromNodeGroupWithCallback(request *UnbindJoinPermissionFromNodeGroupRequest, callback func(response *UnbindJoinPermissionFromNodeGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindJoinPermissionFromNodeGroupResponse
		var err error
		defer close(result)
		response, err = client.UnbindJoinPermissionFromNodeGroup(request)
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

// UnbindJoinPermissionFromNodeGroupRequest is the request struct for api UnbindJoinPermissionFromNodeGroup
type UnbindJoinPermissionFromNodeGroupRequest struct {
	*requests.RpcRequest
	JoinPermissionId string `position:"Query" name:"JoinPermissionId"`
	NodeGroupId      string `position:"Query" name:"NodeGroupId"`
	ApiProduct       string `position:"Body" name:"ApiProduct"`
	ApiRevision      string `position:"Body" name:"ApiRevision"`
}

// UnbindJoinPermissionFromNodeGroupResponse is the response struct for api UnbindJoinPermissionFromNodeGroup
type UnbindJoinPermissionFromNodeGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUnbindJoinPermissionFromNodeGroupRequest creates a request to invoke UnbindJoinPermissionFromNodeGroup API
func CreateUnbindJoinPermissionFromNodeGroupRequest() (request *UnbindJoinPermissionFromNodeGroupRequest) {
	request = &UnbindJoinPermissionFromNodeGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "UnbindJoinPermissionFromNodeGroup", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbindJoinPermissionFromNodeGroupResponse creates a response to parse from UnbindJoinPermissionFromNodeGroup response
func CreateUnbindJoinPermissionFromNodeGroupResponse() (response *UnbindJoinPermissionFromNodeGroupResponse) {
	response = &UnbindJoinPermissionFromNodeGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
