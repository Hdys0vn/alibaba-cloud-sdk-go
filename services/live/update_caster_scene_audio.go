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

// UpdateCasterSceneAudio invokes the live.UpdateCasterSceneAudio API synchronously
func (client *Client) UpdateCasterSceneAudio(request *UpdateCasterSceneAudioRequest) (response *UpdateCasterSceneAudioResponse, err error) {
	response = CreateUpdateCasterSceneAudioResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCasterSceneAudioWithChan invokes the live.UpdateCasterSceneAudio API asynchronously
func (client *Client) UpdateCasterSceneAudioWithChan(request *UpdateCasterSceneAudioRequest) (<-chan *UpdateCasterSceneAudioResponse, <-chan error) {
	responseChan := make(chan *UpdateCasterSceneAudioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCasterSceneAudio(request)
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

// UpdateCasterSceneAudioWithCallback invokes the live.UpdateCasterSceneAudio API asynchronously
func (client *Client) UpdateCasterSceneAudioWithCallback(request *UpdateCasterSceneAudioRequest, callback func(response *UpdateCasterSceneAudioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCasterSceneAudioResponse
		var err error
		defer close(result)
		response, err = client.UpdateCasterSceneAudio(request)
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

// UpdateCasterSceneAudioRequest is the request struct for api UpdateCasterSceneAudio
type UpdateCasterSceneAudioRequest struct {
	*requests.RpcRequest
	CasterId     string                              `position:"Query" name:"CasterId"`
	OwnerId      requests.Integer                    `position:"Query" name:"OwnerId"`
	AudioLayer   *[]UpdateCasterSceneAudioAudioLayer `position:"Query" name:"AudioLayer"  type:"Repeated"`
	SceneId      string                              `position:"Query" name:"SceneId"`
	MixList      *[]string                           `position:"Query" name:"MixList"  type:"Repeated"`
	FollowEnable requests.Integer                    `position:"Query" name:"FollowEnable"`
}

// UpdateCasterSceneAudioAudioLayer is a repeated param struct in UpdateCasterSceneAudioRequest
type UpdateCasterSceneAudioAudioLayer struct {
	VolumeRate         string `name:"VolumeRate"`
	FixedDelayDuration string `name:"FixedDelayDuration"`
	ValidChannel       string `name:"ValidChannel"`
	Filter             string `name:"Filter"`
}

// UpdateCasterSceneAudioResponse is the response struct for api UpdateCasterSceneAudio
type UpdateCasterSceneAudioResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCasterSceneAudioRequest creates a request to invoke UpdateCasterSceneAudio API
func CreateUpdateCasterSceneAudioRequest() (request *UpdateCasterSceneAudioRequest) {
	request = &UpdateCasterSceneAudioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateCasterSceneAudio", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateCasterSceneAudioResponse creates a response to parse from UpdateCasterSceneAudio response
func CreateUpdateCasterSceneAudioResponse() (response *UpdateCasterSceneAudioResponse) {
	response = &UpdateCasterSceneAudioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}