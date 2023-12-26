package antiddos_public

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

// DescribeInstanceIpAddress invokes the antiddos_public.DescribeInstanceIpAddress API synchronously
func (client *Client) DescribeInstanceIpAddress(request *DescribeInstanceIpAddressRequest) (response *DescribeInstanceIpAddressResponse, err error) {
	response = CreateDescribeInstanceIpAddressResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceIpAddressWithChan invokes the antiddos_public.DescribeInstanceIpAddress API asynchronously
func (client *Client) DescribeInstanceIpAddressWithChan(request *DescribeInstanceIpAddressRequest) (<-chan *DescribeInstanceIpAddressResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceIpAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceIpAddress(request)
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

// DescribeInstanceIpAddressWithCallback invokes the antiddos_public.DescribeInstanceIpAddress API asynchronously
func (client *Client) DescribeInstanceIpAddressWithCallback(request *DescribeInstanceIpAddressRequest, callback func(response *DescribeInstanceIpAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceIpAddressResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceIpAddress(request)
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

// DescribeInstanceIpAddressRequest is the request struct for api DescribeInstanceIpAddress
type DescribeInstanceIpAddressRequest struct {
	*requests.RpcRequest
	EagleEyeRpcId    string           `position:"Body" name:"EagleEyeRpcId"`
	SourceIp         string           `position:"Query" name:"SourceIp"`
	EagleEyeTraceId  string           `position:"Body" name:"EagleEyeTraceId"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	DdosRegionId     string           `position:"Query" name:"DdosRegionId"`
	InstanceType     string           `position:"Query" name:"InstanceType"`
	Lang             string           `position:"Query" name:"Lang"`
	DdosStatus       string           `position:"Query" name:"DdosStatus"`
	CurrentPage      requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceName     string           `position:"Query" name:"InstanceName"`
	InstanceId       string           `position:"Query" name:"InstanceId"`
	EagleEyeUserData string           `position:"Body" name:"EagleEyeUserData"`
	InstanceIp       string           `position:"Query" name:"InstanceIp"`
}

// DescribeInstanceIpAddressResponse is the response struct for api DescribeInstanceIpAddress
type DescribeInstanceIpAddressResponse struct {
	*responses.BaseResponse
	Total        int        `json:"Total" xml:"Total"`
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	InstanceList []Instance `json:"InstanceList" xml:"InstanceList"`
}

// CreateDescribeInstanceIpAddressRequest creates a request to invoke DescribeInstanceIpAddress API
func CreateDescribeInstanceIpAddressRequest() (request *DescribeInstanceIpAddressRequest) {
	request = &DescribeInstanceIpAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("antiddos-public", "2017-05-18", "DescribeInstanceIpAddress", "ddosbasic", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceIpAddressResponse creates a response to parse from DescribeInstanceIpAddress response
func CreateDescribeInstanceIpAddressResponse() (response *DescribeInstanceIpAddressResponse) {
	response = &DescribeInstanceIpAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}