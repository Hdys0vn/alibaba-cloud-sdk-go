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

// ModifyVpcFirewallCenSwitchStatus invokes the cloudfw.ModifyVpcFirewallCenSwitchStatus API synchronously
func (client *Client) ModifyVpcFirewallCenSwitchStatus(request *ModifyVpcFirewallCenSwitchStatusRequest) (response *ModifyVpcFirewallCenSwitchStatusResponse, err error) {
	response = CreateModifyVpcFirewallCenSwitchStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpcFirewallCenSwitchStatusWithChan invokes the cloudfw.ModifyVpcFirewallCenSwitchStatus API asynchronously
func (client *Client) ModifyVpcFirewallCenSwitchStatusWithChan(request *ModifyVpcFirewallCenSwitchStatusRequest) (<-chan *ModifyVpcFirewallCenSwitchStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyVpcFirewallCenSwitchStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpcFirewallCenSwitchStatus(request)
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

// ModifyVpcFirewallCenSwitchStatusWithCallback invokes the cloudfw.ModifyVpcFirewallCenSwitchStatus API asynchronously
func (client *Client) ModifyVpcFirewallCenSwitchStatusWithCallback(request *ModifyVpcFirewallCenSwitchStatusRequest, callback func(response *ModifyVpcFirewallCenSwitchStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpcFirewallCenSwitchStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpcFirewallCenSwitchStatus(request)
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

// ModifyVpcFirewallCenSwitchStatusRequest is the request struct for api ModifyVpcFirewallCenSwitchStatus
type ModifyVpcFirewallCenSwitchStatusRequest struct {
	*requests.RpcRequest
	FirewallSwitch string `position:"Query" name:"FirewallSwitch"`
	SourceIp       string `position:"Query" name:"SourceIp"`
	MemberUid      string `position:"Query" name:"MemberUid"`
	Lang           string `position:"Query" name:"Lang"`
	VpcFirewallId  string `position:"Query" name:"VpcFirewallId"`
}

// ModifyVpcFirewallCenSwitchStatusResponse is the response struct for api ModifyVpcFirewallCenSwitchStatus
type ModifyVpcFirewallCenSwitchStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVpcFirewallCenSwitchStatusRequest creates a request to invoke ModifyVpcFirewallCenSwitchStatus API
func CreateModifyVpcFirewallCenSwitchStatusRequest() (request *ModifyVpcFirewallCenSwitchStatusRequest) {
	request = &ModifyVpcFirewallCenSwitchStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "ModifyVpcFirewallCenSwitchStatus", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpcFirewallCenSwitchStatusResponse creates a response to parse from ModifyVpcFirewallCenSwitchStatus response
func CreateModifyVpcFirewallCenSwitchStatusResponse() (response *ModifyVpcFirewallCenSwitchStatusResponse) {
	response = &ModifyVpcFirewallCenSwitchStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
