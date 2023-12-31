package vcs

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

// UpdateCorp invokes the vcs.UpdateCorp API synchronously
func (client *Client) UpdateCorp(request *UpdateCorpRequest) (response *UpdateCorpResponse, err error) {
	response = CreateUpdateCorpResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCorpWithChan invokes the vcs.UpdateCorp API asynchronously
func (client *Client) UpdateCorpWithChan(request *UpdateCorpRequest) (<-chan *UpdateCorpResponse, <-chan error) {
	responseChan := make(chan *UpdateCorpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCorp(request)
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

// UpdateCorpWithCallback invokes the vcs.UpdateCorp API asynchronously
func (client *Client) UpdateCorpWithCallback(request *UpdateCorpRequest, callback func(response *UpdateCorpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCorpResponse
		var err error
		defer close(result)
		response, err = client.UpdateCorp(request)
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

// UpdateCorpRequest is the request struct for api UpdateCorp
type UpdateCorpRequest struct {
	*requests.RpcRequest
	IsvSubId     string `position:"Body" name:"IsvSubId"`
	CorpId       string `position:"Body" name:"CorpId"`
	ParentCorpId string `position:"Body" name:"ParentCorpId"`
	Description  string `position:"Body" name:"Description"`
	IconPath     string `position:"Body" name:"IconPath"`
	AppName      string `position:"Body" name:"AppName"`
	CorpName     string `position:"Body" name:"CorpName"`
}

// UpdateCorpResponse is the response struct for api UpdateCorp
type UpdateCorpResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateUpdateCorpRequest creates a request to invoke UpdateCorp API
func CreateUpdateCorpRequest() (request *UpdateCorpRequest) {
	request = &UpdateCorpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "UpdateCorp", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateCorpResponse creates a response to parse from UpdateCorp response
func CreateUpdateCorpResponse() (response *UpdateCorpResponse) {
	response = &UpdateCorpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
