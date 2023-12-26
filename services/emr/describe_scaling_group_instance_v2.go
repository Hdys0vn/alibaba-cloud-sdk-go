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

// DescribeScalingGroupInstanceV2 invokes the emr.DescribeScalingGroupInstanceV2 API synchronously
func (client *Client) DescribeScalingGroupInstanceV2(request *DescribeScalingGroupInstanceV2Request) (response *DescribeScalingGroupInstanceV2Response, err error) {
	response = CreateDescribeScalingGroupInstanceV2Response()
	err = client.DoAction(request, response)
	return
}

// DescribeScalingGroupInstanceV2WithChan invokes the emr.DescribeScalingGroupInstanceV2 API asynchronously
func (client *Client) DescribeScalingGroupInstanceV2WithChan(request *DescribeScalingGroupInstanceV2Request) (<-chan *DescribeScalingGroupInstanceV2Response, <-chan error) {
	responseChan := make(chan *DescribeScalingGroupInstanceV2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingGroupInstanceV2(request)
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

// DescribeScalingGroupInstanceV2WithCallback invokes the emr.DescribeScalingGroupInstanceV2 API asynchronously
func (client *Client) DescribeScalingGroupInstanceV2WithCallback(request *DescribeScalingGroupInstanceV2Request, callback func(response *DescribeScalingGroupInstanceV2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingGroupInstanceV2Response
		var err error
		defer close(result)
		response, err = client.DescribeScalingGroupInstanceV2(request)
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

// DescribeScalingGroupInstanceV2Request is the request struct for api DescribeScalingGroupInstanceV2
type DescribeScalingGroupInstanceV2Request struct {
	*requests.RpcRequest
	ResourceOwnerId   requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ScalingGroupBizId string           `position:"Query" name:"ScalingGroupBizId"`
	ResourceGroupId   string           `position:"Query" name:"ResourceGroupId"`
	HostGroupBizId    string           `position:"Query" name:"HostGroupBizId"`
}

// DescribeScalingGroupInstanceV2Response is the response struct for api DescribeScalingGroupInstanceV2
type DescribeScalingGroupInstanceV2Response struct {
	*responses.BaseResponse
	RequestId                 string          `json:"RequestId" xml:"RequestId"`
	HostGroupId               string          `json:"HostGroupId" xml:"HostGroupId"`
	ScalingGroupId            string          `json:"ScalingGroupId" xml:"ScalingGroupId"`
	MinSize                   int             `json:"MinSize" xml:"MinSize"`
	MaxSize                   int             `json:"MaxSize" xml:"MaxSize"`
	DefaultCooldown           int             `json:"DefaultCooldown" xml:"DefaultCooldown"`
	ActiveRuleCategory        string          `json:"ActiveRuleCategory" xml:"ActiveRuleCategory"`
	WithGrace                 bool            `json:"WithGrace" xml:"WithGrace"`
	TimeoutWithGrace          int64           `json:"TimeoutWithGrace" xml:"TimeoutWithGrace"`
	MultiAvailablePolicy      string          `json:"MultiAvailablePolicy" xml:"MultiAvailablePolicy"`
	MultiAvailablePolicyParam string          `json:"MultiAvailablePolicyParam" xml:"MultiAvailablePolicyParam"`
	ScalingConfig             ScalingConfig   `json:"ScalingConfig" xml:"ScalingConfig"`
	ScalingRuleList           ScalingRuleList `json:"ScalingRuleList" xml:"ScalingRuleList"`
}

// CreateDescribeScalingGroupInstanceV2Request creates a request to invoke DescribeScalingGroupInstanceV2 API
func CreateDescribeScalingGroupInstanceV2Request() (request *DescribeScalingGroupInstanceV2Request) {
	request = &DescribeScalingGroupInstanceV2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeScalingGroupInstanceV2", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeScalingGroupInstanceV2Response creates a response to parse from DescribeScalingGroupInstanceV2 response
func CreateDescribeScalingGroupInstanceV2Response() (response *DescribeScalingGroupInstanceV2Response) {
	response = &DescribeScalingGroupInstanceV2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
