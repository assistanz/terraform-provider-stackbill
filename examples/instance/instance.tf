
data "stackbill_compute_offering_list" "all" {
  zone_uuid = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
  uuid = "c674ac49-32cd-4aae-96f8-25458bded6ad"
}

# Returns all coffees
output "all_compute_offerings" {
  value = data.stackbill_compute_offering_list.all.computeofferings
}

resource "stackbill_instance" "my-server" {
  compute_offering_uuid = "2f7b1b02-078f-42fb-aece-aa21bfa89e1a"
  cpu_core              = "3"
  # disk_size             = 0
  hypervisor_name       = "string"
  # memory                = "0"
  name                  = "AzTestingVmFour"
  network_uuid          = "f215bcd2-8109-484d-91b3-20fc1cc8293b"
  # security_group_name   = "string"
  # ssh_key_name          = "string"
  # storage_offering_uuid = "string"
  template_uuid         = "e39172b4-715b-44ac-ada2-7f97f733eecc"
  zone_uuid             = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
}

# resource "stackbill_instance_update_name" "update_name" {
#     uuid = "a82c041a-b660-4446-b570-0b387cc95a4c"
#     name = "AzTestingVmFourupdate"
# }

# resource "stackbill_instance_resize" "update_name" {
#     compute_offering_uuid = "2f7b1b02-078f-42fb-aece-aa21bfa89e1a"
#     uuid = "ec3bf2f6-fc16-4157-ae75-71452b6d8d30"
#     cpu_core = "3"
#     memory = "0"
# }

# resource "stackbill_instance_actions" "actions" {
#     uuid = "a82c041a-b660-4446-b570-0b387cc95a4c"
#     action = "Stop"
# }