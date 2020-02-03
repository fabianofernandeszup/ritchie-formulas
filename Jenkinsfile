pipeline{

    stages
    {
      stage("Building formulas and sending them to s3"){
            environment {
                TERRAFORM_AWS_ACCESS_KEY_ID= env["TERRAFORM_AWS_ACCESS_KEY_ID"]
                TERRAFORM_AWS_SECRET_ACCESS_KEY= env["TERRAFORM_AWS_SECRET_ACCESS_KEY"]
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
                        } catch {
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

