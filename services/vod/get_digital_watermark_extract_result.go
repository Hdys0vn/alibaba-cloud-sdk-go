package vod

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

// GetDigitalWatermarkExtractResult invokes the vod.GetDigitalWatermarkExtractResult API synchronously
func (client *Client) GetDigitalWatermarkExtractResult(request *GetDigitalWatermarkExtractResultRequest) (response *GetDigitalWatermarkExtractResultResponse, err error) {
	response = CreateGetDigitalWatermarkExtractResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetDigitalWatermarkExtractResultWithChan invokes the vod.GetDigitalWatermarkExtractResult API asynchronously
func (client *Client) GetDigitalWatermarkExtractResultWithChan(request *GetDigitalWatermarkExtractResultRequest) (<-chan *GetDigitalWatermarkExtractResultResponse, <-chan error) {
	responseChan := make(chan *GetDigitalWatermarkExtractResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDigitalWatermarkExtractResult(request)
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

// GetDigitalWatermarkExtractResultWithCallback invokes the vod.GetDigitalWatermarkExtractResult API asynchronously
func (client *Client) GetDigitalWatermarkExtractResultWithCallback(request *GetDigitalWatermarkExtractResultRequest, callback func(response *GetDigitalWatermarkExtractResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDigitalWatermarkExtractResultResponse
		var err error
		defer close(result)
		response, err = client.GetDigitalWatermarkExtractResult(request)
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

// GetDigitalWatermarkExtractResultRequest is the request struct for api GetDigitalWatermarkExtractResult
type GetDigitalWatermarkExtractResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ExtractType          string `position:"Query" name:"ExtractType"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

// GetDigitalWatermarkExtractResultResponse is the response struct for api GetDigitalWatermarkExtractResult
type GetDigitalWatermarkExtractResultResponse struct {
	*responses.BaseResponse
	RequestId           string            `json:"RequestId" xml:"RequestId"`
	AiExtractResultList []AiExtractResult `json:"AiExtractResultList" xml:"AiExtractResultList"`
}

// CreateGetDigitalWatermarkExtractResultRequest creates a request to invoke GetDigitalWatermarkExtractResult API
func CreateGetDigitalWatermarkExtractResultRequest() (request *GetDigitalWatermarkExtractResultRequest) {
	request = &GetDigitalWatermarkExtractResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetDigitalWatermarkExtractResult", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDigitalWatermarkExtractResultResponse creates a response to parse from GetDigitalWatermarkExtractResult response
func CreateGetDigitalWatermarkExtractResultResponse() (response *GetDigitalWatermarkExtractResultResponse) {
	response = &GetDigitalWatermarkExtractResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
