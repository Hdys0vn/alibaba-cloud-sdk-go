package alb

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

// DeleteHealthCheckTemplates invokes the alb.DeleteHealthCheckTemplates API synchronously
func (client *Client) DeleteHealthCheckTemplates(request *DeleteHealthCheckTemplatesRequest) (response *DeleteHealthCheckTemplatesResponse, err error) {
	response = CreateDeleteHealthCheckTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteHealthCheckTemplatesWithChan invokes the alb.DeleteHealthCheckTemplates API asynchronously
func (client *Client) DeleteHealthCheckTemplatesWithChan(request *DeleteHealthCheckTemplatesRequest) (<-chan *DeleteHealthCheckTemplatesResponse, <-chan error) {
	responseChan := make(chan *DeleteHealthCheckTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHealthCheckTemplates(request)
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

// DeleteHealthCheckTemplatesWithCallback invokes the alb.DeleteHealthCheckTemplates API asynchronously
func (client *Client) DeleteHealthCheckTemplatesWithCallback(request *DeleteHealthCheckTemplatesRequest, callback func(response *DeleteHealthCheckTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHealthCheckTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DeleteHealthCheckTemplates(request)
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

// DeleteHealthCheckTemplatesRequest is the request struct for api DeleteHealthCheckTemplates
type DeleteHealthCheckTemplatesRequest struct {
	*requests.RpcRequest
	ClientToken            string           `position:"Query" name:"ClientToken"`
	DryRun                 requests.Boolean `position:"Query" name:"DryRun"`
	HealthCheckTemplateIds *[]string        `position:"Query" name:"HealthCheckTemplateIds"  type:"Repeated"`
}

// DeleteHealthCheckTemplatesResponse is the response struct for api DeleteHealthCheckTemplates
type DeleteHealthCheckTemplatesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteHealthCheckTemplatesRequest creates a request to invoke DeleteHealthCheckTemplates API
func CreateDeleteHealthCheckTemplatesRequest() (request *DeleteHealthCheckTemplatesRequest) {
	request = &DeleteHealthCheckTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "DeleteHealthCheckTemplates", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteHealthCheckTemplatesResponse creates a response to parse from DeleteHealthCheckTemplates response
func CreateDeleteHealthCheckTemplatesResponse() (response *DeleteHealthCheckTemplatesResponse) {
	response = &DeleteHealthCheckTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
