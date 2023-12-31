package linkface

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

// DeleteDeviceGroup invokes the linkface.DeleteDeviceGroup API synchronously
// api document: https://help.aliyun.com/api/linkface/deletedevicegroup.html
func (client *Client) DeleteDeviceGroup(request *DeleteDeviceGroupRequest) (response *DeleteDeviceGroupResponse, err error) {
	response = CreateDeleteDeviceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDeviceGroupWithChan invokes the linkface.DeleteDeviceGroup API asynchronously
// api document: https://help.aliyun.com/api/linkface/deletedevicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDeviceGroupWithChan(request *DeleteDeviceGroupRequest) (<-chan *DeleteDeviceGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteDeviceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDeviceGroup(request)
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

// DeleteDeviceGroupWithCallback invokes the linkface.DeleteDeviceGroup API asynchronously
// api document: https://help.aliyun.com/api/linkface/deletedevicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDeviceGroupWithCallback(request *DeleteDeviceGroupRequest, callback func(response *DeleteDeviceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDeviceGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteDeviceGroup(request)
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

// DeleteDeviceGroupRequest is the request struct for api DeleteDeviceGroup
type DeleteDeviceGroupRequest struct {
	*requests.RpcRequest
	IotId      string `position:"Body" name:"IotId"`
	GroupId    string `position:"Body" name:"GroupId"`
	DeviceName string `position:"Body" name:"DeviceName"`
	ProductKey string `position:"Body" name:"ProductKey"`
}

// DeleteDeviceGroupResponse is the response struct for api DeleteDeviceGroup
type DeleteDeviceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteDeviceGroupRequest creates a request to invoke DeleteDeviceGroup API
func CreateDeleteDeviceGroupRequest() (request *DeleteDeviceGroupRequest) {
	request = &DeleteDeviceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkFace", "2018-07-20", "DeleteDeviceGroup", "", "")
	return
}

// CreateDeleteDeviceGroupResponse creates a response to parse from DeleteDeviceGroup response
func CreateDeleteDeviceGroupResponse() (response *DeleteDeviceGroupResponse) {
	response = &DeleteDeviceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
