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

// AddFaceDeviceToDeviceGroup invokes the linkvisual.AddFaceDeviceToDeviceGroup API synchronously
func (client *Client) AddFaceDeviceToDeviceGroup(request *AddFaceDeviceToDeviceGroupRequest) (response *AddFaceDeviceToDeviceGroupResponse, err error) {
	response = CreateAddFaceDeviceToDeviceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// AddFaceDeviceToDeviceGroupWithChan invokes the linkvisual.AddFaceDeviceToDeviceGroup API asynchronously
func (client *Client) AddFaceDeviceToDeviceGroupWithChan(request *AddFaceDeviceToDeviceGroupRequest) (<-chan *AddFaceDeviceToDeviceGroupResponse, <-chan error) {
	responseChan := make(chan *AddFaceDeviceToDeviceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddFaceDeviceToDeviceGroup(request)
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

// AddFaceDeviceToDeviceGroupWithCallback invokes the linkvisual.AddFaceDeviceToDeviceGroup API asynchronously
func (client *Client) AddFaceDeviceToDeviceGroupWithCallback(request *AddFaceDeviceToDeviceGroupRequest, callback func(response *AddFaceDeviceToDeviceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddFaceDeviceToDeviceGroupResponse
		var err error
		defer close(result)
		response, err = client.AddFaceDeviceToDeviceGroup(request)
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

// AddFaceDeviceToDeviceGroupRequest is the request struct for api AddFaceDeviceToDeviceGroup
type AddFaceDeviceToDeviceGroupRequest struct {
	*requests.RpcRequest
	IsolationId   string `position:"Query" name:"IsolationId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	DeviceGroupId string `position:"Query" name:"DeviceGroupId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// AddFaceDeviceToDeviceGroupResponse is the response struct for api AddFaceDeviceToDeviceGroup
type AddFaceDeviceToDeviceGroupResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateAddFaceDeviceToDeviceGroupRequest creates a request to invoke AddFaceDeviceToDeviceGroup API
func CreateAddFaceDeviceToDeviceGroupRequest() (request *AddFaceDeviceToDeviceGroupRequest) {
	request = &AddFaceDeviceToDeviceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "AddFaceDeviceToDeviceGroup", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddFaceDeviceToDeviceGroupResponse creates a response to parse from AddFaceDeviceToDeviceGroup response
func CreateAddFaceDeviceToDeviceGroupResponse() (response *AddFaceDeviceToDeviceGroupResponse) {
	response = &AddFaceDeviceToDeviceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
