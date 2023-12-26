package cloudfw

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

// DescribeTrFirewallV2RoutePolicyList invokes the cloudfw.DescribeTrFirewallV2RoutePolicyList API synchronously
func (client *Client) DescribeTrFirewallV2RoutePolicyList(request *DescribeTrFirewallV2RoutePolicyListRequest) (response *DescribeTrFirewallV2RoutePolicyListResponse, err error) {
	response = CreateDescribeTrFirewallV2RoutePolicyListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTrFirewallV2RoutePolicyListWithChan invokes the cloudfw.DescribeTrFirewallV2RoutePolicyList API asynchronously
func (client *Client) DescribeTrFirewallV2RoutePolicyListWithChan(request *DescribeTrFirewallV2RoutePolicyListRequest) (<-chan *DescribeTrFirewallV2RoutePolicyListResponse, <-chan error) {
	responseChan := make(chan *DescribeTrFirewallV2RoutePolicyListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTrFirewallV2RoutePolicyList(request)
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

// DescribeTrFirewallV2RoutePolicyListWithCallback invokes the cloudfw.DescribeTrFirewallV2RoutePolicyList API asynchronously
func (client *Client) DescribeTrFirewallV2RoutePolicyListWithCallback(request *DescribeTrFirewallV2RoutePolicyListRequest, callback func(response *DescribeTrFirewallV2RoutePolicyListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTrFirewallV2RoutePolicyListResponse
		var err error
		defer close(result)
		response, err = client.DescribeTrFirewallV2RoutePolicyList(request)
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

// DescribeTrFirewallV2RoutePolicyListRequest is the request struct for api DescribeTrFirewallV2RoutePolicyList
type DescribeTrFirewallV2RoutePolicyListRequest struct {
	*requests.RpcRequest
	FirewallId  string           `position:"Query" name:"FirewallId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PolicyId    string           `position:"Query" name:"PolicyId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Lang        string           `position:"Query" name:"Lang"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
}

// DescribeTrFirewallV2RoutePolicyListResponse is the response struct for api DescribeTrFirewallV2RoutePolicyList
type DescribeTrFirewallV2RoutePolicyListResponse struct {
	*responses.BaseResponse
	RequestId               string                        `json:"RequestId" xml:"RequestId"`
	TotalCount              string                        `json:"TotalCount" xml:"TotalCount"`
	TrFirewallRoutePolicies []TrFirewallRoutePoliciesItem `json:"TrFirewallRoutePolicies" xml:"TrFirewallRoutePolicies"`
}

// CreateDescribeTrFirewallV2RoutePolicyListRequest creates a request to invoke DescribeTrFirewallV2RoutePolicyList API
func CreateDescribeTrFirewallV2RoutePolicyListRequest() (request *DescribeTrFirewallV2RoutePolicyListRequest) {
	request = &DescribeTrFirewallV2RoutePolicyListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeTrFirewallV2RoutePolicyList", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTrFirewallV2RoutePolicyListResponse creates a response to parse from DescribeTrFirewallV2RoutePolicyList response
func CreateDescribeTrFirewallV2RoutePolicyListResponse() (response *DescribeTrFirewallV2RoutePolicyListResponse) {
	response = &DescribeTrFirewallV2RoutePolicyListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
