provider "oneview" {
	ov_username = "Administrator"
	ov_password = "admin123"
	ov_endpoint = "10.50.9.3"
	ov_sslverify = false
	ov_apiversion = 1200
	ov_ifmatch = "*"
}

data "oneview_scope" "scope_obj" {
        name = "test_scope"
}


resource "oneview_hypervisor_cluster_profile" "HypervisorClusterProfile"{  
    "type":"HypervisorClusterProfileV3",
    "name":"Cluster5",
    "description":"",
    "hypervisorType":"Vmware",
    "hypervisorManagerUri":"/rest/hypervisor-managers/444f88c8-3a66-4f20-9b87-0c1da9542df8",
    "path":"DC2",
    "hypervisorClusterSettings":{  
                                  "type":"Vmware",
                                  "drsEnabled":true,
                                  "haEnabled":false,
                                  "multiNicVMotion":false,
                                  "virtualSwitchType":"Standard"
                               },
    "hypervisorHostProfileTemplate":{  
    "serverProfileTemplateUri":"/rest/server-profile-templates/cd2a2a32-8055-4ac9-9ec4-a10ecfab05d9",
         "deploymentPlan":{  
         "deploymentPlanUri":"/rest/os-deployment-build-plans/1180001",
         "serverPassword":"myPassword"
         },
     "hostprefix":"Test-Cluster-host"
     }
