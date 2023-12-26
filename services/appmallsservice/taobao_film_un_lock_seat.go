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

// TaobaoFilmUnLockSeat invokes the appmallsservice.TaobaoFilmUnLockSeat API synchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmunlockseat.html
func (client *Client) TaobaoFilmUnLockSeat(request *TaobaoFilmUnLockSeatRequest) (response *TaobaoFilmUnLockSeatResponse, err error) {
	response = CreateTaobaoFilmUnLockSeatResponse()
	err = client.DoAction(request, response)
	return
}

// TaobaoFilmUnLockSeatWithChan invokes the appmallsservice.TaobaoFilmUnLockSeat API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmunlockseat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmUnLockSeatWithChan(request *TaobaoFilmUnLockSeatRequest) (<-chan *TaobaoFilmUnLockSeatResponse, <-chan error) {
	responseChan := make(chan *TaobaoFilmUnLockSeatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaobaoFilmUnLockSeat(request)
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

// TaobaoFilmUnLockSeatWithCallback invokes the appmallsservice.TaobaoFilmUnLockSeat API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmunlockseat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmUnLockSeatWithCallback(request *TaobaoFilmUnLockSeatRequest, callback func(response *TaobaoFilmUnLockSeatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaobaoFilmUnLockSeatResponse
		var err error
		defer close(result)
		response, err = client.TaobaoFilmUnLockSeat(request)
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

// TaobaoFilmUnLockSeatRequest is the request struct for api TaobaoFilmUnLockSeat
type TaobaoFilmUnLockSeatRequest struct {
	*requests.RpcRequest
	LockSeatApplyKey string `position:"Query" name:"LockSeatApplyKey"`
	ExtUserId        string `position:"Query" name:"ExtUserId"`
	ParamsJson       string `position:"Query" name:"ParamsJson"`
}

// TaobaoFilmUnLockSeatResponse is the response struct for api TaobaoFilmUnLockSeat
type TaobaoFilmUnLockSeatResponse struct {
	*responses.BaseResponse
	ErrorCode     string `json:"ErrorCode" xml:"ErrorCode"`
	Msg           string `json:"Msg" xml:"Msg"`
	SubCode       string `json:"SubCode" xml:"SubCode"`
	SubMsg        string `json:"SubMsg" xml:"SubMsg"`
	ReturnCode    string `json:"ReturnCode" xml:"ReturnCode"`
	ReturnMessage string `json:"ReturnMessage" xml:"ReturnMessage"`
	ReturnValue   bool   `json:"ReturnValue" xml:"ReturnValue"`
	LogsId        string `json:"LogsId" xml:"LogsId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateTaobaoFilmUnLockSeatRequest creates a request to invoke TaobaoFilmUnLockSeat API
func CreateTaobaoFilmUnLockSeatRequest() (request *TaobaoFilmUnLockSeatRequest) {
	request = &TaobaoFilmUnLockSeatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("AppMallsService", "2018-02-24", "TaobaoFilmUnLockSeat", "", "")
	return
}

// CreateTaobaoFilmUnLockSeatResponse creates a response to parse from TaobaoFilmUnLockSeat response
func CreateTaobaoFilmUnLockSeatResponse() (response *TaobaoFilmUnLockSeatResponse) {
	response = &TaobaoFilmUnLockSeatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}