package cms

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

// PutExporterOutput invokes the cms.PutExporterOutput API synchronously
func (client *Client) PutExporterOutput(request *PutExporterOutputRequest) (response *PutExporterOutputResponse, err error) {
	response = CreatePutExporterOutputResponse()
	err = client.DoAction(request, response)
	return
}

// PutExporterOutputWithChan invokes the cms.PutExporterOutput API asynchronously
func (client *Client) PutExporterOutputWithChan(request *PutExporterOutputRequest) (<-chan *PutExporterOutputResponse, <-chan error) {
	responseChan := make(chan *PutExporterOutputResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutExporterOutput(request)
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

// PutExporterOutputWithCallback invokes the cms.PutExporterOutput API asynchronously
func (client *Client) PutExporterOutputWithCallback(request *PutExporterOutputRequest, callback func(response *PutExporterOutputResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutExporterOutputResponse
		var err error
		defer close(result)
		response, err = client.PutExporterOutput(request)
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

// PutExporterOutputRequest is the request struct for api PutExporterOutput
type PutExporterOutputRequest struct {
	*requests.RpcRequest
	DestName   string `position:"Query" name:"DestName"`
	ConfigJson string `position:"Query" name:"ConfigJson"`
	DestType   string `position:"Query" name:"DestType"`
	Desc       string `position:"Query" name:"Desc"`
}

// PutExporterOutputResponse is the response struct for api PutExporterOutput
type PutExporterOutputResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreatePutExporterOutputRequest creates a request to invoke PutExporterOutput API
func CreatePutExporterOutputRequest() (request *PutExporterOutputRequest) {
	request = &PutExporterOutputRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "PutExporterOutput", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutExporterOutputResponse creates a response to parse from PutExporterOutput response
func CreatePutExporterOutputResponse() (response *PutExporterOutputResponse) {
	response = &PutExporterOutputResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}