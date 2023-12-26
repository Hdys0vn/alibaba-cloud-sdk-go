package sas

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

// DescribeExposedInstanceDetail invokes the sas.DescribeExposedInstanceDetail API synchronously
func (client *Client) DescribeExposedInstanceDetail(request *DescribeExposedInstanceDetailRequest) (response *DescribeExposedInstanceDetailResponse, err error) {
	response = CreateDescribeExposedInstanceDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeExposedInstanceDetailWithChan invokes the sas.DescribeExposedInstanceDetail API asynchronously
func (client *Client) DescribeExposedInstanceDetailWithChan(request *DescribeExposedInstanceDetailRequest) (<-chan *DescribeExposedInstanceDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeExposedInstanceDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExposedInstanceDetail(request)
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

// DescribeExposedInstanceDetailWithCallback invokes the sas.DescribeExposedInstanceDetail API asynchronously
func (client *Client) DescribeExposedInstanceDetailWithCallback(request *DescribeExposedInstanceDetailRequest, callback func(response *DescribeExposedInstanceDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExposedInstanceDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeExposedInstanceDetail(request)
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

// DescribeExposedInstanceDetailRequest is the request struct for api DescribeExposedInstanceDetail
type DescribeExposedInstanceDetailRequest struct {
	*requests.RpcRequest
	Uuid     string `position:"Query" name:"Uuid"`
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeExposedInstanceDetailResponse is the response struct for api DescribeExposedInstanceDetail
type DescribeExposedInstanceDetailResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	ExposedChains []ExposedChain `json:"ExposedChains" xml:"ExposedChains"`
}

// CreateDescribeExposedInstanceDetailRequest creates a request to invoke DescribeExposedInstanceDetail API
func CreateDescribeExposedInstanceDetailRequest() (request *DescribeExposedInstanceDetailRequest) {
	request = &DescribeExposedInstanceDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeExposedInstanceDetail", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeExposedInstanceDetailResponse creates a response to parse from DescribeExposedInstanceDetail response
func CreateDescribeExposedInstanceDetailResponse() (response *DescribeExposedInstanceDetailResponse) {
	response = &DescribeExposedInstanceDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}