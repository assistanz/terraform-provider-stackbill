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

resource "stackbill_network" "my-server" {
  description = "terraform test"
  name                  = "terraform volume teo"
  network_offering_uuid              = "19c28a5d-0237-450d-b47a-dbddc67aa0df"
  virtual_machine_uuid = "4d9150e6-64b5-451a-8851-0a4ea5142242"
  zone_uuid                = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
}
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

# resource "stackbill_volume_actions" "my-server" {
#   uuid = "9fe3dd14-bdb2-4c9a-8742-85ead82475cd"
#   action = "Attach"
#   instance_uuid  = "4d9150e6-64b5-451a-8851-0a4ea5142242"
# }

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
#   network_uuid          = "f215bcd2-8109-484d-91b3-20fc1cc8293b"
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
