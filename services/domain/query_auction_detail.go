package domain

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

// QueryAuctionDetail invokes the domain.QueryAuctionDetail API synchronously
func (client *Client) QueryAuctionDetail(request *QueryAuctionDetailRequest) (response *QueryAuctionDetailResponse, err error) {
	response = CreateQueryAuctionDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAuctionDetailWithChan invokes the domain.QueryAuctionDetail API asynchronously
func (client *Client) QueryAuctionDetailWithChan(request *QueryAuctionDetailRequest) (<-chan *QueryAuctionDetailResponse, <-chan error) {
	responseChan := make(chan *QueryAuctionDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAuctionDetail(request)
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

// QueryAuctionDetailWithCallback invokes the domain.QueryAuctionDetail API asynchronously
func (client *Client) QueryAuctionDetailWithCallback(request *QueryAuctionDetailRequest, callback func(response *QueryAuctionDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAuctionDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryAuctionDetail(request)
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

// QueryAuctionDetailRequest is the request struct for api QueryAuctionDetail
type QueryAuctionDetailRequest struct {
	*requests.RpcRequest
	AuctionId string `position:"Body" name:"AuctionId"`
}

// QueryAuctionDetailResponse is the response struct for api QueryAuctionDetail
type QueryAuctionDetailResponse struct {
	*responses.BaseResponse
	Status          string  `json:"Status" xml:"Status"`
	ReserveMet      bool    `json:"ReserveMet" xml:"ReserveMet"`
	HighBid         float64 `json:"HighBid" xml:"HighBid"`
	DeliveryTime    int64   `json:"DeliveryTime" xml:"DeliveryTime"`
	TransferInPrice float64 `json:"TransferInPrice" xml:"TransferInPrice"`
	NextValidBid    float64 `json:"NextValidBid" xml:"NextValidBid"`
	ProduceStatus   string  `json:"ProduceStatus" xml:"ProduceStatus"`
	HighBidder      string  `json:"HighBidder" xml:"HighBidder"`
	BookedPartner   string  `json:"BookedPartner" xml:"BookedPartner"`
	Currency        string  `json:"Currency" xml:"Currency"`
	DomainName      string  `json:"DomainName" xml:"DomainName"`
	YourCurrentBid  float64 `json:"YourCurrentBid" xml:"YourCurrentBid"`
	FailCode        string  `json:"FailCode" xml:"FailCode"`
	AuctionEndTime  int64   `json:"AuctionEndTime" xml:"AuctionEndTime"`
	BookEndTime     int64   `json:"BookEndTime" xml:"BookEndTime"`
	PayEndTime      int64   `json:"PayEndTime" xml:"PayEndTime"`
	PayStatus       string  `json:"PayStatus" xml:"PayStatus"`
	RequestId       string  `json:"RequestId" xml:"RequestId"`
	YourMaxBid      float64 `json:"YourMaxBid" xml:"YourMaxBid"`
	PayPrice        float64 `json:"PayPrice" xml:"PayPrice"`
	AuctionId       string  `json:"AuctionId" xml:"AuctionId"`
	PartnerType     string  `json:"PartnerType" xml:"PartnerType"`
	DomainType      string  `json:"DomainType" xml:"DomainType"`
	ReservePrice    float64 `json:"ReservePrice" xml:"ReservePrice"`
}

// CreateQueryAuctionDetailRequest creates a request to invoke QueryAuctionDetail API
func CreateQueryAuctionDetailRequest() (request *QueryAuctionDetailRequest) {
	request = &QueryAuctionDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "QueryAuctionDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryAuctionDetailResponse creates a response to parse from QueryAuctionDetail response
func CreateQueryAuctionDetailResponse() (response *QueryAuctionDetailResponse) {
	response = &QueryAuctionDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
