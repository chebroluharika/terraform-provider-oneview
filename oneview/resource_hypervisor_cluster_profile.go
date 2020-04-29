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
)

func resourceHypervisorClusterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceHypervisorClusterProfileCreate,
		Read:   resourceHypervisorClusterProfileRead,
		Update: resourceHypervisorClusterProfileUpdate,
		Delete: resourceHypervisorClusterProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"add_host_requests": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"compliance_state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"e_tag": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hypervisor_cluster_settings": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"distributed_switch_usage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"distributed_switch_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"drs_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ha_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"mulit_nic_v_motion": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_switch_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"hypervisor_host_profile_template": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_manager_type": {
							Type:     schema.TypeString,
							Computed: true},
						"deployment_plan": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deployment_custom_args": {
									        Type:     schema.TypeSet,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										        Computed: true,
										},
									},
									"deployment_plan_description": {
										Type:     schema.TypeString,
										Computed: true},
									"deployment_plan_uri": {
										Type:     schema.TypeString,
										Computed: true},
									"name": {
										Type:     schema.TypeString,
										Computed: true},
									"server_password": {
										Type:     schema.TypeString,
										Computed: true},
								}}},
						"host_config_policy": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"leave_host_in_maintenance": {
										Type:     schema.TypeBool,
										Computed: true},
									"use_host_prefix_as_hostname": {
										Type:     schema.TypeBool,
										Computed: true},
									"use_hostname_to_register": {
										Type:     schema.TypeBool,
										Computed: true},
								}}},
						"host_prefix": {
							Type:     schema.TypeString,
							Computed: true},
						"server_profile_template_uri": {
							Type:     schema.TypeString,
							Computed: true},
						"virtual_switch_config_policy": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"configure_port_group": {
										Type:     schema.TypeBool,
										Computed: true},
									"custom_virtual_switches": {
										Type:     schema.TypeBool,
										Computed: true},
									"manage_virtual_switches": {
										Type:     schema.TypeBool,
										Computed: true},
								}}},
						"virtual_switches": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:     schema.TypeString,
										Computed: true},
									"name": {
										Type:     schema.TypeString,
										Computed: true},
									"newtwork_uris": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString},
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true},
									"virtual_switch_port_groups": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:     schema.TypeString,
													Computed: true},
												"name": {
													Type:     schema.TypeString,
													Computed: true},
												"newtwork_uris": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString},
												},
												"virtual_switch_ports": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"action": {
																Type:     schema.TypeString,
																Computed: true},
															"dhcp": {
																Type:     schema.TypeBool,
																Computed: true},
															"ip_address": {
																Type:     schema.TypeString,
																Computed: true},
															"subnet_mast": {
																Type:     schema.TypeString,
																Computed: true},
															"virtual_port_purpose": {
																Type:     schema.TypeSet,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString},
															},
														},
													}},
												"vlan": {
													Type:     schema.TypeString,
													Computed: true},
											}},
									},
									"virtual_switch_type": {
										Type:     schema.TypeString,
										Computed: true},
									"virtual_switch_uplinks": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:     schema.TypeString,
													Computed: true},
												"active": {
													Type:     schema.TypeBool,
													Computed: true},
												"mac": {
													Type:     schema.TypeString,
													Computed: true},
												"name": {
													Type:     schema.TypeString,
													Computed: true},
												"vmnic": {
													Type:     schema.TypeString,
													Computed: true},
											},
										},
									},
								},
							}},
					}}},
			"hypervisor_cluster_uri": {
				Type: schema.TypeString,
				Computed: true,
			},

			"hypervisor_host_profile_uris": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hypervisor_manager_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hypervisor_type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"ip_pools": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"mgmt_ip_settings_override": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"refresh_state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"shared_storage_volumes": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"state_reason": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"initial_scope_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
		},
	}
}

func resourceHypervisorClusterProfileCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hypCP := ov.HypervisorClusterProfile{
		Name:                 d.Get("name").(string),
		Description:          d.Get("description").(string),
		HypervisorType:       d.Get("hypervisor_type").(string),
		HypervisorManagerUri: d.Get("password").(string),
		Path:                 d.Get("path").(int),
		Type:                 d.Get("type").(string),
		HypervisorHostProfileTemplate: {
			ServerProfileTeamplateUri: d.Get("server_profile_template_uri").(string),
			HostPrefix:                d.Get("host_prefix"),
		},
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for i, raw := range rawInitialScopeUris {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		hypCP.InitialScopeUris = initialScopeUris
	}
	hypCPError := config.ovClient.CreateHypervisorClusterProfile(hypCP)
	d.SetId(d.Get("name").(string))
	if hypCPError != nil {
		d.SetId("")
		return hypCPError
	}
	return resourceHypervisorClusterProfileRead(d, meta)
}

func resourceHypervisorClusterProfileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hypCP, err := config.ovClient.GetHypervisorClusterProfile(d.Id())
	if err != nil || hypCP.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("add_host_requests", hypCP.AddHostRequests)
	d.Set("category", hypCP.Category)
	d.Set("compliance_state", hypCP.ComplianceState)
	d.Set("created", hypCP.Created)
	d.Set("description", hypCP.Description)
	d.Set("e_tag", hypCP.ETag)
	hypCPCS_list := make([]map[string]interface{}, 0, 1)
	hypCPCS_list = append(hypCPCS_list, map[string]interface{}{
		"type":                       hypCP.HypervisorClusterSettings.DistributedSwitchUsage,
		"virtual_switch_type":        hypCP.HypervisorClusterSettings.DistributedSwitchVersion,
		"distributed_switch_version": hypCP.HypervisorClusterSettings.DistributedSwitchVersion,
		"distributed_switch_usage":   hypCP.HypervisorClusterSettings.DistributedSwitchUsage,
		"multi_nic_v_motion":         hypCP.HypervisorClusterSettings.MultiNicVMotion,
		"drs_enabled":                hypCP.HypervisorClusterSettings.DrsEnabled,
		"ha_enabled":                 hypCP.HypervisorClusterSettings.HaEnabled,
	})

	d.Set("hypervisor_cluster_settings", hypCPCS_list)
	d.Set("hypervisor_cluster_uri", hypCP.HypervisorClusterUri)
	dp_list := make([]map[string]interface{}, 0, 1)
	dp_list = append(dp_list, map[string]interface{}{
		
		"deployment_custom_args":      hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentCustomArgs,
		"deployment_plan_description": hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentPlanDescription,
		"deployment_plan_uri":         hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentPlanUri.string(),
		"name":                        hypCP.HypervisorHostProfileTemplate.DeploymentPlan.Name,
		"server_password":             hypCP.HypervisorHostProfileTemplate.DeploymentPlan.ServerPassword,
	})
	hypCPHHPT_list := make([]map[string]interface{}, 0, 1)
	hypCPHHPT_list = append(hypCPHHPT_list, map[string]interface{}{
		"deployment_manager_type":      hypCP.HypervisorHostProfileTemplate.DeploymentManagerType,
		"deployment_plan":              dp_list,
		"host_config_policy":           hypCP.HypervisorHostProfileTemplate.HostConfigPolicy,
		"host_prefix":                  hypCP.HypervisorHostProfileTemplate.Hostprefix,
		"server_profile_template_uri":  hypCP.HypervisorHostProfileTemplate.ServerProfileTemplateUri.string(),
		"virtual_switch_config_policy": hypCP.HypervisorHostProfileTemplate.VirtualSwitchConfigPolicy,
		"virtual_switches":             hypCP.HypervisorHostProfileTemplate.VirtualSwitches,
	})
	d.Set("hypervisor_host_profile_template", hypCP.HypervisorHostProfileTemplate)
	d.Set("hypervisor_host_profile_uris", hypCP.HypervisorHostProfileUris)
	d.Set("hypervisor_manager_uri", hypCP.HypervisorManagerUri)
	d.Set("hypervisor_type", hypCP.HypervisorType)
	d.Set("ip_pools", hypCP.IpPools)
	d.Set("mgmt_ip_settings_override", hypCP.MgmtIpSettingsOverride)
	d.Set("modified", hypCP.Modified)
	d.Set("name", hypCP.Name)
	d.Set("path", hypCP.Path)
	d.Set("refresh_state", hypCP.RefreshState)
	d.Set("scopes_uri", hypCP.ScopesUri)
	d.Set("shared_storage_volumes", hypCP.SharedStorageVolumes)
	d.Set("state", hypCP.State)
	d.Set("state_reason", hypCP.StateReason)
	d.Set("status", hypCP.Status)
	d.Set("type", hypCP.Type)
	d.Set("uri", hypCP.Uri)
	d.Set("initial_scope_uris", hypCP.InitialScopeUris)
	return nil
}

func resourceHypervisorClusterProfileUpdate(d *schema.ResourceData, meta interface{}) error {
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
}

func resourceHypervisorClusterProfileDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteHypervisorClusterProfile(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
