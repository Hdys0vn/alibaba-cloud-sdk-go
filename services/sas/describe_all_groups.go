package sas

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

// DescribeAllGroups invokes the sas.DescribeAllGroups API synchronously
func (client *Client) DescribeAllGroups(request *DescribeAllGroupsRequest) (response *DescribeAllGroupsResponse, err error) {
	response = CreateDescribeAllGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAllGroupsWithChan invokes the sas.DescribeAllGroups API asynchronously
func (client *Client) DescribeAllGroupsWithChan(request *DescribeAllGroupsRequest) (<-chan *DescribeAllGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeAllGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAllGroups(request)
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

// DescribeAllGroupsWithCallback invokes the sas.DescribeAllGroups API asynchronously
func (client *Client) DescribeAllGroupsWithCallback(request *DescribeAllGroupsRequest, callback func(response *DescribeAllGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAllGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAllGroups(request)
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

// DescribeAllGroupsRequest is the request struct for api DescribeAllGroups
type DescribeAllGroupsRequest struct {
	*requests.RpcRequest
	SourceIp                   string `position:"Query" name:"SourceIp"`
	Lang                       string `position:"Query" name:"Lang"`
	ResourceDirectoryAccountId string `position:"Query" name:"ResourceDirectoryAccountId"`
}

// DescribeAllGroupsResponse is the response struct for api DescribeAllGroups
type DescribeAllGroupsResponse struct {
	*responses.BaseResponse
	Count     int     `json:"Count" xml:"Count"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Groups    []Group `json:"Groups" xml:"Groups"`
}

// CreateDescribeAllGroupsRequest creates a request to invoke DescribeAllGroups API
func CreateDescribeAllGroupsRequest() (request *DescribeAllGroupsRequest) {
	request = &DescribeAllGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeAllGroups", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAllGroupsResponse creates a response to parse from DescribeAllGroups response
func CreateDescribeAllGroupsResponse() (response *DescribeAllGroupsResponse) {
	response = &DescribeAllGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
