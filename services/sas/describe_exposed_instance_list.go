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

// DescribeExposedInstanceList invokes the sas.DescribeExposedInstanceList API synchronously
func (client *Client) DescribeExposedInstanceList(request *DescribeExposedInstanceListRequest) (response *DescribeExposedInstanceListResponse, err error) {
	response = CreateDescribeExposedInstanceListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeExposedInstanceListWithChan invokes the sas.DescribeExposedInstanceList API asynchronously
func (client *Client) DescribeExposedInstanceListWithChan(request *DescribeExposedInstanceListRequest) (<-chan *DescribeExposedInstanceListResponse, <-chan error) {
	responseChan := make(chan *DescribeExposedInstanceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExposedInstanceList(request)
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

// DescribeExposedInstanceListWithCallback invokes the sas.DescribeExposedInstanceList API asynchronously
func (client *Client) DescribeExposedInstanceListWithCallback(request *DescribeExposedInstanceListRequest, callback func(response *DescribeExposedInstanceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExposedInstanceListResponse
		var err error
		defer close(result)
		response, err = client.DescribeExposedInstanceList(request)
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

// DescribeExposedInstanceListRequest is the request struct for api DescribeExposedInstanceList
type DescribeExposedInstanceListRequest struct {
	*requests.RpcRequest
	ExposurePort      string           `position:"Query" name:"ExposurePort"`
	SourceIp          string           `position:"Query" name:"SourceIp"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	VulStatus         requests.Boolean `position:"Query" name:"VulStatus"`
	ExposureIp        string           `position:"Query" name:"ExposureIp"`
	GroupId           requests.Integer `position:"Query" name:"GroupId"`
	CurrentPage       requests.Integer `position:"Query" name:"CurrentPage"`
	ExposureComponent string           `position:"Query" name:"ExposureComponent"`
	InstanceId        string           `position:"Query" name:"InstanceId"`
	InstanceName      string           `position:"Query" name:"InstanceName"`
	HealthStatus      requests.Boolean `position:"Query" name:"HealthStatus"`
}

// DescribeExposedInstanceListResponse is the response struct for api DescribeExposedInstanceList
type DescribeExposedInstanceListResponse struct {
	*responses.BaseResponse
	RequestId        string            `json:"RequestId" xml:"RequestId"`
	PageInfo         PageInfo          `json:"PageInfo" xml:"PageInfo"`
	ExposedInstances []ExposedInstance `json:"ExposedInstances" xml:"ExposedInstances"`
}

// CreateDescribeExposedInstanceListRequest creates a request to invoke DescribeExposedInstanceList API
func CreateDescribeExposedInstanceListRequest() (request *DescribeExposedInstanceListRequest) {
	request = &DescribeExposedInstanceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeExposedInstanceList", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeExposedInstanceListResponse creates a response to parse from DescribeExposedInstanceList response
func CreateDescribeExposedInstanceListResponse() (response *DescribeExposedInstanceListResponse) {
	response = &DescribeExposedInstanceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
