package vs

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

// BatchBindTemplate invokes the vs.BatchBindTemplate API synchronously
func (client *Client) BatchBindTemplate(request *BatchBindTemplateRequest) (response *BatchBindTemplateResponse, err error) {
	response = CreateBatchBindTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// BatchBindTemplateWithChan invokes the vs.BatchBindTemplate API asynchronously
func (client *Client) BatchBindTemplateWithChan(request *BatchBindTemplateRequest) (<-chan *BatchBindTemplateResponse, <-chan error) {
	responseChan := make(chan *BatchBindTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchBindTemplate(request)
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

// BatchBindTemplateWithCallback invokes the vs.BatchBindTemplate API asynchronously
func (client *Client) BatchBindTemplateWithCallback(request *BatchBindTemplateRequest, callback func(response *BatchBindTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchBindTemplateResponse
		var err error
		defer close(result)
		response, err = client.BatchBindTemplate(request)
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

// BatchBindTemplateRequest is the request struct for api BatchBindTemplate
type BatchBindTemplateRequest struct {
	*requests.RpcRequest
	Replace      requests.Boolean `position:"Query" name:"Replace"`
	InstanceType string           `position:"Query" name:"InstanceType"`
	ShowLog      string           `position:"Query" name:"ShowLog"`
	ApplyAll     requests.Boolean `position:"Query" name:"ApplyAll"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId   string           `position:"Query" name:"TemplateId"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
}

// BatchBindTemplateResponse is the response struct for api BatchBindTemplate
type BatchBindTemplateResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Bindings  []Binding `json:"Bindings" xml:"Bindings"`
}

// CreateBatchBindTemplateRequest creates a request to invoke BatchBindTemplate API
func CreateBatchBindTemplateRequest() (request *BatchBindTemplateRequest) {
	request = &BatchBindTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "BatchBindTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchBindTemplateResponse creates a response to parse from BatchBindTemplate response
func CreateBatchBindTemplateResponse() (response *BatchBindTemplateResponse) {
	response = &BatchBindTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
