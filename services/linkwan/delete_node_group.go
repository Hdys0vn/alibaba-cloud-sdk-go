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

// DeleteNodeGroup invokes the linkwan.DeleteNodeGroup API synchronously
func (client *Client) DeleteNodeGroup(request *DeleteNodeGroupRequest) (response *DeleteNodeGroupResponse, err error) {
	response = CreateDeleteNodeGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNodeGroupWithChan invokes the linkwan.DeleteNodeGroup API asynchronously
func (client *Client) DeleteNodeGroupWithChan(request *DeleteNodeGroupRequest) (<-chan *DeleteNodeGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteNodeGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNodeGroup(request)
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

// DeleteNodeGroupWithCallback invokes the linkwan.DeleteNodeGroup API asynchronously
func (client *Client) DeleteNodeGroupWithCallback(request *DeleteNodeGroupRequest, callback func(response *DeleteNodeGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNodeGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteNodeGroup(request)
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

// DeleteNodeGroupRequest is the request struct for api DeleteNodeGroup
type DeleteNodeGroupRequest struct {
	*requests.RpcRequest
	NodeGroupId string `position:"Query" name:"NodeGroupId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// DeleteNodeGroupResponse is the response struct for api DeleteNodeGroup
type DeleteNodeGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteNodeGroupRequest creates a request to invoke DeleteNodeGroup API
func CreateDeleteNodeGroupRequest() (request *DeleteNodeGroupRequest) {
	request = &DeleteNodeGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "DeleteNodeGroup", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteNodeGroupResponse creates a response to parse from DeleteNodeGroup response
func CreateDeleteNodeGroupResponse() (response *DeleteNodeGroupResponse) {
	response = &DeleteNodeGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
