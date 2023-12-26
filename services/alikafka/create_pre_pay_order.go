package alikafka

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

// CreatePrePayOrder invokes the alikafka.CreatePrePayOrder API synchronously
func (client *Client) CreatePrePayOrder(request *CreatePrePayOrderRequest) (response *CreatePrePayOrderResponse, err error) {
	response = CreateCreatePrePayOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePrePayOrderWithChan invokes the alikafka.CreatePrePayOrder API asynchronously
func (client *Client) CreatePrePayOrderWithChan(request *CreatePrePayOrderRequest) (<-chan *CreatePrePayOrderResponse, <-chan error) {
	responseChan := make(chan *CreatePrePayOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePrePayOrder(request)
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

// CreatePrePayOrderWithCallback invokes the alikafka.CreatePrePayOrder API asynchronously
func (client *Client) CreatePrePayOrderWithCallback(request *CreatePrePayOrderRequest, callback func(response *CreatePrePayOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePrePayOrderResponse
		var err error
		defer close(result)
		response, err = client.CreatePrePayOrder(request)
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

// CreatePrePayOrderRequest is the request struct for api CreatePrePayOrder
type CreatePrePayOrderRequest struct {
	*requests.RpcRequest
	IoMax           requests.Integer        `position:"Query" name:"IoMax"`
	EipMax          requests.Integer        `position:"Query" name:"EipMax"`
	SpecType        string                  `position:"Query" name:"SpecType"`
	ResourceGroupId string                  `position:"Query" name:"ResourceGroupId"`
	Tag             *[]CreatePrePayOrderTag `position:"Query" name:"Tag"  type:"Repeated"`
	PartitionNum    requests.Integer        `position:"Query" name:"PartitionNum"`
	DiskSize        requests.Integer        `position:"Query" name:"DiskSize"`
	IoMaxSpec       string                  `position:"Query" name:"IoMaxSpec"`
	DiskType        string                  `position:"Query" name:"DiskType"`
	TopicQuota      requests.Integer        `position:"Query" name:"TopicQuota"`
	DeployType      requests.Integer        `position:"Query" name:"DeployType"`
}

// CreatePrePayOrderTag is a repeated param struct in CreatePrePayOrderRequest
type CreatePrePayOrderTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreatePrePayOrderResponse is the response struct for api CreatePrePayOrder
type CreatePrePayOrderResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateCreatePrePayOrderRequest creates a request to invoke CreatePrePayOrder API
func CreateCreatePrePayOrderRequest() (request *CreatePrePayOrderRequest) {
	request = &CreatePrePayOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "CreatePrePayOrder", "", "")
	request.Method = requests.POST
	return
}

// CreateCreatePrePayOrderResponse creates a response to parse from CreatePrePayOrder response
func CreateCreatePrePayOrderResponse() (response *CreatePrePayOrderResponse) {
	response = &CreatePrePayOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}