package elasticsearch

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

// UpdateDescription invokes the elasticsearch.UpdateDescription API synchronously
func (client *Client) UpdateDescription(request *UpdateDescriptionRequest) (response *UpdateDescriptionResponse, err error) {
	response = CreateUpdateDescriptionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDescriptionWithChan invokes the elasticsearch.UpdateDescription API asynchronously
func (client *Client) UpdateDescriptionWithChan(request *UpdateDescriptionRequest) (<-chan *UpdateDescriptionResponse, <-chan error) {
	responseChan := make(chan *UpdateDescriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDescription(request)
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

// UpdateDescriptionWithCallback invokes the elasticsearch.UpdateDescription API asynchronously
func (client *Client) UpdateDescriptionWithCallback(request *UpdateDescriptionRequest, callback func(response *UpdateDescriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDescriptionResponse
		var err error
		defer close(result)
		response, err = client.UpdateDescription(request)
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

// UpdateDescriptionRequest is the request struct for api UpdateDescription
type UpdateDescriptionRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
	Body        string `position:"Body" name:"body"`
}

// UpdateDescriptionResponse is the response struct for api UpdateDescription
type UpdateDescriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdateDescriptionRequest creates a request to invoke UpdateDescription API
func CreateUpdateDescriptionRequest() (request *UpdateDescriptionRequest) {
	request = &UpdateDescriptionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpdateDescription", "/openapi/instances/[InstanceId]/description", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDescriptionResponse creates a response to parse from UpdateDescription response
func CreateUpdateDescriptionResponse() (response *UpdateDescriptionResponse) {
	response = &UpdateDescriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
