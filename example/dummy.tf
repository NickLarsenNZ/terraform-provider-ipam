provider "ipam" {}

resource "ipam_dummy" "d" {
  path       = "./thefile.txt"
  firstline  = "alpha"
  secondline = "bravo"
}