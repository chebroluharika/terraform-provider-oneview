// (C) Copyright 2016 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package oneview

import (
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"path"
	"encoding/json"
	"io/ioutil"
)

func resourceHypervisorClusterProfile() *schema.Resource {
	return &schema.Resource{
		Read: datasourceHypervisorClusterProfileRead,

		Schema: map[string]*schema.Schema{
			"add_host_requests": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString},
				Set: schema.HashString,
			},

			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"compliance_state": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"created": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"e_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"hypervisor_cluster_settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"distributed_switch_usage": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"distributed_switch_version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"drs_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ha_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"multi_nic_v_motion": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"virtual_switch_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"hypervisor_host_profile_template": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_manager_type": {
							Type:     schema.TypeString,
							Optional: true},
						"deployment_plan": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deployment_custom_args": {
										Type:     schema.TypSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"deployment_plan_description": {
										Type:     schema.TypeString,
										Optional: true},
									"deployment_plan_uri": {
										Type:     schema.TypeString,
										Optional: true},
									"name": {
										Type:     schema.TypeString,
										Optional: true},
									"server_password": {
										Type:     schema.TypeString,
										Optional: true},
								}}},
						"host_config_policy": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"leave_host_in_maintenance": {
										Type:     schema.TypeBool,
										Optional: true},
									"use_host_prefix_as_hostname": {
										Type:     schema.TypeBool,
										Optional: true},
									"use_hostname_to_register": {
										Type:     schema.TypeBool,
										Optional: true},
								}}},
						"host_prefix": {
							Type:     schema.TypeString,
							Optional: true},
						"server_profile_template_uri": {
							Type:     schema.TypeString,
							Optional: true},
						"virtual_switch_config_policy": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"configure_port_group": {
										Type:     schema.TypeBool,
										Optional: true},
									"custom_virtual_switches": {
										Type:     schema.TypeBool,
										Optional: true},
									"manage_virtual_switches": {
										Type:     schema.TypeBool,
										Optional: true},
								}}},
						"virtual_switches": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:     schema.TypeString,
										Optional: true},
									"name": {
										Type:     schema.TypeString,
										Optional: true},
									"network_uris": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString},
									},
									"version": {
										Type:     schema.TypeString,
										Optional: true},
									"virtual_switch_port_groups": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:     schema.TypeString,
													Optional: true},
												"name": {
													Type:     schema.TypeString,
													Optional: true},
												"network_uris": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString},
												},
												"virtual_switch_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"action": {
																Type:     schema.TypeString,
																Optional: true},
															"dhcp": {
																Type:     schema.TypeBool,
																Optional: true},
															"ip_address": {
																Type:     schema.TypeString,
																Optional: true},
															"subnet_mask": {
																Type:     schema.TypeString,
																Optional: true},
															"virtual_port_purpose": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString},
															},
														},
													}},
												"vlan": {
													Type:     schema.TypeString,
													Optional: true},
											}},
									},
									"virtual_switch_type": {
										Type:     schema.TypeString,
										Optional: true},
									"virtual_switch_uplinks": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:     schema.TypeString,
													Optional: true},
												"active": {
													Type:     schema.TypeBool,
													Optional: true},
												"mac": {
													Type:     schema.TypeString,
													Optional: true},
												"name": {
													Type:     schema.TypeString,
													Optional: true},
												"vmnic": {
													Type:     schema.TypeString,
													Optional: true},
											},
										},
									},
								},
							}},
					}}},
			"hypervisor_cluster_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"hypervisor_host_profile_uris": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"hypervisor_manager_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"hypervisor_type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"ip_pools": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString},
				Set: schema.HashString,
			},

			"mgmt_ip_settings_override": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"modified": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"path": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"refresh_state": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"scopes_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"shared_storage_volumes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString},
				Set: schema.HashString,
			},

			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"state_reason": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceHypervisorClusterProfileCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hypCP := ov.HypervisorClusterProfile{
		//AddHostRequests: d.Get("add_host_requests").(string),
		Category:        d.Get("category").(string),
		ComplianceState: d.Get("compliance_state").(string),
		Created:         d.Get("created").(string),
		Description:     utils.Nstring(d.Get("description").(string)),
		ETag:            d.Get("e_tag").(string),
		//HypervisorClusterSettings:     d.Get("hypervisor_cluster_settings").(string),
		HypervisorClusterUri: d.Get("hypervisor_cluster_uri").(string),
		//HypervisorHostProfileTemplate: d.Get("hypervisor_host_profile_template").(string),
		HypervisorHostProfileUris: utils.Nstring(d.Get("hypervisor_host_profile_uris").(string)),
		HypervisorManagerUri:      utils.Nstring(d.Get("hypervisor_manager_uri").(string)),
		HypervisorType:            d.Get("hypervisor_type").(string),
		//IpPools:                   utils.Nstring(d.Get("ip_pools").(string)),
		MgmtIpSettingsOverride: d.Get("mgmt_ip_settings_override").(string),
		Modified:               d.Get("modified").(string),
		Name:                   d.Get("name").(string),
		Path:                   d.Get("path").(string),
		RefreshState:           d.Get("refresh_state").(string),
		ScopesUri:              d.Get("scopes_uri").(string),
		//SharedStorageVolumes:      utils.Nstring(d.Get("shared_storage_volumes").(string)),
		State:       d.Get("state").(string),
		StateReason: d.Get("state_reason").(string),
		Status:      d.Get("status").(string),
		Type:        d.Get("type").(string),
		URI:         utils.Nstring(d.Get("uri").(string)),
	}
	HypervisorClusterSettingslist := d.Get("hypervisor_cluster_settings").(*schema.Set).List()
	for _, raw := range HypervisorClusterSettingslist {
		hypervisorClusterSettings := raw.(map[string]interface{})

		hypClusterSettings := ov.HypervisorClusterSettings{
			DistributedSwitchUsage:   hypervisorClusterSettings["distributed_switch_usage"].(string),
			DistributedSwitchVersion: hypervisorClusterSettings["distributed_switch_version"].(string),
			DrsEnabled:               hypervisorClusterSettings["drs_enabled"].(bool),
			HaEnabled:                hypervisorClusterSettings["ha_enabled"].(bool),
			MultiNicVMotion:          hypervisorClusterSettings["multi_nic_v_motion"].(bool),
			Type:                     hypervisorClusterSettings["type"].(string),
			VirtualSwitchType:        hypervisorClusterSettings["virtual_switch_type"].(string),
		}
		hypCP.HypervisorClusterSettings = &hypClusterSettings
	}
	rawHypervisorHostProfileTemplate := d.Get("hypervisor_host_profile_template").(*schema.Set).List()
	hypervisorProfileTemplate := ov.HypervisorProfileTemplate{}
	
	for _, raw := range rawHypervisorHostProfileTemplate {
		/******************* deployment plan start********************/
		var hptdeploymentplan ov.DeploymentPlan
		var dpCustomArgs  []utils.Nstring
		rawHostProfileTemplateItem := raw.(map[string]interface{})
		deploymentpPlan := make([]ov.DeploymentPlan,0)
		rawDeploymentPlan= rawHostProfileTemplateItem ["deployment_plan"].(*schema.Set).List()
		for _,raw2 := range 


		dp_map, _ := (d.Get("hypervisor_host_profile_template")).(map[string]interface{})
		deploymentplanlist := dp_map["deployment_plan"].(*schema.Set).List()
		for _, dp_raw := range deploymentplanlist {
		file1, _ := json.MarshalIndent(dp_raw, "", " ")
		_ = ioutil.WriteFile("dp_raw.json", file1, 0644)
			deploymentPlan := dp_raw.(map[string]interface{})
			/*******************dp_custom-args start***********************/
			if val, ok := deploymentPlan["deployment_custom_args"]; ok {
				dpCustomArgsOrder := val.(*schema.Set).List()
				dpCustomArgs = make([]utils.Nstring, len(dpCustomArgsOrder))
				for i, raw := range dpCustomArgsOrder {
					dpCustomArgs[i] = utils.Nstring(raw.(string))
				}
			}
			/********************dp custom args end**********************/
			hptdeploymentplan = ov.DeploymentPlan{
				DeploymentCustomArgs:      dpCustomArgs,
				DeploymentPlanDescription: deploymentPlan["deployment_plan_description"].(string),
				DeploymentPlanUri:         utils.Nstring(deploymentPlan["deployment_plan_uri"].(string)),
				Name:                      deploymentPlan["name"].(string),
				ServerPassword:            deploymentPlan["server_password"].(string),
			}

		}
		file, _ := json.MarshalIndent(hptdeploymentplan, "", " ")
		_ = ioutil.WriteFile("dp.json", file, 0644)

		/********************deployment plan end**********************************************/
		hostprofiletemplate := raw.(map[string]interface{})
		hypHostProfileTemplate := ov.HypervisorHostProfileTemplate{
			DeploymentManagerType:    hostprofiletemplate["deployment_manager_type"].(string),
			DeploymentPlan:           &hptdeploymentplan,
			Hostprefix:               hostprofiletemplate["host_prefix"].(string),
			ServerProfileTemplateUri: utils.Nstring(hostprofiletemplate["server_profile_template_uri"].(string)),
		}
		hypCP.HypervisorHostProfileTemplate = &hypHostProfileTemplate
	}
	/**********************hypervisor hosr profile end************************************************/
	hypCPError := config.ovClient.CreateHypervisorClusterProfile(hypCP)
	uri := d.Get("URI").(string)
	_, id := path.Split(uri)
        d.SetId(id)
	if hypCPError != nil {
		d.SetId("")
		return hypCPError
	}
	return resourceHypervisorClusterProfileRead(d, meta)
}

func resourceHypervisorClusterProfileRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

/*func resourceHypervisorClusterProfileUpdate(d *schema.ResourceData, meta interface{}) error {%
	config := meta.(*Config)

	hypCP := ov.HypervisorClusterProfile{
		ETAG:        d.Get("etag").(string),
		URI:         utils.NewNstring(d.Get("uri").(string)),
		DisplayName: d.Get("display_name").(string),
		Name:        d.Get("name").(string),
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		Port:        d.Get("port").(int),
	}

	err := config.ovClient.UpdateHypervisorClusterProfile(hypCP)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceHypervisorClusterProfileRead(d, meta)
}*/

func resourceHypervisorClusterProfileDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteHypervisorClusterProfile(d.Id())
	if err != nil {
		return err
	}
	return nil
}
