package pvtz

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

// UpdateSyncEcsHostTask invokes the pvtz.UpdateSyncEcsHostTask API synchronously
func (client *Client) UpdateSyncEcsHostTask(request *UpdateSyncEcsHostTaskRequest) (response *UpdateSyncEcsHostTaskResponse, err error) {
	response = CreateUpdateSyncEcsHostTaskResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSyncEcsHostTaskWithChan invokes the pvtz.UpdateSyncEcsHostTask API asynchronously
func (client *Client) UpdateSyncEcsHostTaskWithChan(request *UpdateSyncEcsHostTaskRequest) (<-chan *UpdateSyncEcsHostTaskResponse, <-chan error) {
	responseChan := make(chan *UpdateSyncEcsHostTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSyncEcsHostTask(request)
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

// UpdateSyncEcsHostTaskWithCallback invokes the pvtz.UpdateSyncEcsHostTask API asynchronously
func (client *Client) UpdateSyncEcsHostTaskWithCallback(request *UpdateSyncEcsHostTaskRequest, callback func(response *UpdateSyncEcsHostTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSyncEcsHostTaskResponse
		var err error
		defer close(result)
		response, err = client.UpdateSyncEcsHostTask(request)
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

// UpdateSyncEcsHostTaskRequest is the request struct for api UpdateSyncEcsHostTask
type UpdateSyncEcsHostTaskRequest struct {
	*requests.RpcRequest
	UserClientIp string                         `position:"Query" name:"UserClientIp"`
	ZoneId       string                         `position:"Query" name:"ZoneId"`
	Lang         string                         `position:"Query" name:"Lang"`
	Region       *[]UpdateSyncEcsHostTaskRegion `position:"Query" name:"Region"  type:"Repeated"`
	Status       string                         `position:"Query" name:"Status"`
}

// UpdateSyncEcsHostTaskRegion is a repeated param struct in UpdateSyncEcsHostTaskRequest
type UpdateSyncEcsHostTaskRegion struct {
	RegionId string `name:"RegionId"`
	UserId   string `name:"UserId"`
}

// UpdateSyncEcsHostTaskResponse is the response struct for api UpdateSyncEcsHostTask
type UpdateSyncEcsHostTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateSyncEcsHostTaskRequest creates a request to invoke UpdateSyncEcsHostTask API
func CreateUpdateSyncEcsHostTaskRequest() (request *UpdateSyncEcsHostTaskRequest) {
	request = &UpdateSyncEcsHostTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "UpdateSyncEcsHostTask", "pvtz", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateSyncEcsHostTaskResponse creates a response to parse from UpdateSyncEcsHostTask response
func CreateUpdateSyncEcsHostTaskResponse() (response *UpdateSyncEcsHostTaskResponse) {
	response = &UpdateSyncEcsHostTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}