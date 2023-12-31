package aliyuncvc

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

// UpdateGonggeLayout invokes the aliyuncvc.UpdateGonggeLayout API synchronously
func (client *Client) UpdateGonggeLayout(request *UpdateGonggeLayoutRequest) (response *UpdateGonggeLayoutResponse, err error) {
	response = CreateUpdateGonggeLayoutResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGonggeLayoutWithChan invokes the aliyuncvc.UpdateGonggeLayout API asynchronously
func (client *Client) UpdateGonggeLayoutWithChan(request *UpdateGonggeLayoutRequest) (<-chan *UpdateGonggeLayoutResponse, <-chan error) {
	responseChan := make(chan *UpdateGonggeLayoutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGonggeLayout(request)
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

// UpdateGonggeLayoutWithCallback invokes the aliyuncvc.UpdateGonggeLayout API asynchronously
func (client *Client) UpdateGonggeLayoutWithCallback(request *UpdateGonggeLayoutRequest, callback func(response *UpdateGonggeLayoutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGonggeLayoutResponse
		var err error
		defer close(result)
		response, err = client.UpdateGonggeLayout(request)
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

// UpdateGonggeLayoutRequest is the request struct for api UpdateGonggeLayout
type UpdateGonggeLayoutRequest struct {
	*requests.RpcRequest
	MeetingUUID string `position:"Body" name:"MeetingUUID"`
	VideoCount  string `position:"Body" name:"VideoCount"`
	Value       string `position:"Body" name:"Value"`
}

// UpdateGonggeLayoutResponse is the response struct for api UpdateGonggeLayout
type UpdateGonggeLayoutResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateGonggeLayoutRequest creates a request to invoke UpdateGonggeLayout API
func CreateUpdateGonggeLayoutRequest() (request *UpdateGonggeLayoutRequest) {
	request = &UpdateGonggeLayoutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "UpdateGonggeLayout", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGonggeLayoutResponse creates a response to parse from UpdateGonggeLayout response
func CreateUpdateGonggeLayoutResponse() (response *UpdateGonggeLayoutResponse) {
	response = &UpdateGonggeLayoutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
