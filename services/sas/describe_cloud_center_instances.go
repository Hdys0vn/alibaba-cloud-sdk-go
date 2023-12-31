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

// DescribeCloudCenterInstances invokes the sas.DescribeCloudCenterInstances API synchronously
func (client *Client) DescribeCloudCenterInstances(request *DescribeCloudCenterInstancesRequest) (response *DescribeCloudCenterInstancesResponse, err error) {
	response = CreateDescribeCloudCenterInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudCenterInstancesWithChan invokes the sas.DescribeCloudCenterInstances API asynchronously
func (client *Client) DescribeCloudCenterInstancesWithChan(request *DescribeCloudCenterInstancesRequest) (<-chan *DescribeCloudCenterInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudCenterInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudCenterInstances(request)
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

// DescribeCloudCenterInstancesWithCallback invokes the sas.DescribeCloudCenterInstances API asynchronously
func (client *Client) DescribeCloudCenterInstancesWithCallback(request *DescribeCloudCenterInstancesRequest, callback func(response *DescribeCloudCenterInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudCenterInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudCenterInstances(request)
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

// DescribeCloudCenterInstancesRequest is the request struct for api DescribeCloudCenterInstances
type DescribeCloudCenterInstancesRequest struct {
	*requests.RpcRequest
	Criteria                   string           `position:"Query" name:"Criteria"`
	Importance                 requests.Integer `position:"Query" name:"Importance"`
	SourceIp                   string           `position:"Query" name:"SourceIp"`
	NoPage                     requests.Boolean `position:"Query" name:"NoPage"`
	PageSize                   requests.Integer `position:"Query" name:"PageSize"`
	LogicalExp                 string           `position:"Query" name:"LogicalExp"`
	ResourceDirectoryAccountId string           `position:"Query" name:"ResourceDirectoryAccountId"`
	CurrentPage                requests.Integer `position:"Query" name:"CurrentPage"`
	MachineTypes               string           `position:"Query" name:"MachineTypes"`
	NoGroupTrace               requests.Boolean `position:"Query" name:"NoGroupTrace"`
}

// DescribeCloudCenterInstancesResponse is the response struct for api DescribeCloudCenterInstances
type DescribeCloudCenterInstancesResponse struct {
	*responses.BaseResponse
	Success   bool       `json:"Success" xml:"Success"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	PageInfo  PageInfo   `json:"PageInfo" xml:"PageInfo"`
	Instances []Instance `json:"Instances" xml:"Instances"`
}

// CreateDescribeCloudCenterInstancesRequest creates a request to invoke DescribeCloudCenterInstances API
func CreateDescribeCloudCenterInstancesRequest() (request *DescribeCloudCenterInstancesRequest) {
	request = &DescribeCloudCenterInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeCloudCenterInstances", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudCenterInstancesResponse creates a response to parse from DescribeCloudCenterInstances response
func CreateDescribeCloudCenterInstancesResponse() (response *DescribeCloudCenterInstancesResponse) {
	response = &DescribeCloudCenterInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
