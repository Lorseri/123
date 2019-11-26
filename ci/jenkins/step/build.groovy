timeout(time: 60, unit: 'MINUTES') {
    dir ("ci/scripts") {
        withCredentials([usernamePassword(credentialsId: "${params.JFROG_CREDENTIALS_ID}", usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
            if ("${env.BINRARY_VERSION}" == "gpu") {
                sh "export JFROG_ARTFACTORY_URL='${params.JFROG_ARTFACTORY_URL}' && export JFROG_USER_NAME='${USERNAME}' && export JFROG_PASSWORD='${PASSWORD}' && ./build.sh -t ${params.BUILD_TYPE} -o /opt/milvus -l -g -u -c"
            } else {
                sh "export JFROG_ARTFACTORY_URL='${params.JFROG_ARTFACTORY_URL}' && export JFROG_USER_NAME='${USERNAME}' && export JFROG_PASSWORD='${PASSWORD}' && ./build.sh -t ${params.BUILD_TYPE} -o /opt/milvus -l -m -u -c"
            }
        }
    }
}
