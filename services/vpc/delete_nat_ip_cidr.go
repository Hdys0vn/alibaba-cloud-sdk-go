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

// DeleteNatIpCidr invokes the vpc.DeleteNatIpCidr API synchronously
func (client *Client) DeleteNatIpCidr(request *DeleteNatIpCidrRequest) (response *DeleteNatIpCidrResponse, err error) {
	response = CreateDeleteNatIpCidrResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNatIpCidrWithChan invokes the vpc.DeleteNatIpCidr API asynchronously
func (client *Client) DeleteNatIpCidrWithChan(request *DeleteNatIpCidrRequest) (<-chan *DeleteNatIpCidrResponse, <-chan error) {
	responseChan := make(chan *DeleteNatIpCidrResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNatIpCidr(request)
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

// DeleteNatIpCidrWithCallback invokes the vpc.DeleteNatIpCidr API asynchronously
func (client *Client) DeleteNatIpCidrWithCallback(request *DeleteNatIpCidrRequest, callback func(response *DeleteNatIpCidrResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNatIpCidrResponse
		var err error
		defer close(result)
		response, err = client.DeleteNatIpCidr(request)
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

// DeleteNatIpCidrRequest is the request struct for api DeleteNatIpCidr
type DeleteNatIpCidrRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	NatGatewayId         string           `position:"Query" name:"NatGatewayId"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NatIpCidr            string           `position:"Query" name:"NatIpCidr"`
}

// DeleteNatIpCidrResponse is the response struct for api DeleteNatIpCidr
type DeleteNatIpCidrResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteNatIpCidrRequest creates a request to invoke DeleteNatIpCidr API
func CreateDeleteNatIpCidrRequest() (request *DeleteNatIpCidrRequest) {
	request = &DeleteNatIpCidrRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteNatIpCidr", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteNatIpCidrResponse creates a response to parse from DeleteNatIpCidr response
func CreateDeleteNatIpCidrResponse() (response *DeleteNatIpCidrResponse) {
	response = &DeleteNatIpCidrResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
