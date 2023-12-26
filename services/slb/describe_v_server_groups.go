package slb

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

// DescribeVServerGroups invokes the slb.DescribeVServerGroups API synchronously
func (client *Client) DescribeVServerGroups(request *DescribeVServerGroupsRequest) (response *DescribeVServerGroupsResponse, err error) {
	response = CreateDescribeVServerGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVServerGroupsWithChan invokes the slb.DescribeVServerGroups API asynchronously
func (client *Client) DescribeVServerGroupsWithChan(request *DescribeVServerGroupsRequest) (<-chan *DescribeVServerGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeVServerGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVServerGroups(request)
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

// DescribeVServerGroupsWithCallback invokes the slb.DescribeVServerGroups API asynchronously
func (client *Client) DescribeVServerGroupsWithCallback(request *DescribeVServerGroupsRequest, callback func(response *DescribeVServerGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVServerGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeVServerGroups(request)
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

// DescribeVServerGroupsRequest is the request struct for api DescribeVServerGroups
type DescribeVServerGroupsRequest struct {
	*requests.RpcRequest
	AccessKeyId          string                      `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer            `position:"Query" name:"ResourceOwnerId"`
	IncludeListener      requests.Boolean            `position:"Query" name:"IncludeListener"`
	IncludeRule          requests.Boolean            `position:"Query" name:"IncludeRule"`
	Tag                  *[]DescribeVServerGroupsTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount string                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                      `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer            `position:"Query" name:"OwnerId"`
	Tags                 string                      `position:"Query" name:"Tags"`
	LoadBalancerId       string                      `position:"Query" name:"LoadBalancerId"`
}

// DescribeVServerGroupsTag is a repeated param struct in DescribeVServerGroupsRequest
type DescribeVServerGroupsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeVServerGroupsResponse is the response struct for api DescribeVServerGroups
type DescribeVServerGroupsResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	VServerGroups VServerGroups `json:"VServerGroups" xml:"VServerGroups"`
}

// CreateDescribeVServerGroupsRequest creates a request to invoke DescribeVServerGroups API
func CreateDescribeVServerGroupsRequest() (request *DescribeVServerGroupsRequest) {
	request = &DescribeVServerGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeVServerGroups", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVServerGroupsResponse creates a response to parse from DescribeVServerGroups response
func CreateDescribeVServerGroupsResponse() (response *DescribeVServerGroupsResponse) {
	response = &DescribeVServerGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}