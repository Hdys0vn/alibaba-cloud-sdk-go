package iot

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

// PrintByTemplate invokes the iot.PrintByTemplate API synchronously
func (client *Client) PrintByTemplate(request *PrintByTemplateRequest) (response *PrintByTemplateResponse, err error) {
	response = CreatePrintByTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// PrintByTemplateWithChan invokes the iot.PrintByTemplate API asynchronously
func (client *Client) PrintByTemplateWithChan(request *PrintByTemplateRequest) (<-chan *PrintByTemplateResponse, <-chan error) {
	responseChan := make(chan *PrintByTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PrintByTemplate(request)
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

// PrintByTemplateWithCallback invokes the iot.PrintByTemplate API asynchronously
func (client *Client) PrintByTemplateWithCallback(request *PrintByTemplateRequest, callback func(response *PrintByTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PrintByTemplateResponse
		var err error
		defer close(result)
		response, err = client.PrintByTemplate(request)
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

// PrintByTemplateRequest is the request struct for api PrintByTemplate
type PrintByTemplateRequest struct {
	*requests.RpcRequest
	TemplateBizCode   string           `position:"Body" name:"TemplateBizCode"`
	IotId             string           `position:"Body" name:"IotId"`
	IotInstanceId     string           `position:"Body" name:"IotInstanceId"`
	HistoryPrintTopic requests.Boolean `position:"Body" name:"HistoryPrintTopic"`
	ProductKey        string           `position:"Body" name:"ProductKey"`
	ParamsJsonString  string           `position:"Body" name:"ParamsJsonString"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
	DeviceName        string           `position:"Body" name:"DeviceName"`
}

// PrintByTemplateResponse is the response struct for api PrintByTemplate
type PrintByTemplateResponse struct {
	*responses.BaseResponse
	RequestId    string                `json:"RequestId" xml:"RequestId"`
	Success      bool                  `json:"Success" xml:"Success"`
	Code         string                `json:"Code" xml:"Code"`
	ErrorMessage string                `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInPrintByTemplate `json:"Data" xml:"Data"`
}

// CreatePrintByTemplateRequest creates a request to invoke PrintByTemplate API
func CreatePrintByTemplateRequest() (request *PrintByTemplateRequest) {
	request = &PrintByTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "PrintByTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreatePrintByTemplateResponse creates a response to parse from PrintByTemplate response
func CreatePrintByTemplateResponse() (response *PrintByTemplateResponse) {
	response = &PrintByTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
