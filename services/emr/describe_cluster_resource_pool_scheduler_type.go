package emr

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

// DescribeClusterResourcePoolSchedulerType invokes the emr.DescribeClusterResourcePoolSchedulerType API synchronously
func (client *Client) DescribeClusterResourcePoolSchedulerType(request *DescribeClusterResourcePoolSchedulerTypeRequest) (response *DescribeClusterResourcePoolSchedulerTypeResponse, err error) {
	response = CreateDescribeClusterResourcePoolSchedulerTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterResourcePoolSchedulerTypeWithChan invokes the emr.DescribeClusterResourcePoolSchedulerType API asynchronously
func (client *Client) DescribeClusterResourcePoolSchedulerTypeWithChan(request *DescribeClusterResourcePoolSchedulerTypeRequest) (<-chan *DescribeClusterResourcePoolSchedulerTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterResourcePoolSchedulerTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterResourcePoolSchedulerType(request)
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

// DescribeClusterResourcePoolSchedulerTypeWithCallback invokes the emr.DescribeClusterResourcePoolSchedulerType API asynchronously
func (client *Client) DescribeClusterResourcePoolSchedulerTypeWithCallback(request *DescribeClusterResourcePoolSchedulerTypeRequest, callback func(response *DescribeClusterResourcePoolSchedulerTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterResourcePoolSchedulerTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterResourcePoolSchedulerType(request)
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

// DescribeClusterResourcePoolSchedulerTypeRequest is the request struct for api DescribeClusterResourcePoolSchedulerType
type DescribeClusterResourcePoolSchedulerTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
}

// DescribeClusterResourcePoolSchedulerTypeResponse is the response struct for api DescribeClusterResourcePoolSchedulerType
type DescribeClusterResourcePoolSchedulerTypeResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	CurrentSchedulerType string `json:"CurrentSchedulerType" xml:"CurrentSchedulerType"`
	SupportSchedulerType string `json:"SupportSchedulerType" xml:"SupportSchedulerType"`
	DefaultSchedulerType string `json:"DefaultSchedulerType" xml:"DefaultSchedulerType"`
}

// CreateDescribeClusterResourcePoolSchedulerTypeRequest creates a request to invoke DescribeClusterResourcePoolSchedulerType API
func CreateDescribeClusterResourcePoolSchedulerTypeRequest() (request *DescribeClusterResourcePoolSchedulerTypeRequest) {
	request = &DescribeClusterResourcePoolSchedulerTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterResourcePoolSchedulerType", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClusterResourcePoolSchedulerTypeResponse creates a response to parse from DescribeClusterResourcePoolSchedulerType response
func CreateDescribeClusterResourcePoolSchedulerTypeResponse() (response *DescribeClusterResourcePoolSchedulerTypeResponse) {
	response = &DescribeClusterResourcePoolSchedulerTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}