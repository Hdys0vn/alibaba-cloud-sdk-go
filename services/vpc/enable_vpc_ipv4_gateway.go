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

// EnableVpcIpv4Gateway invokes the vpc.EnableVpcIpv4Gateway API synchronously
func (client *Client) EnableVpcIpv4Gateway(request *EnableVpcIpv4GatewayRequest) (response *EnableVpcIpv4GatewayResponse, err error) {
	response = CreateEnableVpcIpv4GatewayResponse()
	err = client.DoAction(request, response)
	return
}

// EnableVpcIpv4GatewayWithChan invokes the vpc.EnableVpcIpv4Gateway API asynchronously
func (client *Client) EnableVpcIpv4GatewayWithChan(request *EnableVpcIpv4GatewayRequest) (<-chan *EnableVpcIpv4GatewayResponse, <-chan error) {
	responseChan := make(chan *EnableVpcIpv4GatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableVpcIpv4Gateway(request)
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

// EnableVpcIpv4GatewayWithCallback invokes the vpc.EnableVpcIpv4Gateway API asynchronously
func (client *Client) EnableVpcIpv4GatewayWithCallback(request *EnableVpcIpv4GatewayRequest, callback func(response *EnableVpcIpv4GatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableVpcIpv4GatewayResponse
		var err error
		defer close(result)
		response, err = client.EnableVpcIpv4Gateway(request)
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

// EnableVpcIpv4GatewayRequest is the request struct for api EnableVpcIpv4Gateway
type EnableVpcIpv4GatewayRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Ipv4GatewayId        string           `position:"Query" name:"Ipv4GatewayId"`
	RouteTableList       *[]string        `position:"Query" name:"RouteTableList"  type:"Repeated"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// EnableVpcIpv4GatewayResponse is the response struct for api EnableVpcIpv4Gateway
type EnableVpcIpv4GatewayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableVpcIpv4GatewayRequest creates a request to invoke EnableVpcIpv4Gateway API
func CreateEnableVpcIpv4GatewayRequest() (request *EnableVpcIpv4GatewayRequest) {
	request = &EnableVpcIpv4GatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "EnableVpcIpv4Gateway", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableVpcIpv4GatewayResponse creates a response to parse from EnableVpcIpv4Gateway response
func CreateEnableVpcIpv4GatewayResponse() (response *EnableVpcIpv4GatewayResponse) {
	response = &EnableVpcIpv4GatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
