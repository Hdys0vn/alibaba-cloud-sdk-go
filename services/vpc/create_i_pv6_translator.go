package vpc

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

// CreateIPv6Translator invokes the vpc.CreateIPv6Translator API synchronously
func (client *Client) CreateIPv6Translator(request *CreateIPv6TranslatorRequest) (response *CreateIPv6TranslatorResponse, err error) {
	response = CreateCreateIPv6TranslatorResponse()
	err = client.DoAction(request, response)
	return
}

// CreateIPv6TranslatorWithChan invokes the vpc.CreateIPv6Translator API asynchronously
func (client *Client) CreateIPv6TranslatorWithChan(request *CreateIPv6TranslatorRequest) (<-chan *CreateIPv6TranslatorResponse, <-chan error) {
	responseChan := make(chan *CreateIPv6TranslatorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateIPv6Translator(request)
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

// CreateIPv6TranslatorWithCallback invokes the vpc.CreateIPv6Translator API asynchronously
func (client *Client) CreateIPv6TranslatorWithCallback(request *CreateIPv6TranslatorRequest, callback func(response *CreateIPv6TranslatorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateIPv6TranslatorResponse
		var err error
		defer close(result)
		response, err = client.CreateIPv6Translator(request)
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

// CreateIPv6TranslatorRequest is the request struct for api CreateIPv6Translator
type CreateIPv6TranslatorRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Spec                 string           `position:"Query" name:"Spec"`
	Duration             requests.Integer `position:"Query" name:"Duration"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            requests.Integer `position:"Query" name:"Bandwidth"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
	PayType              string           `position:"Query" name:"PayType"`
	PricingCycle         string           `position:"Query" name:"PricingCycle"`
}

// CreateIPv6TranslatorResponse is the response struct for api CreateIPv6Translator
type CreateIPv6TranslatorResponse struct {
	*responses.BaseResponse
	Ipv6TranslatorId string `json:"Ipv6TranslatorId" xml:"Ipv6TranslatorId"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Spec             string `json:"Spec" xml:"Spec"`
	Name             string `json:"Name" xml:"Name"`
	OrderId          int64  `json:"OrderId" xml:"OrderId"`
}

// CreateCreateIPv6TranslatorRequest creates a request to invoke CreateIPv6Translator API
func CreateCreateIPv6TranslatorRequest() (request *CreateIPv6TranslatorRequest) {
	request = &CreateIPv6TranslatorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateIPv6Translator", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateIPv6TranslatorResponse creates a response to parse from CreateIPv6Translator response
func CreateCreateIPv6TranslatorResponse() (response *CreateIPv6TranslatorResponse) {
	response = &CreateIPv6TranslatorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
