package cdn

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

// UpdateCdnSubTask invokes the cdn.UpdateCdnSubTask API synchronously
func (client *Client) UpdateCdnSubTask(request *UpdateCdnSubTaskRequest) (response *UpdateCdnSubTaskResponse, err error) {
	response = CreateUpdateCdnSubTaskResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCdnSubTaskWithChan invokes the cdn.UpdateCdnSubTask API asynchronously
func (client *Client) UpdateCdnSubTaskWithChan(request *UpdateCdnSubTaskRequest) (<-chan *UpdateCdnSubTaskResponse, <-chan error) {
	responseChan := make(chan *UpdateCdnSubTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCdnSubTask(request)
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

// UpdateCdnSubTaskWithCallback invokes the cdn.UpdateCdnSubTask API asynchronously
func (client *Client) UpdateCdnSubTaskWithCallback(request *UpdateCdnSubTaskRequest, callback func(response *UpdateCdnSubTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCdnSubTaskResponse
		var err error
		defer close(result)
		response, err = client.UpdateCdnSubTask(request)
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

// UpdateCdnSubTaskRequest is the request struct for api UpdateCdnSubTask
type UpdateCdnSubTaskRequest struct {
	*requests.RpcRequest
	ReportIds  string `position:"Body" name:"ReportIds"`
	DomainName string `position:"Body" name:"DomainName"`
	EndTime    string `position:"Body" name:"EndTime"`
	StartTime  string `position:"Body" name:"StartTime"`
}

// UpdateCdnSubTaskResponse is the response struct for api UpdateCdnSubTask
type UpdateCdnSubTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCdnSubTaskRequest creates a request to invoke UpdateCdnSubTask API
func CreateUpdateCdnSubTaskRequest() (request *UpdateCdnSubTaskRequest) {
	request = &UpdateCdnSubTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "UpdateCdnSubTask", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateCdnSubTaskResponse creates a response to parse from UpdateCdnSubTask response
func CreateUpdateCdnSubTaskResponse() (response *UpdateCdnSubTaskResponse) {
	response = &UpdateCdnSubTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
