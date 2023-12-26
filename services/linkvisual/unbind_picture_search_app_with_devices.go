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

// UnbindPictureSearchAppWithDevices invokes the linkvisual.UnbindPictureSearchAppWithDevices API synchronously
func (client *Client) UnbindPictureSearchAppWithDevices(request *UnbindPictureSearchAppWithDevicesRequest) (response *UnbindPictureSearchAppWithDevicesResponse, err error) {
	response = CreateUnbindPictureSearchAppWithDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindPictureSearchAppWithDevicesWithChan invokes the linkvisual.UnbindPictureSearchAppWithDevices API asynchronously
func (client *Client) UnbindPictureSearchAppWithDevicesWithChan(request *UnbindPictureSearchAppWithDevicesRequest) (<-chan *UnbindPictureSearchAppWithDevicesResponse, <-chan error) {
	responseChan := make(chan *UnbindPictureSearchAppWithDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindPictureSearchAppWithDevices(request)
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

// UnbindPictureSearchAppWithDevicesWithCallback invokes the linkvisual.UnbindPictureSearchAppWithDevices API asynchronously
func (client *Client) UnbindPictureSearchAppWithDevicesWithCallback(request *UnbindPictureSearchAppWithDevicesRequest, callback func(response *UnbindPictureSearchAppWithDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindPictureSearchAppWithDevicesResponse
		var err error
		defer close(result)
		response, err = client.UnbindPictureSearchAppWithDevices(request)
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

// UnbindPictureSearchAppWithDevicesRequest is the request struct for api UnbindPictureSearchAppWithDevices
type UnbindPictureSearchAppWithDevicesRequest struct {
	*requests.RpcRequest
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
	DeviceIotIds  *[]string `position:"Query" name:"DeviceIotIds"  type:"Repeated"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
	AppInstanceId string    `position:"Query" name:"AppInstanceId"`
}

// UnbindPictureSearchAppWithDevicesResponse is the response struct for api UnbindPictureSearchAppWithDevices
type UnbindPictureSearchAppWithDevicesResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateUnbindPictureSearchAppWithDevicesRequest creates a request to invoke UnbindPictureSearchAppWithDevices API
func CreateUnbindPictureSearchAppWithDevicesRequest() (request *UnbindPictureSearchAppWithDevicesRequest) {
	request = &UnbindPictureSearchAppWithDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "UnbindPictureSearchAppWithDevices", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbindPictureSearchAppWithDevicesResponse creates a response to parse from UnbindPictureSearchAppWithDevices response
func CreateUnbindPictureSearchAppWithDevicesResponse() (response *UnbindPictureSearchAppWithDevicesResponse) {
	response = &UnbindPictureSearchAppWithDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}