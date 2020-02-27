pipeline{
    agent {
        dockerfile {
            filename 'Dockerfile'
            dir '.'
            args  '--privileged -u 0:0'
        }
    }


    stages
    {
        stage("Setting environment variables"){
          parallel {
              stage('master') {
                  when {
                      expression {
                          return branch_name =~ /^master/
                      }
                  }
                  steps {
                      script{
                        RITCHIE_AWS_ACCESS_KEY_ID = env["DOCKER_AWS_ACCESS_KEY_ID_PRODUCTION_MARTE"]
                        RITCHIE_AWS_SECRET_ACCESS_KEY = env["DOCKER_AWS_SECRET_ACCESS_KEY_PRODUCTION_MARTE"]
                        RITCHIE_AWS_REGION = "sa-east-1"
                        RITCHIE_AWS_BUCKET = "ritchie-cli-bucket152849730126474"
                        buildable = true
                      }
                  }
              }

              stage('qa') {
                  when {
                      expression {
                          return branch_name =~ /^qa/
                      }
                  }
                  steps {
                      script{
                        RITCHIE_AWS_ACCESS_KEY_ID = env["DOCKER_AWS_ACCESS_KEY_ID_QA_MARTE"]
                        RITCHIE_AWS_SECRET_ACCESS_KEY = env["DOCKER_AWS_SECRET_ACCESS_KEY_QA_MARTE"]
                        RITCHIE_AWS_REGION = "sa-east-1"
                        RITCHIE_AWS_BUCKET = "ritchie-cli-bucket234376412767550"
                        buildable = true
                      }
                  }
              }
          }
        }

        stage("Building formulas and sending them to s3"){
           when {
              expression {
                  buildable
              }
            }
            steps {
                script{

                    checkout scm

                    withCredentials(
                      [
                        string(credentialsId: RITCHIE_AWS_ACCESS_KEY_ID, variable: 'aws_access_key_id_unveil'),
                        string(credentialsId: RITCHIE_AWS_SECRET_ACCESS_KEY, variable: 'aws_secret_access_key_unveil'),
                      ]) {
                        try{
                            sh "AWS_ACCESS_KEY_ID=${aws_access_key_id_unveil} AWS_SECRET_ACCESS_KEY=${aws_secret_access_key_unveil} AWS_DEFAULT_REGION=${RITCHIE_AWS_REGION} RITCHIE_AWS_BUCKET=${RITCHIE_AWS_BUCKET} make push-s3"
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

