# This is required for Terraform 0.13+
variable "coffee_name" {
  type    = string
  default = "Vagrante espresso"
}

# resource "stackbill_instance_reset_sshkey" "my-server" {
#   uuid       = "4d9150e6-64b5-451a-8851-0a4ea5142242"
#   ssh_key_uuid = "4d2546b2-2526-4b6a-b25b-9243d13fe0aa"
# }

# resource "stackbill_iso_list" "my-server" {
#   zone_id = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
# }
#
# resource "stackbill_instance_snapshot" "my-server" {
#   description = "terraform test"
#   name                  = "terraform volume"
#   virtual_machine_uuid              = "4d9150e6-64b5-451a-8851-0a4ea5142242"
#   zone_uuid                = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
# }

# resource "stackbill_network" "my-server" {
#   description = "terraform test"
#   name                  = "terraform volume teo"
#   network_offering_uuid              = "19c28a5d-0237-450d-b47a-dbddc67aa0df"
#   virtual_machine_uuid = "4d9150e6-64b5-451a-8851-0a4ea5142242"
#   zone_uuid                = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
# }
# resource "stackbill_network_attach" "my-server" {
#   network_uuid              = "80934451-21e7-4bb3-be39-30a8d14db2ac"
#   virutal_machine_uuid = "4d9150e6-64b5-451a-8851-0a4ea5142242"
#   action = "Add"
# }

# data "stackbill_security_group_list" "all" {
#   # zone_uuid = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
#   uuid = "a860fad9-d1d2-4394-883d-07d8a209a78b"
# }

# Returns all coffees
# output "all_compute_offeringss" {
#   value = data.stackbill_security_group_list.all.securitygroups
# }

# resource "stackbill_instance_snapshot" "my-server" {
#   name = "terraform-test-one"
#   description = "terraform-test-one"
#   virtual_machine_uuid  = "4d9150e6-64b5-451a-8851-0a4ea5142242"
#   zone_uuid = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
#   snapshot_memory = false
# }
# resource "stackbill_volume" "my-server" {
#   disk_size = 1
#   name = "terraform-test-three"
#   storage_offering_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
#   zone_uuid = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
# }

resource "stackbill_volume_actions" "my-server" {
  uuid = "4914ae3d-3f0d-4cf5-9910-b25c4977d3c3"
  action = "Attach"
  instance_uuid  = "49c4843d-9e4e-46ff-a004-518fa34efb2a"
}

# attach network to VM

# resource "stackbill_network_actions" "attach_network" {
#   network_uuid = "f385c6af-a900-415d-a131-44d1fc6ae4e3"
#   virutal_machine_uuid = "a17659f2-37b2-4973-ab52-f1ece7c846bb"
#   action = "Delete"
# }

resource "stackbill_network_actions" "attach_network" {
  network_uuid = "ab8d6f1a-7d6a-4ab6-b564-e625edfb796b"
  virutal_machine_uuid = "49c4843d-9e4e-46ff-a004-518fa34efb2a"
  action = "Delete"
}

# data "stackbill_volume_list" "all" {
#   zone_uuid = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
#   uuid = "b5e24f0e-0ac1-49b2-856a-e316a338ed13"
# }

# # Returns all coffees
# output "all_compute_offeringss" {
#   value = data.stackbill_volume_list.all.volumes
# }

# # Returns all coffees
# output "all_coffees" {
#   value = data.stackbill_instance_list.all.instances
# }


# data "stackbill_sshkey_list" "all" {
# }

# # Returns all coffees
# output "all_compute_offerings" {
#   value = data.stackbill_sshkey_list.all.sshkeys
# }

# output "coffee" {
#   value = {
#     for coffee in data.stackbill_instance_list.all.coffees :
#     coffee.id => coffee
#     if coffee.name == var.coffee_name
#   }
# }

# Returns all coffees
# output "all_coffees" {
#   value = data.hashicups_coffees.all.coffees
# }

# resource "stackbill_instance" "my-server" {
#   compute_offering_uuid = "2f7b1b02-078f-42fb-aece-aa21bfa89e1a"
#   cpu_core              = "3"
#   disk_size             = 0
#   hypervisor_name       = "string"
#   memory                = "0"
#   name                  = "AzTestingVmTwo"
#   network_uuid          = "f4a12b6e-fcde-46d1-a283-235b6d24a647"
#   security_group_name   = "string"
#   ssh_key_name          = "string"
#   storage_offering_uuid = "string"
#   template_uuid         = "e39172b4-715b-44ac-ada2-7f97f733eecc"
#   zone_uuid             = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
# }

