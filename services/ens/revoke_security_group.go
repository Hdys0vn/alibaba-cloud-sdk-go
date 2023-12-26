package ens

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

// RevokeSecurityGroup invokes the ens.RevokeSecurityGroup API synchronously
func (client *Client) RevokeSecurityGroup(request *RevokeSecurityGroupRequest) (response *RevokeSecurityGroupResponse, err error) {
	response = CreateRevokeSecurityGroupResponse()
	err = client.DoAction(request, response)
	return
}

// RevokeSecurityGroupWithChan invokes the ens.RevokeSecurityGroup API asynchronously
func (client *Client) RevokeSecurityGroupWithChan(request *RevokeSecurityGroupRequest) (<-chan *RevokeSecurityGroupResponse, <-chan error) {
	responseChan := make(chan *RevokeSecurityGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeSecurityGroup(request)
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

// RevokeSecurityGroupWithCallback invokes the ens.RevokeSecurityGroup API asynchronously
func (client *Client) RevokeSecurityGroupWithCallback(request *RevokeSecurityGroupRequest, callback func(response *RevokeSecurityGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeSecurityGroupResponse
		var err error
		defer close(result)
		response, err = client.RevokeSecurityGroup(request)
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

// RevokeSecurityGroupRequest is the request struct for api RevokeSecurityGroup
type RevokeSecurityGroupRequest struct {
	*requests.RpcRequest
	SourcePortRange string           `position:"Query" name:"SourcePortRange"`
	PortRange       string           `position:"Query" name:"PortRange"`
	IpProtocol      string           `position:"Query" name:"IpProtocol"`
	SourceCidrIp    string           `position:"Query" name:"SourceCidrIp"`
	Priority        requests.Integer `position:"Query" name:"Priority"`
	SecurityGroupId string           `position:"Query" name:"SecurityGroupId"`
	Policy          string           `position:"Query" name:"Policy"`
}

// RevokeSecurityGroupResponse is the response struct for api RevokeSecurityGroup
type RevokeSecurityGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRevokeSecurityGroupRequest creates a request to invoke RevokeSecurityGroup API
func CreateRevokeSecurityGroupRequest() (request *RevokeSecurityGroupRequest) {
	request = &RevokeSecurityGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "RevokeSecurityGroup", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRevokeSecurityGroupResponse creates a response to parse from RevokeSecurityGroup response
func CreateRevokeSecurityGroupResponse() (response *RevokeSecurityGroupResponse) {
	response = &RevokeSecurityGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
