podTemplate(yaml: """
apiVersion: v1
kind: Pod
spec:
  serviceAccountName: ecr
  containers:
  - name: bazel
    image: gcr.io/cloud-marketplace-containers/google/bazel@sha256:bea7cec14f05aea5a8a8ead0ddbf9a8ef66b4ad2a0f997c8d1d0217804044b8a
    command:
    - cat
    tty: true
"""
) {
    node(POD_LABEL) {
      checkout scm

      container("bazel") {
        sh "bazel run //:docker-password-test"
      }
    }
}
