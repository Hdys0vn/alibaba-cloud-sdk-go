package alimt

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

// GetBatchTranslate invokes the alimt.GetBatchTranslate API synchronously
func (client *Client) GetBatchTranslate(request *GetBatchTranslateRequest) (response *GetBatchTranslateResponse, err error) {
	response = CreateGetBatchTranslateResponse()
	err = client.DoAction(request, response)
	return
}

// GetBatchTranslateWithChan invokes the alimt.GetBatchTranslate API asynchronously
func (client *Client) GetBatchTranslateWithChan(request *GetBatchTranslateRequest) (<-chan *GetBatchTranslateResponse, <-chan error) {
	responseChan := make(chan *GetBatchTranslateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBatchTranslate(request)
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

// GetBatchTranslateWithCallback invokes the alimt.GetBatchTranslate API asynchronously
func (client *Client) GetBatchTranslateWithCallback(request *GetBatchTranslateRequest, callback func(response *GetBatchTranslateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBatchTranslateResponse
		var err error
		defer close(result)
		response, err = client.GetBatchTranslate(request)
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

// GetBatchTranslateRequest is the request struct for api GetBatchTranslate
type GetBatchTranslateRequest struct {
	*requests.RpcRequest
	SourceLanguage string `position:"Body" name:"SourceLanguage"`
	SourceText     string `position:"Body" name:"SourceText"`
	FormatType     string `position:"Body" name:"FormatType"`
	ApiType        string `position:"Body" name:"ApiType"`
	Scene          string `position:"Body" name:"Scene"`
	TargetLanguage string `position:"Body" name:"TargetLanguage"`
}

// GetBatchTranslateResponse is the response struct for api GetBatchTranslate
type GetBatchTranslateResponse struct {
	*responses.BaseResponse
	Code           int                      `json:"Code" xml:"Code"`
	Message        string                   `json:"Message" xml:"Message"`
	RequestId      string                   `json:"RequestId" xml:"RequestId"`
	TranslatedList []map[string]interface{} `json:"TranslatedList" xml:"TranslatedList"`
}

// CreateGetBatchTranslateRequest creates a request to invoke GetBatchTranslate API
func CreateGetBatchTranslateRequest() (request *GetBatchTranslateRequest) {
	request = &GetBatchTranslateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "GetBatchTranslate", "", "")
	request.Method = requests.POST
	return
}

// CreateGetBatchTranslateResponse creates a response to parse from GetBatchTranslate response
func CreateGetBatchTranslateResponse() (response *GetBatchTranslateResponse) {
	response = &GetBatchTranslateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
