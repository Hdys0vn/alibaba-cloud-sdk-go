package linkvisual

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

// DeleteFaceDeviceGroup invokes the linkvisual.DeleteFaceDeviceGroup API synchronously
func (client *Client) DeleteFaceDeviceGroup(request *DeleteFaceDeviceGroupRequest) (response *DeleteFaceDeviceGroupResponse, err error) {
	response = CreateDeleteFaceDeviceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceDeviceGroupWithChan invokes the linkvisual.DeleteFaceDeviceGroup API asynchronously
func (client *Client) DeleteFaceDeviceGroupWithChan(request *DeleteFaceDeviceGroupRequest) (<-chan *DeleteFaceDeviceGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceDeviceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaceDeviceGroup(request)
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

// DeleteFaceDeviceGroupWithCallback invokes the linkvisual.DeleteFaceDeviceGroup API asynchronously
func (client *Client) DeleteFaceDeviceGroupWithCallback(request *DeleteFaceDeviceGroupRequest, callback func(response *DeleteFaceDeviceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceDeviceGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaceDeviceGroup(request)
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

// DeleteFaceDeviceGroupRequest is the request struct for api DeleteFaceDeviceGroup
type DeleteFaceDeviceGroupRequest struct {
	*requests.RpcRequest
	IsolationId   string `position:"Query" name:"IsolationId"`
	DeviceGroupId string `position:"Query" name:"DeviceGroupId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// DeleteFaceDeviceGroupResponse is the response struct for api DeleteFaceDeviceGroup
type DeleteFaceDeviceGroupResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateDeleteFaceDeviceGroupRequest creates a request to invoke DeleteFaceDeviceGroup API
func CreateDeleteFaceDeviceGroupRequest() (request *DeleteFaceDeviceGroupRequest) {
	request = &DeleteFaceDeviceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "DeleteFaceDeviceGroup", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFaceDeviceGroupResponse creates a response to parse from DeleteFaceDeviceGroup response
func CreateDeleteFaceDeviceGroupResponse() (response *DeleteFaceDeviceGroupResponse) {
	response = &DeleteFaceDeviceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
