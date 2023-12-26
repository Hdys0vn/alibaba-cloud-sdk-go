package domain_intl

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

// SaveSingleTaskForCreatingOrderTransfer invokes the domain_intl.SaveSingleTaskForCreatingOrderTransfer API synchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskforcreatingordertransfer.html
func (client *Client) SaveSingleTaskForCreatingOrderTransfer(request *SaveSingleTaskForCreatingOrderTransferRequest) (response *SaveSingleTaskForCreatingOrderTransferResponse, err error) {
	response = CreateSaveSingleTaskForCreatingOrderTransferResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForCreatingOrderTransferWithChan invokes the domain_intl.SaveSingleTaskForCreatingOrderTransfer API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskforcreatingordertransfer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForCreatingOrderTransferWithChan(request *SaveSingleTaskForCreatingOrderTransferRequest) (<-chan *SaveSingleTaskForCreatingOrderTransferResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForCreatingOrderTransferResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForCreatingOrderTransfer(request)
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

// SaveSingleTaskForCreatingOrderTransferWithCallback invokes the domain_intl.SaveSingleTaskForCreatingOrderTransfer API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskforcreatingordertransfer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForCreatingOrderTransferWithCallback(request *SaveSingleTaskForCreatingOrderTransferRequest, callback func(response *SaveSingleTaskForCreatingOrderTransferResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForCreatingOrderTransferResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForCreatingOrderTransfer(request)
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

// SaveSingleTaskForCreatingOrderTransferRequest is the request struct for api SaveSingleTaskForCreatingOrderTransfer
type SaveSingleTaskForCreatingOrderTransferRequest struct {
	*requests.RpcRequest
	PermitPremiumTransfer requests.Boolean `position:"Query" name:"PermitPremiumTransfer"`
	PromotionNo           string           `position:"Query" name:"PromotionNo"`
	AuthorizationCode     string           `position:"Query" name:"AuthorizationCode"`
	UserClientIp          string           `position:"Query" name:"UserClientIp"`
	DomainName            string           `position:"Query" name:"DomainName"`
	RegistrantProfileId   requests.Integer `position:"Query" name:"RegistrantProfileId"`
	CouponNo              string           `position:"Query" name:"CouponNo"`
	UseCoupon             requests.Boolean `position:"Query" name:"UseCoupon"`
	Lang                  string           `position:"Query" name:"Lang"`
	UsePromotion          requests.Boolean `position:"Query" name:"UsePromotion"`
}

// SaveSingleTaskForCreatingOrderTransferResponse is the response struct for api SaveSingleTaskForCreatingOrderTransfer
type SaveSingleTaskForCreatingOrderTransferResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForCreatingOrderTransferRequest creates a request to invoke SaveSingleTaskForCreatingOrderTransfer API
func CreateSaveSingleTaskForCreatingOrderTransferRequest() (request *SaveSingleTaskForCreatingOrderTransferRequest) {
	request = &SaveSingleTaskForCreatingOrderTransferRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForCreatingOrderTransfer", "domain", "openAPI")
	return
}

// CreateSaveSingleTaskForCreatingOrderTransferResponse creates a response to parse from SaveSingleTaskForCreatingOrderTransfer response
func CreateSaveSingleTaskForCreatingOrderTransferResponse() (response *SaveSingleTaskForCreatingOrderTransferResponse) {
	response = &SaveSingleTaskForCreatingOrderTransferResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
