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

// CopyCompanyTemplateView invokes the cloudesl.CopyCompanyTemplateView API synchronously
func (client *Client) CopyCompanyTemplateView(request *CopyCompanyTemplateViewRequest) (response *CopyCompanyTemplateViewResponse, err error) {
	response = CreateCopyCompanyTemplateViewResponse()
	err = client.DoAction(request, response)
	return
}

// CopyCompanyTemplateViewWithChan invokes the cloudesl.CopyCompanyTemplateView API asynchronously
func (client *Client) CopyCompanyTemplateViewWithChan(request *CopyCompanyTemplateViewRequest) (<-chan *CopyCompanyTemplateViewResponse, <-chan error) {
	responseChan := make(chan *CopyCompanyTemplateViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyCompanyTemplateView(request)
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

// CopyCompanyTemplateViewWithCallback invokes the cloudesl.CopyCompanyTemplateView API asynchronously
func (client *Client) CopyCompanyTemplateViewWithCallback(request *CopyCompanyTemplateViewRequest, callback func(response *CopyCompanyTemplateViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyCompanyTemplateViewResponse
		var err error
		defer close(result)
		response, err = client.CopyCompanyTemplateView(request)
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

// CopyCompanyTemplateViewRequest is the request struct for api CopyCompanyTemplateView
type CopyCompanyTemplateViewRequest struct {
	*requests.RpcRequest
	ExtraParams   string           `position:"Body" name:"ExtraParams"`
	TargetName    string           `position:"Body" name:"TargetName"`
	ModelId       string           `position:"Body" name:"ModelId"`
	TargetVersion string           `position:"Body" name:"TargetVersion"`
	TemplateId    string           `position:"Body" name:"TemplateId"`
	TargetGroupId requests.Integer `position:"Body" name:"TargetGroupId"`
}

// CopyCompanyTemplateViewResponse is the response struct for api CopyCompanyTemplateView
type CopyCompanyTemplateViewResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
}

// CreateCopyCompanyTemplateViewRequest creates a request to invoke CopyCompanyTemplateView API
func CreateCopyCompanyTemplateViewRequest() (request *CopyCompanyTemplateViewRequest) {
	request = &CopyCompanyTemplateViewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "CopyCompanyTemplateView", "", "")
	request.Method = requests.POST
	return
}

// CreateCopyCompanyTemplateViewResponse creates a response to parse from CopyCompanyTemplateView response
func CreateCopyCompanyTemplateViewResponse() (response *CopyCompanyTemplateViewResponse) {
	response = &CopyCompanyTemplateViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
