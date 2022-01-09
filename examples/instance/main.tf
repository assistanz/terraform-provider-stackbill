# This is required for Terraform 0.13+
terraform {
  required_providers {
    stackbill = {
      version = "~> 1.0.0"
      source  = "stackbill.com/assistanz/stackbill"
    }
  }
}

variable "coffee_name" {
  type    = string
  default = "Vagrante espresso"
}

provider "stackbill" {
  api_key    = "EFq1o38cXDDq9mn850t-pDXuLq33ahwBFWcsgPFK9mTADA-HROCsfsIhjOuL"
  secret_key = "LKdDxtfboa18le8cJDe98XWKb7Le6C5WrRNSesxasHcg9vZmjA5tS437vYl-"
}

# resource "stackbill_instance_reset_sshkey" "my-server" {
#   uuid       = "4d9150e6-64b5-451a-8851-0a4ea5142242"
#   ssh_key_id = "cdc0d62c-9f9c-41c3-9b8e-37ecd49ef57"
# }

# resource "stackbill_iso_list" "my-server" {
#   zone_id = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
# }
#
# resource "stackbill_instance_resize" "my-server" {
#   compute_offering_uuid = "2f7b1b02-078f-42fb-aece-aa21bfa89e1a"
#   uuid                  = "4d9150e6-64b5-451a-8851-0a4ea5142242"
#   cpu_core              = "3"
#   memory                = "0"
# }

data "stackbill_iso_list" "all" {
  zone_id = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
}

# Returns all coffees
output "all_compute_offeringss" {
  value = data.stackbill_iso_list.all.iso
}

# # Returns all coffees
# output "all_coffees" {
#   value = data.stackbill_instance_list.all.instances
# }


data "stackbill_sshkey_list" "all" {
}

# Returns all coffees
output "all_compute_offerings" {
  value = data.stackbill_sshkey_list.all.sshkeys
}

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
