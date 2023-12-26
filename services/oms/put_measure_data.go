package oms

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

// PutMeasureData invokes the oms.PutMeasureData API synchronously
func (client *Client) PutMeasureData(request *PutMeasureDataRequest) (response *PutMeasureDataResponse, err error) {
	response = CreatePutMeasureDataResponse()
	err = client.DoAction(request, response)
	return
}

// PutMeasureDataWithChan invokes the oms.PutMeasureData API asynchronously
func (client *Client) PutMeasureDataWithChan(request *PutMeasureDataRequest) (<-chan *PutMeasureDataResponse, <-chan error) {
	responseChan := make(chan *PutMeasureDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutMeasureData(request)
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

// PutMeasureDataWithCallback invokes the oms.PutMeasureData API asynchronously
func (client *Client) PutMeasureDataWithCallback(request *PutMeasureDataRequest, callback func(response *PutMeasureDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutMeasureDataResponse
		var err error
		defer close(result)
		response, err = client.PutMeasureData(request)
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

// PutMeasureDataRequest is the request struct for api PutMeasureData
type PutMeasureDataRequest struct {
	*requests.RpcRequest
	Filter          string           `position:"Query" name:"Filter"`
	DomainCode      string           `position:"Query" name:"DomainCode"`
	Data            string           `position:"Query" name:"Data"`
	DataType        string           `position:"Query" name:"DataType"`
	Compressed      requests.Boolean `position:"Query" name:"Compressed"`
	ApiType         string           `position:"Query" name:"ApiType"`
	SourceRequestId string           `position:"Query" name:"SourceRequestId"`
}

// PutMeasureDataResponse is the response struct for api PutMeasureData
type PutMeasureDataResponse struct {
	*responses.BaseResponse
	DataType        string `json:"DataType" xml:"DataType"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	DomainCode      string `json:"DomainCode" xml:"DomainCode"`
	ApiType         string `json:"ApiType" xml:"ApiType"`
	SourceRequestId string `json:"SourceRequestId" xml:"SourceRequestId"`
}

// CreatePutMeasureDataRequest creates a request to invoke PutMeasureData API
func CreatePutMeasureDataRequest() (request *PutMeasureDataRequest) {
	request = &PutMeasureDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "PutMeasureData", "", "")
	request.Method = requests.POST
	return
}

// CreatePutMeasureDataResponse creates a response to parse from PutMeasureData response
func CreatePutMeasureDataResponse() (response *PutMeasureDataResponse) {
	response = &PutMeasureDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}