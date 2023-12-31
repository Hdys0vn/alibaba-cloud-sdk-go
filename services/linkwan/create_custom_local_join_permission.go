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

// CreateCustomLocalJoinPermission invokes the linkwan.CreateCustomLocalJoinPermission API synchronously
func (client *Client) CreateCustomLocalJoinPermission(request *CreateCustomLocalJoinPermissionRequest) (response *CreateCustomLocalJoinPermissionResponse, err error) {
	response = CreateCreateCustomLocalJoinPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCustomLocalJoinPermissionWithChan invokes the linkwan.CreateCustomLocalJoinPermission API asynchronously
func (client *Client) CreateCustomLocalJoinPermissionWithChan(request *CreateCustomLocalJoinPermissionRequest) (<-chan *CreateCustomLocalJoinPermissionResponse, <-chan error) {
	responseChan := make(chan *CreateCustomLocalJoinPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCustomLocalJoinPermission(request)
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

// CreateCustomLocalJoinPermissionWithCallback invokes the linkwan.CreateCustomLocalJoinPermission API asynchronously
func (client *Client) CreateCustomLocalJoinPermissionWithCallback(request *CreateCustomLocalJoinPermissionRequest, callback func(response *CreateCustomLocalJoinPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCustomLocalJoinPermissionResponse
		var err error
		defer close(result)
		response, err = client.CreateCustomLocalJoinPermission(request)
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

// CreateCustomLocalJoinPermissionRequest is the request struct for api CreateCustomLocalJoinPermission
type CreateCustomLocalJoinPermissionRequest struct {
	*requests.RpcRequest
	ClassMode           string           `position:"Query" name:"ClassMode"`
	FreqBandPlanGroupId requests.Integer `position:"Query" name:"FreqBandPlanGroupId"`
	JoinEui             string           `position:"Query" name:"JoinEui"`
	ApiProduct          string           `position:"Body" name:"ApiProduct"`
	ApiRevision         string           `position:"Body" name:"ApiRevision"`
	JoinPermissionName  string           `position:"Query" name:"JoinPermissionName"`
}

// CreateCustomLocalJoinPermissionResponse is the response struct for api CreateCustomLocalJoinPermission
type CreateCustomLocalJoinPermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateCreateCustomLocalJoinPermissionRequest creates a request to invoke CreateCustomLocalJoinPermission API
func CreateCreateCustomLocalJoinPermissionRequest() (request *CreateCustomLocalJoinPermissionRequest) {
	request = &CreateCustomLocalJoinPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "CreateCustomLocalJoinPermission", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCustomLocalJoinPermissionResponse creates a response to parse from CreateCustomLocalJoinPermission response
func CreateCreateCustomLocalJoinPermissionResponse() (response *CreateCustomLocalJoinPermissionResponse) {
	response = &CreateCustomLocalJoinPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
