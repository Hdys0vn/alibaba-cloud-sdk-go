package cloudesl

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

// DescribeCompanyTemplateAttribute invokes the cloudesl.DescribeCompanyTemplateAttribute API synchronously
func (client *Client) DescribeCompanyTemplateAttribute(request *DescribeCompanyTemplateAttributeRequest) (response *DescribeCompanyTemplateAttributeResponse, err error) {
	response = CreateDescribeCompanyTemplateAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCompanyTemplateAttributeWithChan invokes the cloudesl.DescribeCompanyTemplateAttribute API asynchronously
func (client *Client) DescribeCompanyTemplateAttributeWithChan(request *DescribeCompanyTemplateAttributeRequest) (<-chan *DescribeCompanyTemplateAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeCompanyTemplateAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCompanyTemplateAttribute(request)
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

// DescribeCompanyTemplateAttributeWithCallback invokes the cloudesl.DescribeCompanyTemplateAttribute API asynchronously
func (client *Client) DescribeCompanyTemplateAttributeWithCallback(request *DescribeCompanyTemplateAttributeRequest, callback func(response *DescribeCompanyTemplateAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCompanyTemplateAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeCompanyTemplateAttribute(request)
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

// DescribeCompanyTemplateAttributeRequest is the request struct for api DescribeCompanyTemplateAttribute
type DescribeCompanyTemplateAttributeRequest struct {
	*requests.RpcRequest
	ExtraParams string `position:"Body" name:"ExtraParams"`
}

// DescribeCompanyTemplateAttributeResponse is the response struct for api DescribeCompanyTemplateAttribute
type DescribeCompanyTemplateAttributeResponse struct {
	*responses.BaseResponse
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string        `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool          `json:"Success" xml:"Success"`
	ErrorCode      string        `json:"ErrorCode" xml:"ErrorCode"`
	Code           string        `json:"Code" xml:"Code"`
	Message        string        `json:"Message" xml:"Message"`
	DynamicMessage string        `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string        `json:"DynamicCode" xml:"DynamicCode"`
	CategoryField  CategoryField `json:"CategoryField" xml:"CategoryField"`
	FontType       FontType      `json:"FontType" xml:"FontType"`
	DeviceType     DeviceType    `json:"DeviceType" xml:"DeviceType"`
	TemplateType   TemplateType  `json:"TemplateType" xml:"TemplateType"`
	SizeType       SizeType      `json:"SizeType" xml:"SizeType"`
}

// CreateDescribeCompanyTemplateAttributeRequest creates a request to invoke DescribeCompanyTemplateAttribute API
func CreateDescribeCompanyTemplateAttributeRequest() (request *DescribeCompanyTemplateAttributeRequest) {
	request = &DescribeCompanyTemplateAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribeCompanyTemplateAttribute", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCompanyTemplateAttributeResponse creates a response to parse from DescribeCompanyTemplateAttribute response
func CreateDescribeCompanyTemplateAttributeResponse() (response *DescribeCompanyTemplateAttributeResponse) {
	response = &DescribeCompanyTemplateAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
