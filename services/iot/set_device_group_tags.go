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

// SetDeviceGroupTags invokes the iot.SetDeviceGroupTags API synchronously
func (client *Client) SetDeviceGroupTags(request *SetDeviceGroupTagsRequest) (response *SetDeviceGroupTagsResponse, err error) {
	response = CreateSetDeviceGroupTagsResponse()
	err = client.DoAction(request, response)
	return
}

// SetDeviceGroupTagsWithChan invokes the iot.SetDeviceGroupTags API asynchronously
func (client *Client) SetDeviceGroupTagsWithChan(request *SetDeviceGroupTagsRequest) (<-chan *SetDeviceGroupTagsResponse, <-chan error) {
	responseChan := make(chan *SetDeviceGroupTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDeviceGroupTags(request)
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

// SetDeviceGroupTagsWithCallback invokes the iot.SetDeviceGroupTags API asynchronously
func (client *Client) SetDeviceGroupTagsWithCallback(request *SetDeviceGroupTagsRequest, callback func(response *SetDeviceGroupTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDeviceGroupTagsResponse
		var err error
		defer close(result)
		response, err = client.SetDeviceGroupTags(request)
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

// SetDeviceGroupTagsRequest is the request struct for api SetDeviceGroupTags
type SetDeviceGroupTagsRequest struct {
	*requests.RpcRequest
	GroupType     string `position:"Query" name:"GroupType"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	TagString     string `position:"Query" name:"TagString"`
	GroupId       string `position:"Query" name:"GroupId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// SetDeviceGroupTagsResponse is the response struct for api SetDeviceGroupTags
type SetDeviceGroupTagsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateSetDeviceGroupTagsRequest creates a request to invoke SetDeviceGroupTags API
func CreateSetDeviceGroupTagsRequest() (request *SetDeviceGroupTagsRequest) {
	request = &SetDeviceGroupTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "SetDeviceGroupTags", "", "")
	request.Method = requests.POST
	return
}

// CreateSetDeviceGroupTagsResponse creates a response to parse from SetDeviceGroupTags response
func CreateSetDeviceGroupTagsResponse() (response *SetDeviceGroupTagsResponse) {
	response = &SetDeviceGroupTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
