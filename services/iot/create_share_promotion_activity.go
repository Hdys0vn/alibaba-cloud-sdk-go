package iot

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

// CreateSharePromotionActivity invokes the iot.CreateSharePromotionActivity API synchronously
func (client *Client) CreateSharePromotionActivity(request *CreateSharePromotionActivityRequest) (response *CreateSharePromotionActivityResponse, err error) {
	response = CreateCreateSharePromotionActivityResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSharePromotionActivityWithChan invokes the iot.CreateSharePromotionActivity API asynchronously
func (client *Client) CreateSharePromotionActivityWithChan(request *CreateSharePromotionActivityRequest) (<-chan *CreateSharePromotionActivityResponse, <-chan error) {
	responseChan := make(chan *CreateSharePromotionActivityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSharePromotionActivity(request)
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

// CreateSharePromotionActivityWithCallback invokes the iot.CreateSharePromotionActivity API asynchronously
func (client *Client) CreateSharePromotionActivityWithCallback(request *CreateSharePromotionActivityRequest, callback func(response *CreateSharePromotionActivityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSharePromotionActivityResponse
		var err error
		defer close(result)
		response, err = client.CreateSharePromotionActivity(request)
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

// CreateSharePromotionActivityRequest is the request struct for api CreateSharePromotionActivity
type CreateSharePromotionActivityRequest struct {
	*requests.RpcRequest
	StartTime                  requests.Integer `position:"Body" name:"StartTime"`
	IotInstanceId              string           `position:"Body" name:"IotInstanceId"`
	EndTime                    requests.Integer `position:"Body" name:"EndTime"`
	ApiProduct                 string           `position:"Body" name:"ApiProduct"`
	ApiRevision                string           `position:"Body" name:"ApiRevision"`
	SharePromotionActivityName string           `position:"Body" name:"SharePromotionActivityName"`
}

// CreateSharePromotionActivityResponse is the response struct for api CreateSharePromotionActivity
type CreateSharePromotionActivityResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         string `json:"Data" xml:"Data"`
}

// CreateCreateSharePromotionActivityRequest creates a request to invoke CreateSharePromotionActivity API
func CreateCreateSharePromotionActivityRequest() (request *CreateSharePromotionActivityRequest) {
	request = &CreateSharePromotionActivityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateSharePromotionActivity", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateSharePromotionActivityResponse creates a response to parse from CreateSharePromotionActivity response
func CreateCreateSharePromotionActivityResponse() (response *CreateSharePromotionActivityResponse) {
	response = &CreateSharePromotionActivityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}