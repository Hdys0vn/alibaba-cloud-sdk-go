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

// DescribeAppsByApiProduct invokes the cloudapi.DescribeAppsByApiProduct API synchronously
func (client *Client) DescribeAppsByApiProduct(request *DescribeAppsByApiProductRequest) (response *DescribeAppsByApiProductResponse, err error) {
	response = CreateDescribeAppsByApiProductResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppsByApiProductWithChan invokes the cloudapi.DescribeAppsByApiProduct API asynchronously
func (client *Client) DescribeAppsByApiProductWithChan(request *DescribeAppsByApiProductRequest) (<-chan *DescribeAppsByApiProductResponse, <-chan error) {
	responseChan := make(chan *DescribeAppsByApiProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAppsByApiProduct(request)
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

// DescribeAppsByApiProductWithCallback invokes the cloudapi.DescribeAppsByApiProduct API asynchronously
func (client *Client) DescribeAppsByApiProductWithCallback(request *DescribeAppsByApiProductRequest, callback func(response *DescribeAppsByApiProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppsByApiProductResponse
		var err error
		defer close(result)
		response, err = client.DescribeAppsByApiProduct(request)
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

// DescribeAppsByApiProductRequest is the request struct for api DescribeAppsByApiProduct
type DescribeAppsByApiProductRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	ApiProductId  string           `position:"Query" name:"ApiProductId"`
}

// DescribeAppsByApiProductResponse is the response struct for api DescribeAppsByApiProduct
type DescribeAppsByApiProductResponse struct {
	*responses.BaseResponse
	RequestId      string                                   `json:"RequestId" xml:"RequestId"`
	PageSize       int                                      `json:"PageSize" xml:"PageSize"`
	PageNumber     int                                      `json:"PageNumber" xml:"PageNumber"`
	TotalCount     int                                      `json:"TotalCount" xml:"TotalCount"`
	AuthorizedApps AuthorizedAppsInDescribeAppsByApiProduct `json:"AuthorizedApps" xml:"AuthorizedApps"`
}

// CreateDescribeAppsByApiProductRequest creates a request to invoke DescribeAppsByApiProduct API
func CreateDescribeAppsByApiProductRequest() (request *DescribeAppsByApiProductRequest) {
	request = &DescribeAppsByApiProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAppsByApiProduct", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAppsByApiProductResponse creates a response to parse from DescribeAppsByApiProduct response
func CreateDescribeAppsByApiProductResponse() (response *DescribeAppsByApiProductResponse) {
	response = &DescribeAppsByApiProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}