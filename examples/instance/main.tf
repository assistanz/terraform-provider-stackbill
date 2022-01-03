# This is required for Terraform 0.13+
terraform {
  required_providers {
    stackbill = {
      version = "~> 1.0.0"
      source  = "stackbill.com/assistanz/stackbill"
    }
  }
}

provider "stackbill" {
  api_key    = "4plBJ1KEUhaAq9oaZQ7rxxRkM1DRC4M0kU35NeB6W1NPkNfdl7bmVgxHimPA"
  secret_key = "eq63VZ9guuWWoXSU1GNr7Hg94bA2avA7dWWT3avA1uRIcUV6IaN-5s3y8l_e"
}

# resource "stackbill_instance" "my-server" {
#   compute_offering_uuid = "c674ac49-32cd-4aae-96f8-25458bded6ad"
#   cpu_core              = "string"
#   disk_size             = 0
#   hypervisor_name       = "string"
#   memory                = "string"
#   name                  = "AzTestingVmOne"
#   network_uuid          = "f215bcd2-8109-484d-91b3-20fc1cc8293b"
#   security_group_name   = "string"
#   ssh_key_name          = "string"
#   storage_offering_uuid = "string"
#   template_uuid         = "e39172b4-715b-44ac-ada2-7f97f733eecc"
#   zone_uuid             = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
# }

resource "stackbill_actions" "my-server" {
  uuid   = "ffe1168f-569c-41f0-89b0-2f45e8103930"
  action = "Start"
}

# security_group_name
# ssh_key_name
# storage_offering_uuid
# network_uuid
