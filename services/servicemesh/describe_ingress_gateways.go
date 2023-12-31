package servicemesh

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

// DescribeIngressGateways invokes the servicemesh.DescribeIngressGateways API synchronously
func (client *Client) DescribeIngressGateways(request *DescribeIngressGatewaysRequest) (response *DescribeIngressGatewaysResponse, err error) {
	response = CreateDescribeIngressGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIngressGatewaysWithChan invokes the servicemesh.DescribeIngressGateways API asynchronously
func (client *Client) DescribeIngressGatewaysWithChan(request *DescribeIngressGatewaysRequest) (<-chan *DescribeIngressGatewaysResponse, <-chan error) {
	responseChan := make(chan *DescribeIngressGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIngressGateways(request)
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

// DescribeIngressGatewaysWithCallback invokes the servicemesh.DescribeIngressGateways API asynchronously
func (client *Client) DescribeIngressGatewaysWithCallback(request *DescribeIngressGatewaysRequest, callback func(response *DescribeIngressGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIngressGatewaysResponse
		var err error
		defer close(result)
		response, err = client.DescribeIngressGateways(request)
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

// DescribeIngressGatewaysRequest is the request struct for api DescribeIngressGateways
type DescribeIngressGatewaysRequest struct {
	*requests.RpcRequest
	ServiceMeshId string `position:"Query" name:"ServiceMeshId"`
}

// DescribeIngressGatewaysResponse is the response struct for api DescribeIngressGateways
type DescribeIngressGatewaysResponse struct {
	*responses.BaseResponse
	RequestId       string                   `json:"RequestId" xml:"RequestId"`
	IngressGateways []map[string]interface{} `json:"IngressGateways" xml:"IngressGateways"`
}

// CreateDescribeIngressGatewaysRequest creates a request to invoke DescribeIngressGateways API
func CreateDescribeIngressGatewaysRequest() (request *DescribeIngressGatewaysRequest) {
	request = &DescribeIngressGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("servicemesh", "2020-01-11", "DescribeIngressGateways", "servicemesh", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeIngressGatewaysResponse creates a response to parse from DescribeIngressGateways response
func CreateDescribeIngressGatewaysResponse() (response *DescribeIngressGatewaysResponse) {
	response = &DescribeIngressGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
