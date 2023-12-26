package cloudfw

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

// DeleteInstanceMembers invokes the cloudfw.DeleteInstanceMembers API synchronously
func (client *Client) DeleteInstanceMembers(request *DeleteInstanceMembersRequest) (response *DeleteInstanceMembersResponse, err error) {
	response = CreateDeleteInstanceMembersResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteInstanceMembersWithChan invokes the cloudfw.DeleteInstanceMembers API asynchronously
func (client *Client) DeleteInstanceMembersWithChan(request *DeleteInstanceMembersRequest) (<-chan *DeleteInstanceMembersResponse, <-chan error) {
	responseChan := make(chan *DeleteInstanceMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteInstanceMembers(request)
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

// DeleteInstanceMembersWithCallback invokes the cloudfw.DeleteInstanceMembers API asynchronously
func (client *Client) DeleteInstanceMembersWithCallback(request *DeleteInstanceMembersRequest, callback func(response *DeleteInstanceMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteInstanceMembersResponse
		var err error
		defer close(result)
		response, err = client.DeleteInstanceMembers(request)
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

// DeleteInstanceMembersRequest is the request struct for api DeleteInstanceMembers
type DeleteInstanceMembersRequest struct {
	*requests.RpcRequest
	MemberUids *[]string `position:"Query" name:"MemberUids"  type:"Repeated"`
	SourceIp   string    `position:"Query" name:"SourceIp"`
	Lang       string    `position:"Query" name:"Lang"`
}

// DeleteInstanceMembersResponse is the response struct for api DeleteInstanceMembers
type DeleteInstanceMembersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteInstanceMembersRequest creates a request to invoke DeleteInstanceMembers API
func CreateDeleteInstanceMembersRequest() (request *DeleteInstanceMembersRequest) {
	request = &DeleteInstanceMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DeleteInstanceMembers", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteInstanceMembersResponse creates a response to parse from DeleteInstanceMembers response
func CreateDeleteInstanceMembersResponse() (response *DeleteInstanceMembersResponse) {
	response = &DeleteInstanceMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}