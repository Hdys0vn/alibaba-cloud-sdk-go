package ecd

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

// DeleteDesktopGroup invokes the ecd.DeleteDesktopGroup API synchronously
func (client *Client) DeleteDesktopGroup(request *DeleteDesktopGroupRequest) (response *DeleteDesktopGroupResponse, err error) {
	response = CreateDeleteDesktopGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDesktopGroupWithChan invokes the ecd.DeleteDesktopGroup API asynchronously
func (client *Client) DeleteDesktopGroupWithChan(request *DeleteDesktopGroupRequest) (<-chan *DeleteDesktopGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteDesktopGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDesktopGroup(request)
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

// DeleteDesktopGroupWithCallback invokes the ecd.DeleteDesktopGroup API asynchronously
func (client *Client) DeleteDesktopGroupWithCallback(request *DeleteDesktopGroupRequest, callback func(response *DeleteDesktopGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDesktopGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteDesktopGroup(request)
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

// DeleteDesktopGroupRequest is the request struct for api DeleteDesktopGroup
type DeleteDesktopGroupRequest struct {
	*requests.RpcRequest
	DesktopGroupId string `position:"Query" name:"DesktopGroupId"`
}

// DeleteDesktopGroupResponse is the response struct for api DeleteDesktopGroup
type DeleteDesktopGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDesktopGroupRequest creates a request to invoke DeleteDesktopGroup API
func CreateDeleteDesktopGroupRequest() (request *DeleteDesktopGroupRequest) {
	request = &DeleteDesktopGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DeleteDesktopGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDesktopGroupResponse creates a response to parse from DeleteDesktopGroup response
func CreateDeleteDesktopGroupResponse() (response *DeleteDesktopGroupResponse) {
	response = &DeleteDesktopGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
