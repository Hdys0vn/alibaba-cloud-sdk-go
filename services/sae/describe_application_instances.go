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

// DescribeApplicationInstances invokes the sae.DescribeApplicationInstances API synchronously
func (client *Client) DescribeApplicationInstances(request *DescribeApplicationInstancesRequest) (response *DescribeApplicationInstancesResponse, err error) {
	response = CreateDescribeApplicationInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApplicationInstancesWithChan invokes the sae.DescribeApplicationInstances API asynchronously
func (client *Client) DescribeApplicationInstancesWithChan(request *DescribeApplicationInstancesRequest) (<-chan *DescribeApplicationInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeApplicationInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApplicationInstances(request)
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

// DescribeApplicationInstancesWithCallback invokes the sae.DescribeApplicationInstances API asynchronously
func (client *Client) DescribeApplicationInstancesWithCallback(request *DescribeApplicationInstancesRequest, callback func(response *DescribeApplicationInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApplicationInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeApplicationInstances(request)
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

// DescribeApplicationInstancesRequest is the request struct for api DescribeApplicationInstances
type DescribeApplicationInstancesRequest struct {
	*requests.RoaRequest
	JobId       string           `position:"Query" name:"JobId"`
	AppId       string           `position:"Query" name:"AppId"`
	GroupId     string           `position:"Query" name:"GroupId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Reverse     requests.Boolean `position:"Query" name:"Reverse"`
}

// DescribeApplicationInstancesResponse is the response struct for api DescribeApplicationInstances
type DescribeApplicationInstancesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeApplicationInstancesRequest creates a request to invoke DescribeApplicationInstances API
func CreateDescribeApplicationInstancesRequest() (request *DescribeApplicationInstancesRequest) {
	request = &DescribeApplicationInstancesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeApplicationInstances", "/pop/v1/sam/app/describeApplicationInstances", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeApplicationInstancesResponse creates a response to parse from DescribeApplicationInstances response
func CreateDescribeApplicationInstancesResponse() (response *DescribeApplicationInstancesResponse) {
	response = &DescribeApplicationInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
