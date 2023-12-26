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

// InvokeDataAPIService invokes the iot.InvokeDataAPIService API synchronously
func (client *Client) InvokeDataAPIService(request *InvokeDataAPIServiceRequest) (response *InvokeDataAPIServiceResponse, err error) {
	response = CreateInvokeDataAPIServiceResponse()
	err = client.DoAction(request, response)
	return
}

// InvokeDataAPIServiceWithChan invokes the iot.InvokeDataAPIService API asynchronously
func (client *Client) InvokeDataAPIServiceWithChan(request *InvokeDataAPIServiceRequest) (<-chan *InvokeDataAPIServiceResponse, <-chan error) {
	responseChan := make(chan *InvokeDataAPIServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InvokeDataAPIService(request)
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

// InvokeDataAPIServiceWithCallback invokes the iot.InvokeDataAPIService API asynchronously
func (client *Client) InvokeDataAPIServiceWithCallback(request *InvokeDataAPIServiceRequest, callback func(response *InvokeDataAPIServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InvokeDataAPIServiceResponse
		var err error
		defer close(result)
		response, err = client.InvokeDataAPIService(request)
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

// InvokeDataAPIServiceRequest is the request struct for api InvokeDataAPIService
type InvokeDataAPIServiceRequest struct {
	*requests.RpcRequest
	Param         *[]InvokeDataAPIServiceParam `position:"Body" name:"Param"  type:"Repeated"`
	IotInstanceId string                       `position:"Body" name:"IotInstanceId"`
	ApiSrn        string                       `position:"Body" name:"ApiSrn"`
	ApiProduct    string                       `position:"Body" name:"ApiProduct"`
	ApiRevision   string                       `position:"Body" name:"ApiRevision"`
}

// InvokeDataAPIServiceParam is a repeated param struct in InvokeDataAPIServiceRequest
type InvokeDataAPIServiceParam struct {
	ParamType      string    `name:"ParamType"`
	ListParamValue *[]string `name:"ListParamValue" type:"Repeated"`
	ListParamType  string    `name:"ListParamType"`
	ParamName      string    `name:"ParamName"`
	ParamValue     string    `name:"ParamValue"`
}

// InvokeDataAPIServiceResponse is the response struct for api InvokeDataAPIService
type InvokeDataAPIServiceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateInvokeDataAPIServiceRequest creates a request to invoke InvokeDataAPIService API
func CreateInvokeDataAPIServiceRequest() (request *InvokeDataAPIServiceRequest) {
	request = &InvokeDataAPIServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "InvokeDataAPIService", "", "")
	request.Method = requests.POST
	return
}

// CreateInvokeDataAPIServiceResponse creates a response to parse from InvokeDataAPIService response
func CreateInvokeDataAPIServiceResponse() (response *InvokeDataAPIServiceResponse) {
	response = &InvokeDataAPIServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
