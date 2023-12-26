package cloudphoto

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

// EditEvent invokes the cloudphoto.EditEvent API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/editevent.html
func (client *Client) EditEvent(request *EditEventRequest) (response *EditEventResponse, err error) {
	response = CreateEditEventResponse()
	err = client.DoAction(request, response)
	return
}

// EditEventWithChan invokes the cloudphoto.EditEvent API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/editevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EditEventWithChan(request *EditEventRequest) (<-chan *EditEventResponse, <-chan error) {
	responseChan := make(chan *EditEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EditEvent(request)
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

// EditEventWithCallback invokes the cloudphoto.EditEvent API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/editevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EditEventWithCallback(request *EditEventRequest, callback func(response *EditEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EditEventResponse
		var err error
		defer close(result)
		response, err = client.EditEvent(request)
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

// EditEventRequest is the request struct for api EditEvent
type EditEventRequest struct {
	*requests.RpcRequest
	EventId          string           `position:"Query" name:"EventId"`
	BannerPhotoId    string           `position:"Query" name:"BannerPhotoId"`
	WatermarkPhotoId string           `position:"Query" name:"WatermarkPhotoId"`
	Identity         string           `position:"Query" name:"Identity"`
	SplashPhotoId    string           `position:"Query" name:"SplashPhotoId"`
	LibraryId        string           `position:"Query" name:"LibraryId"`
	WeixinTitle      string           `position:"Query" name:"WeixinTitle"`
	StoreName        string           `position:"Query" name:"StoreName"`
	Remark           string           `position:"Query" name:"Remark"`
	Title            string           `position:"Query" name:"Title"`
	EndAt            requests.Integer `position:"Query" name:"EndAt"`
	StartAt          requests.Integer `position:"Query" name:"StartAt"`
}

// EditEventResponse is the response struct for api EditEvent
type EditEventResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Event     Event  `json:"Event" xml:"Event"`
}

// CreateEditEventRequest creates a request to invoke EditEvent API
func CreateEditEventRequest() (request *EditEventRequest) {
	request = &EditEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "EditEvent", "cloudphoto", "openAPI")
	return
}

// CreateEditEventResponse creates a response to parse from EditEvent response
func CreateEditEventResponse() (response *EditEventResponse) {
	response = &EditEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
