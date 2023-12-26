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

// CreateEdgeInstanceChannel invokes the iot.CreateEdgeInstanceChannel API synchronously
func (client *Client) CreateEdgeInstanceChannel(request *CreateEdgeInstanceChannelRequest) (response *CreateEdgeInstanceChannelResponse, err error) {
	response = CreateCreateEdgeInstanceChannelResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEdgeInstanceChannelWithChan invokes the iot.CreateEdgeInstanceChannel API asynchronously
func (client *Client) CreateEdgeInstanceChannelWithChan(request *CreateEdgeInstanceChannelRequest) (<-chan *CreateEdgeInstanceChannelResponse, <-chan error) {
	responseChan := make(chan *CreateEdgeInstanceChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEdgeInstanceChannel(request)
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

// CreateEdgeInstanceChannelWithCallback invokes the iot.CreateEdgeInstanceChannel API asynchronously
func (client *Client) CreateEdgeInstanceChannelWithCallback(request *CreateEdgeInstanceChannelRequest, callback func(response *CreateEdgeInstanceChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEdgeInstanceChannelResponse
		var err error
		defer close(result)
		response, err = client.CreateEdgeInstanceChannel(request)
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

// CreateEdgeInstanceChannelRequest is the request struct for api CreateEdgeInstanceChannel
type CreateEdgeInstanceChannelRequest struct {
	*requests.RpcRequest
	Configs       *[]CreateEdgeInstanceChannelConfigs `position:"Query" name:"Configs"  type:"Repeated"`
	DriverId      string                              `position:"Query" name:"DriverId"`
	IotInstanceId string                              `position:"Query" name:"IotInstanceId"`
	ChannelName   string                              `position:"Query" name:"ChannelName"`
	InstanceId    string                              `position:"Query" name:"InstanceId"`
	ApiProduct    string                              `position:"Body" name:"ApiProduct"`
	ApiRevision   string                              `position:"Body" name:"ApiRevision"`
}

// CreateEdgeInstanceChannelConfigs is a repeated param struct in CreateEdgeInstanceChannelRequest
type CreateEdgeInstanceChannelConfigs struct {
	Format  string `name:"Format"`
	Content string `name:"Content"`
	Key     string `name:"Key"`
}

// CreateEdgeInstanceChannelResponse is the response struct for api CreateEdgeInstanceChannel
type CreateEdgeInstanceChannelResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         string `json:"Data" xml:"Data"`
}

// CreateCreateEdgeInstanceChannelRequest creates a request to invoke CreateEdgeInstanceChannel API
func CreateCreateEdgeInstanceChannelRequest() (request *CreateEdgeInstanceChannelRequest) {
	request = &CreateEdgeInstanceChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateEdgeInstanceChannel", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateEdgeInstanceChannelResponse creates a response to parse from CreateEdgeInstanceChannel response
func CreateCreateEdgeInstanceChannelResponse() (response *CreateEdgeInstanceChannelResponse) {
	response = &CreateEdgeInstanceChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
