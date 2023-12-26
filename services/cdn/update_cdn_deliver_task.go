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

// UpdateCdnDeliverTask invokes the cdn.UpdateCdnDeliverTask API synchronously
func (client *Client) UpdateCdnDeliverTask(request *UpdateCdnDeliverTaskRequest) (response *UpdateCdnDeliverTaskResponse, err error) {
	response = CreateUpdateCdnDeliverTaskResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCdnDeliverTaskWithChan invokes the cdn.UpdateCdnDeliverTask API asynchronously
func (client *Client) UpdateCdnDeliverTaskWithChan(request *UpdateCdnDeliverTaskRequest) (<-chan *UpdateCdnDeliverTaskResponse, <-chan error) {
	responseChan := make(chan *UpdateCdnDeliverTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCdnDeliverTask(request)
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

// UpdateCdnDeliverTaskWithCallback invokes the cdn.UpdateCdnDeliverTask API asynchronously
func (client *Client) UpdateCdnDeliverTaskWithCallback(request *UpdateCdnDeliverTaskRequest, callback func(response *UpdateCdnDeliverTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCdnDeliverTaskResponse
		var err error
		defer close(result)
		response, err = client.UpdateCdnDeliverTask(request)
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

// UpdateCdnDeliverTaskRequest is the request struct for api UpdateCdnDeliverTask
type UpdateCdnDeliverTaskRequest struct {
	*requests.RpcRequest
	Reports    string           `position:"Body" name:"Reports"`
	Schedule   string           `position:"Body" name:"Schedule"`
	Name       string           `position:"Body" name:"Name"`
	DomainName string           `position:"Body" name:"DomainName"`
	Deliver    string           `position:"Body" name:"Deliver"`
	DeliverId  requests.Integer `position:"Body" name:"DeliverId"`
}

// UpdateCdnDeliverTaskResponse is the response struct for api UpdateCdnDeliverTask
type UpdateCdnDeliverTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCdnDeliverTaskRequest creates a request to invoke UpdateCdnDeliverTask API
func CreateUpdateCdnDeliverTaskRequest() (request *UpdateCdnDeliverTaskRequest) {
	request = &UpdateCdnDeliverTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "UpdateCdnDeliverTask", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateCdnDeliverTaskResponse creates a response to parse from UpdateCdnDeliverTask response
func CreateUpdateCdnDeliverTaskResponse() (response *UpdateCdnDeliverTaskResponse) {
	response = &UpdateCdnDeliverTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}