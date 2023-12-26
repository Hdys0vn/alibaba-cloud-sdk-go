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

// TranslateImage invokes the alimt.TranslateImage API synchronously
func (client *Client) TranslateImage(request *TranslateImageRequest) (response *TranslateImageResponse, err error) {
	response = CreateTranslateImageResponse()
	err = client.DoAction(request, response)
	return
}

// TranslateImageWithChan invokes the alimt.TranslateImage API asynchronously
func (client *Client) TranslateImageWithChan(request *TranslateImageRequest) (<-chan *TranslateImageResponse, <-chan error) {
	responseChan := make(chan *TranslateImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TranslateImage(request)
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

// TranslateImageWithCallback invokes the alimt.TranslateImage API asynchronously
func (client *Client) TranslateImageWithCallback(request *TranslateImageRequest, callback func(response *TranslateImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TranslateImageResponse
		var err error
		defer close(result)
		response, err = client.TranslateImage(request)
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

// TranslateImageRequest is the request struct for api TranslateImage
type TranslateImageRequest struct {
	*requests.RpcRequest
	Ext            string `position:"Body" name:"Ext"`
	SourceLanguage string `position:"Body" name:"SourceLanguage"`
	Field          string `position:"Body" name:"Field"`
	ImageUrl       string `position:"Body" name:"ImageUrl"`
	TargetLanguage string `position:"Body" name:"TargetLanguage"`
	ImageBase64    string `position:"Body" name:"ImageBase64"`
}

// TranslateImageResponse is the response struct for api TranslateImage
type TranslateImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateTranslateImageRequest creates a request to invoke TranslateImage API
func CreateTranslateImageRequest() (request *TranslateImageRequest) {
	request = &TranslateImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "TranslateImage", "", "")
	request.Method = requests.POST
	return
}

// CreateTranslateImageResponse creates a response to parse from TranslateImage response
func CreateTranslateImageResponse() (response *TranslateImageResponse) {
	response = &TranslateImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
