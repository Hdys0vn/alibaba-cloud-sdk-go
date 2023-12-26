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

// ModifyIpv6InternetBandwidth invokes the vpc.ModifyIpv6InternetBandwidth API synchronously
func (client *Client) ModifyIpv6InternetBandwidth(request *ModifyIpv6InternetBandwidthRequest) (response *ModifyIpv6InternetBandwidthResponse, err error) {
	response = CreateModifyIpv6InternetBandwidthResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyIpv6InternetBandwidthWithChan invokes the vpc.ModifyIpv6InternetBandwidth API asynchronously
func (client *Client) ModifyIpv6InternetBandwidthWithChan(request *ModifyIpv6InternetBandwidthRequest) (<-chan *ModifyIpv6InternetBandwidthResponse, <-chan error) {
	responseChan := make(chan *ModifyIpv6InternetBandwidthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyIpv6InternetBandwidth(request)
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

// ModifyIpv6InternetBandwidthWithCallback invokes the vpc.ModifyIpv6InternetBandwidth API asynchronously
func (client *Client) ModifyIpv6InternetBandwidthWithCallback(request *ModifyIpv6InternetBandwidthRequest, callback func(response *ModifyIpv6InternetBandwidthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyIpv6InternetBandwidthResponse
		var err error
		defer close(result)
		response, err = client.ModifyIpv6InternetBandwidth(request)
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

// ModifyIpv6InternetBandwidthRequest is the request struct for api ModifyIpv6InternetBandwidth
type ModifyIpv6InternetBandwidthRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken             string           `position:"Query" name:"ClientToken"`
	Ipv6InternetBandwidthId string           `position:"Query" name:"Ipv6InternetBandwidthId"`
	Bandwidth               requests.Integer `position:"Query" name:"Bandwidth"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	Ipv6AddressId           string           `position:"Query" name:"Ipv6AddressId"`
}

// ModifyIpv6InternetBandwidthResponse is the response struct for api ModifyIpv6InternetBandwidth
type ModifyIpv6InternetBandwidthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyIpv6InternetBandwidthRequest creates a request to invoke ModifyIpv6InternetBandwidth API
func CreateModifyIpv6InternetBandwidthRequest() (request *ModifyIpv6InternetBandwidthRequest) {
	request = &ModifyIpv6InternetBandwidthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyIpv6InternetBandwidth", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyIpv6InternetBandwidthResponse creates a response to parse from ModifyIpv6InternetBandwidth response
func CreateModifyIpv6InternetBandwidthResponse() (response *ModifyIpv6InternetBandwidthResponse) {
	response = &ModifyIpv6InternetBandwidthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
