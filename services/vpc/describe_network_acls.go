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

// DescribeNetworkAcls invokes the vpc.DescribeNetworkAcls API synchronously
func (client *Client) DescribeNetworkAcls(request *DescribeNetworkAclsRequest) (response *DescribeNetworkAclsResponse, err error) {
	response = CreateDescribeNetworkAclsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkAclsWithChan invokes the vpc.DescribeNetworkAcls API asynchronously
func (client *Client) DescribeNetworkAclsWithChan(request *DescribeNetworkAclsRequest) (<-chan *DescribeNetworkAclsResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkAclsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkAcls(request)
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

// DescribeNetworkAclsWithCallback invokes the vpc.DescribeNetworkAcls API asynchronously
func (client *Client) DescribeNetworkAclsWithCallback(request *DescribeNetworkAclsRequest, callback func(response *DescribeNetworkAclsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkAclsResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkAcls(request)
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

// DescribeNetworkAclsRequest is the request struct for api DescribeNetworkAcls
type DescribeNetworkAclsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer           `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string                     `position:"Query" name:"ClientToken"`
	PageNumber           requests.Integer           `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer           `position:"Query" name:"PageSize"`
	NetworkAclId         string                     `position:"Query" name:"NetworkAclId"`
	ResourceId           string                     `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string                     `position:"Query" name:"ResourceOwnerAccount"`
	NetworkAclName       string                     `position:"Query" name:"NetworkAclName"`
	OwnerId              requests.Integer           `position:"Query" name:"OwnerId"`
	ResourceType         string                     `position:"Query" name:"ResourceType"`
	Tags                 *[]DescribeNetworkAclsTags `position:"Query" name:"Tags"  type:"Repeated"`
	VpcId                string                     `position:"Query" name:"VpcId"`
}

// DescribeNetworkAclsTags is a repeated param struct in DescribeNetworkAclsRequest
type DescribeNetworkAclsTags struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeNetworkAclsResponse is the response struct for api DescribeNetworkAcls
type DescribeNetworkAclsResponse struct {
	*responses.BaseResponse
	PageSize    string      `json:"PageSize" xml:"PageSize"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageNumber  string      `json:"PageNumber" xml:"PageNumber"`
	TotalCount  string      `json:"TotalCount" xml:"TotalCount"`
	NetworkAcls NetworkAcls `json:"NetworkAcls" xml:"NetworkAcls"`
}

// CreateDescribeNetworkAclsRequest creates a request to invoke DescribeNetworkAcls API
func CreateDescribeNetworkAclsRequest() (request *DescribeNetworkAclsRequest) {
	request = &DescribeNetworkAclsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeNetworkAcls", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkAclsResponse creates a response to parse from DescribeNetworkAcls response
func CreateDescribeNetworkAclsResponse() (response *DescribeNetworkAclsResponse) {
	response = &DescribeNetworkAclsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
