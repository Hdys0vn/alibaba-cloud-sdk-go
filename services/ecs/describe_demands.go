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

// DescribeDemands invokes the ecs.DescribeDemands API synchronously
func (client *Client) DescribeDemands(request *DescribeDemandsRequest) (response *DescribeDemandsResponse, err error) {
	response = CreateDescribeDemandsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDemandsWithChan invokes the ecs.DescribeDemands API asynchronously
func (client *Client) DescribeDemandsWithChan(request *DescribeDemandsRequest) (<-chan *DescribeDemandsResponse, <-chan error) {
	responseChan := make(chan *DescribeDemandsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDemands(request)
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

// DescribeDemandsWithCallback invokes the ecs.DescribeDemands API asynchronously
func (client *Client) DescribeDemandsWithCallback(request *DescribeDemandsRequest, callback func(response *DescribeDemandsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDemandsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDemands(request)
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

// DescribeDemandsRequest is the request struct for api DescribeDemands
type DescribeDemandsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer      `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer      `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer      `position:"Query" name:"PageSize"`
	InstanceType         string                `position:"Query" name:"InstanceType"`
	Tag                  *[]DescribeDemandsTag `position:"Query" name:"Tag"  type:"Repeated"`
	InstanceChargeType   string                `position:"Query" name:"InstanceChargeType"`
	DryRun               requests.Boolean      `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string                `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                `position:"Query" name:"OwnerAccount"`
	InstanceTypeFamily   string                `position:"Query" name:"InstanceTypeFamily"`
	OwnerId              requests.Integer      `position:"Query" name:"OwnerId"`
	DemandStatus         *[]string             `position:"Query" name:"DemandStatus"  type:"Repeated"`
	DemandId             string                `position:"Query" name:"DemandId"`
	ZoneId               string                `position:"Query" name:"ZoneId"`
	DemandType           string                `position:"Query" name:"DemandType"`
}

// DescribeDemandsTag is a repeated param struct in DescribeDemandsRequest
type DescribeDemandsTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeDemandsResponse is the response struct for api DescribeDemands
type DescribeDemandsResponse struct {
	*responses.BaseResponse
	PageSize   int     `json:"PageSize" xml:"PageSize"`
	PageNumber int     `json:"PageNumber" xml:"PageNumber"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	RegionId   string  `json:"RegionId" xml:"RegionId"`
	Demands    Demands `json:"Demands" xml:"Demands"`
}

// CreateDescribeDemandsRequest creates a request to invoke DescribeDemands API
func CreateDescribeDemandsRequest() (request *DescribeDemandsRequest) {
	request = &DescribeDemandsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDemands", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDemandsResponse creates a response to parse from DescribeDemands response
func CreateDescribeDemandsResponse() (response *DescribeDemandsResponse) {
	response = &DescribeDemandsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
