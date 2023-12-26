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

// AttachDhcpOptionsSetToVpc invokes the vpc.AttachDhcpOptionsSetToVpc API synchronously
func (client *Client) AttachDhcpOptionsSetToVpc(request *AttachDhcpOptionsSetToVpcRequest) (response *AttachDhcpOptionsSetToVpcResponse, err error) {
	response = CreateAttachDhcpOptionsSetToVpcResponse()
	err = client.DoAction(request, response)
	return
}

// AttachDhcpOptionsSetToVpcWithChan invokes the vpc.AttachDhcpOptionsSetToVpc API asynchronously
func (client *Client) AttachDhcpOptionsSetToVpcWithChan(request *AttachDhcpOptionsSetToVpcRequest) (<-chan *AttachDhcpOptionsSetToVpcResponse, <-chan error) {
	responseChan := make(chan *AttachDhcpOptionsSetToVpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachDhcpOptionsSetToVpc(request)
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

// AttachDhcpOptionsSetToVpcWithCallback invokes the vpc.AttachDhcpOptionsSetToVpc API asynchronously
func (client *Client) AttachDhcpOptionsSetToVpcWithCallback(request *AttachDhcpOptionsSetToVpcRequest, callback func(response *AttachDhcpOptionsSetToVpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachDhcpOptionsSetToVpcResponse
		var err error
		defer close(result)
		response, err = client.AttachDhcpOptionsSetToVpc(request)
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

// AttachDhcpOptionsSetToVpcRequest is the request struct for api AttachDhcpOptionsSetToVpc
type AttachDhcpOptionsSetToVpcRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	DhcpOptionsSetId     string           `position:"Query" name:"DhcpOptionsSetId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
}

// AttachDhcpOptionsSetToVpcResponse is the response struct for api AttachDhcpOptionsSetToVpc
type AttachDhcpOptionsSetToVpcResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachDhcpOptionsSetToVpcRequest creates a request to invoke AttachDhcpOptionsSetToVpc API
func CreateAttachDhcpOptionsSetToVpcRequest() (request *AttachDhcpOptionsSetToVpcRequest) {
	request = &AttachDhcpOptionsSetToVpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AttachDhcpOptionsSetToVpc", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachDhcpOptionsSetToVpcResponse creates a response to parse from AttachDhcpOptionsSetToVpc response
func CreateAttachDhcpOptionsSetToVpcResponse() (response *AttachDhcpOptionsSetToVpcResponse) {
	response = &AttachDhcpOptionsSetToVpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
