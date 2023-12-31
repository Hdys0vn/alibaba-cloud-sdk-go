package idrsservice

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

// GetPreSignedUrl invokes the idrsservice.GetPreSignedUrl API synchronously
func (client *Client) GetPreSignedUrl(request *GetPreSignedUrlRequest) (response *GetPreSignedUrlResponse, err error) {
	response = CreateGetPreSignedUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetPreSignedUrlWithChan invokes the idrsservice.GetPreSignedUrl API asynchronously
func (client *Client) GetPreSignedUrlWithChan(request *GetPreSignedUrlRequest) (<-chan *GetPreSignedUrlResponse, <-chan error) {
	responseChan := make(chan *GetPreSignedUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPreSignedUrl(request)
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

// GetPreSignedUrlWithCallback invokes the idrsservice.GetPreSignedUrl API asynchronously
func (client *Client) GetPreSignedUrlWithCallback(request *GetPreSignedUrlRequest, callback func(response *GetPreSignedUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPreSignedUrlResponse
		var err error
		defer close(result)
		response, err = client.GetPreSignedUrl(request)
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

// GetPreSignedUrlRequest is the request struct for api GetPreSignedUrl
type GetPreSignedUrlRequest struct {
	*requests.RpcRequest
	Prefix string `position:"Body" name:"Prefix"`
}

// GetPreSignedUrlResponse is the response struct for api GetPreSignedUrl
type GetPreSignedUrlResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetPreSignedUrlRequest creates a request to invoke GetPreSignedUrl API
func CreateGetPreSignedUrlRequest() (request *GetPreSignedUrlRequest) {
	request = &GetPreSignedUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "GetPreSignedUrl", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetPreSignedUrlResponse creates a response to parse from GetPreSignedUrl response
func CreateGetPreSignedUrlResponse() (response *GetPreSignedUrlResponse) {
	response = &GetPreSignedUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
