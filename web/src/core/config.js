/**
 * 网站配置文件
 */

const config = {
  appName: 'Go Enrollapi',
  appLogo: '../assets/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用go-enrollapi`
      )
    )
   
  }
}

export default config
