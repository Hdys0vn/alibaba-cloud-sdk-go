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

// ModifyVpcFirewallControlPolicy invokes the cloudfw.ModifyVpcFirewallControlPolicy API synchronously
func (client *Client) ModifyVpcFirewallControlPolicy(request *ModifyVpcFirewallControlPolicyRequest) (response *ModifyVpcFirewallControlPolicyResponse, err error) {
	response = CreateModifyVpcFirewallControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpcFirewallControlPolicyWithChan invokes the cloudfw.ModifyVpcFirewallControlPolicy API asynchronously
func (client *Client) ModifyVpcFirewallControlPolicyWithChan(request *ModifyVpcFirewallControlPolicyRequest) (<-chan *ModifyVpcFirewallControlPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyVpcFirewallControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpcFirewallControlPolicy(request)
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

// ModifyVpcFirewallControlPolicyWithCallback invokes the cloudfw.ModifyVpcFirewallControlPolicy API asynchronously
func (client *Client) ModifyVpcFirewallControlPolicyWithCallback(request *ModifyVpcFirewallControlPolicyRequest, callback func(response *ModifyVpcFirewallControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpcFirewallControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpcFirewallControlPolicy(request)
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

// ModifyVpcFirewallControlPolicyRequest is the request struct for api ModifyVpcFirewallControlPolicy
type ModifyVpcFirewallControlPolicyRequest struct {
	*requests.RpcRequest
	DestPortType        string    `position:"Query" name:"DestPortType"`
	Release             string    `position:"Query" name:"Release"`
	Destination         string    `position:"Query" name:"Destination"`
	DestinationType     string    `position:"Query" name:"DestinationType"`
	DestPortGroup       string    `position:"Query" name:"DestPortGroup"`
	ApplicationNameList *[]string `position:"Query" name:"ApplicationNameList"  type:"Repeated"`
	Description         string    `position:"Query" name:"Description"`
	Source              string    `position:"Query" name:"Source"`
	AclUuid             string    `position:"Query" name:"AclUuid"`
	AclAction           string    `position:"Query" name:"AclAction"`
	SourceIp            string    `position:"Query" name:"SourceIp"`
	SourceType          string    `position:"Query" name:"SourceType"`
	Lang                string    `position:"Query" name:"Lang"`
	VpcFirewallId       string    `position:"Query" name:"VpcFirewallId"`
	ApplicationName     string    `position:"Query" name:"ApplicationName"`
	Proto               string    `position:"Query" name:"Proto"`
	DestPort            string    `position:"Query" name:"DestPort"`
}

// ModifyVpcFirewallControlPolicyResponse is the response struct for api ModifyVpcFirewallControlPolicy
type ModifyVpcFirewallControlPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVpcFirewallControlPolicyRequest creates a request to invoke ModifyVpcFirewallControlPolicy API
func CreateModifyVpcFirewallControlPolicyRequest() (request *ModifyVpcFirewallControlPolicyRequest) {
	request = &ModifyVpcFirewallControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "ModifyVpcFirewallControlPolicy", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpcFirewallControlPolicyResponse creates a response to parse from ModifyVpcFirewallControlPolicy response
func CreateModifyVpcFirewallControlPolicyResponse() (response *ModifyVpcFirewallControlPolicyResponse) {
	response = &ModifyVpcFirewallControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}