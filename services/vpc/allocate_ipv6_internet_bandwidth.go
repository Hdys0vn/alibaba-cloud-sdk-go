package vpc

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

// AllocateIpv6InternetBandwidth invokes the vpc.AllocateIpv6InternetBandwidth API synchronously
func (client *Client) AllocateIpv6InternetBandwidth(request *AllocateIpv6InternetBandwidthRequest) (response *AllocateIpv6InternetBandwidthResponse, err error) {
	response = CreateAllocateIpv6InternetBandwidthResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateIpv6InternetBandwidthWithChan invokes the vpc.AllocateIpv6InternetBandwidth API asynchronously
func (client *Client) AllocateIpv6InternetBandwidthWithChan(request *AllocateIpv6InternetBandwidthRequest) (<-chan *AllocateIpv6InternetBandwidthResponse, <-chan error) {
	responseChan := make(chan *AllocateIpv6InternetBandwidthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateIpv6InternetBandwidth(request)
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

// AllocateIpv6InternetBandwidthWithCallback invokes the vpc.AllocateIpv6InternetBandwidth API asynchronously
func (client *Client) AllocateIpv6InternetBandwidthWithCallback(request *AllocateIpv6InternetBandwidthRequest, callback func(response *AllocateIpv6InternetBandwidthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateIpv6InternetBandwidthResponse
		var err error
		defer close(result)
		response, err = client.AllocateIpv6InternetBandwidth(request)
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

// AllocateIpv6InternetBandwidthRequest is the request struct for api AllocateIpv6InternetBandwidth
type AllocateIpv6InternetBandwidthRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Bandwidth            requests.Integer `position:"Query" name:"Bandwidth"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Ipv6AddressId        string           `position:"Query" name:"Ipv6AddressId"`
	InternetChargeType   string           `position:"Query" name:"InternetChargeType"`
	Ipv6GatewayId        string           `position:"Query" name:"Ipv6GatewayId"`
}

// AllocateIpv6InternetBandwidthResponse is the response struct for api AllocateIpv6InternetBandwidth
type AllocateIpv6InternetBandwidthResponse struct {
	*responses.BaseResponse
	Ipv6AddressId       string `json:"Ipv6AddressId" xml:"Ipv6AddressId"`
	RequestId           string `json:"RequestId" xml:"RequestId"`
	InternetBandwidthId string `json:"InternetBandwidthId" xml:"InternetBandwidthId"`
}

// CreateAllocateIpv6InternetBandwidthRequest creates a request to invoke AllocateIpv6InternetBandwidth API
func CreateAllocateIpv6InternetBandwidthRequest() (request *AllocateIpv6InternetBandwidthRequest) {
	request = &AllocateIpv6InternetBandwidthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AllocateIpv6InternetBandwidth", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAllocateIpv6InternetBandwidthResponse creates a response to parse from AllocateIpv6InternetBandwidth response
func CreateAllocateIpv6InternetBandwidthResponse() (response *AllocateIpv6InternetBandwidthResponse) {
	response = &AllocateIpv6InternetBandwidthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
