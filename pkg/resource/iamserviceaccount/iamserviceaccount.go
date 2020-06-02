package iamserviceaccount

import (
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mumoshu/terraform-provider-eksctl/pkg/resource"
)

const KeyNamespace = "namespace"
const KeyName = "name"
const KeyCluster = "cluster"
const KeyOverrideExistingServiceAccounts = "override_existing_serviceaccounts"
const KeyAttachPolicyARNs = "attach_policy_arns"

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: func(d *schema.ResourceData, meta interface{}) error {
			a := ReadIAMServiceAccount(d)

			args := []string{
				"create",
				"iamserviceaccount",
				"--cluster", a.Cluster,
				"--name", a.Name,
				"--namespace", a.Namespace,
			}

			if a.OverrideExistingServiceAccounts {
				args = append(args,
					"--override-existing-serviceaccounts",
				)
			}

			for _, policyARN := range a.AttachPolicyARNs {
				args = append(args,
					"--attach-policy-arn", policyARN,
				)
			}

			return resource.Create(exec.Command("eksctl", args...), d, "")
		},
		Delete: func(d *schema.ResourceData, meta interface{}) error {
			a := ReadIAMServiceAccount(d)

			args := []string{
				"delete",
				"iamserviceaccount",
				"--cluster", a.Cluster,
				"--name", a.Name,
				"--namespace", a.Namespace,
			}

			return resource.Delete(exec.Command("eksctl", args...), d)
		},
		Read: func(d *schema.ResourceData, meta interface{}) error {
			return nil
		},
		Schema: map[string]*schema.Schema{
			KeyNamespace: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "default",
			},
			KeyName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			KeyCluster: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			KeyOverrideExistingServiceAccounts: {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			KeyAttachPolicyARNs: {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			resource.KeyOutput: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

type IAMServiceAccount struct {
	Name                            string
	Namespace                       string
	Cluster                         string
	AttachPolicyARNs                []string
	OverrideExistingServiceAccounts bool
	Output                          string
}

func ReadIAMServiceAccount(d *schema.ResourceData) *IAMServiceAccount {
	a := IAMServiceAccount{}
	a.Namespace = d.Get(KeyNamespace).(string)
	a.Name = d.Get(KeyName).(string)
	a.Cluster = d.Get(KeyCluster).(string)
	a.AttachPolicyARNs = d.Get(KeyAttachPolicyARNs).([]string)
	a.OverrideExistingServiceAccounts = d.Get(KeyOverrideExistingServiceAccounts).(bool)
	return &a
}
