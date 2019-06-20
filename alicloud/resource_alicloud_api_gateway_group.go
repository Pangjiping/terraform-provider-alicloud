package alicloud

import (
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cloudapi"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func resourceAliyunApigatewayGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliyunApigatewayGroupCreate,
		Read:   resourceAliyunApigatewayGroupRead,
		Update: resourceAliyunApigatewayGroupUpdate,
		Delete: resourceAliyunApigatewayGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliyunApigatewayGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	request := cloudapi.CreateCreateApiGroupRequest()
	request.GroupName = d.Get("name").(string)
	request.Description = d.Get("description").(string)

	if err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err := client.WithCloudApiClient(func(cloudApiClient *cloudapi.Client) (interface{}, error) {
			return cloudApiClient.CreateApiGroup(request)
		})
		if err != nil {
			if IsExceptedError(err, RepeatedCommit) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(request.GetActionName(), raw)
		response, _ := raw.(*cloudapi.CreateApiGroupResponse)
		d.SetId(response.GroupId)
		return nil
	}); err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_api_gateway_group", request.GetActionName(), AlibabaCloudSdkGoERROR)
	}

	return resourceAliyunApigatewayGroupRead(d, meta)
}

func resourceAliyunApigatewayGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudApiService := CloudApiService{client}
	apiGroup, err := cloudApiService.DescribeApiGatewayGroup(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("name", apiGroup.GroupName)
	d.Set("description", apiGroup.Description)

	return nil
}

func resourceAliyunApigatewayGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	if d.HasChange("name") || d.HasChange("description") {
		request := cloudapi.CreateModifyApiGroupRequest()
		request.Description = d.Get("description").(string)
		request.GroupName = d.Get("name").(string)
		request.GroupId = d.Id()
		raw, err := client.WithCloudApiClient(func(cloudApiClient *cloudapi.Client) (interface{}, error) {
			return cloudApiClient.ModifyApiGroup(request)
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
		}
		addDebug(request.GetActionName(), raw)
	}
	return resourceAliyunApigatewayGroupRead(d, meta)
}

func resourceAliyunApigatewayGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudApiService := CloudApiService{client}
	request := cloudapi.CreateDeleteApiGroupRequest()
	request.GroupId = d.Id()

	raw, err := client.WithCloudApiClient(func(cloudApiClient *cloudapi.Client) (interface{}, error) {
		return cloudApiClient.DeleteApiGroup(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw)
	return WrapError(cloudApiService.WaitForApiGatewayGroup(d.Id(), Deleted, DefaultTimeout))

}
