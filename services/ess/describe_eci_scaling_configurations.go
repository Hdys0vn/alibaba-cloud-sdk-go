package ess

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

// DescribeEciScalingConfigurations invokes the ess.DescribeEciScalingConfigurations API synchronously
func (client *Client) DescribeEciScalingConfigurations(request *DescribeEciScalingConfigurationsRequest) (response *DescribeEciScalingConfigurationsResponse, err error) {
	response = CreateDescribeEciScalingConfigurationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEciScalingConfigurationsWithChan invokes the ess.DescribeEciScalingConfigurations API asynchronously
func (client *Client) DescribeEciScalingConfigurationsWithChan(request *DescribeEciScalingConfigurationsRequest) (<-chan *DescribeEciScalingConfigurationsResponse, <-chan error) {
	responseChan := make(chan *DescribeEciScalingConfigurationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEciScalingConfigurations(request)
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

// DescribeEciScalingConfigurationsWithCallback invokes the ess.DescribeEciScalingConfigurations API asynchronously
func (client *Client) DescribeEciScalingConfigurationsWithCallback(request *DescribeEciScalingConfigurationsRequest, callback func(response *DescribeEciScalingConfigurationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEciScalingConfigurationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeEciScalingConfigurations(request)
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

// DescribeEciScalingConfigurationsRequest is the request struct for api DescribeEciScalingConfigurations
type DescribeEciScalingConfigurationsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ScalingGroupId           string           `position:"Query" name:"ScalingGroupId"`
	PageNumber               requests.Integer `position:"Query" name:"PageNumber"`
	PageSize                 requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	ScalingConfigurationName *[]string        `position:"Query" name:"ScalingConfigurationName"  type:"Repeated"`
	ScalingConfigurationId   *[]string        `position:"Query" name:"ScalingConfigurationId"  type:"Repeated"`
}

// DescribeEciScalingConfigurationsResponse is the response struct for api DescribeEciScalingConfigurations
type DescribeEciScalingConfigurationsResponse struct {
	*responses.BaseResponse
	PageNumber            int                    `json:"PageNumber" xml:"PageNumber"`
	PageSize              int                    `json:"PageSize" xml:"PageSize"`
	RequestId             string                 `json:"RequestId" xml:"RequestId"`
	TotalCount            int                    `json:"TotalCount" xml:"TotalCount"`
	ScalingConfigurations []ScalingConfiguration `json:"ScalingConfigurations" xml:"ScalingConfigurations"`
}

// CreateDescribeEciScalingConfigurationsRequest creates a request to invoke DescribeEciScalingConfigurations API
func CreateDescribeEciScalingConfigurationsRequest() (request *DescribeEciScalingConfigurationsRequest) {
	request = &DescribeEciScalingConfigurationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeEciScalingConfigurations", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEciScalingConfigurationsResponse creates a response to parse from DescribeEciScalingConfigurations response
func CreateDescribeEciScalingConfigurationsResponse() (response *DescribeEciScalingConfigurationsResponse) {
	response = &DescribeEciScalingConfigurationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
