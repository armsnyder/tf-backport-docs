# By doot
resource "foo_deploy_key" "example_1" {
  project = "example/deploying"
  title   = "Example deploy key"
}

# By blah
resource "foo_deploy_key" "example_2" {
  project = "example/deploying2"
  title   = "Example deploy key 2"
}
