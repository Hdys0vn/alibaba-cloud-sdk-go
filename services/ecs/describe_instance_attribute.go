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

// DescribeInstanceAttribute invokes the ecs.DescribeInstanceAttribute API synchronously
func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (response *DescribeInstanceAttributeResponse, err error) {
	response = CreateDescribeInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceAttributeWithChan invokes the ecs.DescribeInstanceAttribute API asynchronously
func (client *Client) DescribeInstanceAttributeWithChan(request *DescribeInstanceAttributeRequest) (<-chan *DescribeInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceAttribute(request)
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

// DescribeInstanceAttributeWithCallback invokes the ecs.DescribeInstanceAttribute API asynchronously
func (client *Client) DescribeInstanceAttributeWithCallback(request *DescribeInstanceAttributeRequest, callback func(response *DescribeInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceAttribute(request)
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

// DescribeInstanceAttributeRequest is the request struct for api DescribeInstanceAttribute
type DescribeInstanceAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeInstanceAttributeResponse is the response struct for api DescribeInstanceAttribute
type DescribeInstanceAttributeResponse struct {
	*responses.BaseResponse
	Status                  string                                      `json:"Status" xml:"Status"`
	SerialNumber            string                                      `json:"SerialNumber" xml:"SerialNumber"`
	CreationTime            string                                      `json:"CreationTime" xml:"CreationTime"`
	RequestId               string                                      `json:"RequestId" xml:"RequestId"`
	Description             string                                      `json:"Description" xml:"Description"`
	InstanceName            string                                      `json:"InstanceName" xml:"InstanceName"`
	InstanceNetworkType     string                                      `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	Memory                  int                                         `json:"Memory" xml:"Memory"`
	ImageId                 string                                      `json:"ImageId" xml:"ImageId"`
	ClusterId               string                                      `json:"ClusterId" xml:"ClusterId"`
	VlanId                  string                                      `json:"VlanId" xml:"VlanId"`
	StoppedMode             string                                      `json:"StoppedMode" xml:"StoppedMode"`
	HostName                string                                      `json:"HostName" xml:"HostName"`
	InstanceId              string                                      `json:"InstanceId" xml:"InstanceId"`
	InstanceType            string                                      `json:"InstanceType" xml:"InstanceType"`
	InternetMaxBandwidthIn  int                                         `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut int                                         `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	RegionId                string                                      `json:"RegionId" xml:"RegionId"`
	InstanceChargeType      string                                      `json:"InstanceChargeType" xml:"InstanceChargeType"`
	IoOptimized             string                                      `json:"IoOptimized" xml:"IoOptimized"`
	Cpu                     int                                         `json:"Cpu" xml:"Cpu"`
	ExpiredTime             string                                      `json:"ExpiredTime" xml:"ExpiredTime"`
	ZoneId                  string                                      `json:"ZoneId" xml:"ZoneId"`
	InternetChargeType      string                                      `json:"InternetChargeType" xml:"InternetChargeType"`
	CreditSpecification     string                                      `json:"CreditSpecification" xml:"CreditSpecification"`
	EnableJumboFrame        bool                                        `json:"EnableJumboFrame" xml:"EnableJumboFrame"`
	SecurityGroupIds        SecurityGroupIdsInDescribeInstanceAttribute `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	PublicIpAddress         PublicIpAddressInDescribeInstanceAttribute  `json:"PublicIpAddress" xml:"PublicIpAddress"`
	InnerIpAddress          InnerIpAddressInDescribeInstanceAttribute   `json:"InnerIpAddress" xml:"InnerIpAddress"`
	VpcAttributes           VpcAttributes                               `json:"VpcAttributes" xml:"VpcAttributes"`
	EipAddress              EipAddressInDescribeInstanceAttribute       `json:"EipAddress" xml:"EipAddress"`
	DedicatedHostAttribute  DedicatedHostAttribute                      `json:"DedicatedHostAttribute" xml:"DedicatedHostAttribute"`
	OperationLocks          OperationLocksInDescribeInstanceAttribute   `json:"OperationLocks" xml:"OperationLocks"`
}

// CreateDescribeInstanceAttributeRequest creates a request to invoke DescribeInstanceAttribute API
func CreateDescribeInstanceAttributeRequest() (request *DescribeInstanceAttributeRequest) {
	request = &DescribeInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceAttribute", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceAttributeResponse creates a response to parse from DescribeInstanceAttribute response
func CreateDescribeInstanceAttributeResponse() (response *DescribeInstanceAttributeResponse) {
	response = &DescribeInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
