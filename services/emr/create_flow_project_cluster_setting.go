package emr

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

// CreateFlowProjectClusterSetting invokes the emr.CreateFlowProjectClusterSetting API synchronously
func (client *Client) CreateFlowProjectClusterSetting(request *CreateFlowProjectClusterSettingRequest) (response *CreateFlowProjectClusterSettingResponse, err error) {
	response = CreateCreateFlowProjectClusterSettingResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFlowProjectClusterSettingWithChan invokes the emr.CreateFlowProjectClusterSetting API asynchronously
func (client *Client) CreateFlowProjectClusterSettingWithChan(request *CreateFlowProjectClusterSettingRequest) (<-chan *CreateFlowProjectClusterSettingResponse, <-chan error) {
	responseChan := make(chan *CreateFlowProjectClusterSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFlowProjectClusterSetting(request)
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

// CreateFlowProjectClusterSettingWithCallback invokes the emr.CreateFlowProjectClusterSetting API asynchronously
func (client *Client) CreateFlowProjectClusterSettingWithCallback(request *CreateFlowProjectClusterSettingRequest, callback func(response *CreateFlowProjectClusterSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFlowProjectClusterSettingResponse
		var err error
		defer close(result)
		response, err = client.CreateFlowProjectClusterSetting(request)
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

// CreateFlowProjectClusterSettingRequest is the request struct for api CreateFlowProjectClusterSetting
type CreateFlowProjectClusterSettingRequest struct {
	*requests.RpcRequest
	UserList     *[]string `position:"Query" name:"UserList"  type:"Repeated"`
	HostList     *[]string `position:"Query" name:"HostList"  type:"Repeated"`
	ClusterId    string    `position:"Query" name:"ClusterId"`
	DefaultQueue string    `position:"Query" name:"DefaultQueue"`
	DefaultUser  string    `position:"Query" name:"DefaultUser"`
	QueueList    *[]string `position:"Query" name:"QueueList"  type:"Repeated"`
	ProjectId    string    `position:"Query" name:"ProjectId"`
}

// CreateFlowProjectClusterSettingResponse is the response struct for api CreateFlowProjectClusterSetting
type CreateFlowProjectClusterSettingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateCreateFlowProjectClusterSettingRequest creates a request to invoke CreateFlowProjectClusterSetting API
func CreateCreateFlowProjectClusterSettingRequest() (request *CreateFlowProjectClusterSettingRequest) {
	request = &CreateFlowProjectClusterSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateFlowProjectClusterSetting", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateFlowProjectClusterSettingResponse creates a response to parse from CreateFlowProjectClusterSetting response
func CreateCreateFlowProjectClusterSettingResponse() (response *CreateFlowProjectClusterSettingResponse) {
	response = &CreateFlowProjectClusterSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
