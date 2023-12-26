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

// QueryEventRecordPlans invokes the linkvisual.QueryEventRecordPlans API synchronously
func (client *Client) QueryEventRecordPlans(request *QueryEventRecordPlansRequest) (response *QueryEventRecordPlansResponse, err error) {
	response = CreateQueryEventRecordPlansResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEventRecordPlansWithChan invokes the linkvisual.QueryEventRecordPlans API asynchronously
func (client *Client) QueryEventRecordPlansWithChan(request *QueryEventRecordPlansRequest) (<-chan *QueryEventRecordPlansResponse, <-chan error) {
	responseChan := make(chan *QueryEventRecordPlansResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEventRecordPlans(request)
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

// QueryEventRecordPlansWithCallback invokes the linkvisual.QueryEventRecordPlans API asynchronously
func (client *Client) QueryEventRecordPlansWithCallback(request *QueryEventRecordPlansRequest, callback func(response *QueryEventRecordPlansResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEventRecordPlansResponse
		var err error
		defer close(result)
		response, err = client.QueryEventRecordPlans(request)
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

// QueryEventRecordPlansRequest is the request struct for api QueryEventRecordPlans
type QueryEventRecordPlansRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	ApiProduct  string           `position:"Body" name:"ApiProduct"`
	ApiRevision string           `position:"Body" name:"ApiRevision"`
}

// QueryEventRecordPlansResponse is the response struct for api QueryEventRecordPlans
type QueryEventRecordPlansResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryEventRecordPlansRequest creates a request to invoke QueryEventRecordPlans API
func CreateQueryEventRecordPlansRequest() (request *QueryEventRecordPlansRequest) {
	request = &QueryEventRecordPlansRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryEventRecordPlans", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryEventRecordPlansResponse creates a response to parse from QueryEventRecordPlans response
func CreateQueryEventRecordPlansResponse() (response *QueryEventRecordPlansResponse) {
	response = &QueryEventRecordPlansResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
