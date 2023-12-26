package imm

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

// DecodeBlindWatermark invokes the imm.DecodeBlindWatermark API synchronously
func (client *Client) DecodeBlindWatermark(request *DecodeBlindWatermarkRequest) (response *DecodeBlindWatermarkResponse, err error) {
	response = CreateDecodeBlindWatermarkResponse()
	err = client.DoAction(request, response)
	return
}

// DecodeBlindWatermarkWithChan invokes the imm.DecodeBlindWatermark API asynchronously
func (client *Client) DecodeBlindWatermarkWithChan(request *DecodeBlindWatermarkRequest) (<-chan *DecodeBlindWatermarkResponse, <-chan error) {
	responseChan := make(chan *DecodeBlindWatermarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DecodeBlindWatermark(request)
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

// DecodeBlindWatermarkWithCallback invokes the imm.DecodeBlindWatermark API asynchronously
func (client *Client) DecodeBlindWatermarkWithCallback(request *DecodeBlindWatermarkRequest, callback func(response *DecodeBlindWatermarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DecodeBlindWatermarkResponse
		var err error
		defer close(result)
		response, err = client.DecodeBlindWatermark(request)
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

// DecodeBlindWatermarkRequest is the request struct for api DecodeBlindWatermark
type DecodeBlindWatermarkRequest struct {
	*requests.RpcRequest
	ImageQuality     requests.Integer `position:"Query" name:"ImageQuality"`
	Project          string           `position:"Query" name:"Project"`
	WatermarkType    string           `position:"Query" name:"WatermarkType"`
	TargetUri        string           `position:"Query" name:"TargetUri"`
	Model            string           `position:"Query" name:"Model"`
	ImageUri         string           `position:"Query" name:"ImageUri"`
	OriginalImageUri string           `position:"Query" name:"OriginalImageUri"`
}

// DecodeBlindWatermarkResponse is the response struct for api DecodeBlindWatermark
type DecodeBlindWatermarkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Content   string `json:"Content" xml:"Content"`
	TargetUri string `json:"TargetUri" xml:"TargetUri"`
}

// CreateDecodeBlindWatermarkRequest creates a request to invoke DecodeBlindWatermark API
func CreateDecodeBlindWatermarkRequest() (request *DecodeBlindWatermarkRequest) {
	request = &DecodeBlindWatermarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DecodeBlindWatermark", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDecodeBlindWatermarkResponse creates a response to parse from DecodeBlindWatermark response
func CreateDecodeBlindWatermarkResponse() (response *DecodeBlindWatermarkResponse) {
	response = &DecodeBlindWatermarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
