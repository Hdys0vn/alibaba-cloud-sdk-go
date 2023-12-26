package linkwan

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

// CreateLocalJoinPermission invokes the linkwan.CreateLocalJoinPermission API synchronously
func (client *Client) CreateLocalJoinPermission(request *CreateLocalJoinPermissionRequest) (response *CreateLocalJoinPermissionResponse, err error) {
	response = CreateCreateLocalJoinPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLocalJoinPermissionWithChan invokes the linkwan.CreateLocalJoinPermission API asynchronously
func (client *Client) CreateLocalJoinPermissionWithChan(request *CreateLocalJoinPermissionRequest) (<-chan *CreateLocalJoinPermissionResponse, <-chan error) {
	responseChan := make(chan *CreateLocalJoinPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLocalJoinPermission(request)
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

// CreateLocalJoinPermissionWithCallback invokes the linkwan.CreateLocalJoinPermission API asynchronously
func (client *Client) CreateLocalJoinPermissionWithCallback(request *CreateLocalJoinPermissionRequest, callback func(response *CreateLocalJoinPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLocalJoinPermissionResponse
		var err error
		defer close(result)
		response, err = client.CreateLocalJoinPermission(request)
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

// CreateLocalJoinPermissionRequest is the request struct for api CreateLocalJoinPermission
type CreateLocalJoinPermissionRequest struct {
	*requests.RpcRequest
	RxDelay             requests.Integer `position:"Query" name:"RxDelay"`
	IotInstanceId       string           `position:"Query" name:"IotInstanceId"`
	UseDefaultJoinEui   requests.Boolean `position:"Query" name:"UseDefaultJoinEui"`
	ClassMode           string           `position:"Query" name:"ClassMode"`
	FreqBandPlanGroupId requests.Integer `position:"Query" name:"FreqBandPlanGroupId"`
	JoinEui             string           `position:"Query" name:"JoinEui"`
	ApiProduct          string           `position:"Body" name:"ApiProduct"`
	ApiRevision         string           `position:"Body" name:"ApiRevision"`
	JoinPermissionName  string           `position:"Query" name:"JoinPermissionName"`
	DataRate            requests.Integer `position:"Query" name:"DataRate"`
}

// CreateLocalJoinPermissionResponse is the response struct for api CreateLocalJoinPermission
type CreateLocalJoinPermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateCreateLocalJoinPermissionRequest creates a request to invoke CreateLocalJoinPermission API
func CreateCreateLocalJoinPermissionRequest() (request *CreateLocalJoinPermissionRequest) {
	request = &CreateLocalJoinPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "CreateLocalJoinPermission", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLocalJoinPermissionResponse creates a response to parse from CreateLocalJoinPermission response
func CreateCreateLocalJoinPermissionResponse() (response *CreateLocalJoinPermissionResponse) {
	response = &CreateLocalJoinPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
