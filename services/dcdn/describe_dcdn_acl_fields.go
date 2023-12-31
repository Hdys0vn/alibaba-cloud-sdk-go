package dcdn

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

// DescribeDcdnAclFields invokes the dcdn.DescribeDcdnAclFields API synchronously
func (client *Client) DescribeDcdnAclFields(request *DescribeDcdnAclFieldsRequest) (response *DescribeDcdnAclFieldsResponse, err error) {
	response = CreateDescribeDcdnAclFieldsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnAclFieldsWithChan invokes the dcdn.DescribeDcdnAclFields API asynchronously
func (client *Client) DescribeDcdnAclFieldsWithChan(request *DescribeDcdnAclFieldsRequest) (<-chan *DescribeDcdnAclFieldsResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnAclFieldsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnAclFields(request)
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

// DescribeDcdnAclFieldsWithCallback invokes the dcdn.DescribeDcdnAclFields API asynchronously
func (client *Client) DescribeDcdnAclFieldsWithCallback(request *DescribeDcdnAclFieldsRequest, callback func(response *DescribeDcdnAclFieldsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnAclFieldsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnAclFields(request)
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

// DescribeDcdnAclFieldsRequest is the request struct for api DescribeDcdnAclFields
type DescribeDcdnAclFieldsRequest struct {
	*requests.RpcRequest
	Lang string `position:"Query" name:"Lang"`
}

// DescribeDcdnAclFieldsResponse is the response struct for api DescribeDcdnAclFields
type DescribeDcdnAclFieldsResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Content   []ContentItem `json:"Content" xml:"Content"`
}

// CreateDescribeDcdnAclFieldsRequest creates a request to invoke DescribeDcdnAclFields API
func CreateDescribeDcdnAclFieldsRequest() (request *DescribeDcdnAclFieldsRequest) {
	request = &DescribeDcdnAclFieldsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnAclFields", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnAclFieldsResponse creates a response to parse from DescribeDcdnAclFields response
func CreateDescribeDcdnAclFieldsResponse() (response *DescribeDcdnAclFieldsResponse) {
	response = &DescribeDcdnAclFieldsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
