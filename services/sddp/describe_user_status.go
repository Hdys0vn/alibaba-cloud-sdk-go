package sddp

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

// DescribeUserStatus invokes the sddp.DescribeUserStatus API synchronously
func (client *Client) DescribeUserStatus(request *DescribeUserStatusRequest) (response *DescribeUserStatusResponse, err error) {
	response = CreateDescribeUserStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserStatusWithChan invokes the sddp.DescribeUserStatus API asynchronously
func (client *Client) DescribeUserStatusWithChan(request *DescribeUserStatusRequest) (<-chan *DescribeUserStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeUserStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserStatus(request)
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

// DescribeUserStatusWithCallback invokes the sddp.DescribeUserStatus API asynchronously
func (client *Client) DescribeUserStatusWithCallback(request *DescribeUserStatusRequest, callback func(response *DescribeUserStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserStatus(request)
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

// DescribeUserStatusRequest is the request struct for api DescribeUserStatus
type DescribeUserStatusRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribeUserStatusResponse is the response struct for api DescribeUserStatus
type DescribeUserStatusResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	UserStatus UserStatus `json:"UserStatus" xml:"UserStatus"`
}

// CreateDescribeUserStatusRequest creates a request to invoke DescribeUserStatus API
func CreateDescribeUserStatusRequest() (request *DescribeUserStatusRequest) {
	request = &DescribeUserStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeUserStatus", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUserStatusResponse creates a response to parse from DescribeUserStatus response
func CreateDescribeUserStatusResponse() (response *DescribeUserStatusResponse) {
	response = &DescribeUserStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
