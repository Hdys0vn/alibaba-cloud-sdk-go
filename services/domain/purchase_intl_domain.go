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

// PurchaseIntlDomain invokes the domain.PurchaseIntlDomain API synchronously
func (client *Client) PurchaseIntlDomain(request *PurchaseIntlDomainRequest) (response *PurchaseIntlDomainResponse, err error) {
	response = CreatePurchaseIntlDomainResponse()
	err = client.DoAction(request, response)
	return
}

// PurchaseIntlDomainWithChan invokes the domain.PurchaseIntlDomain API asynchronously
func (client *Client) PurchaseIntlDomainWithChan(request *PurchaseIntlDomainRequest) (<-chan *PurchaseIntlDomainResponse, <-chan error) {
	responseChan := make(chan *PurchaseIntlDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PurchaseIntlDomain(request)
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

// PurchaseIntlDomainWithCallback invokes the domain.PurchaseIntlDomain API asynchronously
func (client *Client) PurchaseIntlDomainWithCallback(request *PurchaseIntlDomainRequest, callback func(response *PurchaseIntlDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PurchaseIntlDomainResponse
		var err error
		defer close(result)
		response, err = client.PurchaseIntlDomain(request)
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

// PurchaseIntlDomainRequest is the request struct for api PurchaseIntlDomain
type PurchaseIntlDomainRequest struct {
	*requests.RpcRequest
	AuctionId string `position:"Body" name:"AuctionId"`
	Price     string `position:"Body" name:"Price"`
	Currency  string `position:"Body" name:"Currency"`
}

// PurchaseIntlDomainResponse is the response struct for api PurchaseIntlDomain
type PurchaseIntlDomainResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string   `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string   `json:"DynamicMessage" xml:"DynamicMessage"`
	ErrorMsg       string   `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode      string   `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool     `json:"Success" xml:"Success"`
	AllowRetry     bool     `json:"AllowRetry" xml:"AllowRetry"`
	AppName        string   `json:"AppName" xml:"AppName"`
	AuctionId      string   `json:"AuctionId" xml:"AuctionId"`
	ErrorArgs      []string `json:"ErrorArgs" xml:"ErrorArgs"`
}

// CreatePurchaseIntlDomainRequest creates a request to invoke PurchaseIntlDomain API
func CreatePurchaseIntlDomainRequest() (request *PurchaseIntlDomainRequest) {
	request = &PurchaseIntlDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "PurchaseIntlDomain", "", "")
	request.Method = requests.POST
	return
}

// CreatePurchaseIntlDomainResponse creates a response to parse from PurchaseIntlDomain response
func CreatePurchaseIntlDomainResponse() (response *PurchaseIntlDomainResponse) {
	response = &PurchaseIntlDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}