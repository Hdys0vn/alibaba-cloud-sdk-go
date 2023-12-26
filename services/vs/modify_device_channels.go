package vs

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

// ModifyDeviceChannels invokes the vs.ModifyDeviceChannels API synchronously
func (client *Client) ModifyDeviceChannels(request *ModifyDeviceChannelsRequest) (response *ModifyDeviceChannelsResponse, err error) {
	response = CreateModifyDeviceChannelsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDeviceChannelsWithChan invokes the vs.ModifyDeviceChannels API asynchronously
func (client *Client) ModifyDeviceChannelsWithChan(request *ModifyDeviceChannelsRequest) (<-chan *ModifyDeviceChannelsResponse, <-chan error) {
	responseChan := make(chan *ModifyDeviceChannelsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDeviceChannels(request)
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

// ModifyDeviceChannelsWithCallback invokes the vs.ModifyDeviceChannels API asynchronously
func (client *Client) ModifyDeviceChannelsWithCallback(request *ModifyDeviceChannelsRequest, callback func(response *ModifyDeviceChannelsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDeviceChannelsResponse
		var err error
		defer close(result)
		response, err = client.ModifyDeviceChannels(request)
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

// ModifyDeviceChannelsRequest is the request struct for api ModifyDeviceChannels
type ModifyDeviceChannelsRequest struct {
	*requests.RpcRequest
	DeviceStatus string           `position:"Query" name:"DeviceStatus"`
	Id           string           `position:"Query" name:"Id"`
	ShowLog      string           `position:"Query" name:"ShowLog"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	Channels     string           `position:"Query" name:"Channels"`
	Dsn          string           `position:"Query" name:"Dsn"`
}

// ModifyDeviceChannelsResponse is the response struct for api ModifyDeviceChannels
type ModifyDeviceChannelsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDeviceChannelsRequest creates a request to invoke ModifyDeviceChannels API
func CreateModifyDeviceChannelsRequest() (request *ModifyDeviceChannelsRequest) {
	request = &ModifyDeviceChannelsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ModifyDeviceChannels", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDeviceChannelsResponse creates a response to parse from ModifyDeviceChannels response
func CreateModifyDeviceChannelsResponse() (response *ModifyDeviceChannelsResponse) {
	response = &ModifyDeviceChannelsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
