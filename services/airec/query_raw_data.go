package airec

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

// QueryRawData invokes the airec.QueryRawData API synchronously
func (client *Client) QueryRawData(request *QueryRawDataRequest) (response *QueryRawDataResponse, err error) {
	response = CreateQueryRawDataResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRawDataWithChan invokes the airec.QueryRawData API asynchronously
func (client *Client) QueryRawDataWithChan(request *QueryRawDataRequest) (<-chan *QueryRawDataResponse, <-chan error) {
	responseChan := make(chan *QueryRawDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRawData(request)
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

// QueryRawDataWithCallback invokes the airec.QueryRawData API asynchronously
func (client *Client) QueryRawDataWithCallback(request *QueryRawDataRequest, callback func(response *QueryRawDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRawDataResponse
		var err error
		defer close(result)
		response, err = client.QueryRawData(request)
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

// QueryRawDataRequest is the request struct for api QueryRawData
type QueryRawDataRequest struct {
	*requests.RoaRequest
	ItemId     string `position:"Query" name:"itemId"`
	InstanceId string `position:"Path" name:"instanceId"`
	ItemType   string `position:"Query" name:"itemType"`
	UserType   string `position:"Query" name:"userType"`
	UserId     string `position:"Query" name:"userId"`
	Table      string `position:"Path" name:"table"`
}

// QueryRawDataResponse is the response struct for api QueryRawData
type QueryRawDataResponse struct {
	*responses.BaseResponse
	Code      string                 `json:"code" xml:"code"`
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    map[string]interface{} `json:"result" xml:"result"`
}

// CreateQueryRawDataRequest creates a request to invoke QueryRawData API
func CreateQueryRawDataRequest() (request *QueryRawDataRequest) {
	request = &QueryRawDataRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "QueryRawData", "/v2/openapi/instances/[instanceId]/tables/[table]/raw-data", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryRawDataResponse creates a response to parse from QueryRawData response
func CreateQueryRawDataResponse() (response *QueryRawDataResponse) {
	response = &QueryRawDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
