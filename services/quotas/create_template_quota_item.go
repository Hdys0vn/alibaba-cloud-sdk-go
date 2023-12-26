package quotas

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

// CreateTemplateQuotaItem invokes the quotas.CreateTemplateQuotaItem API synchronously
func (client *Client) CreateTemplateQuotaItem(request *CreateTemplateQuotaItemRequest) (response *CreateTemplateQuotaItemResponse, err error) {
	response = CreateCreateTemplateQuotaItemResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTemplateQuotaItemWithChan invokes the quotas.CreateTemplateQuotaItem API asynchronously
func (client *Client) CreateTemplateQuotaItemWithChan(request *CreateTemplateQuotaItemRequest) (<-chan *CreateTemplateQuotaItemResponse, <-chan error) {
	responseChan := make(chan *CreateTemplateQuotaItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTemplateQuotaItem(request)
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

// CreateTemplateQuotaItemWithCallback invokes the quotas.CreateTemplateQuotaItem API asynchronously
func (client *Client) CreateTemplateQuotaItemWithCallback(request *CreateTemplateQuotaItemRequest, callback func(response *CreateTemplateQuotaItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTemplateQuotaItemResponse
		var err error
		defer close(result)
		response, err = client.CreateTemplateQuotaItem(request)
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

// CreateTemplateQuotaItemRequest is the request struct for api CreateTemplateQuotaItem
type CreateTemplateQuotaItemRequest struct {
	*requests.RpcRequest
	ProductCode     string                               `position:"Body" name:"ProductCode"`
	QuotaActionCode string                               `position:"Body" name:"QuotaActionCode"`
	DesireValue     requests.Float                       `position:"Body" name:"DesireValue"`
	EffectiveTime   string                               `position:"Body" name:"EffectiveTime"`
	QuotaCategory   string                               `position:"Body" name:"QuotaCategory"`
	OriginalContext string                               `position:"Body" name:"OriginalContext"`
	ExpireTime      string                               `position:"Body" name:"ExpireTime"`
	EnvLanguage     string                               `position:"Body" name:"EnvLanguage"`
	NoticeType      requests.Integer                     `position:"Body" name:"NoticeType"`
	Dimensions      *[]CreateTemplateQuotaItemDimensions `position:"Body" name:"Dimensions"  type:"Repeated"`
}

// CreateTemplateQuotaItemDimensions is a repeated param struct in CreateTemplateQuotaItemRequest
type CreateTemplateQuotaItemDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateTemplateQuotaItemResponse is the response struct for api CreateTemplateQuotaItem
type CreateTemplateQuotaItemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateCreateTemplateQuotaItemRequest creates a request to invoke CreateTemplateQuotaItem API
func CreateCreateTemplateQuotaItemRequest() (request *CreateTemplateQuotaItemRequest) {
	request = &CreateTemplateQuotaItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quotas", "2020-05-10", "CreateTemplateQuotaItem", "quotas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTemplateQuotaItemResponse creates a response to parse from CreateTemplateQuotaItem response
func CreateCreateTemplateQuotaItemResponse() (response *CreateTemplateQuotaItemResponse) {
	response = &CreateTemplateQuotaItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}