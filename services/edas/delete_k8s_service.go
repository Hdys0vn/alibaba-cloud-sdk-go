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

// DeleteK8sService invokes the edas.DeleteK8sService API synchronously
func (client *Client) DeleteK8sService(request *DeleteK8sServiceRequest) (response *DeleteK8sServiceResponse, err error) {
	response = CreateDeleteK8sServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteK8sServiceWithChan invokes the edas.DeleteK8sService API asynchronously
func (client *Client) DeleteK8sServiceWithChan(request *DeleteK8sServiceRequest) (<-chan *DeleteK8sServiceResponse, <-chan error) {
	responseChan := make(chan *DeleteK8sServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteK8sService(request)
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

// DeleteK8sServiceWithCallback invokes the edas.DeleteK8sService API asynchronously
func (client *Client) DeleteK8sServiceWithCallback(request *DeleteK8sServiceRequest, callback func(response *DeleteK8sServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteK8sServiceResponse
		var err error
		defer close(result)
		response, err = client.DeleteK8sService(request)
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

// DeleteK8sServiceRequest is the request struct for api DeleteK8sService
type DeleteK8sServiceRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
	Name  string `position:"Query" name:"Name"`
}

// DeleteK8sServiceResponse is the response struct for api DeleteK8sService
type DeleteK8sServiceResponse struct {
	*responses.BaseResponse
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteK8sServiceRequest creates a request to invoke DeleteK8sService API
func CreateDeleteK8sServiceRequest() (request *DeleteK8sServiceRequest) {
	request = &DeleteK8sServiceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DeleteK8sService", "/pop/v5/k8s/acs/k8s_service", "Edas", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteK8sServiceResponse creates a response to parse from DeleteK8sService response
func CreateDeleteK8sServiceResponse() (response *DeleteK8sServiceResponse) {
	response = &DeleteK8sServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
