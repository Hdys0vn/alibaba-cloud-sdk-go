package vpc

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

// ListVpcEndpointServicesByEndUser invokes the vpc.ListVpcEndpointServicesByEndUser API synchronously
func (client *Client) ListVpcEndpointServicesByEndUser(request *ListVpcEndpointServicesByEndUserRequest) (response *ListVpcEndpointServicesByEndUserResponse, err error) {
	response = CreateListVpcEndpointServicesByEndUserResponse()
	err = client.DoAction(request, response)
	return
}

// ListVpcEndpointServicesByEndUserWithChan invokes the vpc.ListVpcEndpointServicesByEndUser API asynchronously
func (client *Client) ListVpcEndpointServicesByEndUserWithChan(request *ListVpcEndpointServicesByEndUserRequest) (<-chan *ListVpcEndpointServicesByEndUserResponse, <-chan error) {
	responseChan := make(chan *ListVpcEndpointServicesByEndUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVpcEndpointServicesByEndUser(request)
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

// ListVpcEndpointServicesByEndUserWithCallback invokes the vpc.ListVpcEndpointServicesByEndUser API asynchronously
func (client *Client) ListVpcEndpointServicesByEndUserWithCallback(request *ListVpcEndpointServicesByEndUserRequest, callback func(response *ListVpcEndpointServicesByEndUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVpcEndpointServicesByEndUserResponse
		var err error
		defer close(result)
		response, err = client.ListVpcEndpointServicesByEndUser(request)
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

// ListVpcEndpointServicesByEndUserRequest is the request struct for api ListVpcEndpointServicesByEndUser
type ListVpcEndpointServicesByEndUserRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NextToken            string           `position:"Query" name:"NextToken"`
	MaxResults           requests.Integer `position:"Query" name:"MaxResults"`
	ServiceName          string           `position:"Query" name:"ServiceName"`
}

// ListVpcEndpointServicesByEndUserResponse is the response struct for api ListVpcEndpointServicesByEndUser
type ListVpcEndpointServicesByEndUserResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	NextToken  string    `json:"NextToken" xml:"NextToken"`
	MaxResults int64     `json:"MaxResults" xml:"MaxResults"`
	Services   []Service `json:"Services" xml:"Services"`
}

// CreateListVpcEndpointServicesByEndUserRequest creates a request to invoke ListVpcEndpointServicesByEndUser API
func CreateListVpcEndpointServicesByEndUserRequest() (request *ListVpcEndpointServicesByEndUserRequest) {
	request = &ListVpcEndpointServicesByEndUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ListVpcEndpointServicesByEndUser", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVpcEndpointServicesByEndUserResponse creates a response to parse from ListVpcEndpointServicesByEndUser response
func CreateListVpcEndpointServicesByEndUserResponse() (response *ListVpcEndpointServicesByEndUserResponse) {
	response = &ListVpcEndpointServicesByEndUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}