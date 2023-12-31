package webplus

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

// DescribeConfigTemplates invokes the webplus.DescribeConfigTemplates API synchronously
// api document: https://help.aliyun.com/api/webplus/describeconfigtemplates.html
func (client *Client) DescribeConfigTemplates(request *DescribeConfigTemplatesRequest) (response *DescribeConfigTemplatesResponse, err error) {
	response = CreateDescribeConfigTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConfigTemplatesWithChan invokes the webplus.DescribeConfigTemplates API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeconfigtemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConfigTemplatesWithChan(request *DescribeConfigTemplatesRequest) (<-chan *DescribeConfigTemplatesResponse, <-chan error) {
	responseChan := make(chan *DescribeConfigTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConfigTemplates(request)
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

// DescribeConfigTemplatesWithCallback invokes the webplus.DescribeConfigTemplates API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeconfigtemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConfigTemplatesWithCallback(request *DescribeConfigTemplatesRequest, callback func(response *DescribeConfigTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConfigTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeConfigTemplates(request)
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

// DescribeConfigTemplatesRequest is the request struct for api DescribeConfigTemplates
type DescribeConfigTemplatesRequest struct {
	*requests.RoaRequest
	TemplateSearch string           `position:"Query" name:"TemplateSearch"`
	AppId          string           `position:"Query" name:"AppId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	TemplateName   string           `position:"Query" name:"TemplateName"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeConfigTemplatesResponse is the response struct for api DescribeConfigTemplates
type DescribeConfigTemplatesResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Code            string          `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	PageNumber      int             `json:"PageNumber" xml:"PageNumber"`
	PageSize        int             `json:"PageSize" xml:"PageSize"`
	TotalCount      int             `json:"TotalCount" xml:"TotalCount"`
	ConfigTemplates ConfigTemplates `json:"ConfigTemplates" xml:"ConfigTemplates"`
}

// CreateDescribeConfigTemplatesRequest creates a request to invoke DescribeConfigTemplates API
func CreateDescribeConfigTemplatesRequest() (request *DescribeConfigTemplatesRequest) {
	request = &DescribeConfigTemplatesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DescribeConfigTemplates", "/pop/v1/wam/configTemplate", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeConfigTemplatesResponse creates a response to parse from DescribeConfigTemplates response
func CreateDescribeConfigTemplatesResponse() (response *DescribeConfigTemplatesResponse) {
	response = &DescribeConfigTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
