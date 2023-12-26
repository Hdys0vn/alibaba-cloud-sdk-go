package iot

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

// ListDistributedProduct invokes the iot.ListDistributedProduct API synchronously
func (client *Client) ListDistributedProduct(request *ListDistributedProductRequest) (response *ListDistributedProductResponse, err error) {
	response = CreateListDistributedProductResponse()
	err = client.DoAction(request, response)
	return
}

// ListDistributedProductWithChan invokes the iot.ListDistributedProduct API asynchronously
func (client *Client) ListDistributedProductWithChan(request *ListDistributedProductRequest) (<-chan *ListDistributedProductResponse, <-chan error) {
	responseChan := make(chan *ListDistributedProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDistributedProduct(request)
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

// ListDistributedProductWithCallback invokes the iot.ListDistributedProduct API asynchronously
func (client *Client) ListDistributedProductWithCallback(request *ListDistributedProductRequest, callback func(response *ListDistributedProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDistributedProductResponse
		var err error
		defer close(result)
		response, err = client.ListDistributedProduct(request)
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

// ListDistributedProductRequest is the request struct for api ListDistributedProduct
type ListDistributedProductRequest struct {
	*requests.RpcRequest
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	SourceInstanceId string           `position:"Query" name:"SourceInstanceId"`
	CurrentPage      requests.Integer `position:"Query" name:"CurrentPage"`
	ProductKey       string           `position:"Query" name:"ProductKey"`
	TargetInstanceId string           `position:"Query" name:"TargetInstanceId"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
	TargetUid        string           `position:"Query" name:"TargetUid"`
}

// ListDistributedProductResponse is the response struct for api ListDistributedProduct
type ListDistributedProductResponse struct {
	*responses.BaseResponse
	RequestId    string                       `json:"RequestId" xml:"RequestId"`
	Success      bool                         `json:"Success" xml:"Success"`
	Code         string                       `json:"Code" xml:"Code"`
	ErrorMessage string                       `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInListDistributedProduct `json:"Data" xml:"Data"`
}

// CreateListDistributedProductRequest creates a request to invoke ListDistributedProduct API
func CreateListDistributedProductRequest() (request *ListDistributedProductRequest) {
	request = &ListDistributedProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ListDistributedProduct", "", "")
	request.Method = requests.POST
	return
}

// CreateListDistributedProductResponse creates a response to parse from ListDistributedProduct response
func CreateListDistributedProductResponse() (response *ListDistributedProductResponse) {
	response = &ListDistributedProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
