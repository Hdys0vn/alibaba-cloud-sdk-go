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

// DescribeInstanceTypes invokes the ecs.DescribeInstanceTypes API synchronously
func (client *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
	response = CreateDescribeInstanceTypesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceTypesWithChan invokes the ecs.DescribeInstanceTypes API asynchronously
func (client *Client) DescribeInstanceTypesWithChan(request *DescribeInstanceTypesRequest) (<-chan *DescribeInstanceTypesResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceTypes(request)
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

// DescribeInstanceTypesWithCallback invokes the ecs.DescribeInstanceTypes API asynchronously
func (client *Client) DescribeInstanceTypesWithCallback(request *DescribeInstanceTypesRequest, callback func(response *DescribeInstanceTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceTypesResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceTypes(request)
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

// DescribeInstanceTypesRequest is the request struct for api DescribeInstanceTypes
type DescribeInstanceTypesRequest struct {
	*requests.RpcRequest
	GPUSpec                            string           `position:"Query" name:"GPUSpec"`
	ResourceOwnerId                    requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MaximumCpuCoreCount                requests.Integer `position:"Query" name:"MaximumCpuCoreCount"`
	MaximumGPUAmount                   requests.Integer `position:"Query" name:"MaximumGPUAmount"`
	LocalStorageCategory               string           `position:"Query" name:"LocalStorageCategory"`
	MaximumMemorySize                  requests.Float   `position:"Query" name:"MaximumMemorySize"`
	InstanceCategory                   string           `position:"Query" name:"InstanceCategory"`
	MinimumInstancePpsTx               requests.Integer `position:"Query" name:"MinimumInstancePpsTx"`
	MinimumCpuCoreCount                requests.Integer `position:"Query" name:"MinimumCpuCoreCount"`
	MinimumPrimaryEniQueueNumber       requests.Integer `position:"Query" name:"MinimumPrimaryEniQueueNumber"`
	MinimumBaselineCredit              requests.Integer `position:"Query" name:"MinimumBaselineCredit"`
	MinimumSecondaryEniQueueNumber     requests.Integer `position:"Query" name:"MinimumSecondaryEniQueueNumber"`
	MinimumInstanceBandwidthTx         requests.Integer `position:"Query" name:"MinimumInstanceBandwidthTx"`
	MinimumGPUAmount                   requests.Integer `position:"Query" name:"MinimumGPUAmount"`
	MaximumCpuSpeedFrequency           requests.Float   `position:"Query" name:"MaximumCpuSpeedFrequency"`
	CpuArchitecture                    string           `position:"Query" name:"CpuArchitecture"`
	OwnerId                            requests.Integer `position:"Query" name:"OwnerId"`
	MinimumMemorySize                  requests.Float   `position:"Query" name:"MinimumMemorySize"`
	MinimumEniQuantity                 requests.Integer `position:"Query" name:"MinimumEniQuantity"`
	InstanceFamilyLevel                string           `position:"Query" name:"InstanceFamilyLevel"`
	MinimumQueuePairNumber             requests.Integer `position:"Query" name:"MinimumQueuePairNumber"`
	MinimumLocalStorageAmount          requests.Integer `position:"Query" name:"MinimumLocalStorageAmount"`
	MaxResults                         requests.Integer `position:"Query" name:"MaxResults"`
	PhysicalProcessorModel             string           `position:"Query" name:"PhysicalProcessorModel"`
	MaximumCpuTurboFrequency           requests.Float   `position:"Query" name:"MaximumCpuTurboFrequency"`
	InstanceTypes                      *[]string        `position:"Query" name:"InstanceTypes"  type:"Repeated"`
	MinimumInstancePpsRx               requests.Integer `position:"Query" name:"MinimumInstancePpsRx"`
	MinimumEniIpv6AddressQuantity      requests.Integer `position:"Query" name:"MinimumEniIpv6AddressQuantity"`
	MinimumEriQuantity                 requests.Integer `position:"Query" name:"MinimumEriQuantity"`
	MinimumDiskQuantity                requests.Integer `position:"Query" name:"MinimumDiskQuantity"`
	MinimumCpuTurboFrequency           requests.Float   `position:"Query" name:"MinimumCpuTurboFrequency"`
	NextToken                          string           `position:"Query" name:"NextToken"`
	MinimumInstanceBandwidthRx         requests.Integer `position:"Query" name:"MinimumInstanceBandwidthRx"`
	MinimumCpuSpeedFrequency           requests.Float   `position:"Query" name:"MinimumCpuSpeedFrequency"`
	NvmeSupport                        string           `position:"Query" name:"NvmeSupport"`
	ResourceOwnerAccount               string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                       string           `position:"Query" name:"OwnerAccount"`
	MinimumInitialCredit               requests.Integer `position:"Query" name:"MinimumInitialCredit"`
	InstanceTypeFamily                 string           `position:"Query" name:"InstanceTypeFamily"`
	MinimumEniPrivateIpAddressQuantity requests.Integer `position:"Query" name:"MinimumEniPrivateIpAddressQuantity"`
	MinimumLocalStorageCapacity        requests.Integer `position:"Query" name:"MinimumLocalStorageCapacity"`
}

// DescribeInstanceTypesResponse is the response struct for api DescribeInstanceTypes
type DescribeInstanceTypesResponse struct {
	*responses.BaseResponse
	RequestId     string                               `json:"RequestId" xml:"RequestId"`
	NextToken     string                               `json:"NextToken" xml:"NextToken"`
	InstanceTypes InstanceTypesInDescribeInstanceTypes `json:"InstanceTypes" xml:"InstanceTypes"`
}

// CreateDescribeInstanceTypesRequest creates a request to invoke DescribeInstanceTypes API
func CreateDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
	request = &DescribeInstanceTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceTypes", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceTypesResponse creates a response to parse from DescribeInstanceTypes response
func CreateDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
	response = &DescribeInstanceTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
