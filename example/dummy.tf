provider "ipam" {}

resource "ipam_dummy" "d" {
  path = "./thefile.txt"
}