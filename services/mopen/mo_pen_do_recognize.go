package mopen

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

// MoPenDoRecognize invokes the mopen.MoPenDoRecognize API synchronously
// api document: https://help.aliyun.com/api/mopen/mopendorecognize.html
func (client *Client) MoPenDoRecognize(request *MoPenDoRecognizeRequest) (response *MoPenDoRecognizeResponse, err error) {
	response = CreateMoPenDoRecognizeResponse()
	err = client.DoAction(request, response)
	return
}

// MoPenDoRecognizeWithChan invokes the mopen.MoPenDoRecognize API asynchronously
// api document: https://help.aliyun.com/api/mopen/mopendorecognize.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoPenDoRecognizeWithChan(request *MoPenDoRecognizeRequest) (<-chan *MoPenDoRecognizeResponse, <-chan error) {
	responseChan := make(chan *MoPenDoRecognizeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MoPenDoRecognize(request)
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

// MoPenDoRecognizeWithCallback invokes the mopen.MoPenDoRecognize API asynchronously
// api document: https://help.aliyun.com/api/mopen/mopendorecognize.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoPenDoRecognizeWithCallback(request *MoPenDoRecognizeRequest, callback func(response *MoPenDoRecognizeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MoPenDoRecognizeResponse
		var err error
		defer close(result)
		response, err = client.MoPenDoRecognize(request)
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

// MoPenDoRecognizeRequest is the request struct for api MoPenDoRecognize
type MoPenDoRecognizeRequest struct {
	*requests.RpcRequest
	CanvasId   requests.Integer `position:"Body" name:"CanvasId"`
	EndY       requests.Integer `position:"Body" name:"EndY"`
	EndX       requests.Integer `position:"Body" name:"EndX"`
	JsonConf   string           `position:"Body" name:"JsonConf"`
	ExportType string           `position:"Body" name:"ExportType"`
	StartY     requests.Integer `position:"Body" name:"StartY"`
	StartX     requests.Integer `position:"Body" name:"StartX"`
}

// MoPenDoRecognizeResponse is the response struct for api MoPenDoRecognize
type MoPenDoRecognizeResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Code        string `json:"Code" xml:"Code"`
	Message     string `json:"Message" xml:"Message"`
	Success     bool   `json:"Success" xml:"Success"`
	Description string `json:"Description" xml:"Description"`
	Data        Data   `json:"Data" xml:"Data"`
}

// CreateMoPenDoRecognizeRequest creates a request to invoke MoPenDoRecognize API
func CreateMoPenDoRecognizeRequest() (request *MoPenDoRecognizeRequest) {
	request = &MoPenDoRecognizeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("MoPen", "2018-02-11", "MoPenDoRecognize", "mopen", "openAPI")
	return
}

// CreateMoPenDoRecognizeResponse creates a response to parse from MoPenDoRecognize response
func CreateMoPenDoRecognizeResponse() (response *MoPenDoRecognizeResponse) {
	response = &MoPenDoRecognizeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
