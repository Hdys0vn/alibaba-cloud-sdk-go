package appstream_center

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

// GetResourcePrice invokes the appstream_center.GetResourcePrice API synchronously
func (client *Client) GetResourcePrice(request *GetResourcePriceRequest) (response *GetResourcePriceResponse, err error) {
	response = CreateGetResourcePriceResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourcePriceWithChan invokes the appstream_center.GetResourcePrice API asynchronously
func (client *Client) GetResourcePriceWithChan(request *GetResourcePriceRequest) (<-chan *GetResourcePriceResponse, <-chan error) {
	responseChan := make(chan *GetResourcePriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourcePrice(request)
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

// GetResourcePriceWithCallback invokes the appstream_center.GetResourcePrice API asynchronously
func (client *Client) GetResourcePriceWithCallback(request *GetResourcePriceRequest, callback func(response *GetResourcePriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourcePriceResponse
		var err error
		defer close(result)
		response, err = client.GetResourcePrice(request)
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

// GetResourcePriceRequest is the request struct for api GetResourcePrice
type GetResourcePriceRequest struct {
	*requests.RpcRequest
	BizRegionId      string           `position:"Query" name:"BizRegionId"`
	Period           requests.Integer `position:"Query" name:"Period"`
	Amount           requests.Integer `position:"Query" name:"Amount"`
	NodeInstanceType string           `position:"Query" name:"NodeInstanceType"`
	ProductType      string           `position:"Query" name:"ProductType"`
	PeriodUnit       string           `position:"Query" name:"PeriodUnit"`
	ChargeType       string           `position:"Query" name:"ChargeType"`
}

// GetResourcePriceResponse is the response struct for api GetResourcePrice
type GetResourcePriceResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Code       string     `json:"Code" xml:"Code"`
	Message    string     `json:"Message" xml:"Message"`
	PriceModel PriceModel `json:"PriceModel" xml:"PriceModel"`
}

// CreateGetResourcePriceRequest creates a request to invoke GetResourcePrice API
func CreateGetResourcePriceRequest() (request *GetResourcePriceRequest) {
	request = &GetResourcePriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("appstream-center", "2021-09-01", "GetResourcePrice", "", "")
	request.Method = requests.POST
	return
}

// CreateGetResourcePriceResponse creates a response to parse from GetResourcePrice response
func CreateGetResourcePriceResponse() (response *GetResourcePriceResponse) {
	response = &GetResourcePriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
