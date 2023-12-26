package alidns

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

// UpdateDnsGtmInstanceGlobalConfig invokes the alidns.UpdateDnsGtmInstanceGlobalConfig API synchronously
func (client *Client) UpdateDnsGtmInstanceGlobalConfig(request *UpdateDnsGtmInstanceGlobalConfigRequest) (response *UpdateDnsGtmInstanceGlobalConfigResponse, err error) {
	response = CreateUpdateDnsGtmInstanceGlobalConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDnsGtmInstanceGlobalConfigWithChan invokes the alidns.UpdateDnsGtmInstanceGlobalConfig API asynchronously
func (client *Client) UpdateDnsGtmInstanceGlobalConfigWithChan(request *UpdateDnsGtmInstanceGlobalConfigRequest) (<-chan *UpdateDnsGtmInstanceGlobalConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateDnsGtmInstanceGlobalConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDnsGtmInstanceGlobalConfig(request)
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

// UpdateDnsGtmInstanceGlobalConfigWithCallback invokes the alidns.UpdateDnsGtmInstanceGlobalConfig API asynchronously
func (client *Client) UpdateDnsGtmInstanceGlobalConfigWithCallback(request *UpdateDnsGtmInstanceGlobalConfigRequest, callback func(response *UpdateDnsGtmInstanceGlobalConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDnsGtmInstanceGlobalConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateDnsGtmInstanceGlobalConfig(request)
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

// UpdateDnsGtmInstanceGlobalConfigRequest is the request struct for api UpdateDnsGtmInstanceGlobalConfig
type UpdateDnsGtmInstanceGlobalConfigRequest struct {
	*requests.RpcRequest
	AlertGroup           string                                         `position:"Query" name:"AlertGroup"`
	CnameType            string                                         `position:"Query" name:"CnameType"`
	Lang                 string                                         `position:"Query" name:"Lang"`
	AlertConfig          *[]UpdateDnsGtmInstanceGlobalConfigAlertConfig `position:"Query" name:"AlertConfig"  type:"Repeated"`
	PublicCnameMode      string                                         `position:"Query" name:"PublicCnameMode"`
	PublicUserDomainName string                                         `position:"Query" name:"PublicUserDomainName"`
	Ttl                  requests.Integer                               `position:"Query" name:"Ttl"`
	ForceUpdate          requests.Boolean                               `position:"Query" name:"ForceUpdate"`
	InstanceId           string                                         `position:"Query" name:"InstanceId"`
	InstanceName         string                                         `position:"Query" name:"InstanceName"`
	PublicRr             string                                         `position:"Query" name:"PublicRr"`
	UserClientIp         string                                         `position:"Query" name:"UserClientIp"`
	PublicZoneName       string                                         `position:"Query" name:"PublicZoneName"`
}

// UpdateDnsGtmInstanceGlobalConfigAlertConfig is a repeated param struct in UpdateDnsGtmInstanceGlobalConfigRequest
type UpdateDnsGtmInstanceGlobalConfigAlertConfig struct {
	DingtalkNotice string `name:"DingtalkNotice"`
	SmsNotice      string `name:"SmsNotice"`
	NoticeType     string `name:"NoticeType"`
	EmailNotice    string `name:"EmailNotice"`
}

// UpdateDnsGtmInstanceGlobalConfigResponse is the response struct for api UpdateDnsGtmInstanceGlobalConfig
type UpdateDnsGtmInstanceGlobalConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDnsGtmInstanceGlobalConfigRequest creates a request to invoke UpdateDnsGtmInstanceGlobalConfig API
func CreateUpdateDnsGtmInstanceGlobalConfigRequest() (request *UpdateDnsGtmInstanceGlobalConfigRequest) {
	request = &UpdateDnsGtmInstanceGlobalConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDnsGtmInstanceGlobalConfig", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDnsGtmInstanceGlobalConfigResponse creates a response to parse from UpdateDnsGtmInstanceGlobalConfig response
func CreateUpdateDnsGtmInstanceGlobalConfigResponse() (response *UpdateDnsGtmInstanceGlobalConfigResponse) {
	response = &UpdateDnsGtmInstanceGlobalConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}