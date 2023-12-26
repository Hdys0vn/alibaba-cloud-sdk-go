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

// CreateDeliveryChannel invokes the config.CreateDeliveryChannel API synchronously
func (client *Client) CreateDeliveryChannel(request *CreateDeliveryChannelRequest) (response *CreateDeliveryChannelResponse, err error) {
	response = CreateCreateDeliveryChannelResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDeliveryChannelWithChan invokes the config.CreateDeliveryChannel API asynchronously
func (client *Client) CreateDeliveryChannelWithChan(request *CreateDeliveryChannelRequest) (<-chan *CreateDeliveryChannelResponse, <-chan error) {
	responseChan := make(chan *CreateDeliveryChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDeliveryChannel(request)
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

// CreateDeliveryChannelWithCallback invokes the config.CreateDeliveryChannel API asynchronously
func (client *Client) CreateDeliveryChannelWithCallback(request *CreateDeliveryChannelRequest, callback func(response *CreateDeliveryChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDeliveryChannelResponse
		var err error
		defer close(result)
		response, err = client.CreateDeliveryChannel(request)
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

// CreateDeliveryChannelRequest is the request struct for api CreateDeliveryChannel
type CreateDeliveryChannelRequest struct {
	*requests.RpcRequest
	NonCompliantNotification            requests.Boolean `position:"Body" name:"NonCompliantNotification"`
	ClientToken                         string           `position:"Body" name:"ClientToken"`
	ConfigurationSnapshot               requests.Boolean `position:"Body" name:"ConfigurationSnapshot"`
	Description                         string           `position:"Body" name:"Description"`
	DeliveryChannelTargetArn            string           `position:"Body" name:"DeliveryChannelTargetArn"`
	DeliveryChannelCondition            string           `position:"Body" name:"DeliveryChannelCondition"`
	ConfigurationItemChangeNotification requests.Boolean `position:"Body" name:"ConfigurationItemChangeNotification"`
	DeliveryChannelAssumeRoleArn        string           `position:"Body" name:"DeliveryChannelAssumeRoleArn"`
	DeliveryChannelName                 string           `position:"Body" name:"DeliveryChannelName"`
	OversizedDataOSSTargetArn           string           `position:"Body" name:"OversizedDataOSSTargetArn"`
	DeliveryChannelType                 string           `position:"Body" name:"DeliveryChannelType"`
}

// CreateDeliveryChannelResponse is the response struct for api CreateDeliveryChannel
type CreateDeliveryChannelResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DeliveryChannelId string `json:"DeliveryChannelId" xml:"DeliveryChannelId"`
}

// CreateCreateDeliveryChannelRequest creates a request to invoke CreateDeliveryChannel API
func CreateCreateDeliveryChannelRequest() (request *CreateDeliveryChannelRequest) {
	request = &CreateDeliveryChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "CreateDeliveryChannel", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDeliveryChannelResponse creates a response to parse from CreateDeliveryChannel response
func CreateCreateDeliveryChannelResponse() (response *CreateDeliveryChannelResponse) {
	response = &CreateDeliveryChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}