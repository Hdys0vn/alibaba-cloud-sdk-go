package pvtz

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

// UpdateResolverEndpoint invokes the pvtz.UpdateResolverEndpoint API synchronously
func (client *Client) UpdateResolverEndpoint(request *UpdateResolverEndpointRequest) (response *UpdateResolverEndpointResponse, err error) {
	response = CreateUpdateResolverEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateResolverEndpointWithChan invokes the pvtz.UpdateResolverEndpoint API asynchronously
func (client *Client) UpdateResolverEndpointWithChan(request *UpdateResolverEndpointRequest) (<-chan *UpdateResolverEndpointResponse, <-chan error) {
	responseChan := make(chan *UpdateResolverEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateResolverEndpoint(request)
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

// UpdateResolverEndpointWithCallback invokes the pvtz.UpdateResolverEndpoint API asynchronously
func (client *Client) UpdateResolverEndpointWithCallback(request *UpdateResolverEndpointRequest, callback func(response *UpdateResolverEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateResolverEndpointResponse
		var err error
		defer close(result)
		response, err = client.UpdateResolverEndpoint(request)
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

// UpdateResolverEndpointRequest is the request struct for api UpdateResolverEndpoint
type UpdateResolverEndpointRequest struct {
	*requests.RpcRequest
	EndpointId   string                            `position:"Query" name:"EndpointId"`
	UserClientIp string                            `position:"Query" name:"UserClientIp"`
	Name         string                            `position:"Query" name:"Name"`
	Lang         string                            `position:"Query" name:"Lang"`
	IpConfig     *[]UpdateResolverEndpointIpConfig `position:"Query" name:"IpConfig"  type:"Repeated"`
}

// UpdateResolverEndpointIpConfig is a repeated param struct in UpdateResolverEndpointRequest
type UpdateResolverEndpointIpConfig struct {
	VSwitchId string `name:"VSwitchId"`
	Ip        string `name:"Ip"`
	CidrBlock string `name:"CidrBlock"`
	AzId      string `name:"AzId"`
}

// UpdateResolverEndpointResponse is the response struct for api UpdateResolverEndpoint
type UpdateResolverEndpointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateResolverEndpointRequest creates a request to invoke UpdateResolverEndpoint API
func CreateUpdateResolverEndpointRequest() (request *UpdateResolverEndpointRequest) {
	request = &UpdateResolverEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "UpdateResolverEndpoint", "pvtz", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateResolverEndpointResponse creates a response to parse from UpdateResolverEndpoint response
func CreateUpdateResolverEndpointResponse() (response *UpdateResolverEndpointResponse) {
	response = &UpdateResolverEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}