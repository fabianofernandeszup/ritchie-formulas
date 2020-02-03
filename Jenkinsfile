def TERRAFORM_AWS_ACCESS_KEY_ID = env["DOCKER_AWS_ACCESS_KEY_ID_PRODUCTION_MARTE"]
def TERRAFORM_AWS_SECRET_ACCESS_KEY = env["DOCKER_AWS_SECRET_ACCESS_KEY_PRODUCTION_MARTE"]

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
                        string(credentialsId: TERRAFORM_AWS_ACCESS_KEY_ID, variable: 'aws_access_key_id_unveil'),
                        string(credentialsId: TERRAFORM_AWS_SECRET_ACCESS_KEY, variable: 'aws_secret_access_key_unveil'),
                      ]) {
                        try{
                            sh "AWS_ACCESS_KEY_ID=${aws_access_key_id_unveil} AWS_SECRET_ACCESS_KEY=${aws_secret_access_key_unveil} echo Here"
                        } catch (Error error){
                            echo "Error while building or pushing to s3"
                        }
                      }
                }
            }
      }
    }
    post {
        failure {
            echo "Build failed"
        }
    }
}

