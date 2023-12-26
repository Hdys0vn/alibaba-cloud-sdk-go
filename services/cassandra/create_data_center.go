package cassandra

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

// CreateDataCenter invokes the cassandra.CreateDataCenter API synchronously
func (client *Client) CreateDataCenter(request *CreateDataCenterRequest) (response *CreateDataCenterResponse, err error) {
	response = CreateCreateDataCenterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataCenterWithChan invokes the cassandra.CreateDataCenter API asynchronously
func (client *Client) CreateDataCenterWithChan(request *CreateDataCenterRequest) (<-chan *CreateDataCenterResponse, <-chan error) {
	responseChan := make(chan *CreateDataCenterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataCenter(request)
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

// CreateDataCenterWithCallback invokes the cassandra.CreateDataCenter API asynchronously
func (client *Client) CreateDataCenterWithCallback(request *CreateDataCenterRequest, callback func(response *CreateDataCenterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataCenterResponse
		var err error
		defer close(result)
		response, err = client.CreateDataCenter(request)
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

// CreateDataCenterRequest is the request struct for api CreateDataCenter
type CreateDataCenterRequest struct {
	*requests.RpcRequest
	ClientToken     string           `position:"Query" name:"ClientToken"`
	InstanceType    string           `position:"Query" name:"InstanceType"`
	AutoRenewPeriod requests.Integer `position:"Query" name:"AutoRenewPeriod"`
	Period          requests.Integer `position:"Query" name:"Period"`
	DiskSize        requests.Integer `position:"Query" name:"DiskSize"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	DiskType        string           `position:"Query" name:"DiskType"`
	VswitchId       string           `position:"Query" name:"VswitchId"`
	PeriodUnit      string           `position:"Query" name:"PeriodUnit"`
	AutoRenew       requests.Boolean `position:"Query" name:"AutoRenew"`
	DataCenterName  string           `position:"Query" name:"DataCenterName"`
	NodeCount       requests.Integer `position:"Query" name:"NodeCount"`
	VpcId           string           `position:"Query" name:"VpcId"`
	ZoneId          string           `position:"Query" name:"ZoneId"`
	PayType         string           `position:"Query" name:"PayType"`
}

// CreateDataCenterResponse is the response struct for api CreateDataCenter
type CreateDataCenterResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DataCenterId string `json:"DataCenterId" xml:"DataCenterId"`
}

// CreateCreateDataCenterRequest creates a request to invoke CreateDataCenter API
func CreateCreateDataCenterRequest() (request *CreateDataCenterRequest) {
	request = &CreateDataCenterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "CreateDataCenter", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDataCenterResponse creates a response to parse from CreateDataCenter response
func CreateCreateDataCenterResponse() (response *CreateDataCenterResponse) {
	response = &CreateDataCenterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
