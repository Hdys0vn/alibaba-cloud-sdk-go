package ecs

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

// DescribeTags invokes the ecs.DescribeTags API synchronously
func (client *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
	response = CreateDescribeTagsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTagsWithChan invokes the ecs.DescribeTags API asynchronously
func (client *Client) DescribeTagsWithChan(request *DescribeTagsRequest) (<-chan *DescribeTagsResponse, <-chan error) {
	responseChan := make(chan *DescribeTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTags(request)
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

// DescribeTagsWithCallback invokes the ecs.DescribeTags API asynchronously
func (client *Client) DescribeTagsWithCallback(request *DescribeTagsRequest, callback func(response *DescribeTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTagsResponse
		var err error
		defer close(result)
		response, err = client.DescribeTags(request)
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

// DescribeTagsRequest is the request struct for api DescribeTags
type DescribeTagsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer   `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer   `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer   `position:"Query" name:"PageSize"`
	Tag                  *[]DescribeTagsTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceId           string             `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer   `position:"Query" name:"OwnerId"`
	ResourceType         string             `position:"Query" name:"ResourceType"`
	Category             string             `position:"Query" name:"Category"`
}

// DescribeTagsTag is a repeated param struct in DescribeTagsRequest
type DescribeTagsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeTagsResponse is the response struct for api DescribeTags
type DescribeTagsResponse struct {
	*responses.BaseResponse
	RequestId  string             `json:"RequestId" xml:"RequestId"`
	PageSize   int                `json:"PageSize" xml:"PageSize"`
	PageNumber int                `json:"PageNumber" xml:"PageNumber"`
	TotalCount int                `json:"TotalCount" xml:"TotalCount"`
	Tags       TagsInDescribeTags `json:"Tags" xml:"Tags"`
}

// CreateDescribeTagsRequest creates a request to invoke DescribeTags API
func CreateDescribeTagsRequest() (request *DescribeTagsRequest) {
	request = &DescribeTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeTags", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTagsResponse creates a response to parse from DescribeTags response
func CreateDescribeTagsResponse() (response *DescribeTagsResponse) {
	response = &DescribeTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
