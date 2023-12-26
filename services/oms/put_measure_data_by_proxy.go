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

// PutMeasureDataByProxy invokes the oms.PutMeasureDataByProxy API synchronously
func (client *Client) PutMeasureDataByProxy(request *PutMeasureDataByProxyRequest) (response *PutMeasureDataByProxyResponse, err error) {
	response = CreatePutMeasureDataByProxyResponse()
	err = client.DoAction(request, response)
	return
}

// PutMeasureDataByProxyWithChan invokes the oms.PutMeasureDataByProxy API asynchronously
func (client *Client) PutMeasureDataByProxyWithChan(request *PutMeasureDataByProxyRequest) (<-chan *PutMeasureDataByProxyResponse, <-chan error) {
	responseChan := make(chan *PutMeasureDataByProxyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutMeasureDataByProxy(request)
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

// PutMeasureDataByProxyWithCallback invokes the oms.PutMeasureDataByProxy API asynchronously
func (client *Client) PutMeasureDataByProxyWithCallback(request *PutMeasureDataByProxyRequest, callback func(response *PutMeasureDataByProxyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutMeasureDataByProxyResponse
		var err error
		defer close(result)
		response, err = client.PutMeasureDataByProxy(request)
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

// PutMeasureDataByProxyRequest is the request struct for api PutMeasureDataByProxy
type PutMeasureDataByProxyRequest struct {
	*requests.RpcRequest
	Filter     string           `position:"Query" name:"Filter"`
	DomainCode string           `position:"Query" name:"DomainCode"`
	Data       string           `position:"Query" name:"Data"`
	DataType   string           `position:"Query" name:"DataType"`
	Compressed requests.Boolean `position:"Query" name:"Compressed"`
	ApiType    string           `position:"Query" name:"ApiType"`
}

// PutMeasureDataByProxyResponse is the response struct for api PutMeasureDataByProxy
type PutMeasureDataByProxyResponse struct {
	*responses.BaseResponse
	DataType   string `json:"DataType" xml:"DataType"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainCode string `json:"DomainCode" xml:"DomainCode"`
	ApiType    string `json:"ApiType" xml:"ApiType"`
}

// CreatePutMeasureDataByProxyRequest creates a request to invoke PutMeasureDataByProxy API
func CreatePutMeasureDataByProxyRequest() (request *PutMeasureDataByProxyRequest) {
	request = &PutMeasureDataByProxyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "PutMeasureDataByProxy", "", "")
	request.Method = requests.POST
	return
}

// CreatePutMeasureDataByProxyResponse creates a response to parse from PutMeasureDataByProxy response
func CreatePutMeasureDataByProxyResponse() (response *PutMeasureDataByProxyResponse) {
	response = &PutMeasureDataByProxyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}