package openanalytics_open

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

// AddEndPoint invokes the openanalytics_open.AddEndPoint API synchronously
func (client *Client) AddEndPoint(request *AddEndPointRequest) (response *AddEndPointResponse, err error) {
	response = CreateAddEndPointResponse()
	err = client.DoAction(request, response)
	return
}

// AddEndPointWithChan invokes the openanalytics_open.AddEndPoint API asynchronously
func (client *Client) AddEndPointWithChan(request *AddEndPointRequest) (<-chan *AddEndPointResponse, <-chan error) {
	responseChan := make(chan *AddEndPointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddEndPoint(request)
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

// AddEndPointWithCallback invokes the openanalytics_open.AddEndPoint API asynchronously
func (client *Client) AddEndPointWithCallback(request *AddEndPointRequest, callback func(response *AddEndPointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddEndPointResponse
		var err error
		defer close(result)
		response, err = client.AddEndPoint(request)
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

// AddEndPointRequest is the request struct for api AddEndPoint
type AddEndPointRequest struct {
	*requests.RpcRequest
	Product     string `position:"Body" name:"Product"`
	NetworkType string `position:"Body" name:"NetworkType"`
	Vswitch     string `position:"Body" name:"Vswitch"`
	Zone        string `position:"Body" name:"Zone"`
	VpcID       string `position:"Body" name:"VpcID"`
}

// AddEndPointResponse is the response struct for api AddEndPoint
type AddEndPointResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	RegionId     string       `json:"RegionId" xml:"RegionId"`
	EndPointInfo EndPointInfo `json:"EndPointInfo" xml:"EndPointInfo"`
}

// CreateAddEndPointRequest creates a request to invoke AddEndPoint API
func CreateAddEndPointRequest() (request *AddEndPointRequest) {
	request = &AddEndPointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "AddEndPoint", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddEndPointResponse creates a response to parse from AddEndPoint response
func CreateAddEndPointResponse() (response *AddEndPointResponse) {
	response = &AddEndPointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
