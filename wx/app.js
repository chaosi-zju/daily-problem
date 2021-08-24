// app.js
App({
  globalData: {
    problems: []
  },
  towxml: require('/towxml/index'),
  onLaunch() {
    // 登录
    wx.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      }
    })
    //显示转发分享
    wx.showShareMenu({
      menus: ['shareAppMessage', 'shareTimeline'],
      success(res) {
        util.toast(wx, '转发成功', 800)
      },
      fail(e) {
        util.toast(wx, '转发失败', 800)
      }
    })
  }
})