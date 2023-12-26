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

// DescribeGroupedContainerInstances invokes the sas.DescribeGroupedContainerInstances API synchronously
func (client *Client) DescribeGroupedContainerInstances(request *DescribeGroupedContainerInstancesRequest) (response *DescribeGroupedContainerInstancesResponse, err error) {
	response = CreateDescribeGroupedContainerInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGroupedContainerInstancesWithChan invokes the sas.DescribeGroupedContainerInstances API asynchronously
func (client *Client) DescribeGroupedContainerInstancesWithChan(request *DescribeGroupedContainerInstancesRequest) (<-chan *DescribeGroupedContainerInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeGroupedContainerInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGroupedContainerInstances(request)
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

// DescribeGroupedContainerInstancesWithCallback invokes the sas.DescribeGroupedContainerInstances API asynchronously
func (client *Client) DescribeGroupedContainerInstancesWithCallback(request *DescribeGroupedContainerInstancesRequest, callback func(response *DescribeGroupedContainerInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGroupedContainerInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGroupedContainerInstances(request)
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

// DescribeGroupedContainerInstancesRequest is the request struct for api DescribeGroupedContainerInstances
type DescribeGroupedContainerInstancesRequest struct {
	*requests.RpcRequest
	Criteria    string           `position:"Query" name:"Criteria"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	GroupField  string           `position:"Query" name:"GroupField"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	LogicalExp  string           `position:"Query" name:"LogicalExp"`
	FieldValue  string           `position:"Query" name:"FieldValue"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
}

// DescribeGroupedContainerInstancesResponse is the response struct for api DescribeGroupedContainerInstances
type DescribeGroupedContainerInstancesResponse struct {
	*responses.BaseResponse
	RequestId                    string                     `json:"RequestId" xml:"RequestId"`
	PageInfo                     PageInfo                   `json:"PageInfo" xml:"PageInfo"`
	GroupedContainerInstanceList []GroupedContainerInstance `json:"GroupedContainerInstanceList" xml:"GroupedContainerInstanceList"`
}

// CreateDescribeGroupedContainerInstancesRequest creates a request to invoke DescribeGroupedContainerInstances API
func CreateDescribeGroupedContainerInstancesRequest() (request *DescribeGroupedContainerInstancesRequest) {
	request = &DescribeGroupedContainerInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeGroupedContainerInstances", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGroupedContainerInstancesResponse creates a response to parse from DescribeGroupedContainerInstances response
func CreateDescribeGroupedContainerInstancesResponse() (response *DescribeGroupedContainerInstancesResponse) {
	response = &DescribeGroupedContainerInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}