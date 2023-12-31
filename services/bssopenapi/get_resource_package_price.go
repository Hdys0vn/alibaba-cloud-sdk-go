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

// GetResourcePackagePrice invokes the bssopenapi.GetResourcePackagePrice API synchronously
func (client *Client) GetResourcePackagePrice(request *GetResourcePackagePriceRequest) (response *GetResourcePackagePriceResponse, err error) {
	response = CreateGetResourcePackagePriceResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourcePackagePriceWithChan invokes the bssopenapi.GetResourcePackagePrice API asynchronously
func (client *Client) GetResourcePackagePriceWithChan(request *GetResourcePackagePriceRequest) (<-chan *GetResourcePackagePriceResponse, <-chan error) {
	responseChan := make(chan *GetResourcePackagePriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourcePackagePrice(request)
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

// GetResourcePackagePriceWithCallback invokes the bssopenapi.GetResourcePackagePrice API asynchronously
func (client *Client) GetResourcePackagePriceWithCallback(request *GetResourcePackagePriceRequest, callback func(response *GetResourcePackagePriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourcePackagePriceResponse
		var err error
		defer close(result)
		response, err = client.GetResourcePackagePrice(request)
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

// GetResourcePackagePriceRequest is the request struct for api GetResourcePackagePrice
type GetResourcePackagePriceRequest struct {
	*requests.RpcRequest
	ProductCode   string           `position:"Query" name:"ProductCode"`
	Specification string           `position:"Query" name:"Specification"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	EffectiveDate string           `position:"Query" name:"EffectiveDate"`
	Duration      requests.Integer `position:"Query" name:"Duration"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	PackageType   string           `position:"Query" name:"PackageType"`
	PricingCycle  string           `position:"Query" name:"PricingCycle"`
	OrderType     string           `position:"Query" name:"OrderType"`
}

// GetResourcePackagePriceResponse is the response struct for api GetResourcePackagePrice
type GetResourcePackagePriceResponse struct {
	*responses.BaseResponse
	Code      string                        `json:"Code" xml:"Code"`
	Message   string                        `json:"Message" xml:"Message"`
	RequestId string                        `json:"RequestId" xml:"RequestId"`
	Success   bool                          `json:"Success" xml:"Success"`
	Data      DataInGetResourcePackagePrice `json:"Data" xml:"Data"`
}

// CreateGetResourcePackagePriceRequest creates a request to invoke GetResourcePackagePrice API
func CreateGetResourcePackagePriceRequest() (request *GetResourcePackagePriceRequest) {
	request = &GetResourcePackagePriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "GetResourcePackagePrice", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetResourcePackagePriceResponse creates a response to parse from GetResourcePackagePrice response
func CreateGetResourcePackagePriceResponse() (response *GetResourcePackagePriceResponse) {
	response = &GetResourcePackagePriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
