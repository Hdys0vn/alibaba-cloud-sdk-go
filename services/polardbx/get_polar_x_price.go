package polardbx

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

// GetPolarXPrice invokes the polardbx.GetPolarXPrice API synchronously
func (client *Client) GetPolarXPrice(request *GetPolarXPriceRequest) (response *GetPolarXPriceResponse, err error) {
	response = CreateGetPolarXPriceResponse()
	err = client.DoAction(request, response)
	return
}

// GetPolarXPriceWithChan invokes the polardbx.GetPolarXPrice API asynchronously
func (client *Client) GetPolarXPriceWithChan(request *GetPolarXPriceRequest) (<-chan *GetPolarXPriceResponse, <-chan error) {
	responseChan := make(chan *GetPolarXPriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPolarXPrice(request)
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

// GetPolarXPriceWithCallback invokes the polardbx.GetPolarXPrice API asynchronously
func (client *Client) GetPolarXPriceWithCallback(request *GetPolarXPriceRequest, callback func(response *GetPolarXPriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPolarXPriceResponse
		var err error
		defer close(result)
		response, err = client.GetPolarXPrice(request)
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

// GetPolarXPriceRequest is the request struct for api GetPolarXPrice
type GetPolarXPriceRequest struct {
	*requests.RpcRequest
	DBInstanceName string `position:"Query" name:"DBInstanceName"`
	NodeCount      string `position:"Query" name:"NodeCount"`
}

// GetPolarXPriceResponse is the response struct for api GetPolarXPrice
type GetPolarXPriceResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	OrderPriceList []OrderPrice `json:"OrderPriceList" xml:"OrderPriceList"`
}

// CreateGetPolarXPriceRequest creates a request to invoke GetPolarXPrice API
func CreateGetPolarXPriceRequest() (request *GetPolarXPriceRequest) {
	request = &GetPolarXPriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "GetPolarXPrice", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetPolarXPriceResponse creates a response to parse from GetPolarXPrice response
func CreateGetPolarXPriceResponse() (response *GetPolarXPriceResponse) {
	response = &GetPolarXPriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
