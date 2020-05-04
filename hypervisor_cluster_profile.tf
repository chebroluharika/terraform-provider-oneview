provider "oneview" {
	ov_username = "Administrator"
	ov_password = "admin123"
	ov_endpoint = "https://10.50.9.31"
	ov_sslverify = false
	ov_apiversion = 1200
	ov_ifmatch = "*"
}

data "oneview_scope" "scope_obj" {
        name = "test_scope"
}


resource "oneview_hypervisor_cluster_profile" "HypervisorClusterProfile"{  
    "type"="HypervisorClusterProfileV3",
    "name"="Cluster7",
    "description"="",
    "hypervisor_type"="Vmware",
    "hypervisor_manager_uri"="/rest/hypervisor-managers/c6ef2d7a-e84a-4749-96e2-1bbc59543fce",
    "path"="DC2",
    "hypervisor_cluster_settings"={  
                                  "type"="Vmware",
                                  "drs_enabled"=true,
                                  "ha_enabled"=false,
                                  "multi_nic_v_motion"=false,
                                  "virtual_switch_type"="Standard"
                               },
    "hypervisor_host_profile_template"={  
    "server_profile_template_uri"="/rest/server-profile-templates/278cadfb-2e86-4a05-8932-972553518259",
     "host_prefix"="Test-Cluster-host"
     },
}
