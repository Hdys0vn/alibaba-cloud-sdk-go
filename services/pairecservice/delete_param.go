package pairecservice

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

// DeleteParam invokes the pairecservice.DeleteParam API synchronously
func (client *Client) DeleteParam(request *DeleteParamRequest) (response *DeleteParamResponse, err error) {
	response = CreateDeleteParamResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteParamWithChan invokes the pairecservice.DeleteParam API asynchronously
func (client *Client) DeleteParamWithChan(request *DeleteParamRequest) (<-chan *DeleteParamResponse, <-chan error) {
	responseChan := make(chan *DeleteParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteParam(request)
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

// DeleteParamWithCallback invokes the pairecservice.DeleteParam API asynchronously
func (client *Client) DeleteParamWithCallback(request *DeleteParamRequest, callback func(response *DeleteParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteParamResponse
		var err error
		defer close(result)
		response, err = client.DeleteParam(request)
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

// DeleteParamRequest is the request struct for api DeleteParam
type DeleteParamRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	ParamId    string `position:"Path" name:"ParamId"`
}

// DeleteParamResponse is the response struct for api DeleteParam
type DeleteParamResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteParamRequest creates a request to invoke DeleteParam API
func CreateDeleteParamRequest() (request *DeleteParamRequest) {
	request = &DeleteParamRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "DeleteParam", "/api/v1/params/[ParamId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteParamResponse creates a response to parse from DeleteParam response
func CreateDeleteParamResponse() (response *DeleteParamResponse) {
	response = &DeleteParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
