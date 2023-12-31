package appmallsservice

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

// TaobaoFilmIssueOrder invokes the appmallsservice.TaobaoFilmIssueOrder API synchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmissueorder.html
func (client *Client) TaobaoFilmIssueOrder(request *TaobaoFilmIssueOrderRequest) (response *TaobaoFilmIssueOrderResponse, err error) {
	response = CreateTaobaoFilmIssueOrderResponse()
	err = client.DoAction(request, response)
	return
}

// TaobaoFilmIssueOrderWithChan invokes the appmallsservice.TaobaoFilmIssueOrder API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmissueorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmIssueOrderWithChan(request *TaobaoFilmIssueOrderRequest) (<-chan *TaobaoFilmIssueOrderResponse, <-chan error) {
	responseChan := make(chan *TaobaoFilmIssueOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaobaoFilmIssueOrder(request)
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

// TaobaoFilmIssueOrderWithCallback invokes the appmallsservice.TaobaoFilmIssueOrder API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmissueorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmIssueOrderWithCallback(request *TaobaoFilmIssueOrderRequest, callback func(response *TaobaoFilmIssueOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaobaoFilmIssueOrderResponse
		var err error
		defer close(result)
		response, err = client.TaobaoFilmIssueOrder(request)
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

// TaobaoFilmIssueOrderRequest is the request struct for api TaobaoFilmIssueOrder
type TaobaoFilmIssueOrderRequest struct {
	*requests.RpcRequest
	LockSeatApplyKey string           `position:"Query" name:"LockSeatApplyKey"`
	ExtUserId        string           `position:"Query" name:"ExtUserId"`
	ExtOrderId       string           `position:"Query" name:"ExtOrderId"`
	TotalPrice       requests.Integer `position:"Query" name:"TotalPrice"`
	ParamsJson       string           `position:"Query" name:"ParamsJson"`
}

// TaobaoFilmIssueOrderResponse is the response struct for api TaobaoFilmIssueOrder
type TaobaoFilmIssueOrderResponse struct {
	*responses.BaseResponse
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Msg            string `json:"Msg" xml:"Msg"`
	SubCode        string `json:"SubCode" xml:"SubCode"`
	SubMsg         string `json:"SubMsg" xml:"SubMsg"`
	Message        string `json:"Message" xml:"Message"`
	Status         string `json:"Status" xml:"Status"`
	TbOrderId      string `json:"TbOrderId" xml:"TbOrderId"`
	TicketContents string `json:"TicketContents" xml:"TicketContents"`
	LogsId         string `json:"LogsId" xml:"LogsId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateTaobaoFilmIssueOrderRequest creates a request to invoke TaobaoFilmIssueOrder API
func CreateTaobaoFilmIssueOrderRequest() (request *TaobaoFilmIssueOrderRequest) {
	request = &TaobaoFilmIssueOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("AppMallsService", "2018-02-24", "TaobaoFilmIssueOrder", "", "")
	return
}

// CreateTaobaoFilmIssueOrderResponse creates a response to parse from TaobaoFilmIssueOrder response
func CreateTaobaoFilmIssueOrderResponse() (response *TaobaoFilmIssueOrderResponse) {
	response = &TaobaoFilmIssueOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
