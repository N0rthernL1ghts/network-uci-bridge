group "default" {
  targets = [
    "1_0_0"
  ]
}

target "build-dockerfile" {
  dockerfile = "Dockerfile"
}

target "build-platforms" {
  platforms = ["linux/amd64", "linux/armhf", "linux/aarch64"]
}

target "build-common" {
  pull = true
}

variable "REGISTRY_CACHE" {
  default = "docker.io/nlss/uci-bridge-cache"
}

######################
# Define the functions
######################

# Get the cache-from configuration
function "get-cache-from" {
  params = [version]
  result = [
    "type=registry,ref=${REGISTRY_CACHE}:${sha1("${version}-${BAKE_LOCAL_PLATFORM}")}"
  ]
}

# Get the cache-to configuration
function "get-cache-to" {
  params = [version]
  result = [
    "type=registry,mode=max,ref=${REGISTRY_CACHE}:${sha1("${version}-${BAKE_LOCAL_PLATFORM}")}"
  ]
}

# Get list of image tags and registries
# Takes a version and a list of extra versions to tag
# eg. get-tags("1.0.0", ["1", "1.0", "latest"])
function "get-tags" {
  params = [version, extra_versions]
  result = concat(
    [
      "docker.io/nlss/uci-bridge:${version}",
      "ghcr.io/n0rthernl1ghts/uci-bridge:${version}"
    ],
    flatten([
      for extra_version in extra_versions : [
        "docker.io/nlss/uci-bridge:${extra_version}",
        "ghcr.io/n0rthernl1ghts/uci-bridge:${extra_version}"
      ]
    ])
  )
}

##########################
# Define the build targets
##########################

target "1_0_0" {
  inherits   = ["build-dockerfile", "build-platforms", "build-common"]
  cache-from = get-cache-from("1.0.0")
  cache-to   = get-cache-to("1.0.0")
  tags       = get-tags("1.0.0", ["1", "1.0", "latest"])
}