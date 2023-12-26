package acm

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

// DescribeNamespaces invokes the acm.DescribeNamespaces API synchronously
// api document: https://help.aliyun.com/api/acm/describenamespaces.html
func (client *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
	response = CreateDescribeNamespacesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNamespacesWithChan invokes the acm.DescribeNamespaces API asynchronously
// api document: https://help.aliyun.com/api/acm/describenamespaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNamespacesWithChan(request *DescribeNamespacesRequest) (<-chan *DescribeNamespacesResponse, <-chan error) {
	responseChan := make(chan *DescribeNamespacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNamespaces(request)
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

// DescribeNamespacesWithCallback invokes the acm.DescribeNamespaces API asynchronously
// api document: https://help.aliyun.com/api/acm/describenamespaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNamespacesWithCallback(request *DescribeNamespacesRequest, callback func(response *DescribeNamespacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNamespacesResponse
		var err error
		defer close(result)
		response, err = client.DescribeNamespaces(request)
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

// DescribeNamespacesRequest is the request struct for api DescribeNamespaces
type DescribeNamespacesRequest struct {
	*requests.RoaRequest
}

// DescribeNamespacesResponse is the response struct for api DescribeNamespaces
type DescribeNamespacesResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	Code       string      `json:"Code" xml:"Code"`
	Message    string      `json:"Message" xml:"Message"`
	Namespaces []Namespace `json:"Namespaces" xml:"Namespaces"`
}

// CreateDescribeNamespacesRequest creates a request to invoke DescribeNamespaces API
func CreateDescribeNamespacesRequest() (request *DescribeNamespacesRequest) {
	request = &DescribeNamespacesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("acm", "2020-02-06", "DescribeNamespaces", "/diamond-ops/pop/namespace/list", "acms", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeNamespacesResponse creates a response to parse from DescribeNamespaces response
func CreateDescribeNamespacesResponse() (response *DescribeNamespacesResponse) {
	response = &DescribeNamespacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
