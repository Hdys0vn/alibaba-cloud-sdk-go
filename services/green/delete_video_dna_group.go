package green

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

// DeleteVideoDnaGroup invokes the green.DeleteVideoDnaGroup API synchronously
func (client *Client) DeleteVideoDnaGroup(request *DeleteVideoDnaGroupRequest) (response *DeleteVideoDnaGroupResponse, err error) {
	response = CreateDeleteVideoDnaGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVideoDnaGroupWithChan invokes the green.DeleteVideoDnaGroup API asynchronously
func (client *Client) DeleteVideoDnaGroupWithChan(request *DeleteVideoDnaGroupRequest) (<-chan *DeleteVideoDnaGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteVideoDnaGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVideoDnaGroup(request)
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

// DeleteVideoDnaGroupWithCallback invokes the green.DeleteVideoDnaGroup API asynchronously
func (client *Client) DeleteVideoDnaGroupWithCallback(request *DeleteVideoDnaGroupRequest, callback func(response *DeleteVideoDnaGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVideoDnaGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteVideoDnaGroup(request)
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

// DeleteVideoDnaGroupRequest is the request struct for api DeleteVideoDnaGroup
type DeleteVideoDnaGroupRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// DeleteVideoDnaGroupResponse is the response struct for api DeleteVideoDnaGroup
type DeleteVideoDnaGroupResponse struct {
	*responses.BaseResponse
}

// CreateDeleteVideoDnaGroupRequest creates a request to invoke DeleteVideoDnaGroup API
func CreateDeleteVideoDnaGroupRequest() (request *DeleteVideoDnaGroupRequest) {
	request = &DeleteVideoDnaGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "DeleteVideoDnaGroup", "/green/video/dna/group/delete", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteVideoDnaGroupResponse creates a response to parse from DeleteVideoDnaGroup response
func CreateDeleteVideoDnaGroupResponse() (response *DeleteVideoDnaGroupResponse) {
	response = &DeleteVideoDnaGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
