package config

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

// CreateConfigDeliveryChannel invokes the config.CreateConfigDeliveryChannel API synchronously
func (client *Client) CreateConfigDeliveryChannel(request *CreateConfigDeliveryChannelRequest) (response *CreateConfigDeliveryChannelResponse, err error) {
	response = CreateCreateConfigDeliveryChannelResponse()
	err = client.DoAction(request, response)
	return
}

// CreateConfigDeliveryChannelWithChan invokes the config.CreateConfigDeliveryChannel API asynchronously
func (client *Client) CreateConfigDeliveryChannelWithChan(request *CreateConfigDeliveryChannelRequest) (<-chan *CreateConfigDeliveryChannelResponse, <-chan error) {
	responseChan := make(chan *CreateConfigDeliveryChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateConfigDeliveryChannel(request)
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

// CreateConfigDeliveryChannelWithCallback invokes the config.CreateConfigDeliveryChannel API asynchronously
func (client *Client) CreateConfigDeliveryChannelWithCallback(request *CreateConfigDeliveryChannelRequest, callback func(response *CreateConfigDeliveryChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateConfigDeliveryChannelResponse
		var err error
		defer close(result)
		response, err = client.CreateConfigDeliveryChannel(request)
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

// CreateConfigDeliveryChannelRequest is the request struct for api CreateConfigDeliveryChannel
type CreateConfigDeliveryChannelRequest struct {
	*requests.RpcRequest
	NonCompliantNotification            requests.Boolean `position:"Query" name:"NonCompliantNotification"`
	ClientToken                         string           `position:"Query" name:"ClientToken"`
	ConfigurationSnapshot               requests.Boolean `position:"Query" name:"ConfigurationSnapshot"`
	Description                         string           `position:"Query" name:"Description"`
	DeliveryChannelTargetArn            string           `position:"Query" name:"DeliveryChannelTargetArn"`
	DeliveryChannelCondition            string           `position:"Query" name:"DeliveryChannelCondition"`
	ConfigurationItemChangeNotification requests.Boolean `position:"Query" name:"ConfigurationItemChangeNotification"`
	DeliveryChannelName                 string           `position:"Query" name:"DeliveryChannelName"`
	DeliverySnapshotTime                string           `position:"Query" name:"DeliverySnapshotTime"`
	OversizedDataOSSTargetArn           string           `position:"Query" name:"OversizedDataOSSTargetArn"`
	DeliveryChannelType                 string           `position:"Query" name:"DeliveryChannelType"`
}

// CreateConfigDeliveryChannelResponse is the response struct for api CreateConfigDeliveryChannel
type CreateConfigDeliveryChannelResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DeliveryChannelId string `json:"DeliveryChannelId" xml:"DeliveryChannelId"`
}

// CreateCreateConfigDeliveryChannelRequest creates a request to invoke CreateConfigDeliveryChannel API
func CreateCreateConfigDeliveryChannelRequest() (request *CreateConfigDeliveryChannelRequest) {
	request = &CreateConfigDeliveryChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "CreateConfigDeliveryChannel", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateConfigDeliveryChannelResponse creates a response to parse from CreateConfigDeliveryChannel response
func CreateCreateConfigDeliveryChannelResponse() (response *CreateConfigDeliveryChannelResponse) {
	response = &CreateConfigDeliveryChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
