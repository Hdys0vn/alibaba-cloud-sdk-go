package trademark

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

// QueryTrademarkPrice invokes the trademark.QueryTrademarkPrice API synchronously
// api document: https://help.aliyun.com/api/trademark/querytrademarkprice.html
func (client *Client) QueryTrademarkPrice(request *QueryTrademarkPriceRequest) (response *QueryTrademarkPriceResponse, err error) {
	response = CreateQueryTrademarkPriceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTrademarkPriceWithChan invokes the trademark.QueryTrademarkPrice API asynchronously
// api document: https://help.aliyun.com/api/trademark/querytrademarkprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTrademarkPriceWithChan(request *QueryTrademarkPriceRequest) (<-chan *QueryTrademarkPriceResponse, <-chan error) {
	responseChan := make(chan *QueryTrademarkPriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTrademarkPrice(request)
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

// QueryTrademarkPriceWithCallback invokes the trademark.QueryTrademarkPrice API asynchronously
// api document: https://help.aliyun.com/api/trademark/querytrademarkprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTrademarkPriceWithCallback(request *QueryTrademarkPriceRequest, callback func(response *QueryTrademarkPriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTrademarkPriceResponse
		var err error
		defer close(result)
		response, err = client.QueryTrademarkPrice(request)
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

// QueryTrademarkPriceRequest is the request struct for api QueryTrademarkPrice
type QueryTrademarkPriceRequest struct {
	*requests.RpcRequest
	TmName    string           `position:"Query" name:"TmName"`
	TmIcon    string           `position:"Query" name:"TmIcon"`
	OrderData string           `position:"Query" name:"OrderData"`
	Type      requests.Integer `position:"Query" name:"Type"`
	UserId    requests.Integer `position:"Query" name:"UserId"`
}

// QueryTrademarkPriceResponse is the response struct for api QueryTrademarkPrice
type QueryTrademarkPriceResponse struct {
	*responses.BaseResponse
	RequestId     string  `json:"RequestId" xml:"RequestId"`
	DiscountPrice float64 `json:"DiscountPrice" xml:"DiscountPrice"`
	OriginalPrice float64 `json:"OriginalPrice" xml:"OriginalPrice"`
	TradePrice    float64 `json:"TradePrice" xml:"TradePrice"`
	Currency      string  `json:"Currency" xml:"Currency"`
	Prices        Prices  `json:"Prices" xml:"Prices"`
}

// CreateQueryTrademarkPriceRequest creates a request to invoke QueryTrademarkPrice API
func CreateQueryTrademarkPriceRequest() (request *QueryTrademarkPriceRequest) {
	request = &QueryTrademarkPriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "QueryTrademarkPrice", "trademark", "openAPI")
	return
}

// CreateQueryTrademarkPriceResponse creates a response to parse from QueryTrademarkPrice response
func CreateQueryTrademarkPriceResponse() (response *QueryTrademarkPriceResponse) {
	response = &QueryTrademarkPriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
