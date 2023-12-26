package iot

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

// ReleaseProduct invokes the iot.ReleaseProduct API synchronously
func (client *Client) ReleaseProduct(request *ReleaseProductRequest) (response *ReleaseProductResponse, err error) {
	response = CreateReleaseProductResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseProductWithChan invokes the iot.ReleaseProduct API asynchronously
func (client *Client) ReleaseProductWithChan(request *ReleaseProductRequest) (<-chan *ReleaseProductResponse, <-chan error) {
	responseChan := make(chan *ReleaseProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseProduct(request)
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

// ReleaseProductWithCallback invokes the iot.ReleaseProduct API asynchronously
func (client *Client) ReleaseProductWithCallback(request *ReleaseProductRequest, callback func(response *ReleaseProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseProductResponse
		var err error
		defer close(result)
		response, err = client.ReleaseProduct(request)
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

// ReleaseProductRequest is the request struct for api ReleaseProduct
type ReleaseProductRequest struct {
	*requests.RpcRequest
	Template           requests.Integer `position:"Query" name:"Template"`
	RealTenantId       string           `position:"Query" name:"RealTenantId"`
	CategoryKey        string           `position:"Query" name:"CategoryKey"`
	RealTripartiteKey  string           `position:"Query" name:"RealTripartiteKey"`
	TemplateIdentifier string           `position:"Query" name:"TemplateIdentifier"`
	IotInstanceId      string           `position:"Query" name:"IotInstanceId"`
	TemplateName       string           `position:"Query" name:"TemplateName"`
	ProductKey         string           `position:"Query" name:"ProductKey"`
	ApiProduct         string           `position:"Body" name:"ApiProduct"`
	ApiRevision        string           `position:"Body" name:"ApiRevision"`
	CategoryName       string           `position:"Query" name:"CategoryName"`
	BizTenantId        string           `position:"Query" name:"BizTenantId"`
}

// ReleaseProductResponse is the response struct for api ReleaseProduct
type ReleaseProductResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
}

// CreateReleaseProductRequest creates a request to invoke ReleaseProduct API
func CreateReleaseProductRequest() (request *ReleaseProductRequest) {
	request = &ReleaseProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ReleaseProduct", "", "")
	request.Method = requests.POST
	return
}

// CreateReleaseProductResponse creates a response to parse from ReleaseProduct response
func CreateReleaseProductResponse() (response *ReleaseProductResponse) {
	response = &ReleaseProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
