package rtc

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

// AddRecordTemplate invokes the rtc.AddRecordTemplate API synchronously
func (client *Client) AddRecordTemplate(request *AddRecordTemplateRequest) (response *AddRecordTemplateResponse, err error) {
	response = CreateAddRecordTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// AddRecordTemplateWithChan invokes the rtc.AddRecordTemplate API asynchronously
func (client *Client) AddRecordTemplateWithChan(request *AddRecordTemplateRequest) (<-chan *AddRecordTemplateResponse, <-chan error) {
	responseChan := make(chan *AddRecordTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddRecordTemplate(request)
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

// AddRecordTemplateWithCallback invokes the rtc.AddRecordTemplate API asynchronously
func (client *Client) AddRecordTemplateWithCallback(request *AddRecordTemplateRequest, callback func(response *AddRecordTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddRecordTemplateResponse
		var err error
		defer close(result)
		response, err = client.AddRecordTemplate(request)
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

// AddRecordTemplateRequest is the request struct for api AddRecordTemplate
type AddRecordTemplateRequest struct {
	*requests.RpcRequest
	Formats            *[]string                        `position:"Query" name:"Formats"  type:"Repeated"`
	OssFilePrefix      string                           `position:"Query" name:"OssFilePrefix"`
	BackgroundColor    requests.Integer                 `position:"Query" name:"BackgroundColor"`
	TaskProfile        string                           `position:"Query" name:"TaskProfile"`
	LayoutIds          *[]string                        `position:"Query" name:"LayoutIds"  type:"Repeated"`
	ClockWidgets       *[]AddRecordTemplateClockWidgets `position:"Query" name:"ClockWidgets"  type:"Repeated"`
	ShowLog            string                           `position:"Query" name:"ShowLog"`
	OssBucket          string                           `position:"Query" name:"OssBucket"`
	DelayStopTime      requests.Integer                 `position:"Query" name:"DelayStopTime"`
	FileSplitInterval  requests.Integer                 `position:"Query" name:"FileSplitInterval"`
	MnsQueue           string                           `position:"Query" name:"MnsQueue"`
	HttpCallbackUrl    string                           `position:"Query" name:"HttpCallbackUrl"`
	Watermarks         *[]AddRecordTemplateWatermarks   `position:"Query" name:"Watermarks"  type:"Repeated"`
	OwnerId            requests.Integer                 `position:"Query" name:"OwnerId"`
	EnableM3u8DateTime requests.Boolean                 `position:"Query" name:"EnableM3u8DateTime"`
	AppId              string                           `position:"Query" name:"AppId"`
	Backgrounds        *[]AddRecordTemplateBackgrounds  `position:"Query" name:"Backgrounds"  type:"Repeated"`
	Name               string                           `position:"Query" name:"Name"`
	MediaEncode        requests.Integer                 `position:"Query" name:"MediaEncode"`
}

// AddRecordTemplateClockWidgets is a repeated param struct in AddRecordTemplateRequest
type AddRecordTemplateClockWidgets struct {
	FontType  string `name:"FontType"`
	FontColor string `name:"FontColor"`
	Y         string `name:"Y"`
	ZOrder    string `name:"ZOrder"`
	X         string `name:"X"`
	FontSize  string `name:"FontSize"`
}

// AddRecordTemplateWatermarks is a repeated param struct in AddRecordTemplateRequest
type AddRecordTemplateWatermarks struct {
	Alpha   string `name:"Alpha"`
	Width   string `name:"Width"`
	Height  string `name:"Height"`
	Y       string `name:"Y"`
	Url     string `name:"Url"`
	Display string `name:"Display"`
	ZOrder  string `name:"ZOrder"`
	X       string `name:"X"`
}

// AddRecordTemplateBackgrounds is a repeated param struct in AddRecordTemplateRequest
type AddRecordTemplateBackgrounds struct {
	Width   string `name:"Width"`
	Height  string `name:"Height"`
	Y       string `name:"Y"`
	Url     string `name:"Url"`
	Display string `name:"Display"`
	ZOrder  string `name:"ZOrder"`
	X       string `name:"X"`
}

// AddRecordTemplateResponse is the response struct for api AddRecordTemplate
type AddRecordTemplateResponse struct {
	*responses.BaseResponse
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateAddRecordTemplateRequest creates a request to invoke AddRecordTemplate API
func CreateAddRecordTemplateRequest() (request *AddRecordTemplateRequest) {
	request = &AddRecordTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "AddRecordTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateAddRecordTemplateResponse creates a response to parse from AddRecordTemplate response
func CreateAddRecordTemplateResponse() (response *AddRecordTemplateResponse) {
	response = &AddRecordTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
