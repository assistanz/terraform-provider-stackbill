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
  url    = "http://wolfapp.assistanz24x7.com/restapis"
  api_key    = "NqkTfkADuBFCNid2ypPViqq3M2vMOQddtjLyIQJ84LygYkppwfO1DZFCZX5H"
  secret_key = "H_eRcufZShygDM2ZV6g94gA9ecmgE9cKGTZ31VOkGXE72EiupuMnhuMZd0hs"
}
