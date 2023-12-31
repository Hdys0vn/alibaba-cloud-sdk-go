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

// DescribeImageFromFamily invokes the ecs.DescribeImageFromFamily API synchronously
func (client *Client) DescribeImageFromFamily(request *DescribeImageFromFamilyRequest) (response *DescribeImageFromFamilyResponse, err error) {
	response = CreateDescribeImageFromFamilyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageFromFamilyWithChan invokes the ecs.DescribeImageFromFamily API asynchronously
func (client *Client) DescribeImageFromFamilyWithChan(request *DescribeImageFromFamilyRequest) (<-chan *DescribeImageFromFamilyResponse, <-chan error) {
	responseChan := make(chan *DescribeImageFromFamilyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageFromFamily(request)
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

// DescribeImageFromFamilyWithCallback invokes the ecs.DescribeImageFromFamily API asynchronously
func (client *Client) DescribeImageFromFamilyWithCallback(request *DescribeImageFromFamilyRequest, callback func(response *DescribeImageFromFamilyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageFromFamilyResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageFromFamily(request)
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

// DescribeImageFromFamilyRequest is the request struct for api DescribeImageFromFamily
type DescribeImageFromFamilyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ImageFamily          string           `position:"Query" name:"ImageFamily"`
}

// DescribeImageFromFamilyResponse is the response struct for api DescribeImageFromFamily
type DescribeImageFromFamilyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Image     Image  `json:"Image" xml:"Image"`
}

// CreateDescribeImageFromFamilyRequest creates a request to invoke DescribeImageFromFamily API
func CreateDescribeImageFromFamilyRequest() (request *DescribeImageFromFamilyRequest) {
	request = &DescribeImageFromFamilyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImageFromFamily", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeImageFromFamilyResponse creates a response to parse from DescribeImageFromFamily response
func CreateDescribeImageFromFamilyResponse() (response *DescribeImageFromFamilyResponse) {
	response = &DescribeImageFromFamilyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
