package cloudapi

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

// DescribeApisBySignature invokes the cloudapi.DescribeApisBySignature API synchronously
func (client *Client) DescribeApisBySignature(request *DescribeApisBySignatureRequest) (response *DescribeApisBySignatureResponse, err error) {
	response = CreateDescribeApisBySignatureResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApisBySignatureWithChan invokes the cloudapi.DescribeApisBySignature API asynchronously
func (client *Client) DescribeApisBySignatureWithChan(request *DescribeApisBySignatureRequest) (<-chan *DescribeApisBySignatureResponse, <-chan error) {
	responseChan := make(chan *DescribeApisBySignatureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApisBySignature(request)
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

// DescribeApisBySignatureWithCallback invokes the cloudapi.DescribeApisBySignature API asynchronously
func (client *Client) DescribeApisBySignatureWithCallback(request *DescribeApisBySignatureRequest, callback func(response *DescribeApisBySignatureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApisBySignatureResponse
		var err error
		defer close(result)
		response, err = client.DescribeApisBySignature(request)
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

// DescribeApisBySignatureRequest is the request struct for api DescribeApisBySignature
type DescribeApisBySignatureRequest struct {
	*requests.RpcRequest
	SignatureId   string           `position:"Query" name:"SignatureId"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeApisBySignatureResponse is the response struct for api DescribeApisBySignature
type DescribeApisBySignatureResponse struct {
	*responses.BaseResponse
	PageNumber int                               `json:"PageNumber" xml:"PageNumber"`
	RequestId  string                            `json:"RequestId" xml:"RequestId"`
	PageSize   int                               `json:"PageSize" xml:"PageSize"`
	TotalCount int                               `json:"TotalCount" xml:"TotalCount"`
	ApiInfos   ApiInfosInDescribeApisBySignature `json:"ApiInfos" xml:"ApiInfos"`
}

// CreateDescribeApisBySignatureRequest creates a request to invoke DescribeApisBySignature API
func CreateDescribeApisBySignatureRequest() (request *DescribeApisBySignatureRequest) {
	request = &DescribeApisBySignatureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApisBySignature", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeApisBySignatureResponse creates a response to parse from DescribeApisBySignature response
func CreateDescribeApisBySignatureResponse() (response *DescribeApisBySignatureResponse) {
	response = &DescribeApisBySignatureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
