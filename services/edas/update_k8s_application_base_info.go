package edas

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

// UpdateK8sApplicationBaseInfo invokes the edas.UpdateK8sApplicationBaseInfo API synchronously
func (client *Client) UpdateK8sApplicationBaseInfo(request *UpdateK8sApplicationBaseInfoRequest) (response *UpdateK8sApplicationBaseInfoResponse, err error) {
	response = CreateUpdateK8sApplicationBaseInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateK8sApplicationBaseInfoWithChan invokes the edas.UpdateK8sApplicationBaseInfo API asynchronously
func (client *Client) UpdateK8sApplicationBaseInfoWithChan(request *UpdateK8sApplicationBaseInfoRequest) (<-chan *UpdateK8sApplicationBaseInfoResponse, <-chan error) {
	responseChan := make(chan *UpdateK8sApplicationBaseInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateK8sApplicationBaseInfo(request)
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

// UpdateK8sApplicationBaseInfoWithCallback invokes the edas.UpdateK8sApplicationBaseInfo API asynchronously
func (client *Client) UpdateK8sApplicationBaseInfoWithCallback(request *UpdateK8sApplicationBaseInfoRequest, callback func(response *UpdateK8sApplicationBaseInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateK8sApplicationBaseInfoResponse
		var err error
		defer close(result)
		response, err = client.UpdateK8sApplicationBaseInfo(request)
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

// UpdateK8sApplicationBaseInfoRequest is the request struct for api UpdateK8sApplicationBaseInfo
type UpdateK8sApplicationBaseInfoRequest struct {
	*requests.RoaRequest
	Owner       string `position:"Query" name:"Owner"`
	AppId       string `position:"Query" name:"AppId"`
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
	Description string `position:"Query" name:"Description"`
	Email       string `position:"Query" name:"Email"`
}

// UpdateK8sApplicationBaseInfoResponse is the response struct for api UpdateK8sApplicationBaseInfo
type UpdateK8sApplicationBaseInfoResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateUpdateK8sApplicationBaseInfoRequest creates a request to invoke UpdateK8sApplicationBaseInfo API
func CreateUpdateK8sApplicationBaseInfoRequest() (request *UpdateK8sApplicationBaseInfoRequest) {
	request = &UpdateK8sApplicationBaseInfoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateK8sApplicationBaseInfo", "/pop/v5/oam/update_app_basic_info", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateK8sApplicationBaseInfoResponse creates a response to parse from UpdateK8sApplicationBaseInfo response
func CreateUpdateK8sApplicationBaseInfoResponse() (response *UpdateK8sApplicationBaseInfoResponse) {
	response = &UpdateK8sApplicationBaseInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
