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

// CreateFault invokes the cloudcallcenter.CreateFault API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createfault.html
func (client *Client) CreateFault(request *CreateFaultRequest) (response *CreateFaultResponse, err error) {
	response = CreateCreateFaultResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFaultWithChan invokes the cloudcallcenter.CreateFault API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createfault.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFaultWithChan(request *CreateFaultRequest) (<-chan *CreateFaultResponse, <-chan error) {
	responseChan := make(chan *CreateFaultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFault(request)
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

// CreateFaultWithCallback invokes the cloudcallcenter.CreateFault API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createfault.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFaultWithCallback(request *CreateFaultRequest, callback func(response *CreateFaultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFaultResponse
		var err error
		defer close(result)
		response, err = client.CreateFault(request)
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

// CreateFaultRequest is the request struct for api CreateFault
type CreateFaultRequest struct {
	*requests.RpcRequest
	AgentOssFileName       string           `position:"Query" name:"AgentOssFileName"`
	Description            string           `position:"Query" name:"Description"`
	OperatingSystemVersion string           `position:"Query" name:"OperatingSystemVersion"`
	StartTime              requests.Integer `position:"Query" name:"StartTime"`
	MicrophoneList         string           `position:"Query" name:"MicrophoneList"`
	ClientPort             string           `position:"Query" name:"ClientPort"`
	CustomFilePath         string           `position:"Query" name:"CustomFilePath"`
	ClientIp               string           `position:"Query" name:"ClientIp"`
	SpeakerList            string           `position:"Query" name:"SpeakerList"`
	AgentId                requests.Integer `position:"Query" name:"AgentId"`
	EndTime                requests.Integer `position:"Query" name:"EndTime"`
	SpeakerEquipment       string           `position:"Query" name:"SpeakerEquipment"`
	ServicePort            string           `position:"Query" name:"ServicePort"`
	ServiceIp              string           `position:"Query" name:"ServiceIp"`
	InstanceId             string           `position:"Query" name:"InstanceId"`
	AgentFilePath          string           `position:"Query" name:"AgentFilePath"`
	ConnectId              string           `position:"Query" name:"ConnectId"`
	CustomOssFileName      string           `position:"Query" name:"CustomOssFileName"`
	MicrophoneEquipment    string           `position:"Query" name:"MicrophoneEquipment"`
	BrowserVersion         string           `position:"Query" name:"BrowserVersion"`
}

// CreateFaultResponse is the response struct for api CreateFault
type CreateFaultResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCreateFaultRequest creates a request to invoke CreateFault API
func CreateCreateFaultRequest() (request *CreateFaultRequest) {
	request = &CreateFaultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateFault", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateFaultResponse creates a response to parse from CreateFault response
func CreateCreateFaultResponse() (response *CreateFaultResponse) {
	response = &CreateFaultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
