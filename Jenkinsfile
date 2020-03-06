def RITCHIE_AWS_ACCESS_KEY_ID = env["DOCKER_AWS_ACCESS_KEY_ID_PRODUCTION_MARTE"]
def RITCHIE_AWS_SECRET_ACCESS_KEY = env["DOCKER_AWS_SECRET_ACCESS_KEY_PRODUCTION_MARTE"]
def RITCHIE_AWS_REGION_PRODUCTION_MARTE = "sa-east-1"

def jobNameParts = JOB_NAME.tokenize('/') as String[];
def githubDestinationRepo = jobNameParts[0];
def githubDestinationBranch = "marte"
def githubDestinationOrg = "martetech"

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
      stage("Building formulas and sending them to s3"){
           when {
              expression {
                  return branch_name =~ /^master/
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
                            sh "AWS_ACCESS_KEY_ID=${aws_access_key_id_unveil} AWS_SECRET_ACCESS_KEY=${aws_secret_access_key_unveil} AWS_DEFAULT_REGION=${RITCHIE_AWS_REGION_PRODUCTION_MARTE} make push-s3"
                        } catch (Error error){
                            echo "Error while building or pushing to s3"
                        }
                      }
                }
            }
      }
        stage("Sync with martetech repo") {
            when {
              branch 'marte'
            }
            steps {
                withCredentials([usernamePassword(credentialsId: 'github-ci-marte-zup', passwordVariable: 'git_passwd', usernameVariable: 'git_user')]) {
                    sh "git config --global user.name ${git_user}"
                    sh "git remote rm upstream || exit 0"
                    sh "git remote add upstream https://${git_user}:${git_passwd}@github.com/${githubDestinationOrg}/${githubDestinationRepo}.git"
                    sh "git remote -v"
                    sh "rm -rf vivo"
                    sh "git add . && git commit -m \"jenkins: rm unnecessary files\""
                    sh "git fetch upstream"
                    sh "git push -u upstream HEAD:${githubDestinationBranch} -f"
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

