package baas

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

// DescribeAntChainTransactionReceipt invokes the baas.DescribeAntChainTransactionReceipt API synchronously
// api document: https://help.aliyun.com/api/baas/describeantchaintransactionreceipt.html
func (client *Client) DescribeAntChainTransactionReceipt(request *DescribeAntChainTransactionReceiptRequest) (response *DescribeAntChainTransactionReceiptResponse, err error) {
	response = CreateDescribeAntChainTransactionReceiptResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAntChainTransactionReceiptWithChan invokes the baas.DescribeAntChainTransactionReceipt API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchaintransactionreceipt.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainTransactionReceiptWithChan(request *DescribeAntChainTransactionReceiptRequest) (<-chan *DescribeAntChainTransactionReceiptResponse, <-chan error) {
	responseChan := make(chan *DescribeAntChainTransactionReceiptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAntChainTransactionReceipt(request)
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

// DescribeAntChainTransactionReceiptWithCallback invokes the baas.DescribeAntChainTransactionReceipt API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchaintransactionreceipt.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainTransactionReceiptWithCallback(request *DescribeAntChainTransactionReceiptRequest, callback func(response *DescribeAntChainTransactionReceiptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAntChainTransactionReceiptResponse
		var err error
		defer close(result)
		response, err = client.DescribeAntChainTransactionReceipt(request)
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

// DescribeAntChainTransactionReceiptRequest is the request struct for api DescribeAntChainTransactionReceipt
type DescribeAntChainTransactionReceiptRequest struct {
	*requests.RpcRequest
	AntChainId string `position:"Body" name:"AntChainId"`
	Hash       string `position:"Body" name:"Hash"`
}

// DescribeAntChainTransactionReceiptResponse is the response struct for api DescribeAntChainTransactionReceipt
type DescribeAntChainTransactionReceiptResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeAntChainTransactionReceiptRequest creates a request to invoke DescribeAntChainTransactionReceipt API
func CreateDescribeAntChainTransactionReceiptRequest() (request *DescribeAntChainTransactionReceiptRequest) {
	request = &DescribeAntChainTransactionReceiptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeAntChainTransactionReceipt", "baas", "openAPI")
	return
}

// CreateDescribeAntChainTransactionReceiptResponse creates a response to parse from DescribeAntChainTransactionReceipt response
func CreateDescribeAntChainTransactionReceiptResponse() (response *DescribeAntChainTransactionReceiptResponse) {
	response = &DescribeAntChainTransactionReceiptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
