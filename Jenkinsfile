//创建一个Pod的模板，label为jenkins-slave
podTemplate(label: 'jenkins-slave', cloud: 'kubernetes', containers: [
    containerTemplate(
        name: 'jnlp',
        image: "jenkins/jnlp-slave:latest"
    )
  ]
)
{
//引用jenkins-slave的pod模块来构建Jenkins-Slave的pod
node("jenkins-slave"){
      // 第一步
      stage('测试'){
    sh '''
            echo "hello world"
        '''
      }
  }
}