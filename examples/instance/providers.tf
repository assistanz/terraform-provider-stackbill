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
  api_key    = "EFq1o38cXDDq9mn850t-pDXuLq33ahwBFWcsgPFK9mTADA-HROCsfsIhjOuL"
  secret_key = "LKdDxtfboa18le8cJDe98XWKb7Le6C5WrRNSesxasHcg9vZmjA5tS437vYl-"
}
