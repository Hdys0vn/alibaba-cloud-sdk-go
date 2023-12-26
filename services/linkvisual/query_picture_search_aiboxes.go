package linkvisual

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

// QueryPictureSearchAiboxes invokes the linkvisual.QueryPictureSearchAiboxes API synchronously
func (client *Client) QueryPictureSearchAiboxes(request *QueryPictureSearchAiboxesRequest) (response *QueryPictureSearchAiboxesResponse, err error) {
	response = CreateQueryPictureSearchAiboxesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPictureSearchAiboxesWithChan invokes the linkvisual.QueryPictureSearchAiboxes API asynchronously
func (client *Client) QueryPictureSearchAiboxesWithChan(request *QueryPictureSearchAiboxesRequest) (<-chan *QueryPictureSearchAiboxesResponse, <-chan error) {
	responseChan := make(chan *QueryPictureSearchAiboxesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPictureSearchAiboxes(request)
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

// QueryPictureSearchAiboxesWithCallback invokes the linkvisual.QueryPictureSearchAiboxes API asynchronously
func (client *Client) QueryPictureSearchAiboxesWithCallback(request *QueryPictureSearchAiboxesRequest, callback func(response *QueryPictureSearchAiboxesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPictureSearchAiboxesResponse
		var err error
		defer close(result)
		response, err = client.QueryPictureSearchAiboxes(request)
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

// QueryPictureSearchAiboxesRequest is the request struct for api QueryPictureSearchAiboxes
type QueryPictureSearchAiboxesRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	AppInstanceId string           `position:"Query" name:"AppInstanceId"`
}

// QueryPictureSearchAiboxesResponse is the response struct for api QueryPictureSearchAiboxes
type QueryPictureSearchAiboxesResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryPictureSearchAiboxesRequest creates a request to invoke QueryPictureSearchAiboxes API
func CreateQueryPictureSearchAiboxesRequest() (request *QueryPictureSearchAiboxesRequest) {
	request = &QueryPictureSearchAiboxesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryPictureSearchAiboxes", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryPictureSearchAiboxesResponse creates a response to parse from QueryPictureSearchAiboxes response
func CreateQueryPictureSearchAiboxesResponse() (response *QueryPictureSearchAiboxesResponse) {
	response = &QueryPictureSearchAiboxesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
