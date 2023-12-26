package edas

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

// MigrateEcu invokes the edas.MigrateEcu API synchronously
func (client *Client) MigrateEcu(request *MigrateEcuRequest) (response *MigrateEcuResponse, err error) {
	response = CreateMigrateEcuResponse()
	err = client.DoAction(request, response)
	return
}

// MigrateEcuWithChan invokes the edas.MigrateEcu API asynchronously
func (client *Client) MigrateEcuWithChan(request *MigrateEcuRequest) (<-chan *MigrateEcuResponse, <-chan error) {
	responseChan := make(chan *MigrateEcuResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MigrateEcu(request)
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

// MigrateEcuWithCallback invokes the edas.MigrateEcu API asynchronously
func (client *Client) MigrateEcuWithCallback(request *MigrateEcuRequest, callback func(response *MigrateEcuResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MigrateEcuResponse
		var err error
		defer close(result)
		response, err = client.MigrateEcu(request)
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

// MigrateEcuRequest is the request struct for api MigrateEcu
type MigrateEcuRequest struct {
	*requests.RoaRequest
	InstanceIds     string `position:"Query" name:"InstanceIds"`
	LogicalRegionId string `position:"Query" name:"LogicalRegionId"`
}

// MigrateEcuResponse is the response struct for api MigrateEcu
type MigrateEcuResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMigrateEcuRequest creates a request to invoke MigrateEcu API
func CreateMigrateEcuRequest() (request *MigrateEcuRequest) {
	request = &MigrateEcuRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "MigrateEcu", "/pop/v5/resource/migrate_ecu", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMigrateEcuResponse creates a response to parse from MigrateEcu response
func CreateMigrateEcuResponse() (response *MigrateEcuResponse) {
	response = &MigrateEcuResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
