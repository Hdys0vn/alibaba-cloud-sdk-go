package cloudwf

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

// InnerRefund invokes the cloudwf.InnerRefund API synchronously
// api document: https://help.aliyun.com/api/cloudwf/innerrefund.html
func (client *Client) InnerRefund(request *InnerRefundRequest) (response *InnerRefundResponse, err error) {
	response = CreateInnerRefundResponse()
	err = client.DoAction(request, response)
	return
}

// InnerRefundWithChan invokes the cloudwf.InnerRefund API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/innerrefund.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InnerRefundWithChan(request *InnerRefundRequest) (<-chan *InnerRefundResponse, <-chan error) {
	responseChan := make(chan *InnerRefundResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InnerRefund(request)
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

// InnerRefundWithCallback invokes the cloudwf.InnerRefund API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/innerrefund.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InnerRefundWithCallback(request *InnerRefundRequest, callback func(response *InnerRefundResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InnerRefundResponse
		var err error
		defer close(result)
		response, err = client.InnerRefund(request)
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

// InnerRefundRequest is the request struct for api InnerRefund
type InnerRefundRequest struct {
	*requests.RpcRequest
	Data string `position:"Query" name:"data"`
}

// InnerRefundResponse is the response struct for api InnerRefund
type InnerRefundResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"success" xml:"success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"data" xml:"data"`
}

// CreateInnerRefundRequest creates a request to invoke InnerRefund API
func CreateInnerRefundRequest() (request *InnerRefundRequest) {
	request = &InnerRefundRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "InnerRefund", "cloudwf", "openAPI")
	return
}

// CreateInnerRefundResponse creates a response to parse from InnerRefund response
func CreateInnerRefundResponse() (response *InnerRefundResponse) {
	response = &InnerRefundResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
