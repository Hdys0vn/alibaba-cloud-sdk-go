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

// DescribeGatewayStock invokes the sgw.DescribeGatewayStock API synchronously
func (client *Client) DescribeGatewayStock(request *DescribeGatewayStockRequest) (response *DescribeGatewayStockResponse, err error) {
	response = CreateDescribeGatewayStockResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGatewayStockWithChan invokes the sgw.DescribeGatewayStock API asynchronously
func (client *Client) DescribeGatewayStockWithChan(request *DescribeGatewayStockRequest) (<-chan *DescribeGatewayStockResponse, <-chan error) {
	responseChan := make(chan *DescribeGatewayStockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGatewayStock(request)
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

// DescribeGatewayStockWithCallback invokes the sgw.DescribeGatewayStock API asynchronously
func (client *Client) DescribeGatewayStockWithCallback(request *DescribeGatewayStockRequest, callback func(response *DescribeGatewayStockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGatewayStockResponse
		var err error
		defer close(result)
		response, err = client.DescribeGatewayStock(request)
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

// DescribeGatewayStockRequest is the request struct for api DescribeGatewayStock
type DescribeGatewayStockRequest struct {
	*requests.RpcRequest
	GatewayRegionId string `position:"Query" name:"GatewayRegionId"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
}

// DescribeGatewayStockResponse is the response struct for api DescribeGatewayStock
type DescribeGatewayStockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Stocks    Stocks `json:"Stocks" xml:"Stocks"`
}

// CreateDescribeGatewayStockRequest creates a request to invoke DescribeGatewayStock API
func CreateDescribeGatewayStockRequest() (request *DescribeGatewayStockRequest) {
	request = &DescribeGatewayStockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeGatewayStock", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGatewayStockResponse creates a response to parse from DescribeGatewayStock response
func CreateDescribeGatewayStockResponse() (response *DescribeGatewayStockResponse) {
	response = &DescribeGatewayStockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
