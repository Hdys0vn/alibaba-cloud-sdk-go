package dms_enterprise

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

// GetClassificationTemplate invokes the dms_enterprise.GetClassificationTemplate API synchronously
func (client *Client) GetClassificationTemplate(request *GetClassificationTemplateRequest) (response *GetClassificationTemplateResponse, err error) {
	response = CreateGetClassificationTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetClassificationTemplateWithChan invokes the dms_enterprise.GetClassificationTemplate API asynchronously
func (client *Client) GetClassificationTemplateWithChan(request *GetClassificationTemplateRequest) (<-chan *GetClassificationTemplateResponse, <-chan error) {
	responseChan := make(chan *GetClassificationTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetClassificationTemplate(request)
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

// GetClassificationTemplateWithCallback invokes the dms_enterprise.GetClassificationTemplate API asynchronously
func (client *Client) GetClassificationTemplateWithCallback(request *GetClassificationTemplateRequest, callback func(response *GetClassificationTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetClassificationTemplateResponse
		var err error
		defer close(result)
		response, err = client.GetClassificationTemplate(request)
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

// GetClassificationTemplateRequest is the request struct for api GetClassificationTemplate
type GetClassificationTemplateRequest struct {
	*requests.RpcRequest
	Tid        requests.Integer `position:"Query" name:"Tid"`
	InstanceId requests.Integer `position:"Query" name:"InstanceId"`
}

// GetClassificationTemplateResponse is the response struct for api GetClassificationTemplate
type GetClassificationTemplateResponse struct {
	*responses.BaseResponse
	RequestId                         string                            `json:"RequestId" xml:"RequestId"`
	ErrorCode                         string                            `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage                      string                            `json:"ErrorMessage" xml:"ErrorMessage"`
	Success                           bool                              `json:"Success" xml:"Success"`
	ClassificationResourceTemplateMap ClassificationResourceTemplateMap `json:"ClassificationResourceTemplateMap" xml:"ClassificationResourceTemplateMap"`
}

// CreateGetClassificationTemplateRequest creates a request to invoke GetClassificationTemplate API
func CreateGetClassificationTemplateRequest() (request *GetClassificationTemplateRequest) {
	request = &GetClassificationTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetClassificationTemplate", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetClassificationTemplateResponse creates a response to parse from GetClassificationTemplate response
func CreateGetClassificationTemplateResponse() (response *GetClassificationTemplateResponse) {
	response = &GetClassificationTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
