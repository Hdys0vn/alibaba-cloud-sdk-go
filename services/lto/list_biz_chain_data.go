package lto

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

// ListBizChainData invokes the lto.ListBizChainData API synchronously
func (client *Client) ListBizChainData(request *ListBizChainDataRequest) (response *ListBizChainDataResponse, err error) {
	response = CreateListBizChainDataResponse()
	err = client.DoAction(request, response)
	return
}

// ListBizChainDataWithChan invokes the lto.ListBizChainData API asynchronously
func (client *Client) ListBizChainDataWithChan(request *ListBizChainDataRequest) (<-chan *ListBizChainDataResponse, <-chan error) {
	responseChan := make(chan *ListBizChainDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBizChainData(request)
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

// ListBizChainDataWithCallback invokes the lto.ListBizChainData API asynchronously
func (client *Client) ListBizChainDataWithCallback(request *ListBizChainDataRequest, callback func(response *ListBizChainDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBizChainDataResponse
		var err error
		defer close(result)
		response, err = client.ListBizChainData(request)
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

// ListBizChainDataRequest is the request struct for api ListBizChainData
type ListBizChainDataRequest struct {
	*requests.RpcRequest
	Num        requests.Integer `position:"Query" name:"Num"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	BizChainId string           `position:"Query" name:"BizChainId"`
	Size       requests.Integer `position:"Query" name:"Size"`
	IoTDataDID string           `position:"Query" name:"IoTDataDID"`
	MemberId   string           `position:"Query" name:"MemberId"`
}

// ListBizChainDataResponse is the response struct for api ListBizChainData
type ListBizChainDataResponse struct {
	*responses.BaseResponse
	Code           string                 `json:"Code" xml:"Code"`
	HttpStatusCode int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                 `json:"Message" xml:"Message"`
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	Success        bool                   `json:"Success" xml:"Success"`
	Data           DataInListBizChainData `json:"Data" xml:"Data"`
}

// CreateListBizChainDataRequest creates a request to invoke ListBizChainData API
func CreateListBizChainDataRequest() (request *ListBizChainDataRequest) {
	request = &ListBizChainDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "ListBizChainData", "", "")
	request.Method = requests.POST
	return
}

// CreateListBizChainDataResponse creates a response to parse from ListBizChainData response
func CreateListBizChainDataResponse() (response *ListBizChainDataResponse) {
	response = &ListBizChainDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
