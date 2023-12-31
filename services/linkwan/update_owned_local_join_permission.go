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

// UpdateOwnedLocalJoinPermission invokes the linkwan.UpdateOwnedLocalJoinPermission API synchronously
func (client *Client) UpdateOwnedLocalJoinPermission(request *UpdateOwnedLocalJoinPermissionRequest) (response *UpdateOwnedLocalJoinPermissionResponse, err error) {
	response = CreateUpdateOwnedLocalJoinPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateOwnedLocalJoinPermissionWithChan invokes the linkwan.UpdateOwnedLocalJoinPermission API asynchronously
func (client *Client) UpdateOwnedLocalJoinPermissionWithChan(request *UpdateOwnedLocalJoinPermissionRequest) (<-chan *UpdateOwnedLocalJoinPermissionResponse, <-chan error) {
	responseChan := make(chan *UpdateOwnedLocalJoinPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateOwnedLocalJoinPermission(request)
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

// UpdateOwnedLocalJoinPermissionWithCallback invokes the linkwan.UpdateOwnedLocalJoinPermission API asynchronously
func (client *Client) UpdateOwnedLocalJoinPermissionWithCallback(request *UpdateOwnedLocalJoinPermissionRequest, callback func(response *UpdateOwnedLocalJoinPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateOwnedLocalJoinPermissionResponse
		var err error
		defer close(result)
		response, err = client.UpdateOwnedLocalJoinPermission(request)
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

// UpdateOwnedLocalJoinPermissionRequest is the request struct for api UpdateOwnedLocalJoinPermission
type UpdateOwnedLocalJoinPermissionRequest struct {
	*requests.RpcRequest
	RxDelay             string           `position:"Query" name:"RxDelay"`
	JoinPermissionId    string           `position:"Query" name:"JoinPermissionId"`
	IotInstanceId       string           `position:"Query" name:"IotInstanceId"`
	ClassMode           string           `position:"Query" name:"ClassMode"`
	FreqBandPlanGroupId requests.Integer `position:"Query" name:"FreqBandPlanGroupId"`
	JoinEui             string           `position:"Query" name:"JoinEui"`
	ApiProduct          string           `position:"Body" name:"ApiProduct"`
	ApiRevision         string           `position:"Body" name:"ApiRevision"`
	JoinPermissionName  string           `position:"Query" name:"JoinPermissionName"`
	DataRate            string           `position:"Query" name:"DataRate"`
}

// UpdateOwnedLocalJoinPermissionResponse is the response struct for api UpdateOwnedLocalJoinPermission
type UpdateOwnedLocalJoinPermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateOwnedLocalJoinPermissionRequest creates a request to invoke UpdateOwnedLocalJoinPermission API
func CreateUpdateOwnedLocalJoinPermissionRequest() (request *UpdateOwnedLocalJoinPermissionRequest) {
	request = &UpdateOwnedLocalJoinPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "UpdateOwnedLocalJoinPermission", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateOwnedLocalJoinPermissionResponse creates a response to parse from UpdateOwnedLocalJoinPermission response
func CreateUpdateOwnedLocalJoinPermissionResponse() (response *UpdateOwnedLocalJoinPermissionResponse) {
	response = &UpdateOwnedLocalJoinPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
