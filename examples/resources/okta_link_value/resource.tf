resource "okta_link_definition" "padishah" {
  primary_name           = "emperor"
  primary_title          = "Emperor"
  primary_description    = "Hereditary ruler of the Imperium and the Known Universe"
  associated_name        = "sardaukar"
  associated_title       = "Sardaukar"
  associated_description = "Elite military force member"
}

resource "okta_user" "emperor" {
  first_name = "Shaddam"
  last_name  = "Corrino IV"
  login      = "shaddam.corrino.iv@salusa-secundus.planet"
  email      = "shaddam.corrino.iv@salusa-secundus.planet"
}

resource "okta_user" "sardaukars" {
  count      = 5
  first_name = "Amrit"
  last_name  = "Sardaukar_${count.index}"
  login      = "amritsardaukar_${count.index}@salusa-secundus.planet"
  email      = "amritsardaukar_${count.index}@salusa-secundus.planet"
}

resource "okta_link_value" "example" {
  primary_name    = okta_link_definition.padishah.primary_name
  primary_user_id = okta_user.emperor.id
  associated_user_ids = [
    okta_user.sardaukars[0].id,
    okta_user.sardaukars[1].id,
    okta_user.sardaukars[2].id,
    okta_user.sardaukars[3].id,
    okta_user.sardaukars[4].id,
  ]
}
