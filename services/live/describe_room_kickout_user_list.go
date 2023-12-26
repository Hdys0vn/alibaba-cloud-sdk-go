package live

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

// DescribeRoomKickoutUserList invokes the live.DescribeRoomKickoutUserList API synchronously
func (client *Client) DescribeRoomKickoutUserList(request *DescribeRoomKickoutUserListRequest) (response *DescribeRoomKickoutUserListResponse, err error) {
	response = CreateDescribeRoomKickoutUserListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRoomKickoutUserListWithChan invokes the live.DescribeRoomKickoutUserList API asynchronously
func (client *Client) DescribeRoomKickoutUserListWithChan(request *DescribeRoomKickoutUserListRequest) (<-chan *DescribeRoomKickoutUserListResponse, <-chan error) {
	responseChan := make(chan *DescribeRoomKickoutUserListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRoomKickoutUserList(request)
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

// DescribeRoomKickoutUserListWithCallback invokes the live.DescribeRoomKickoutUserList API asynchronously
func (client *Client) DescribeRoomKickoutUserListWithCallback(request *DescribeRoomKickoutUserListRequest, callback func(response *DescribeRoomKickoutUserListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRoomKickoutUserListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRoomKickoutUserList(request)
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

// DescribeRoomKickoutUserListRequest is the request struct for api DescribeRoomKickoutUserList
type DescribeRoomKickoutUserListRequest struct {
	*requests.RpcRequest
	PageNum  requests.Integer `position:"Query" name:"PageNum"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	Order    string           `position:"Query" name:"Order"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	RoomId   string           `position:"Query" name:"RoomId"`
	AppId    string           `position:"Query" name:"AppId"`
}

// DescribeRoomKickoutUserListResponse is the response struct for api DescribeRoomKickoutUserList
type DescribeRoomKickoutUserListResponse struct {
	*responses.BaseResponse
	TotalPage int    `json:"TotalPage" xml:"TotalPage"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	TotalNum  int    `json:"TotalNum" xml:"TotalNum"`
	UserList  []User `json:"UserList" xml:"UserList"`
}

// CreateDescribeRoomKickoutUserListRequest creates a request to invoke DescribeRoomKickoutUserList API
func CreateDescribeRoomKickoutUserListRequest() (request *DescribeRoomKickoutUserListRequest) {
	request = &DescribeRoomKickoutUserListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeRoomKickoutUserList", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRoomKickoutUserListResponse creates a response to parse from DescribeRoomKickoutUserList response
func CreateDescribeRoomKickoutUserListResponse() (response *DescribeRoomKickoutUserListResponse) {
	response = &DescribeRoomKickoutUserListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
