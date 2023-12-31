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

// GetGroups invokes the green.GetGroups API synchronously
func (client *Client) GetGroups(request *GetGroupsRequest) (response *GetGroupsResponse, err error) {
	response = CreateGetGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// GetGroupsWithChan invokes the green.GetGroups API asynchronously
func (client *Client) GetGroupsWithChan(request *GetGroupsRequest) (<-chan *GetGroupsResponse, <-chan error) {
	responseChan := make(chan *GetGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGroups(request)
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

// GetGroupsWithCallback invokes the green.GetGroups API asynchronously
func (client *Client) GetGroupsWithCallback(request *GetGroupsRequest, callback func(response *GetGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGroupsResponse
		var err error
		defer close(result)
		response, err = client.GetGroups(request)
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

// GetGroupsRequest is the request struct for api GetGroups
type GetGroupsRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// GetGroupsResponse is the response struct for api GetGroups
type GetGroupsResponse struct {
	*responses.BaseResponse
}

// CreateGetGroupsRequest creates a request to invoke GetGroups API
func CreateGetGroupsRequest() (request *GetGroupsRequest) {
	request = &GetGroupsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "GetGroups", "/green/sface/groups", "", "")
	request.Method = requests.POST
	return
}

// CreateGetGroupsResponse creates a response to parse from GetGroups response
func CreateGetGroupsResponse() (response *GetGroupsResponse) {
	response = &GetGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
