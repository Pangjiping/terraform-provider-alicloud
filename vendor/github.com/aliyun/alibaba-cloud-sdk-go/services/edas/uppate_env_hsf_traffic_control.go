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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UppateEnvHsfTrafficControl invokes the edas.UppateEnvHsfTrafficControl API synchronously
func (client *Client) UppateEnvHsfTrafficControl(request *UppateEnvHsfTrafficControlRequest) (response *UppateEnvHsfTrafficControlResponse, err error) {
	response = CreateUppateEnvHsfTrafficControlResponse()
	err = client.DoAction(request, response)
	return
}

// UppateEnvHsfTrafficControlWithChan invokes the edas.UppateEnvHsfTrafficControl API asynchronously
func (client *Client) UppateEnvHsfTrafficControlWithChan(request *UppateEnvHsfTrafficControlRequest) (<-chan *UppateEnvHsfTrafficControlResponse, <-chan error) {
	responseChan := make(chan *UppateEnvHsfTrafficControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UppateEnvHsfTrafficControl(request)
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

// UppateEnvHsfTrafficControlWithCallback invokes the edas.UppateEnvHsfTrafficControl API asynchronously
func (client *Client) UppateEnvHsfTrafficControlWithCallback(request *UppateEnvHsfTrafficControlRequest, callback func(response *UppateEnvHsfTrafficControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UppateEnvHsfTrafficControlResponse
		var err error
		defer close(result)
		response, err = client.UppateEnvHsfTrafficControl(request)
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

// UppateEnvHsfTrafficControlRequest is the request struct for api UppateEnvHsfTrafficControl
type UppateEnvHsfTrafficControlRequest struct {
	*requests.RoaRequest
	ParamTypes      string `position:"Body" name:"ParamTypes"`
	Condition       string `position:"Body" name:"Condition"`
	AppId           string `position:"Body" name:"AppId"`
	LabelAdviceName string `position:"Body" name:"LabelAdviceName"`
	PointcutName    string `position:"Body" name:"PointcutName"`
	ServiceName     string `position:"Body" name:"ServiceName"`
	TriggerPolicy   string `position:"Body" name:"TriggerPolicy"`
	Group           string `position:"Body" name:"Group"`
	MethodName      string `position:"Body" name:"MethodName"`
}

// UppateEnvHsfTrafficControlResponse is the response struct for api UppateEnvHsfTrafficControl
type UppateEnvHsfTrafficControlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUppateEnvHsfTrafficControlRequest creates a request to invoke UppateEnvHsfTrafficControl API
func CreateUppateEnvHsfTrafficControlRequest() (request *UppateEnvHsfTrafficControlRequest) {
	request = &UppateEnvHsfTrafficControlRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UppateEnvHsfTrafficControl", "/pop/v5/gray/env_hsf_traffic_control", "Edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUppateEnvHsfTrafficControlResponse creates a response to parse from UppateEnvHsfTrafficControl response
func CreateUppateEnvHsfTrafficControlResponse() (response *UppateEnvHsfTrafficControlResponse) {
	response = &UppateEnvHsfTrafficControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
