#!/bin/bash

for f in api/environments/v1alpha1/build.go \
          api/environments/v1alpha1/deployment.go \
          api/environments/v1alpha1/environment.go \
          api/environments/v1alpha1/package.go \
          api/environments/v1alpha1/serviceunit.go \
          api/events/v1alpha1/githubevent.go \
          api/networks/v1alpha1/domain.go \
          api/networks/v1alpha1/route.go \
          api/sources/v1alpha1/gitrepository.go; do
  sed -i 's|^// \*\/$|// */\n// +k8s:openapi-gen=true|' $f || \
  sed -i '0,/^package /s/^package /\/\/ +k8s:openapi-gen=true\n\npackage /' $f
done