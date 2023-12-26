package cloudcallcenter

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

// GetAggregateTaobaoOrder invokes the cloudcallcenter.GetAggregateTaobaoOrder API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getaggregatetaobaoorder.html
func (client *Client) GetAggregateTaobaoOrder(request *GetAggregateTaobaoOrderRequest) (response *GetAggregateTaobaoOrderResponse, err error) {
	response = CreateGetAggregateTaobaoOrderResponse()
	err = client.DoAction(request, response)
	return
}

// GetAggregateTaobaoOrderWithChan invokes the cloudcallcenter.GetAggregateTaobaoOrder API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getaggregatetaobaoorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAggregateTaobaoOrderWithChan(request *GetAggregateTaobaoOrderRequest) (<-chan *GetAggregateTaobaoOrderResponse, <-chan error) {
	responseChan := make(chan *GetAggregateTaobaoOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAggregateTaobaoOrder(request)
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

// GetAggregateTaobaoOrderWithCallback invokes the cloudcallcenter.GetAggregateTaobaoOrder API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getaggregatetaobaoorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAggregateTaobaoOrderWithCallback(request *GetAggregateTaobaoOrderRequest, callback func(response *GetAggregateTaobaoOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAggregateTaobaoOrderResponse
		var err error
		defer close(result)
		response, err = client.GetAggregateTaobaoOrder(request)
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

// GetAggregateTaobaoOrderRequest is the request struct for api GetAggregateTaobaoOrder
type GetAggregateTaobaoOrderRequest struct {
	*requests.RpcRequest
	TaobaoUid requests.Integer `position:"Query" name:"TaobaoUid"`
}

// GetAggregateTaobaoOrderResponse is the response struct for api GetAggregateTaobaoOrder
type GetAggregateTaobaoOrderResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Success          bool   `json:"Success" xml:"Success"`
	Code             string `json:"Code" xml:"Code"`
	Message          string `json:"Message" xml:"Message"`
	HttpStatusCode   int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	IncomingAccount  int    `json:"IncomingAccount" xml:"IncomingAccount"`
	OutcomingAccount int    `json:"OutcomingAccount" xml:"OutcomingAccount"`
}

// CreateGetAggregateTaobaoOrderRequest creates a request to invoke GetAggregateTaobaoOrder API
func CreateGetAggregateTaobaoOrderRequest() (request *GetAggregateTaobaoOrderRequest) {
	request = &GetAggregateTaobaoOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetAggregateTaobaoOrder", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAggregateTaobaoOrderResponse creates a response to parse from GetAggregateTaobaoOrder response
func CreateGetAggregateTaobaoOrderResponse() (response *GetAggregateTaobaoOrderResponse) {
	response = &GetAggregateTaobaoOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}