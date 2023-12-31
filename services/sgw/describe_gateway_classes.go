package sgw

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

// DescribeGatewayClasses invokes the sgw.DescribeGatewayClasses API synchronously
func (client *Client) DescribeGatewayClasses(request *DescribeGatewayClassesRequest) (response *DescribeGatewayClassesResponse, err error) {
	response = CreateDescribeGatewayClassesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGatewayClassesWithChan invokes the sgw.DescribeGatewayClasses API asynchronously
func (client *Client) DescribeGatewayClassesWithChan(request *DescribeGatewayClassesRequest) (<-chan *DescribeGatewayClassesResponse, <-chan error) {
	responseChan := make(chan *DescribeGatewayClassesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGatewayClasses(request)
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

// DescribeGatewayClassesWithCallback invokes the sgw.DescribeGatewayClasses API asynchronously
func (client *Client) DescribeGatewayClassesWithCallback(request *DescribeGatewayClassesRequest, callback func(response *DescribeGatewayClassesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGatewayClassesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGatewayClasses(request)
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

// DescribeGatewayClassesRequest is the request struct for api DescribeGatewayClasses
type DescribeGatewayClassesRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribeGatewayClassesResponse is the response struct for api DescribeGatewayClasses
type DescribeGatewayClassesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Classes   string `json:"Classes" xml:"Classes"`
}

// CreateDescribeGatewayClassesRequest creates a request to invoke DescribeGatewayClasses API
func CreateDescribeGatewayClassesRequest() (request *DescribeGatewayClassesRequest) {
	request = &DescribeGatewayClassesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeGatewayClasses", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGatewayClassesResponse creates a response to parse from DescribeGatewayClasses response
func CreateDescribeGatewayClassesResponse() (response *DescribeGatewayClassesResponse) {
	response = &DescribeGatewayClassesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