# resource "stackbill_actions" "my-server" {
#   uuid   = "ffe1168f-569c-41f0-89b0-2f45e8103930"
#   action = "Start"
# }

# security_group_name
# ssh_key_name
# storage_offering_uuid
# network_uuid


# This is required for Terraform 0.13+
# terraform {
#   required_providers {
#     stackbill = {
#       version = "~> 1.0.0"
#       source  = "stackbill.com/assistanz/stackbill"
#     }
#   }
# }

# provider "stackbill" {
#   url    = "http://wolfapp.assistanz24x7.com/restapi"
#   api_key    = "NqkTfkADuBFCNid2ypPViqq3M2vMOQddtjLyIQJ84LygYkppwfO1DZFCZX5H"
#   secret_key = "H_eRcufZShygDM2ZV6g94gA9ecmgE9cKGTZ31VOkGXE72EiupuMnhuMZd0hs"
# }


#creating stackbill_instance
# resource "stackbill_instance" "my-server" {
#   compute_offering_uuid = "c674ac49-32cd-4aae-96f8-25458bded6ad"
#   cpu_core              = "3"
#   disk_size             = 0
#   hypervisor_name       = "XenServer"
#   # memory                = "0"
#   name                  = "AzTestingVmOne"
#   network_uuid          = "f385c6af-a900-415d-a131-44d1fc6ae4e3"

#   # security_group_name   = "string"
#   # ssh_key_name          = "string"
#   storage_offering_uuid = "45ac1fe4-9e40-47d2-81a3-7dd98811bf27"
#   template_uuid         = "670dc0de-8958-449e-a217-262f666c2fa6"
#   zone_uuid             = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
# }

# #updating VM status
# resource "stackbill_instance_actions" "status" {
#   uuid  = stackbill_instance.my-server.id
#   action = "Stop"
# }

#updating display name
resource "stackbill_instance_update_name" "update_name" {
  uuid  = "a17659f2-37b2-4973-ab52-f1ece7c846bb"
  name = "pra-vm3"
}

#resizing VM
resource "stackbill_instance_resize" "resize" {
  uuid  = "a17659f2-37b2-4973-ab52-f1ece7c846bb"
  compute_offering_uuid = "597b163b-18f0-40b1-82c0-e499bff5a2d5"
  cpu_core = "1"
  memory = "1024"
}

resource "stackbill_instance_reset_sshkey" "reset" {
  uuid = "a17659f2-37b2-4973-ab52-f1ece7c846bb"
  ssh_key_uuid = "a6903126-ff89-4878-971f-fba79ab4c8bc"
}

# # creating network
# resource "stackbill_network" "TerraformNetwork" {
#   description = "network created to test terraform"
#   name = "TerraformNetwork"
#   # virtual_machine_uuid = stackbill_instance.my-server.id
#   network_offering_uuid = "19c28a5d-0237-450d-b47a-dbddc67aa0df"
#   zone_uuid = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
#   is_public = "true"
# }

# # #creating new volume
# resource "stackbill_volume" "newvolume" {
#   disk_size = 1024
#   name = "ExtraVolumeForTerrformVM"
#   storage_offering_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
#   zone_uuid = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
# }


# # #creating snapshot
# resource "stackbill_instance_snapshot" "snapshot" {
#   name = "terraform_VM_Snapshot"
#   description = "snapshot of VM created with terraform"
#   virtual_machine_uuid = "c7e8ba23-0c79-47ab-9fd3-5158ba275c64"
#   zone_uuid = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
#   snapshot_memory = "0"
# }

#updating display name
# resource "stackbill_instance_update_name" "update_name" {
#   uuid  = "c7e8ba23-0c79-47ab-9fd3-5158ba275c64"
#   name = "TerraformTestVMOne"
# }

#resizing VM
# resource "stackbill_instance_resize" "resize" {
#   uuid  = "c7e8ba23-0c79-47ab-9fd3-5158ba275c64"
#   compute_offering_uuid = "c674ac49-32cd-4aae-96f8-25458bded6ad"
#   cpu_core = "5"
#   memory = "512"
# }