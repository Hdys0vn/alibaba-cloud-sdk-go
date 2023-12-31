package rtc

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

// DescribeChannelParticipants invokes the rtc.DescribeChannelParticipants API synchronously
func (client *Client) DescribeChannelParticipants(request *DescribeChannelParticipantsRequest) (response *DescribeChannelParticipantsResponse, err error) {
	response = CreateDescribeChannelParticipantsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeChannelParticipantsWithChan invokes the rtc.DescribeChannelParticipants API asynchronously
func (client *Client) DescribeChannelParticipantsWithChan(request *DescribeChannelParticipantsRequest) (<-chan *DescribeChannelParticipantsResponse, <-chan error) {
	responseChan := make(chan *DescribeChannelParticipantsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeChannelParticipants(request)
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

// DescribeChannelParticipantsWithCallback invokes the rtc.DescribeChannelParticipants API asynchronously
func (client *Client) DescribeChannelParticipantsWithCallback(request *DescribeChannelParticipantsRequest, callback func(response *DescribeChannelParticipantsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeChannelParticipantsResponse
		var err error
		defer close(result)
		response, err = client.DescribeChannelParticipants(request)
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

// DescribeChannelParticipantsRequest is the request struct for api DescribeChannelParticipants
type DescribeChannelParticipantsRequest struct {
	*requests.RpcRequest
	PageNum   requests.Integer `position:"Query" name:"PageNum"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	Order     string           `position:"Query" name:"Order"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	AppId     string           `position:"Query" name:"AppId"`
	ChannelId string           `position:"Query" name:"ChannelId"`
}

// DescribeChannelParticipantsResponse is the response struct for api DescribeChannelParticipants
type DescribeChannelParticipantsResponse struct {
	*responses.BaseResponse
	TotalPage int                                   `json:"TotalPage" xml:"TotalPage"`
	RequestId string                                `json:"RequestId" xml:"RequestId"`
	Timestamp int                                   `json:"Timestamp" xml:"Timestamp"`
	TotalNum  int                                   `json:"TotalNum" xml:"TotalNum"`
	UserList  UserListInDescribeChannelParticipants `json:"UserList" xml:"UserList"`
}

// CreateDescribeChannelParticipantsRequest creates a request to invoke DescribeChannelParticipants API
func CreateDescribeChannelParticipantsRequest() (request *DescribeChannelParticipantsRequest) {
	request = &DescribeChannelParticipantsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeChannelParticipants", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeChannelParticipantsResponse creates a response to parse from DescribeChannelParticipants response
func CreateDescribeChannelParticipantsResponse() (response *DescribeChannelParticipantsResponse) {
	response = &DescribeChannelParticipantsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
