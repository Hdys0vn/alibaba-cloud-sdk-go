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

// DeleteConfigTemplate invokes the webplus.DeleteConfigTemplate API synchronously
// api document: https://help.aliyun.com/api/webplus/deleteconfigtemplate.html
func (client *Client) DeleteConfigTemplate(request *DeleteConfigTemplateRequest) (response *DeleteConfigTemplateResponse, err error) {
	response = CreateDeleteConfigTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteConfigTemplateWithChan invokes the webplus.DeleteConfigTemplate API asynchronously
// api document: https://help.aliyun.com/api/webplus/deleteconfigtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteConfigTemplateWithChan(request *DeleteConfigTemplateRequest) (<-chan *DeleteConfigTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteConfigTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteConfigTemplate(request)
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

// DeleteConfigTemplateWithCallback invokes the webplus.DeleteConfigTemplate API asynchronously
// api document: https://help.aliyun.com/api/webplus/deleteconfigtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteConfigTemplateWithCallback(request *DeleteConfigTemplateRequest, callback func(response *DeleteConfigTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteConfigTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteConfigTemplate(request)
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

// DeleteConfigTemplateRequest is the request struct for api DeleteConfigTemplate
type DeleteConfigTemplateRequest struct {
	*requests.RoaRequest
	TemplateId string `position:"Query" name:"TemplateId"`
}

// DeleteConfigTemplateResponse is the response struct for api DeleteConfigTemplate
type DeleteConfigTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteConfigTemplateRequest creates a request to invoke DeleteConfigTemplate API
func CreateDeleteConfigTemplateRequest() (request *DeleteConfigTemplateRequest) {
	request = &DeleteConfigTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DeleteConfigTemplate", "/pop/v1/wam/configTemplate", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteConfigTemplateResponse creates a response to parse from DeleteConfigTemplate response
func CreateDeleteConfigTemplateResponse() (response *DeleteConfigTemplateResponse) {
	response = &DeleteConfigTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
