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

// StopK8sApplication invokes the edas.StopK8sApplication API synchronously
func (client *Client) StopK8sApplication(request *StopK8sApplicationRequest) (response *StopK8sApplicationResponse, err error) {
	response = CreateStopK8sApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// StopK8sApplicationWithChan invokes the edas.StopK8sApplication API asynchronously
func (client *Client) StopK8sApplicationWithChan(request *StopK8sApplicationRequest) (<-chan *StopK8sApplicationResponse, <-chan error) {
	responseChan := make(chan *StopK8sApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopK8sApplication(request)
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

// StopK8sApplicationWithCallback invokes the edas.StopK8sApplication API asynchronously
func (client *Client) StopK8sApplicationWithCallback(request *StopK8sApplicationRequest, callback func(response *StopK8sApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopK8sApplicationResponse
		var err error
		defer close(result)
		response, err = client.StopK8sApplication(request)
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

// StopK8sApplicationRequest is the request struct for api StopK8sApplication
type StopK8sApplicationRequest struct {
	*requests.RoaRequest
	AppId   string           `position:"Query" name:"AppId"`
	Timeout requests.Integer `position:"Query" name:"Timeout"`
}

// StopK8sApplicationResponse is the response struct for api StopK8sApplication
type StopK8sApplicationResponse struct {
	*responses.BaseResponse
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateStopK8sApplicationRequest creates a request to invoke StopK8sApplication API
func CreateStopK8sApplicationRequest() (request *StopK8sApplicationRequest) {
	request = &StopK8sApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "StopK8sApplication", "/pop/v5/k8s/acs/stop_k8s_app", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopK8sApplicationResponse creates a response to parse from StopK8sApplication response
func CreateStopK8sApplicationResponse() (response *StopK8sApplicationResponse) {
	response = &StopK8sApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
