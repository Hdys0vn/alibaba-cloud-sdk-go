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

// ModifyCasterProgram invokes the live.ModifyCasterProgram API synchronously
func (client *Client) ModifyCasterProgram(request *ModifyCasterProgramRequest) (response *ModifyCasterProgramResponse, err error) {
	response = CreateModifyCasterProgramResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCasterProgramWithChan invokes the live.ModifyCasterProgram API asynchronously
func (client *Client) ModifyCasterProgramWithChan(request *ModifyCasterProgramRequest) (<-chan *ModifyCasterProgramResponse, <-chan error) {
	responseChan := make(chan *ModifyCasterProgramResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCasterProgram(request)
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

// ModifyCasterProgramWithCallback invokes the live.ModifyCasterProgram API asynchronously
func (client *Client) ModifyCasterProgramWithCallback(request *ModifyCasterProgramRequest, callback func(response *ModifyCasterProgramResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCasterProgramResponse
		var err error
		defer close(result)
		response, err = client.ModifyCasterProgram(request)
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

// ModifyCasterProgramRequest is the request struct for api ModifyCasterProgram
type ModifyCasterProgramRequest struct {
	*requests.RpcRequest
	Episode  *[]ModifyCasterProgramEpisode `position:"Query" name:"Episode"  type:"Repeated"`
	CasterId string                        `position:"Query" name:"CasterId"`
	OwnerId  requests.Integer              `position:"Query" name:"OwnerId"`
}

// ModifyCasterProgramEpisode is a repeated param struct in ModifyCasterProgramRequest
type ModifyCasterProgramEpisode struct {
	EndTime     string    `name:"EndTime"`
	StartTime   string    `name:"StartTime"`
	EpisodeName string    `name:"EpisodeName"`
	EpisodeType string    `name:"EpisodeType"`
	EpisodeId   string    `name:"EpisodeId"`
	ResourceId  string    `name:"ResourceId"`
	ComponentId *[]string `name:"ComponentId" type:"Repeated"`
	SwitchType  string    `name:"SwitchType"`
}

// ModifyCasterProgramResponse is the response struct for api ModifyCasterProgram
type ModifyCasterProgramResponse struct {
	*responses.BaseResponse
	CasterId  string `json:"CasterId" xml:"CasterId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCasterProgramRequest creates a request to invoke ModifyCasterProgram API
func CreateModifyCasterProgramRequest() (request *ModifyCasterProgramRequest) {
	request = &ModifyCasterProgramRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ModifyCasterProgram", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyCasterProgramResponse creates a response to parse from ModifyCasterProgram response
func CreateModifyCasterProgramResponse() (response *ModifyCasterProgramResponse) {
	response = &ModifyCasterProgramResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
