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

// ChangeResourceGroup invokes the arms.ChangeResourceGroup API synchronously
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (response *ChangeResourceGroupResponse, err error) {
	response = CreateChangeResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeResourceGroupWithChan invokes the arms.ChangeResourceGroup API asynchronously
func (client *Client) ChangeResourceGroupWithChan(request *ChangeResourceGroupRequest) (<-chan *ChangeResourceGroupResponse, <-chan error) {
	responseChan := make(chan *ChangeResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeResourceGroup(request)
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

// ChangeResourceGroupWithCallback invokes the arms.ChangeResourceGroup API asynchronously
func (client *Client) ChangeResourceGroupWithCallback(request *ChangeResourceGroupRequest, callback func(response *ChangeResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.ChangeResourceGroup(request)
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

// ChangeResourceGroupRequest is the request struct for api ChangeResourceGroup
type ChangeResourceGroupRequest struct {
	*requests.RpcRequest
	ResourceId         string `position:"Query" name:"ResourceId"`
	NewResourceGroupId string `position:"Query" name:"NewResourceGroupId"`
	ResourceType       string `position:"Query" name:"ResourceType"`
}

// ChangeResourceGroupResponse is the response struct for api ChangeResourceGroup
type ChangeResourceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateChangeResourceGroupRequest creates a request to invoke ChangeResourceGroup API
func CreateChangeResourceGroupRequest() (request *ChangeResourceGroupRequest) {
	request = &ChangeResourceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ChangeResourceGroup", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateChangeResourceGroupResponse creates a response to parse from ChangeResourceGroup response
func CreateChangeResourceGroupResponse() (response *ChangeResourceGroupResponse) {
	response = &ChangeResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
