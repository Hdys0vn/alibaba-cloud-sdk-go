package live

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

// AddRtsLiveStreamTranscode invokes the live.AddRtsLiveStreamTranscode API synchronously
func (client *Client) AddRtsLiveStreamTranscode(request *AddRtsLiveStreamTranscodeRequest) (response *AddRtsLiveStreamTranscodeResponse, err error) {
	response = CreateAddRtsLiveStreamTranscodeResponse()
	err = client.DoAction(request, response)
	return
}

// AddRtsLiveStreamTranscodeWithChan invokes the live.AddRtsLiveStreamTranscode API asynchronously
func (client *Client) AddRtsLiveStreamTranscodeWithChan(request *AddRtsLiveStreamTranscodeRequest) (<-chan *AddRtsLiveStreamTranscodeResponse, <-chan error) {
	responseChan := make(chan *AddRtsLiveStreamTranscodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddRtsLiveStreamTranscode(request)
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

// AddRtsLiveStreamTranscodeWithCallback invokes the live.AddRtsLiveStreamTranscode API asynchronously
func (client *Client) AddRtsLiveStreamTranscodeWithCallback(request *AddRtsLiveStreamTranscodeRequest, callback func(response *AddRtsLiveStreamTranscodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddRtsLiveStreamTranscodeResponse
		var err error
		defer close(result)
		response, err = client.AddRtsLiveStreamTranscode(request)
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

// AddRtsLiveStreamTranscodeRequest is the request struct for api AddRtsLiveStreamTranscode
type AddRtsLiveStreamTranscodeRequest struct {
	*requests.RpcRequest
	Template        string           `position:"Query" name:"Template"`
	DeleteBframes   requests.Boolean `position:"Query" name:"DeleteBframes"`
	Lazy            string           `position:"Query" name:"Lazy"`
	Gop             string           `position:"Query" name:"Gop"`
	Opus            requests.Boolean `position:"Query" name:"Opus"`
	AudioCodec      string           `position:"Query" name:"AudioCodec"`
	TemplateType    string           `position:"Query" name:"TemplateType"`
	AudioProfile    string           `position:"Query" name:"AudioProfile"`
	Height          requests.Integer `position:"Query" name:"Height"`
	App             string           `position:"Query" name:"App"`
	AudioChannelNum requests.Integer `position:"Query" name:"AudioChannelNum"`
	Profile         requests.Integer `position:"Query" name:"Profile"`
	FPS             requests.Integer `position:"Query" name:"FPS"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	AudioRate       requests.Integer `position:"Query" name:"AudioRate"`
	AudioBitrate    requests.Integer `position:"Query" name:"AudioBitrate"`
	Domain          string           `position:"Query" name:"Domain"`
	Width           requests.Integer `position:"Query" name:"Width"`
	VideoBitrate    requests.Integer `position:"Query" name:"VideoBitrate"`
}

// AddRtsLiveStreamTranscodeResponse is the response struct for api AddRtsLiveStreamTranscode
type AddRtsLiveStreamTranscodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddRtsLiveStreamTranscodeRequest creates a request to invoke AddRtsLiveStreamTranscode API
func CreateAddRtsLiveStreamTranscodeRequest() (request *AddRtsLiveStreamTranscodeRequest) {
	request = &AddRtsLiveStreamTranscodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddRtsLiveStreamTranscode", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddRtsLiveStreamTranscodeResponse creates a response to parse from AddRtsLiveStreamTranscode response
func CreateAddRtsLiveStreamTranscodeResponse() (response *AddRtsLiveStreamTranscodeResponse) {
	response = &AddRtsLiveStreamTranscodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
