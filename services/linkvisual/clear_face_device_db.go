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

// ClearFaceDeviceDB invokes the linkvisual.ClearFaceDeviceDB API synchronously
func (client *Client) ClearFaceDeviceDB(request *ClearFaceDeviceDBRequest) (response *ClearFaceDeviceDBResponse, err error) {
	response = CreateClearFaceDeviceDBResponse()
	err = client.DoAction(request, response)
	return
}

// ClearFaceDeviceDBWithChan invokes the linkvisual.ClearFaceDeviceDB API asynchronously
func (client *Client) ClearFaceDeviceDBWithChan(request *ClearFaceDeviceDBRequest) (<-chan *ClearFaceDeviceDBResponse, <-chan error) {
	responseChan := make(chan *ClearFaceDeviceDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ClearFaceDeviceDB(request)
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

// ClearFaceDeviceDBWithCallback invokes the linkvisual.ClearFaceDeviceDB API asynchronously
func (client *Client) ClearFaceDeviceDBWithCallback(request *ClearFaceDeviceDBRequest, callback func(response *ClearFaceDeviceDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ClearFaceDeviceDBResponse
		var err error
		defer close(result)
		response, err = client.ClearFaceDeviceDB(request)
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

// ClearFaceDeviceDBRequest is the request struct for api ClearFaceDeviceDB
type ClearFaceDeviceDBRequest struct {
	*requests.RpcRequest
	IsolationId   string `position:"Query" name:"IsolationId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// ClearFaceDeviceDBResponse is the response struct for api ClearFaceDeviceDB
type ClearFaceDeviceDBResponse struct {
	*responses.BaseResponse
	Code         string                 `json:"Code" xml:"Code"`
	Data         map[string]interface{} `json:"Data" xml:"Data"`
	ErrorMessage string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string                 `json:"RequestId" xml:"RequestId"`
	Success      bool                   `json:"Success" xml:"Success"`
}

// CreateClearFaceDeviceDBRequest creates a request to invoke ClearFaceDeviceDB API
func CreateClearFaceDeviceDBRequest() (request *ClearFaceDeviceDBRequest) {
	request = &ClearFaceDeviceDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "ClearFaceDeviceDB", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateClearFaceDeviceDBResponse creates a response to parse from ClearFaceDeviceDB response
func CreateClearFaceDeviceDBResponse() (response *ClearFaceDeviceDBResponse) {
	response = &ClearFaceDeviceDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
