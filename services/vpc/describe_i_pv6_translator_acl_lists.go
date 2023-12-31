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

// DescribeIPv6TranslatorAclLists invokes the vpc.DescribeIPv6TranslatorAclLists API synchronously
func (client *Client) DescribeIPv6TranslatorAclLists(request *DescribeIPv6TranslatorAclListsRequest) (response *DescribeIPv6TranslatorAclListsResponse, err error) {
	response = CreateDescribeIPv6TranslatorAclListsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIPv6TranslatorAclListsWithChan invokes the vpc.DescribeIPv6TranslatorAclLists API asynchronously
func (client *Client) DescribeIPv6TranslatorAclListsWithChan(request *DescribeIPv6TranslatorAclListsRequest) (<-chan *DescribeIPv6TranslatorAclListsResponse, <-chan error) {
	responseChan := make(chan *DescribeIPv6TranslatorAclListsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIPv6TranslatorAclLists(request)
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

// DescribeIPv6TranslatorAclListsWithCallback invokes the vpc.DescribeIPv6TranslatorAclLists API asynchronously
func (client *Client) DescribeIPv6TranslatorAclListsWithCallback(request *DescribeIPv6TranslatorAclListsRequest, callback func(response *DescribeIPv6TranslatorAclListsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIPv6TranslatorAclListsResponse
		var err error
		defer close(result)
		response, err = client.DescribeIPv6TranslatorAclLists(request)
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

// DescribeIPv6TranslatorAclListsRequest is the request struct for api DescribeIPv6TranslatorAclLists
type DescribeIPv6TranslatorAclListsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AclName              string           `position:"Query" name:"AclName"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeIPv6TranslatorAclListsResponse is the response struct for api DescribeIPv6TranslatorAclLists
type DescribeIPv6TranslatorAclListsResponse struct {
	*responses.BaseResponse
	PageSize           int                `json:"PageSize" xml:"PageSize"`
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	PageNumber         int                `json:"PageNumber" xml:"PageNumber"`
	TotalCount         int                `json:"TotalCount" xml:"TotalCount"`
	Ipv6TranslatorAcls Ipv6TranslatorAcls `json:"Ipv6TranslatorAcls" xml:"Ipv6TranslatorAcls"`
}

// CreateDescribeIPv6TranslatorAclListsRequest creates a request to invoke DescribeIPv6TranslatorAclLists API
func CreateDescribeIPv6TranslatorAclListsRequest() (request *DescribeIPv6TranslatorAclListsRequest) {
	request = &DescribeIPv6TranslatorAclListsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeIPv6TranslatorAclLists", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIPv6TranslatorAclListsResponse creates a response to parse from DescribeIPv6TranslatorAclLists response
func CreateDescribeIPv6TranslatorAclListsResponse() (response *DescribeIPv6TranslatorAclListsResponse) {
	response = &DescribeIPv6TranslatorAclListsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
