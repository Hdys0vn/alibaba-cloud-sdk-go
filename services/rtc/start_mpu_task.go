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

// StartMPUTask invokes the rtc.StartMPUTask API synchronously
func (client *Client) StartMPUTask(request *StartMPUTaskRequest) (response *StartMPUTaskResponse, err error) {
	response = CreateStartMPUTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StartMPUTaskWithChan invokes the rtc.StartMPUTask API asynchronously
func (client *Client) StartMPUTaskWithChan(request *StartMPUTaskRequest) (<-chan *StartMPUTaskResponse, <-chan error) {
	responseChan := make(chan *StartMPUTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartMPUTask(request)
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

// StartMPUTaskWithCallback invokes the rtc.StartMPUTask API asynchronously
func (client *Client) StartMPUTaskWithCallback(request *StartMPUTaskRequest, callback func(response *StartMPUTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartMPUTaskResponse
		var err error
		defer close(result)
		response, err = client.StartMPUTask(request)
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

// StartMPUTaskRequest is the request struct for api StartMPUTask
type StartMPUTaskRequest struct {
	*requests.RpcRequest
	PayloadType               requests.Integer                  `position:"Query" name:"PayloadType"`
	UserPanes                 *[]StartMPUTaskUserPanes          `position:"Query" name:"UserPanes"  type:"Repeated"`
	BackgroundColor           requests.Integer                  `position:"Query" name:"BackgroundColor"`
	ReportVad                 requests.Integer                  `position:"Query" name:"ReportVad"`
	SourceType                string                            `position:"Query" name:"SourceType"`
	TaskId                    string                            `position:"Query" name:"TaskId"`
	ClockWidgets              *[]StartMPUTaskClockWidgets       `position:"Query" name:"ClockWidgets"  type:"Repeated"`
	ShowLog                   string                            `position:"Query" name:"ShowLog"`
	UnsubSpecCameraUsers      *[]string                         `position:"Query" name:"UnsubSpecCameraUsers"  type:"Repeated"`
	TaskType                  requests.Integer                  `position:"Query" name:"TaskType"`
	UnsubSpecAudioUsers       *[]string                         `position:"Query" name:"UnsubSpecAudioUsers"  type:"Repeated"`
	VadInterval               requests.Integer                  `position:"Query" name:"VadInterval"`
	Watermarks                *[]StartMPUTaskWatermarks         `position:"Query" name:"Watermarks"  type:"Repeated"`
	OwnerId                   requests.Integer                  `position:"Query" name:"OwnerId"`
	SubSpecAudioUsers         *[]string                         `position:"Query" name:"SubSpecAudioUsers"  type:"Repeated"`
	MediaEncode               requests.Integer                  `position:"Query" name:"MediaEncode"`
	EnhancedParam             StartMPUTaskEnhancedParam         `position:"Body" name:"EnhancedParam"  type:"Struct"`
	RtpExtInfo                requests.Integer                  `position:"Query" name:"RtpExtInfo"`
	CropMode                  requests.Integer                  `position:"Query" name:"CropMode"`
	SubSpecCameraUsers        *[]string                         `position:"Query" name:"SubSpecCameraUsers"  type:"Repeated"`
	OutputStreamParams        *[]StartMPUTaskOutputStreamParams `position:"Query" name:"OutputStreamParams"  type:"Repeated"`
	TaskProfile               string                            `position:"Query" name:"TaskProfile"`
	LayoutIds                 *[]string                         `position:"Query" name:"LayoutIds"  type:"Repeated"`
	StreamURL                 string                            `position:"Query" name:"StreamURL"`
	StreamType                requests.Integer                  `position:"Query" name:"StreamType"`
	UnsubSpecShareScreenUsers *[]string                         `position:"Query" name:"UnsubSpecShareScreenUsers"  type:"Repeated"`
	SubSpecShareScreenUsers   *[]string                         `position:"Query" name:"SubSpecShareScreenUsers"  type:"Repeated"`
	SubSpecUsers              *[]string                         `position:"Query" name:"SubSpecUsers"  type:"Repeated"`
	AppId                     string                            `position:"Query" name:"AppId"`
	Backgrounds               *[]StartMPUTaskBackgrounds        `position:"Query" name:"Backgrounds"  type:"Repeated"`
	TimeStampRef              requests.Integer                  `position:"Query" name:"TimeStampRef"`
	MixMode                   requests.Integer                  `position:"Query" name:"MixMode"`
	ChannelId                 string                            `position:"Query" name:"ChannelId"`
}

// StartMPUTaskUserPanes is a repeated param struct in StartMPUTaskRequest
type StartMPUTaskUserPanes struct {
	Images      *[]StartMPUTaskUserPanesImages `name:"Images" type:"Repeated"`
	SegmentType string                         `name:"SegmentType"`
	UserId      string                         `name:"UserId"`
	Texts       *[]StartMPUTaskUserPanesTexts  `name:"Texts" type:"Repeated"`
	SourceType  string                         `name:"SourceType"`
	PaneId      string                         `name:"PaneId"`
}

// StartMPUTaskClockWidgets is a repeated param struct in StartMPUTaskRequest
type StartMPUTaskClockWidgets struct {
	FontType       string `name:"FontType"`
	FontColor      string `name:"FontColor"`
	Y              string `name:"Y"`
	ZOrder         string `name:"ZOrder"`
	X              string `name:"X"`
	FontSize       string `name:"FontSize"`
	BorderWidth    string `name:"BorderWidth"`
	BorderColor    string `name:"BorderColor"`
	Box            string `name:"Box"`
	BoxColor       string `name:"BoxColor"`
	BoxBorderWidth string `name:"BoxBorderWidth"`
	Alpha          string `name:"Alpha"`
}

// StartMPUTaskWatermarks is a repeated param struct in StartMPUTaskRequest
type StartMPUTaskWatermarks struct {
	Alpha   string `name:"Alpha"`
	Width   string `name:"Width"`
	Height  string `name:"Height"`
	Y       string `name:"Y"`
	Url     string `name:"Url"`
	Display string `name:"Display"`
	ZOrder  string `name:"ZOrder"`
	X       string `name:"X"`
}

// StartMPUTaskEnhancedParam is a repeated param struct in StartMPUTaskRequest
type StartMPUTaskEnhancedParam struct {
	EnablePortraitSegmentation string `name:"EnablePortraitSegmentation"`
	EnableVirtualConference    string `name:"EnableVirtualConference"`
	EnableCartoonPortrait      string `name:"EnableCartoonPortrait"`
	EnableVoiceChanger         string `name:"EnableVoiceChanger"`
	EnableUserPaneBackground   string `name:"EnableUserPaneBackground"`
	BackgroundPath             string `name:"BackgroundPath"`
}

// StartMPUTaskOutputStreamParams is a repeated param struct in StartMPUTaskRequest
type StartMPUTaskOutputStreamParams struct {
	StreamURL string `name:"StreamURL"`
}

// StartMPUTaskBackgrounds is a repeated param struct in StartMPUTaskRequest
type StartMPUTaskBackgrounds struct {
	Width   string `name:"Width"`
	Height  string `name:"Height"`
	Y       string `name:"Y"`
	Url     string `name:"Url"`
	Display string `name:"Display"`
	ZOrder  string `name:"ZOrder"`
	X       string `name:"X"`
}

// StartMPUTaskUserPanesImages is a repeated param struct in StartMPUTaskRequest
type StartMPUTaskUserPanesImages struct {
	Width   string `name:"Width"`
	Height  string `name:"Height"`
	Y       string `name:"Y"`
	Url     string `name:"Url"`
	Display string `name:"Display"`
	ZOrder  string `name:"ZOrder"`
	X       string `name:"X"`
}

// StartMPUTaskUserPanesTexts is a repeated param struct in StartMPUTaskRequest
type StartMPUTaskUserPanesTexts struct {
	FontType       string `name:"FontType"`
	FontColor      string `name:"FontColor"`
	Y              string `name:"Y"`
	Text           string `name:"Text"`
	ZOrder         string `name:"ZOrder"`
	X              string `name:"X"`
	FontSize       string `name:"FontSize"`
	BorderWidth    string `name:"BorderWidth"`
	BorderColor    string `name:"BorderColor"`
	Box            string `name:"Box"`
	BoxColor       string `name:"BoxColor"`
	BoxBorderWidth string `name:"BoxBorderWidth"`
	Alpha          string `name:"Alpha"`
}

// StartMPUTaskResponse is the response struct for api StartMPUTask
type StartMPUTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartMPUTaskRequest creates a request to invoke StartMPUTask API
func CreateStartMPUTaskRequest() (request *StartMPUTaskRequest) {
	request = &StartMPUTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "StartMPUTask", "", "")
	request.Method = requests.POST
	return
}

// CreateStartMPUTaskResponse creates a response to parse from StartMPUTask response
func CreateStartMPUTaskResponse() (response *StartMPUTaskResponse) {
	response = &StartMPUTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}