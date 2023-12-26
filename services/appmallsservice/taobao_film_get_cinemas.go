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

// TaobaoFilmGetCinemas invokes the appmallsservice.TaobaoFilmGetCinemas API synchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmgetcinemas.html
func (client *Client) TaobaoFilmGetCinemas(request *TaobaoFilmGetCinemasRequest) (response *TaobaoFilmGetCinemasResponse, err error) {
	response = CreateTaobaoFilmGetCinemasResponse()
	err = client.DoAction(request, response)
	return
}

// TaobaoFilmGetCinemasWithChan invokes the appmallsservice.TaobaoFilmGetCinemas API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmgetcinemas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmGetCinemasWithChan(request *TaobaoFilmGetCinemasRequest) (<-chan *TaobaoFilmGetCinemasResponse, <-chan error) {
	responseChan := make(chan *TaobaoFilmGetCinemasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaobaoFilmGetCinemas(request)
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

// TaobaoFilmGetCinemasWithCallback invokes the appmallsservice.TaobaoFilmGetCinemas API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmgetcinemas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmGetCinemasWithCallback(request *TaobaoFilmGetCinemasRequest, callback func(response *TaobaoFilmGetCinemasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaobaoFilmGetCinemasResponse
		var err error
		defer close(result)
		response, err = client.TaobaoFilmGetCinemas(request)
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

// TaobaoFilmGetCinemasRequest is the request struct for api TaobaoFilmGetCinemas
type TaobaoFilmGetCinemasRequest struct {
	*requests.RpcRequest
	PageIndex  requests.Integer `position:"Query" name:"PageIndex"`
	ParamsJson string           `position:"Query" name:"ParamsJson"`
}

// TaobaoFilmGetCinemasResponse is the response struct for api TaobaoFilmGetCinemas
type TaobaoFilmGetCinemasResponse struct {
	*responses.BaseResponse
	ErrorCode  string        `json:"ErrorCode" xml:"ErrorCode"`
	Msg        string        `json:"Msg" xml:"Msg"`
	SubCode    string        `json:"SubCode" xml:"SubCode"`
	SubMsg     string        `json:"SubMsg" xml:"SubMsg"`
	TotalCount int64         `json:"TotalCount" xml:"TotalCount"`
	LogsId     string        `json:"LogsId" xml:"LogsId"`
	RequestId  string        `json:"RequestId" xml:"RequestId"`
	Cinemas    []CinemasItem `json:"Cinemas" xml:"Cinemas"`
}

// CreateTaobaoFilmGetCinemasRequest creates a request to invoke TaobaoFilmGetCinemas API
func CreateTaobaoFilmGetCinemasRequest() (request *TaobaoFilmGetCinemasRequest) {
	request = &TaobaoFilmGetCinemasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("AppMallsService", "2018-02-24", "TaobaoFilmGetCinemas", "", "")
	return
}

// CreateTaobaoFilmGetCinemasResponse creates a response to parse from TaobaoFilmGetCinemas response
func CreateTaobaoFilmGetCinemasResponse() (response *TaobaoFilmGetCinemasResponse) {
	response = &TaobaoFilmGetCinemasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
