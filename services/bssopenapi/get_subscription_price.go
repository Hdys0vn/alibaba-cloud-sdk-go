package bssopenapi

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

// GetSubscriptionPrice invokes the bssopenapi.GetSubscriptionPrice API synchronously
func (client *Client) GetSubscriptionPrice(request *GetSubscriptionPriceRequest) (response *GetSubscriptionPriceResponse, err error) {
	response = CreateGetSubscriptionPriceResponse()
	err = client.DoAction(request, response)
	return
}

// GetSubscriptionPriceWithChan invokes the bssopenapi.GetSubscriptionPrice API asynchronously
func (client *Client) GetSubscriptionPriceWithChan(request *GetSubscriptionPriceRequest) (<-chan *GetSubscriptionPriceResponse, <-chan error) {
	responseChan := make(chan *GetSubscriptionPriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSubscriptionPrice(request)
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

// GetSubscriptionPriceWithCallback invokes the bssopenapi.GetSubscriptionPrice API asynchronously
func (client *Client) GetSubscriptionPriceWithCallback(request *GetSubscriptionPriceRequest, callback func(response *GetSubscriptionPriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSubscriptionPriceResponse
		var err error
		defer close(result)
		response, err = client.GetSubscriptionPrice(request)
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

// GetSubscriptionPriceRequest is the request struct for api GetSubscriptionPrice
type GetSubscriptionPriceRequest struct {
	*requests.RpcRequest
	ProductCode           string                            `position:"Query" name:"ProductCode"`
	Quantity              requests.Integer                  `position:"Query" name:"Quantity"`
	SubscriptionType      string                            `position:"Query" name:"SubscriptionType"`
	ModuleList            *[]GetSubscriptionPriceModuleList `position:"Query" name:"ModuleList"  type:"Repeated"`
	OwnerId               requests.Integer                  `position:"Query" name:"OwnerId"`
	ProductType           string                            `position:"Query" name:"ProductType"`
	ServicePeriodQuantity requests.Integer                  `position:"Query" name:"ServicePeriodQuantity"`
	InstanceId            string                            `position:"Query" name:"InstanceId"`
	ServicePeriodUnit     string                            `position:"Query" name:"ServicePeriodUnit"`
	Region                string                            `position:"Query" name:"Region"`
	OrderType             string                            `position:"Query" name:"OrderType"`
}

// GetSubscriptionPriceModuleList is a repeated param struct in GetSubscriptionPriceRequest
type GetSubscriptionPriceModuleList struct {
	ModuleCode   string `name:"ModuleCode"`
	ModuleStatus string `name:"ModuleStatus"`
	Tag          string `name:"Tag"`
	Config       string `name:"Config"`
}

// GetSubscriptionPriceResponse is the response struct for api GetSubscriptionPrice
type GetSubscriptionPriceResponse struct {
	*responses.BaseResponse
	Code      string                     `json:"Code" xml:"Code"`
	Message   string                     `json:"Message" xml:"Message"`
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Success   bool                       `json:"Success" xml:"Success"`
	Data      DataInGetSubscriptionPrice `json:"Data" xml:"Data"`
}

// CreateGetSubscriptionPriceRequest creates a request to invoke GetSubscriptionPrice API
func CreateGetSubscriptionPriceRequest() (request *GetSubscriptionPriceRequest) {
	request = &GetSubscriptionPriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "GetSubscriptionPrice", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSubscriptionPriceResponse creates a response to parse from GetSubscriptionPrice response
func CreateGetSubscriptionPriceResponse() (response *GetSubscriptionPriceResponse) {
	response = &GetSubscriptionPriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
