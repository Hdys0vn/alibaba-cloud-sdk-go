package sae

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

// DescribeGreyTagRoute invokes the sae.DescribeGreyTagRoute API synchronously
func (client *Client) DescribeGreyTagRoute(request *DescribeGreyTagRouteRequest) (response *DescribeGreyTagRouteResponse, err error) {
	response = CreateDescribeGreyTagRouteResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGreyTagRouteWithChan invokes the sae.DescribeGreyTagRoute API asynchronously
func (client *Client) DescribeGreyTagRouteWithChan(request *DescribeGreyTagRouteRequest) (<-chan *DescribeGreyTagRouteResponse, <-chan error) {
	responseChan := make(chan *DescribeGreyTagRouteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGreyTagRoute(request)
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

// DescribeGreyTagRouteWithCallback invokes the sae.DescribeGreyTagRoute API asynchronously
func (client *Client) DescribeGreyTagRouteWithCallback(request *DescribeGreyTagRouteRequest, callback func(response *DescribeGreyTagRouteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGreyTagRouteResponse
		var err error
		defer close(result)
		response, err = client.DescribeGreyTagRoute(request)
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

// DescribeGreyTagRouteRequest is the request struct for api DescribeGreyTagRoute
type DescribeGreyTagRouteRequest struct {
	*requests.RoaRequest
	GreyTagRouteId requests.Integer `position:"Query" name:"GreyTagRouteId"`
}

// DescribeGreyTagRouteResponse is the response struct for api DescribeGreyTagRoute
type DescribeGreyTagRouteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeGreyTagRouteRequest creates a request to invoke DescribeGreyTagRoute API
func CreateDescribeGreyTagRouteRequest() (request *DescribeGreyTagRouteRequest) {
	request = &DescribeGreyTagRouteRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeGreyTagRoute", "/pop/v1/sam/tagroute/greyTagRoute", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeGreyTagRouteResponse creates a response to parse from DescribeGreyTagRoute response
func CreateDescribeGreyTagRouteResponse() (response *DescribeGreyTagRouteResponse) {
	response = &DescribeGreyTagRouteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
