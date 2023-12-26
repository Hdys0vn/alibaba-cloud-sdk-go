package cloudcallcenter

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

// ModifyNotificationConfig invokes the cloudcallcenter.ModifyNotificationConfig API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifynotificationconfig.html
func (client *Client) ModifyNotificationConfig(request *ModifyNotificationConfigRequest) (response *ModifyNotificationConfigResponse, err error) {
	response = CreateModifyNotificationConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyNotificationConfigWithChan invokes the cloudcallcenter.ModifyNotificationConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifynotificationconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyNotificationConfigWithChan(request *ModifyNotificationConfigRequest) (<-chan *ModifyNotificationConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyNotificationConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyNotificationConfig(request)
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

// ModifyNotificationConfigWithCallback invokes the cloudcallcenter.ModifyNotificationConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifynotificationconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyNotificationConfigWithCallback(request *ModifyNotificationConfigRequest, callback func(response *ModifyNotificationConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyNotificationConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyNotificationConfig(request)
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

// ModifyNotificationConfigRequest is the request struct for api ModifyNotificationConfig
type ModifyNotificationConfigRequest struct {
	*requests.RpcRequest
	Subscriptions *[]ModifyNotificationConfigSubscriptions `position:"Query" name:"Subscriptions"  type:"Repeated"`
	AccessPoint   string                                   `position:"Query" name:"AccessPoint"`
	InstanceId    string                                   `position:"Query" name:"InstanceId"`
	Topic         string                                   `position:"Query" name:"Topic"`
	ProducerId    string                                   `position:"Query" name:"ProducerId"`
}

// ModifyNotificationConfigSubscriptions is a repeated param struct in ModifyNotificationConfigRequest
type ModifyNotificationConfigSubscriptions struct {
	DisplayName string `name:"DisplayName"`
	Name        string `name:"Name"`
	Selected    string `name:"Selected"`
}

// ModifyNotificationConfigResponse is the response struct for api ModifyNotificationConfig
type ModifyNotificationConfigResponse struct {
	*responses.BaseResponse
	RequestId      string                                  `json:"RequestId" xml:"RequestId"`
	Success        bool                                    `json:"Success" xml:"Success"`
	Code           string                                  `json:"Code" xml:"Code"`
	Message        string                                  `json:"Message" xml:"Message"`
	HttpStatusCode int                                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ProducerId     string                                  `json:"ProducerId" xml:"ProducerId"`
	AccessPoint    string                                  `json:"AccessPoint" xml:"AccessPoint"`
	Topic          string                                  `json:"Topic" xml:"Topic"`
	Subscriptions  SubscriptionsInModifyNotificationConfig `json:"Subscriptions" xml:"Subscriptions"`
}

// CreateModifyNotificationConfigRequest creates a request to invoke ModifyNotificationConfig API
func CreateModifyNotificationConfigRequest() (request *ModifyNotificationConfigRequest) {
	request = &ModifyNotificationConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyNotificationConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyNotificationConfigResponse creates a response to parse from ModifyNotificationConfig response
func CreateModifyNotificationConfigResponse() (response *ModifyNotificationConfigResponse) {
	response = &ModifyNotificationConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
