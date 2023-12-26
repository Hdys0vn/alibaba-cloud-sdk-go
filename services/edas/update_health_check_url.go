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

// UpdateHealthCheckUrl invokes the edas.UpdateHealthCheckUrl API synchronously
func (client *Client) UpdateHealthCheckUrl(request *UpdateHealthCheckUrlRequest) (response *UpdateHealthCheckUrlResponse, err error) {
	response = CreateUpdateHealthCheckUrlResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateHealthCheckUrlWithChan invokes the edas.UpdateHealthCheckUrl API asynchronously
func (client *Client) UpdateHealthCheckUrlWithChan(request *UpdateHealthCheckUrlRequest) (<-chan *UpdateHealthCheckUrlResponse, <-chan error) {
	responseChan := make(chan *UpdateHealthCheckUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateHealthCheckUrl(request)
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

// UpdateHealthCheckUrlWithCallback invokes the edas.UpdateHealthCheckUrl API asynchronously
func (client *Client) UpdateHealthCheckUrlWithCallback(request *UpdateHealthCheckUrlRequest, callback func(response *UpdateHealthCheckUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateHealthCheckUrlResponse
		var err error
		defer close(result)
		response, err = client.UpdateHealthCheckUrl(request)
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

// UpdateHealthCheckUrlRequest is the request struct for api UpdateHealthCheckUrl
type UpdateHealthCheckUrlRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
	HcURL string `position:"Query" name:"hcURL"`
}

// UpdateHealthCheckUrlResponse is the response struct for api UpdateHealthCheckUrl
type UpdateHealthCheckUrlResponse struct {
	*responses.BaseResponse
	Code           int    `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HealthCheckURL string `json:"HealthCheckURL" xml:"HealthCheckURL"`
}

// CreateUpdateHealthCheckUrlRequest creates a request to invoke UpdateHealthCheckUrl API
func CreateUpdateHealthCheckUrlRequest() (request *UpdateHealthCheckUrlRequest) {
	request = &UpdateHealthCheckUrlRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateHealthCheckUrl", "/pop/v5/app/modify_hc_url", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateHealthCheckUrlResponse creates a response to parse from UpdateHealthCheckUrl response
func CreateUpdateHealthCheckUrlResponse() (response *UpdateHealthCheckUrlResponse) {
	response = &UpdateHealthCheckUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
