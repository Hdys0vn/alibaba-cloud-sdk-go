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

// ModifyVpcAttribute invokes the vpc.ModifyVpcAttribute API synchronously
func (client *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
	response = CreateModifyVpcAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpcAttributeWithChan invokes the vpc.ModifyVpcAttribute API asynchronously
func (client *Client) ModifyVpcAttributeWithChan(request *ModifyVpcAttributeRequest) (<-chan *ModifyVpcAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVpcAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpcAttribute(request)
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

// ModifyVpcAttributeWithCallback invokes the vpc.ModifyVpcAttribute API asynchronously
func (client *Client) ModifyVpcAttributeWithCallback(request *ModifyVpcAttributeRequest, callback func(response *ModifyVpcAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpcAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpcAttribute(request)
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

// ModifyVpcAttributeRequest is the request struct for api ModifyVpcAttribute
type ModifyVpcAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EnableIPv6           requests.Boolean `position:"Query" name:"EnableIPv6"`
	Description          string           `position:"Query" name:"Description"`
	VpcName              string           `position:"Query" name:"VpcName"`
	Ipv6Isp              string           `position:"Query" name:"Ipv6Isp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Ipv6CidrBlock        string           `position:"Query" name:"Ipv6CidrBlock"`
	VpcId                string           `position:"Query" name:"VpcId"`
	CidrBlock            string           `position:"Query" name:"CidrBlock"`
}

// ModifyVpcAttributeResponse is the response struct for api ModifyVpcAttribute
type ModifyVpcAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVpcAttributeRequest creates a request to invoke ModifyVpcAttribute API
func CreateModifyVpcAttributeRequest() (request *ModifyVpcAttributeRequest) {
	request = &ModifyVpcAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpcAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpcAttributeResponse creates a response to parse from ModifyVpcAttribute response
func CreateModifyVpcAttributeResponse() (response *ModifyVpcAttributeResponse) {
	response = &ModifyVpcAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
