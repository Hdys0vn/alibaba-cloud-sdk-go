package ens

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

// DescribeImageInfos invokes the ens.DescribeImageInfos API synchronously
func (client *Client) DescribeImageInfos(request *DescribeImageInfosRequest) (response *DescribeImageInfosResponse, err error) {
	response = CreateDescribeImageInfosResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageInfosWithChan invokes the ens.DescribeImageInfos API asynchronously
func (client *Client) DescribeImageInfosWithChan(request *DescribeImageInfosRequest) (<-chan *DescribeImageInfosResponse, <-chan error) {
	responseChan := make(chan *DescribeImageInfosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageInfos(request)
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

// DescribeImageInfosWithCallback invokes the ens.DescribeImageInfos API asynchronously
func (client *Client) DescribeImageInfosWithCallback(request *DescribeImageInfosRequest, callback func(response *DescribeImageInfosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageInfosResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageInfos(request)
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

// DescribeImageInfosRequest is the request struct for api DescribeImageInfos
type DescribeImageInfosRequest struct {
	*requests.RpcRequest
	OsType string `position:"Query" name:"OsType"`
}

// DescribeImageInfosResponse is the response struct for api DescribeImageInfos
type DescribeImageInfosResponse struct {
	*responses.BaseResponse
	Code      int                        `json:"Code" xml:"Code"`
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Images    ImagesInDescribeImageInfos `json:"Images" xml:"Images"`
}

// CreateDescribeImageInfosRequest creates a request to invoke DescribeImageInfos API
func CreateDescribeImageInfosRequest() (request *DescribeImageInfosRequest) {
	request = &DescribeImageInfosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeImageInfos", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeImageInfosResponse creates a response to parse from DescribeImageInfos response
func CreateDescribeImageInfosResponse() (response *DescribeImageInfosResponse) {
	response = &DescribeImageInfosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
