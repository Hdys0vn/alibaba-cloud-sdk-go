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

// QueryRecordPlans invokes the linkvisual.QueryRecordPlans API synchronously
func (client *Client) QueryRecordPlans(request *QueryRecordPlansRequest) (response *QueryRecordPlansResponse, err error) {
	response = CreateQueryRecordPlansResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRecordPlansWithChan invokes the linkvisual.QueryRecordPlans API asynchronously
func (client *Client) QueryRecordPlansWithChan(request *QueryRecordPlansRequest) (<-chan *QueryRecordPlansResponse, <-chan error) {
	responseChan := make(chan *QueryRecordPlansResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRecordPlans(request)
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

// QueryRecordPlansWithCallback invokes the linkvisual.QueryRecordPlans API asynchronously
func (client *Client) QueryRecordPlansWithCallback(request *QueryRecordPlansRequest, callback func(response *QueryRecordPlansResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRecordPlansResponse
		var err error
		defer close(result)
		response, err = client.QueryRecordPlans(request)
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

// QueryRecordPlansRequest is the request struct for api QueryRecordPlans
type QueryRecordPlansRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	ApiProduct  string           `position:"Body" name:"ApiProduct"`
	ApiRevision string           `position:"Body" name:"ApiRevision"`
}

// QueryRecordPlansResponse is the response struct for api QueryRecordPlans
type QueryRecordPlansResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryRecordPlansRequest creates a request to invoke QueryRecordPlans API
func CreateQueryRecordPlansRequest() (request *QueryRecordPlansRequest) {
	request = &QueryRecordPlansRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryRecordPlans", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryRecordPlansResponse creates a response to parse from QueryRecordPlans response
func CreateQueryRecordPlansResponse() (response *QueryRecordPlansResponse) {
	response = &QueryRecordPlansResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
