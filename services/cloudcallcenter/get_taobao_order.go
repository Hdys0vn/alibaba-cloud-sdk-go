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

// GetTaobaoOrder invokes the cloudcallcenter.GetTaobaoOrder API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/gettaobaoorder.html
func (client *Client) GetTaobaoOrder(request *GetTaobaoOrderRequest) (response *GetTaobaoOrderResponse, err error) {
	response = CreateGetTaobaoOrderResponse()
	err = client.DoAction(request, response)
	return
}

// GetTaobaoOrderWithChan invokes the cloudcallcenter.GetTaobaoOrder API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/gettaobaoorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTaobaoOrderWithChan(request *GetTaobaoOrderRequest) (<-chan *GetTaobaoOrderResponse, <-chan error) {
	responseChan := make(chan *GetTaobaoOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTaobaoOrder(request)
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

// GetTaobaoOrderWithCallback invokes the cloudcallcenter.GetTaobaoOrder API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/gettaobaoorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTaobaoOrderWithCallback(request *GetTaobaoOrderRequest, callback func(response *GetTaobaoOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTaobaoOrderResponse
		var err error
		defer close(result)
		response, err = client.GetTaobaoOrder(request)
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

// GetTaobaoOrderRequest is the request struct for api GetTaobaoOrder
type GetTaobaoOrderRequest struct {
	*requests.RpcRequest
	EndTime         requests.Integer `position:"Query" name:"endTime"`
	StartTime       requests.Integer `position:"Query" name:"startTime"`
	TaobaoUid       requests.Integer `position:"Query" name:"TaobaoUid"`
	Type            requests.Integer `position:"Query" name:"type"`
	IncludeAll      requests.Boolean `position:"Query" name:"includeAll"`
	ArticleCode     string           `position:"Query" name:"articleCode"`
	ArticleItemCode string           `position:"Query" name:"articleItemCode"`
	ExpiredOnly     requests.Boolean `position:"Query" name:"expiredOnly"`
}

// GetTaobaoOrderResponse is the response struct for api GetTaobaoOrder
type GetTaobaoOrderResponse struct {
	*responses.BaseResponse
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	Success        bool                   `json:"Success" xml:"Success"`
	Code           string                 `json:"Code" xml:"Code"`
	Message        string                 `json:"Message" xml:"Message"`
	HttpStatusCode int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Orders         OrdersInGetTaobaoOrder `json:"Orders" xml:"Orders"`
}

// CreateGetTaobaoOrderRequest creates a request to invoke GetTaobaoOrder API
func CreateGetTaobaoOrderRequest() (request *GetTaobaoOrderRequest) {
	request = &GetTaobaoOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetTaobaoOrder", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTaobaoOrderResponse creates a response to parse from GetTaobaoOrder response
func CreateGetTaobaoOrderResponse() (response *GetTaobaoOrderResponse) {
	response = &GetTaobaoOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
