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

// DeleteVpcFirewallControlPolicy invokes the cloudfw.DeleteVpcFirewallControlPolicy API synchronously
func (client *Client) DeleteVpcFirewallControlPolicy(request *DeleteVpcFirewallControlPolicyRequest) (response *DeleteVpcFirewallControlPolicyResponse, err error) {
	response = CreateDeleteVpcFirewallControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVpcFirewallControlPolicyWithChan invokes the cloudfw.DeleteVpcFirewallControlPolicy API asynchronously
func (client *Client) DeleteVpcFirewallControlPolicyWithChan(request *DeleteVpcFirewallControlPolicyRequest) (<-chan *DeleteVpcFirewallControlPolicyResponse, <-chan error) {
	responseChan := make(chan *DeleteVpcFirewallControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVpcFirewallControlPolicy(request)
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

// DeleteVpcFirewallControlPolicyWithCallback invokes the cloudfw.DeleteVpcFirewallControlPolicy API asynchronously
func (client *Client) DeleteVpcFirewallControlPolicyWithCallback(request *DeleteVpcFirewallControlPolicyRequest, callback func(response *DeleteVpcFirewallControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVpcFirewallControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.DeleteVpcFirewallControlPolicy(request)
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

// DeleteVpcFirewallControlPolicyRequest is the request struct for api DeleteVpcFirewallControlPolicy
type DeleteVpcFirewallControlPolicyRequest struct {
	*requests.RpcRequest
	AclUuid       string `position:"Query" name:"AclUuid"`
	SourceIp      string `position:"Query" name:"SourceIp"`
	Lang          string `position:"Query" name:"Lang"`
	VpcFirewallId string `position:"Query" name:"VpcFirewallId"`
}

// DeleteVpcFirewallControlPolicyResponse is the response struct for api DeleteVpcFirewallControlPolicy
type DeleteVpcFirewallControlPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVpcFirewallControlPolicyRequest creates a request to invoke DeleteVpcFirewallControlPolicy API
func CreateDeleteVpcFirewallControlPolicyRequest() (request *DeleteVpcFirewallControlPolicyRequest) {
	request = &DeleteVpcFirewallControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DeleteVpcFirewallControlPolicy", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVpcFirewallControlPolicyResponse creates a response to parse from DeleteVpcFirewallControlPolicy response
func CreateDeleteVpcFirewallControlPolicyResponse() (response *DeleteVpcFirewallControlPolicyResponse) {
	response = &DeleteVpcFirewallControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
