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

// ResetVpcFirewallRuleHitCount invokes the cloudfw.ResetVpcFirewallRuleHitCount API synchronously
func (client *Client) ResetVpcFirewallRuleHitCount(request *ResetVpcFirewallRuleHitCountRequest) (response *ResetVpcFirewallRuleHitCountResponse, err error) {
	response = CreateResetVpcFirewallRuleHitCountResponse()
	err = client.DoAction(request, response)
	return
}

// ResetVpcFirewallRuleHitCountWithChan invokes the cloudfw.ResetVpcFirewallRuleHitCount API asynchronously
func (client *Client) ResetVpcFirewallRuleHitCountWithChan(request *ResetVpcFirewallRuleHitCountRequest) (<-chan *ResetVpcFirewallRuleHitCountResponse, <-chan error) {
	responseChan := make(chan *ResetVpcFirewallRuleHitCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetVpcFirewallRuleHitCount(request)
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

// ResetVpcFirewallRuleHitCountWithCallback invokes the cloudfw.ResetVpcFirewallRuleHitCount API asynchronously
func (client *Client) ResetVpcFirewallRuleHitCountWithCallback(request *ResetVpcFirewallRuleHitCountRequest, callback func(response *ResetVpcFirewallRuleHitCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetVpcFirewallRuleHitCountResponse
		var err error
		defer close(result)
		response, err = client.ResetVpcFirewallRuleHitCount(request)
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

// ResetVpcFirewallRuleHitCountRequest is the request struct for api ResetVpcFirewallRuleHitCount
type ResetVpcFirewallRuleHitCountRequest struct {
	*requests.RpcRequest
	AclUuid  string `position:"Query" name:"AclUuid"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
}

// ResetVpcFirewallRuleHitCountResponse is the response struct for api ResetVpcFirewallRuleHitCount
type ResetVpcFirewallRuleHitCountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetVpcFirewallRuleHitCountRequest creates a request to invoke ResetVpcFirewallRuleHitCount API
func CreateResetVpcFirewallRuleHitCountRequest() (request *ResetVpcFirewallRuleHitCountRequest) {
	request = &ResetVpcFirewallRuleHitCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "ResetVpcFirewallRuleHitCount", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetVpcFirewallRuleHitCountResponse creates a response to parse from ResetVpcFirewallRuleHitCount response
func CreateResetVpcFirewallRuleHitCountResponse() (response *ResetVpcFirewallRuleHitCountResponse) {
	response = &ResetVpcFirewallRuleHitCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
