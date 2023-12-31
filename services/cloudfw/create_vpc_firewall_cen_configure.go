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

// CreateVpcFirewallCenConfigure invokes the cloudfw.CreateVpcFirewallCenConfigure API synchronously
func (client *Client) CreateVpcFirewallCenConfigure(request *CreateVpcFirewallCenConfigureRequest) (response *CreateVpcFirewallCenConfigureResponse, err error) {
	response = CreateCreateVpcFirewallCenConfigureResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVpcFirewallCenConfigureWithChan invokes the cloudfw.CreateVpcFirewallCenConfigure API asynchronously
func (client *Client) CreateVpcFirewallCenConfigureWithChan(request *CreateVpcFirewallCenConfigureRequest) (<-chan *CreateVpcFirewallCenConfigureResponse, <-chan error) {
	responseChan := make(chan *CreateVpcFirewallCenConfigureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVpcFirewallCenConfigure(request)
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

// CreateVpcFirewallCenConfigureWithCallback invokes the cloudfw.CreateVpcFirewallCenConfigure API asynchronously
func (client *Client) CreateVpcFirewallCenConfigureWithCallback(request *CreateVpcFirewallCenConfigureRequest, callback func(response *CreateVpcFirewallCenConfigureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVpcFirewallCenConfigureResponse
		var err error
		defer close(result)
		response, err = client.CreateVpcFirewallCenConfigure(request)
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

// CreateVpcFirewallCenConfigureRequest is the request struct for api CreateVpcFirewallCenConfigure
type CreateVpcFirewallCenConfigureRequest struct {
	*requests.RpcRequest
	VpcRegion                string `position:"Query" name:"VpcRegion"`
	CenId                    string `position:"Query" name:"CenId"`
	FirewallVSwitchCidrBlock string `position:"Query" name:"FirewallVSwitchCidrBlock"`
	NetworkInstanceId        string `position:"Query" name:"NetworkInstanceId"`
	FirewallVpcCidrBlock     string `position:"Query" name:"FirewallVpcCidrBlock"`
	VpcFirewallName          string `position:"Query" name:"VpcFirewallName"`
	FirewallVpcZoneId        string `position:"Query" name:"FirewallVpcZoneId"`
	SourceIp                 string `position:"Query" name:"SourceIp"`
	Lang                     string `position:"Query" name:"Lang"`
	FirewallSwitch           string `position:"Query" name:"FirewallSwitch"`
	VSwitchId                string `position:"Query" name:"VSwitchId"`
	MemberUid                string `position:"Query" name:"MemberUid"`
}

// CreateVpcFirewallCenConfigureResponse is the response struct for api CreateVpcFirewallCenConfigure
type CreateVpcFirewallCenConfigureResponse struct {
	*responses.BaseResponse
	VpcFirewallId string `json:"VpcFirewallId" xml:"VpcFirewallId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateVpcFirewallCenConfigureRequest creates a request to invoke CreateVpcFirewallCenConfigure API
func CreateCreateVpcFirewallCenConfigureRequest() (request *CreateVpcFirewallCenConfigureRequest) {
	request = &CreateVpcFirewallCenConfigureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "CreateVpcFirewallCenConfigure", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVpcFirewallCenConfigureResponse creates a response to parse from CreateVpcFirewallCenConfigure response
func CreateCreateVpcFirewallCenConfigureResponse() (response *CreateVpcFirewallCenConfigureResponse) {
	response = &CreateVpcFirewallCenConfigureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
