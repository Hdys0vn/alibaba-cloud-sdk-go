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

// CreateDefaultVpc invokes the vpc.CreateDefaultVpc API synchronously
func (client *Client) CreateDefaultVpc(request *CreateDefaultVpcRequest) (response *CreateDefaultVpcResponse, err error) {
	response = CreateCreateDefaultVpcResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDefaultVpcWithChan invokes the vpc.CreateDefaultVpc API asynchronously
func (client *Client) CreateDefaultVpcWithChan(request *CreateDefaultVpcRequest) (<-chan *CreateDefaultVpcResponse, <-chan error) {
	responseChan := make(chan *CreateDefaultVpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDefaultVpc(request)
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

// CreateDefaultVpcWithCallback invokes the vpc.CreateDefaultVpc API asynchronously
func (client *Client) CreateDefaultVpcWithCallback(request *CreateDefaultVpcRequest, callback func(response *CreateDefaultVpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDefaultVpcResponse
		var err error
		defer close(result)
		response, err = client.CreateDefaultVpc(request)
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

// CreateDefaultVpcRequest is the request struct for api CreateDefaultVpc
type CreateDefaultVpcRequest struct {
	*requests.RpcRequest
	EnableDefaultVSwitch requests.Boolean `position:"Query" name:"EnableDefaultVSwitch"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	EnableIpv6           requests.Boolean `position:"Query" name:"EnableIpv6"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Ipv6CidrBlock        string           `position:"Query" name:"Ipv6CidrBlock"`
	ZoneId               *[]string        `position:"Query" name:"ZoneId"  type:"Repeated"`
}

// CreateDefaultVpcResponse is the response struct for api CreateDefaultVpc
type CreateDefaultVpcResponse struct {
	*responses.BaseResponse
	VpcId           string           `json:"VpcId" xml:"VpcId"`
	VRouterId       string           `json:"VRouterId" xml:"VRouterId"`
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	RouteTableId    string           `json:"RouteTableId" xml:"RouteTableId"`
	DefaultVSwitchs []DefaultVSwitch `json:"DefaultVSwitchs" xml:"DefaultVSwitchs"`
}

// CreateCreateDefaultVpcRequest creates a request to invoke CreateDefaultVpc API
func CreateCreateDefaultVpcRequest() (request *CreateDefaultVpcRequest) {
	request = &CreateDefaultVpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateDefaultVpc", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDefaultVpcResponse creates a response to parse from CreateDefaultVpc response
func CreateCreateDefaultVpcResponse() (response *CreateDefaultVpcResponse) {
	response = &CreateDefaultVpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
