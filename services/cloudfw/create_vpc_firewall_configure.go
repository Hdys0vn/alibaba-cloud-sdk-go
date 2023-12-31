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

// CreateVpcFirewallConfigure invokes the cloudfw.CreateVpcFirewallConfigure API synchronously
func (client *Client) CreateVpcFirewallConfigure(request *CreateVpcFirewallConfigureRequest) (response *CreateVpcFirewallConfigureResponse, err error) {
	response = CreateCreateVpcFirewallConfigureResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVpcFirewallConfigureWithChan invokes the cloudfw.CreateVpcFirewallConfigure API asynchronously
func (client *Client) CreateVpcFirewallConfigureWithChan(request *CreateVpcFirewallConfigureRequest) (<-chan *CreateVpcFirewallConfigureResponse, <-chan error) {
	responseChan := make(chan *CreateVpcFirewallConfigureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVpcFirewallConfigure(request)
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

// CreateVpcFirewallConfigureWithCallback invokes the cloudfw.CreateVpcFirewallConfigure API asynchronously
func (client *Client) CreateVpcFirewallConfigureWithCallback(request *CreateVpcFirewallConfigureRequest, callback func(response *CreateVpcFirewallConfigureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVpcFirewallConfigureResponse
		var err error
		defer close(result)
		response, err = client.CreateVpcFirewallConfigure(request)
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

// CreateVpcFirewallConfigureRequest is the request struct for api CreateVpcFirewallConfigure
type CreateVpcFirewallConfigureRequest struct {
	*requests.RpcRequest
	PeerVpcCidrTableList  string `position:"Query" name:"PeerVpcCidrTableList"`
	LocalVpcCidrTableList string `position:"Query" name:"LocalVpcCidrTableList"`
	LocalVpcRegion        string `position:"Query" name:"LocalVpcRegion"`
	PeerVpcRegion         string `position:"Query" name:"PeerVpcRegion"`
	VpcFirewallName       string `position:"Query" name:"VpcFirewallName"`
	LocalVpcId            string `position:"Query" name:"LocalVpcId"`
	SourceIp              string `position:"Query" name:"SourceIp"`
	Lang                  string `position:"Query" name:"Lang"`
	FirewallSwitch        string `position:"Query" name:"FirewallSwitch"`
	MemberUid             string `position:"Query" name:"MemberUid"`
	PeerVpcId             string `position:"Query" name:"PeerVpcId"`
}

// CreateVpcFirewallConfigureResponse is the response struct for api CreateVpcFirewallConfigure
type CreateVpcFirewallConfigureResponse struct {
	*responses.BaseResponse
	VpcFirewallId string `json:"VpcFirewallId" xml:"VpcFirewallId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateVpcFirewallConfigureRequest creates a request to invoke CreateVpcFirewallConfigure API
func CreateCreateVpcFirewallConfigureRequest() (request *CreateVpcFirewallConfigureRequest) {
	request = &CreateVpcFirewallConfigureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "CreateVpcFirewallConfigure", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVpcFirewallConfigureResponse creates a response to parse from CreateVpcFirewallConfigure response
func CreateCreateVpcFirewallConfigureResponse() (response *CreateVpcFirewallConfigureResponse) {
	response = &CreateVpcFirewallConfigureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
