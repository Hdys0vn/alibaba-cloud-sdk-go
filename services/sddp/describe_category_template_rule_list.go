package sddp

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

// DescribeCategoryTemplateRuleList invokes the sddp.DescribeCategoryTemplateRuleList API synchronously
func (client *Client) DescribeCategoryTemplateRuleList(request *DescribeCategoryTemplateRuleListRequest) (response *DescribeCategoryTemplateRuleListResponse, err error) {
	response = CreateDescribeCategoryTemplateRuleListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCategoryTemplateRuleListWithChan invokes the sddp.DescribeCategoryTemplateRuleList API asynchronously
func (client *Client) DescribeCategoryTemplateRuleListWithChan(request *DescribeCategoryTemplateRuleListRequest) (<-chan *DescribeCategoryTemplateRuleListResponse, <-chan error) {
	responseChan := make(chan *DescribeCategoryTemplateRuleListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCategoryTemplateRuleList(request)
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

// DescribeCategoryTemplateRuleListWithCallback invokes the sddp.DescribeCategoryTemplateRuleList API asynchronously
func (client *Client) DescribeCategoryTemplateRuleListWithCallback(request *DescribeCategoryTemplateRuleListRequest, callback func(response *DescribeCategoryTemplateRuleListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCategoryTemplateRuleListResponse
		var err error
		defer close(result)
		response, err = client.DescribeCategoryTemplateRuleList(request)
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

// DescribeCategoryTemplateRuleListRequest is the request struct for api DescribeCategoryTemplateRuleList
type DescribeCategoryTemplateRuleListRequest struct {
	*requests.RpcRequest
	RiskLevelId          requests.Integer `position:"Query" name:"RiskLevelId"`
	SourceIp             string           `position:"Query" name:"SourceIp"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	Lang                 string           `position:"Query" name:"Lang"`
	ParentCategoryIdList string           `position:"Query" name:"ParentCategoryIdList"`
	CurrentPage          requests.Integer `position:"Query" name:"CurrentPage"`
	TemplateId           requests.Integer `position:"Query" name:"TemplateId"`
	CustomType           requests.Integer `position:"Query" name:"CustomType"`
	Status               requests.Integer `position:"Query" name:"Status"`
}

// DescribeCategoryTemplateRuleListResponse is the response struct for api DescribeCategoryTemplateRuleList
type DescribeCategoryTemplateRuleListResponse struct {
	*responses.BaseResponse
	CurrentPage int         `json:"CurrentPage" xml:"CurrentPage"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	Items       []DataLimit `json:"Items" xml:"Items"`
}

// CreateDescribeCategoryTemplateRuleListRequest creates a request to invoke DescribeCategoryTemplateRuleList API
func CreateDescribeCategoryTemplateRuleListRequest() (request *DescribeCategoryTemplateRuleListRequest) {
	request = &DescribeCategoryTemplateRuleListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeCategoryTemplateRuleList", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCategoryTemplateRuleListResponse creates a response to parse from DescribeCategoryTemplateRuleList response
func CreateDescribeCategoryTemplateRuleListResponse() (response *DescribeCategoryTemplateRuleListResponse) {
	response = &DescribeCategoryTemplateRuleListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}