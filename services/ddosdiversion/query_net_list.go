package ddosdiversion

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

// QueryNetList invokes the ddosdiversion.QueryNetList API synchronously
func (client *Client) QueryNetList(request *QueryNetListRequest) (response *QueryNetListResponse, err error) {
	response = CreateQueryNetListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryNetListWithChan invokes the ddosdiversion.QueryNetList API asynchronously
func (client *Client) QueryNetListWithChan(request *QueryNetListRequest) (<-chan *QueryNetListResponse, <-chan error) {
	responseChan := make(chan *QueryNetListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryNetList(request)
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

// QueryNetListWithCallback invokes the ddosdiversion.QueryNetList API asynchronously
func (client *Client) QueryNetListWithCallback(request *QueryNetListRequest, callback func(response *QueryNetListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryNetListResponse
		var err error
		defer close(result)
		response, err = client.QueryNetList(request)
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

// QueryNetListRequest is the request struct for api QueryNetList
type QueryNetListRequest struct {
	*requests.RpcRequest
	Num     requests.Integer `position:"Query" name:"Num"`
	MainNet string           `position:"Query" name:"MainNet"`
	Mode    string           `position:"Query" name:"Mode"`
	Net     string           `position:"Query" name:"Net"`
	SaleId  string           `position:"Query" name:"SaleId"`
	Page    requests.Integer `position:"Query" name:"Page"`
}

// QueryNetListResponse is the response struct for api QueryNetList
type QueryNetListResponse struct {
	*responses.BaseResponse
	Code      int64      `json:"Code" xml:"Code"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Message   string     `json:"Message" xml:"Message"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateQueryNetListRequest creates a request to invoke QueryNetList API
func CreateQueryNetListRequest() (request *QueryNetListRequest) {
	request = &QueryNetListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DdosDiversion", "2023-07-01", "QueryNetList", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryNetListResponse creates a response to parse from QueryNetList response
func CreateQueryNetListResponse() (response *QueryNetListResponse) {
	response = &QueryNetListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}