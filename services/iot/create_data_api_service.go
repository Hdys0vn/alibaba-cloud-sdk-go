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

// CreateDataAPIService invokes the iot.CreateDataAPIService API synchronously
func (client *Client) CreateDataAPIService(request *CreateDataAPIServiceRequest) (response *CreateDataAPIServiceResponse, err error) {
	response = CreateCreateDataAPIServiceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataAPIServiceWithChan invokes the iot.CreateDataAPIService API asynchronously
func (client *Client) CreateDataAPIServiceWithChan(request *CreateDataAPIServiceRequest) (<-chan *CreateDataAPIServiceResponse, <-chan error) {
	responseChan := make(chan *CreateDataAPIServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataAPIService(request)
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

// CreateDataAPIServiceWithCallback invokes the iot.CreateDataAPIService API asynchronously
func (client *Client) CreateDataAPIServiceWithCallback(request *CreateDataAPIServiceRequest, callback func(response *CreateDataAPIServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataAPIServiceResponse
		var err error
		defer close(result)
		response, err = client.CreateDataAPIService(request)
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

// CreateDataAPIServiceRequest is the request struct for api CreateDataAPIService
type CreateDataAPIServiceRequest struct {
	*requests.RpcRequest
	RequestParam  *[]CreateDataAPIServiceRequestParam  `position:"Body" name:"RequestParam"  type:"Repeated"`
	IotInstanceId string                               `position:"Body" name:"IotInstanceId"`
	ApiPath       string                               `position:"Body" name:"ApiPath"`
	TemplateSql   string                               `position:"Body" name:"TemplateSql"`
	ResponseParam *[]CreateDataAPIServiceResponseParam `position:"Body" name:"ResponseParam"  type:"Repeated"`
	OriginSql     string                               `position:"Body" name:"OriginSql"`
	DisplayName   string                               `position:"Body" name:"DisplayName"`
	ApiProduct    string                               `position:"Body" name:"ApiProduct"`
	ApiRevision   string                               `position:"Body" name:"ApiRevision"`
	Desc          string                               `position:"Body" name:"Desc"`
}

// CreateDataAPIServiceRequestParam is a repeated param struct in CreateDataAPIServiceRequest
type CreateDataAPIServiceRequestParam struct {
	Name     string `name:"Name"`
	Type     string `name:"Type"`
	Desc     string `name:"Desc"`
	Example  string `name:"Example"`
	Required string `name:"Required"`
}

// CreateDataAPIServiceResponseParam is a repeated param struct in CreateDataAPIServiceRequest
type CreateDataAPIServiceResponseParam struct {
	Name     string `name:"Name"`
	Type     string `name:"Type"`
	Desc     string `name:"Desc"`
	Example  string `name:"Example"`
	Required string `name:"Required"`
}

// CreateDataAPIServiceResponse is the response struct for api CreateDataAPIService
type CreateDataAPIServiceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateCreateDataAPIServiceRequest creates a request to invoke CreateDataAPIService API
func CreateCreateDataAPIServiceRequest() (request *CreateDataAPIServiceRequest) {
	request = &CreateDataAPIServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateDataAPIService", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDataAPIServiceResponse creates a response to parse from CreateDataAPIService response
func CreateCreateDataAPIServiceResponse() (response *CreateDataAPIServiceResponse) {
	response = &CreateDataAPIServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
