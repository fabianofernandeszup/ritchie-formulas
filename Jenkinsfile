def RITCHIE_AWS_ACCESS_KEY_ID = env["DOCKER_AWS_ACCESS_KEY_ID_PRODUCTION_MARTE"]
def RITCHIE_AWS_SECRET_ACCESS_KEY = env["DOCKER_AWS_SECRET_ACCESS_KEY_PRODUCTION_MARTE"]
def RITCHIE_AWS_REGION_PRODUCTION_MARTE = "sa-east-1"


pipeline{
    agent any
    stages
    {
      stage("Building formulas and sending them to s3"){
            environment {
                TERRAFORM_AWS_ACCESS_KEY_ID= "${TERRAFORM_AWS_ACCESS_KEY_ID}"
                TERRAFORM_AWS_SECRET_ACCESS_KEY= "${TERRAFORM_AWS_SECRET_ACCESS_KEY}"
            }
            steps {
                script{
                    withCredentials(
                      [
                        string(credentialsId: RITCHIE_AWS_ACCESS_KEY_ID, variable: 'aws_access_key_id_unveil'),
                        string(credentialsId: RITCHIE_AWS_SECRET_ACCESS_KEY, variable: 'aws_secret_access_key_unveil'),
                      ]) {
                        try{
                            sh "AWS_ACCESS_KEY_ID=${aws_access_key_id_unveil} AWS_SECRET_ACCESS_KEY=${aws_secret_access_key_unveil} AWS_DEFAULT_REGION=${RITCHIE_AWS_REGION_PRODUCTION_MARTE} make push-s3"
                        } catch (Error error){
                            echo "Error while building or pushing to s3"
                        }
                      }
                }
            }
      }
    }
    post {
        success {
            echo "Build and push successfully executed by Jenkins"
        }
        failure {
            echo "Build failed"
        }
    }
}

